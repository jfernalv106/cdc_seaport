package bl

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ================== BL PRINCIPAL ==================
func ConvertToBL(src *model.BlTopic, base *model.BL) *model.BL {
	bl := cloneBase(base)

	var topic model.BlT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}

	bl.ID = topic.ID
	bl.NroBl = topic.NroBl
	bl.ManifiestoNroMftoInterno = topic.ManifiestoNroMftoInterno
	bl.TipoAccion = topic.TipoAccion
	bl.TipoBl = topic.TipoBl
	bl.FechaRecepcion = utils.ToFormattedDateTime(topic.FechaRecepcion)
	bl.Servicio = topic.Servicio
	bl.TipoServicio = topic.TipoServicio
	bl.CondTransporte = topic.CondTransporte
	bl.TotalBultos = topic.TotalBultos

	bl.TotalPeso = utils.DecodeBinaryDecimal(topic.TotalPeso, 3)
	bl.UnidadPeso = topic.UnidadPeso
	bl.TotalVolumen = utils.ParseToFloat(topic.TotalVolumen)
	bl.UnidadVolumen = topic.UnidadVolumen
	bl.TotalItem = topic.TotalItem
	bl.UsuarioCR = topic.UsuarioCR
	bl.FechaCR = utils.ToFormattedDateTime(topic.FechaCR)
	bl.UsuarioUp = topic.UsuarioUp
	bl.FechaUp = utils.ToFormattedDateTime(topic.FechaUp)
	bl.BlAforo = topic.BlAforo
	bl.BlSag = topic.BlSag
	bl.NroBlOriginal = topic.NroBlOriginal
	bl.MotivoAnulacion = topic.MotivoAnulacion
	bl.Estado = topic.Estado
	bl.Etiqueta = topic.Etiqueta
	bl.Aforo = topic.Aforo
	bl.Evento = utils.MapOperation(*src.Op)
	if *src.Op == "r" || *src.Op == "c" {
		bl.FechaEvento = utils.ToFormattedDateTime(topic.FechaCR)
	} else {
		bl.FechaEvento = utils.ToFormattedDateTimeEvento(src.TsMS)
	}

	return bl
}

// ================== FECHAS ==================
func ConvertToBLFecha(src *model.BlFechaTopic) *model.BlFecha {
	var topic model.BlFechaT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	fecha := model.BlFecha{
		ID:            topic.ID,
		Nombre:        topic.Nombre,
		Valor:         utils.ToFormattedDateTime(topic.Valor),
		BlNroBl:       topic.BlNroBl,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTime(src.TsMS),
		FechaTraspaso: utils.ToFormattedDateTime(topic.FechaTraspaso),
	}
	return &fecha
}

