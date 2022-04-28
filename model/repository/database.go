package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB //This is your database handle.

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"winelist-app", "winelist-password", "gewineapi-db:3306", "winelist",
	)//DB/docker-compose.yml and value of Dockerfile 
	Db, err = sql.Open("mysql", dataSourceName)//Call sql.Open to initialize the db variable, passing the return value of dataSourceName.
	if err != nil {
		panic(err)
	}
}