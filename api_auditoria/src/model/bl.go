package model

import "go.mongodb.org/mongo-driver/bson/primitive"

/**BL**/
type BlTopic struct {
	Before      *BlT        `json:"before,omitempty"`
	After       *BlT        `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type BlT struct {
	ID                       *int64  `json:"id,omitempty"`
	NroBl                    *string `json:"nro_bl,omitempty"`
	ManifiestoNroMftoInterno *string `json:"manifiesto_nro_mfto_interno,omitempty"`
	TipoAccion               *string `json:"tipo_accion,omitempty"`
	TipoBl                   *string `json:"tipo_bl"`
	FechaRecepcion           *int64  `json:"fecha_recepcion"`
	Servicio                 *string `json:"servicio,omitempty"`
	TipoServicio             *string `json:"tipo_servicio,omitempty"`
	CondTransporte           *string `json:"cond_transporte,omitempty"`
	TotalBultos              *int64  `json:"total_bultos,omitempty"`
	TotalPeso                *string `json:"total_peso,omitempty"`
	UnidadPeso               *string `json:"unidad_peso,omitempty"`
	TotalVolumen             *string `json:"total_volumen"`
	UnidadVolumen            *string `json:"unidad_volumen"`
	TotalItem                *int64  `json:"total_item,omitempty"`
	UsuarioCR                *string `json:"usuario_cr,omitempty"`
	FechaCR                  *int64  `json:"fecha_cr,omitempty"`
	UsuarioUp                *string `json:"usuario_up"`
	FechaUp                  *int64  `json:"fecha_up"`
	BlAforo                  *string `json:"bl_aforo"`
	BlSag                    *string `json:"bl_sag"`
	NroBlOriginal            *string `json:"nro_bl_original,omitempty"`
	MotivoAnulacion          *string `json:"motivo_anulacion"`
	Estado                   *string `json:"estado,omitempty"`
	Etiqueta                 *string `json:"etiqueta,omitempty"`
	Aforo                    *string `json:"aforo"`
}

/**BL Fecha**/
type BlFechaTopic struct {
	Before      *BlFechaT   `json:"before,omitempty"`
	After       *BlFechaT   `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type BlFechaT struct {
	ID            *int64  `json:"id,omitempty"`
	Nombre        *string `json:"nombre,omitempty"`
	Valor         *int64  `json:"valor,omitempty"`
	BlNroBl       *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Flete**/
type BlFleteTopic struct {
	Before      *BlFleteT   `json:"before,omitempty"`
	After       *BlFleteT   `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type BlFleteT struct {
	Descripcion   *string `json:"descripcion"`
	Tipo          *string `json:"tipo,omitempty"`
	BlNroBl       *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Item_Imo**/
type BlItemImoTopic struct {
	Before      *BlItemImoT `json:"before,omitempty"`
	After       *BlItemImoT `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type BlItemImoT struct {
	ID            *int64  `json:"id,omitempty"`
	ClaseImo      *string `json:"clase_imo,omitempty"`
	NumeroImo     *int64  `json:"numero_imo,omitempty"`
	BlItemID      *int64  `json:"bl_item_id,omitempty"`
	FechaTraspaso *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Item**/
type BlItemTopic struct {
	Before      *BlItemT    `json:"before,omitempty"`
	After       *BlItemT    `json:"after,omitempty"`
	Source      *Source     `json:"source,omitempty"`
	Op          *string     `json:"op,omitempty"`
	TsMS        *int64      `json:"ts_ms,omitempty"`
	Transaction interface{} `json:"transaction"`
}

type BlItemT struct {
	ID             *int64  `json:"id,omitempty"`
	BlNroBl        *int64  `json:"bl_nro_bl,omitempty"`
	NroItem        *int64  `json:"nro_item,omitempty"`
	Marcas         *string `json:"marcas,omitempty"`
	CargaPeligrosa *string `json:"carga_peligrosa,omitempty"`
	TipoBulto      *string `json:"tipo_bulto,omitempty"`
	Descripcion    *string `json:"descripcion,omitempty"`
	Cantidad       *int64  `json:"cantidad,omitempty"`
	PesoBruto      *string `json:"peso_bruto,omitempty"`
	UnidadPeso     *string `json:"unidad_peso,omitempty"`
	Volumen        *string `json:"volumen,omitempty"`
	UnidadVolumen  *string `json:"unidad_volumen,omitempty"`
	Observaciones  *string `json:"observaciones"`
	CargaCnt       *string `json:"carga_cnt"`
	FechaTraspaso  *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Item Contenedor**/
type BlItemContenedorTopic struct {
	Before      *BlItemContenedorT `json:"before,omitempty"`
	After       *BlItemContenedorT `json:"after,omitempty"`
	Source      *Source            `json:"source,omitempty"`
	Op          *string            `json:"op,omitempty"`
	TsMS        *int64             `json:"ts_ms,omitempty"`
	Transaction interface{}        `json:"transaction"`
}

type BlItemContenedorT struct {
	ID             *int64  `json:"id,omitempty"`
	BlItemID       *int64  `json:"bl_item_id,omitempty"`
	Sigla          *string `json:"sigla,omitempty"`
	Numero         *string `json:"numero,omitempty"`
	Digito         *string `json:"digito,omitempty"`
	TipoCnt        *string `json:"tipo_cnt,omitempty"`
	CntSo          *string `json:"cnt_so"`
	Peso           *string `json:"peso,omitempty"`
	ValorIDOp      *string `json:"valor_id_op"`
	NombreOperador *string `json:"nombre_operador,omitempty"`
	Estado         *string `json:"estado,omitempty"`
	FechaTraspaso  *int64  `json:"fecha_traspaso,omitempty"`
}

/**Bl Item Contenedor Imo**/
type BlItemContenedorImoTopic struct {
	Before      *BlItemContenedorImoT `json:"before,omitempty"`
	After       *BlItemContenedorImoT `json:"after,omitempty"`
	Source      *Source               `json:"source,omitempty"`
	Op          *string               `json:"op,omitempty"`
	TsMS        *int64                `json:"ts_ms,omitempty"`
	Transaction interface{}           `json:"transaction"`
}

type BlItemContenedorImoT struct {
	ID                 *int64  `json:"id,omitempty"`
	ClaseImo           *string `json:"clase_imo,omitempty"`
	NumeroImo          *int64  `json:"numero_imo,omitempty"`
	BlItemContenedorID *int64  `json:"bl_item_contenedor_id,omitempty"`
	FechaTraspaso      *int64  `json:"fecha_traspaso,omitempty"`
}

/**Bl Item Contenedor Sello**/

type BlItemContenedorSelloTopic struct {
	Before      *BlItemContenedorSelloT `json:"before,omitempty"`
	After       *BlItemContenedorSelloT `json:"after,omitempty"`
	Source      *Source                 `json:"source,omitempty"`
	Op          *string                 `json:"op,omitempty"`
	TsMS        *int64                  `json:"ts_ms,omitempty"`
	Transaction interface{}             `json:"transaction"`
}

type BlItemContenedorSelloT struct {
	ID                 *int64  `json:"id,omitempty"`
	Codigo             *string `json:"codigo,omitempty"`
	Emisor             *string `json:"emisor,omitempty"`
	Numero             *string `json:"numero,omitempty"`
	BlItemContenedorID *int64  `json:"bl_item_contenedor_id,omitempty"`
	FechaTraspaso      *int64  `json:"fecha_traspaso,omitempty"`
}

/**Bl Locacion**/

type BlLocacionTopic struct {
	Before      *BlLocacionT `json:"before,omitempty"`
	After       *BlLocacionT `json:"after,omitempty"`
	Source      *Source      `json:"source,omitempty"`
	Op          *string      `json:"op,omitempty"`
	TsMS        *int64       `json:"ts_ms,omitempty"`
	Transaction interface{}  `json:"transaction"`
}

type BlLocacionT struct {
	ID            *int64  `json:"id,omitempty"`
	Codigo        *string `json:"codigo,omitempty"`
	Descripcion   *string `json:"descripcion,omitempty"`
	Nombre        *string `json:"nombre,omitempty"`
	BlNroBl       *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso *int64  `json:"fecha_traspaso,omitempty"`
}

/**Bl Observacion**/
type BlObservacionTopic struct {
	Before      *BlObservacionT `json:"before,omitempty"`
	After       *BlObservacionT `json:"after,omitempty"`
	Source      *Source         `json:"source,omitempty"`
	Op          *string         `json:"op,omitempty"`
	TsMS        *int64          `json:"ts_ms,omitempty"`
	Transaction interface{}     `json:"transaction"`
}

type BlObservacionT struct {
	ID            *int64  `json:"id,omitempty"`
	Contenido     *string `json:"contenido,omitempty"`
	Nombre        *string `json:"nombre,omitempty"`
	BlNroBl       *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Participante**/
type BlParticipanteTopic struct {
	Before      *BlParticipanteT `json:"before,omitempty"`
	After       *BlParticipanteT `json:"after,omitempty"`
	Source      *Source          `json:"source,omitempty"`
	Op          *string          `json:"op,omitempty"`
	TsMS        *int64           `json:"ts_ms,omitempty"`
	Transaction interface{}      `json:"transaction"`
}

type BlParticipanteT struct {
	ID               *int64  `json:"id,omitempty"`
	CodigoAlmacen    *string `json:"codigo_almacen"`
	CodigoPais       *string `json:"codigo_pais"`
	Comuna           *string `json:"comuna,omitempty"`
	Direccion        *string `json:"direccion,omitempty"`
	Email            *string `json:"email"`
	NacionID         *string `json:"nacion_id"`
	Nombres          *string `json:"nombres,omitempty"`
	Telefono         *string `json:"telefono,omitempty"`
	TipoID           *string `json:"tipo_id"`
	TipoParticipante *string `json:"tipo_participante,omitempty"`
	ValorID          *string `json:"valor_id"`
	BlNroBl          *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso    *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Referencia**/
type BlReferenciaTopic struct {
	Before      *BlReferenciaT `json:"before,omitempty"`
	After       *BlReferenciaT `json:"after,omitempty"`
	Source      *Source        `json:"source,omitempty"`
	Op          *string        `json:"op,omitempty"`
	TsMS        *int64         `json:"ts_ms,omitempty"`
	Transaction interface{}    `json:"transaction"`
}

type BlReferenciaT struct {
	ID             *int64  `json:"id,omitempty"`
	Emisor         *string `json:"emisor,omitempty"`
	Fecha          *int64  `json:"fecha,omitempty"`
	NacIDEmisor    *string `json:"nac_id_emisor,omitempty"`
	Numero         *string `json:"numero,omitempty"`
	TipoDocumento  *string `json:"tipo_documento,omitempty"`
	TipoIDEmisor   *string `json:"tipo_id_emisor,omitempty"`
	TipoReferencia *string `json:"tipo_referencia,omitempty"`
	ValorIDEmisor  *string `json:"valor_id_emisor,omitempty"`
	BlNroBl        *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso  *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Transbordo**/
type BlTransbordoTopic struct {
	Before      *BlTransbordoT `json:"before,omitempty"`
	After       *BlTransbordoT `json:"after,omitempty"`
	Source      *Source        `json:"source,omitempty"`
	Op          *string        `json:"op,omitempty"`
	TsMS        *int64         `json:"ts_ms,omitempty"`
	Transaction interface{}    `json:"transaction"`
}

type BlTransbordoT struct {
	ID               *int64  `json:"id,omitempty"`
	CodigoLugar      *string `json:"codigo_lugar,omitempty"`
	DescripcionLugar *string `json:"descripcion_lugar,omitempty"`
	FechaArribo      *int64  `json:"fecha_arribo"`
	BlNroBl          *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso    *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Transporte**/
type BlTransporteTopic struct {
	Before      *BlTransporteT `json:"before,omitempty"`
	After       *BlTransporteT `json:"after,omitempty"`
	Source      *Source        `json:"source,omitempty"`
	Op          *string        `json:"op,omitempty"`
	TsMS        *int64         `json:"ts_ms,omitempty"`
	Transaction interface{}    `json:"transaction"`
}

type BlTransporteT struct {
	NombreNave       *string `json:"nombre_nave,omitempty"`
	SentidoOperacion *string `json:"sentido_operacion,omitempty"`
	BlNroBl          *int64  `json:"bl_nro_bl,omitempty"`
	FechaTraspaso    *int64  `json:"fecha_traspaso,omitempty"`
}

/**BL Modelo**/
type BL struct {
	IDMongo                  primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	ID                       *int64             `bson:"id,omitempty" json:"id,omitempty"`
	NroBl                    *string            `bson:"nro_bl,omitempty" json:"nro_bl,omitempty"`
	ManifiestoNroMftoInterno *string            `bson:"manifiesto_nro_mfto_interno,omitempty" json:"manifiesto_nro_mfto_interno,omitempty"`
	TipoAccion               *string            `bson:"tipo_accion,omitempty" json:"tipo_accion,omitempty"`
	TipoBl                   *string            `bson:"tipo_bl,omitempty" json:"tipo_bl,omitempty"`
	FechaRecepcion           *string            `bson:"fecha_recepcion,omitempty" json:"fecha_recepcion,omitempty"`
	Servicio                 *string            `bson:"servicio,omitempty" json:"servicio,omitempty"`
	TipoServicio             *string            `bson:"tipo_servicio,omitempty" json:"tipo_servicio,omitempty"`
	CondTransporte           *string            `bson:"cond_transporte,omitempty" json:"cond_transporte,omitempty"`
	TotalBultos              *int64             `bson:"total_bultos,omitempty" json:"total_bultos,omitempty"`
	TotalPeso                *float64           `bson:"total_peso,omitempty" json:"total_peso,omitempty"`
	UnidadPeso               *string            `bson:"unidad_peso,omitempty" json:"unidad_peso,omitempty"`
	TotalVolumen             *float64           `bson:"total_volumen,omitempty" json:"total_volumen,omitempty"`
	UnidadVolumen            *string            `bson:"unidad_volumen,omitempty" json:"unidad_volumen,omitempty"`
	TotalItem                *int64             `bson:"total_item,omitempty" json:"total_item,omitempty"`
	UsuarioCR                *string            `bson:"usuario_cr,omitempty" json:"usuario_cr,omitempty"`
	FechaCR                  *string            `bson:"fecha_cr,omitempty" json:"fecha_cr,omitempty"`
	UsuarioUp                *string            `bson:"usuario_up,omitempty" json:"usuario_up,omitempty"`
	FechaUp                  *string            `bson:"fecha_up,omitempty" json:"fecha_up,omitempty"`
	BlAforo                  *string            `bson:"bl_aforo,omitempty" json:"bl_aforo,omitempty"`
	BlSag                    *string            `bson:"bl_sag,omitempty" json:"bl_sag,omitempty"`
	NroBlOriginal            *string            `bson:"nro_bl_original,omitempty" json:"nro_bl_original,omitempty"`
	MotivoAnulacion          *string            `bson:"motivo_anulacion,omitempty" json:"motivo_anulacion,omitempty"`
	Estado                   *string            `bson:"estado,omitempty" json:"estado,omitempty"`
	Etiqueta                 *string            `bson:"etiqueta,omitempty" json:"etiqueta,omitempty"`
	Aforo                    *string            `bson:"aforo,omitempty" json:"aforo,omitempty"`
	Evento                   string             `bson:"evento,omitempty" json:"evento,omitempty"`
	FechaEvento              *string            `bson:"fecha_evento,omitempty" json:"fecha_evento,omitempty"`

	BlFechas        *[]BlFecha        `bson:"bl_fechas,omitempty" json:"bl_fechas,omitempty"`
	BlFletes        *[]BlFlete        `bson:"bl_fletes,omitempty" json:"bl_fletes,omitempty"`
	BlItems         *[]BlItem         `bson:"bl_items,omitempty" json:"bl_items,omitempty"`
	BlParticipantes *[]BlParticipante `bson:"bl_participantes,omitempty" json:"bl_participantes,omitempty"`
	BlReferencias   *[]BlReferencia   `bson:"bl_referencias,omitempty" json:"bl_referencias,omitempty"`
	BlTransbordos   *[]BlTransbordo   `bson:"bl_transbordos,omitempty" json:"bl_transbordos,omitempty"`
	BlTransportes   *[]BlTransporte   `bson:"bl_transportes,omitempty" json:"bl_transportes,omitempty"`
	BlObservaciones *[]BlObservacion  `bson:"bl_observaciones,omitempty" json:"bl_observaciones,omitempty"`
	BlLocaciones    *[]BlLocacion     `bson:"bl_locaciones,omitempty" json:"bl_locaciones,omitempty"`
}

type BlFecha struct {
	IDMongo       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID            *int64             `json:"id,omitempty" bson:"id,omitempty"`
	Nombre        *string            `json:"nombre,omitempty" bson:"nombre,omitempty"`
	Valor         *string            `json:"valor,omitempty" bson:"valor,omitempty"`
	BlNroBl       *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento        string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento   *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL FLETE ========== **/
type BlFlete struct {
	IDMongo       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Descripcion   *string            `json:"descripcion" bson:"descripcion"`
	Tipo          *string            `json:"tipo,omitempty" bson:"tipo,omitempty"`
	BlNroBl       *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento        string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento   *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL ITEM IMO ========== **/
type BlItemImo struct {
	IDMongo       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID            *int64             `json:"id,omitempty" bson:"id,omitempty"`
	ClaseImo      *string            `json:"clase_imo,omitempty" bson:"clase_imo,omitempty"`
	NumeroImo     *int64             `json:"numero_imo,omitempty" bson:"numero_imo,omitempty"`
	BlItemID      *int64             `json:"bl_item_id,omitempty" bson:"bl_item_id,omitempty"`
	Evento        string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento   *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL ITEM ========== **/
type BlItem struct {
	IDMongo            primitive.ObjectID  `json:"_id,omitempty" bson:"_id,omitempty"`
	ID                 *int64              `json:"id,omitempty" bson:"id,omitempty"`
	BlNroBl            *int64              `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	NroItem            *int64              `json:"nro_item,omitempty" bson:"nro_item,omitempty"`
	Marcas             *string             `json:"marcas,omitempty" bson:"marcas,omitempty"`
	CargaPeligrosa     *string             `json:"carga_peligrosa,omitempty" bson:"carga_peligrosa,omitempty"`
	TipoBulto          *string             `json:"tipo_bulto,omitempty" bson:"tipo_bulto,omitempty"`
	Descripcion        *string             `json:"descripcion,omitempty" bson:"descripcion,omitempty"`
	Cantidad           *int64              `json:"cantidad,omitempty" bson:"cantidad,omitempty"`
	PesoBruto          *float64            `json:"peso_bruto,omitempty" bson:"peso_bruto,omitempty"`
	UnidadPeso         *string             `json:"unidad_peso,omitempty" bson:"unidad_peso,omitempty"`
	Volumen            *float64            `json:"volumen,omitempty" bson:"volumen,omitempty"`
	UnidadVolumen      *string             `json:"unidad_volumen,omitempty" bson:"unidad_volumen,omitempty"`
	Observaciones      *string             `json:"observaciones" bson:"observaciones"`
	CargaCnt           *string             `json:"carga_cnt" bson:"carga_cnt"`
	Evento             string              `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento        *string             `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	BlItemImos         *[]BlItemImo        `json:"bl_item_imos,omitempty" bson:"bl_item_imos,omitempty"`
	BlItemContenedores *[]BlItemContenedor `json:"bl_item_contenedores,omitempty" bson:"bl_item_contenedores,omitempty"`
	FechaTraspaso      *string             `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL ITEM CONTENEDOR ========== **/
type BlItemContenedor struct {
	IDMongo                primitive.ObjectID       `json:"_id,omitempty" bson:"_id,omitempty"`
	ID                     *int64                   `json:"id,omitempty" bson:"id,omitempty"`
	BlItemID               *int64                   `json:"bl_item_id,omitempty" bson:"bl_item_id,omitempty"`
	Sigla                  *string                  `json:"sigla,omitempty" bson:"sigla,omitempty"`
	Numero                 *string                  `json:"numero,omitempty" bson:"numero,omitempty"`
	Digito                 *string                  `json:"digito,omitempty" bson:"digito,omitempty"`
	TipoCnt                *string                  `json:"tipo_cnt,omitempty" bson:"tipo_cnt,omitempty"`
	CntSo                  *string                  `json:"cnt_so" bson:"cnt_so"`
	Peso                   *float64                 `json:"peso,omitempty" bson:"peso,omitempty"`
	ValorIDOp              *string                  `json:"valor_id_op" bson:"valor_id_op"`
	NombreOperador         *string                  `json:"nombre_operador,omitempty" bson:"nombre_operador,omitempty"`
	Estado                 *string                  `json:"estado,omitempty" bson:"estado,omitempty"`
	Evento                 string                   `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento            *string                  `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	BlItemContenedorImos   *[]BlItemContenedorImo   `json:"bl_item_contenedor_imos,omitempty" bson:"bl_item_contenedor_imos,omitempty"`
	BlItemContenedorSellos *[]BlItemContenedorSello `json:"bl_item_contenedor_sellos,omitempty" bson:"bl_item_contenedor_sellos,omitempty"`
	FechaTraspaso          *string                  `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL ITEM CONTENEDOR IMO ========== **/
type BlItemContenedorImo struct {
	IDMongo            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID                 *int64             `json:"id,omitempty" bson:"id,omitempty"`
	ClaseImo           *string            `json:"clase_imo,omitempty" bson:"clase_imo,omitempty"`
	NumeroImo          *int64             `json:"numero_imo,omitempty" bson:"numero_imo,omitempty"`
	BlItemContenedorID *int64             `json:"bl_item_contenedor_id,omitempty" bson:"bl_item_contenedor_id,omitempty"`
	Evento             string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento        *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso      *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL ITEM CONTENEDOR SELLO ========== **/
type BlItemContenedorSello struct {
	IDMongo            primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID                 *int64             `json:"id,omitempty" bson:"id,omitempty"`
	Codigo             *string            `json:"codigo,omitempty" bson:"codigo,omitempty"`
	Emisor             *string            `json:"emisor,omitempty" bson:"emisor,omitempty"`
	Numero             *string            `json:"numero,omitempty" bson:"numero,omitempty"`
	BlItemContenedorID *int64             `json:"bl_item_contenedor_id,omitempty" bson:"bl_item_contenedor_id,omitempty"`
	Evento             string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento        *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso      *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL LOCACION ========== **/
type BlLocacion struct {
	IDMongo       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID            *int64             `json:"id,omitempty" bson:"id,omitempty"`
	Codigo        *string            `json:"codigo,omitempty" bson:"codigo,omitempty"`
	Descripcion   *string            `json:"descripcion,omitempty" bson:"descripcion,omitempty"`
	Nombre        *string            `json:"nombre,omitempty" bson:"nombre,omitempty"`
	BlNroBl       *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento        string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento   *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL OBSERVACION ========== **/
type BlObservacion struct {
	IDMongo       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID            *int64             `json:"id,omitempty" bson:"id,omitempty"`
	Contenido     *string            `json:"contenido,omitempty" bson:"contenido,omitempty"`
	Nombre        *string            `json:"nombre,omitempty" bson:"nombre,omitempty"`
	BlNroBl       *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento        string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento   *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL PARTICIPANTE ========== **/
type BlParticipante struct {
	IDMongo          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID               *int64             `json:"id,omitempty" bson:"id,omitempty"`
	CodigoAlmacen    *string            `json:"codigo_almacen" bson:"codigo_almacen"`
	CodigoPais       *string            `json:"codigo_pais" bson:"codigo_pais"`
	Comuna           *string            `json:"comuna,omitempty" bson:"comuna,omitempty"`
	Direccion        *string            `json:"direccion,omitempty" bson:"direccion,omitempty"`
	Email            *string            `json:"email" bson:"email"`
	NacionID         *string            `json:"nacion_id" bson:"nacion_id"`
	Nombres          *string            `json:"nombres,omitempty" bson:"nombres,omitempty"`
	Telefono         *string            `json:"telefono,omitempty" bson:"telefono,omitempty"`
	TipoID           *string            `json:"tipo_id" bson:"tipo_id"`
	TipoParticipante *string            `json:"tipo_participante,omitempty" bson:"tipo_participante,omitempty"`
	ValorID          *string            `json:"valor_id" bson:"valor_id"`
	BlNroBl          *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento           string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento      *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso    *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL REFERENCIA ========== **/
type BlReferencia struct {
	IDMongo        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID             *int64             `json:"id,omitempty" bson:"id,omitempty"`
	Emisor         *string            `json:"emisor,omitempty" bson:"emisor,omitempty"`
	Fecha          *string            `json:"fecha,omitempty" bson:"fecha,omitempty"`
	NacIDEmisor    *string            `json:"nac_id_emisor,omitempty" bson:"nac_id_emisor,omitempty"`
	Numero         *string            `json:"numero,omitempty" bson:"numero,omitempty"`
	TipoDocumento  *string            `json:"tipo_documento,omitempty" bson:"tipo_documento,omitempty"`
	TipoIDEmisor   *string            `json:"tipo_id_emisor,omitempty" bson:"tipo_id_emisor,omitempty"`
	TipoReferencia *string            `json:"tipo_referencia,omitempty" bson:"tipo_referencia,omitempty"`
	ValorIDEmisor  *string            `json:"valor_id_emisor,omitempty" bson:"valor_id_emisor,omitempty"`
	BlNroBl        *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento         string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento    *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso  *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL TRANSBORDO ========== **/
type BlTransbordo struct {
	IDMongo          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ID               *int64             `json:"id,omitempty" bson:"id,omitempty"`
	CodigoLugar      *string            `json:"codigo_lugar,omitempty" bson:"codigo_lugar,omitempty"`
	DescripcionLugar *string            `json:"descripcion_lugar,omitempty" bson:"descripcion_lugar,omitempty"`
	FechaArribo      *string            `json:"fecha_arribo" bson:"fecha_arribo"`
	BlNroBl          *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento           string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento      *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso    *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}

/** ========== BL TRANSPORTE ========== **/
type BlTransporte struct {
	IDMongo          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	NombreNave       *string            `json:"nombre_nave,omitempty" bson:"nombre_nave,omitempty"`
	SentidoOperacion *string            `json:"sentido_operacion,omitempty" bson:"sentido_operacion,omitempty"`
	BlNroBl          *int64             `json:"bl_nro_bl,omitempty" bson:"bl_nro_bl,omitempty"`
	Evento           string             `json:"evento,omitempty" bson:"evento,omitempty"`
	FechaEvento      *string            `json:"fecha_evento,omitempty" bson:"fecha_evento,omitempty"`
	FechaTraspaso    *string            `json:"fecha_traspaso,omitempty" bson:"fecha_traspaso,omitempty"`
}
