package bl

import (
	"api_auditoria/src/model"
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ================== BL FECHAS ==================
func ReemplazarBlFecha(bl *model.BL, nueva model.BlFecha) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlFechas == nil {
		fmt.Println("Inicializando BL FECHAS ")
		bl.BlFechas = &[]model.BlFecha{}
	}
	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlFechas) == 0 {
		*bl.BlFechas = append(*bl.BlFechas, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlFecha

		for _, detalle := range *bl.BlFechas {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlFechas = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlFechas {
			if *detalle.ID == *nueva.ID {
				(*bl.BlFechas)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlFechas = append(*bl.BlFechas, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL FLETES ==================
func ReemplazarBlFlete(bl *model.BL, nueva model.BlFlete) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlFletes == nil {
		fmt.Println("Inicializando BL FLETES ")
		bl.BlFletes = &[]model.BlFlete{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlFletes) == 0 {
		*bl.BlFletes = append(*bl.BlFletes, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlFlete

		for _, detalle := range *bl.BlFletes {
			if *detalle.BlNroBl != *nueva.BlNroBl {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlFletes = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlFletes {
			if *detalle.BlNroBl == *nueva.BlNroBl {
				(*bl.BlFletes)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlFletes = append(*bl.BlFletes, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL ITEMS ==================
func ReemplazarBlItem(bl *model.BL, nueva model.BlItem) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlItems == nil {
		fmt.Println("Inicializando BL ITEMS ")
		bl.BlItems = &[]model.BlItem{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlItems) == 0 {
		*bl.BlItems = append(*bl.BlItems, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlItem

		for _, detalle := range *bl.BlItems {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlItems = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlItems {
			if *detalle.ID == *nueva.ID {
				(*bl.BlItems)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlItems = append(*bl.BlItems, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL ITEM IMO ==================
func ReemplazarBlItemImo(item *model.BlItem, nueva model.BlItemImo) (*model.BlItem, error) {
	if item == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if item.BlItemImos == nil {
		fmt.Println("Inicializando BL ITEM IMO ")
		item.BlItemImos = &[]model.BlItemImo{}
	}

	if nueva.Evento == "CREATE" || len(*item.BlItemImos) == 0 {
		*item.BlItemImos = append(*item.BlItemImos, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlItemImo

		for _, detalle := range *item.BlItemImos {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		item.BlItemImos = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *item.BlItemImos {
			if *detalle.ID == *nueva.ID {
				(*item.BlItemImos)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*item.BlItemImos = append(*item.BlItemImos, nueva)
		}
	}

	item.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		item.Evento = "UPDATE"
	}
	item.FechaEvento = nueva.FechaEvento
	return item, nil
}

// ================== BL ITEM CONTENEDOR ==================
func ReemplazarBlItemContenedor(item *model.BlItem, nueva model.BlItemContenedor) (*model.BlItem, error) {
	if item == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if item.BlItemContenedores == nil {
		fmt.Println("Inicializando BL ITEM IMO ")
		item.BlItemContenedores = &[]model.BlItemContenedor{}
	}

	if nueva.Evento == "CREATE" || len(*item.BlItemContenedores) == 0 {
		*item.BlItemContenedores = append(*item.BlItemContenedores, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlItemContenedor

		for _, detalle := range *item.BlItemContenedores {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		item.BlItemContenedores = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *item.BlItemContenedores {
			if *detalle.ID == *nueva.ID {
				(*item.BlItemContenedores)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*item.BlItemContenedores = append(*item.BlItemContenedores, nueva)
		}
	}

	item.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		item.Evento = "UPDATE"
	}
	item.FechaEvento = nueva.FechaEvento
	return item, nil
}

// ================== BL ITEM CONTENEDOR IMO ==================
func ReemplazarBlItemContenedorImo(cnt *model.BlItemContenedor, nueva model.BlItemContenedorImo) (*model.BlItemContenedor, error) {
	if cnt == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if cnt.BlItemContenedorImos == nil {
		fmt.Println("Inicializando BL ITEM CONTENEDOR IMO ")
		cnt.BlItemContenedorImos = &[]model.BlItemContenedorImo{}
	}

	if nueva.Evento == "CREATE" || len(*cnt.BlItemContenedorImos) == 0 {
		*cnt.BlItemContenedorImos = append(*cnt.BlItemContenedorImos, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlItemContenedorImo

		for _, detalle := range *cnt.BlItemContenedorImos {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		cnt.BlItemContenedorImos = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *cnt.BlItemContenedorImos {
			if *detalle.ID == *nueva.ID {
				(*cnt.BlItemContenedorImos)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*cnt.BlItemContenedorImos = append(*cnt.BlItemContenedorImos, nueva)
		}
	}

	cnt.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		cnt.Evento = "UPDATE"
	}
	cnt.FechaEvento = nueva.FechaEvento
	return cnt, nil
}

// ================== BL ITEM CONTENEDOR SELLO ==================
func ReemplazarBlItemContenedorSello(cnt *model.BlItemContenedor, nueva model.BlItemContenedorSello) (*model.BlItemContenedor, error) {
	if cnt == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if cnt.BlItemContenedorSellos == nil {
		fmt.Println("Inicializando BL ITEM CONTENEDOR IMO ")
		cnt.BlItemContenedorSellos = &[]model.BlItemContenedorSello{}
	}

	if nueva.Evento == "CREATE" || len(*cnt.BlItemContenedorSellos) == 0 {
		*cnt.BlItemContenedorSellos = append(*cnt.BlItemContenedorSellos, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlItemContenedorSello

		for _, detalle := range *cnt.BlItemContenedorSellos {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		cnt.BlItemContenedorSellos = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *cnt.BlItemContenedorSellos {
			if *detalle.ID == *nueva.ID {
				(*cnt.BlItemContenedorSellos)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*cnt.BlItemContenedorSellos = append(*cnt.BlItemContenedorSellos, nueva)
		}
	}

	cnt.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		cnt.Evento = "UPDATE"
	}
	cnt.FechaEvento = nueva.FechaEvento
	return cnt, nil
}

// ================== BL PARTICIPANTES ==================
func ReemplazarBlParticipante(bl *model.BL, nueva model.BlParticipante) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlParticipantes == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlParticipantes = &[]model.BlParticipante{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlParticipantes) == 0 {
		*bl.BlParticipantes = append(*bl.BlParticipantes, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlParticipante

		for _, detalle := range *bl.BlParticipantes {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlParticipantes = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlParticipantes {
			if *detalle.ID == *nueva.ID {
				(*bl.BlParticipantes)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlParticipantes = append(*bl.BlParticipantes, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL OBSERVACIONES ==================
func ReemplazarBlObservacion(bl *model.BL, nueva model.BlObservacion) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlObservaciones == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlObservaciones = &[]model.BlObservacion{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlObservaciones) == 0 {
		*bl.BlObservaciones = append(*bl.BlObservaciones, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlObservacion

		for _, detalle := range *bl.BlObservaciones {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlObservaciones = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlObservaciones {
			if *detalle.ID == *nueva.ID {
				(*bl.BlObservaciones)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlObservaciones = append(*bl.BlObservaciones, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL LOCACIONES ==================
func ReemplazarBlLocacion(bl *model.BL, nueva model.BlLocacion) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlLocaciones == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlLocaciones = &[]model.BlLocacion{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlLocaciones) == 0 {
		*bl.BlLocaciones = append(*bl.BlLocaciones, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlLocacion

		for _, detalle := range *bl.BlLocaciones {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlLocaciones = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlLocaciones {
			if *detalle.ID == *nueva.ID {
				(*bl.BlLocaciones)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlLocaciones = append(*bl.BlLocaciones, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL REFERENCIAS ==================
func ReemplazarBlReferencia(bl *model.BL, nueva model.BlReferencia) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlReferencias == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlReferencias = &[]model.BlReferencia{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlReferencias) == 0 {
		*bl.BlReferencias = append(*bl.BlReferencias, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlReferencia

		for _, detalle := range *bl.BlReferencias {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlReferencias = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlReferencias {
			if *detalle.ID == *nueva.ID {
				(*bl.BlReferencias)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlReferencias = append(*bl.BlReferencias, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL TRANSBORDOS ==================
func ReemplazarBlTransbordo(bl *model.BL, nueva model.BlTransbordo) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlTransbordos == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlTransbordos = &[]model.BlTransbordo{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlTransbordos) == 0 {
		*bl.BlTransbordos = append(*bl.BlTransbordos, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlTransbordo

		for _, detalle := range *bl.BlTransbordos {
			if *detalle.ID != *nueva.ID {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlTransbordos = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlTransbordos {
			if *detalle.ID == *nueva.ID {
				(*bl.BlTransbordos)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlTransbordos = append(*bl.BlTransbordos, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}

// ================== BL TRANSPORTES ==================
func ReemplazarBlTransporte(bl *model.BL, nueva model.BlTransporte) (*model.BL, error) {
	if bl == nil {
		return nil, fmt.Errorf("El BL no puede ser nulo")
	}
	if bl.BlTransportes == nil {
		fmt.Println("Inicializando bl_fechas ")
		bl.BlTransportes = &[]model.BlTransporte{}
	}

	if nueva.Evento != "CREATE" || bl.Evento != "CREATE" {
		bl.IDMongo = primitive.NewObjectID()
	}
	if nueva.Evento == "CREATE" || len(*bl.BlTransportes) == 0 {
		*bl.BlTransportes = append(*bl.BlTransportes, nueva)
	}
	// Si el evento es DELETE, eliminar el detalle con el ID correspondiente
	if nueva.Evento == "DELETE" {
		var nuevosDetalles []model.BlTransporte

		for _, detalle := range *bl.BlTransportes {
			if *detalle.BlNroBl != *nueva.BlNroBl {
				nuevosDetalles = append(nuevosDetalles, detalle)
				continue
			}
		}
		bl.BlTransportes = &nuevosDetalles

	}
	if nueva.Evento == "UPDATE" {
		var encontrado bool = false
		// Reemplazar el detalle existente con el nuevo detalle basado en el ID
		for i, detalle := range *bl.BlTransportes {
			if *detalle.BlNroBl == *nueva.BlNroBl {
				(*bl.BlTransportes)[i] = nueva
				encontrado = true
				break
			}
		}
		if !encontrado {
			*bl.BlTransportes = append(*bl.BlTransportes, nueva)
		}
	}

	bl.Evento = nueva.Evento
	if nueva.Evento == "DELETE" {
		bl.Evento = "UPDATE"
	}
	bl.FechaEvento = nueva.FechaEvento
	return bl, nil
}
