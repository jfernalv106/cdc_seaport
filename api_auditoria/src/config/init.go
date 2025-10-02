package config

import (
	"encoding/json"
	"log"
	"os"
)

func NewConfig() Mongo {
	file, err := os.Open("db.json")
	if err != nil {
		log.Fatalf("No se pudo abrir db.json: %v", err)
	}
	defer file.Close()

	var db Mongo
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&db); err != nil {
		log.Fatalf("No se pudo decodificar db.json: %v", err)
	}
	return db
}
