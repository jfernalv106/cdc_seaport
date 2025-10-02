package papeleta

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
		GuardaPapeleta                 Controller
		GuardaPapeletaRecepcionDetalle Controller
		GetPapeletaRecepcion           Controller
		GetPapeletaRecepcionDetalle    Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}
	GetPapeletaRequest struct {
		NroPapeleta *string `json:"nro_papeleta,omitempty"`
		Bl          *string `json:"bl,omitempty"`
		Manifiesto  *string `json:"manifiesto,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardaPapeleta:                 GuardaPapeleta(s),
		GuardaPapeletaRecepcionDetalle: GuardaPapeletaRecepcionDetalle(s),
		GetPapeletaRecepcion:           GetPapeletaRecepcion(s),
		GetPapeletaRecepcionDetalle:    GetPapeletaRecepcionDetalle(s),
	}
}
func GuardaPapeleta(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.PapeletaRecepcionTopic)

		var id string
		if req.After != nil && req.After.NroPapeleta != nil {
			id = *req.After.NroPapeleta
		} else {
			id = *req.Before.NroPapeleta
		}
		fmt.Println("Guardando Papeleta de Recepcion:", req.After.NroPapeleta)

		if req.Op == "r" || req.Op == "c" {
			_, err := s.GuardarPapeleta(*ConvertToPapeletaRecepcion(&req, nil))
			if err != nil {
				log.Println("Error al guardar la papeleta:", req.Op, err)
				return response.BadRequest("Error al guardar la papeleta"), nil
			}

		} else {
			pr, _ := s.GetUltimaPapeletaPorEvento(id)
			_, err := s.GuardarPapeleta(*ConvertToPapeletaRecepcion(&req, pr))
			if err != nil {
				log.Println("Error al guardar la papeleta:", req.Op)
				return response.BadRequest("Error al guardar la papeleta"), nil
			}
		}
		fmt.Println("Papeleta guardada exitosamente:", req.Op)

		return response.OK("Papeleta Recepción registrada", req, nil), nil

	}
}
func GuardaPapeletaRecepcionDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.PapeletaRecepcionDetalleTopic)
		s.GuardarPapeletaDetalle(*ConvertToPapeletaRecepcionDetalle(&req))
		fmt.Println("Guardando Detalle Papeleta Recepcion :", req)

		return response.OK("Detalle papeleta Recepción detalle registrada", req, nil), nil

	}
}
func GetPapeletaRecepcion(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPapeletaRequest)

		fmt.Println("Obteniendo Papeleta Recepción:", req.NroPapeleta)
		papeletas, err := s.GetPapeletaRecepcion(req.NroPapeleta, req.Manifiesto, req.Bl)
		if err != nil {
			log.Println("Error al obtener la papeleta recepción:", err)
			return response.BadRequest("Error al obtener la papeleta recepción"), nil
		}

		fmt.Println("Papeleta Recepción obtenida exitosamente:", req.NroPapeleta)
		return response.OK("Papeleta Recepción obtenida", &papeletas, nil), nil
	}
}
func GetPapeletaRecepcionDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPapeletaRequest)

		fmt.Println("Obteniendo Papeleta Recepción:", req.NroPapeleta)
		papeletas, err := s.GetPapeletaDetalleAll()

		if err != nil {
			log.Println("Error al obtener la Papeleta Recepción Detalle:", err)
			return response.BadRequest("Error al obtener la papeleta recepción"), nil
		}

		fmt.Println("Papeleta Recepción obtenida exitosamente:", req.NroPapeleta)
		return response.OK("Detalles Papeleta Recepción obtenida", papeletas, nil), nil
	}
}
