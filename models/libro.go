package models

import "time"

type Libro struct {
	ISBN        int       `json:"isbn" db:"ISBN"`
	Titulo      string    `json:"titulo" db:"TITULO"`
	AnioEdicion time.Time `json:"anioEdicion" db:"ANIOEDICION"`
	EditorialID int       `json:"editorialId" db:"EDITORIAL_IDEDITORIAL"`
}

type Editorial struct {
	IDEditorial int    `json:"idEditorial" db:"IDEDITORIAL"`
	Nombre      string `json:"nombre" db:"NOMBRE"`
	Pais        string `json:"pais" db:"PAIS"`
}

type Autor struct {
	IDAutor      int    `json:"idAutor" db:"IDAUTOR"`
	Nombre       string `json:"nombre" db:"NOMBRE"`
	Apellido     string `json:"apellido" db:"APELLIDO"`
	Nacionalidad string `json:"nacionalidad" db:"NACIONALIDAD"`
}

type LibroAutor struct {
	IDLibroAutor int    `json:"idLibroAutor" db:"IDLIBROAUTOR"`
	TipoAutor    string `json:"tipoAutor" db:"TIPOAUTOR"`
	AutorID      int    `json:"autorId" db:"AUTOR_IDAUTOR"`
	LibroISBN    int    `json:"libroIsbn" db:"LIBRO_ISBN"`
}

type LibroConDetalles struct {
	Libro
	Editorial   string   `json:"editorial"`
	Autores     []string `json:"autores"`
	Disponibles int      `json:"disponibles"`
	Total       int      `json:"total"`
}
