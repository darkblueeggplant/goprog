package helper

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Age      int    `json:"age"`
	Info     string `json:"info"`
}

func ReadConfig(filename string) config {

	jsonBlob, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	cnf := config{}

	json.Unmarshal(jsonBlob, &cnf)

	return cnf

}
