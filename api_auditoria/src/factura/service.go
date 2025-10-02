package factura

import (
	"api_auditoria/src/model"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Service interface {
		GuardarFactura(factura interface{}) (interface{}, error)
		GuardarFacturaDetalle(facturaDetalle interface{}) (interface{}, error)
		GetFactura(folio *int64, manifiesto *string) (model.Facturas, error)
		GetFacturaDetalle(id *int64) (model.FacturaDetalles, error)
		GetUltimaFacturaPorEvento(id int64) (*model.Factura, error)
		UpdateFactura(factura *model.Factura) (*model.Factura, error)
		BorrarFacturaDetalle(id string) error
		GetFacturaDetalleAll() ([]*model.FacturaDetalle, error)
		AgruparFacturaDetallePorIdFactura()
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

func (s service) GuardarFactura(factura interface{}) (interface{}, error) {
	b, err := s.repo.GuardarFactura(factura)
	if err != nil {
		s.log.Println("Error al guardar la Factura:", err)
		return nil, err
	}
	fmt.Println("Factura guardada exitosamente:", b)
	return b, nil
}
func (s service) GuardarFacturaDetalle(facturaDetalle interface{}) (interface{}, error) {
	b, err := s.repo.GuardarFacturaDetalle(facturaDetalle)
	if err != nil {
		s.log.Println("Error al guardar la factura detalle:", err)
		return nil, err
	}
	fmt.Println("Factura detalle guardada exitosamente:", b)
	return b, nil
}
func (s service) GetFactura(folio *int64, manifiesto *string) (model.Facturas, error) {
	facturas, err := s.repo.GetFactura(folio, manifiesto)
	if err != nil {
		s.log.Println("Error al obtener la Facturas:", err)
		return nil, err
	}
	fmt.Printf("Facturas obtenida exitosamente:%+v\n", facturas)
	return facturas, nil
}
func (s service) GetFacturaDetalle(id *int64) (model.FacturaDetalles, error) {
	facturaDetalle, err := s.repo.GetFacturaDetalle(id)
	if err != nil {
		s.log.Println("Error al obtener la Facturas:", err)
		return nil, err
	}
	fmt.Printf("Facturas obtenida exitosamente:%+v\n", facturaDetalle)
	return facturaDetalle, nil
}
func (s service) GetUltimaFacturaPorEvento(id int64) (*model.Factura, error) {
	factura, err := s.repo.GetUltimaFacturaPorEvento(id)
	if err != nil {
		s.log.Println("Error al obtener la última factura por evento:", err)
		return nil, err
	}
	fmt.Printf("Última factura por evento obtenida exitosamente:%+v\n", factura)
	return factura, nil
}
func (s service) UpdateFactura(factura *model.Factura) (*model.Factura, error) {
	facturaActualizada, err := s.repo.UpdateFactura(factura)
	if err != nil {
		s.log.Println("Error al actualizar la factura:", err)
		return nil, err
	}
	fmt.Printf("Factura actualizada exitosamente:%+v\n", facturaActualizada)
	return facturaActualizada, nil
}
func (s service) BorrarFacturaDetalle(id string) error {
	err := s.repo.BorrarFacturaDetalle(id)
	if err != nil {
		s.log.Println("Error al borrar el detalle de la factura:", err)
		return err
	}
	fmt.Println("Detalle de la factura borrado exitosamente:", id)
	return nil
}
func (s service) GetFacturaDetalleAll() ([]*model.FacturaDetalle, error) {
	detalles, err := s.repo.GetFacturaDetalleAll()
	if err != nil {
		s.log.Println("Error al obtener todos los detalles de la factura:", err)
		return nil, err
	}
	fmt.Printf("Detalles de la factura obtenidos exitosamente:%+v\n", detalles)
	return detalles, nil
}
func (s service) AgruparFacturaDetallePorIdFactura() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger), // espera a que termine
			cron.Recover(cron.DefaultLogger),             // recupera panics
		),
	)

	// saca el cuerpo del job a un método para reutilizar
	job := cron.FuncJob(func() {
		fmt.Println("agrupando detalles de factura por id_factura...", time.Now().Format("2006-01-02T15:04:05"))
		fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
		detalles, err := s.GetFacturaDetalleAll()
		if err != nil {
			s.log.Println("Error al obtener todos los detalles de la factura:", err)
			return
		}
		for _, d := range detalles {
			if d != nil {
				fmt.Printf("Detalle de factura: %+v\n", *d)
				if d.IDFactura != nil {
					factura, err := s.repo.GetUltimaFacturaPorEvento(*d.IDFactura)
					if err != nil {
						s.log.Println("Error al obtener la última factura por evento:", err)
					}
					facturaActualizada, err := ReemplazarDetalle(factura, d)
					if err != nil {
						s.log.Println("Error al reemplazar el detalle en la factura:", err)
						continue
					}

					if *d.Evento == "CREATE" && factura.Evento == "CREATE" {
						_, err = s.repo.ActualizaFactura(facturaActualizada.IDMongo.Hex(), facturaActualizada)
						if err != nil {
							fmt.Println("error Guardar BL:", err)
							continue
						}
					} else {

						_, err = s.repo.GuardarFactura(*facturaActualizada)
						if err != nil {
							fmt.Println("error Guardar BL:", err)
							continue
						}
					}

					s.repo.BorrarFacturaDetalle(d.IDMongo.Hex())

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
