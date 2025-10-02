package model

type Despacho struct {
	IDMongo               string  `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                    *int64  `bson:"id,omitempty" json:"id,omitempty"`
	Visacion              *int64  `bson:"visacion,omitempty" json:"visacion,omitempty"`
	UsuarioCR             *string `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	Empresa               *string `bson:"empresa,omitempty" json:"empresa,omitempty"`
	FechaCR               *string `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	Patente               *string `bson:"patente,omitempty" json:"patente,omitempty"`
	RutChofer             *string `bson:"rut_chofer,omitempty" json:"rut_chofer,omitempty"`
	Autorizado            *bool   `bson:"autorizado,omitempty" json:"autorizado,omitempty"`
	UsuarioAutoriza       *string `bson:"usuario_autoriza,omitempty" json:"usuario_autoriza,omitempty"`
	FechaAutorizacion     *string `bson:"fecha_autorizacion,omitempty" json:"fecha_autorizacion,omitempty"`
	Correlativo           *int64  `bson:"correlativo,omitempty" json:"correlativo,omitempty"`
	Estado                *string `bson:"estado,omitempty" json:"estado,omitempty"`
	FechaUp               *string `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	NotivoAnulacion       *string `bson:"notivo_anulacion,omitempty" json:"notivo_anulacion,omitempty"`
	PapeletaExpo          *int64  `bson:"papeleta_expo,omitempty" json:"papeleta_expo,omitempty"`
	Contenedor            *string `bson:"contenedor,omitempty" json:"contenedor,omitempty"`
	GuiaDespacho          *int64  `bson:"guia_despacho,omitempty" json:"guia_despacho,omitempty"`
	TipoDespacho          *string `bson:"tipo_despacho,omitempty" json:"tipo_despacho,omitempty"`
	MotivoAnulacionSalida *string `bson:"motivo_anulacion_salida,omitempty" json:"motivo_anulacion_salida,omitempty"`
	UsuarioAnulaSalida    *string `bson:"usuario_anula_salida,omitempty" json:"usuario_anula_salida,omitempty"`
	FechaAnulacionSalida  *string `bson:"fecha_anulacion_salida,omitempty" json:"fecha_anulacion_salida,omitempty"`
	Evento                string  `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento           *string `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type DespachoT struct {
	IDMongo               string  `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                    *int64  `bson:"id,omitempty" json:"id,omitempty"`
	Visacion              *int64  `bson:"visacion,omitempty" json:"visacion,omitempty"`
	UsuarioCR             *string `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	Empresa               *string `bson:"empresa,omitempty" json:"empresa,omitempty"`
	FechaCR               *int64  `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	Patente               *string `bson:"patente,omitempty" json:"patente,omitempty"`
	RutChofer             *string `bson:"rut_chofer,omitempty" json:"rut_chofer,omitempty"`
	Autorizado            *bool   `bson:"autorizado,omitempty" json:"autorizado,omitempty"`
	UsuarioAutoriza       *string `bson:"usuario_autoriza,omitempty" json:"usuario_autoriza,omitempty"`
	FechaAutorizacion     *int64  `bson:"fecha_autorizacion,omitempty" json:"fecha_autorizacion,omitempty"`
	Correlativo           *int64  `bson:"correlativo,omitempty" json:"correlativo,omitempty"`
	Estado                *string `bson:"estado,omitempty" json:"estado,omitempty"`
	FechaUp               *int64  `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	NotivoAnulacion       *string `bson:"notivo_anulacion,omitempty" json:"notivo_anulacion,omitempty"`
	PapeletaExpo          *int64  `bson:"papeleta_expo,omitempty" json:"papeleta_expo,omitempty"`
	Contenedor            *string `bson:"contenedor,omitempty" json:"contenedor,omitempty"`
	GuiaDespacho          *int64  `bson:"guia_despacho,omitempty" json:"guia_despacho,omitempty"`
	TipoDespacho          *string `bson:"tipo_despacho,omitempty" json:"tipo_despacho,omitempty"`
	MotivoAnulacionSalida *string `bson:"motivo_anulacion_salida,omitempty" json:"motivo_anulacion_salida,omitempty"`
	UsuarioAnulaSalida    *string `bson:"usuario_anula_salida,omitempty" json:"usuario_anula_salida,omitempty"`
	FechaAnulacionSalida  *string `bson:"fecha_anulacion_salida,omitempty" json:"fecha_anulacion_salida,omitempty"`
	Evento                string  `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento           *string `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}

type Despachos []*Despacho

type DespachoTopic struct {
	Before      *DespachoT  `json:"before,omitempty"`
	After       *DespachoT  `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          string      `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}
