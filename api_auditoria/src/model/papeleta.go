package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PapeletaRecepcion struct {
	IDMongo                  primitive.ObjectID          `bson:"_id,omitempty" json:"_id,omitempty"`
	NroPapeleta              *string                     `bson:"nro_papeleta,omitempty" json:"nro_papeleta,omitempty"`
	Etiqueta                 *string                     `bson:"etiqueta,omitempty" json:"etiqueta,omitempty"`
	Operacion                *string                     `bson:"operacion,omitempty" json:"operacion,omitempty"`
	Zona                     *string                     `bson:"zona,omitempty" json:"zona,omitempty"`
	NroBl                    *string                     `bson:"nro_bl,omitempty" json:"nro_bl,omitempty"`
	PtoOrigen                *string                     `bson:"pto_origen,omitempty" json:"pto_origen,omitempty"`
	PtoDescarga              *string                     `bson:"pto_descarga,omitempty" json:"pto_descarga,omitempty"`
	PtoDestino               *string                     `bson:"pto_destino,omitempty" json:"pto_destino,omitempty"`
	Consiganatario           *string                     `bson:"consiganatario,omitempty" json:"consiganatario,omitempty"`
	DireccionConsignatario   *string                     `bson:"direccion_consignatario,omitempty" json:"direccion_consignatario,omitempty"`
	RutConsignatario         *string                     `bson:"rut_consignatario,omitempty" json:"rut_consignatario,omitempty"`
	ZonaAlmacenaje           *string                     `bson:"zona_almacenaje,omitempty" json:"zona_almacenaje,omitempty"`
	Ubicacion                *string                     `bson:"ubicacion,omitempty" json:"ubicacion,omitempty"`
	GuardaAlmacen            *string                     `bson:"guarda_almacen,omitempty" json:"guarda_almacen,omitempty"`
	InicioAlmacenaje         *string                     `bson:"inicio_almacenaje,omitempty" json:"inicio_almacenaje,omitempty"`
	ManifiestoNroMftoInterno *string                     `bson:"manifiesto_nro_mfto_interno,omitempty" json:"manifiesto_nro_mfto_interno,omitempty"`
	Forwarder                *string                     `bson:"forwarder,omitempty" json:"forwarder,omitempty"`
	Estado                   *string                     `bson:"estado,omitempty" json:"estado,omitempty"`
	TipoCarga                *string                     `bson:"tipo_carga,omitempty" json:"tipo_carga,omitempty"`
	UsuarioCR                *string                     `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	FechaCR                  *string                     `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	FechaUP                  *string                     `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	MotivoLiberacion         *string                     `bson:"motivo_liberacion,omitempty" json:"motivo_liberacion,omitempty"`
	FechaDescon              *string                     `bson:"fecha_descon,omitempty" json:"fecha_descon,omitempty"`
	Tipo                     *string                     `bson:"tipo,omitempty" json:"tipo,omitempty"`
	PrePapeleta              *int64                      `bson:"pre_papeleta,omitempty" json:"pre_papeleta,omitempty"`
	PesoManifestado          *int64                      `bson:"peso_manifestado,omitempty" json:"peso_manifestado,omitempty"`
	Noty                     *string                     `bson:"noty,omitempty" json:"noty,omitempty"`
	NroPapeletaEmpresa       *string                     `bson:"nro_papeleta_empresa,omitempty" json:"nro_papeleta_empresa,omitempty"`
	Aga                      *string                     `bson:"aga,omitempty" json:"aga,omitempty"`
	RutForw                  *string                     `bson:"rut_forw,omitempty" json:"rut_forw,omitempty"`
	Liberada                 *bool                       `bson:"liberada,omitempty" json:"liberada,omitempty"`
	VAnticipada              *bool                       `bson:"v_anticipada,omitempty" json:"v_anticipada,omitempty"`
	Evento                   string                      `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento              *string                     `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
	Detalles                 *[]PapeletaRecepcionDetalle `bson:"detalles,omitempty" json:"detalles,omitempty"`
}
type PapeletasRecepcion []*PapeletaRecepcion

type PapeletaRecepcionDetalle struct {
	IDMongo       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID            *int64             `bson:"id,omitempty" json:"id,omitempty"`
	CantidadItem  *int64             `bson:"cantidad_item,omitempty" json:"cantidad_item,omitempty"`
	Contenedor    *string            `bson:"contenedor,omitempty" json:"contenedor,omitempty"`
	Estado        *string            `bson:"estado,omitempty" json:"estado,omitempty"`
	IDItem        *int64             `bson:"id_item,omitempty" json:"id_item,omitempty"`
	Marcas        *string            `bson:"marcas,omitempty" json:"marcas,omitempty"`
	Observaciones *string            `bson:"observaciones,omitempty" json:"observaciones,omitempty"`
	Peso          *float64           `bson:"peso,omitempty" json:"peso,omitempty"`
	Sellos        *string            `bson:"sellos,omitempty" json:"sellos,omitempty"`
	Situacion     *string            `bson:"situacion,omitempty" json:"situacion,omitempty"`
	Volumen       *float64           `bson:"volumen,omitempty" json:"volumen,omitempty"`
	NroPapeleta   *string            `bson:"nro_papeleta,omitempty" json:"nro_papeleta,omitempty"`
	TipoBulto     *string            `bson:"tipo_bulto,omitempty" json:"tipo_bulto,omitempty"`
	Chassis       *string            `bson:"chassis,omitempty" json:"chassis,omitempty"`
	Identificador *int64             `bson:"identificador,omitempty" json:"identificador,omitempty"`
	Evento        string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento   *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`
}
type PapeletaRecepcionDetalles []*PapeletaRecepcionDetalle

