package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq" // blank import to get the postgres package
)

func GetDBConn() *sql.DB {
	postgresSource := "host=db user=postgres password=postgres dbname=petstore sslmode=disable"
	for i := 0; i < 10; i++ {
		db, err := sql.Open("postgres", postgresSource)
		if err == nil {
			return db
		}
		fmt.Println(err)
		fmt.Println("cant connect, retrying!")
		time.Sleep(1 * time.Second)
	}
	return nil
}
