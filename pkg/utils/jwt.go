package utils

import (
	"errors"
	"proyecto-bd-final/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret []byte

// InitJWT inicializa el secreto JWT
func InitJWT(secret string) {
	jwtSecret = []byte(secret)
}

// Claims representa las claims del JWT
type Claims struct {
	UsuarioID int      `json:"usuarioId"`
	Correo    string   `json:"correo"`
	Roles     []string `json:"roles"`
	jwt.RegisteredClaims
}

// GenerateToken genera un nuevo token JWT para el usuario
func GenerateToken(usuario models.UsuarioConRoles) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &Claims{
		UsuarioID: usuario.IDUsuario,
		Correo:    usuario.Correo,
		Roles:     usuario.Roles,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateToken valida y parsea un token JWT
func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token inválido")
	}

	return claims, nil
}
