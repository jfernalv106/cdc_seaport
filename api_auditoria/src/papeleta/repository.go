package papeleta

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
	GuardarPapeletaRecepcion(papeletaRecepcion interface{}) (interface{}, error)
	GuardarPapeletaRecepcionDetalle(papeletaRecepcionDetalle interface{}) (interface{}, error)
	ActualizaPapeletaRecepcion(id string, papeleta interface{}) (interface{}, error)
	GetUltimaPapeletaPorEvento(nroPapeleta string) (*model.PapeletaRecepcion, error)
	UpdatePapeletaRecepcion(papeletaRecepcion *model.PapeletaRecepcion) (*model.PapeletaRecepcion, error)
	GetPapeletaRecepcion(nroPapeletaEmpresa *string, manifiesto *string, bl *string) (model.PapeletasRecepcion, error)
	GetPapeletaRecepcionDetalle(nroPapeleta string) (model.PapeletaRecepcionDetalles, error)
	GetPapeletaDetalleAll() ([]model.PapeletaRecepcionDetalle, error)
	BorrarPapeletaDetalle(id string) error
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
func (r *repo) GuardarPapeletaRecepcion(papeletaRecepcion interface{}) (interface{}, error) {
	b, err := r.db.Collection("papeleta_recepcion").InsertOne(*r.ctx, papeletaRecepcion)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) ActualizaPapeletaRecepcion(id string, papeleta interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("papeleta_recepcion").ReplaceOne(*r.ctx, bson.M{"_id": objID}, papeleta, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) UpdatePapeletaRecepcion(papeletaRecepcion *model.PapeletaRecepcion) (*model.PapeletaRecepcion, error) {
	if papeletaRecepcion == nil {
		return nil, fmt.Errorf("la papeleta de recepción no puede ser nula")
	}

	if papeletaRecepcion.IDMongo.IsZero() {
		return nil, fmt.Errorf("se requiere el ID de la papeleta para actualización")
	}

	// Preparar el documento de actualización
	update := bson.M{
		"$set": bson.M{
			"detalles": papeletaRecepcion.Detalles,
			// Agrega aquí todos los campos que deseas actualizar
		},
	}

	// Filtrar campos nil
	for k, v := range update["$set"].(bson.M) {
		if v == nil {
			delete(update["$set"].(bson.M), k)
		}
	}

	filter := bson.M{"_id": papeletaRecepcion.IDMongo}
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var updatedDoc model.PapeletaRecepcion
	err := r.db.Collection("papeleta_recepcion").
		FindOneAndUpdate(*r.ctx, filter, update, opts).
		Decode(&updatedDoc)

	if err != nil {
		return nil, fmt.Errorf("error al actualizar papeleta: %v", err)
	}

	return &updatedDoc, nil
}
func (r *repo) GuardarPapeletaRecepcionDetalle(papeletaRecepcionDetalle interface{}) (interface{}, error) {
	b, err := r.db.Collection("papeleta_recepcion_detalle").InsertOne(*r.ctx, papeletaRecepcionDetalle)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GetPapeletaRecepcion(nroPapeletaEmpresa *string, manifiesto *string, bl *string) (model.PapeletasRecepcion, error) {
	var papeletas model.PapeletasRecepcion
	filter := buildFilter(nroPapeletaEmpresa, manifiesto, bl)

	cur, err := r.db.Collection("papeleta_recepcion").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var pr model.PapeletaRecepcion
		err = cur.Decode(&pr)

		if err != nil {

			return nil, err
		}

		papeletas = append(papeletas, &pr)
	}
	defer cur.Close(*r.ctx)

	return papeletas, nil

}
func (r *repo) GetUltimaPapeletaPorEvento(nroPapeleta string) (*model.PapeletaRecepcion, error) {
	filter := bson.M{}

	filter["nro_papeleta"] = nroPapeleta

	cur, err := r.db.Collection("papeleta_recepcion").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.PapeletaRecepcion
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró Papeleta con id %s", nroPapeleta)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}
func (r *repo) GetPapeletaRecepcionDetalle(nroPapeleta string) (model.PapeletaRecepcionDetalles, error) {
	var detalles model.PapeletaRecepcionDetalles
	filter := bson.M{"nro_papeleta": nroPapeleta}

	cur, err := r.db.Collection("papeleta_recepcion_detalle").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var pr model.PapeletaRecepcionDetalle
		err = cur.Decode(&pr)

		if err != nil {
			return nil, err
		}

		detalles = append(detalles, &pr)
	}
	defer cur.Close(*r.ctx)

	return detalles, nil

}

func (r *repo) GetPapeletaDetalleAll() ([]model.PapeletaRecepcionDetalle, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("papeleta_recepcion_detalle").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []model.PapeletaRecepcionDetalle
	for cur.Next(*r.ctx) {
		var dt model.PapeletaRecepcionDetalle
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, dt)
	}

	return detalles, nil
}
func (r *repo) BorrarPapeletaDetalle(id string) error {
	fmt.Println("Borrando detalle de papeleta con id:", id)

	coll := r.db.Collection("papeleta_recepcion_detalle")

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
		return fmt.Errorf("no se encontró el detalle de la papeleta recepcion con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}

func buildFilter(nroPapeletaEmpresa *string, manifiesto *string, bl *string) bson.M {
	var conditions []bson.M

	if nroPapeletaEmpresa != nil && *nroPapeletaEmpresa != "" {
		conditions = append(conditions, bson.M{"nro_papeleta_empresa": *nroPapeletaEmpresa})
	}

	if manifiesto != nil && *manifiesto != "" {
		conditions = append(conditions, bson.M{"manifiesto_nro_mfto_interno": *manifiesto})
	}

	if bl != nil && *bl != "" {
		regexPattern := fmt.Sprintf(".*%s.*", regexp.QuoteMeta(*bl))
		conditions = append(conditions, bson.M{"nro_bl": bson.M{"$regex": regexPattern, "$options": "i"}}) // "i" = case-insensitive
	}

	if len(conditions) == 0 {
		return bson.M{}
	}

	if len(conditions) == 1 {
		return conditions[0]
	}

	return bson.M{"$or": conditions}
}
