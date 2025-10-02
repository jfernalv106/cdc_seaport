package model

type PapeletaRecepcionTopic struct {
	Before      *PapeletaRecepcion `json:"before,omitempty"`
	After       *PapeletaRecepcion `json:"after,omitempty"`
	Source      *Source            `json:"source,omitempty"`
	Op          *string            `json:"op,omitempty"`
	TsMS        *int64             `json:"ts_ms,omitempty"`
	Transaction interface{}        `json:"transaction"`
}
type PapeletaRecepcionDetalleTopic struct {
	Before      *PapeletaRecepcionDetalle `json:"before,omitempty"`
	After       *PapeletaRecepcionDetalle `json:"after,omitempty"`
	Source      *Source                   `json:"source,omitempty"`
	Op          *string                   `json:"op,omitempty"`
	TsMS        *int64                    `json:"ts_ms,omitempty"`
	Transaction interface{}               `json:"transaction"`
}

type PapeletaRecepcion struct {
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

type PapeletaRecepcionDetalle struct {
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
