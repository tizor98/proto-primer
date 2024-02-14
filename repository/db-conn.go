package repository

import (
	"context"
	"database/sql"
	"log"
)

var dbConn *sql.DB

func getDBConn(ctx context.Context) *sql.DB {
	if dbConn == nil {
		open, err := sql.Open("postgres", "host")
		if err != nil {
			log.Fatalln(err)
			return nil
		}
		dbConn = open
	}
	return dbConn
}
