package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Visacion struct {
	IDMongo       primitive.ObjectID      `bson:"_id,omitempty" json:"_id,omitempty"`
	ID            *int64                  `bson:"id,omitempty" json:"id,omitempty"`
	Aux           *string                 `bson:"aux,omitempty" json:"aux,omitempty"`
	Doc           *string                 `bson:"doc,omitempty" json:"doc,omitempty"`
	Nro           *string                 `bson:"nro,omitempty" json:"nro,omitempty"`
	Aga           *string                 `bson:"aga,omitempty" json:"aga,omitempty"`
	NroPapeleta   *string                 `bson:"nro_papeleta,omitempty" json:"nro_papeleta,omitempty"`
	FechaRetiro   *string                 `bson:"fecha_retiro,omitempty" json:"fecha_retiro,omitempty"`
	FechaCR       *string                 `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	CorrelativoPE *string                 `bson:"correlativo_pe,omitempty" json:"correlativo_pe,omitempty"`
	FechaDoc      *string                 `bson:"fecha_doc,omitempty" json:"fecha_doc,omitempty"`
	Evento        string                  `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento   *string                 `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
	Mercacias     *[]MercanciasDespachada `bson:"mercacias,omitempty" json:"mercacias,omitempty"`
}
type Visaciones []*Visacion

type MercanciasDespachada struct {
	IDMongo           primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                *int64             `bson:"id,omitempty" json:"id,omitempty"`
	Cantidad          *int64             `bson:"cantidad,omitempty" json:"cantidad,omitempty"`
	FechaDespacho     *string            `bson:"fecha_despacho,omitempty" json:"fecha_despacho,omitempty"`
	Peso              *float64           `bson:"peso,omitempty" json:"peso,omitempty"`
	IDDetallePapeleta *int64             `bson:"id_detalle_papeleta,omitempty" json:"id_detalle_papeleta,omitempty"`
	IDVisaje          *int64             `bson:"id_visaje,omitempty" json:"id_visaje,omitempty"`
	Volumen           *float64           `bson:"volumen,omitempty" json:"volumen,omitempty"`
	Bulto             *string            `bson:"bulto,omitempty" json:"bulto,omitempty"`
	UsuarioCR         *string            `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	Evento            string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento       *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type MercanciasDespachadas []*MercanciasDespachada

type VisacionTopic struct {
	Before      *VisacionT  `json:"before,omitempty"`
	After       *VisacionT  `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          string      `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}
type MercanciasDespachadasTopic struct {
	Before      *MercanciasDespachadasT `json:"before,omitempty"`
	After       *MercanciasDespachadasT `json:"after,omitempty"`
	Source      *Source                 `json:"source,omitempty"`
	Op          string                  `json:"op,omitempty"`
	TsMS        *int64                  `json:"ts_ms,omitempty"`
	Transaction interface{}             `json:"transaction"`
}
type VisacionT struct {
	ID                 *int64  `json:"id,omitempty"`
	Aux                *string `json:"aux,omitempty"`
	Doc                *string `json:"doc,omitempty"`
	Nro                *string `json:"nro,omitempty"`
	Ppt                *string `json:"ppt,omitempty"`
	Aga                *string `json:"aga,omitempty"`
	NroPapeleta        *string `json:"nro_papeleta,omitempty"`
	FechaRetiro        *int64  `json:"fecha_retiro,omitempty"`
	FechaCR            *int64  `json:"fecha_cr,omitempty"`
	CorrelativoPE      *string `json:"correlativo_pe,omitempty"`
	FechaDoc           *int64  `json:"fecha_doc,omitempty"`
	Fob                *string `json:"fob,omitempty"`
	VisacionAnticipada *bool   `json:"visacion_anticipada,omitempty"`
}

type MercanciasDespachadasT struct {
	ID            *int64  `json:"id,omitempty"`
	Cantidad      *int64  `json:"cantidad,omitempty"`
	FechaDespacho *int64  `json:"fecha_despacho,omitempty"`
	Peso          *string `json:"peso,omitempty"`
	UsuarioCR     *string `json:"usuario_cr,omitempty"`
	Despachado    *int64  `json:"despachado,omitempty"`
	Visaje        *int64  `json:"visaje,omitempty"`
	Volumen       *string `json:"volumen,omitempty"`
}
