package papeleta

import (
	"api_auditoria/src/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReemplazarDetalle(papeleta *model.PapeletaRecepcion, nuevoDetalle *model.PapeletaRecepcionDetalle) (*model.PapeletaRecepcion, error) {
	if papeleta == nil {
		return nil, fmt.Errorf("la factura no puede ser nula")
	}
	if papeleta.Detalles == nil {
		fmt.Println("Inicializando detalles de la papeleta recepcion")
		papeleta.Detalles = &[]model.PapeletaRecepcionDetalle{}
	}

	if nuevoDetalle == nil {
		return nil, fmt.Errorf("el nuevo detalle no puede ser nulo")
	}
	if papeleta.Evento != "CREATE" || nuevoDetalle.Evento != "CREATE" {
		papeleta.IDMongo = primitive.NewObjectID()
	}

	if nuevoDetalle.Evento == "CREATE" || len(*papeleta.Detalles) == 0 {
		*papeleta.Detalles = append(*papeleta.Detalles, *nuevoDetalle)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nuevoDetalle.Evento == "DELETE" {
		var nuevosDetalles []model.PapeletaRecepcionDetalle

		for _, detalle := range *papeleta.Detalles {
			if *detalle.ID != *nuevoDetalle.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		papeleta.Detalles = &nuevosDetalles

	}
	if nuevoDetalle.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *papeleta.Detalles {
			if *detalle.ID == *nuevoDetalle.ID {
				(*papeleta.Detalles)[i] = *nuevoDetalle
				encontrado = true
				break
			}
		}
		if !encontrado {
			*papeleta.Detalles = append(*papeleta.Detalles, *nuevoDetalle)
		}
	}

	papeleta.Evento = nuevoDetalle.Evento
	if nuevoDetalle.Evento == "DELETE" {
		papeleta.Evento = "UPDATE"
	}
	papeleta.FechaEvento = nuevoDetalle.FechaEvento
	return papeleta, nil
}
