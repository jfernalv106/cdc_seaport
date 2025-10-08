package papeleta

import (
	"api_auditoria/src/bultos"
	"api_auditoria/src/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Service interface {
		GuardarPapeleta(papeletaRecepcion model.PapeletaRecepcion) (interface{}, error)
		GuardarPapeletaDetalle(papeletaRecepcionDetalle model.PapeletaRecepcionDetalle) interface{}
		GetUltimaPapeletaPorEvento(nroPapeleta string) (*model.PapeletaRecepcion, error)
		UpdatePapeletaRecepcion(papeletaRecepcion *model.PapeletaRecepcion) (*model.PapeletaRecepcion, error)
		GetPapeletaRecepcion(nroPapeleta *string, manifiesto *string, bl *string) (model.PapeletasRecepcion, error)
		GetPapeletaDetalleAll() ([]model.PapeletaRecepcionDetalle, error)
		BorrarPapeletaDetalle(id string) error
		AgruparPapeletaDetallePorIdPapeleta()
	}
	service struct {
		log       *log.Logger
		repo      Repository
		repoBulto bultos.Repository
	}
)

func NewService(log *log.Logger, repo Repository, repoBulto bultos.Repository) Service {
	return &service{
		log:       log,
		repo:      repo,
		repoBulto: repoBulto,
	}
}
func (s service) GuardarPapeleta(papeletaRecepcion model.PapeletaRecepcion) (interface{}, error) {
	b, err := s.repo.GuardarPapeletaRecepcion(papeletaRecepcion)
	if err != nil {
		s.log.Println("Error al guardar la Papeleta Recepción :", err)
		return nil, err
	}
	fmt.Println("Papeleta Recepción  guardada exitosamente:", b)
	return b, nil
}
func (s service) GuardarPapeletaDetalle(papeletaRecepcionDetalle model.PapeletaRecepcionDetalle) interface{} {
	b, err := s.repo.GuardarPapeletaRecepcionDetalle(papeletaRecepcionDetalle)
	if err != nil {
		s.log.Println("Error al guardar la papeleta Recepcion Detalle:", err)
		return nil
	}

	return b
}
func (s service) GetUltimaPapeletaPorEvento(nroPapeleta string) (*model.PapeletaRecepcion, error) {
	if nroPapeleta == "" {
		return nil, fmt.Errorf("el número de papeleta es requerido")
	}
	papeleta, err := s.repo.GetUltimaPapeletaPorEvento(nroPapeleta)
	if err != nil {
		s.log.Println("Error al obtener la última papeleta por evento:", err)
		return nil, err
	}
	if papeleta == nil {
		s.log.Println("No se encontró la última papeleta por evento para el número:", nroPapeleta)
		return nil, nil
	}
	fmt.Println("Última papeleta por evento obtenida exitosamente:", papeleta)
	return papeleta, nil
}
func (s service) UpdatePapeletaRecepcion(papeletaRecepcion *model.PapeletaRecepcion) (*model.PapeletaRecepcion, error) {
	if papeletaRecepcion == nil {
		return nil, fmt.Errorf("la papeleta de recepción no puede ser nula")
	}

	updatedPapeleta, err := s.repo.UpdatePapeletaRecepcion(papeletaRecepcion)
	if err != nil {
		s.log.Println("Error al actualizar la papeleta de recepción:", err)
		return nil, err
	}
	fmt.Println("Papeleta de recepción actualizada exitosamente:", updatedPapeleta)
	return updatedPapeleta, nil
}

func (s service) GetPapeletaRecepcion(nroPapeleta *string, manifiesto *string, bl *string) (model.PapeletasRecepcion, error) {

	papeletas, err := s.repo.GetPapeletaRecepcion(nroPapeleta, manifiesto, bl)
	if err != nil {
		s.log.Println("Error al obtener la visacion:", err)
		return nil, err
	}
	return papeletas, nil
}
func (s service) GetPapeletaDetalleAll() ([]model.PapeletaRecepcionDetalle, error) {

	detalles, err := s.repo.GetPapeletaDetalleAll()
	if err != nil {
		s.log.Println("Error al obtener los detalles de papeleta de recepción:", err)
		return nil, err
	}
	return detalles, nil
}
func (s service) BorrarPapeletaDetalle(id string) error {
	err := s.repo.BorrarPapeletaDetalle(id)
	if err != nil {
		s.log.Println("Error al borrar el detalle de la papeleta de recepción:", err)
		return err
	}
	return nil
}

func (s *service) AgruparPapeletaDetallePorIdPapeleta() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger), // espera a que termine
			cron.Recover(cron.DefaultLogger),             // recupera panics
		),
	)

	// saca el cuerpo del job a un método para reutilizar
	job := cron.FuncJob(func() {
		fmt.Println("agrupando detalles de papeleta recepcion por nro...", time.Now().Format("2006-01-02T15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
		detalles, err := s.GetPapeletaDetalleAll()
		if err != nil {
			s.log.Println("error GetPapeletaDetalleAll:", err)
			return
		}

		for _, d := range detalles {
			if d.NroPapeleta != nil {
				px, err := s.repo.GetUltimaPapeletaPorEvento(*d.NroPapeleta)
				if err != nil {
					s.log.Println("error GetUltimaPapeletaRecepcionPorEvento:", err)
					continue
				}
				if px == nil {
					s.log.Println("no se encontró papeleta recepcion con nro_papeleta:", *d.NroPapeleta)
					continue
				}
				pxActualizada, err := ReemplazarDetalle(px, &d)
				if err != nil {
					s.log.Println("error ReemplazarDetalle:", err)
					continue
				}

				if d.Evento == "CREATE" && px.Evento == "CREATE" {
					_, err = s.repo.ActualizaPapeletaRecepcion(pxActualizada.IDMongo.Hex(), pxActualizada)
					if err != nil {
						fmt.Println("error Guardar BL:", err)
						continue
					}
				} else {

					_, err = s.repo.GuardarPapeletaRecepcion(pxActualizada)
					if err != nil {
						fmt.Printf("error Guardar Papeleta Recepcion :", err, *pxActualizada.NroPapeleta)
						continue
					}
				}
				s.BorrarPapeletaDetalle(d.IDMongo.Hex())

				fmt.Println("Detalle agregado a Papeleta Expo ID:", *d.NroPapeleta)
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
