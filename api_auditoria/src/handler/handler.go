package handler

import (
	"api_auditoria/src/bl"
	"api_auditoria/src/bultos"
	"api_auditoria/src/despacho"
	"api_auditoria/src/expo"
	"api_auditoria/src/factura"
	"api_auditoria/src/manifiesto"
	"api_auditoria/src/model"
	notacredito "api_auditoria/src/nota_credito"
	"api_auditoria/src/papeleta"
	"api_auditoria/src/visacion"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/jfernalv106/response_go/response"
)

func encodeResponse(ctx context.Context, w http.ResponseWriter, resp interface{}) error {
	r := resp.(response.Response)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(r.StatusCode())
	return json.NewEncoder(w).Encode(resp)
}
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	resp := err.(response.Response)

	w.WriteHeader(resp.StatusCode())

	_ = json.NewEncoder(w).Encode(resp)
}

/**Papeleta**/
func decodeStorePapeletaRecepcion(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.PapeletaRecepcionTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeStorePapeletaRecepcionDetalle(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.PapeletaRecepcionDetalleTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeGetPapeleta(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	nroPapeleta := v.Get("nro_papeleta")
	bl := v.Get("bl")
	manifiesto := v.Get("manifiesto")

	req := papeleta.GetPapeletaRequest{
		NroPapeleta: &nroPapeleta,
		Bl:          &bl,
		Manifiesto:  &manifiesto,
	}

	return req, nil
}

/*Visacion*/
func decodeStoreVisacion(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.VisacionTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeStoreVisacionMercancias(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.MercanciasDespachadasTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}

func decodeGetVisacion(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	nroPapeleta := v.Get("nro_papeleta")
	fmt.Println("NroPapeleta:", nroPapeleta)
	req := visacion.GetVisacionRequest{
		NroPapeleta: &nroPapeleta,
	}

	return req, nil
}

/*Bultos*/
func decodeStoreBulto(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BultoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeGetBulto(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	cod := v.Get("cod")
	fmt.Println("cod:", cod)
	req := bultos.GetBultoRequest{
		Cod: cod,
	}

	return req, nil
}

/*maniesto*/

func decodeStoreManiesto(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.ManifiestoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeGetManifiesto(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	nro := v.Get("nro")
	num, err := strconv.ParseInt(nro, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("no se pudo convertir '%s' a int64: %v", nro, err)
	}
	req := manifiesto.GetManifiestoRequest{
		Nro: num,
	}

	return req, nil
}

/*despacho*/

func decodeStoreDespacho(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.DespachoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeGetDespacho(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	id := v.Get("id")
	num, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("no se pudo convertir '%s' a int64: %v", id, err)
	}
	req := despacho.GetDespachoRequest{
		ID: &num,
	}

	return req, nil
}

/**Factura**/
func decodeStoreFactura(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.FacturaTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeStoreFacturaDetalle(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.FacturaDetalleTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}

	return req, nil
}
func decodeGetFactura(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	id := v.Get("id")
	folioStr := v.Get("folio")
	manifiesto := v.Get("manifiesto")
	var folioNum *int64
	var idFactura *int64
	var man *string

	if id != "" {
		num, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("no se pudo convertir '%s' a int64: %v", id, err)
		}
		idFactura = &num
	}
	if folioStr != "" {
		f, err := strconv.ParseInt(folioStr, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("no se pudo convertir 'folio' ('%s') a int64: %v", folioStr, err)
		}
		folioNum = &f
	}
	if manifiesto != "" {
		man = &manifiesto
	}

	req := factura.GetFacturaRequest{
		Id:         idFactura,
		Folio:      folioNum,
		Manifiesto: man,
	}

	return req, nil
}

/**Expo**/
func decodeStorePapeletaExpo(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.PapeletaExpoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}
func decodeStorePapeletaExpoDetalle(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.PapeletaExpoDetalleTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}
func decodeGetPapeletaExpo(_ context.Context, r *http.Request) (interface{}, error) {

	v := r.URL.Query()

	i := v.Get("id")
	b := v.Get("booking")
	p := v.Get("nro_papeleta")

	var id *int64
	var booking *string
	var papeleta *string

	if i != "" {
		num, err := strconv.ParseInt(i, 10, 64)
		if err != nil {
			return nil, fmt.Errorf("no se pudo convertir '%s' a int64: %v", id, err)
		}
		id = &num
	}

	if b != "" {
		booking = &b
	}

	if p != "" {
		papeleta = &p
	}

	req := expo.GetExpoRequest{
		ID:       id,
		Booking:  booking,
		Papeleta: papeleta,
	}

	return req, nil
}

/*NotaCredito*/
func decodeStoreNotaCredito(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.NotaCreditoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreNotaCreditoServicio(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.NotaCreditoServTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeGetNotaCredito(_ context.Context, r *http.Request) (interface{}, error) {
	v := r.URL.Query()

	var idNroNc *int64
	var idFolNc *int64
	var idFactura *int64
	var idFolio *int64

	if val := v.Get("id_nro_nc"); val != "" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			idNroNc = &num
		}
	}
	if val := v.Get("id_fol_nc"); val != "" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			idFolNc = &num
		}
	}
	if val := v.Get("id_factura"); val != "" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			idFactura = &num
		}
	}
	if val := v.Get("id_folio"); val != "" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			idFolio = &num
		}
	}

	req := notacredito.GetNotaCreditoRequest{
		IDNroNc:   idNroNc,
		IDFolNc:   idFolNc,
		IDFactura: idFactura,
		IDFolio:   idFolio,
	}
	return req, nil
}

/* BL */

func decodeStoreBL(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlFecha(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlFechaTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlFlete(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlFleteTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlItem(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlItemTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlItemImo(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlItemImoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlItemContenedor(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlItemContenedorTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlItemContenedorImo(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlItemContenedorImoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlItemContenedorSello(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlItemContenedorSelloTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlLocacion(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlLocacionTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlObservacion(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlObservacionTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlParticipante(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlParticipanteTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlReferencia(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlReferenciaTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlTransbordo(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlTransbordoTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeStoreBlTransporte(_ context.Context, r *http.Request) (interface{}, error) {
	var req model.BlTransporteTopic
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return nil, response.BadRequest(fmt.Sprintf("invalid request format: '%v'", err.Error()))
	}
	return req, nil
}

func decodeGetBL(_ context.Context, r *http.Request) (interface{}, error) {
	v := r.URL.Query()
	var id *int64
	var nroBl *string
	var manifiesto *string

	if val := v.Get("id"); val != "" {
		num, err := strconv.ParseInt(val, 10, 64)
		if err == nil {
			id = &num
		}
	}
	if val := v.Get("nro_bl"); val != "" {
		nroBl = &val
	}
	if val := v.Get("manifiesto"); val != "" {
		manifiesto = &val
	}

	req := bl.GetByNroBlRequest{
		ID:         id,
		NroBl:      nroBl,
		Manifiesto: manifiesto,
	}
	return req, nil
}
