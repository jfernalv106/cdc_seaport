package manifiesto

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
)

type (
	Service interface {
		GuardarManifiesto(manifiesto interface{}) (interface{}, error)
		GetManifiesto(nro int64) (model.Manifiestos, error)
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

func (s service) GuardarManifiesto(manifiesto interface{}) (interface{}, error) {
	b, err := s.repo.GuardarManifiesto(manifiesto)
	if err != nil {
		s.log.Println("Error al guardar manifiesto:", err)
		return nil, err
	}
	s.log.Println("Manifiesto guardado exitosamente:", b)
	return b, nil
}
func (s service) GetManifiesto(nro int64) (model.Manifiestos, error) {
	manifiestos, err := s.repo.GetManifiesto(nro)
	if err != nil {
		s.log.Println("Error al obtener Manifiesto:", err)
		return nil, err
	}
	fmt.Printf("Manifiesto obtenidos exitosamente:%+v\n", manifiestos)
	return manifiestos, nil
}
