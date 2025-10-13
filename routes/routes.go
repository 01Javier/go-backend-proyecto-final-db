package routes

import (
	"proyecto-bd-final/handlers"
	"proyecto-bd-final/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	// Aplicar middleware CORS
	router.Use(middleware.CORSMiddleware())

	// Inicializar handlers
	authHandler := handlers.NewAuthHandler()
	adminHandler := handlers.NewAdminHandler()
	estadisticasHandler := handlers.NewEstadisticasHandler()
	bitacoraHandler := handlers.NewBitacoraHandler()

	// Rutas públicas de autenticación
	auth := router.Group("/api/auth")
	{
		auth.POST("/register", authHandler.Register)
		auth.POST("/login", authHandler.Login)
	}

	// Rutas protegidas (requieren autenticación)
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		// Perfil del usuario
		protected.GET("/profile", authHandler.GetProfile)

		// Estadísticas (disponible para usuarios autenticados)
		protected.GET("/estadisticas", estadisticasHandler.GetEstadisticas)
		protected.GET("/prestamos", estadisticasHandler.GetAllPrestamos)
		protected.GET("/prestamos/vencidos", estadisticasHandler.GetPrestamosVencidos)

		// Bitácora del usuario actual
		protected.POST("/bitacora", bitacoraHandler.CreateRegistro)
	}

	// Rutas de administrador (requieren autenticación + rol admin)
	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware())
	admin.Use(middleware.AdminMiddleware())
	{
		// Gestión de usuarios
		admin.GET("/usuarios", adminHandler.GetAllUsuarios)
		admin.GET("/usuarios/:id", adminHandler.GetUsuarioByID)
		admin.PUT("/usuarios/:id", adminHandler.UpdateUsuario)
		admin.DELETE("/usuarios/:id", adminHandler.DeleteUsuario)

		// Bitácora completa (solo admin)
		admin.GET("/bitacora", bitacoraHandler.GetAllRegistros)
		admin.GET("/bitacora/usuario/:id", bitacoraHandler.GetRegistrosByUsuario)
		admin.GET("/bitacora/entidad/:entidad", bitacoraHandler.GetRegistrosByEntidad)
	}

	// Ruta de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"message": "Biblioteca API is running",
		})
	})
}
