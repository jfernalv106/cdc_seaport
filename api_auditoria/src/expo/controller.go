package expo

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
		GuardaPapeleta        Controller
		GuardaPapeletaDetalle Controller
		GetPapeleta           Controller
		GetPapeletaDetalle    Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}
	Ok struct {
		Mensaje string `json:"mensaje"`
	}

	// Para GET /papeleta_expo y /papeleta_expo_detalle
	GetExpoRequest struct {
		ID       *int64  `json:"id,omitempty"`
		Booking  *string `json:"booking,omitempty"`
		Papeleta *string `json:"nro_papeleta,omitempty"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardaPapeleta:        GuardaPapeleta(s),
		GuardaPapeletaDetalle: GuardaPapeletaDetalle(s),
		GetPapeleta:           GetPapeleta(s),
		GetPapeletaDetalle:    GetPapeletaDetalle(s),
	}
}

// ------------------- CREATE / UPSERT -------------------

func GuardaPapeleta(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.PapeletaExpoTopic)

		var id int64
		if req.After != nil && req.After.ID != nil {
			id = *req.After.ID
		} else {
			id = *req.Before.ID
		}

		// "r" (snapshot/read) o "c" (create) => insert inicial desde el tópico
		if req.Op == "r" || req.Op == "c" {
			_, err := s.GuardarPapeleta(*ConvertToPapeletaExpo(&req, nil))
			if err != nil {
				log.Println("Error al guardar la papeleta expo:", err)
				return response.BadRequest("Error al guardar la papeleta expo"), nil
			}

		} else {

			pr, _ := s.GetUltimaPapeletaPorEvento(id)
			_, err := s.GuardarPapeleta(*ConvertToPapeletaExpo(&req, pr))
			if err != nil {
				log.Println("Error al guardar la papeleta expo:", err)
				return response.BadRequest("Error al guardar la papeleta expo"), nil
			}

		}

		fmt.Println("Papeleta Expo guardada exitosamente:", req.Op)
		return response.OK("Papeleta Expo registrada", req, nil), nil
	}
}

func GuardaPapeletaDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.PapeletaExpoDetalleTopic)

		_, err := s.GuardarPapeletaDetalle(*ConvertToPapeletaExpoDetalle(&req))
		if err != nil {
			log.Println("Error al guardar el detalle de la papeleta expo:", err)
			return response.BadRequest("Error al guardar el detalle de la papeleta expo"), nil
		}

		return response.OK("Detalle Papeleta Expo registrado", req, nil), nil
	}
}

func GetPapeleta(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetExpoRequest)

		fmt.Println("Obteniendo Papeleta Expo: id =", req.ID, " booking =", req.Booking)
		papeletas, err := s.GetPapeleta(req.ID, req.Booking, req.Papeleta)
		if err != nil {
			log.Println("Error al obtener la Papeleta Expo:", err)
			return response.BadRequest("Error al obtener la Papeleta Expo"), nil
		}

		log.Println("Papeleta Expo obtenida exitosamente:", req.ID, req.Booking)
		return response.OK("Papeleta Expo obtenida", &papeletas, nil), nil
	}
}

func GetPapeletaDetalle(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetExpoRequest)

		if req.ID == nil {
			log.Println("id es requerido para obtener detalles de Papeleta Expo")
			return response.BadRequest("id es requerido"), nil
		}

		log.Println("Obteniendo Papeleta Expo Detalle: id =", *req.ID)
		detalles, err := s.GetPapeletaDetalle(*req.ID)
		if err != nil {
			log.Println("Error al obtener la Papeleta Expo Detalle:", err)
			return response.BadRequest("Error al obtener la Papeleta Expo Detalle"), nil
		}

		log.Println("Papeleta Expo Detalle obtenida exitosamente:", *req.ID)
		return response.OK("Detalles Papeleta Expo obtenida", detalles, nil), nil
	}
}
