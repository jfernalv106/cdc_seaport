package factura

import (
	"api_auditoria/src/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReemplazarDetalle(factura *model.Factura, nuevoDetalle *model.FacturaDetalle) (*model.Factura, error) {
	if factura == nil {
		return nil, fmt.Errorf("la factura no puede ser nula")
	}
	if factura.Detalles == nil {
		fmt.Println("Inicializando detalles de la factura")
		factura.Detalles = &[]model.FacturaDetalle{}
	}

	if nuevoDetalle == nil {
		return nil, fmt.Errorf("el nuevo detalle no puede ser nulo")
	}

	if factura.Evento != "CREATE" || (nuevoDetalle.Evento == nil || *nuevoDetalle.Evento != "CREATE") {
		factura.IDMongo = primitive.NewObjectID()
	}

	if nuevoDetalle.Evento != nil && *nuevoDetalle.Evento == "CREATE" || len(*factura.Detalles) == 0 {
		*factura.Detalles = append(*factura.Detalles, *nuevoDetalle)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nuevoDetalle.Evento != nil && *nuevoDetalle.Evento == "DELETE" {
		var nuevosDetalles []model.FacturaDetalle

		for _, detalle := range *factura.Detalles {
			if *detalle.ID != *nuevoDetalle.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		factura.Detalles = &nuevosDetalles

	}
	if nuevoDetalle.Evento != nil && *nuevoDetalle.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *factura.Detalles {
			if *detalle.ID == *nuevoDetalle.ID {
				(*factura.Detalles)[i] = *nuevoDetalle
				encontrado = true
				break
			}
		}
		if !encontrado {
			*factura.Detalles = append(*factura.Detalles, *nuevoDetalle)
		}
	}

	factura.Evento = *nuevoDetalle.Evento
	if nuevoDetalle.Evento != nil && *nuevoDetalle.Evento == "DELETE" {
		factura.Evento = "UPDATE"
	}
	factura.FechaEvento = nuevoDetalle.FechaEvento
	return factura, nil
}
