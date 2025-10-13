package models

type Estadisticas struct {
	TotalLibros           int `json:"totalLibros"`
	TotalPrestamos        int `json:"totalPrestamos"`
	PrestamosActivos      int `json:"prestamosActivos"`
	PrestamosVencidos     int `json:"prestamosVencidos"`
	TotalUsuarios         int `json:"totalUsuarios"`
	TotalEjemplares       int `json:"totalEjemplares"`
	EjemplaresDisponibles int `json:"ejemplaresDisponibles"`
}
