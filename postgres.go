package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // using PostgreSQL impl
)

var postgresConnectionString string
var databaseTable string
var databaseColumns string

func readDatabaseData() ([]map[string]interface{}, error) {
	connStr := postgresConnectionString
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	whereNotEmpty := "1=1 "
	columns := strings.Split(databaseColumns, ",")
	for _, cl := range columns {
		whereNotEmpty += fmt.Sprintf("AND %s <> '' ", cl)
	}
	rows, err := db.Queryx(fmt.Sprintf("SELECT %s FROM %s WHERE %s ORDER BY %s ASC", databaseColumns, databaseTable, whereNotEmpty, databaseColumns))

	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	resultSet := make([]map[string]interface{}, 0)
	for rows.Next() {
		resultRow := make(map[string]interface{}, 0)
		err := rows.MapScan(resultRow)
		if err != nil {
			log.Fatalln(err)
		}
		resultSet = append(resultSet, resultRow)
	}

	return resultSet, nil
}
