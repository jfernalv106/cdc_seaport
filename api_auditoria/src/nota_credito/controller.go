package notacredito

import (
	"api_auditoria/src/model"
	"context"
	"fmt"

	"github.com/jfernalv106/response_go/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		GuardaNotaCredito         Controller
		GuardaNotaCreditoServicio Controller
		GetNotaCredito            Controller
		GetNotaCreditoServicios   Controller
	}

	GetNotaCreditoRequest struct {
		IDNroNc   *int64 `json:"id_nro_nc,omitempty"`
		IDFolNc   *int64 `json:"id_fol_nc,omitempty"`
		IDFactura *int64 `json:"id_factura,omitempty"`
		IDFolio   *int64 `json:"id_folio,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardaNotaCredito:         GuardaNotaCredito(s),
		GuardaNotaCreditoServicio: GuardaNotaCreditoServicio(s),
		GetNotaCredito:            GetNotaCredito(s),
		GetNotaCreditoServicios:   GetNotaCreditoServicios(s),
	}
}

func GuardaNotaCredito(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.NotaCreditoTopic)
		var id int64
		if req.After != nil {
			id = *req.After.IDFolNc
		} else {
			id = *req.Before.IDFolNc
		}
		if *req.Op == "r" || *req.Op == "c" {
			_, err := s.GuardarNotaCredito(*ConvertToNotaCredito(&req, nil))
			if err != nil {
				fmt.Println("Error al guardar NotaCredito:", err)
				return response.BadRequest("Error al guardar NotaCredito"), nil
			}
		} else {
			nc, _ := s.GetUltimaNotaCreditoPorEvento(id)
			_, err := s.GuardarNotaCredito(*ConvertToNotaCredito(&req, nc))
			if err != nil {
				fmt.Println("Error al guardar NotaCredito:", err)
				return response.BadRequest("Error al guardar NotaCredito"), nil
			}
		}
		return response.OK("NotaCredito registrada", req, nil), nil
	}
}

func GuardaNotaCreditoServicio(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.NotaCreditoServTopic)

		_, err := s.GuardarNotaCreditoServicio(*ConvertToNotaCreditoServicio(&req))
		if err != nil {
			return response.BadRequest("Error al guardar NotaCreditoServicio"), nil
		}

		return response.OK("NotaCreditoServicio registrado", req, nil), nil
	}
}

func GetNotaCredito(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNotaCreditoRequest)

		if req.IDNroNc == nil && req.IDFolNc == nil && req.IDFactura == nil && req.IDFolio == nil {
			return response.BadRequest("Debe enviar al menos un criterio de búsqueda"), nil
		}

		nc, err := s.GetNotaCredito(&req)
		if err != nil {
			return response.BadRequest("Error al obtener NotaCredito"), nil
		}
		return response.OK("NotaCredito obtenida", nc, nil), nil
	}
}

func GetNotaCreditoServicios(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetNotaCreditoRequest)
		if req.IDNroNc == nil {
			return response.BadRequest("id_nro_nc es requerido"), nil
		}

		serv, err := s.GetNotaCreditoServicios(*req.IDNroNc)
		if err != nil {
			return response.BadRequest("Error al obtener servicios"), nil
		}
		return response.OK("Servicios NotaCredito obtenidos", serv, nil), nil
	}
}
