package bl

import (
	"api_auditoria/src/model"
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/jfernalv106/response_go/response"
)

type (
	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	Endpoints struct {
		GuardarBl                    Controller
		GuardarBlFecha               Controller
		GuardarBlFlete               Controller
		GuardarBlItem                Controller
		GuardarBlItemImo             Controller
		GuardarBlItemContenedor      Controller
		GuardarBlItemContenedorImo   Controller
		GuardarBlItemContenedorSello Controller
		GuardarBlLocacion            Controller
		GuardarBlObservacion         Controller
		GuardarBlParticipante        Controller
		GuardarBlReferencia          Controller
		GuardarBlTransbordo          Controller
		GuardarBlTransporte          Controller
		GetByID                      Controller
		GetByNroBl                   Controller
		GetAll                       Controller
	}

	Response struct {
		Status int         `json:"status"`
		Data   interface{} `json:"data,omitempty"`
		Error  string      `json:"error,omitempty"`
	}

	Ok struct {
		Mensaje string `json:"mensaje"`
	}

	GetByIDRequest struct {
		ID string `json:"id"`
	}
	GetByNroBlRequest struct {
		ID         *int64  `json:"id,omitempty"`
		NroBl      *string `json:"nro_bl"`
		Manifiesto *string `json:"manifiesto,omitempty"`
	}
	GetAllRequest struct {
		Filter map[string]interface{} `json:"filter"`
	}
)

func MakeEndPoint(s Service) Endpoints {
	return Endpoints{
		GuardarBl:                    GuardarBl(s),
		GuardarBlFecha:               GuardarBlFecha(s),
		GuardarBlFlete:               GuardarBlFlete(s),
		GuardarBlItem:                GuardarBlItem(s),
		GuardarBlItemImo:             GuardarBlItemImo(s),
		GuardarBlItemContenedor:      GuardarBlItemContenedor(s),
		GuardarBlItemContenedorImo:   GuardarBlItemContenedorImo(s),
		GuardarBlItemContenedorSello: GuardarBlItemContenedorSello(s),
		GuardarBlLocacion:            GuardarBlLocacion(s),
		GuardarBlObservacion:         GuardarBlObservacion(s),
		GuardarBlParticipante:        GuardarBlParticipante(s),
		GuardarBlReferencia:          GuardarBlReferencia(s),
		GuardarBlTransbordo:          GuardarBlTransbordo(s),
		GuardarBlTransporte:          GuardarBlTransporte(s),
		GetByID:                      GetByID(s),
		GetByNroBl:                   GetByNroBl(s),
		GetAll:                       GetAll(s),
	}
}
func GuardarBl(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlTopic)

		res, err := s.GuardarBl(req)
		if err != nil {
			log.Println("Error GuardarBl:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

// ---------- Subdocumentos ----------
func GuardarBlFecha(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlFechaTopic)

		res, err := s.GuardarBlFecha(req)
		if err != nil {
			log.Println("Error al obtener el BL base para BlFecha:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}

		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlFlete(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		req := request.(model.BlFleteTopic)

		res, err := s.GuardarBlFlete(req)
		if err != nil {
			log.Println("Error GuardarBlFlete:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlItem(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlItemTopic)

		res, err := s.GuardarBlItem(req)
		if err != nil {
			log.Println("Error GuardarBlItem:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlItemImo(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlItemImoTopic)

		res, err := s.GuardarBlItemImo(req)
		if err != nil {
			log.Println("Error GuardarBlItemImo:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlItemContenedor(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlItemContenedorTopic)

		res, err := s.GuardarBlItemContenedor(req)
		if err != nil {
			log.Println("Error GuardarBlItemContenedor:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlItemContenedorImo(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlItemContenedorImoTopic)

		res, err := s.GuardarBlItemContenedorImo(req)
		if err != nil {
			log.Println("Error GuardarBlItemContenedorImo:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlItemContenedorSello(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlItemContenedorSelloTopic)

		res, err := s.GuardarBlItemContenedorSello(req)
		if err != nil {
			log.Println("Error GuardarBlItemContenedorSello:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlLocacion(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlLocacionTopic)

		res, err := s.GuardarBlLocacion(req)
		if err != nil {
			log.Println("Error GuardarBlLocacion:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlObservacion(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlObservacionTopic)

		res, err := s.GuardarBlObservacion(req)
		if err != nil {
			log.Println("Error GuardarBlObservacion:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlParticipante(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlParticipanteTopic)

		res, err := s.GuardarBlParticipante(req)
		if err != nil {
			log.Println("Error GuardarBlParticipante:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlReferencia(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlReferenciaTopic)

		res, err := s.GuardarBlReferencia(req)
		if err != nil {
			log.Println("Error GuardarBlReferencia:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlTransbordo(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlTransbordoTopic)

		res, err := s.GuardarBlTransbordo(req)
		if err != nil {
			log.Println("Error GuardarBlTransbordo:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

func GuardarBlTransporte(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(model.BlTransporteTopic)

		res, err := s.GuardarBlTransporte(req)
		if err != nil {
			log.Println("Error GuardarBlTransporte:", err)
			return Response{Status: http.StatusInternalServerError, Error: err.Error()}, nil
		}
		return response.OK("BL registrada", res, nil), nil
	}
}

// ---------- Consultas ----------
func GetByID(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByIDRequest)

		bl, err := s.GetByID(req.ID)
		if err != nil {
			log.Println("Error al obtener BLs:", err)
			return response.BadRequest("Error al obtener BLs"), nil
		}
		return response.OK("BLs", bl, nil), nil
	}
}

func GetByNroBl(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetByNroBlRequest)
		fmt.Println("Request GetByNroBl:", req)
		bls, err := s.GetByNroBl(req.ID, req.NroBl, req.Manifiesto)

		if err != nil {
			log.Println("Error al obtener BLs:", err)
			fmt.Println("error BL:", err)
			return response.BadRequest("Error al obtener BLs"), nil
		}
		return response.OK("BLs", bls, nil), nil
	}
}

func GetAll(s Service) Controller {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAllRequest)

		bls, err := s.GetAll(req.Filter)
		if err != nil {
			log.Println("Error al obtener BLs:", err)
			return response.BadRequest("Error al obtener BLs"), nil
		}
		return response.OK("BLs", bls, nil), nil
	}
}
