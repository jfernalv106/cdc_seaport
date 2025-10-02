package manifiesto

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
		GuardarManiesto Controller
		GetManiesto     Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetManifiestoRequest struct {
		Nro int64 `json:"nro"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardarManiesto: GuardarManiesto(s),
		GetManiesto:     GetManiesto(s),
	}
}

func GuardarManiesto(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.ManifiestoTopic)

		bs, err := s.GuardarManifiesto(ConvertManifiestoTopicToManifiesto(&req))
		if err != nil {
			log.Println("Error al guardar Manifiesto:", err)
			return response.BadRequest("Error al guardar Manifiesto"), nil
		}
		if bs == nil {

			return response.BadRequest("Error al guardar Manifiesto"), nil
		}

		return response.OK("Manifiesto registrado", bs, nil), nil

	}
}
func GetManiesto(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetManifiestoRequest)

		bultos, err := s.GetManifiesto(req.Nro)
		if err != nil {
			log.Println("Error al obtener los Bultos:", err)
			return response.BadRequest("Error al obtener los Bultos"), nil
		}
		fmt.Println("Manifiesto obtenida exitosamente:", req.Nro)
		return response.OK("Manifiesto obtenidos", bultos, nil), nil
	}
}
