package visacion

import (
	"api_auditoria/src/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReemplazarMercacia(visacion *model.Visacion, nuevaMercancia *model.MercanciasDespachada) (*model.Visacion, error) {
	if visacion == nil {
		return nil, fmt.Errorf("la factura no puede ser nula")
	}
	if visacion.Mercacias == nil {
		fmt.Println("Inicializando detalles de la papeleta recepcion")
		visacion.Mercacias = &[]model.MercanciasDespachada{}
	}

	if nuevaMercancia == nil {
		return nil, fmt.Errorf("el nuevo detalle no puede ser nulo")
	}
	if visacion.Evento != "CREATE" && nuevaMercancia.Evento != "CREATE" {
		visacion.IDMongo = primitive.NewObjectID()
	}

	if nuevaMercancia.Evento == "CREATE" || len(*visacion.Mercacias) == 0 {
		*visacion.Mercacias = append(*visacion.Mercacias, *nuevaMercancia)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nuevaMercancia.Evento == "DELETE" {
		var nuevosDetalles []model.MercanciasDespachada

		for _, detalle := range *visacion.Mercacias {
			if *detalle.ID != *nuevaMercancia.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		visacion.Mercacias = &nuevosDetalles

	}
	if nuevaMercancia.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *visacion.Mercacias {
			if *detalle.ID == *nuevaMercancia.ID {
				(*visacion.Mercacias)[i] = *nuevaMercancia
				encontrado = true
				break
			}
		}
		if !encontrado {
			*visacion.Mercacias = append(*visacion.Mercacias, *nuevaMercancia)
		}
	}

	if nuevaMercancia.Evento != "DELETE" {
		visacion.Evento = nuevaMercancia.Evento
	}
	visacion.FechaEvento = nuevaMercancia.FechaEvento
	return visacion, nil
}