type PapeletaRecepcionTopic struct {
	Before      *PapeletaRecepcionT `json:"before,omitempty"`
	After       *PapeletaRecepcionT `json:"after,omitempty"`
	Source      *Source             `json:"source,omitempty"`
	Op          string              `json:"op,omitempty"`
	TsMS        *int64              `json:"ts_ms,omitempty"`
	Transaction interface{}         `json:"transaction"`
}
type PapeletaRecepcionT struct {
	NroPapeleta              *string `json:"nro_papeleta,omitempty"`
	Etiqueta                 *string `json:"etiqueta,omitempty"`
	Operacion                *string `json:"operacion,omitempty"`
	Zona                     *string `json:"zona,omitempty"`
	NroBl                    *string `json:"nro_bl,omitempty"`
	PtoOrigen                *string `json:"pto_origen,omitempty"`
	PtoDescarga              *string `json:"pto_descarga,omitempty"`
	PtoDestino               *string `json:"pto_destino,omitempty"`
	Consiganatario           *string `json:"consiganatario,omitempty"`
	DireccionConsignatario   *string `json:"direccion_consignatario,omitempty"`
	RutConsignatario         *string `json:"rut_consignatario,omitempty"`
	Despachador              *string `json:"despachador"`
	ZonaAlmacenaje           *int64  `json:"zona_almacenaje,omitempty"`
	Ubicacion                *int64  `json:"ubicacion,omitempty"`
	GuardaAlmacen            *string `json:"guarda_almacen,omitempty"`
	InicioAlmacenaje         *int64  `json:"inicio_almacenaje,omitempty"`
	ManifiestoNroMftoInterno *string `json:"manifiesto_nro_mfto_interno,omitempty"`
	Forwarder                *string `json:"forwarder,omitempty"`
	Estado                   *string `json:"estado,omitempty"`
	TipoCarga                *string `json:"tipo_carga,omitempty"`
	UsuarioCR                *string `json:"usuario_cr,omitempty"`
	FechaCR                  *int64  `json:"fecha_cr,omitempty"`
	UsuarioUp                *string `json:"usuario_up,omitempty"`
	FechaUp                  *int64  `json:"fecha_up,omitempty"`
	FechaDescon              *int64  `json:"fecha_descon,omitempty"`
	Transbordo               *string `json:"transbordo,omitempty"`
	Tipo                     *string `json:"tipo,omitempty"`
	PrePapeleta              *int64  `json:"pre_papeleta,omitempty"`
	PesoManifestado          *string `json:"peso_manifestado,omitempty"`
	Noty                     *string `json:"noty,omitempty"`
	DireccionNoty            *string `json:"direccion_noty,omitempty"`
	RutNoty                  *string `json:"rut_noty,omitempty"`
	NroPapeletaEmpresa       *string `json:"nro_papeleta_empresa,omitempty"`
	Aga                      *string `json:"aga,omitempty"`
	RutForw                  *string `json:"rut_forw,omitempty"`
	Liberada                 *bool   `json:"liberada,omitempty"`
	MotivoLiberacion         *string `json:"motivo_liberacion,omitempty"`
	VAnticipada              *bool   `json:"v_anticipada,omitempty"`
}
type PapeletaRecepcionDetalleTopic struct {
	Before      *PapeletaRecepcionDetalleT `json:"before,omitempty"`
	After       *PapeletaRecepcionDetalleT `json:"after,omitempty"`
	Source      *Source                    `json:"source,omitempty"`
	Op          string                     `json:"op,omitempty"`
	TsMS        *int64                     `json:"ts_ms,omitempty"`
	Transaction interface{}                `json:"transaction"`
}

type PapeletaRecepcionDetalleT struct {
	ID            *int64  `json:"id,omitempty"`
	CantidadItem  *int64  `json:"cantidad_item,omitempty"`
	Contenedor    *string `json:"contenedor,omitempty"`
	Estado        *string `json:"estado,omitempty"`
	IDItem        *int64  `json:"id_item,omitempty"`
	Marcas        *string `json:"marcas"`
	Observaciones *string `json:"observaciones"`
	Peso          *string `json:"peso,omitempty"`
	Sellos        *string `json:"sellos"`
	Situacion     *string `json:"situacion,omitempty"`
	Tara          *string `json:"tara,omitempty"`
	Volumen       *string `json:"volumen,omitempty"`
	NroPapeleta   *string `json:"nro_papeleta,omitempty"`
	TipoBulto     *string `json:"tipo_bulto,omitempty"`
	Chassis       *string `json:"chassis,omitempty"`
	Identificador *int64  `json:"identificador"`
}
