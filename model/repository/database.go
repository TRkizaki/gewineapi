package repository

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		"todo-app", "todo-password", "sample-api-db:3306", "todo",
	)//DB/docker-compose.yml and value of Dockerfile 
	Db, err = sql.Open("mysql", dataSourceName)//Call sql.Open to initialize the db variable, passing the return value of dataSourceName.
	if err != nil {
		panic(err)
	}
}