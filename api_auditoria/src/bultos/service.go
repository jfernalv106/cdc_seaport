package bultos

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
)

type (
	Service interface {
		GuardarBulto(bulto interface{}) (interface{}, error)
		GetBultos(cod string) (model.Bultos, error)
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

func (s service) GuardarBulto(bulto interface{}) (interface{}, error) {
	b, err := s.repo.GuardarBulto(bulto)
	if err != nil {
		s.log.Println("Error al guardar la visaciones:", err)
		return nil, err
	}
	fmt.Println("Bulto guardado exitosamente:", b)
	return b, nil
}
func (s service) GetBultos(cod string) (model.Bultos, error) {
	visaciones, err := s.repo.GetBultos(cod)
	if err != nil {
		s.log.Println("Error al obtener Bulto:", err)
		return nil, err
	}
	fmt.Printf("Bultos obtenidos exitosamente:%+v\n", visaciones)
	return visaciones, nil
}
