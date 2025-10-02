package visacion

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"context"
	"fmt"
	"log"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GuardarVisacion(visacion interface{}) (interface{}, error)
	GuardarVisacionMercancias(mercancias interface{}) (interface{}, error)
	ActualizaVisacion(id string, visacion interface{}) (interface{}, error)
	GetVisacion(nroPapeleta string) (model.Visaciones, error)
	GetVisacionMercancias(visaje int64) (model.MercanciasDespachadas, error)
	GetUltimaVisacionEvento(id int64) (*model.Visacion, error)
	UpdateVisaje(visacion *model.Visacion) (*model.Visacion, error)
	GetMercanciasAll() ([]model.MercanciasDespachada, error)
	BorrarMercancia(id string) error
}
type repo struct {
	log *log.Logger
	db  *mongo.Database
	ctx *context.Context
}

func NewRepo(log *log.Logger, db *mongo.Database, ctx *context.Context) Repository {
	return &repo{
		log: log,
		db:  db,
		ctx: ctx,
	}
}

func (r *repo) GuardarVisacion(visacion interface{}) (interface{}, error) {
	b, err := r.db.Collection("visacion").InsertOne(*r.ctx, visacion)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GuardarVisacionMercancias(mercancias interface{}) (interface{}, error) {
	b, err := r.db.Collection("visacion-mercancias").InsertOne(*r.ctx, mercancias)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) ActualizaVisacion(id string, visacion interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("visacion").ReplaceOne(*r.ctx, bson.M{"_id": objID}, visacion, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) GetVisacion(nroPapeleta string) (model.Visaciones, error) {
	var visaciones model.Visaciones
	filter := bson.M{"nro_papeleta": nroPapeleta}
	fmt.Println("Filter:", nroPapeleta)
	cur, err := r.db.Collection("visacion").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var pr model.Visacion
		err = cur.Decode(&pr)

		if err != nil {
			fmt.Println("Error decoding visacion:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", pr)
		visaciones = append(visaciones, &pr)
	}
	defer cur.Close(*r.ctx)

	return visaciones, nil

}

func (r *repo) GetUltimaVisacionEvento(id int64) (*model.Visacion, error) {

	filter := bson.M{}

	filter["id"] = id

	cur, err := r.db.Collection("visacion").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.Visacion
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró visacion con id %s", id)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}

func (r *repo) GetVisacionMercancias(visaje int64) (model.MercanciasDespachadas, error) {
	var mercancias model.MercanciasDespachadas
	filter := bson.M{"nro_papeleta_empresa": visaje}
	fmt.Println("Filter:", filter)
	cur, err := r.db.Collection("visacion-mercancias").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var mr model.MercanciasDespachada
		err = cur.Decode(&mr)

		if err != nil {
			fmt.Println("Error decoding visacion:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", mr)
		mercancias = append(mercancias, &mr)
	}
	defer cur.Close(*r.ctx)

	return mercancias, nil

}
func (r *repo) UpdateVisaje(visacion *model.Visacion) (*model.Visacion, error) {
	if visacion == nil {
		return nil, fmt.Errorf("la visacion no puede ser nula")
	}

	if visacion.IDMongo.IsZero() {
		return nil, fmt.Errorf("se requiere el ID de la papeleta para actualización")
	}

	// Preparar el documento de actualización
	update := bson.M{
		"$set": bson.M{
			"mercacias": visacion.Mercacias,
			// Agrega aquí todos los campos que deseas actualizar
		},
	}

	// Filtrar campos nil
	for k, v := range update["$set"].(bson.M) {
		if v == nil {
			delete(update["$set"].(bson.M), k)
		}
	}

	filter := bson.M{"_id": visacion.IDMongo}
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var visacionDoc model.Visacion
	err := r.db.Collection("visacion").
		FindOneAndUpdate(*r.ctx, filter, update, opts).
		Decode(&visacionDoc)

	if err != nil {
		return nil, fmt.Errorf("error al actualizar papeleta: %v", err)
	}

	return &visacionDoc, nil
}
func (r *repo) GetMercanciasAll() ([]model.MercanciasDespachada, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("visacion-mercancias").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var mercancias []model.MercanciasDespachada
	for cur.Next(*r.ctx) {
		var dt model.MercanciasDespachada
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		mercancias = append(mercancias, dt)
	}

	return mercancias, nil
}
func (r *repo) BorrarMercancia(id string) error {
	fmt.Println("Borrando mercancias  con id:", id)

	coll := r.db.Collection("visacion-mercancias")

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
		return fmt.Errorf("no se encontró mercancia con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
