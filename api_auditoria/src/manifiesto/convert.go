package manifiesto

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
)

func ConvertManifiestoTopicToManifiesto(topic *model.ManifiestoTopic) *model.Manifiesto {

	var sourceManifiesto model.ManifiestoT
	if topic.Op == "d" {
		sourceManifiesto = *topic.Before
	} else {
		sourceManifiesto = *topic.After
	}
	var fechaEvento *string
	if topic.Op == "r" || topic.Op == "c" {
		fechaEvento = utils.ToFormattedDateTime(sourceManifiesto.FechaCR)
	} else {
		fechaEvento = utils.ToFormattedDateTimeEvento(topic.TsMS)
	}
	return &model.Manifiesto{
		NroMftoInterno:        sourceManifiesto.NroMftoInterno,
		NroMfto:               sourceManifiesto.NroMfto,
		Nave:                  sourceManifiesto.Nave,
		Viaje:                 sourceManifiesto.Viaje,
		Servicio:              sourceManifiesto.Servicio,
		Agencia:               sourceManifiesto.Agencia,
		FechaMfto:             utils.ToFormattedDateTime(sourceManifiesto.FechaMfto),
		Estado:                sourceManifiesto.Estado,
		InscripcionDesde:      utils.ToFormattedDateTime(sourceManifiesto.InscripcionDesde),
		InscripcionHasta:      utils.ToFormattedDateTime(sourceManifiesto.InscripcionHasta),
		TrasladoPuerto:        utils.ToFormattedDateTime(sourceManifiesto.TrasladoPuerto),
		TipoMfto:              sourceManifiesto.TipoMfto,
		TipoAccion:            sourceManifiesto.TipoAccion,
		Version:               sourceManifiesto.Version,
		CondCarga:             sourceManifiesto.CondCarga,
		SitioAtraque:          sourceManifiesto.SitioAtraque,
		Almacen:               sourceManifiesto.Almacen,
		UsuarioCR:             sourceManifiesto.UsuarioCR,
		FechaCR:               utils.ToFormattedDateTime(sourceManifiesto.FechaCR),
		UsuarioUp:             sourceManifiesto.UsuarioUp,
		FechaUp:               utils.ToFormattedDateTime(sourceManifiesto.FechaUp),
		FechaUltLec:           utils.ToFormattedDateTime(sourceManifiesto.FechaUltLec),
		TipoCarga:             sourceManifiesto.TipoCarga,
		MotivoAnulacion:       sourceManifiesto.MotivoAnulacion,
		FechaTraspaso:         utils.ToFormattedDateTime(sourceManifiesto.FechaTraspaso),
		Zarpe:                 utils.ToFormattedDateTime(sourceManifiesto.Zarpe),
		NroMftoInternoEmpresa: sourceManifiesto.NroMftoInternoEmpresa,
		ArriboEfectivo:        utils.ToFormattedDateTime(sourceManifiesto.ArriboEfectivo),
		Evento:                utils.MapOperation(topic.Op),
		FechaEvento:           fechaEvento,
	}
}
