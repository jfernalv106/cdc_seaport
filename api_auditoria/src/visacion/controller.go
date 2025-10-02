package visacion

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
		GuardarVisacion           Controller
		GuardarVisacionMercancias Controller
		GetVisacion               Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetVisacionRequest struct {
		NroPapeleta *string `json:"nro_papeleta,omitempty"`
		Manifiesto  *string `json:"manifiesto,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardarVisacion:           GuardarVisacion(s),
		GuardarVisacionMercancias: GuardarVisacionMercancias(s),
		GetVisacion:               GetVisacion(s),
	}
}

func GuardarVisacion(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.VisacionTopic)
		var id int64
		if req.After != nil && req.After.ID != nil {
			id = *req.After.ID
		} else {
			id = *req.Before.ID
		}
		fmt.Println("Guardando Visacion :", req)
		if req.Op == "r" || req.Op == "c" {
			bs, err := s.GuardarVisacion(TransformVisacion(&req, nil))
			if err != nil {
				log.Println("Error al guardar Visacion:", err)
				return response.BadRequest("Error al guardar Visacion"), nil
			}
			if bs == nil {
				log.Println("Error al guardar la papeleta:", req.Op)
				return response.BadRequest("Error al guardar la papeleta"), nil
			}
			return response.OK("Visacion registrada", bs, nil), nil
		} else {
			vs, _ := s.GetUltimaVisacionEvento(id)
			bs, _ := s.GuardarVisacion(TransformVisacion(&req, vs))
			if bs == nil {
				log.Println("Error al guardar la papeleta:", req.Op)
				return response.BadRequest("Error al guardar la papeleta"), nil
			}
			return response.OK("Visacion registrada", bs, nil), nil

		}

	}
}
func GetVisacion(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetVisacionRequest)

		fmt.Println("Obteniendo visacion:", req.NroPapeleta)
		visaciones, err := s.GetVisacion(*req.NroPapeleta)
		if err != nil {
			log.Println("Error al obtener la visacion:", err)
			return response.BadRequest("Error al obtener la visacion"), nil
		}
		fmt.Println("visacion obtenida exitosamente:", req.NroPapeleta)
		return response.OK("visacion obtenida", visaciones, nil), nil
	}
}

func GuardarVisacionMercancias(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.MercanciasDespachadasTopic)
		s.GuardarVisacionMercancias(TransformMercanciasDespachadas(&req))
		fmt.Println("Guardando Mercancias despachadas :", req.After.ID)

		return response.OK("mercancias despachadas registrada", req, nil), nil
	}
}
