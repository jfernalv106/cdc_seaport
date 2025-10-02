package model

import "go.mongodb.org/mongo-driver/bson/primitive"

/**Nota Credito**/
type NotaCreditoTopic struct {
	Before      *NotaCreditoT `json:"before,omitempty"`
	After       *NotaCreditoT `json:"after,omitempty"`
	Source      *Source       `json:"source,omitempty"`
	Op          *string       `json:"op,omitempty"`
	TsMS        *int64        `json:"ts_ms,omitempty"`
	Transaction interface{}   `json:"transaction"`
}

type NotaCreditoT struct {
	IDNroNc       *int64  `json:"id_nro_nc,omitempty"`
	IDFolNc       *int64  `json:"id_fol_nc,omitempty"`
	IDFactura     *int64  `json:"id_factura,omitempty"`
	IDFolio       *int64  `json:"id_folio,omitempty"`
	FechaCR       *int64  `json:"fecha_cr,omitempty"`
	UsuarioCR     *string `json:"usuario_cr,omitempty"`
	Observaciones *string `json:"observaciones"`
	ValorNeto     *string `json:"valor_neto,omitempty"`
	Tipo          *string `json:"tipo,omitempty"`
	Empresa       *string `json:"empresa,omitempty"`
	Traspasada    *bool   `json:"traspasada,omitempty"`
	IDSap         *int64  `json:"id_sap"`
}
type NotaCreditoServTopic struct {
	Before      *NotaCreditoServT `json:"before,omitempty"`
	After       *NotaCreditoServT `json:"after,omitempty"`
	Source      *Source           `json:"source,omitempty"`
	Op          *string           `json:"op,omitempty"`
	TsMS        *int64            `json:"ts_ms,omitempty"`
	Transaction interface{}       `json:"transaction"`
}

type NotaCreditoServT struct {
	ID            *int64  `json:"id,omitempty"`
	IDNroNc       *int64  `json:"id_nro_nc,omitempty"`
	CodServ       *int64  `json:"cod_serv,omitempty"`
	ValorAplicado *string `json:"valor_aplicado,omitempty"`
	Visaje        *int64  `json:"visaje"`
	DfID          *int64  `json:"df_id,omitempty"`
}

type NotaCredito struct {
	IDMongo       primitive.ObjectID     `bson:"_id,omitempty" json:"_id,omitempty"`
	IDNroNc       *int64                 `bson:"id_nro_nc,omitempty" json:"id_nro_nc,omitempty"`
	IDFolNc       *int64                 `bson:"id_fol_nc,omitempty" json:"id_fol_nc,omitempty"`
	IDFactura     *int64                 `bson:"id_factura,omitempty" json:"id_factura,omitempty"`
	IDFolio       *int64                 `bson:"id_folio,omitempty" json:"id_folio,omitempty"`
	FechaCR       *string                `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	UsuarioCR     *string                `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	Observaciones *string                `bson:"observaciones,omitempty" json:"observaciones,omitempty"`
	ValorNeto     *float64               `bson:"valor_neto,omitempty" json:"valor_neto,omitempty"`
	Tipo          *string                `bson:"tipo,omitempty" json:"tipo,omitempty"`
	Empresa       *string                `bson:"empresa,omitempty" json:"empresa,omitempty"`
	Traspasada    *bool                  `bson:"traspasada,omitempty" json:"traspasada,omitempty"`
	IDSap         *int64                 `bson:"id_sap,omitempty" json:"id_sap,omitempty"`
	Evento        string                 `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento   *string                `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
	Servicios     *[]NotaCreditoServicio `bson:"servicios,omitempty" json:"servicios,omitempty"`
}
type NotaCreditoServicio struct {
	IDMongo       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID            *int64             `bson:"id,omitempty" json:"id,omitempty"`
	IDNroNc       *int64             `bson:"id_nro_nc,omitempty" json:"id_nro_nc,omitempty"`
	CodServ       *int64             `bson:"cod_serv,omitempty" json:"cod_serv,omitempty"`
	ValorAplicado *float64           `bson:"valor_aplicado,omitempty" json:"valor_aplicado,omitempty"`
	Visaje        *int64             `bson:"visaje,omitempty" json:"visaje,omitempty"`
	DfID          *int64             `bson:"df_id,omitempty" json:"df_id,omitempty"`
	Evento        string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento   *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
