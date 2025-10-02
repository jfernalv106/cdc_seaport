package bl

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"context"
	"errors"
	"fmt"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	Guardar(bl interface{}) (interface{}, error)
	ActualizaBl(id string, bl interface{}) (interface{}, error)
	GuardarParticipante(blParticipante interface{}) (interface{}, error)
	GetByID(id primitive.ObjectID) (*model.BL, error)
	GetByIDString(id string) (*model.BL, error)
	GetByNroBl(id *int64, nroBl *string, manifiesto *string) ([]*model.BL, error)
	GetAll(filter map[string]interface{}) ([]*model.BL, error)
	GetUltimoPorNroBl(nroBl string) (*model.BL, error)
	GetUltimoPorId(id *int64) (*model.BL, error)
	GetUltimoPorIdItem(id *int64) (*model.BlItem, error)
	GetUltimoPorIdItemContenedor(id *int64) (*model.BlItemContenedor, error)
	GuardarBlFecha(blFecha interface{}) (interface{}, error)
	GuardarBlFlete(blFlete interface{}) (interface{}, error)
	GuardarBlItem(blItem interface{}) (interface{}, error)
	GuardarBlItemImo(blItemImo interface{}) (interface{}, error)
	GuardarBlItemContenedor(blItemContenedor interface{}) (interface{}, error)
	GuardarBlItemContenedorImo(blItemContenedorImo interface{}) (interface{}, error)
	GuardarBlItemContenedorSello(blItemContenedorSello interface{}) (interface{}, error)
	GuardarBlLocacion(blLocacion interface{}) (interface{}, error)
	GuardarBlObservacion(blObservacion interface{}) (interface{}, error)
	GuardarBlParticipante(blParticipante interface{}) (interface{}, error)
	GuardarBlReferencia(blReferencia interface{}) (interface{}, error)
	GuardarBlTransbordo(blTransbordo interface{}) (interface{}, error)
	GuardarBlTransporte(blTransporte interface{}) (interface{}, error)
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
}

type repo struct {
	db  *mongo.Database
	ctx *context.Context
}

func NewRepository(db *mongo.Database, ctx *context.Context) Repository {
	return &repo{db: db, ctx: ctx}
}

// 👉 Siempre se inserta un documento nuevo
func (r *repo) Guardar(bl interface{}) (interface{}, error) {

	b, err := r.db.Collection("bl").InsertOne(*r.ctx, bl)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) ActualizaBl(id string, bl interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("bl").ReplaceOne(*r.ctx, bson.M{"_id": objID}, bl, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repo) GuardarParticipante(blParticipante interface{}) (interface{}, error) {

	b, err := r.db.Collection("bl_participante").InsertOne(*r.ctx, blParticipante)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) GetByID(id primitive.ObjectID) (*model.BL, error) {
	var bl model.BL
	if err := r.db.Collection("bl").FindOne(*r.ctx, bson.M{"_id": id}).Decode(&bl); err != nil {
		return nil, err
	}
	return &bl, nil
}

func (r *repo) GetByIDString(id string) (*model.BL, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("id inválido")
	}
	return r.GetByID(oid)
}

