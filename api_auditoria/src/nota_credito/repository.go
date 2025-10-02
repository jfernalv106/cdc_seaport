package notacredito

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
	GuardarNotaCredito(nc interface{}) (interface{}, error)
	GuardarNotaCreditoServicio(serv interface{}) (interface{}, error)
	ActualizaNotaCredito(id string, nota interface{}) (interface{}, error)
	GetUltimaNotaCreditoPorEvento(idNroNc int64) (*model.NotaCredito, error)
	UpdateNotaCredito(nc *model.NotaCredito) (*model.NotaCredito, error)
	GetNotaCredito(req *GetNotaCreditoRequest) (*model.NotaCredito, error)
	GetNotaCreditoServicios(idNroNc int64) ([]*model.NotaCreditoServicio, error)
	GetNotaCreditoServiciosAll() ([]*model.NotaCreditoServicio, error)
	BorrarNotaCreditoServicio(id string) error
}

type repo struct {
	log *log.Logger
	db  *mongo.Database
	ctx *context.Context
}

func NewRepo(log *log.Logger, db *mongo.Database, ctx *context.Context) Repository {
	return &repo{log: log, db: db, ctx: ctx}
}

func (r *repo) GuardarNotaCredito(nc interface{}) (interface{}, error) {
	res, err := r.db.Collection("nota_credito").InsertOne(*r.ctx, nc)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *repo) GuardarNotaCreditoServicio(serv interface{}) (interface{}, error) {
	res, err := r.db.Collection("nota_credito_servicio").InsertOne(*r.ctx, serv)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) ActualizaNotaCredito(id string, nota interface{}) (interface{}, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	opts := options.Replace().SetUpsert(false)
	res, err := r.db.Collection("nota_credito").ReplaceOne(*r.ctx, bson.M{"_id": objID}, nota, opts)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (r *repo) UpdateNotaCredito(nc *model.NotaCredito) (*model.NotaCredito, error) {
	if nc == nil {
		return nil, fmt.Errorf("la nota de crédito no puede ser nula")
	}

	update := bson.M{
		"$set": bson.M{
			"id_nro_nc":     nc.IDNroNc,
			"id_fol_nc":     nc.IDFolNc,
			"id_factura":    nc.IDFactura,
			"id_folio":      nc.IDFolio,
			"fecha_cr":      nc.FechaCR,
			"usuario_cr":    nc.UsuarioCR,
			"observaciones": nc.Observaciones,
			"valor_neto":    nc.ValorNeto,
			"tipo":          nc.Tipo,
			"empresa":       nc.Empresa,
			"traspasada":    nc.Traspasada,
			"id_sap":        nc.IDSap,
			"evento":        nc.Evento,
			"fecha_evento":  nc.FechaEvento,
			"servicios":     nc.Servicios,
		},
	}

	// Elimina campos nulos
	for k, v := range update["$set"].(bson.M) {
		if v == nil {
			delete(update["$set"].(bson.M), k)
		}
	}

	filter := bson.M{"_id": nc.IDMongo}
	opts := options.FindOneAndUpdate().SetReturnDocument(options.After)

	var updated model.NotaCredito
	if err := r.db.Collection("nota_credito").FindOneAndUpdate(*r.ctx, filter, update, opts).Decode(&updated); err != nil {
		return nil, fmt.Errorf("error al actualizar nota de crédito: %v", err)
	}
	return &updated, nil
}

func (r *repo) GetUltimaNotaCreditoPorEvento(idNroNc int64) (*model.NotaCredito, error) {
	filter := bson.M{}

	filter["id_nro_nc"] = idNroNc

	cur, err := r.db.Collection("nota_credito").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var result []*model.NotaCredito
	if err = cur.All(*r.ctx, &result); err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no se encontró La Nota de Credito con id %d", idNroNc)
	}
	sort.Slice(result, func(i, j int) bool {
		fechaI, _ := utils.ParseFecha(*result[i].FechaEvento)
		fechaJ, _ := utils.ParseFecha(*result[j].FechaEvento)
		return fechaI.After(fechaJ)
	})

	return result[0], nil
}
func (r *repo) GetNotaCreditoServiciosAll() ([]*model.NotaCreditoServicio, error) {
	filter := bson.M{}

	cur, err := r.db.Collection("nota_credito_servicio").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var detalles []*model.NotaCreditoServicio
	for cur.Next(*r.ctx) {
		var dt model.NotaCreditoServicio
		if err := cur.Decode(&dt); err != nil {
			return nil, err
		}
		detalles = append(detalles, &dt)
	}

	return detalles, nil
}
func (r *repo) BorrarNotaCreditoServicio(id string) error {
	fmt.Println("Borrando nota credito servicio con id:", id)

	coll := r.db.Collection("nota_credito_servicio")

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
		return fmt.Errorf("no se encontrónota credito servicio con id %s", id)
	}
	fmt.Println("Borrado OK (string), count:", res.DeletedCount)
	return nil
}
func (r *repo) GetNotaCredito(req *GetNotaCreditoRequest) (*model.NotaCredito, error) {
	filter := bson.M{}

	if req.IDNroNc != nil {
		filter["id_nro_nc"] = *req.IDNroNc
	}
	if req.IDFolNc != nil {
		filter["id_fol_nc"] = *req.IDFolNc
	}
	if req.IDFactura != nil {
		filter["id_factura"] = *req.IDFactura
	}
	if req.IDFolio != nil {
		filter["id_folio"] = *req.IDFolio
	}

	var nc model.NotaCredito
	if err := r.db.Collection("nota_credito").FindOne(*r.ctx, filter).Decode(&nc); err != nil {
		return nil, err
	}
	return &nc, nil
}

func (r *repo) GetNotaCreditoServicios(idNroNc int64) ([]*model.NotaCreditoServicio, error) {
	filter := bson.M{"id_nro_nc": idNroNc}

	cur, err := r.db.Collection("nota_credito_servicio").Find(*r.ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(*r.ctx)

	var servicios []*model.NotaCreditoServicio
	for cur.Next(*r.ctx) {
		var s model.NotaCreditoServicio
		if err := cur.Decode(&s); err != nil {
			return nil, err
		}
		servicios = append(servicios, &s)
	}
	return servicios, nil
}
