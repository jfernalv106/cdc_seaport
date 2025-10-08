package notacredito

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
)

func ConvertToNotaCredito(src *model.NotaCreditoTopic, nc *model.NotaCredito) *model.NotaCredito {

	var servicios *[]model.NotaCreditoServicio
	if nc != nil && nc.Servicios != nil {
		copiedServicios := make([]model.NotaCreditoServicio, len(*nc.Servicios))
		copy(copiedServicios, *nc.Servicios)
		servicios = &copiedServicios
	}

	var topic model.NotaCreditoT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	var fechaEvento *string
	if *src.Op == "r" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(src.TsMS)
	}
	return &model.NotaCredito{
		IDNroNc:       topic.IDNroNc,
		IDFolNc:       topic.IDFolNc,
		IDFactura:     topic.IDFactura,
		IDFolio:       topic.IDFolio,
		FechaCR:       utils.ToFormattedDateTime(topic.FechaCR),
		UsuarioCR:     topic.UsuarioCR,
		Observaciones: topic.Observaciones,
		ValorNeto:     utils.ParseToFloat(topic.ValorNeto),
		Tipo:          topic.Tipo,
		Empresa:       topic.Empresa,
		Traspasada:    topic.Traspasada,
		IDSap:         topic.IDSap,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   fechaEvento,
		Servicios:     servicios,
	}
}

func ConvertToNotaCreditoServicio(src *model.NotaCreditoServTopic) *model.NotaCreditoServicio {
	var topic model.NotaCreditoServT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	return &model.NotaCreditoServicio{
		ID:            topic.ID,
		IDNroNc:       topic.IDNroNc,
		CodServ:       topic.CodServ,
		ValorAplicado: func(v float64) *float64 { return &v }(utils.DecodeBinaryDecimal(topic.ValorAplicado, 2)),
		Visaje:        topic.Visaje,
		DfID:          topic.DfID,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTime(src.TsMS),
	}
}
