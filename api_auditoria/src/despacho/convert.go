package despacho

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
)

func TransformDespacho(t *model.DespachoTopic, visacion *model.Despacho) *model.Despacho {
	var topic model.DespachoT
	if t.Op == "d" {
		topic = *t.Before
	} else {
		topic = *t.After
	}
	var fechaEvento *string
	if t.Op == "r" || t.Op == "c" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(t.TsMS)
	}
	return &model.Despacho{
		ID:                    topic.ID,
		Visacion:              topic.Visacion,
		UsuarioCR:             topic.UsuarioCR,
		Empresa:               topic.Empresa,
		FechaCR:               utils.ToFormattedDateTime(topic.FechaCR),
		Patente:               topic.Patente,
		RutChofer:             topic.RutChofer,
		Autorizado:            topic.Autorizado,
		UsuarioAutoriza:       topic.UsuarioAutoriza,
		FechaAutorizacion:     utils.ToFormattedDateTime(topic.FechaAutorizacion),
		Correlativo:           topic.Correlativo,
		Estado:                topic.Estado,
		FechaUp:               utils.ToFormattedDateTime(topic.FechaUp),
		NotivoAnulacion:       topic.NotivoAnulacion,
		PapeletaExpo:          topic.PapeletaExpo,
		Contenedor:            topic.Contenedor,
		GuiaDespacho:          topic.GuiaDespacho,
		TipoDespacho:          topic.TipoDespacho,
		MotivoAnulacionSalida: topic.MotivoAnulacionSalida,
		UsuarioAnulaSalida:    topic.UsuarioAnulaSalida,
		FechaAnulacionSalida:  topic.FechaAnulacionSalida,
		Evento:                utils.MapOperation(t.Op),
		FechaEvento:           fechaEvento,
	}
}
