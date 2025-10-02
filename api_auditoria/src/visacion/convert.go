package visacion

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
)

func TransformMercanciasDespachadas(src *model.MercanciasDespachadasTopic) *model.MercanciasDespachada {
	var topic model.MercanciasDespachadasT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}

	// Convertir peso de string a float64

	return &model.MercanciasDespachada{
		ID:                topic.ID,
		Cantidad:          topic.Cantidad,
		FechaDespacho:     utils.ToFormattedDateTime(topic.FechaDespacho),
		Peso:              func(f float64) *float64 { return &f }(utils.DecodeBinaryDecimal(topic.Peso, 2)),
		IDDetallePapeleta: topic.Despachado,
		IDVisaje:          topic.Visaje,
		Volumen:           func(f float64) *float64 { return &f }(utils.DecodeBinaryDecimal(topic.Volumen, 2)),
		Bulto:             nil,
		UsuarioCR:         topic.UsuarioCR,
		Evento:            utils.MapOperation(src.Op),
		FechaEvento:       utils.ToFormattedDateTime(src.TsMS),
	}
}

func TransformVisacion(src *model.VisacionTopic, visacion *model.Visacion) *model.Visacion {
	var topic model.VisacionT
	if src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	var fechaEvento *string
	if src.Op == "r" {
		fechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTime(src.TsMS)
	}
	var mercancias *[]model.MercanciasDespachada
	if visacion != nil && visacion.Mercacias != nil {
		copiedMercancias := make([]model.MercanciasDespachada, len(*visacion.Mercacias))
		copy(copiedMercancias, *visacion.Mercacias)
		mercancias = &copiedMercancias
	}
	return &model.Visacion{
		ID:          topic.ID,
		Aux:         topic.Aux,
		Doc:         topic.Doc,
		Nro:         topic.Nro,
		Aga:         topic.Aga,
		NroPapeleta: topic.NroPapeleta,

		FechaRetiro:   utils.ToFormattedDate(topic.FechaRetiro),
		FechaCR:       utils.ToFormattedDateTime(topic.FechaCR),
		CorrelativoPE: topic.CorrelativoPE,
		FechaDoc:      utils.ToFormattedDate(topic.FechaDoc),
		Evento:        utils.MapOperation(src.Op),
		FechaEvento:   fechaEvento,
		Mercacias:     mercancias,
	}
}
