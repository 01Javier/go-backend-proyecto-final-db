package models

import "time"

type Bitacora struct {
	IDBitacora int       `json:"idBitacora" db:"IDBITACORA"`
	Accion     string    `json:"accion" db:"ACCION"`
	FechaHora  time.Time `json:"fechaHora" db:"FECHAHORA"`
	Detalle    string    `json:"detalle" db:"DETALLE"`
	Entidad    string    `json:"entidad" db:"ENTIDAD"`
	UsuarioID  int       `json:"usuarioId" db:"USUARIO_IDUSUARIO"`
}

type BitacoraConUsuario struct {
	Bitacora
	NombreUsuario   string `json:"nombreUsuario"`
	ApellidoUsuario string `json:"apellidoUsuario"`
	CorreoUsuario   string `json:"correoUsuario"`
}

type RegistroBitacoraRequest struct {
	Accion  string `json:"accion" binding:"required"`
	Detalle string `json:"detalle"`
	Entidad string `json:"entidad" binding:"required"`
}
