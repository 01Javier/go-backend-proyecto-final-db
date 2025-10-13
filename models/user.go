package models

import "time"

// Usuario representa la tabla Usuario en la base de datos
type Usuario struct {
	IDUsuario     int       `json:"idUsuario" db:"IDUSUARIO"`
	Nombre        string    `json:"nombre" db:"NOMBRE"`
	Apellido      string    `json:"apellido" db:"APELLIDO"`
	Contrasenia   string    `json:"contrasenia,omitempty" db:"CONTRASENIA"`
	Correo        string    `json:"correo" db:"CORREO"`
	Telefono      int       `json:"telefono" db:"TELEFONO"`
	FechaRegistro time.Time `json:"fechaRegistro" db:"FECHAREGISTRO"`
}

// Estudiante representa la tabla Estudiante
type Estudiante struct {
	Carnet    int    `json:"carnet" db:"CARNET"`
	Carrera   string `json:"carrera" db:"CARRERA"`
	Semestre  int    `json:"semestre" db:"SEMESTRE"`
	UsuarioID int    `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

// Profesor representa la tabla Profesor
type Profesor struct {
	CodigoDocencia int    `json:"codigoDocencia" db:"CODIGODOCENCIA"`
	Facultad       string `json:"facultad" db:"FACULTAD"`
	UsuarioID      int    `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

// Personal representa la tabla Personal
type Personal struct {
	CodigoEmpleado int    `json:"codigoEmpleado" db:"CODIGOEMPLEADO"`
	Puesto         string `json:"puesto" db:"PUESTO"`
	UsuarioID      int    `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

// Rol representa la tabla Roles
type Rol struct {
	IDRol     int    `json:"idRol" db:"IDROL"`
	NombreRol string `json:"nombreRol" db:"NOMBREROL"`
}

// UsuarioRol representa la relación entre Usuario y Rol
type UsuarioRol struct {
	IDUsuarioRol int `json:"idUsuarioRol" db:"IDUSUARIOROL"`
	UsuarioID    int `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
	RolID        int `json:"rolId" db:"ROLES_IDROL"`
}

// UsuarioConRoles es un modelo extendido que incluye los roles del usuario
type UsuarioConRoles struct {
	Usuario
	Roles []string `json:"roles"`
}

// LoginRequest representa la petición de login
type LoginRequest struct {
	Correo      string `json:"correo" binding:"required,email"`
	Contrasenia string `json:"contrasenia" binding:"required"`
}

// RegisterRequest representa la petición de registro
type RegisterRequest struct {
	Nombre      string `json:"nombre" binding:"required"`
	Apellido    string `json:"apellido" binding:"required"`
	Correo      string `json:"correo" binding:"required,email"`
	Contrasenia string `json:"contrasenia" binding:"required,min=6"`
	Telefono    int    `json:"telefono"`
	RolID       int    `json:"rolId" binding:"required"` // 1=estudiante, 2=profesor, 3=admin, etc.

	// Campos específicos según el rol
	Carnet   *int    `json:"carnet,omitempty"`   // Para estudiante
	Carrera  *string `json:"carrera,omitempty"`  // Para estudiante
	Semestre *int    `json:"semestre,omitempty"` // Para estudiante
	Facultad *string `json:"facultad,omitempty"` // Para profesor
	Puesto   *string `json:"puesto,omitempty"`   // Para personal
}

// AuthResponse representa la respuesta de autenticación
type AuthResponse struct {
	Token   string          `json:"token"`
	Usuario UsuarioConRoles `json:"usuario"`
}
