package repository

import (
	"database/sql"
	"proyecto-bd-final/models"
	"time"
)

type PrestamoRepository struct {
	db *sql.DB
}

func NewPrestamoRepository(db *sql.DB) *PrestamoRepository {
	return &PrestamoRepository{db: db}
}

func (r *PrestamoRepository) GetPrestamosVencidos() ([]models.PrestamoVencido, error) {
	query := `
		SELECT p.idPrestamo, p.fechaPrestamo, p.fechaDevolucionPrevista,
		       u.nombre, u.apellido, u.correo,
		       l.titulo, e.codigo,
		       TRUNC(SYSDATE - p.fechaDevolucionPrevista) as diasVencido
		FROM Prestamo p
		JOIN Usuario u ON p.Usuario_idUsuario = u.idUsuario
		JOIN Ejemplar e ON e.Prestamo_idPrestamo = p.idPrestamo
		JOIN Libro l ON e.Libro_ISBN = l.ISBN
		WHERE p.estado = 'activo' 
		AND p.fechaDevolucionPrevista < SYSDATE
		AND p.fechaDevolucionReal IS NULL
		ORDER BY p.fechaDevolucionPrevista ASC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prestamos []models.PrestamoVencido
	for rows.Next() {
		var prestamo models.PrestamoVencido
		err := rows.Scan(
			&prestamo.IDPrestamo,
			&prestamo.FechaPrestamo,
			&prestamo.FechaDevolucionPrevista,
			&prestamo.NombreUsuario,
			&prestamo.ApellidoUsuario,
			&prestamo.CorreoUsuario,
			&prestamo.TituloLibro,
			&prestamo.CodigoEjemplar,
			&prestamo.DiasVencido,
		)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}

// GetEstadisticas obtiene las estadísticas generales del sistema
func (r *PrestamoRepository) GetEstadisticas() (*models.Estadisticas, error) {
	var stats models.Estadisticas

	// Total de libros
	err := r.db.QueryRow("SELECT COUNT(*) FROM Libro").Scan(&stats.TotalLibros)
	if err != nil {
		return nil, err
	}

	// Total de préstamos
	err = r.db.QueryRow("SELECT COUNT(*) FROM Prestamo").Scan(&stats.TotalPrestamos)
	if err != nil {
		return nil, err
	}

	// Préstamos activos
	err = r.db.QueryRow(`
		SELECT COUNT(*) FROM Prestamo 
		WHERE estado = 'activo' AND fechaDevolucionReal IS NULL
	`).Scan(&stats.PrestamosActivos)
	if err != nil {
		return nil, err
	}

	// Préstamos vencidos
	err = r.db.QueryRow(`
		SELECT COUNT(*) FROM Prestamo 
		WHERE estado = 'activo' 
		AND fechaDevolucionPrevista < SYSDATE
		AND fechaDevolucionReal IS NULL
	`).Scan(&stats.PrestamosVencidos)
	if err != nil {
		return nil, err
	}

	// Total de usuarios
	err = r.db.QueryRow("SELECT COUNT(*) FROM Usuario").Scan(&stats.TotalUsuarios)
	if err != nil {
		return nil, err
	}

	// Total de ejemplares
	err = r.db.QueryRow("SELECT COUNT(*) FROM Ejemplar").Scan(&stats.TotalEjemplares)
	if err != nil {
		return nil, err
	}

	// Ejemplares disponibles
	err = r.db.QueryRow(`
		SELECT COUNT(*) FROM Ejemplar 
		WHERE estado = 'disponible'
	`).Scan(&stats.EjemplaresDisponibles)
	if err != nil {
		return nil, err
	}

	return &stats, nil
}

func (r *PrestamoRepository) GetAllPrestamos() ([]models.PrestamoConDetalles, error) {
	query := `
		SELECT p.idPrestamo, p.fechaPrestamo, p.fechaDevolucionPrevista, 
		       p.fechaDevolucionReal, p.estado, p.Usuario_idUsuario,
		       u.nombre, u.apellido, l.titulo, e.codigo
		FROM Prestamo p
		JOIN Usuario u ON p.Usuario_idUsuario = u.idUsuario
		JOIN Ejemplar e ON e.Prestamo_idPrestamo = p.idPrestamo
		JOIN Libro l ON e.Libro_ISBN = l.ISBN
		ORDER BY p.fechaPrestamo DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prestamos []models.PrestamoConDetalles
	for rows.Next() {
		var prestamo models.PrestamoConDetalles
		var fechaDevReal sql.NullTime

		err := rows.Scan(
			&prestamo.IDPrestamo,
			&prestamo.FechaPrestamo,
			&prestamo.FechaDevolucionPrevista,
			&fechaDevReal,
			&prestamo.Estado,
			&prestamo.UsuarioID,
			&prestamo.NombreUsuario,
			&prestamo.ApellidoUsuario,
			&prestamo.TituloLibro,
			&prestamo.CodigoEjemplar,
		)
		if err != nil {
			return nil, err
		}

		if fechaDevReal.Valid {
			prestamo.FechaDevolucionReal = &fechaDevReal.Time
		}

		if prestamo.Estado == "activo" && time.Now().After(prestamo.FechaDevolucionPrevista) {
			prestamo.DiasVencido = int(time.Since(prestamo.FechaDevolucionPrevista).Hours() / 24)
		}

		prestamos = append(prestamos, prestamo)
	}

	return prestamos, nil
}
