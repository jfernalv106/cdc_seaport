package bl

import (
	"api_auditoria/src/model"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

type (
	Service interface {
		GuardarBl(topic model.BlTopic) (interface{}, error)
		GuardarBlFecha(topic model.BlFechaTopic) (interface{}, error)
		GuardarBlFlete(topic model.BlFleteTopic) (interface{}, error)
		GuardarBlItem(topic model.BlItemTopic) (interface{}, error)
		GuardarBlItemImo(topic model.BlItemImoTopic) (interface{}, error)
		GuardarBlItemContenedor(topic model.BlItemContenedorTopic) (interface{}, error)
		GuardarBlItemContenedorImo(topic model.BlItemContenedorImoTopic) (interface{}, error)
		GuardarBlItemContenedorSello(topic model.BlItemContenedorSelloTopic) (interface{}, error)
		GuardarBlLocacion(topic model.BlLocacionTopic) (interface{}, error)
		GuardarBlObservacion(topic model.BlObservacionTopic) (interface{}, error)
		GuardarBlParticipante(topic model.BlParticipanteTopic) (interface{}, error)
		GuardarBlReferencia(topic model.BlReferenciaTopic) (interface{}, error)
		GuardarBlTransbordo(topic model.BlTransbordoTopic) (interface{}, error)
		GuardarBlTransporte(topic model.BlTransporteTopic) (interface{}, error)

		GetByID(id string) (*model.BL, error)
		GetByNroBl(id *int64, nroBl *string, manifiesto *string) ([]*model.BL, error)
		GetAll(filter map[string]interface{}) ([]*model.BL, error)
		GetUltimoBl(id *int64) (*model.BL, error)
		GetBlFechaAll() ([]*model.BlFecha, error)
		GetBlFleteAll() ([]*model.BlFlete, error)
		GetBlItemAll() ([]*model.BlItem, error)
		GetBlItemImoAll() ([]*model.BlItemImo, error)
		GetBlItemContenedorAll() ([]*model.BlItemContenedor, error)
		GetBlItemContenedorImoAll() ([]*model.BlItemContenedorImo, error)
		GetBlItemContenedorSelloAll() ([]*model.BlItemContenedorSello, error)
		GetBlLocacionAll() ([]*model.BlLocacion, error)
		GetBlObservacionAll() ([]*model.BlObservacion, error)
		GetBlParticipanteAll() ([]*model.BlParticipante, error)
		GetBlReferenciaAll() ([]*model.BlReferencia, error)
		GetBlTransbordoAll() ([]*model.BlTransbordo, error)
		GetBlTransporteAll() ([]*model.BlTransporte, error)
		BorraBlFecha(id string) error
		BorraBlFlete(id string) error
		BorraBlItem(id string) error
		BorraBlReferencia(id string) error
		BorraBlTransbordo(id string) error
		BorraBlTransporte(id string) error
		BorraBlLocacion(id string) error
		BorraBlObservacion(id string) error
		BorraBlParticipante(id string) error
		BorraBlItemImo(id string) error
		BorraBlItemContenedor(id string) error
		BorraBlItemContenedorImo(id string) error
		BorraBlItemContenedorSello(id string) error
		AgruparBlFecha()
		AgruparBlFlete()
		AgruparBlItem()
		AgruparBlReferencia()
		AgruparBlTransbordo()
		AgruparBlTransporte()
		AgruparBlLocacion()
		AgruparBlObservacion()
		AgruparBlParticipante()
		AgruparBlItemImo()
		AgruparBlItemContenedor()
		AgruparBlItemContenedorImo()
		AgruparBlItemContenedorSello()
		AgruparBl()
	}

	service struct {
		logger *log.Logger
		repo   Repository
	}
)

func NewService(logger *log.Logger, repo Repository) Service {
	return &service{
		logger: logger,
		repo:   repo,
	}
}

func (s *service) GuardarBl(topic model.BlTopic) (interface{}, error) {
	var id int64
	if topic.After != nil && topic.After.ID != nil {
		id = *topic.After.ID
	} else {
		id = *topic.Before.ID
	}
	base, _ := s.repo.GetUltimoPorId(&id)
	bl := ConvertToBL(&topic, base)
	return s.repo.Guardar(bl)
}

func (s *service) GuardarBlFecha(topic model.BlFechaTopic) (interface{}, error) {

	bl, err := s.repo.GuardarBlFecha(ConvertToBLFecha(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl fecha:", err)
		return nil, err
	}

	return bl, nil
}

func (s *service) GuardarBlFlete(topic model.BlFleteTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlFlete(ConvertToBLFlete(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl flete:", err)
		return nil, err
	}

	return bl, nil
}

func (s *service) GuardarBlItem(topic model.BlItemTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlItem(ConvertToBLItem(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl item:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlItemImo(topic model.BlItemImoTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlItemImo(ConvertToBLItemImo(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl item imo:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlItemContenedor(topic model.BlItemContenedorTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlItemContenedor(ConvertToBLItemContenedor(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl item contenedor:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlItemContenedorImo(topic model.BlItemContenedorImoTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlItemContenedorImo(ConvertToBLItemContenedorImo(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl item contenedor imo:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlItemContenedorSello(topic model.BlItemContenedorSelloTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlItemContenedorSello(ConvertToBLItemContenedorSello(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl item contenedor sello:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlLocacion(topic model.BlLocacionTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlLocacion(ConvertToBLLocacion(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl locacion:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlObservacion(topic model.BlObservacionTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlObservacion(ConvertToBLObservacion(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl observacion:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlParticipante(topic model.BlParticipanteTopic) (interface{}, error) {

	bl, err := s.repo.GuardarBlParticipante(ConvertToBLParticipante(&topic))
	if err != nil {
		s.logger.Println("error al reemplazar participante:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlReferencia(topic model.BlReferenciaTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlReferencia(ConvertToBLReferencia(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl referencia:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlTransbordo(topic model.BlTransbordoTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlTransbordo(ConvertToBLTransbordo(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl transbordo:", err)
		return nil, err
	}
	return bl, nil
}

func (s *service) GuardarBlTransporte(topic model.BlTransporteTopic) (interface{}, error) {
	bl, err := s.repo.GuardarBlTransporte(ConvertToBLTransporte(&topic))
	if err != nil {
		s.logger.Println("error al guardar bl transporte:", err)
		return nil, err
	}

	return bl, nil
}

// 📌 Métodos de lectura
func (s *service) GetByID(id string) (*model.BL, error) {
	return s.repo.GetByIDString(id)
}

func (s *service) GetByNroBl(id *int64, nroBl *string, manifiesto *string) ([]*model.BL, error) {
	var list []*model.BL
	bls, err := s.repo.GetByNroBl(id, nroBl, manifiesto)
	if err != nil {
		return nil, err
	}
	for _, b := range bls {

		list = append(list, InicilizaVacio(b))

	}
	return list, nil
}

func (s *service) GetAll(filter map[string]interface{}) ([]*model.BL, error) {
	return s.repo.GetAll(filter)
}
func (s *service) GetUltimoBl(id *int64) (*model.BL, error) {
	base, _ := s.repo.GetUltimoPorId(id)
	jsonBytes, err := json.MarshalIndent(base, "", "  ")
	if err != nil {
		panic(err)
	}
	jsonStr := string(jsonBytes)

	fmt.Println("BL como string JSON:")
	fmt.Println(jsonStr)
	return base, nil
}
func (s *service) GetBlFechaAll() ([]*model.BlFecha, error) {
	return s.repo.GetBlFechaAll()
}
func (s *service) GetBlFleteAll() ([]*model.BlFlete, error) {
	return s.repo.GetBlFleteAll()
}
func (s *service) GetBlItemAll() ([]*model.BlItem, error) {
	return s.repo.GetBlItemAll()
}
func (s *service) GetBlItemImoAll() ([]*model.BlItemImo, error) {
	return s.repo.GetBlItemImoAll()
}
func (s *service) GetBlItemContenedorAll() ([]*model.BlItemContenedor, error) {
	return s.repo.GetBlItemContenedorAll()
}
func (s *service) GetBlItemContenedorImoAll() ([]*model.BlItemContenedorImo, error) {
	return s.repo.GetBlItemContenedorImoAll()
}
func (s *service) GetBlItemContenedorSelloAll() ([]*model.BlItemContenedorSello, error) {
	return s.repo.GetBlItemContenedorSelloAll()
}
func (s *service) GetBlLocacionAll() ([]*model.BlLocacion, error) {
	return s.repo.GetBlLocacionAll()
}
func (s *service) GetBlObservacionAll() ([]*model.BlObservacion, error) {
	return s.repo.GetBlObservacionAll()
}
func (s *service) GetBlParticipanteAll() ([]*model.BlParticipante, error) {
	return s.repo.GetBlParticipanteAll()
}
func (s *service) GetBlReferenciaAll() ([]*model.BlReferencia, error) {
	return s.repo.GetBlReferenciaAll()
}
func (s *service) GetBlTransbordoAll() ([]*model.BlTransbordo, error) {
	return s.repo.GetBlTransbordoAll()
}
func (s *service) GetBlTransporteAll() ([]*model.BlTransporte, error) {
	return s.repo.GetBlTransporteAll()
}

// 📌 Métodos de borrado
func (s *service) BorraBlFecha(id string) error {
	return s.repo.BorraBlFecha(id)
}
func (s *service) BorraBlFlete(id string) error {
	return s.repo.BorraBlFlete(id)
}
func (s *service) BorraBlItem(id string) error {
	return s.repo.BorraBlItem(id)
}
func (s *service) BorraBlLocacion(id string) error {
	return s.repo.BorraBlLocacion(id)
}
func (s *service) BorraBlObservacion(id string) error {
	return s.repo.BorraBlObservacion(id)
}
func (s *service) BorraBlParticipante(id string) error {
	return s.repo.BorraBlParticipante(id)
}
func (s *service) BorraBlItemImo(id string) error {
	return s.repo.BorraBlItemImo(id)
}
func (s *service) BorraBlItemContenedor(id string) error {
	return s.repo.BorraBlItemContenedor(id)
}
func (s *service) BorraBlItemContenedorImo(id string) error {
	return s.repo.BorraBlItemContenedorImo(id)
}
func (s *service) BorraBlItemContenedorSello(id string) error {
	return s.repo.BorraBlItemContenedorSello(id)
}
func (s *service) BorraBlReferencia(id string) error {
	return s.repo.BorraBlReferencia(id)
}
func (s *service) BorraBlTransbordo(id string) error {
	return s.repo.BorraBlTransbordo(id)
}
func (s *service) BorraBlTransporte(id string) error {
	return s.repo.BorraBlTransporte(id)
}

// 📌 Métodos de agrupació
func (s *service) AgruparBlFecha() {

	fmt.Println("agrupando Bl Fechas...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	fechas, err := s.GetBlFechaAll()
	if err != nil {
		fmt.Println("error GetBlFechaAll:", err)
		return
	}

	for _, d := range fechas {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(*&d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlFecha(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlFecha:", err)
				continue
			}
			_, err = s.repo.Guardar(*pxActualizada)
			s.BorraBlFecha(d.IDMongo.Hex())
			if err != nil {
				fmt.Println("error Guardar BL:", err)
				continue
			}
			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlFlete() {

	fmt.Println("agrupando Bl flete...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	fletes, err := s.GetBlFleteAll()
	if err != nil {
		fmt.Println("error GetBlFleteAll:", err)
		return
	}

	for _, d := range fletes {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlFlete(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlFecha:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlFlete(d.IDMongo.Hex())
			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlItem() {

	fmt.Println("agrupando Bl item...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	items, err := s.GetBlItemAll()
	if err != nil {
		fmt.Println("error GetBlItemAll:", err)
		return
	}

	for _, d := range items {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlItem(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlItem:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlItem(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlReferencia() {

	fmt.Println("agrupando Bl item...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	referencias, err := s.GetBlReferenciaAll()
	if err != nil {
		fmt.Println("error GetBlReferenciaAll:", err)
		return
	}

	for _, d := range referencias {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(*&d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlReferencia(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlReferencia:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlReferencia(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlTransbordo() {

	fmt.Println("agrupando Bl transbordo...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	transbordos, err := s.GetBlTransbordoAll()
	if err != nil {
		fmt.Println("error GetBlTransbordoAll:", err)
		return
	}

	for _, d := range transbordos {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlTransbordo(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlReferencia:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlTransbordo(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlTransporte() {

	fmt.Println("agrupando Bl transbordo...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	transportes, err := s.GetBlTransporteAll()
	if err != nil {
		fmt.Println("error GetBlTransporteAll:", err)
		return
	}

	for _, d := range transportes {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlTransporte(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlTransporte:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlTransporte(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlLocacion() {

	fmt.Println("agrupando Bl locaciones...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	locaciones, err := s.GetBlLocacionAll()
	if err != nil {
		fmt.Println("error GetBlTransporteAll:", err)
		return
	}

	for _, d := range locaciones {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlLocacion(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlTransporte:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlLocacion(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlObservacion() {

	fmt.Println("agrupando Bl locaciones...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	observaciones, err := s.GetBlObservacionAll()
	if err != nil {
		fmt.Println("error GetBlTransporteAll:", err)
		return
	}

	for _, d := range observaciones {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlObservacion(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlObservacion:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlObservacion(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}
func (s *service) AgruparBlParticipante() {

	fmt.Println("agrupando Bl participantes...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	participantes, err := s.GetBlParticipanteAll()
	if err != nil {
		fmt.Println("error GetBlTransporteAll:", err)
		return
	}

	for _, d := range participantes {
		if d.BlNroBl != nil {
			px, err := s.repo.GetUltimoPorId(*&d.BlNroBl)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con id:", *d.BlNroBl)
				continue
			}
			pxActualizada, err := ReemplazarBlParticipante(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlObservacion:", err)
				continue
			}
			if d.Evento == "CREATE" && px.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(pxActualizada.IDMongo.Hex(), pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*pxActualizada)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlParticipante(d.IDMongo.Hex())

			fmt.Println("BL fecha agregado a BL ID:", *d.BlNroBl)
		}
	}

}

// agrupacion nietos  GetUltimoPorIdItem
func (s *service) AgruparBlItemImo() {

	fmt.Println("agrupando Item IMO...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	imos, err := s.GetBlItemImoAll()
	if err != nil {
		fmt.Println("error GetBlTransporteAll:", err)
		return
	}

	for _, d := range imos {
		if d.BlItemID != nil {
			px, err := s.repo.GetUltimoPorIdItem(d.BlItemID)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL con item id:", *d.BlItemID)
				continue
			}
			itemActualizado, err := ReemplazarBlItemImo(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlObservacion:", err)
				continue
			}
			b, err := s.repo.GetUltimoPorId(itemActualizado.BlNroBl)
			if err != nil {
				fmt.Println("GetUltimoPorId AgruparBlItemImo:", *d.BlItemID)
			}
			bl, err := ReemplazarBlItem(b, *itemActualizado)
			if err != nil {
				fmt.Println("GetUltimoPorId ReemplazarBlItem   :", *d.BlItemID)
			}
			if d.Evento == "CREATE" && bl.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(bl.IDMongo.Hex(), bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlItemImo(d.IDMongo.Hex())

			fmt.Println("BL Item IMO agregado a BL ID:", *d.BlItemID)
		}
	}

}
func (s *service) AgruparBlItemContenedor() {

	fmt.Println("agrupando Item Contenedor...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	imos, err := s.GetBlItemContenedorAll()
	if err != nil {
		fmt.Println("error GetBlItemContenedorAll:", err)
		return
	}

	for _, d := range imos {
		if d.BlItemID != nil {
			px, err := s.repo.GetUltimoPorIdItem(d.BlItemID)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if px == nil {
				fmt.Println("no se encontró BL Item con item id:", *d.BlItemID)
				continue
			}
			itemActualizado, err := ReemplazarBlItemContenedor(px, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlItemContenedor:", err)
				continue
			}
			b, err := s.repo.GetUltimoPorId(itemActualizado.BlNroBl)
			if err != nil {
				fmt.Println("GetUltimoPorId AgruparBlItemImo:", *d.BlItemID)
			}
			bl, err := ReemplazarBlItem(b, *itemActualizado)
			if err != nil {
				fmt.Println("GetUltimoPorId ReemplazarBlItem   :", *d.BlItemID)
			}
			if d.Evento == "CREATE" && bl.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(bl.IDMongo.Hex(), bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlItemContenedor(d.IDMongo.Hex())

			fmt.Println("BL Item Contenedor agregado a BL ID:", *d.BlItemID)
		}
	}

}
func (s *service) AgruparBlItemContenedorImo() {

	fmt.Println("agrupando Item Contenedor Imo...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	imos, err := s.GetBlItemContenedorImoAll()
	if err != nil {
		fmt.Println("error GetBlItemContenedorAll:", err)
		return
	}

	for _, d := range imos {
		if d.BlItemContenedorID != nil {
			cnt, err := s.repo.GetUltimoPorIdItemContenedor(d.BlItemContenedorID)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}
			if cnt == nil {
				fmt.Println("no se encontró BL Item con item id:", *d.BlItemContenedorID)
				continue
			}
			cntActualizado, err := ReemplazarBlItemContenedorImo(cnt, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlItemContenedor:", err)
				continue
			}
			item, err := s.repo.GetUltimoPorIdItem(cntActualizado.BlItemID)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}

			itemActualizado, err := ReemplazarBlItemContenedor(item, *cntActualizado)
			if err != nil {
				fmt.Println("error ReemplazarBlItemContenedor:", err)
				continue
			}
			b, err := s.repo.GetUltimoPorId(itemActualizado.BlNroBl)
			if err != nil {
				fmt.Println("GetUltimoPorId AgruparBlItemImo:", *d.BlItemContenedorID)
			}
			bl, err := ReemplazarBlItem(b, *itemActualizado)
			if err != nil {
				fmt.Println("GetUltimoPorId ReemplazarBlItem   :", *d.BlItemContenedorID)
			}
			if d.Evento == "CREATE" && bl.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(bl.IDMongo.Hex(), bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlItemContenedorImo(d.IDMongo.Hex())

			fmt.Println("BL Item Contenedor imo agregado a BL ID:", *d.BlItemContenedorID)
		}
	}

}
func (s *service) AgruparBlItemContenedorSello() {
	fmt.Println("agrupando AgruparBlItemContenedorSello...", time.Now().Format("2006-01-02T15:04:05"))
	fmt.Println(time.Now().Format("2006-01-02T15:04:05"))
	imos, err := s.GetBlItemContenedorSelloAll()
	if err != nil {
		fmt.Println("error GetBlItemContenedorSelloAll:", err)
		return
	}

	for _, d := range imos {
		if d.BlItemContenedorID != nil {
			cnt, err := s.repo.GetUltimoPorIdItemContenedor(d.BlItemContenedorID)
			if err != nil {
				fmt.Println("error GetUltimoPorIdItemContenedor:", err)
				continue
			}
			if cnt == nil {
				fmt.Println("no se encontró BL Item con item id:", *d.BlItemContenedorID)
				continue
			}
			cntActualizado, err := ReemplazarBlItemContenedorSello(cnt, *d)
			if err != nil {
				fmt.Println("error ReemplazarBlItemContenedor:", err)
				continue
			}
			item, err := s.repo.GetUltimoPorIdItem(cntActualizado.BlItemID)
			if err != nil {
				fmt.Println("error GetUltimoPorId:", err)
				continue
			}

			itemActualizado, err := ReemplazarBlItemContenedor(item, *cntActualizado)
			if err != nil {
				fmt.Println("error ReemplazarBlItemContenedor:", err)
				continue
			}
			b, err := s.repo.GetUltimoPorId(itemActualizado.BlNroBl)
			if err != nil {
				fmt.Println("GetUltimoPorId AgruparBlItemImo:", *d.BlItemContenedorID)
			}
			bl, err := ReemplazarBlItem(b, *itemActualizado)
			if err != nil {
				fmt.Println("GetUltimoPorId ReemplazarBlItem   :", *d.BlItemContenedorID)
			}
			if d.Evento == "CREATE" && bl.Evento == "CREATE" {
				_, err = s.repo.ActualizaBl(bl.IDMongo.Hex(), bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			} else {

				_, err = s.repo.Guardar(*bl)
				if err != nil {
					fmt.Println("error Guardar BL:", err)
					continue
				}
			}
			s.BorraBlItemContenedorSello(d.IDMongo.Hex())

			fmt.Println("BL Item Contenedor Sello agregado a BL ID:", *d.BlItemContenedorID)
		}
	}
}

func (s *service) AgruparBl() {
	c := cron.New(
		cron.WithSeconds(),
		cron.WithChain(
			cron.DelayIfStillRunning(cron.DefaultLogger), // espera a que termine
			cron.Recover(cron.DefaultLogger),             // recupera panics
		),
	)

	// saca el cuerpo del job a un método para reutilizar
	job := cron.FuncJob(func() {
		s.AgruparBlFecha()
		s.AgruparBlFlete()
		s.AgruparBlItem()
		s.AgruparBlReferencia()
		s.AgruparBlTransbordo()
		s.AgruparBlTransporte()
		s.AgruparBlLocacion()
		s.AgruparBlObservacion()
		s.AgruparBlParticipante()
		s.AgruparBlItemImo()
		s.AgruparBlItemContenedor()
		s.AgruparBlItemContenedorImo()
		s.AgruparBlItemContenedorSello()
	})

	_, err := c.AddJob("0/3 * * * * *", job) // cada 50 min al segundo 0
	if err != nil {
		fmt.Println("Error al agregar tarea:", err)
		return
	}
	c.Start()
	select {} // ma
}
