package expo

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"fmt"
)

func ConvertToPapeletaExpo(src *model.PapeletaExpoTopic, papeleta *model.PapeletaExpo) *model.PapeletaExpo {
	var detalles *[]model.PapeletaExpoDetalle
	if papeleta != nil && papeleta.Detalles != nil {
		fmt.Printf("Reemplazando detalles de la factura existente: %d\n", len(*papeleta.Detalles))
		copiedDetalles := make([]model.PapeletaExpoDetalle, len(*papeleta.Detalles))
		copy(copiedDetalles, *papeleta.Detalles)
		detalles = &copiedDetalles
	}
	var topic model.PapeletaExpoT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	var fechaEvento *string
	if src.Op == "r" || src.Op == "c" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(&src.TsMS)
	}
	return &model.PapeletaExpo{
		ID:                topic.ID,
		Empresa:           topic.Empresa,
		UsuarioCR:         topic.UsuarioCR,
		FechaCR:           utils.ToFormattedDateTime(topic.FechaCR),
		UsuarioUp:         topic.UsuarioUp,
		FechaUp:           utils.ToFormattedDateTime(topic.FechaUp),
		NroPapeleta:       topic.NroPapeleta,
		Operacion:         topic.Operacion,
		Zona:              topic.Zona,
		Documento:         topic.Documento,
		NroDoc:            topic.NroDoc,
		Emisor:            topic.Emisor,
		FechaDoc:          utils.ToFormattedDate(topic.FechaDoc),
		FechaRecepcion:    utils.ToFormattedDate(topic.FechaRecepcion),
		Booking:           topic.Booking,
		DeclaracionAduana: topic.DeclaracionAduana,
		NroDocAduana:      topic.NroDocAduana,
		FechaDocAduana:    utils.ToFormattedDate(topic.FechaDocAduana),
		Producto:          topic.Producto,
		Codigo:            topic.Codigo,
		Lote:              topic.Lote,
		Cosecha:           topic.Cosecha,
		CodigoFert:        topic.CodigoFert,
		Observaciones:     topic.Observaciones,
		IngresoExpo:       topic.IngresoExpo,
		Estado:            topic.Estado,
		Forwarder:         topic.Forwarder,
		MotivoAnulacion:   topic.MotivoAnulacion,
		Aga:               topic.Aga,
		Marcas:            topic.Marcas,
		PtoDesmb:          topic.PtoDesmb,
		NroViaje:          topic.NroViaje,
		Nave:              topic.Nave,
		Sku:               topic.Sku,
		Liberada:          topic.Liberada,
		MotivoLiberacion:  topic.MotivoLiberacion,
		IDRecepcion:       topic.IDRecepcion,
		Evento: utils.MapOperation(func(s *string) string {
			if s != nil {
				return *s
			}
			return ""
		}(&src.Op)),

		FechaEvento: fechaEvento,
		Detalles:    detalles,
	}
}

func ConvertToPapeletaExpoDetalle(
	src *model.PapeletaExpoDetalleTopic,
) *model.PapeletaExpoDetalle {
	var topic model.PapeletaExpoDetalleT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	return &model.PapeletaExpoDetalle{
		ID:         topic.ID,
		IDPapeleta: topic.IDPapeleta,
		Cantidad:   topic.Cantidad,
		Bulto:      topic.Bulto,
		Marca:      topic.Marca,
		Peso:       utils.DecodeBinaryDecimal(topic.Peso, 2),
		Volumen:    utils.DecodeBinaryDecimal(topic.Volumen, 2),
		Estado:     topic.Estado,
		Contenedor: topic.Contenedor,
		Sello:      topic.Sello,
		Chassis:    topic.Chassis,
		Situacion:  topic.Situacion,

		FechaEvento: utils.ToFormattedDateTimeEvento(src.TsMS),
		Evento: utils.MapOperation(func(s *string) string {
			if s != nil {
				return *s
			} else {
				return ""
			}
		}(&src.Op)),
	}
}
