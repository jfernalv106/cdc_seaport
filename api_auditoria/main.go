package main

import (
	"api_auditoria/src/bl"
	"api_auditoria/src/bultos"
	"api_auditoria/src/config"
	"api_auditoria/src/db"
	"api_auditoria/src/despacho"
	"api_auditoria/src/expo"
	"api_auditoria/src/factura"
	"api_auditoria/src/handler"
	"api_auditoria/src/manifiesto"
	notacredito "api_auditoria/src/nota_credito"
	"api_auditoria/src/papeleta"
	"api_auditoria/src/visacion"
	"context"
	"log"
	"net/http"
	"os"
)

func main() {

	l := db.InitLogger()
	f, err := os.OpenFile("logfile", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)

	database := config.NewConfig()

	var collection = db.Connect("auditoria", database)
	var ctx = context.Background()
	// Inicialización de repositorios y servicios para Bultos
	bultoRepo := bultos.NewRepo(l, collection, &ctx)
	bultoService := bultos.NewService(l, bultoRepo)
	bultoEndpoints := bultos.MakeEndPoint(bultoService)
	// Inicialización de repositorios y servicios para papeleta
	papeletaRepo := papeleta.NewRepo(l, collection, &ctx)
	papeletaService := papeleta.NewService(l, papeletaRepo, bultoRepo)
	go papeletaService.AgruparPapeletaDetallePorIdPapeleta()
	papeletaEndpoints := papeleta.MakeEndPoint(papeletaService)

	// Inicialización de repositorios y servicios para visación
	visacionRepo := visacion.NewRepo(l, collection, &ctx)
	visacionService := visacion.NewService(l, visacionRepo)
	go visacionService.AgruparMercancias()
	visacionEndpoints := visacion.MakeEndPoint(visacionService)
	// Inicialización de repositorios y servicios para manifiesto
	manifiestoRepo := manifiesto.NewRepo(l, collection, &ctx)
	manifiestoService := manifiesto.NewService(l, manifiestoRepo)
	manifiestoEndpoints := manifiesto.MakeEndPoint(manifiestoService)
	// Inicialización de repositorios y servicios para despacho
	despachoRepo := despacho.NewRepo(l, collection, &ctx)
	despachoService := despacho.NewService(l, despachoRepo)
	despachoEndpoints := despacho.MakeEndPoint(despachoService)
	// Inicialización de repositorios y servicios para factura
	facturaRepo := factura.NewRepo(l, collection, &ctx)
	facturaService := factura.NewService(l, facturaRepo)
	go facturaService.AgruparFacturaDetallePorIdFactura()
	facturaEndpoints := factura.MakeEndPoint(facturaService)
	// Inicialización de repositorios y servicios para papeleta expo
	papeletaExpoRepo := expo.NewRepo(l, collection, &ctx)
	papeletaExpoService := expo.NewService(l, papeletaExpoRepo)
	go papeletaExpoService.AgruparPapeletaDetallePorIdPapeleta()
	papeletaExpoEndpoints := expo.MakeEndPoint(papeletaExpoService)
	// Inicialización de repositorios y servicios para papeleta expo
	notaCreditoExpoRepo := notacredito.NewRepo(l, collection, &ctx)
	notaCreditoExpoService := notacredito.NewService(l, notaCreditoExpoRepo)
	go notaCreditoExpoService.AgruparNotaCreditoServicioPorIdNC()
	notaCreditoExpoEndpoints := notacredito.MakeEndPoint(notaCreditoExpoService)
	// Inicialización de repositorios y servicios para papeleta expo
	blRepo := bl.NewRepository(collection, &ctx)
	blService := bl.NewService(l, blRepo)
	go blService.AgruparBl()

	blEndpoints := bl.MakeEndPoint(blService)
	h := handler.NewRouter(ctx, papeletaEndpoints, visacionEndpoints, bultoEndpoints, manifiestoEndpoints, despachoEndpoints, facturaEndpoints, papeletaExpoEndpoints, notaCreditoExpoEndpoints, blEndpoints)

	srv := &http.Server{
		Addr:    ":8181",
		Handler: h,
	}
	log.Fatal(srv.ListenAndServe())
}
