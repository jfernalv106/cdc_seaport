package visacion

import (
	"api_auditoria/src/model"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Service interface {
		GuardarVisacion(visacion interface{}) (interface{}, error)
		GetVisacion(nroPapeleta string) (model.Visaciones, error)
		GetUltimaVisacionEvento(id int64) (*model.Visacion, error)
		GuardarVisacionMercancias(mercancias interface{}) (interface{}, error)
		GetVisacionMercancias(id int64) (model.MercanciasDespachadas, error)
		UpdateVisaje(visacion *model.Visacion) (*model.Visacion, error)
		GetMercanciasAll() ([]model.MercanciasDespachada, error)
		BorrarMercancia(id string) error
		AgruparMercancias()
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

func (s service) GuardarVisacion(visacion interface{}) (interface{}, error) {
	b, err := s.repo.GuardarVisacion(visacion)
	if err != nil {
		s.log.Println("Error al guardar la visaciones:", err)
		return nil, err
	}
	s.log.Println("visaciones guardada exitosamente:", b)
	return b, nil
}
func (s service) GetVisacion(nroPapeleta string) (model.Visaciones, error) {
	visaciones, err := s.repo.GetVisacion(nroPapeleta)
	if err != nil {
		s.log.Println("Error al obtener la visacion:", err)
		return nil, err
	}
	fmt.Printf("Visacion obtenida exitosamente:%+v\n", visaciones)
	return visaciones, nil
}
func (s service) GuardarVisacionMercancias(mercancias interface{}) (interface{}, error) {
	b, err := s.repo.GuardarVisacionMercancias(mercancias)
	if err != nil {
		s.log.Println("Error al guardar las mercancias de la visacion:", err)
		return nil, err
	}
	s.log.Println("Mercancias de la visacion guardadas exitosamente:", b)
	return b, nil
}
func (s service) GetUltimaVisacionEvento(id int64) (*model.Visacion, error) {

	visacion, err := s.repo.GetUltimaVisacionEvento(id)

	if err != nil {
		s.log.Println("Error al obtener la última visacion por evento:", err)
		return nil, err
	}
	if visacion == nil {
		s.log.Println("No existe una visacion para el evento", err)
		return nil, errors.New("no existe una visacion para el evento")
	}
	fmt.Println("Última visacion por evento obtenida exitosamente:", visacion)
	return visacion, nil
}

func (s service) GetVisacionMercancias(id int64) (model.MercanciasDespachadas, error) {
	mercancias, err := s.repo.GetVisacionMercancias(id)
	if err != nil {
		s.log.Println("Error al obtener las mercancias de la visacion:", err)
		return nil, err
	}
	fmt.Printf("Mercancias de la visacion obtenidas exitosamente:%+v\n", mercancias)
	return mercancias, nil
}
func (s service) UpdateVisaje(visacion *model.Visacion) (*model.Visacion, error) {
	updatedVisacion, err := s.repo.UpdateVisaje(visacion)
	if err != nil {
		s.log.Println("Error al actualizar la visacion:", err)
		return nil, err
	}
	fmt.Println("Visacion actualizada exitosamente:", updatedVisacion)
	return updatedVisacion, nil
}
func (s service) GetMercanciasAll() ([]model.MercanciasDespachada, error) {

	detalles, err := s.repo.GetMercanciasAll()
	if err != nil {
		s.log.Println("Error al obtener los detalles de mercancias de la visacion:", err)
		return nil, err
	}
	return detalles, nil
}
func (s service) BorrarMercancia(id string) error {

	err := s.repo.BorrarMercancia(id)
	if err != nil {
		s.log.Println("Error al borrar el detalle de mercancias de la visacion:", err)
	}
	return err
}

func (s *service) AgruparMercancias() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger), // espera a que termine
			cron.Recover(cron.DefaultLogger),             // recupera panics
		),
	)

	// saca el cuerpo del job a un método para reutilizar
	job := cron.FuncJob(func() {
		fmt.Println("agrupando Mercancias visacion...", time.Now().Format("2006-01-02T15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
		detalles, err := s.GetMercanciasAll()
		if err != nil {
			s.log.Println("error GetPapeletaDetalleAll:", err)
			return
		}

		for _, d := range detalles {
			if d.IDVisaje != nil {
				px, err := s.repo.GetUltimaVisacionEvento(*d.IDVisaje)
				if err != nil {
					s.log.Println("error GetUltimaVisacionEvento:", err)
					continue
				}
				if px == nil {
					s.log.Println("no se encontró papeleta recepcion con id:", *d.IDVisaje)
					continue
				}
				pxActualizada, err := ReemplazarMercacia(px, &d)
				if err != nil {
					s.log.Println("error ReemplazarDetalle:", err)
					continue
				}

				if d.Evento == "CREATE" && px.Evento == "CREATE" {
					_, err = s.repo.ActualizaVisacion(pxActualizada.IDMongo.Hex(), pxActualizada)
					if err != nil {
						fmt.Println("error Guardar BL:", err)
						continue
					}
				} else {

					_, err = s.repo.GuardarVisacion(pxActualizada)
					if err != nil {
						fmt.Println("error Guardar BL:", err)
						continue
					}
				}

				s.BorrarMercancia(d.IDMongo.Hex())

				fmt.Println("Detalle agregado a Papeleta Expo ID:", *d.IDVisaje)
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
