package model

type ManifiestoTopic struct {
	Before      *Manifiesto `json:"before,omitempty"`
	After       *Manifiesto `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type Manifiesto struct {
	NroMftoInterno        *string `json:"nro_mfto_interno,omitempty"`
	NroMfto               *int64  `json:"nro_mfto,omitempty"`
	Nave                  *string `json:"nave,omitempty"`
	Viaje                 *string `json:"viaje,omitempty"`
	Servicio              *string `json:"servicio,omitempty"`
	Agencia               *string `json:"agencia,omitempty"`
	FechaMfto             *int64  `json:"fecha_mfto,omitempty"`
	Estado                *string `json:"estado,omitempty"`
	InscripcionDesde      *int64  `json:"inscripcion_desde,omitempty"`
	InscripcionHasta      *int64  `json:"inscripcion_hasta,omitempty"`
	TrasladoPuerto        *int64  `json:"traslado_puerto,omitempty"`
	TipoMfto              *string `json:"tipo_mfto"`
	TipoAccion            *string `json:"tipo_accion"`
	Version               *int64  `json:"version"`
	CondCarga             *string `json:"cond_carga"`
	SitioAtraque          *string `json:"sitio_atraque"`
	Almacen               *string `json:"almacen,omitempty"`
	UsuarioCR             *string `json:"usuario_cr,omitempty"`
	FechaCR               *int64  `json:"fecha_cr,omitempty"`
	UsuarioUp             *string `json:"usuario_up,omitempty"`
	FechaUp               *int64  `json:"fecha_up,omitempty"`
	FechaUltLec           *int64  `json:"fecha_ult_lec,omitempty"`
	TipoCarga             *string `json:"tipo_carga,omitempty"`
	MotivoAnulacion       *string `json:"motivo_anulacion,omitempty"`
	FechaTraspaso         *int64  `json:"fecha_traspaso,omitempty"`
	Zarpe                 *int64  `json:"zarpe,omitempty"`
	NroMftoInternoEmpresa *string `json:"nro_mfto_interno_empresa,omitempty"`
	ArriboEfectivo        *int64  `json:"arribo_efectivo,omitempty"`
}
