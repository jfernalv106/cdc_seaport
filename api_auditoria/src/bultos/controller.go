package bultos

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"context"
	"fmt"
	"log"

	"github.com/jfernalv106/response_go/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		GuardarBulto Controller
		GetBultos    Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetBultoRequest struct {
		Cod string `json:"cod"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardarBulto: GuardarBulto(s),
		GetBultos:    GetBultos(s),
	}
}

func GuardarBulto(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.BultoTopic)
		var topic model.Bulto
		if *req.Op == "d" {
			topic = *req.Before
		} else {
			topic = *req.After
		}
		topic.Evento = utils.MapOperation(func(s *string) string {
			if s != nil {
				return *s
			}
			return ""
		}(req.Op))

		topic.FechaEvento = utils.ToFormattedDateTime(req.TsMS)

		fmt.Println("Guardando Bulto:", topic.Nombre)
		bs, err := s.GuardarBulto(topic)
		if err != nil {
			log.Println("Error al guardar Bulto:", err)
			return response.BadRequest("Error al guardar Bulto"), nil
		}
		if bs == nil {
			log.Println("Error al guardar detalle Bulto:", topic.Nombre)
			return response.BadRequest("Error al guardar Bulto"), nil
		}
		fmt.Println("Bulto guardado exitosamente:", topic.Nombre)

		return response.OK("Bulto registrado", bs, nil), nil

	}
}
func GetBultos(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetBultoRequest)

		bultos, err := s.GetBultos(req.Cod)
		if err != nil {
			log.Println("Error al obtener los Bultos:", err)
			return response.BadRequest("Error al obtener los Bultos"), nil
		}
		fmt.Println("Bultos obtenida exitosamente:", req.Cod)
		return response.OK("Bultos obtenidos", bultos, nil), nil
	}
}
