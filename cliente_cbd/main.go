package main

import (
	"cliente_cbc/src/config"
	"cliente_cbc/src/service"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	l := config.InitLogger()
	f, err := os.OpenFile("logfile.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	cn := config.Newconfig()
	httpClient := &http.Client{
		Timeout: time.Duration(30) * time.Second,
	}

	service := service.NewService(l, *httpClient, *cn)
	/*tablas padres*/
	go service.EscuchaMsjCDCBultos()
	go service.EscuchaMsjCDCManifiesto()
	go service.EscuchaMsjCDCPapeletaRecepcion()
	go service.EscuchaMsjCDCVisacion()
	go service.EscuchaMsjCDCDespacho()
	go service.EscuchaMsjCDCFactura()
	go service.EscuchaMsjCDCPapeletaExpo()
	go service.EscuchaMsjCDCBl()
	go service.EscuchaMsjCDCNotaCredito()

	/*tablas hijas*/
	go service.EscuchaMsjCDCPapeletaRecepcionDetalle()
	go service.EscuchaMsjCDCFacturaDetalle()
	go service.EscuchaMsjCDCMercacias()
	go service.EscuchaMsjCDCPapeletaExpoDetalle()
	go service.EscuchaMsjCDCBlFecha()
	go service.EscuchaMsjCDCBlFlete()
	go service.EscuchaMsjCDCBlItem()
	go service.EscuchaMsjCDCBlLocacion()
	go service.EscuchaMsjCDCBlObservacion()
	go service.EscuchaMsjCDCBlParticipante()
	go service.EscuchaMsjCDCBlReferencia()
	go service.EscuchaMsjCDCBlTransbordo()
	go service.EscuchaMsjCDCBlTransporte()
	go service.EscuchaMsjCDCNotaCreditoServicios()

	/*tablas nietos*/
	go service.EscuchaMsjCDCBlItemImo()
	go service.EscuchaMsjCDCBlItemContenedor()
	/* tablas bisnietos*/
	go service.EscuchaMsjCDCBlItemContenedorSello()
	go service.EscuchaMsjCDCBlItemContenedorImo()
	select {}
}
