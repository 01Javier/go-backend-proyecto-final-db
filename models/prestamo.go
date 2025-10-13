package models

import "time"

// Prestamo representa la tabla Prestamo
type Prestamo struct {
	IDPrestamo              int        `json:"idPrestamo" db:"IDPRESTAMO"`
	FechaPrestamo           time.Time  `json:"fechaPrestamo" db:"FECHAPRESTAMO"`
	FechaDevolucionPrevista time.Time  `json:"fechaDevolucionPrevista" db:"FECHADEVOLUCIONPREVISTA"`
	FechaDevolucionReal     *time.Time `json:"fechaDevolucionReal,omitempty" db:"FECHADEVOLUCIONREAL"`
	Estado                  string     `json:"estado" db:"ESTADO"`
	UsuarioID               int        `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

// Ejemplar representa la tabla Ejemplar
type Ejemplar struct {
	Codigo     int    `json:"codigo" db:"CODIGO"`
	Estado     string `json:"estado" db:"ESTADO"`
	LibroISBN  int    `json:"libroIsbn" db:"LIBRO_ISBN"`
	PrestamoID int    `json:"prestamoId" db:"PRESTAMO_IDPRESTAMO"`
}

// PrestamoConDetalles es un modelo extendido con información completa del préstamo
type PrestamoConDetalles struct {
	Prestamo
	NombreUsuario   string `json:"nombreUsuario"`
	ApellidoUsuario string `json:"apellidoUsuario"`
	TituloLibro     string `json:"tituloLibro"`
	CodigoEjemplar  int    `json:"codigoEjemplar"`
	DiasVencido     int    `json:"diasVencido,omitempty"`
}

// PrestamoVencido representa un préstamo vencido
type PrestamoVencido struct {
	IDPrestamo              int       `json:"idPrestamo"`
	FechaPrestamo           time.Time `json:"fechaPrestamo"`
	FechaDevolucionPrevista time.Time `json:"fechaDevolucionPrevista"`
	NombreUsuario           string    `json:"nombreUsuario"`
	ApellidoUsuario         string    `json:"apellidoUsuario"`
	CorreoUsuario           string    `json:"correoUsuario"`
	TituloLibro             string    `json:"tituloLibro"`
	CodigoEjemplar          int       `json:"codigoEjemplar"`
	DiasVencido             int       `json:"diasVencido"`
}
