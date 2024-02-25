package repository

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var dbConn *sql.DB

func getDBConn(ctx context.Context) *sql.DB {
	if dbConn == nil {
		info := fmt.Sprintf("host=%s password=%s user=%s sslmode=disable", "127.0.0.1", "12345", "postgres")
		open, err := sql.Open("postgres", info)
		if err != nil {
			log.Fatalln(err)
			return nil
		}
		dbConn = open
	}
	return dbConn
}
