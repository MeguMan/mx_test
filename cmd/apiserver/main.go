package main

import (
	"encoding/json"
	"github.com/MeguMan/mx_test/internal/app/apiserver"
	"log"
	"os"
)

func main() {
	dbConfig := apiserver.NewConfig()
	configFile, err := os.Open("configs/db.json")
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(configFile)
	if err := jsonParser.Decode(dbConfig); err != nil {
		log.Fatal(err)
	}

	if err := apiserver.Start(dbConfig); err != nil {
		log.Fatal(err)
	}
}
