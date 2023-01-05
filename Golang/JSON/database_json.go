package JSON

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type DatabaseConf struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Database string `json:"database"`
}

var DBConf = &DatabaseConf{}

func FillDatabaseConf() {
	file, err := ioutil.ReadFile("./Golang/JSON/database.json")
	if err != nil {
		log.Fatal("ERROR in reading file: ", err.Error())
	}

	err = json.Unmarshal(file, DBConf)
	if err != nil {
		log.Fatal("ERROR in unmarshalling: ", string(file))
	}
}
