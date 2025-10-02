package model

type BultoTopic struct {
	Before      *Bulto      `json:"before,omitempty"`
	After       *Bulto      `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type Bulto struct {
	IDMongo     string  `bson:"_id,omitempty" json:"_id,omitempty"`
	Cod         *string `bson:"cod,omitempty" json:"cod,omitempty"`
	Nombre      *string `bson:"nombre,omitempty" json:"nombre,omitempty"`
	Evento      string  `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento *string `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type Bultos []*Bulto
