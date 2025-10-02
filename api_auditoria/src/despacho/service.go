package despacho

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
)

type (
	Service interface {
		GuardarDespacho(visacion interface{}) (interface{}, error)
		GetDespacho(id *int64, visacion *int64, expo *int64, guia *int64) (model.Despachos, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s service) GuardarDespacho(despacho interface{}) (interface{}, error) {
	b, err := s.repo.GuardarDespacho(despacho)
	if err != nil {
		s.log.Println("Error al guardar el despacho:", err)
		return nil, err
	}
	fmt.Println("Despacho guardado exitosamente:", b)
	return b, nil
}
func (s service) GetDespacho(id *int64, visacion *int64, expo *int64, guia *int64) (model.Despachos, error) {
	despachos, err := s.repo.GetDespacho(id, visacion, expo, guia)
	if err != nil {
		fmt.Println("Error al obtener el despacho:", err)
		return nil, err
	}
	fmt.Printf("Despacho obtenido exitosamente:%+v\n", despachos)
	return despachos, nil
}
