package main

import (
	"log"
	"proyecto-bd-final/config"
	"proyecto-bd-final/database"
	"proyecto-bd-final/pkg/utils"
	"proyecto-bd-final/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.LoadConfig()

	utils.InitJWT(cfg.JWTSecret)

	if err := database.ConnectDatabase(cfg); err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer database.CloseDatabase()

	server := gin.Default()

	routes.SetupRoutes(server)

	log.Printf("Servidor corriendo en el puerto %s", cfg.ServerPort)
	if err := server.Run(":" + cfg.ServerPort); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
} ////
/////
