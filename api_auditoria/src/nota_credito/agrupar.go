package notacredito

import (
	"api_auditoria/src/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ReemplazarDetalle(nota *model.NotaCredito, nuevoServicio *model.NotaCreditoServicio) (*model.NotaCredito, error) {
	if nota == nil {
		return nil, fmt.Errorf("la Nota no puede ser nula")
	}
	if nota.Servicios == nil {
		fmt.Println("Inicializando servicios de la nota de crédito")
		nota.Servicios = &[]model.NotaCreditoServicio{}
	}

	if nuevoServicio == nil {
		return nil, fmt.Errorf("el nuevo Servicio no puede ser nulo")
	}
	if nota.Evento != "CREATE" && nuevoServicio.Evento != "CREATE" {
		nota.IDMongo = primitive.NewObjectID()
	}

	if nuevoServicio.Evento == "CREATE" || len(*nota.Servicios) == 0 {
		*nota.Servicios = append(*nota.Servicios, *nuevoServicio)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nuevoServicio.Evento == "DELETE" {
		var nuevoServicios []model.NotaCreditoServicio

		for _, detalle := range *nota.Servicios {
			if *detalle.ID != *nuevoServicio.ID {
				nuevoServicios = append(nuevoServicios, detalle)
				continue
			}
		}
		nota.Servicios = &nuevoServicios

	}
	if nuevoServicio.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *nota.Servicios {
			if *detalle.ID == *nuevoServicio.ID {
				(*nota.Servicios)[i] = *nuevoServicio
				encontrado = true
				break
			}
		}
		if !encontrado {
			*nota.Servicios = append(*nota.Servicios, *nuevoServicio)
		}
	}

	if nuevoServicio.Evento != "DELETE" {
		nota.Evento = nuevoServicio.Evento
	}
	nota.FechaEvento = nuevoServicio.FechaEvento
	return nota, nil
}
