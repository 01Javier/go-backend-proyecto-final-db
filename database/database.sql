DROP TABLE Autor CASCADE CONSTRAINT;
DROP TABLE Bitacora CASCADE CONSTRAINT;
DROP TABLE Editorial CASCADE CONSTRAINT;
DROP TABLE Ejemplar CASCADE CONSTRAINT;
DROP TABLE Estudiante CASCADE CONSTRAINT;
DROP TABLE Libro CASCADE CONSTRAINT;
DROP TABLE LibroAutor CASCADE CONSTRAINT;
DROP TABLE Permiso CASCADE CONSTRAINT;
DROP TABLE Personal CASCADE CONSTRAINT;
DROP TABLE Prestamo CASCADE CONSTRAINT;
DROP TABLE Profesor CASCADE CONSTRAINT;
DROP TABLE Roles CASCADE CONSTRAINT;
DROP TABLE RolPermiso CASCADE CONSTRAINT;
DROP TABLE Usuario CASCADE CONSTRAINT;
DROP TABLE UsuarioRol CASCADE CONSTRAINT;

CREATE TABLE Autor 
    ( 
     idAutor      INTEGER  NOT NULL , 
     nombre       VARCHAR2 (100) , 
     apellido     VARCHAR2 (100) , 
     nacionalidad VARCHAR2 (100) 
    ) 
;

ALTER TABLE Autor 
    ADD CONSTRAINT Autor_PK PRIMARY KEY ( idAutor ) ;

