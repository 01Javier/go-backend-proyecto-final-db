package models

import "time"

type Prestamo struct {
	IDPrestamo              int        `json:"idPrestamo" db:"IDPRESTAMO"`
	FechaPrestamo           time.Time  `json:"fechaPrestamo" db:"FECHAPRESTAMO"`
	FechaDevolucionPrevista time.Time  `json:"fechaDevolucionPrevista" db:"FECHADEVOLUCIONPREVISTA"`
	FechaDevolucionReal     *time.Time `json:"fechaDevolucionReal,omitempty" db:"FECHADEVOLUCIONREAL"`
	Estado                  string     `json:"estado" db:"ESTADO"`
	UsuarioID               int        `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

type Ejemplar struct {
	Codigo     int    `json:"codigo" db:"CODIGO"`
	Estado     string `json:"estado" db:"ESTADO"`
	LibroISBN  int    `json:"libroIsbn" db:"LIBRO_ISBN"`
	PrestamoID int    `json:"prestamoId" db:"PRESTAMO_IDPRESTAMO"`
}

type PrestamoConDetalles struct {
	Prestamo
	NombreUsuario   string `json:"nombreUsuario"`
	ApellidoUsuario string `json:"apellidoUsuario"`
	TituloLibro     string `json:"tituloLibro"`
	CodigoEjemplar  int    `json:"codigoEjemplar"`
	DiasVencido     int    `json:"diasVencido,omitempty"`
}

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
