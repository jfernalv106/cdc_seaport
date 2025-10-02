package despacho

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
		GuardarDespacho Controller
		GetDespacho     Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetDespachoRequest struct {
		ID       *int64 `json:"id,omitempty"`
		Visacion *int64 `json:"visacion,omitempty"`
		Expo     *int64 `json:"expo,omitempty"`
		Guia     *int64 `json:"guia,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardarDespacho: GuardarDespacho(s),
		GetDespacho:     GetDespacho(s),
	}
}

func GuardarDespacho(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.DespachoTopic)

		bs, err := s.GuardarDespacho(TransformDespacho(&req, nil))
		if err != nil {
			log.Println("Error al guardar Despacho:", err)
			return response.BadRequest("Error al guardar Despacho"), nil
		}
		if bs == nil {
			log.Println("Error al guardar la papeleta:", req.Op)
			return response.BadRequest("Error al guardar la papeleta"), nil
		}
		return response.OK("Despacho registrado", bs, nil), nil

	}
}
func GetDespacho(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetDespachoRequest)

		despachos, err := s.GetDespacho(req.ID, req.Visacion, req.Expo, req.Guia)
		if err != nil {
			log.Println("Error al obtener el despacho:", err)
			return response.BadRequest("Error al obtener el despacho"), nil
		}
		fmt.Println("despacho obtenido exitosamente:", req.ID)
		return response.OK("despacho obtenido", despachos, nil), nil
	}
}
