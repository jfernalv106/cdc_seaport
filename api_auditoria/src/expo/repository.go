package expo

import (
	"api_auditoria/src/model"
	"api_auditoria/src/utils"
	"context"
	"fmt"
	"log"
	"regexp"
	"sort"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	GuardarPapeletaExpo(papeletaExpo interface{}) (interface{}, error)
	ActualizaPapeletaExpo(id string, papeletaExpo interface{}) (interface{}, error)
	GuardarPapeletaExpoDetalle(papeletaExpoDetalle interface{}) (interface{}, error)
	GetUltimaPapeletaExpoPorEvento(id int64) (*model.PapeletaExpo, error)
	UpdatePapeletaExpo(papeletaExpo *model.PapeletaExpo) (*model.PapeletaExpo, error)
	GetPapeletaExpo(id *int64, booking *string, papeleta *string) ([]*model.PapeletaExpo, error)
	GetPapeletaExpoDetalle(idPapeleta int64) ([]*model.PapeletaExpoDetalle, error)
	GetPapeletaExpoDetalleAll() ([]*model.PapeletaExpoDetalle, error)
	BorrarPapeletaExpoDetalle(id string) error
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

/*** Create ***/

func (r *repo) GuardarPapeletaExpo(papeletaExpo interface{}) (interface{}, error) {
	res, err := r.db.Collection("papeleta_expo").InsertOne(*r.ctx, papeletaExpo)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) ActualizaPapeletaExpo(id string, papeletaExpo interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("papeleta_expo").ReplaceOne(*r.ctx, bson.M{"_id": objID}, papeletaExpo, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) GuardarPapeletaExpoDetalle(papeletaExpoDetalle interface{}) (interface{}, error) {
	res, err := r.db.Collection("papeleta_expo_detalle").InsertOne(*r.ctx, papeletaExpoDetalle)
	if err != nil {
		return nil, err
	}
	return res, nil
}

/*** Read (list + last-by-event) ***/

func (r *repo) GetPapeletaExpo(id *int64, booking *string, papeleta *string) ([]*model.PapeletaExpo, error) {
	filter := buildFilterExpo(id, booking, papeleta)

	cur, err := r.db.Collection("papeleta_expo").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var papeletas []*model.PapeletaExpo
	for cur.Next(*r.ctx) {
		var pr model.PapeletaExpo
		if err := cur.Decode(&pr); err != nil {
			return nil, err
		}
		papeletas = append(papeletas, &pr)
	}

	return papeletas, nil
}

func (r *repo) GetUltimaPapeletaExpoPorEvento(id int64) (*model.PapeletaExpo, error) {
	filter := bson.M{}

	filter["id"] = id

	cur, err := r.db.Collection("papeleta_expo").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.PapeletaExpo
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró Papeleta Expo con id %d", id)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}

func (r *repo) GetPapeletaExpoDetalle(idPapeleta int64) ([]*model.PapeletaExpoDetalle, error) {
	filter := bson.M{"id_papeleta": idPapeleta}

	cur, err := r.db.Collection("papeleta_expo_detalle").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.PapeletaExpoDetalle
	for cur.Next(*r.ctx) {
		var dt model.PapeletaExpoDetalle
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) GetPapeletaExpoDetalleAll() ([]*model.PapeletaExpoDetalle, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("papeleta_expo_detalle").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.PapeletaExpoDetalle
	for cur.Next(*r.ctx) {
		var dt model.PapeletaExpoDetalle
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) BorrarPapeletaExpoDetalle(id string) error {
	fmt.Println("Borrando detalle de papeleta expo con id:", id)

	coll := r.db.Collection("papeleta_expo_detalle")

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
		return fmt.Errorf("no se encontró el detalle de la papeleta expo con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}

func (r *repo) UpdatePapeletaExpo(papeletaExpo *model.PapeletaExpo) (*model.PapeletaExpo, error) {
	if papeletaExpo == nil {
		return nil, fmt.Errorf("la papeleta expo no puede ser nula")
	}
	if papeletaExpo.IDMongo.IsZero() {
		return nil, fmt.Errorf("se requiere el _id (IDMongo) de la papeleta expo para actualización")
	}

	// Solo actualizamos campos conocidos y no nulos para evitar sobreescrituras accidentales.
	set := bson.M{
		"detalles":         papeletaExpo.Detalles,
		"booking":          papeletaExpo.Booking,
		"estado":           papeletaExpo.Estado,
		"fecha_evento":     papeletaExpo.FechaEvento,
		"usuario_up":       papeletaExpo.UsuarioUp,
		"motivo_anulacion": papeletaExpo.MotivoAnulacion,
		// agrega aquí otros campos que quieras permitir actualizar
	}

	// Filtrar nils
	for k, v := range set {
		if v == nil {
			delete(set, k)
		}
	}

	update := bson.M{"$set": set}
	filter := bson.M{"_id": papeletaExpo.IDMongo}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.PapeletaExpo
	if err := r.db.Collection("papeleta_expo").
		FindOneAndUpdate(*r.ctx, filter, update, opts).
		Decode(&updated); err != nil {
		return nil, fmt.Errorf("error al actualizar papeleta expo: %v", err)
	}

	return &updated, nil
}

func buildFilterExpo(id *int64, booking *string, papeleta *string) bson.M {
	conditions := make([]bson.M, 0)

	if id != nil && *id != 0 {
		conditions = append(conditions, bson.M{"id": *id})
	}
	if papeleta != nil && *papeleta != "" {
		conditions = append(conditions, bson.M{"nro_papeleta": *papeleta})
	}

	if booking != nil && *booking != "" {
		regex := fmt.Sprintf(".*%s.*", regexp.QuoteMeta(*booking))
		conditions = append(conditions, bson.M{
			"booking": bson.M{"$regex": regex, "$options": "i"},
		})
	}

	if len(conditions) == 0 {
		return bson.M{}
	}
	if len(conditions) == 1 {
		return conditions[0]
	}
	return bson.M{"$or": conditions}
}
