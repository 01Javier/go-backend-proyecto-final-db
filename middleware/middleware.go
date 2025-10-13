package middleware

import (
	"net/http"
	"proyecto-bd-final/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware verifica que el usuario esté autenticado
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			utils.ErrorResponse(c, "Token de autenticación requerido", http.StatusUnauthorized)
			c.Abort()
			return
		}

		// El token debe venir en formato "Bearer <token>"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.ErrorResponse(c, "Formato de token inválido", http.StatusUnauthorized)
			c.Abort()
			return
		}

		token := parts[1]
		claims, err := utils.ValidateToken(token)

		if err != nil {
			utils.ErrorResponseWithDetail(c, "Token inválido o expirado", http.StatusUnauthorized, err)
			c.Abort()
			return
		}

		// Guardar la información del usuario en el contexto
		c.Set("usuarioId", claims.UsuarioID)
		c.Set("correo", claims.Correo)
		c.Set("roles", claims.Roles)

		c.Next()
	}
}

// AdminMiddleware verifica que el usuario tenga rol de administrador
func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")

		if !exists {
			utils.ErrorResponse(c, "No se pudo verificar los roles del usuario", http.StatusForbidden)
			c.Abort()
			return
		}

		rolesSlice, ok := roles.([]string)
		if !ok {
			utils.ErrorResponse(c, "Error en la estructura de roles", http.StatusForbidden)
			c.Abort()
			return
		}

		// Verificar si el usuario tiene rol de admin
		isAdmin := false
		for _, rol := range rolesSlice {
			if rol == "admin" || rol == "administrador" {
				isAdmin = true
				break
			}
		}

		if !isAdmin {
			utils.ErrorResponse(c, "Acceso denegado: se requieren permisos de administrador", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}

// RoleMiddleware verifica que el usuario tenga uno de los roles especificados
func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		roles, exists := c.Get("roles")

		if !exists {
			utils.ErrorResponse(c, "No se pudo verificar los roles del usuario", http.StatusForbidden)
			c.Abort()
			return
		}

		rolesSlice, ok := roles.([]string)
		if !ok {
			utils.ErrorResponse(c, "Error en la estructura de roles", http.StatusForbidden)
			c.Abort()
			return
		}

		// Verificar si el usuario tiene alguno de los roles permitidos
		hasRole := false
		for _, userRole := range rolesSlice {
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					hasRole = true
					break
				}
			}
			if hasRole {
				break
			}
		}

		if !hasRole {
			utils.ErrorResponse(c, "Acceso denegado: no tienes permisos suficientes", http.StatusForbidden)
			c.Abort()
			return
		}

		c.Next()
	}
}

// CORSMiddleware agrega headers CORS a las respuestas
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
