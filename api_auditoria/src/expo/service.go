package expo

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type Service interface {
	// Crea / actualiza
	GuardarPapeleta(model.PapeletaExpo) (interface{}, error)
	GuardarPapeletaDetalle(model.PapeletaExpoDetalle) (interface{}, error)
	UpdatePapeleta(*model.PapeletaExpo) (*model.PapeletaExpo, error)

	// Consultas
	GetUltimaPapeletaPorEvento(id int64) (*model.PapeletaExpo, error)
	GetPapeleta(id *int64, booking *string, papeleta *string) ([]model.PapeletaExpo, error)
	GetPapeletaDetalle(id int64) ([]model.PapeletaExpoDetalle, error)
	GetPapeletaDetalleAll() ([]model.PapeletaExpoDetalle, error)
	BorrarPapeletaDetalle(id string) error
	AgruparPapeletaDetallePorIdPapeleta()
}

type service struct {
	log  *log.Logger
	repo Repository
}

func NewService(l *log.Logger, r Repository) Service {
	return &service{log: l, repo: r}
}

func (s *service) GuardarPapeleta(p model.PapeletaExpo) (interface{}, error) {
	res, err := s.repo.GuardarPapeletaExpo(p)
	if err != nil {
		s.log.Println("error GuardarPapeleta:", err)
		return nil, err
	}
	return res, nil
}

func (s *service) GuardarPapeletaDetalle(d model.PapeletaExpoDetalle) (interface{}, error) {
	res, err := s.repo.GuardarPapeletaExpoDetalle(d)
	if err != nil {
		s.log.Println("error GuardarPapeletaDetalle:", err)
		return nil, err
	}
	return res, nil
}

func (s *service) UpdatePapeleta(p *model.PapeletaExpo) (*model.PapeletaExpo, error) {
	if p == nil {
		return nil, fmt.Errorf("papeleta expo es nil")
	}
	return s.repo.UpdatePapeletaExpo(p)
}

func (s *service) GetUltimaPapeletaPorEvento(id int64) (*model.PapeletaExpo, error) {
	if id == 0 {
		return nil, fmt.Errorf("id es requerido")
	}
	return s.repo.GetUltimaPapeletaExpoPorEvento(id)
}

func (s *service) GetPapeleta(id *int64, booking *string, papeleta *string) ([]model.PapeletaExpo, error) {
	listPtr, err := s.repo.GetPapeletaExpo(id, booking, papeleta)
	if err != nil {
		return nil, err
	}

	out := make([]model.PapeletaExpo, 0, len(listPtr))
	for _, p := range listPtr {
		if p != nil {
			out = append(out, *p)
		}
	}
	return out, nil
}

func (s *service) GetPapeletaDetalle(id int64) ([]model.PapeletaExpoDetalle, error) {
	if id == 0 {
		return nil, fmt.Errorf("id es requerido")
	}
	listPtr, err := s.repo.GetPapeletaExpoDetalle(id)
	if err != nil {
		return nil, err
	}

	out := make([]model.PapeletaExpoDetalle, 0, len(listPtr))
	for _, d := range listPtr {
		if d != nil {
			out = append(out, *d)
		}
	}
	return out, nil
}
func (s *service) GetPapeletaDetalleAll() ([]model.PapeletaExpoDetalle, error) {
	listPtr, err := s.repo.GetPapeletaExpoDetalleAll()
	if err != nil {
		return nil, err
	}

	out := make([]model.PapeletaExpoDetalle, 0, len(listPtr))
	for _, d := range listPtr {
		if d != nil {
			out = append(out, *d)
		}
	}
	return out, nil
}
func (s *service) BorrarPapeletaDetalle(id string) error {
	if id == "" {
		return fmt.Errorf("id es requerido")
	}
	return s.repo.BorrarPapeletaExpoDetalle(id)
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
		fmt.Println("agrupando detalles de papeleta expo por id_papeleta...", time.Now().Format("2006-01-02T15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
		detalles, err := s.GetPapeletaDetalleAll()
		if err != nil {
			s.log.Println("error GetPapeletaDetalleAll:", err)
			return
		}

		for _, d := range detalles {
			if d.IDPapeleta != nil {
				px, err := s.repo.GetUltimaPapeletaExpoPorEvento(*d.IDPapeleta)
				if err != nil {
					s.log.Println("error GetUltimaPapeletaExpoPorEvento:", err)
					continue
				}
				if px == nil {
					s.log.Println("no se encontró papeleta expo con id:", *d.IDPapeleta)
					continue
				}
				pxActualizada, err := ReemplazarDetalle(px, &d)
				if err != nil {
					s.log.Println("error ReemplazarDetalle:", err)
					continue
				}

				if d.Evento == "CREATE" && px.Evento == "CREATE" {
					_, err = s.repo.ActualizaPapeletaExpo(pxActualizada.IDMongo.Hex(), pxActualizada)
					if err != nil {
						fmt.Println("error Guardar BL:", err)
						continue
					}
				} else {

					_, err = s.repo.GuardarPapeletaExpo(*pxActualizada)
					if err != nil {
						fmt.Println("error Guardar BL:", err)
						continue
					}
				}

				s.BorrarPapeletaDetalle(d.IDMongo.Hex())

				fmt.Println("Detalle agregado a Papeleta Expo ID:", *d.IDPapeleta)
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
