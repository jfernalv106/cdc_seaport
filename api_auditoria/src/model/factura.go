package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type FacturaTopic struct {
	Before      *FacturaT   `json:"before,omitempty"`
	After       *FacturaT   `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          string      `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}
type FacturaDetalleTopic struct {
	Before      *FacturaDetalleT `json:"before,omitempty"`
	After       *FacturaDetalleT `json:"after,omitempty"`
	Source      *Source          `json:"source,omitempty"`
	Op          string           `json:"op,omitempty"`
	TsMS        *int64           `json:"ts_ms,omitempty"`
	Transaction interface{}      `json:"transaction"`
}
type Factura struct {
	IDMongo                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                       *int64             `bson:"id,omitempty" json:"id,omitempty"`
	RutCliente               *string            `bson:"rut_cliente,omitempty" json:"rut_cliente,omitempty"`
	IDFolio                  *int64             `bson:"id_folio,omitempty" json:"id_folio,omitempty"`
	ManifiestoNroMftoInterno *string            `bson:"manifiesto_nro_mfto_interno,omitempty" json:"manifiesto_nro_mfto_interno,omitempty"`
	Origen                   *string            `bson:"origen,omitempty" json:"origen,omitempty"`
	Moneda                   *string            `bson:"moneda,omitempty" json:"moneda,omitempty"`
	TipoVenta                *string            `bson:"tipo_venta,omitempty" json:"tipo_venta,omitempty"`
	FechaCambio              *string            `bson:"fecha_cambio,omitempty" json:"fecha_cambio,omitempty"`
	ValorCambio              *float64           `bson:"valor_cambio,omitempty" json:"valor_cambio,omitempty"`
	UsuarioCR                *string            `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	FechaCR                  *string            `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	UsuarioUp                *string            `bson:"usuario_up,omitempty" json:"usuario_up,omitempty"`
	FechaUp                  *string            `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	Folio                    *int64             `bson:"folio,omitempty" json:"folio,omitempty"`
	Comentario               *string            `bson:"comentario,omitempty" json:"comentario,omitempty"`
	Empresa                  *string            `bson:"empresa,omitempty" json:"empresa,omitempty"`
	Estado                   *string            `bson:"estado,omitempty" json:"estado,omitempty"`
	TotalNeto                *int64             `bson:"total_neto,omitempty" json:"total_neto,omitempty"`
	Aga                      *string            `bson:"aga,omitempty" json:"aga,omitempty"`
	NroIngRecaudacion        *string            `bson:"nro_ing_recaudacion,omitempty" json:"nro_ing_recaudacion,omitempty"`
	NroBl                    *string            `bson:"nro_bl,omitempty" json:"nro_bl,omitempty"`
	FacturaSap               *int64             `bson:"factura_sap,omitempty" json:"factura_sap,omitempty"`
	CorreoDespacho           *bool              `bson:"correo_despacho,omitempty" json:"correo_despacho,omitempty"`
	Evento                   string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento              *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
	Tipo                     *string            `bson:"tipo,omitempty" json:"tipo,omitempty"`
	NroOportunidad           *string            `bson:"nro_oportunidad,omitempty" json:"nro_oportunidad,omitempty"`
	Forw                     *string            `bson:"forw,omitempty" json:"forw,omitempty"`
	OrdenCompra              *string            `bson:"orden_compra,omitempty" json:"orden_compra,omitempty"`
	Despacho                 *int64             `bson:"despacho,omitempty" json:"despacho,omitempty"`
	Mandante                 *string            `bson:"mandante,omitempty" json:"mandante,omitempty"`
	Zona                     *string            `bson:"zona,omitempty" json:"zona,omitempty"`

	Detalles *[]FacturaDetalle `bson:"detalles,omitempty" json:"detalles,omitempty"`
}
type FacturaT struct {
	ID                       *int64  `json:"id,omitempty"`
	RutCliente               *string `json:"rut_cliente,omitempty"`
	IDFolio                  *int64  `json:"id_folio,omitempty"`
	ManifiestoNroMftoInterno *string `json:"manifiesto_nro_mfto_interno,omitempty"`
	Origen                   *string `json:"origen,omitempty"`
	Zona                     *string `json:"zona,omitempty"`
	Moneda                   *string `json:"moneda,omitempty"`
	TipoVenta                *string `json:"tipo_venta,omitempty"`
	FechaCambio              *int64  `json:"fecha_cambio,omitempty"`
	ValorCambio              *string `json:"valor_cambio,omitempty"`
	CobrarA                  *string `json:"cobrar_a"`
	UsuarioCR                *string `json:"usuario_cr,omitempty"`
	FechaCR                  *int64  `json:"fecha_cr,omitempty"`
	UsuarioUp                *string `json:"usuario_up,omitempty"`
	FechaUp                  *int64  `json:"fecha_up,omitempty"`
	Folio                    *int64  `json:"folio,omitempty"`
	Comentario               *string `json:"comentario,omitempty"`
	Empresa                  *string `json:"empresa,omitempty"`
	Estado                   *string `json:"estado,omitempty"`
	TotalNeto                *string `json:"total_neto,omitempty"`
	Aga                      *string `json:"aga,omitempty"`
	NroIngRecaudacion        *string `json:"nro_ing_recaudacion,omitempty"`
	Forw                     *string `json:"forw"`
	NroBl                    *string `json:"nro_bl,omitempty"`
	OrdenCompra              *string `json:"orden_compra"`
	FacturaSap               *int64  `json:"factura_sap,omitempty"`
	Despacho                 *int64  `json:"despacho"`
	Mandante                 *string `json:"mandante"`
	Tipo                     *string `json:"tipo"`
	NroOportunidad           *string `json:"nro_oportunidad"`
	CorreoDespacho           *bool   `json:"correo_despacho,omitempty"`
}

type Facturas []*Factura

type FacturaDetalle struct {
	IDMongo         primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID              *int64             `bson:"id,omitempty" json:"id,omitempty"`
	Cantidad        *int64             `bson:"cantidad,omitempty" json:"cantidad,omitempty"`
	Descuento       *float64           `bson:"descuento,omitempty" json:"descuento,omitempty"`
	Detalle         *string            `bson:"detalle,omitempty" json:"detalle,omitempty"`
	Dias            *int64             `bson:"dias,omitempty" json:"dias,omitempty"`
	FechaCR         *string            `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	GlosaFactura    *string            `bson:"glosa_factura,omitempty" json:"glosa_factura,omitempty"`
	Imo             *string            `bson:"imo,omitempty" json:"imo,omitempty"`
	Moneda          *string            `bson:"moneda,omitempty" json:"moneda,omitempty"`
	Neto            *float64           `bson:"neto,omitempty" json:"neto,omitempty"`
	Peso            *float64           `bson:"peso,omitempty" json:"peso,omitempty"`
	Recargo         *float64           `bson:"recargo,omitempty" json:"recargo,omitempty"`
	Tarifa          *string            `bson:"tarifa,omitempty" json:"tarifa,omitempty"`
	Unidad          *string            `bson:"unidad,omitempty" json:"unidad,omitempty"`
	UsuarioCR       *string            `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	ValorUnitario   *float64           `bson:"valor_unitario,omitempty" json:"valor_unitario,omitempty"`
	Volumen         *float64           `bson:"volumen,omitempty" json:"volumen,omitempty"`
	IDFactura       *int64             `bson:"id_factura,omitempty" json:"id_factura,omitempty"`
	NroPapeleta     *string            `bson:"nro_papeleta,omitempty" json:"nro_papeleta,omitempty"`
	Servicio        *int64             `bson:"servicio,omitempty" json:"servicio,omitempty"`
	Visaje          *int64             `bson:"visaje,omitempty" json:"visaje,omitempty"`
	NetoPeso        *float64           `bson:"neto_peso,omitempty" json:"neto_peso,omitempty"`
	MotivoDescuento *string            `bson:"motivo_descuento,omitempty" json:"motivo_descuento,omitempty"`
	MotivoRecargo   *string            `bson:"motivo_recargo,omitempty" json:"motivo_recargo,omitempty"`
	PapeletaExpo    *int64             `bson:"papeleta_expo,omitempty" json:"papeleta_expo,omitempty"`
	VisajeVa        *int64             `bson:"visaje_va,omitempty" json:"visaje_va,omitempty"`
	DiasReales      *int64             `bson:"dias_reales,omitempty" json:"dias_reales,omitempty"`
	Evento          *string            `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento     *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type FacturaDetalleT struct {
	ID              *int64  `json:"id,omitempty"`
	Cantidad        *int64  `json:"cantidad,omitempty"`
	Descuento       *string `json:"descuento,omitempty"`
	Detalle         *string `json:"detalle,omitempty"`
	Dias            *int64  `json:"dias,omitempty"`
	FechaCR         *int64  `json:"fecha_cr,omitempty"`
	FechaInicio     *int64  `json:"fecha_inicio,omitempty"`
	GlosaFactura    *string `json:"glosa_factura,omitempty"`
	Imo             *string `json:"imo,omitempty"`
	Moneda          *string `json:"moneda,omitempty"`
	Neto            *string `json:"neto,omitempty"`
	Peso            *string `json:"peso,omitempty"`
	Recargo         *string `json:"recargo,omitempty"`
	Tarifa          *string `json:"tarifa,omitempty"`
	Unidad          *string `json:"unidad,omitempty"`
	UsuarioCR       *string `json:"usuario_cr,omitempty"`
	ValorUnitario   *string `json:"valor_unitario,omitempty"`
	Volumen         *string `json:"volumen,omitempty"`
	IDFactura       *int64  `json:"id_factura,omitempty"`
	NroPapeleta     *string `json:"nro_papeleta,omitempty"`
	Servicio        *int64  `json:"servicio,omitempty"`
	Visaje          *int64  `json:"visaje,omitempty"`
	NetoPeso        *string `json:"neto_peso,omitempty"`
	MotivoDescuento *string `json:"motivo_descuento,omitempty"`
	MotivoRecargo   *string `json:"motivo_recargo,omitempty"`
	PapeletaExpo    *int64  `json:"papeleta_expo,omitempty"`
	VisajeVa        *int64  `json:"visaje_va,omitempty"`
	DiasReales      *int64  `json:"dias_reales,omitempty"`
}

type FacturaDetalles []*FacturaDetalle
