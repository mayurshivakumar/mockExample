package models

import (
	"context"
	"database/sql"
	"fmt"
	"log"
)

var server = "localhost"
var port = 1433
var user = "sa"
var password = "yourStrong(!)Password"
var database = "SampleDB"

//Datastore ...
type Datastore interface {
	GetAllEmployees() ([]*Employee, error)
}

//DB ...
type DB struct {
	*sql.DB
}

//NewDB ...
func NewDB(dataSourceName string) (*DB, error) {
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;",
		server, user, password, port, database)

	var err error

	// Create connection pool
	db, err := sql.Open(dataSourceName, connString)
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	ctx := context.Background()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatal(err.Error())
	}

	return &DB{db}, nil
}
