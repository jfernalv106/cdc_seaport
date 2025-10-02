package config

import (
	"encoding/json"
	"log"
	"os"
)

func Newconfig() *Config {
	file, err := os.Open("topic.json")
	if err != nil {
		log.Fatalf("No se pudo abrir topic.json: %v", err)
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatalf("Error al decodificar topic.json: %v", err)
	}
	return &config
}
func InitLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}