// ================== FLETES ==================
func ConvertToBLFlete(src *model.BlFleteTopic) *model.BlFlete {
	var topic model.BlFleteT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	flete := model.BlFlete{
		Descripcion:   topic.Descripcion,
		Tipo:          topic.Tipo,
		BlNroBl:       topic.BlNroBl,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso: utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &flete
}

// ================== ITEMS ==================
func ConvertToBLItem(src *model.BlItemTopic) *model.BlItem {
	var topic model.BlItemT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	item := model.BlItem{
		ID:             topic.ID,
		BlNroBl:        topic.BlNroBl,
		NroItem:        topic.NroItem,
		Marcas:         topic.Marcas,
		CargaPeligrosa: topic.CargaPeligrosa,
		TipoBulto:      topic.TipoBulto,
		Descripcion:    topic.Descripcion,
		Cantidad:       topic.Cantidad,
		PesoBruto:      utils.DecodeBinaryDecimal(topic.PesoBruto, 3),
		UnidadPeso:     topic.UnidadPeso,
		Volumen:        utils.DecodeBinaryDecimal(topic.Volumen, 2),
		UnidadVolumen:  topic.UnidadVolumen,
		Observaciones:  topic.Observaciones,
		CargaCnt:       topic.CargaCnt,
		Evento:         utils.MapOperation(*src.Op),
		FechaEvento:    utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:  utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &item
}

func ConvertToBLItemImo(src *model.BlItemImoTopic) *model.BlItemImo {
	var topic model.BlItemImoT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	imo := model.BlItemImo{
		ID:            topic.ID,
		ClaseImo:      topic.ClaseImo,
		NumeroImo:     topic.NumeroImo,
		BlItemID:      topic.BlItemID,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso: utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &imo
}

// ================== CONTENEDORES ==================
func ConvertToBLItemContenedor(src *model.BlItemContenedorTopic) *model.BlItemContenedor {
	var topic model.BlItemContenedorT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	cnt := model.BlItemContenedor{
		ID:             topic.ID,
		BlItemID:       topic.BlItemID,
		Sigla:          topic.Sigla,
		Numero:         topic.Numero,
		Digito:         topic.Digito,
		TipoCnt:        topic.TipoCnt,
		CntSo:          topic.CntSo,
		Peso:           utils.ParseToFloat(topic.Peso),
		ValorIDOp:      topic.ValorIDOp,
		NombreOperador: topic.NombreOperador,
		Estado:         topic.Estado,
		Evento:         utils.MapOperation(*src.Op),
		FechaEvento:    utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:  utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &cnt
}

func ConvertToBLItemContenedorImo(src *model.BlItemContenedorImoTopic) *model.BlItemContenedorImo {
	var topic model.BlItemContenedorImoT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	imo := model.BlItemContenedorImo{
		ID:                 topic.ID,
		ClaseImo:           topic.ClaseImo,
		NumeroImo:          topic.NumeroImo,
		BlItemContenedorID: topic.BlItemContenedorID,
		Evento:             utils.MapOperation(*src.Op),
		FechaEvento:        utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:      utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &imo
}

func ConvertToBLItemContenedorSello(src *model.BlItemContenedorSelloTopic) *model.BlItemContenedorSello {
	var topic model.BlItemContenedorSelloT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	sello := model.BlItemContenedorSello{
		ID:                 topic.ID,
		Codigo:             topic.Codigo,
		Emisor:             topic.Emisor,
		Numero:             topic.Numero,
		BlItemContenedorID: topic.BlItemContenedorID,
		Evento:             utils.MapOperation(*src.Op),
		FechaEvento:        utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:      utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &sello
}

// ================== VARIOS ==================
func ConvertToBLLocacion(src *model.BlLocacionTopic) *model.BlLocacion {
	var topic model.BlLocacionT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	loc := model.BlLocacion{
		ID:            topic.ID,
		Codigo:        topic.Codigo,
		Descripcion:   topic.Descripcion,
		Nombre:        topic.Nombre,
		BlNroBl:       topic.BlNroBl,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso: utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &loc
}

func ConvertToBLObservacion(src *model.BlObservacionTopic) *model.BlObservacion {
	var topic model.BlObservacionT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	obs := model.BlObservacion{
		ID:            topic.ID,
		Contenido:     topic.Contenido,
		Nombre:        topic.Nombre,
		BlNroBl:       topic.BlNroBl,
		Evento:        utils.MapOperation(*src.Op),
		FechaEvento:   utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso: utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &obs
}

func ConvertToBLParticipante(topic *model.BlParticipanteTopic) *model.BlParticipante {
	var participante *model.BlParticipanteT

	if topic.After != nil {
		participante = topic.After
	} else {
		participante = topic.Before
	}
	if participante.FechaTraspaso == nil {
		return nil
	}
	return &model.BlParticipante{
		ID:               participante.ID,
		CodigoAlmacen:    participante.CodigoAlmacen,
		CodigoPais:       participante.CodigoPais,
		Comuna:           participante.Comuna,
		Direccion:        participante.Direccion,
		Email:            participante.Email,
		NacionID:         participante.NacionID,
		Nombres:          participante.Nombres,
		Telefono:         participante.Telefono,
		TipoID:           participante.TipoID,
		TipoParticipante: participante.TipoParticipante,
		ValorID:          participante.ValorID,
		BlNroBl:          participante.BlNroBl,
		Evento:           utils.MapOperation(*topic.Op),
		FechaEvento:      utils.ToFormattedDateTimeEvento(topic.TsMS),
		FechaTraspaso:    utils.ToFormattedDateTime(participante.FechaTraspaso),
	}
}

func ConvertToBLReferencia(src *model.BlReferenciaTopic) *model.BlReferencia {
	var topic model.BlReferenciaT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	ref := model.BlReferencia{
		ID:             topic.ID,
		Emisor:         topic.Emisor,
		Fecha:          utils.ToFormattedDate(topic.Fecha),
		NacIDEmisor:    topic.NacIDEmisor,
		Numero:         topic.Numero,
		TipoDocumento:  topic.TipoDocumento,
		TipoIDEmisor:   topic.TipoIDEmisor,
		TipoReferencia: topic.TipoReferencia,
		ValorIDEmisor:  topic.ValorIDEmisor,
		BlNroBl:        topic.BlNroBl,
		Evento:         utils.MapOperation(*src.Op),
		FechaEvento:    utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:  utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &ref
}

func ConvertToBLTransbordo(src *model.BlTransbordoTopic) *model.BlTransbordo {
	var topic model.BlTransbordoT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	t := model.BlTransbordo{
		ID:               topic.ID,
		CodigoLugar:      topic.CodigoLugar,
		DescripcionLugar: topic.DescripcionLugar,
		FechaArribo:      utils.ToFormattedDateTime(topic.FechaArribo),
		BlNroBl:          topic.BlNroBl,
		Evento:           utils.MapOperation(*src.Op),
		FechaEvento:      utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:    utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &t
}

func ConvertToBLTransporte(src *model.BlTransporteTopic) *model.BlTransporte {
	var topic model.BlTransporteT
	if *src.Op == "d" {
		topic = *src.Before
	} else {
		topic = *src.After
	}
	if topic.FechaTraspaso == nil {
		return nil
	}
	tr := model.BlTransporte{
		NombreNave:       topic.NombreNave,
		SentidoOperacion: topic.SentidoOperacion,
		BlNroBl:          topic.BlNroBl,
		Evento:           utils.MapOperation(*src.Op),
		FechaEvento:      utils.ToFormattedDateTimeEvento(src.TsMS),
		FechaTraspaso:    utils.ToFormattedDateTime(topic.FechaTraspaso),
	}

	return &tr
}

// ================== CLONADOR ==================
func cloneBase(bl *model.BL) *model.BL {
	if bl == nil {

		return &model.BL{IDMongo: primitive.NewObjectID()}
	}

	clone := *bl
	clone.IDMongo = primitive.NewObjectID()

	if bl.BlParticipantes != nil {
		participantesCopy := make([]model.BlParticipante, len(*bl.BlParticipantes))
		copy(participantesCopy, *bl.BlParticipantes)
		clone.BlParticipantes = &participantesCopy
	}

	if bl.BlFechas != nil {
		cp := make([]model.BlFecha, len(*bl.BlFechas))
		copy(cp, *bl.BlFechas)
		clone.BlFechas = &cp
	}

	if bl.BlFletes != nil {
		cp := make([]model.BlFlete, len(*bl.BlFletes))
		copy(cp, *bl.BlFletes)
		clone.BlFletes = &cp
	}

	if bl.BlItems != nil {
		cp := make([]model.BlItem, len(*bl.BlItems))
		copy(cp, *bl.BlItems)
		clone.BlItems = &cp
	}

	if bl.BlLocaciones != nil {
		cp := make([]model.BlLocacion, len(*bl.BlLocaciones))
		copy(cp, *bl.BlLocaciones)
		clone.BlLocaciones = &cp
	}

	if bl.BlObservaciones != nil {
		cp := make([]model.BlObservacion, len(*bl.BlObservaciones))
		copy(cp, *bl.BlObservaciones)
		clone.BlObservaciones = &cp
	}

	if bl.BlReferencias != nil {
		cp := make([]model.BlReferencia, len(*bl.BlReferencias))
		copy(cp, *bl.BlReferencias)
		clone.BlReferencias = &cp
	}

	if bl.BlTransbordos != nil {
		cp := make([]model.BlTransbordo, len(*bl.BlTransbordos))
		copy(cp, *bl.BlTransbordos)
		clone.BlTransbordos = &cp
	}

	if bl.BlTransportes != nil {
		cp := make([]model.BlTransporte, len(*bl.BlTransportes))
		copy(cp, *bl.BlTransportes)
		clone.BlTransportes = &cp
	}

	return &clone
}
func InicilizaVacio(bl *model.BL) *model.BL {

	if bl.BlParticipantes == nil {
		bl.BlParticipantes = &[]model.BlParticipante{}

	}

	if bl.BlFechas == nil {
		bl.BlFechas = &[]model.BlFecha{}
	}

	if bl.BlFletes == nil {
		bl.BlFletes = &[]model.BlFlete{}
	}

	if bl.BlItems == nil {
		bl.BlItems = &[]model.BlItem{}
	}

	if bl.BlLocaciones == nil {
		bl.BlLocaciones = &[]model.BlLocacion{}
	}

	if bl.BlObservaciones == nil {
		bl.BlObservaciones = &[]model.BlObservacion{}
	}

	if bl.BlReferencias == nil {
		bl.BlReferencias = &[]model.BlReferencia{}
	}

	if bl.BlTransbordos == nil {
		bl.BlTransbordos = &[]model.BlTransbordo{}
	}

	if bl.BlTransportes == nil {
		bl.BlTransportes = &[]model.BlTransporte{}
	}

	return bl
}
