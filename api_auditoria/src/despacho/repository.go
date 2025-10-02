package despacho

import (
	"api_auditoria/src/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GuardarDespacho(despacho interface{}) (interface{}, error)

	GetDespacho(id *int64, visacion *int64, expo *int64, guia *int64) (model.Despachos, error)
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

func (r *repo) GuardarDespacho(despacho interface{}) (interface{}, error) {
	b, err := r.db.Collection("despacho").InsertOne(*r.ctx, despacho)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) GetDespacho(id *int64, visacion *int64, expo *int64, guia *int64) (model.Despachos, error) {
	var despachos model.Despachos
	filter := buildFilter(id, visacion, expo, guia)
	cur, err := r.db.Collection("despacho").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var pr model.Despacho
		err = cur.Decode(&pr)

		if err != nil {
			fmt.Println("Error decoding despacho:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", pr)
		despachos = append(despachos, &pr)
	}
	defer cur.Close(*r.ctx)

	return despachos, nil

}
func buildFilter(id *int64, visacion *int64, expo *int64, guia *int64) bson.M {
	var conditions []bson.M

	if id != nil {
		conditions = append(conditions, bson.M{"id": id})
	}

	if visacion != nil {
		conditions = append(conditions, bson.M{"visacion": visacion})
	}
	if expo != nil {
		conditions = append(conditions, bson.M{"papeletaexpo": expo})
	}
	if guia != nil {
		conditions = append(conditions, bson.M{"guiadespacho": guia})
	}

	if len(conditions) == 0 {
		return bson.M{}
	}

	if len(conditions) == 1 {
		return conditions[0]
	}

	return bson.M{"$or": conditions}
}
