package papeleta

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
)

func ConvertToPapeletaRecepcion(src *model.PapeletaRecepcionTopic, papeleta *model.PapeletaRecepcion) *model.PapeletaRecepcion {
	var detalles *[]model.PapeletaRecepcionDetalle
	if papeleta != nil && papeleta.Detalles != nil {
		copiedDetalles := make([]model.PapeletaRecepcionDetalle, len(*papeleta.Detalles))
		copy(copiedDetalles, *papeleta.Detalles)
		detalles = &copiedDetalles
	}
	var topic model.PapeletaRecepcionT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	var fechaEvento *string
	if src.Op == "r" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(src.TsMS)
	}
	return &model.PapeletaRecepcion{
		NroPapeleta:              topic.NroPapeleta,
		Etiqueta:                 topic.Etiqueta,
		Operacion:                topic.Operacion,
		Zona:                     topic.Zona,
		NroBl:                    topic.NroBl,
		PtoOrigen:                topic.PtoOrigen,
		PtoDescarga:              topic.PtoDescarga,
		PtoDestino:               topic.PtoDestino,
		Consiganatario:           topic.Consiganatario,
		DireccionConsignatario:   topic.DireccionConsignatario,
		RutConsignatario:         topic.RutConsignatario,
		GuardaAlmacen:            topic.GuardaAlmacen,
		ZonaAlmacenaje:           utils.ToStr(topic.ZonaAlmacenaje),
		Ubicacion:                utils.ToStr(topic.Ubicacion),
		InicioAlmacenaje:         utils.ToFormattedDateTime(topic.InicioAlmacenaje),
		ManifiestoNroMftoInterno: topic.ManifiestoNroMftoInterno,
		Forwarder:                topic.Forwarder,
		Estado:                   topic.Estado,
		TipoCarga:                topic.TipoCarga,
		UsuarioCR:                topic.UsuarioCR,
		FechaCR:                  utils.ToFormattedDateTime(topic.FechaCR),
		FechaUP:                  utils.ToFormattedDateTime(topic.FechaUp),
		FechaDescon:              utils.ToFormattedDateTime(topic.FechaDescon),
		Tipo:                     topic.Tipo,
		PrePapeleta:              topic.PrePapeleta,
		PesoManifestado:          utils.ToInt(topic.PesoManifestado),
		Noty:                     topic.Noty,
		NroPapeletaEmpresa:       topic.NroPapeletaEmpresa,
		Aga:                      topic.Aga,
		RutForw:                  topic.RutForw,
		Liberada:                 topic.Liberada,
		MotivoLiberacion:         topic.MotivoLiberacion,
		VAnticipada:              topic.VAnticipada,
		Evento:                   utils.MapOperation(src.Op),
		FechaEvento:              fechaEvento,
		Detalles:                 detalles,
	}
}
func ConvertToPapeletaRecepcionDetalle(
	src *model.PapeletaRecepcionDetalleTopic,
) *model.PapeletaRecepcionDetalle {
	var topic model.PapeletaRecepcionDetalleT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	return &model.PapeletaRecepcionDetalle{
		ID:            topic.ID,
		CantidadItem:  topic.CantidadItem,
		Contenedor:    topic.Contenedor,
		Estado:        topic.Estado,
		IDItem:        topic.IDItem,
		Marcas:        topic.Marcas,
		Observaciones: topic.Observaciones,
		Peso:          utils.DecodeBinaryDecimal(topic.Peso, 2),
		Sellos:        topic.Sellos,
		Situacion:     topic.Situacion,
		Volumen:       utils.DecodeBinaryDecimal(topic.Volumen, 2),
		NroPapeleta:   topic.NroPapeleta,
		TipoBulto:     topic.TipoBulto,
		Chassis:       topic.Chassis,
		Identificador: topic.Identificador,
		Evento:        utils.MapOperation(src.Op),
		FechaEvento:   utils.ToFormattedDateTime(src.TsMS),
	}
}
