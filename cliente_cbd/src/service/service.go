package service

import (
	"bytes"
	"cliente_cbc/src/config"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/segmentio/kafka-go"
)

type (
	Service interface {
		EscuchaMsjCDCPapeletaRecepcion()
		EscuchaMsjCDCPapeletaRecepcionDetalle()
		EscuchaMsjCDCVisacion()
		EscuchaMsjCDCBultos()
		EscuchaMsjCDCManifiesto()
		EscuchaMsjCDCMercacias()
		EscuchaMsjCDCDespacho()
		EscuchaMsjCDCFactura()
		EscuchaMsjCDCFacturaDetalle()
		EscuchaMsjCDCPapeletaExpo()
		EscuchaMsjCDCPapeletaExpoDetalle()
		EscuchaMsjCDCBl()
		EscuchaMsjCDCBlFecha()
		EscuchaMsjCDCBlFlete()
		EscuchaMsjCDCBlItemImo()
		EscuchaMsjCDCBlItem()
		EscuchaMsjCDCBlItemContenedor()
		EscuchaMsjCDCBlItemContenedorSello()
		EscuchaMsjCDCBlItemContenedorImo()
		EscuchaMsjCDCBlLocacion()
		EscuchaMsjCDCBlObservacion()
		EscuchaMsjCDCBlParticipante()
		EscuchaMsjCDCBlReferencia()
		EscuchaMsjCDCBlTransbordo()
		EscuchaMsjCDCBlTransporte()
		EscuchaMsjCDCNotaCredito()
		EscuchaMsjCDCNotaCreditoServicios()
	}
	service struct {
		log        *log.Logger
		httpClient http.Client
		config     config.Config
	}
)

func NewService(log *log.Logger, httpClient http.Client, config config.Config) Service {
	return &service{
		log:        log,
		httpClient: httpClient,
		config:     config,
	}
}

func (s *service) EscuchaMsjCDCPapeletaRecepcion() {
	s.escucharKafka(
		"Papeleta Recepcion",
		s.config.TopicPapeletaRecepcion,
		s.config.Url+"/api/papeleta_recepcion",
	)
}

func (s *service) EscuchaMsjCDCPapeletaRecepcionDetalle() {
	s.escucharKafka(
		"Papeleta Recepcion Detalle",
		s.config.TopicPapeletaRecepcionDetalle,
		s.config.Url+"/api/papeleta_recepcion_detalle",
	)
}