func (r *repo) GetByNroBl(id *int64, nroBl *string, manifiesto *string) ([]*model.BL, error) {
	filter := bson.M{}

	if id != nil {
		filter["id"] = *id
	}
	if nroBl != nil && *nroBl != "" {
		filter["nro_bl"] = *nroBl
	}
	if manifiesto != nil && *manifiesto != "" {
		filter["manifiesto_nro_mfto_interno"] = *manifiesto
	}

	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) GetAll(filter map[string]interface{}) ([]*model.BL, error) {
	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (r *repo) GetUltimoPorId(id *int64) (*model.BL, error) {
	filter := bson.M{}

	if id != nil {
		filter["id"] = *id
	}

	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró BL con id %d", *id)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}
func (r *repo) GetUltimoPorNroBl(nroBl string) (*model.BL, error) {

	filter := bson.M{}

	filter["nro_bl"] = nroBl

	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró BL con id %d", nroBl)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}

func (r *repo) GetUltimoPorIdItem(id *int64) (*model.BlItem, error) {
	if id == nil {
		return nil, fmt.Errorf("id no puede ser nil")
	}

	// 1) usar el valor, no el puntero
	filter := bson.M{"bl_items.id": *id}

	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró BL con bl_items.id=%d", *id)
	}

	// ordenar por fecha_evento desc
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	if result[0].BlItems == nil {
		return nil, fmt.Errorf("BL sin items")
	}
	items := *result[0].BlItems

	// 2) comparar valores (punteros seguros)
	for i := range items {
		if items[i].ID != nil && *items[i].ID == *id {
			// 3) devolver puntero al elemento del slice (no al range var)
			return &items[i], nil
		}
	}

	return nil, fmt.Errorf("no se encontró bl_items.id=%d dentro del BL más reciente", *id)
}

func (r *repo) GetUltimoPorIdItemContenedor(id *int64) (*model.BlItemContenedor, error) {

	if id == nil {
		return nil, fmt.Errorf("id no puede ser nil")
	}

	filter := bson.M{"bl_items.bl_item_contenedores.id": *id}

	cur, err := r.db.Collection("bl").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.BL
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró BL con bl_items.bl_item_contenedores.id=%d", *id)
	}

	// Ordena por fecha_evento desc del BL (mismo criterio que ya usas)
	sort.Slice(result, func(i, j int) bool {
		fi, _ := utils.ParseFecha(*result[i].FechaEvento)
		fj, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fi.After(fj)
	})

	// Buscar dentro del BL más reciente
	if result[0].BlItems == nil {
		return nil, fmt.Errorf("BL sin items")
	}
	items := *result[0].BlItems

	for i := range items {
		contsPtr := items[i].BlItemContenedores
		if contsPtr == nil {
			continue
		}
		conts := *contsPtr
		for j := range conts {
			// si tu campo es *int64
			if conts[j].ID != nil && *conts[j].ID == *id {
				// devolver puntero al elemento del slice (no a la var del range)
				return &conts[j], nil
			}
			// si fuera int64 no puntero: if conts[j].ID == *id { ... }
		}
	}

	return nil, fmt.Errorf("no se encontró bl_item_contenedores.id=%d dentro del BL más reciente", *id)
}

func (r *repo) GuardarBlFecha(blFecha interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_fecha").InsertOne(*r.ctx, blFecha)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlFlete(blFlete interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_flete").InsertOne(*r.ctx, blFlete)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlItem(blItem interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_item").InsertOne(*r.ctx, blItem)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlItemImo(blItemImo interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_item_imo").InsertOne(*r.ctx, blItemImo)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlItemContenedor(blItemContenedor interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_item_contenedor").InsertOne(*r.ctx, blItemContenedor)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlItemContenedorImo(blItemContenedorImo interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_contenedor_imo").InsertOne(*r.ctx, blItemContenedorImo)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlItemContenedorSello(blItemContenedorSello interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_contenedor_sello").InsertOne(*r.ctx, blItemContenedorSello)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlLocacion(blLocacion interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_locacion").InsertOne(*r.ctx, blLocacion)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlObservacion(blObservacion interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_observacion").InsertOne(*r.ctx, blObservacion)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlParticipante(blParticipante interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_participante").InsertOne(*r.ctx, blParticipante)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlReferencia(blReferencia interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_referencia").InsertOne(*r.ctx, blReferencia)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlTransbordo(blTransbordo interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_transbordo").InsertOne(*r.ctx, blTransbordo)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarBlTransporte(blTransporte interface{}) (interface{}, error) {
	b, err := r.db.Collection("bl_transporte").InsertOne(*r.ctx, blTransporte)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) GetBlFechaAll() ([]*model.BlFecha, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_fecha").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlFecha
	for cur.Next(*r.ctx) {
		var dt model.BlFecha
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlFleteAll() ([]*model.BlFlete, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_flete").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlFlete
	for cur.Next(*r.ctx) {
		var dt model.BlFlete
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlItemAll() ([]*model.BlItem, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_item").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlItem
	for cur.Next(*r.ctx) {
		var dt model.BlItem
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlItemImoAll() ([]*model.BlItemImo, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_item_imo").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlItemImo
	for cur.Next(*r.ctx) {
		var dt model.BlItemImo
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlItemContenedorAll() ([]*model.BlItemContenedor, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_item_contenedor").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlItemContenedor
	for cur.Next(*r.ctx) {
		var dt model.BlItemContenedor
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlItemContenedorImoAll() ([]*model.BlItemContenedorImo, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_contenedor_imo").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlItemContenedorImo
	for cur.Next(*r.ctx) {
		var dt model.BlItemContenedorImo
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlItemContenedorSelloAll() ([]*model.BlItemContenedorSello, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_contenedor_sello").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlItemContenedorSello
	for cur.Next(*r.ctx) {
		var dt model.BlItemContenedorSello
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlLocacionAll() ([]*model.BlLocacion, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_locacion").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlLocacion
	for cur.Next(*r.ctx) {
		var dt model.BlLocacion
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlObservacionAll() ([]*model.BlObservacion, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_observacion").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlObservacion
	for cur.Next(*r.ctx) {
		var dt model.BlObservacion
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlParticipanteAll() ([]*model.BlParticipante, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_participante").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlParticipante
	for cur.Next(*r.ctx) {
		var dt model.BlParticipante
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlReferenciaAll() ([]*model.BlReferencia, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_referencia").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlReferencia
	for cur.Next(*r.ctx) {
		var dt model.BlReferencia
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlTransbordoAll() ([]*model.BlTransbordo, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_transbordo").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlTransbordo
	for cur.Next(*r.ctx) {
		var dt model.BlTransbordo
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetBlTransporteAll() ([]*model.BlTransporte, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("bl_transporte").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.BlTransporte
	for cur.Next(*r.ctx) {
		var dt model.BlTransporte
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}

func (r *repo) BorraBlFecha(id string) error {
	fmt.Println("Borrando detalle de bl_fecha con id:", id)
	coll := r.db.Collection("bl_fecha")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_fecha con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlFlete(id string) error {
	fmt.Println("Borrando detalle de bl_flete con id:", id)
	coll := r.db.Collection("bl_flete")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_flete con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlItem(id string) error {
	fmt.Println("Borrando detalle de bl_item con id:", id)
	coll := r.db.Collection("bl_item")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_item con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlItemImo(id string) error {
	fmt.Println("Borrando detalle de bl_item_imo con id:", id)
	coll := r.db.Collection("bl_item_imo")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_item_imo con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlItemContenedor(id string) error {
	fmt.Println("Borrando detalle de bl_fecha con id:", id)
	coll := r.db.Collection("bl_item_contenedor")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_item_contenedor con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlItemContenedorImo(id string) error {
	fmt.Println("Borrando detalle de bl_fecha con id:", id)
	coll := r.db.Collection("bl_contenedor_imo")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_contenedor_imo con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlItemContenedorSello(id string) error {
	fmt.Println("Borrando detalle de bl_contenedor_sello con id:", id)
	coll := r.db.Collection("bl_contenedor_sello")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_contenedor_sello con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlLocacion(id string) error {
	fmt.Println("Borrando detalle de bl_locacion con id:", id)
	coll := r.db.Collection("bl_locacion")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_locacion con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlObservacion(id string) error {
	fmt.Println("Borrando detalle de bl_observacion con id:", id)
	coll := r.db.Collection("bl_observacion")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_observacion con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlParticipante(id string) error {
	fmt.Println("Borrando detalle de bl_fecha con id:", id)
	coll := r.db.Collection("bl_participante")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_participante con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlReferencia(id string) error {
	fmt.Println("Borrando detalle de bl_referencia con id:", id)
	coll := r.db.Collection("bl_referencia")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_referencia con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlTransbordo(id string) error {
	fmt.Println("Borrando detalle de bl_transbordo con id:", id)
	coll := r.db.Collection("bl_transbordo")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_transbordo con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) BorraBlTransporte(id string) error {
	fmt.Println("Borrando detalle de bl_transporte con id:", id)
	coll := r.db.Collection("bl_transporte")

	if oid, err := primitive.ObjectIDFromHex(id); err == nil {
		res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": oid})
		if err != nil {
			return fmt.Errorf("error al borrar (ObjectID): %w", err)
		}
		if res.DeletedCount > 0 {
			fmt.Println("Borrado OK (ObjectID), count:", res.DeletedCount)
			return nil
		}
	}

	res, err := coll.DeleteOne(*r.ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("error al borrar (string): %w", err)
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no se encontró el detalle de bl_transporte con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
