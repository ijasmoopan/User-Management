package repository

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase() *sql.DB {
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "ijasmoopan"
	dbname := "users"
	
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatalln("Can't open database", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatalln("Can't connect to database", err)
	}
	// fmt.Println("Database connected succesfully")
	return db
}

func CloseDatabase(db *sql.DB) {
	db.Close()
}