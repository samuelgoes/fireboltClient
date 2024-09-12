package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/firebolt-db/firebolt-go-sdk"
)

const (
	clientId     = "*******"       // Service Account ID
	clientSecret = "*******"       // Service Account Secret
	accountName  = "test-account"  // Account name (organization)
	databaseName = "test-database" // Database name
	engine       = "test-engine"   // Engine name
)

func main() {
	checkConnection()
	insert()
}

func checkConnection() {
	dsn := fmt.Sprintf("firebolt:///%s?account_name=%s&client_id=%s&client_secret=%s", databaseName, accountName, clientId, clientSecret)

	// Open the Firebolt connection
	db, err := sql.Open("firebolt", dsn)
	if err != nil {
		log.Fatal("Error opening connection:", err)
	}

	// Ping the database to check the connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to connect to Firebolt:", err)
	}

	fmt.Println("Connection successful!")
	defer db.Close()
}

func insert() {
	// Firebolt connection string
	dsn := fmt.Sprintf("firebolt:///%s?account_name=%s&client_id=%s&client_secret=%s&engine=%s", databaseName, accountName, clientId, clientSecret, engine)

	// Open the connection
	db, err := sql.Open("firebolt", dsn)
	if err != nil {
		log.Fatal("Failed to connect to Firebolt:", err)
	}
	defer db.Close()

	// Insert event data
	query := `
		INSERT INTO events (id, type, user_id, ts, metadata) 
		VALUES (?, ?, ?, ?, ?)
	`

	_, err = db.Exec(query, 123, "click", 456, "2024-09-02 12:34:56", "sample metadata")
	if err != nil {
		log.Fatal("Failed to insert event:", err)
	}

	fmt.Println("Event ingested successfully")
}
