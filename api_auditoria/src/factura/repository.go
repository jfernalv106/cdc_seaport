package factura

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
	GuardarFactura(factura interface{}) (interface{}, error)
	GuardarFacturaDetalle(factura interface{}) (interface{}, error)
	ActualizaFactura(id string, factura interface{}) (interface{}, error)
	GetFactura(folio *int64, manifiesto *string) (model.Facturas, error)
	GetFacturaDetalle(factura *int64) (model.FacturaDetalles, error)
	GetUltimaFacturaPorEvento(id int64) (*model.Factura, error)
	UpdateFactura(factura *model.Factura) (*model.Factura, error)
	BorrarFacturaDetalle(id string) error
	GetFacturaDetalleAll() ([]*model.FacturaDetalle, error)
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

func (r *repo) GuardarFactura(factura interface{}) (interface{}, error) {
	b, err := r.db.Collection("factura").InsertOne(*r.ctx, factura)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) ActualizaFactura(id string, factura interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("factura").ReplaceOne(*r.ctx, bson.M{"_id": objID}, factura, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) GuardarFacturaDetalle(detalle interface{}) (interface{}, error) {
	b, err := r.db.Collection("factura_detalle").InsertOne(*r.ctx, detalle)
	if err != nil {
		return nil, err
	}
	return b, nil
}
func (r *repo) GetFactura(folio *int64, manifiesto *string) (model.Facturas, error) {
	var facturas model.Facturas
	filter := buildFilter(folio, manifiesto)
	fmt.Println("Filter:", filter)
	cur, err := r.db.Collection("factura").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var fc model.Factura
		err = cur.Decode(&fc)

		if err != nil {
			fmt.Println("Error decoding factura:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", fc)
		facturas = append(facturas, &fc)
	}
	defer cur.Close(*r.ctx)

	return facturas, nil

}
func (r *repo) GetFacturaDetalle(factura *int64) (model.FacturaDetalles, error) {
	var detalles model.FacturaDetalles
	filter := bson.M{"idfactura": factura}
	fmt.Println("Filter:", filter)
	cur, err := r.db.Collection("factura_detalle").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var fc model.FacturaDetalle
		err = cur.Decode(&fc)

		if err != nil {
			fmt.Println("Error decoding factura:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", fc)
		detalles = append(detalles, &fc)
	}
	defer cur.Close(*r.ctx)

	return detalles, nil

}
func (r *repo) GetFacturaDetalleAll() ([]*model.FacturaDetalle, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("factura_detalle").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.FacturaDetalle
	for cur.Next(*r.ctx) {
		var dt model.FacturaDetalle
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) BorrarFacturaDetalle(id string) error {
	fmt.Println("Borrando detalle de factura con id:", id)

	coll := r.db.Collection("factura_detalle")

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
		return fmt.Errorf("no se encontró el detalle de la factura con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) GetUltimaFacturaPorEvento(id int64) (*model.Factura, error) {

	filter := bson.M{}

	filter["id"] = id

	cur, err := r.db.Collection("factura").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.Factura
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró BL con id %d", id)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}

func (r *repo) UpdateFactura(factura *model.Factura) (*model.Factura, error) {
	if factura == nil {
		return nil, fmt.Errorf("la factura no puede ser nula")
	}

	if factura.IDMongo.IsZero() {
		return nil, fmt.Errorf("se requiere el ID de la factura para actualización")
	}

	// Preparar el documento de actualización
	update := bson.M{
		"$set": bson.M{
			"detalles": factura.Detalles,
			// Agrega aquí todos los campos que deseas actualizar
		},
	}

	// Filtrar campos nil
	for k, v := range update["$set"].(bson.M) {
		if v == nil {
			delete(update["$set"].(bson.M), k)
		}
	}

	filter := bson.M{"_id": factura.IDMongo}
	opts := options.FindOneAndUpdate().
		SetReturnDocument(options.After)

	var facturaDoc model.Factura
	err := r.db.Collection("factura").
		FindOneAndUpdate(*r.ctx, filter, update, opts).
		Decode(&facturaDoc)

	if err != nil {
		return nil, fmt.Errorf("error al actualizar factura: %v", err)
	}

	return &facturaDoc, nil
}
func buildFilter(folio *int64, manifiesto *string) bson.M {
	var conditions []bson.M

	if folio != nil {
		conditions = append(conditions, bson.M{"folio": folio})
	}

	if manifiesto != nil {
		conditions = append(conditions, bson.M{"manifiesto": manifiesto})
	}

	if len(conditions) == 0 {
		return bson.M{}
	}

	if len(conditions) == 1 {
		return conditions[0]
	}

	return bson.M{"$or": conditions}
}
