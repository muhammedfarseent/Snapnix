package main

import (
	"SNAPNIX/config"
	"SNAPNIX/database"
	"SNAPNIX/migration"
	"SNAPNIX/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading the config file")
	}

	db, err := database.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("error to connecting database")
	}

	migration.Migrate(db)

	routes.RegisterRoutes(r, db)
	log.Printf("Starting server at %s\n", cfg.ServerAdr)
	if err := r.Run(cfg.ServerAdr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}

}
