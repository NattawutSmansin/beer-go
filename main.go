package main

import (
	"beer/database"
	"beer/server"
	"log"
)

// @title Beer API Title
// @version 1.0
// @description This is a sample Beer API
// @host localhost:8000
// @BasePath /api/v1
func main() {

	// Initialize MySQL Database
	db := database.ConMySQLDatabase()
	if db == nil {
		log.Fatal("MySQL connection failed")
	} else {
		// database.Migrate()
		log.Println("MySQL connected and migration completed.")
	}

	// Initialize MongoDB Database
	dbMongo := database.ConMongoDatabase()
	if dbMongo.GetDb() == nil {
		log.Fatal("MongoDB connection failed")
	} else {
		log.Println("MongoDB connected.")
	}

	// Start the server with database instances
	server := server.NewGinServer(db, dbMongo)
	if server == nil {
		log.Fatal("Failed to initialize server")
	} else {
		server.Start()
		log.Println("Server started Document API on http://localhost:8000/api/v1/docs/index.html")
	}
}
