package MySQL

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"snake_and_ladder/JSON"
)

var DBConn = &sql.DB{}

func CreateConnection() {
	var err error
	dbConf := JSON.DBConf
	dataSourceName := fmt.Sprintf("%v:%v@%v/%v", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Database)
	DBConn, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Fatal("ERROR in making connection to database", err.Error())
	}
	fmt.Println("Successfully connected to MySQL Database")
}
