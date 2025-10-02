package model

type VisacionTopic struct {
	Before      *Visacion   `json:"before,omitempty"`
	After       *Visacion   `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}
type MercanciasDespachadasTopic struct {
	Before      *MercanciasDespachadas `json:"before,omitempty"`
	After       *MercanciasDespachadas `json:"after,omitempty"`
	Source      *Source                `json:"source,omitempty"`
	Op          *string                `json:"op,omitempty"`
	TsMS        *int64                 `json:"ts_ms,omitempty"`
	Transaction interface{}            `json:"transaction"`
}
type Visacion struct {
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

type MercanciasDespachadas struct {
	ID            *int64  `json:"id,omitempty"`
	Cantidad      *int64  `json:"cantidad,omitempty"`
	FechaDespacho *int64  `json:"fecha_despacho,omitempty"`
	Peso          *string `json:"peso,omitempty"`
	UsuarioCR     *string `json:"usuario_cr,omitempty"`
	Despachado    *int64  `json:"despachado,omitempty"`
	Visaje        *int64  `json:"visaje,omitempty"`
	Volumen       *string `json:"volumen,omitempty"`
}
