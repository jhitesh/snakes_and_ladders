package JSON

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Players struct {
	Players []string `json:"players"`
}

func GetPlayers() []string {
	players_ := &Players{}

	file, err := ioutil.ReadFile("./JSON/players.json")
	if err != nil {
		log.Fatal("ERROR in reading file: ", err.Error())
	}

	err = json.Unmarshal(file, players_)
	if err != nil {
		log.Fatal("ERROR in unmarshalling data: ", string(file), err.Error())
	}
	return players_.Players
}
