package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PapeletaExpoTopic struct {
	Before      *PapeletaExpoT `json:"before,omitempty"`
	After       *PapeletaExpoT `json:"after,omitempty"`
	Source      *Source        `json:"source,omitempty"`
	Op          string         `json:"op,omitempty"`
	TsMS        int64          `json:"ts_ms,omitempty"`
	Transaction interface{}    `json:"transaction"`
}
type PapeletaExpoDetalleTopic struct {
	Before      *PapeletaExpoDetalleT `json:"before,omitempty"`
	After       *PapeletaExpoDetalleT `json:"after,omitempty"`
	Source      *Source               `json:"source,omitempty"`
	Op          string                `json:"op,omitempty"`
	TsMS        *int64                `json:"ts_ms,omitempty"`
	Transaction interface{}           `json:"transaction"`
}

type PapeletaExpoT struct {
	ID                *int64  `json:"id,omitempty"`
	NroPapeleta       *string `json:"nro_papeleta,omitempty"`
	Operacion         *string `json:"operacion,omitempty"`
	Zona              *string `json:"zona,omitempty"`
	Documento         *int64  `json:"documento,omitempty"`
	NroDoc            *string `json:"nro_doc,omitempty"`
	Emisor            *string `json:"emisor,omitempty"`
	FechaDoc          *int64  `json:"fecha_doc,omitempty"`
	FechaRecepcion    *int64  `json:"fecha_recepcion,omitempty"`
	Booking           *string `json:"booking,omitempty"`
	FechaCR           *int64  `json:"fecha_cr,omitempty"`
	FechaUp           *int64  `json:"fecha_up,omitempty"`
	UsuarioCR         *string `json:"usuario_cr,omitempty"`
	UsuarioUp         *string `json:"usuario_up"`
	DeclaracionAduana *string `json:"declaracion_aduana"`
	NroDocAduana      *string `json:"nro_doc_aduana"`
	FechaDocAduana    *int64  `json:"fecha_doc_aduana,omitempty"`
	Producto          *string `json:"producto"`
	Codigo            *string `json:"codigo"`
	Lote              *string `json:"lote"`
	Cosecha           *string `json:"cosecha"`
	CodigoFert        *string `json:"codigo_fert"`
	Observaciones     *string `json:"observaciones,omitempty"`
	Empresa           *string `json:"empresa,omitempty"`
	IngresoExpo       *int64  `json:"ingreso_expo,omitempty"`
	Estado            *string `json:"estado,omitempty"`
	Forwarder         *string `json:"forwarder"`
	MotivoAnulacion   *string `json:"motivo_anulacion"`
	Aga               *string `json:"aga"`
	Marcas            *string `json:"marcas"`
	PtoDesmb          *string `json:"pto_desmb"`
	NroViaje          *string `json:"nro_viaje"`
	Nave              *string `json:"nave"`
	Sku               *string `json:"sku"`
	Liberada          *bool   `json:"liberada,omitempty"`
	MotivoLiberacion  *string `json:"motivo_liberacion,omitempty"`
	IDRecepcion       *int64  `json:"id_recepcion"`
}
type PapeletaExpoDetalleT struct {
	ID         *int64  `json:"id,omitempty"`
	IDPapeleta *int64  `json:"id_papeleta,omitempty"`
	Bulto      *string `json:"bulto,omitempty"`
	Cantidad   *int64  `json:"cantidad,omitempty"`
	Peso       *string `json:"peso,omitempty"`
	Estado     *string `json:"estado,omitempty"`
	Contenedor *string `json:"contenedor"`
	Sello      *string `json:"sello"`
	Chassis    *string `json:"chassis,omitempty"`
	Volumen    *string `json:"volumen,omitempty"`
	Situacion  *string `json:"situacion,omitempty"`
	Marca      *string `json:"marca"`
}
type PapeletaExpoDetalle struct {
	IDMongo     primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID          *int64             `bson:"id,omitempty" json:"id,omitempty"`
	IDPapeleta  *int64             `bson:"id_papeleta,omitempty" json:"id_papeleta,omitempty"`
	Bulto       *string            `bson:"bulto,omitempty" json:"bulto,omitempty"`
	Cantidad    *int64             `bson:"cantidad,omitempty" json:"cantidad,omitempty"`
	Peso        *float64           `bson:"peso,omitempty" json:"peso,omitempty"`
	Estado      *string            `bson:"estado,omitempty" json:"estado,omitempty"`
	Contenedor  *string            `bson:"contenedor,omitempty" json:"contenedor"`
	Sello       *string            `bson:"sello,omitempty" json:"sello"`
	Chassis     *string            `bson:"chassis,omitempty" json:"chassis,omitempty"`
	Volumen     *float64           `bson:"volumen,omitempty" json:"volumen,omitempty"`
	Situacion   *string            `bson:"situacion,omitempty" json:"situacion,omitempty"`
	Marca       *string            `bson:"marca,omitempty" json:"marca"`
	Evento      string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}

