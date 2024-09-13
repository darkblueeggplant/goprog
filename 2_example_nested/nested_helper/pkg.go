package nested_helper

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

type config struct {
	Name      string    `json:"name"`
	Height    string    `json:"height"`
	Mass      string    `json:"mass"`
	HairColor string    `json:"hair_color"`
	SkinColor string    `json:"skin_color"`
	EyeColor  string    `json:"eye_color"`
	BirthYear string    `json:"birth_year"`
	Gender    string    `json:"gender"`
	Homeworld string    `json:"homeworld"`
	Films     []string  `json:"films"`
	Species   []any     `json:"species"`
	Vehicles  []string  `json:"vehicles"`
	Starships []string  `json:"starships"`
	Created   time.Time `json:"created"`
	Edited    time.Time `json:"edited"`
	URL       string    `json:"url"`
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
