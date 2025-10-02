package factura

import (
	"api_auditoria/src/model"
	"context"
	"fmt"
	"log"

	"github.com/jfernalv106/response_go/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		GuardaFactura        Controller
		GuardaFacturaDetalle Controller
		GetFactura           Controller
		GetFacturaDetalle    Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetFacturaRequest struct {
		Id         *int64  `json:"id,omitempty"`
		Folio      *int64  `json:"folio,omitempty"`
		Manifiesto *string `json:"manifiesto,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardaFactura:        GuardaFactura(s),
		GuardaFacturaDetalle: GuardaFacturaDetalle(s),
		GetFactura:           GetFactura(s),
		GetFacturaDetalle:    GetFacturaDetalle(s),
	}
}
func GuardaFactura(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.FacturaTopic)

		var id int64
		if req.After != nil && req.After.ID != nil {
			id = *req.After.ID
		} else {
			id = *req.Before.ID
		}

		if req.Op == "r" || req.Op == "c" {
			_, err := s.GuardarFactura(*ConvertToFactura(&req, nil))
			if err != nil {
				log.Println("Error al guardar la factura:", err)
				return response.BadRequest("Error al guardar la factura"), nil
			}

		} else {
			fc, _ := s.GetUltimaFacturaPorEvento(id)
			bs, err := s.GuardarFactura(*ConvertToFactura(&req, fc))
			if err != nil || bs == nil {
				log.Println("Error al guardar la factura:", req.Op, err)
				return response.BadRequest("Error al guardar la factura"), nil
			}
		}
		fmt.Println("Factura guardada exitosamente:", req.Op)

		return response.OK("Factura registrada", req, nil), nil

	}
}
func GuardaFacturaDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.FacturaDetalleTopic)
		fmt.Println("Guardando Detalle Factura:", req)
		p, err := s.GuardarFacturaDetalle(*ConvertToFacturaDetalle(&req))
		if err != nil {
			log.Println("Error al guardar el detalle de la factura:", err)
			fmt.Println("Error al guardar el detalle de la factura:", err)
			return response.BadRequest("Error al guardar el detalle de la factura"), nil
		}
		return response.OK("Detalle Factura registrada", p, nil), nil

	}
}
func GetFactura(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFacturaRequest)

		fmt.Println("Obteniendo Factura:", req.Folio)
		papeletas, err := s.GetFactura(req.Folio, req.Manifiesto)
		if err != nil {
			log.Println("Error al obtener la factura:", err)
			return response.BadRequest("Error al obtener la factura"), nil
		}

		fmt.Println("Factura obtenida exitosamente:", req.Folio)
		return response.OK("Factura obtenida", &papeletas, nil), nil
	}
}
func GetFacturaDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetFacturaRequest)

		fmt.Println("Obteniendo Factura Detalle:", req.Id)
		detalles, err := s.GetFacturaDetalle(req.Id)

		if err != nil {
			log.Println("Error al obtener la Factura Detalle:", err)
			return response.BadRequest("Error al obtener la factura detalle"), nil
		}

		fmt.Println("Factura Detalle obtenida exitosamente:", req.Id)
		return response.OK("Detalles Factura obtenida", detalles, nil), nil
	}
}
