package model

type ManifiestoTopic struct {
	Before      *ManifiestoT `json:"before,omitempty"`
	After       *ManifiestoT `json:"after,omitempty"`
	Source      *Source      `json:"source,omitempty"`
	Op          string       `json:"op,omitempty"`
	TsMS        *int64       `json:"ts_ms,omitempty"`
	Transaction interface{}  `json:"transaction"`
}

type ManifiestoT struct {
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
type Manifiesto struct {
	IDMongo               string  `bson:"_id,omitempty" json:"_id,omitempty"`
	NroMftoInterno        *string `bson:"nro_mfto_interno,omitempty" json:"nro_mfto_interno,omitempty"`
	NroMfto               *int64  `bson:"nro_mfto,omitempty" json:"nro_mfto,omitempty"`
	Nave                  *string `bson:"nave,omitempty" json:"nave,omitempty"`
	Viaje                 *string `bson:"viaje,omitempty" json:"viaje,omitempty"`
	Servicio              *string `bson:"servicio,omitempty" json:"servicio,omitempty"`
	Agencia               *string `bson:"agencia,omitempty" json:"agencia,omitempty"`
	FechaMfto             *string `bson:"fecha_mfto,omitempty" json:"fecha_mfto,omitempty"`
	Estado                *string `bson:"estado,omitempty" json:"estado,omitempty"`
	InscripcionDesde      *string `bson:"inscripcion_desde,omitempty" json:"inscripcion_desde,omitempty"`
	InscripcionHasta      *string `bson:"inscripcion_hasta,omitempty" json:"inscripcion_hasta,omitempty"`
	TrasladoPuerto        *string `bson:"traslado_puerto,omitempty" json:"traslado_puerto,omitempty"`
	TipoMfto              *string `bson:"tipo_mfto,omitempty" json:"tipo_mfto,omitempty"`
	TipoAccion            *string `bson:"tipo_accion,omitempty" json:"tipo_accion,omitempty"`
	Version               *int64  `bson:"version,omitempty" json:"version,omitempty"`
	CondCarga             *string `bson:"cond_carga,omitempty" json:"cond_carga,omitempty"`
	SitioAtraque          *string `bson:"sitio_atraque,omitempty" json:"sitio_atraque,omitempty"`
	Almacen               *string `bson:"almacen,omitempty" json:"almacen,omitempty"`
	UsuarioCR             *string `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	FechaCR               *string `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	UsuarioUp             *string `bson:"usuario_up,omitempty" json:"usuario_up,omitempty"`
	FechaUp               *string `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	FechaUltLec           *string `bson:"fecha_ult_lec,omitempty" json:"fecha_ult_lec,omitempty"`
	TipoCarga             *string `bson:"tipo_carga,omitempty" json:"tipo_carga,omitempty"`
	MotivoAnulacion       *string `bson:"motivo_anulacion,omitempty" json:"motivo_anulacion,omitempty"`
	FechaTraspaso         *string `bson:"fecha_traspaso,omitempty" json:"fecha_traspaso,omitempty"`
	Zarpe                 *string `bson:"zarpe,omitempty" json:"zarpe,omitempty"`
	NroMftoInternoEmpresa *string `bson:"nro_mfto_interno_empresa,omitempty" json:"nro_mfto_interno_empresa,omitempty"`
	ArriboEfectivo        *string `bson:"arribo_efectivo,omitempty" json:"arribo_efectivo,omitempty"`
	Evento                string  `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento           *string `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type Manifiestos []*Manifiesto
