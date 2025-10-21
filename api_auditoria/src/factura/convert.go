package factura

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"fmt"
)

func ConvertToFactura(src *model.FacturaTopic, factura *model.Factura) *model.Factura {
	var detalles *[]model.FacturaDetalle
	var topic model.FacturaT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	var fechaEvento *string
	if src.Op == "r" || src.Op == "c" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(src.TsMS)
	}
	if factura != nil && factura.Detalles != nil {
		fmt.Printf("Reemplazando detalles de la factura existente: %d\n", len(*factura.Detalles))
		copiedDetalles := make([]model.FacturaDetalle, len(*factura.Detalles))
		copy(copiedDetalles, *factura.Detalles)
		detalles = &copiedDetalles
	}

	return &model.Factura{
		ID:                       topic.ID,
		Folio:                    topic.Folio,
		RutCliente:               topic.RutCliente,
		IDFolio:                  topic.IDFolio,
		ManifiestoNroMftoInterno: topic.ManifiestoNroMftoInterno,
		Origen:                   topic.Origen,
		Moneda:                   topic.Moneda,
		TipoVenta:                topic.TipoVenta,
		FechaCambio:              utils.ToFormattedDate(topic.FechaCambio),
		ValorCambio:              utils.DecodeBinaryDecimal(topic.ValorCambio, 3),
		UsuarioCR:                topic.UsuarioCR,
		UsuarioUp:                topic.UsuarioUp,
		FechaCR:                  utils.ToFormattedDateTime(topic.FechaCR),
		FechaUp:                  utils.ToFormattedDateTime(topic.FechaUp),
		Comentario:               topic.Comentario,
		Empresa:                  topic.Empresa,
		Estado:                   topic.Estado,
		TotalNeto:                utils.ToInt(topic.TotalNeto),
		Aga:                      topic.Aga,
		NroIngRecaudacion:        topic.NroIngRecaudacion,
		NroBl:                    topic.NroBl,
		FacturaSap:               topic.FacturaSap,
		CorreoDespacho:           topic.CorreoDespacho,
		Evento:                   utils.MapOperation(src.Op),
		FechaEvento:              fechaEvento,
		Tipo:                     topic.Tipo,
		NroOportunidad:           topic.NroOportunidad,
		Forw:                     topic.Forw,
		OrdenCompra:              topic.OrdenCompra,
		Despacho:                 topic.Despacho,
		Mandante:                 topic.Mandante,
		Zona:                     topic.Zona,
		Detalles:                 detalles,
	}
}

func ConvertToFacturaDetalle(
	src *model.FacturaDetalleTopic,
) *model.FacturaDetalle {
	var topic model.FacturaDetalleT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	return &model.FacturaDetalle{
		ID:              topic.ID,
		Cantidad:        topic.Cantidad,
		Descuento:       utils.DecodeBinaryDecimal(topic.Descuento, 2),
		Detalle:         topic.Detalle,
		Dias:            topic.Dias,
		FechaCR:         utils.ToFormattedDateTime(topic.FechaCR),
		GlosaFactura:    topic.GlosaFactura,
		Imo:             topic.Imo,
		Moneda:          topic.Moneda,
		Neto:            utils.DecodeBinaryDecimal(topic.Neto, 2),
		Peso:            utils.DecodeBinaryDecimal(topic.Peso, 2),
		Recargo:         utils.DecodeBinaryDecimal(topic.Recargo, 2),
		Tarifa:          topic.Tarifa,
		Unidad:          topic.Unidad,
		UsuarioCR:       topic.UsuarioCR,
		ValorUnitario:   utils.DecodeBinaryDecimal(topic.ValorUnitario, 2),
		Volumen:         utils.DecodeBinaryDecimal(topic.Volumen, 2),
		IDFactura:       topic.IDFactura,
		NroPapeleta:     topic.NroPapeleta,
		Servicio:        topic.Servicio,
		Visaje:          topic.Visaje,
		NetoPeso:        utils.DecodeBinaryDecimal(topic.NetoPeso, 2),
		MotivoDescuento: topic.MotivoDescuento,
		MotivoRecargo:   topic.MotivoRecargo,
		PapeletaExpo:    topic.PapeletaExpo,
		VisajeVa:        topic.VisajeVa,
		DiasReales:      topic.DiasReales,
		Evento:          func(s string) *string { return &s }(utils.MapOperation(src.Op)),
		FechaEvento:     utils.ToFormattedDateTimeEvento(src.TsMS),
	}
}