type PapeletaExpo struct {
	IDMongo           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                *int64             `bson:"id,omitempty" json:"id,omitempty"`
	NroPapeleta       *string            `bson:"nro_papeleta,omitempty" json:"nro_papeleta,omitempty"`
	Operacion         *string            `bson:"operacion,omitempty" json:"operacion,omitempty"`
	Zona              *string            `bson:"zona,omitempty" json:"zona,omitempty"`
	Documento         *int64             `bson:"documento,omitempty" json:"documento,omitempty"`
	NroDoc            *string            `bson:"nro_doc,omitempty" json:"nro_doc,omitempty"`
	Emisor            *string            `bson:"emisor,omitempty" json:"emisor,omitempty"`
	FechaDoc          *string            `bson:"fecha_doc,omitempty" json:"fecha_doc,omitempty"`
	FechaRecepcion    *string            `bson:"fecha_recepcion,omitempty" json:"fecha_recepcion,omitempty"`
	Booking           *string            `bson:"booking,omitempty" json:"booking,omitempty"`
	FechaCR           *string            `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	FechaUp           *string            `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	UsuarioCR         *string            `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	UsuarioUp         *string            `bson:"usuario_up,omitempty" json:"usuario_up"`
	DeclaracionAduana *string            `bson:"declaracion_aduana,omitempty" json:"declaracion_aduana"`
	NroDocAduana      *string            `bson:"nro_doc_aduana,omitempty" json:"nro_doc_aduana"`
	FechaDocAduana    *string            `bson:"fecha_doc_aduana,omitempty" json:"fecha_doc_aduana,omitempty"`
	Producto          *string            `bson:"producto,omitempty" json:"producto"`
	Codigo            *string            `bson:"codigo,omitempty" json:"codigo"`
	Lote              *string            `bson:"lote,omitempty" json:"lote"`
	Cosecha           *string            `bson:"cosecha,omitempty" json:"cosecha"`
	CodigoFert        *string            `bson:"codigo_fert,omitempty" json:"codigo_fert"`
	Observaciones     *string            `bson:"observaciones,omitempty" json:"observaciones,omitempty"`
	Empresa           *string            `bson:"empresa,omitempty" json:"empresa,omitempty"`
	IngresoExpo       *int64             `bson:"ingreso_expo,omitempty" json:"ingreso_expo,omitempty"`
	Estado            *string            `bson:"estado,omitempty" json:"estado,omitempty"`
	Forwarder         *string            `bson:"forwarder,omitempty" json:"forwarder"`
	MotivoAnulacion   *string            `bson:"motivo_anulacion,omitempty" json:"motivo_anulacion"`
	Aga               *string            `bson:"aga,omitempty" json:"aga"`
	Marcas            *string            `bson:"marcas,omitempty" json:"marcas"`
	PtoDesmb          *string            `bson:"pto_desmb,omitempty" json:"pto_desmb"`
	NroViaje          *string            `bson:"nro_viaje,omitempty" json:"nro_viaje"`
	Nave              *string            `bson:"nave,omitempty" json:"nave"`
	Sku               *string            `bson:"sku,omitempty" json:"sku"`
	Liberada          *bool              `bson:"liberada,omitempty" json:"liberada,omitempty"`
	MotivoLiberacion  *string            `bson:"motivo_liberacion,omitempty" json:"motivo_liberacion,omitempty"`
	IDRecepcion       *int64             `bson:"id_recepcion,omitempty" json:"id_recepcion"`
	Evento            string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento       *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`

	Detalles *[]PapeletaExpoDetalle `bson:"detalles,omitempty" json:"detalles,omitempty"`
}
