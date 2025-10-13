package repository

import (
	"database/sql"
	"proyecto-bd-final/models"
	"time"
)

type BitacoraRepository struct {
	db *sql.DB
}

func NewBitacoraRepository(db *sql.DB) *BitacoraRepository {
	return &BitacoraRepository{db: db}
}

func (r *BitacoraRepository) CreateRegistro(bitacora *models.Bitacora) error {
	query := `
		INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario)
		VALUES (bitacora_seq.NEXTVAL, :1, :2, :3, :4, :5)
		RETURNING idBitacora INTO :6
	`

	var id int
	err := r.db.QueryRow(query,
		bitacora.Accion,
		time.Now(),
		bitacora.Detalle,
		bitacora.Entidad,
		bitacora.UsuarioID,
	).Scan(&id)

	if err != nil {
		return err
	}

	bitacora.IDBitacora = id
	bitacora.FechaHora = time.Now()
	return nil
}

func (r *BitacoraRepository) GetAllRegistros(limit int) ([]models.BitacoraConUsuario, error) {
	query := `
		SELECT b.idBitacora, b.accion, b.fechaHora, b.detalle, b.entidad, b.Usuario_idUsuario,
		       u.nombre, u.apellido, u.correo
		FROM Bitacora b
		JOIN Usuario u ON b.Usuario_idUsuario = u.idUsuario
		ORDER BY b.fechaHora DESC
		FETCH FIRST :1 ROWS ONLY
	`

	rows, err := r.db.Query(query, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registros []models.BitacoraConUsuario
	for rows.Next() {
		var registro models.BitacoraConUsuario
		err := rows.Scan(
			&registro.IDBitacora,
			&registro.Accion,
			&registro.FechaHora,
			&registro.Detalle,
			&registro.Entidad,
			&registro.UsuarioID,
			&registro.NombreUsuario,
			&registro.ApellidoUsuario,
			&registro.CorreoUsuario,
		)
		if err != nil {
			return nil, err
		}
		registros = append(registros, registro)
	}

	return registros, nil
}

func (r *BitacoraRepository) GetRegistrosByUsuario(usuarioID int) ([]models.Bitacora, error) {
	query := `
		SELECT idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario
		FROM Bitacora
		WHERE Usuario_idUsuario = :1
		ORDER BY fechaHora DESC
	`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registros []models.Bitacora
	for rows.Next() {
		var registro models.Bitacora
		err := rows.Scan(
			&registro.IDBitacora,
			&registro.Accion,
			&registro.FechaHora,
			&registro.Detalle,
			&registro.Entidad,
			&registro.UsuarioID,
		)
		if err != nil {
			return nil, err
		}
		registros = append(registros, registro)
	}

	return registros, nil
}

func (r *BitacoraRepository) GetRegistrosByEntidad(entidad string) ([]models.BitacoraConUsuario, error) {
	query := `
		SELECT b.idBitacora, b.accion, b.fechaHora, b.detalle, b.entidad, b.Usuario_idUsuario,
		       u.nombre, u.apellido, u.correo
		FROM Bitacora b
		JOIN Usuario u ON b.Usuario_idUsuario = u.idUsuario
		WHERE b.entidad = :1
		ORDER BY b.fechaHora DESC
	`

	rows, err := r.db.Query(query, entidad)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var registros []models.BitacoraConUsuario
	for rows.Next() {
		var registro models.BitacoraConUsuario
		err := rows.Scan(
			&registro.IDBitacora,
			&registro.Accion,
			&registro.FechaHora,
			&registro.Detalle,
			&registro.Entidad,
			&registro.UsuarioID,
			&registro.NombreUsuario,
			&registro.ApellidoUsuario,
			&registro.CorreoUsuario,
		)
		if err != nil {
			return nil, err
		}
		registros = append(registros, registro)
	}

	return registros, nil
}