CREATE TABLE Bitacora 
    ( 
     idBitacora        INTEGER  NOT NULL , 
     accion            VARCHAR2 (100) , 
     fechaHora         DATE , 
     detalle           VARCHAR2 (500) , 
     entidad           VARCHAR2 (50) , 
     Usuario_idUsuario INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Bitacora 
    ADD CONSTRAINT Bitacora_PK PRIMARY KEY ( idBitacora ) ;

CREATE TABLE Editorial 
    ( 
     idEditorial INTEGER  NOT NULL , 
     nombre      VARCHAR2 (100) , 
     pais        VARCHAR2 (100) 
    ) 
;

ALTER TABLE Editorial 
    ADD CONSTRAINT Editorial_PK PRIMARY KEY ( idEditorial ) ;

CREATE TABLE Ejemplar 
    ( 
     codigo              INTEGER  NOT NULL , 
     estado              VARCHAR2 (50) , 
     Libro_ISBN          INTEGER  NOT NULL , 
     Prestamo_idPrestamo INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Ejemplar 
    ADD CONSTRAINT Ejemplar_PK PRIMARY KEY ( codigo ) ;

CREATE TABLE Estudiante 
    ( 
     carnet            INTEGER  NOT NULL , 
     carrera           VARCHAR2 (100) , 
     semestre          INTEGER , 
     Usuario_idUsuario INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Estudiante 
    ADD CONSTRAINT Estudiante_PK PRIMARY KEY ( carnet ) ;

CREATE TABLE Libro 
    ( 
     ISBN                  INTEGER  NOT NULL , 
     titulo                VARCHAR2 (200) , 
     anioEdicion           DATE , 
     Editorial_idEditorial INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Libro 
    ADD CONSTRAINT Libro_PK PRIMARY KEY ( ISBN ) ;

CREATE TABLE LibroAutor 
    ( 
     idLibroAutor  INTEGER  NOT NULL , 
     tipoAutor     VARCHAR2 (100) , 
     Autor_idAutor INTEGER  NOT NULL , 
     Libro_ISBN    INTEGER  NOT NULL 
    ) 
;

ALTER TABLE LibroAutor 
    ADD CONSTRAINT LibroAutor_PK PRIMARY KEY ( idLibroAutor ) ;

CREATE TABLE Permiso 
    ( 
     idPermiso   INTEGER  NOT NULL , 
     descripcion VARCHAR2 (100) 
    ) 
;

ALTER TABLE Permiso 
    ADD CONSTRAINT Permiso_PK PRIMARY KEY ( idPermiso ) ;

CREATE TABLE Personal 
    ( 
     codigoEmpleado    INTEGER  NOT NULL , 
     puesto            VARCHAR2 (100) , 
     Usuario_idUsuario INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Personal 
    ADD CONSTRAINT Personal_PK PRIMARY KEY ( codigoEmpleado ) ;

CREATE TABLE Prestamo 
    ( 
     idPrestamo              INTEGER  NOT NULL , 
     fechaPrestamo           DATE , 
     fechaDevolucionPrevista DATE , 
     fechaDevolucionReal     DATE , 
     estado                  VARCHAR2 (100) , 
     Usuario_idUsuario       INTEGER  NOT NULL
    ) 
;

ALTER TABLE Prestamo 
    ADD CONSTRAINT Prestamo_PK PRIMARY KEY ( idPrestamo ) ;

CREATE TABLE Profesor 
    ( 
     codigoDocencia    INTEGER  NOT NULL , 
     facultad          VARCHAR2 (100) , 
     Usuario_idUsuario INTEGER  NOT NULL 
    ) 
;

ALTER TABLE Profesor 
    ADD CONSTRAINT Profesor_PK PRIMARY KEY ( codigoDocencia ) ;

CREATE TABLE Roles 
    ( 
     idRol     INTEGER  NOT NULL , 
     nombreRol VARCHAR2 (100) 
    ) 
;

ALTER TABLE Roles 
    ADD CONSTRAINT Roles_PK PRIMARY KEY ( idRol ) ;

CREATE TABLE RolPermiso 
    ( 
     idRolPermiso      INTEGER  NOT NULL , 
     Roles_idRol       INTEGER  NOT NULL , 
     Permiso_idPermiso INTEGER  NOT NULL 
    ) 
;

ALTER TABLE RolPermiso 
    ADD CONSTRAINT RolPermiso_PK PRIMARY KEY ( idRolPermiso ) ;

CREATE TABLE Usuario 
    ( 
     idUsuario     INTEGER  NOT NULL , 
     nombre        VARCHAR2 (50) , 
     apellido      VARCHAR2 (50) , 
     contrasenia   VARCHAR2 (200) , 
     correo        VARCHAR2 (200) , 
     telefono      INTEGER , 
     fechaRegistro DATE 
    ) 
;

ALTER TABLE Usuario 
    ADD CONSTRAINT Usuario_PK PRIMARY KEY ( idUsuario ) ;

CREATE TABLE UsuarioRol 
    ( 
     idUsuarioRol      INTEGER  NOT NULL , 
     Usuario_idUsuario INTEGER  NOT NULL , 
     Roles_idRol       INTEGER  NOT NULL 
    ) 
;

ALTER TABLE UsuarioRol 
    ADD CONSTRAINT UsuarioRol_PK PRIMARY KEY ( idUsuarioRol ) ;

ALTER TABLE Bitacora 
    ADD CONSTRAINT Bitacora_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

ALTER TABLE Ejemplar 
    ADD CONSTRAINT Ejemplar_Libro_FK FOREIGN KEY 
    ( 
     Libro_ISBN
    ) 
    REFERENCES Libro 
    ( 
     ISBN
    ) 
;

ALTER TABLE Ejemplar 
    ADD CONSTRAINT Ejemplar_Prestamo_FK FOREIGN KEY 
    ( 
     Prestamo_idPrestamo
    ) 
    REFERENCES Prestamo 
    ( 
     idPrestamo
    ) 
;

ALTER TABLE Estudiante 
    ADD CONSTRAINT Estudiante_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

ALTER TABLE Libro 
    ADD CONSTRAINT Libro_Editorial_FK FOREIGN KEY 
    ( 
     Editorial_idEditorial
    ) 
    REFERENCES Editorial 
    ( 
     idEditorial
    ) 
;

ALTER TABLE LibroAutor 
    ADD CONSTRAINT LibroAutor_Autor_FK FOREIGN KEY 
    ( 
     Autor_idAutor
    ) 
    REFERENCES Autor 
    ( 
     idAutor
    ) 
;

ALTER TABLE LibroAutor 
    ADD CONSTRAINT LibroAutor_Libro_FK FOREIGN KEY 
    ( 
     Libro_ISBN
    ) 
    REFERENCES Libro 
    ( 
     ISBN
    ) 
;

ALTER TABLE Personal 
    ADD CONSTRAINT Personal_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

ALTER TABLE Prestamo 
    ADD CONSTRAINT Prestamo_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

ALTER TABLE Profesor 
    ADD CONSTRAINT Profesor_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

ALTER TABLE RolPermiso 
    ADD CONSTRAINT RolPermiso_Permiso_FK FOREIGN KEY 
    ( 
     Permiso_idPermiso
    ) 
    REFERENCES Permiso 
    ( 
     idPermiso
    ) 
;

ALTER TABLE RolPermiso 
    ADD CONSTRAINT RolPermiso_Roles_FK FOREIGN KEY 
    ( 
     Roles_idRol
    ) 
    REFERENCES Roles 
    ( 
     idRol
    ) 
;

ALTER TABLE UsuarioRol 
    ADD CONSTRAINT UsuarioRol_Roles_FK FOREIGN KEY 
    ( 
     Roles_idRol
    ) 
    REFERENCES Roles 
    ( 
     idRol
    ) 
;

ALTER TABLE UsuarioRol 
    ADD CONSTRAINT UsuarioRol_Usuario_FK FOREIGN KEY 
    ( 
     Usuario_idUsuario
    ) 
    REFERENCES Usuario 
    ( 
     idUsuario
    ) 
;

