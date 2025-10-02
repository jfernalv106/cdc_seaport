package handler

import (
	"api_auditoria/src/bl"
	"api_auditoria/src/bultos"
	"api_auditoria/src/despacho"
	"api_auditoria/src/expo"
	"api_auditoria/src/factura"
	"api_auditoria/src/manifiesto"
	notacredito "api_auditoria/src/nota_credito"
	"api_auditoria/src/papeleta"
	"api_auditoria/src/visacion"
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func NewRouter(ctx context.Context, endpointPr papeleta.Endpoints, endpointVs visacion.Endpoints, endpointBulto bultos.Endpoints, endpointManifiesto manifiesto.Endpoints, endpointDespacho despacho.Endpoints, endpointsFactura factura.Endpoints, endpointsExpo expo.Endpoints, endpointsNotaCredito notacredito.Endpoints, endpointsBL bl.Endpoints) http.Handler {

	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()
	opts := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(encodeError),
	}
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir cualquier origen
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	api.Handle("/papeleta_recepcion", httptransport.NewServer(
		endpoint.Endpoint(endpointPr.GuardaPapeleta),
		decodeStorePapeletaRecepcion,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/papeleta_recepcion", httptransport.NewServer(
		endpoint.Endpoint(endpointPr.GetPapeletaRecepcion),
		decodeGetPapeleta,
		encodeResponse,
		opts...,
	)).Methods("GET")
	api.Handle("/papeleta_recepcion_detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointPr.GuardaPapeletaRecepcionDetalle),
		decodeStorePapeletaRecepcionDetalle,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/papeleta_recepcion_detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointPr.GetPapeletaRecepcionDetalle),
		decodeGetPapeleta,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**Visacion */
	api.Handle("/visacion", httptransport.NewServer(
		endpoint.Endpoint(endpointVs.GuardarVisacion),
		decodeStoreVisacion,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/visacion-mercancias", httptransport.NewServer(
		endpoint.Endpoint(endpointVs.GuardarVisacionMercancias),
		decodeStoreVisacionMercancias,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/visacion", httptransport.NewServer(
		endpoint.Endpoint(endpointVs.GetVisacion),
		decodeGetVisacion,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**Bulto*/
	api.Handle("/bultos", httptransport.NewServer(
		endpoint.Endpoint(endpointBulto.GuardarBulto),
		decodeStoreBulto,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/bultos", httptransport.NewServer(
		endpoint.Endpoint(endpointBulto.GetBultos),
		decodeGetBulto,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**Manifiesto */
	api.Handle("/manifiesto", httptransport.NewServer(
		endpoint.Endpoint(endpointManifiesto.GuardarManiesto),
		decodeStoreManiesto,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/manifiesto", httptransport.NewServer(
		endpoint.Endpoint(endpointManifiesto.GetManiesto),
		decodeGetManifiesto,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**despacho */
	api.Handle("/despacho", httptransport.NewServer(
		endpoint.Endpoint(endpointDespacho.GuardarDespacho),
		decodeStoreDespacho,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/despacho", httptransport.NewServer(
		endpoint.Endpoint(endpointDespacho.GetDespacho),
		decodeGetDespacho,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**factura */
	api.Handle("/factura", httptransport.NewServer(
		endpoint.Endpoint(endpointsFactura.GuardaFactura),
		decodeStoreFactura,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/factura", httptransport.NewServer(
		endpoint.Endpoint(endpointsFactura.GetFactura),
		decodeGetFactura,
		encodeResponse,
		opts...,
	)).Methods("GET")
	api.Handle("/factura-detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointsFactura.GuardaFacturaDetalle),
		decodeStoreFacturaDetalle,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/factura-detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointsFactura.GetFacturaDetalle),
		decodeGetFactura,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**expo**/
	api.Handle("/papeleta-expo", httptransport.NewServer(
		endpoint.Endpoint(endpointsExpo.GuardaPapeleta),
		decodeStorePapeletaExpo,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/papeleta-expo", httptransport.NewServer(
		endpoint.Endpoint(endpointsExpo.GetPapeleta),
		decodeGetPapeletaExpo,
		encodeResponse,
		opts...,
	)).Methods("GET")
	api.Handle("/papeleta-expo-detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointsExpo.GuardaPapeletaDetalle),
		decodeStorePapeletaExpoDetalle,
		encodeResponse,
		opts...,
	)).Methods("POST")
	api.Handle("/papeleta-expo-detalle", httptransport.NewServer(
		endpoint.Endpoint(endpointsExpo.GetPapeletaDetalle),
		decodeGetPapeletaExpo,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**NotaCredito */
	api.Handle("/nota_credito", httptransport.NewServer(
		endpoint.Endpoint(endpointsNotaCredito.GuardaNotaCredito),
		decodeStoreNotaCredito,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/nota-credito-servicio", httptransport.NewServer(
		endpoint.Endpoint(endpointsNotaCredito.GuardaNotaCreditoServicio),
		decodeStoreNotaCreditoServicio,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/nota_credito", httptransport.NewServer(
		endpoint.Endpoint(endpointsNotaCredito.GetNotaCredito),
		decodeGetNotaCredito,
		encodeResponse,
		opts...,
	)).Methods("GET")

	api.Handle("/nota-credito-servicio", httptransport.NewServer(
		endpoint.Endpoint(endpointsNotaCredito.GetNotaCreditoServicios),
		decodeGetNotaCredito,
		encodeResponse,
		opts...,
	)).Methods("GET")
	/**BL */
	api.Handle("/bl", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBl),
		decodeStoreBL,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GetByNroBl),
		decodeGetBL,
		encodeResponse,
		opts...,
	)).Methods("GET")

	api.Handle("/bl-fecha", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlFecha),
		decodeStoreBlFecha,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-flete", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlFlete),
		decodeStoreBlFlete,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-item", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlItem),
		decodeStoreBlItem,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-item-imo", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlItemImo),
		decodeStoreBlItemImo,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-item-contenedor", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlItemContenedor),
		decodeStoreBlItemContenedor,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-item-contenedor-imo", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlItemContenedorImo),
		decodeStoreBlItemContenedorImo,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-item-contenedor-sello", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlItemContenedorSello),
		decodeStoreBlItemContenedorSello,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-locacion", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlLocacion),
		decodeStoreBlLocacion,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-observacion", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlObservacion),
		decodeStoreBlObservacion,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-participante", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlParticipante),
		decodeStoreBlParticipante,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-referencia", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlReferencia),
		decodeStoreBlReferencia,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-transbordo", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlTransbordo),
		decodeStoreBlTransbordo,
		encodeResponse,
		opts...,
	)).Methods("POST")

	api.Handle("/bl-transporte", httptransport.NewServer(
		endpoint.Endpoint(endpointsBL.GuardarBlTransporte),
		decodeStoreBlTransporte,
		encodeResponse,
		opts...,
	)).Methods("POST")

	return corsHandler.Handler(api)
}
