package model

type FacturaTopic struct {
	Before      *Factura    `json:"before,omitempty"`
	After       *Factura    `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}
type FacturaDetalleTopic struct {
	Before      *FacturaDetalle `json:"before,omitempty"`
	After       *FacturaDetalle `json:"after,omitempty"`
	Source      *Source         `json:"source,omitempty"`
	Op          *string         `json:"op,omitempty"`
	TsMS        *int64          `json:"ts_ms,omitempty"`
	Transaction interface{}     `json:"transaction"`
}
type Factura struct {
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
	CobrarA                  *string `json:"cobrar_a,omitempty"`
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
	NroIngRecaudacion        *string `json:"nro_ing_recaudacion"`
	Forw                     *string `json:"forw,omitempty"`
	NroBl                    *string `json:"nro_bl,omitempty"`
	OrdenCompra              *string `json:"orden_compra"`
	FacturaSap               *int64  `json:"factura_sap,omitempty"`
	Despacho                 *int64  `json:"despacho,omitempty"`
	Mandante                 *string `json:"mandante,omitempty"`
	Tipo                     *string `json:"tipo,omitempty"`
	NroOportunidad           *string `json:"nro_oportunidad,omitempty"`
	CorreoDespacho           *bool   `json:"correo_despacho,omitempty"`
}

type FacturaDetalle struct {
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
