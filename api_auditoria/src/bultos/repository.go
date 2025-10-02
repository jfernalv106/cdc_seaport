package bultos

import (
	"api_auditoria/src/model"
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	GuardarBulto(bulto interface{}) (interface{}, error)
	GetBultos(cod string) (model.Bultos, error)
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
func (r *repo) GuardarBulto(bulto interface{}) (interface{}, error) {
	b, err := r.db.Collection("bulto").InsertOne(*r.ctx, bulto)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (r *repo) GetBultos(cod string) (model.Bultos, error) {
	var bultos model.Bultos
	filter := bson.M{"cod": cod}
	fmt.Println("Filter:", filter)
	cur, err := r.db.Collection("bulto").Find(*r.ctx, filter)

	if err != nil {
		return nil, err
	}

	for cur.Next(*r.ctx) {

		var bulto model.Bulto
		err = cur.Decode(&bulto)

		if err != nil {
			fmt.Println("Error decoding bulto:", err)
			return nil, err
		}
		fmt.Printf("Filter:%+v\n", bulto)
		bultos = append(bultos, &bulto)
	}
	defer cur.Close(*r.ctx)

	return bultos, nil

}
