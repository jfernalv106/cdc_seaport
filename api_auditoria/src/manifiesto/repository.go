package manifiesto

import (
	"api_auditoria/src/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GuardarManifiesto(manifiesto interface{}) (interface{}, error)

	GetManifiesto(nro int64) (model.Manifiestos, error)
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

func (r *repo) GuardarManifiesto(manifiesto interface{}) (interface{}, error) {
	b, err := r.db.Collection("manifiesto").InsertOne(*r.ctx, manifiesto)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) GetManifiesto(nro int64) (model.Manifiestos, error) {
	var manifiestos model.Manifiestos
	filter := bson.M{"nro_mfto": nro}
	fmt.Println("Filter:", filter)
	cur, err := r.db.Collection("manifiesto").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var mn model.Manifiesto
		err = cur.Decode(&mn)

		if err != nil {
			fmt.Println("Error decoding manifiesto:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", mn)
		manifiestos = append(manifiestos, &mn)
	}
	defer cur.Close(*r.ctx)

	return manifiestos, nil

}
