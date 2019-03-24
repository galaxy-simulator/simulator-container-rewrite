package main

import (
	"database/sql"
	"fmt"
	"log"
)

func connectToDatabase() {
	log.Println("[Database] Initializing the database connection")
	db = connectToDB()
	db.SetMaxOpenConns(75)
	pingDB()
}

// connectToDB returns a pointer to an sql database writing to the database
func connectToDB() *sql.DB {
	//connStr := fmt.Sprintf("user=%s dbname=%s sslmode=%s", DBUSER, DBNAME, DBSSLMODE)
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		DBHOST, DBPORT, DBUSER, DBPASSWD, DBNAME)
	db := dbConnect(connStr)
	return db
}

// dbConnect connects to a PostgreSQL database
func dbConnect(connStr string) *sql.DB {
	// connect to the database
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("[ E ] connection: %v", err)
	}

	return db
}

func pingDB() {
	// ping the db
	err := db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("[ ] Done Pinging the DB")
}
