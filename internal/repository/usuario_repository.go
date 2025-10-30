package repository

import (
	"database/sql"
	"errors"
	"proyecto-bd-final/models"
	"time"
)

type UsuarioRepository struct {
	db *sql.DB
}

func NewUsuarioRepository(db *sql.DB) *UsuarioRepository {
	return &UsuarioRepository{db: db}
}

func (r *UsuarioRepository) CreateUsuario(usuario *models.Usuario) error {
	query := `
		INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro)
		VALUES (usuario_seq.NEXTVAL, :1, :2, :3, :4, :5, :6)
		RETURNING idUsuario INTO :7
	`

	var id int
	err := r.db.QueryRow(query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.Contrasenia,
		usuario.Correo,
		usuario.Telefono,
		time.Now(),
	).Scan(&id)

	if err != nil {
		return err
	}

	usuario.IDUsuario = id
	usuario.FechaRegistro = time.Now()
	return nil
}

func (r *UsuarioRepository) GetUsuarioByCorreo(correo string) (*models.Usuario, error) {
	query := `
		SELECT idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro
		FROM Usuario
		WHERE correo = :1
	`

	var usuario models.Usuario
	err := r.db.QueryRow(query, correo).Scan(
		&usuario.IDUsuario,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.Contrasenia,
		&usuario.Correo,
		&usuario.Telefono,
		&usuario.FechaRegistro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) GetUsuarioByID(id int) (*models.Usuario, error) {
	query := `
		SELECT idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro
		FROM Usuario
		WHERE idUsuario = :1
	`

	var usuario models.Usuario
	err := r.db.QueryRow(query, id).Scan(
		&usuario.IDUsuario,
		&usuario.Nombre,
		&usuario.Apellido,
		&usuario.Contrasenia,
		&usuario.Correo,
		&usuario.Telefono,
		&usuario.FechaRegistro,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	return &usuario, nil
}

func (r *UsuarioRepository) GetUsuarioRoles(usuarioID int) ([]string, error) {
	query := `
		SELECT r.nombreRol
		FROM UsuarioRol ur
		JOIN Roles r ON ur.Roles_idRol = r.idRol
		WHERE ur.Usuario_idUsuario = :1
	`

	rows, err := r.db.Query(query, usuarioID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []string
	for rows.Next() {
		var rol string
		if err := rows.Scan(&rol); err != nil {
			return nil, err
		}
		roles = append(roles, rol)
	}

	return roles, nil
}

func (r *UsuarioRepository) AssignRoleToUsuario(usuarioID, rolID int) error {
	query := `
		INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol)
		VALUES (usuariorol_seq.NEXTVAL, :1, :2)
	`

	_, err := r.db.Exec(query, usuarioID, rolID)
	return err
}

func (r *UsuarioRepository) CreateEstudiante(estudiante *models.Estudiante) error {
	query := `
		INSERT INTO Estudiante (carnet, carrera, semestre, Usuario_idUsuario)
		VALUES (estudiante_seq.NEXTVAL, :1, :2, :3)
		RETURNING carnet INTO :4
	`

	var carnet int
	err := r.db.QueryRow(query,
		estudiante.Carrera,
		estudiante.Semestre,
		estudiante.UsuarioID,
	).Scan(&carnet)

	if err != nil {
		return err
	}

	estudiante.Carnet = carnet
	return nil
}

func (r *UsuarioRepository) CreateProfesor(profesor *models.Profesor) error {
	query := `
		INSERT INTO Profesor (codigoDocencia, facultad, Usuario_idUsuario)
		VALUES (profesor_seq.NEXTVAL, :1, :2)
		RETURNING codigoDocencia INTO :3
	`

	var codigo int
	err := r.db.QueryRow(query,
		profesor.Facultad,
		profesor.UsuarioID,
	).Scan(&codigo)

	if err != nil {
		return err
	}

	profesor.CodigoDocencia = codigo
	return nil
}

func (r *UsuarioRepository) CreatePersonal(personal *models.Personal) error {
	query := `
		INSERT INTO Personal (codigoEmpleado, puesto, Usuario_idUsuario)
		VALUES (personal_seq.NEXTVAL, :1, :2)
		RETURNING codigoEmpleado INTO :3
	`

	var codigo int
	err := r.db.QueryRow(query,
		personal.Puesto,
		personal.UsuarioID,
	).Scan(&codigo)

	if err != nil {
		return err
	}

	personal.CodigoEmpleado = codigo
	return nil
}

func (r *UsuarioRepository) GetAllUsuarios() ([]models.UsuarioConRoles, error) {
	query := `
		SELECT DISTINCT u.idUsuario, u.nombre, u.apellido, u.correo, u.telefono, u.fechaRegistro
		FROM Usuario u
		ORDER BY u.fechaRegistro DESC
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.UsuarioConRoles
	for rows.Next() {
		var usuario models.UsuarioConRoles
		err := rows.Scan(
			&usuario.IDUsuario,
			&usuario.Nombre,
			&usuario.Apellido,
			&usuario.Correo,
			&usuario.Telefono,
			&usuario.FechaRegistro,
		)
		if err != nil {
			return nil, err
		}

		roles, err := r.GetUsuarioRoles(usuario.IDUsuario)
		if err != nil {
			return nil, err
		}
		usuario.Roles = roles

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (r *UsuarioRepository) DeleteUsuario(id int) error {
	query := `DELETE FROM Usuario WHERE idUsuario = :1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}

func (r *UsuarioRepository) UpdateUsuario(usuario *models.Usuario) error {
	query := `
		UPDATE Usuario
		SET nombre = :1, apellido = :2, correo = :3, telefono = :4
		WHERE idUsuario = :5
	`

	result, err := r.db.Exec(query,
		usuario.Nombre,
		usuario.Apellido,
		usuario.Correo,
		usuario.Telefono,
		usuario.IDUsuario,
	)

	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("usuario no encontrado")
	}

	return nil
}
func (r *UsuarioRepository) GetPrestamosByUsuarioID(id int) ([]models.Prestamo, error) {
	query := `
        SELECT idPrestamo, Libro_idLibro, fechaPrestamo, fechaDevolucion, estado
        FROM Prestamo
        WHERE Usuario_idUsuario = :1
        ORDER BY fechaPrestamo DESC
    `

	rows, err := r.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var prestamos []models.Prestamo
	for rows.Next() {
		var p models.Prestamo
		err := rows.Scan(&p.IDPrestamo, &p.FechaPrestamo, &p.FechaDevolucionPrevista, &p.Estado)
		if err != nil {
			return nil, err
		}
		prestamos = append(prestamos, p)
	}

	return prestamos, nil
}

func (r *UsuarioRepository) BuscarUsuariosPorNombreOCorreo(termino string) ([]models.UsuarioConRoles, error) {
	query := `
        SELECT idUsuario, nombre, apellido, correo, telefono, fechaRegistro
        FROM Usuario
        WHERE LOWER(nombre) LIKE '%' || LOWER(:1) || '%'
           OR LOWER(apellido) LIKE '%' || LOWER(:1) || '%'
           OR LOWER(correo) LIKE '%' || LOWER(:1) || '%'
    `

	rows, err := r.db.Query(query, termino)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []models.UsuarioConRoles
	for rows.Next() {
		var u models.UsuarioConRoles
		err := rows.Scan(&u.IDUsuario, &u.Nombre, &u.Apellido, &u.Correo, &u.Telefono, &u.FechaRegistro)
		if err != nil {
			return nil, err
		}
		u.Roles, _ = r.GetUsuarioRoles(u.IDUsuario)
		usuarios = append(usuarios, u)
	}

	return usuarios, nil
}
