package notacredito

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Service interface {
		GuardarNotaCredito(nc model.NotaCredito) (interface{}, error)
		GuardarNotaCreditoServicio(serv model.NotaCreditoServicio) (interface{}, error)
		GetUltimaNotaCreditoPorEvento(idNroNc int64) (*model.NotaCredito, error)
		UpdateNotaCredito(nc *model.NotaCredito) (*model.NotaCredito, error)
		GetNotaCredito(req *GetNotaCreditoRequest) (*model.NotaCredito, error)
		GetNotaCreditoServicios(idNroNc int64) ([]*model.NotaCreditoServicio, error)
		GetNotaCreditoServiciosAll() ([]*model.NotaCreditoServicio, error)
		BorrarNotaCreditoServicio(id string) error
		AgruparNotaCreditoServicioPorIdNC()
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{log: log, repo: repo}
}

func (s service) GuardarNotaCredito(nc model.NotaCredito) (interface{}, error) {
	res, err := s.repo.GuardarNotaCredito(nc)
	if err != nil {
		s.log.Println("Error al guardar NotaCredito:", err)
		return nil, err
	}
	return res, nil
}

func (s service) GuardarNotaCreditoServicio(serv model.NotaCreditoServicio) (interface{}, error) {
	res, err := s.repo.GuardarNotaCreditoServicio(serv)
	if err != nil {
		s.log.Println("Error al guardar NotaCreditoServicio:", err)
		return nil, err
	}
	return res, nil
}

func (s service) GetUltimaNotaCreditoPorEvento(idNroNc int64) (*model.NotaCredito, error) {
	if idNroNc == 0 {
		return nil, fmt.Errorf("idNroNc es requerido")
	}
	return s.repo.GetUltimaNotaCreditoPorEvento(idNroNc)
}

func (s service) UpdateNotaCredito(nc *model.NotaCredito) (*model.NotaCredito, error) {
	if nc == nil {
		return nil, fmt.Errorf("NotaCredito no puede ser nula")
	}
	return s.repo.UpdateNotaCredito(nc)
}

func (s service) GetNotaCredito(req *GetNotaCreditoRequest) (*model.NotaCredito, error) {
	return s.repo.GetNotaCredito(req)
}

func (s service) GetNotaCreditoServicios(idNroNc int64) ([]*model.NotaCreditoServicio, error) {
	return s.repo.GetNotaCreditoServicios(idNroNc)
}
func (s service) GetNotaCreditoServiciosAll() ([]*model.NotaCreditoServicio, error) {
	return s.repo.GetNotaCreditoServiciosAll()
}
func (s service) BorrarNotaCreditoServicio(id string) error {
	return s.repo.BorrarNotaCreditoServicio(id)
}
func (s service) AgruparNotaCreditoServicioPorIdNC() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger), // espera a que termine
			cron.Recover(cron.DefaultLogger),             // recupera panics
		),
	)

	// saca el cuerpo del job a un método para reutilizar
	job := cron.FuncJob(func() {
		fmt.Println("agrupando detalles de nota credito...", time.Now().Format("2006-01-02T15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
		servicios, err := s.GetNotaCreditoServiciosAll()
		if err != nil {
			s.log.Println("Error al obtener todos los detalles de la nota credito:", err)
			return
		}
		for _, d := range servicios {
			if d != nil {
				fmt.Printf("Detalle de nota credito: %+v\n", *d)
				if d.IDNroNc != nil {
					nota, err := s.GetUltimaNotaCreditoPorEvento(*d.IDNroNc)
					if err != nil {
						s.log.Println("Error al obtener la última Nota de Credito por evento:", err)
					}
					notaActualizada, err := ReemplazarDetalle(nota, d)
					if err != nil {
						s.log.Println("Error al reemplazar el Servicio en la nota de credito:", err)
						continue
					}

					if d.Evento == "CREATE" && nota.Evento == "CREATE" {
						_, err = s.repo.ActualizaNotaCredito(notaActualizada.IDMongo.Hex(), notaActualizada)
						if err != nil {
							fmt.Println("error Guardar BL:", err)
							continue
						}
					} else {

						_, err = s.repo.GuardarNotaCredito(notaActualizada)
						if err != nil {
							fmt.Println("error Guardar BL:", err)
							continue
						}
					}
					s.BorrarNotaCreditoServicio(d.IDMongo.Hex())
					if err != nil {
						s.log.Println("Error al actualizar la Nota de Credito:", err)
						continue
					}

				}
			}

		}

	})
	_, err := c.AddJob("0/3 * * * * *", job) // cada 50 min al segundo 0
	if err != nil {
		fmt.Println("Error al agregar tarea:", err)
		return
	}
	c.Start()
	select {}
}