func (s *service) EscuchaMsjCDCVisacion() {
	s.escucharKafka(
		"Visacion",
		s.config.TopicVisacion,
		s.config.Url+"/api/visacion",
	)
}
func (s *service) EscuchaMsjCDCBultos() {
	s.escucharKafka(
		"Bultos",
		s.config.TopicBlItemTipoBulto,
		s.config.Url+"/api/bultos",
	)
}
func (s *service) EscuchaMsjCDCManifiesto() {
	s.escucharKafka(
		"Manifiesto",
		s.config.TopicManifiesto,
		s.config.Url+"/api/manifiesto",
	)
}
func (s *service) EscuchaMsjCDCMercacias() {
	s.escucharKafka(
		"Mercacias_despachadas",
		s.config.TopicMercanciasDespachadas,
		s.config.Url+"/api/visacion-mercancias",
	)
}
func (s *service) EscuchaMsjCDCDespacho() {
	s.escucharKafka(
		"despacho",
		s.config.TopicDespacho,
		s.config.Url+"/api/despacho",
	)
}
func (s *service) EscuchaMsjCDCFactura() {
	s.escucharKafka(
		"factura",
		s.config.TopicFactura,
		s.config.Url+"/api/factura",
	)
}
func (s *service) EscuchaMsjCDCFacturaDetalle() {
	s.escucharKafka(
		"factura_detalle",
		s.config.TopicDetalleFactura,
		s.config.Url+"/api/factura-detalle",
	)
}
func (s *service) EscuchaMsjCDCPapeletaExpo() {
	s.escucharKafka(
		"papeleta_expo",
		s.config.TopicPapeletaRecepcionExpo,
		s.config.Url+"/api/papeleta-expo",
	)
}
func (s *service) EscuchaMsjCDCPapeletaExpoDetalle() {
	s.escucharKafka(
		"papeleta_expo_detalle",
		s.config.TopicPapeletaRecepcionExpoDetalle,
		s.config.Url+"/api/papeleta-expo-detalle",
	)
}
func (s *service) EscuchaMsjCDCBl() {
	s.escucharKafka(
		"bl",
		s.config.TopicBl,
		s.config.Url+"/api/bl",
	)
}
func (s *service) EscuchaMsjCDCBlFecha() {
	s.escucharKafka(
		"bl_fecha",
		s.config.TopicBlFecha,
		s.config.Url+"/api/bl-fecha",
	)
}
func (s *service) EscuchaMsjCDCBlFlete() {
	s.escucharKafka(
		"bl_flete",
		s.config.TopicBlFlete,
		s.config.Url+"/api/bl-flete",
	)
}
func (s *service) EscuchaMsjCDCBlItemImo() {
	s.escucharKafka(
		"bl_item_imo",
		s.config.TopicBlItemImo,
		s.config.Url+"/api/bl-item-imo",
	)
}
func (s *service) EscuchaMsjCDCBlItem() {
	s.escucharKafka(
		"bl_item",
		s.config.TopicBlItem,
		s.config.Url+"/api/bl-item",
	)
}
func (s *service) EscuchaMsjCDCBlItemContenedor() {
	s.escucharKafka(
		"bl_item_contenedor",
		s.config.TopicBlItemContenedor,
		s.config.Url+"/api/bl-item-contenedor",
	)
}
func (s *service) EscuchaMsjCDCBlItemContenedorSello() {
	s.escucharKafka(
		"bl_item_contenedor_sello",
		s.config.TopicBlItemContenedorSello,
		s.config.Url+"/api/bl-item-contenedor-sello",
	)
}
func (s *service) EscuchaMsjCDCBlItemContenedorImo() {
	s.escucharKafka(
		"bl_item_contenedor_imo",
		s.config.TopicBlItemContenedorImo,
		s.config.Url+"/api/bl-item-contenedor-imo",
	)
}
func (s *service) EscuchaMsjCDCBlLocacion() {
	s.escucharKafka(
		"bl_locacion",
		s.config.TopicBlLocacion,
		s.config.Url+"/api/bl-locacion",
	)
}
func (s *service) EscuchaMsjCDCBlObservacion() {
	s.escucharKafka(
		"bl_observacion",
		s.config.TopicBlObservacion,
		s.config.Url+"/api/bl-observacion",
	)
}
func (s *service) EscuchaMsjCDCBlParticipante() {
	s.escucharKafka(
		"bl_participante",
		s.config.TopicBlParticipante,
		s.config.Url+"/api/bl-participante",
	)
}
func (s *service) EscuchaMsjCDCBlReferencia() {
	s.escucharKafka(
		"bl_referencia",
		s.config.TopicBlReferencia,
		s.config.Url+"/api/bl-referencia",
	)
}
func (s *service) EscuchaMsjCDCBlTransbordo() {
	s.escucharKafka(
		"bl_transbordo",
		s.config.TopicBlTransbordo,
		s.config.Url+"/api/bl-transbordo",
	)
}
func (s *service) EscuchaMsjCDCBlTransporte() {
	s.escucharKafka(
		"bl_transporte",
		s.config.TopicBlTransporte,
		s.config.Url+"/api/bl-transporte",
	)
}
func (s *service) EscuchaMsjCDCNotaCredito() {
	s.escucharKafka(
		"nota_credito",
		s.config.TopicNotaCredito,
		s.config.Url+"/api/nota_credito",
	)
}
func (s *service) EscuchaMsjCDCNotaCreditoServicios() {
	s.escucharKafka(
		"nota_credito_servicios",
		s.config.TopicNotaCreditoServ,
		s.config.Url+"/api/nota-credito-servicio",
	)
}

func (s *service) escucharKafka(serviceName, topic, apiEndpoint string) {
	fmt.Printf("Escuchando mensajes de %s...\n", serviceName)
	fmt.Println("Configuración:", s.config.Broker)
	fmt.Println("Configuración:", topic)

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{s.config.Broker},
		Topic:     topic,
		GroupID:   "prod_texval_test_huerfanos",
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
		MaxWait:   1 * time.Second,
	})

	defer r.Close()

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			s.log.Printf("Error leyendo mensaje: %v", err)
			continue
		}

		if err := s.RegistraApi(apiEndpoint, m.Value); err != nil {
			s.log.Printf("Error al registrar en API: %v", err)
		}
	}
}

func (s *service) RegistraApi(url string, body []byte) error {
	fmt.Println("Enviando a la API:", string(body))

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		s.log.Println(err)
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		s.log.Printf("Error al enviar solicitud HTTP: %v", err)
		fmt.Println("Error al enviar solicitud HTTP: ", err)
		return err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		s.log.Printf("Error al leer el cuerpo de la respuesta: %v", err)
		fmt.Println("Error al leer el cuerpo de la respuesta: ", err)
		return err
	}

	s.log.Println(string(bodyBytes))

	return nil
}
