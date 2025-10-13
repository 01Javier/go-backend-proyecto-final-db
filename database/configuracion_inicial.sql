CREATE SEQUENCE usuario_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE estudiante_seq START WITH 1000 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE profesor_seq START WITH 2000 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE personal_seq START WITH 3000 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE roles_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE usuariorol_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE bitacora_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE libro_seq START WITH 1000 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE editorial_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE autor_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE libroautor_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE prestamo_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE ejemplar_seq START WITH 1000 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE permiso_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;
CREATE SEQUENCE rolpermiso_seq START WITH 1 INCREMENT BY 1 NOCACHE NOCYCLE;

INSERT INTO Roles (idRol, nombreRol) VALUES (1, 'estudiante');
INSERT INTO Roles (idRol, nombreRol) VALUES (2, 'profesor');
INSERT INTO Roles (idRol, nombreRol) VALUES (3, 'admin');
INSERT INTO Roles (idRol, nombreRol) VALUES (4, 'personal');





-- ============================================
-- 1. EDITORIALES
-- ============================================
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Pearson', 'Estados Unidos');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'McGraw-Hill', 'Estados Unidos');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Oxford University Press', 'Reino Unido');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Springer', 'Alemania');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Editorial Universitaria', 'Chile');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Alfaomega', 'México');
INSERT INTO Editorial (idEditorial, nombre, pais) VALUES (editorial_seq.NEXTVAL, 'Anaya Multimedia', 'España');

-- ============================================
-- 2. AUTORES
-- ============================================
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Abraham', 'Silberschatz', 'Estados Unidos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Andrew', 'Tanenbaum', 'Países Bajos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Bjarne', 'Stroustrup', 'Dinamarca');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Robert', 'Martin', 'Estados Unidos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Martin', 'Fowler', 'Reino Unido');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Donald', 'Knuth', 'Estados Unidos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Thomas', 'Cormen', 'Estados Unidos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Grady', 'Booch', 'Estados Unidos');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Gang of', 'Four', 'Internacional');
INSERT INTO Autor (idAutor, nombre, apellido, nacionalidad) VALUES (autor_seq.NEXTVAL, 'Eric', 'Evans', 'Estados Unidos');

-- ============================================
-- 3. LIBROS
-- ============================================
INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Database System Concepts', TO_DATE('2020', 'YYYY'), 1);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Modern Operating Systems', TO_DATE('2021', 'YYYY'), 1);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'The C++ Programming Language', TO_DATE('2019', 'YYYY'), 2);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Clean Code', TO_DATE('2022', 'YYYY'), 1);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Refactoring: Improving the Design of Existing Code', TO_DATE('2021', 'YYYY'), 2);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'The Art of Computer Programming Vol. 1', TO_DATE('2018', 'YYYY'), 2);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Introduction to Algorithms', TO_DATE('2022', 'YYYY'), 4);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Design Patterns', TO_DATE('2020', 'YYYY'), 2);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Domain-Driven Design', TO_DATE('2019', 'YYYY'), 2);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Computer Networks', TO_DATE('2021', 'YYYY'), 1);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Artificial Intelligence: A Modern Approach', TO_DATE('2023', 'YYYY'), 1);

INSERT INTO Libro (ISBN, titulo, anioEdicion, Editorial_idEditorial) 
VALUES (libro_seq.NEXTVAL, 'Software Engineering', TO_DATE('2020', 'YYYY'), 3);

-- ============================================
-- 4. LIBRO-AUTOR (Relación muchos a muchos)
-- ============================================
INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 1, 1000);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 2, 1001);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 3, 1002);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 4, 1003);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 5, 1004);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 6, 1005);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 7, 1006);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Coautor', 8, 1006);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 9, 1007);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 10, 1008);

INSERT INTO LibroAutor (idLibroAutor, tipoAutor, Autor_idAutor, Libro_ISBN) 
VALUES (libroautor_seq.NEXTVAL, 'Principal', 2, 1009);

-- ============================================
-- 5. USUARIOS (Base)
-- ============================================
-- Admin
INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Carlos', 'Administrador', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'admin@biblioteca.usac.edu.gt', 12345678, SYSDATE);

-- Estudiantes
INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Ana', 'García', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'ana.garcia@est.usac.edu.gt', 23456789, SYSDATE);

INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Luis', 'Martínez', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'luis.martinez@est.usac.edu.gt', 34567890, SYSDATE);

INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'María', 'López', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'maria.lopez@est.usac.edu.gt', 45678901, SYSDATE);

INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Jorge', 'Ramírez', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'jorge.ramirez@est.usac.edu.gt', 56789012, SYSDATE);

-- Profesores
INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Dr. Roberto', 'Pérez', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'roberto.perez@prof.usac.edu.gt', 67890123, SYSDATE);

INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Dra. Carmen', 'Morales', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'carmen.morales@prof.usac.edu.gt', 78901234, SYSDATE);

-- Personal
INSERT INTO Usuario (idUsuario, nombre, apellido, contrasenia, correo, telefono, fechaRegistro) 
VALUES (usuario_seq.NEXTVAL, 'Pedro', 'Bibliotecario', '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy', 'pedro.bibliotecario@usac.edu.gt', 89012345, SYSDATE);

-- ============================================
-- 6. ESTUDIANTES
-- ============================================
INSERT INTO Estudiante (carnet, carrera, semestre, Usuario_idUsuario) 
VALUES (estudiante_seq.NEXTVAL, 'Ingeniería en Ciencias y Sistemas', 6, 2);

INSERT INTO Estudiante (carnet, carrera, semestre, Usuario_idUsuario) 
VALUES (estudiante_seq.NEXTVAL, 'Ingeniería en Ciencias y Sistemas', 4, 3);

INSERT INTO Estudiante (carnet, carrera, semestre, Usuario_idUsuario) 
VALUES (estudiante_seq.NEXTVAL, 'Ingeniería Industrial', 3, 4);

INSERT INTO Estudiante (carnet, carrera, semestre, Usuario_idUsuario) 
VALUES (estudiante_seq.NEXTVAL, 'Ingeniería Mecánica', 5, 5);

-- ============================================
-- 7. PROFESORES
-- ============================================
INSERT INTO Profesor (codigoDocencia, facultad, Usuario_idUsuario) 
VALUES (profesor_seq.NEXTVAL, 'Ingeniería', 6);

INSERT INTO Profesor (codigoDocencia, facultad, Usuario_idUsuario) 
VALUES (profesor_seq.NEXTVAL, 'Ingeniería', 7);

-- ============================================
-- 8. PERSONAL
-- ============================================
INSERT INTO Personal (codigoEmpleado, puesto, Usuario_idUsuario) 
VALUES (personal_seq.NEXTVAL, 'Bibliotecario', 8);

-- ============================================
-- 9. USUARIO-ROL (Asignar roles)
-- ============================================
-- Admin
INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 1, 3);

-- Estudiantes
INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 2, 1);

INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 3, 1);

INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 4, 1);

INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 5, 1);

-- Profesores
INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 6, 2);

INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 7, 2);

-- Personal
INSERT INTO UsuarioRol (idUsuarioRol, Usuario_idUsuario, Roles_idRol) 
VALUES (usuariorol_seq.NEXTVAL, 8, 4);

-- ============================================
-- 10. PRÉSTAMOS (Algunos activos, algunos vencidos)
-- ============================================
-- Préstamo activo
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 5, SYSDATE + 10, NULL, 'activo', 2, 0);

-- Préstamo activo
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 3, SYSDATE + 12, NULL, 'activo', 3, 0);

-- Préstamo vencido (fecha prevista ya pasó)
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 20, SYSDATE - 5, NULL, 'vencido', 4, 0);

-- Préstamo vencido
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 25, SYSDATE - 10, NULL, 'vencido', 5, 0);

-- Préstamo devuelto (completado)
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 30, SYSDATE - 16, SYSDATE - 15, 'devuelto', 2, 0);

-- Préstamo activo profesor
INSERT INTO Prestamo (idPrestamo, fechaPrestamo, fechaDevolucionPrevista, fechaDevolucionReal, estado, Usuario_idUsuario, Devolucion_idDevolucion) 
VALUES (prestamo_seq.NEXTVAL, SYSDATE - 2, SYSDATE + 28, NULL, 'activo', 6, 0);

-- ============================================
-- 11. EJEMPLARES (Copias físicas de libros)
-- ============================================
-- Database System Concepts - 3 copias
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'prestado', 1000, 1);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1000, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1000, 0);

-- Modern Operating Systems - 2 copias
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'prestado', 1001, 2);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1001, 0);

-- Clean Code - 3 copias
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'prestado', 1003, 3);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1003, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'mantenimiento', 1003, 0);

-- Introduction to Algorithms - 2 copias
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'prestado', 1006, 4);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1006, 0);

-- Design Patterns - 2 copias
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'prestado', 1007, 6);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1007, 0);

-- Resto de libros con 1-2 copias disponibles
INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1002, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1004, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1005, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1008, 0);

INSERT INTO Ejemplar (codigo, estado, Libro_ISBN, Prestamo_idPrestamo) 
VALUES (ejemplar_seq.NEXTVAL, 'disponible', 1009, 0);

-- ============================================
-- 12. PERMISOS
-- ============================================
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Gestionar Usuarios');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Gestionar Libros');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Gestionar Préstamos');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Ver Estadísticas');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Ver Bitácora');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Realizar Préstamos');
INSERT INTO Permiso (idPermiso, descripcion) VALUES (permiso_seq.NEXTVAL, 'Devolver Libros');

-- ============================================
-- 13. ROL-PERMISO
-- ============================================
-- Admin (todos los permisos)
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 1);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 2);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 3);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 4);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 5);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 6);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 3, 7);

-- Estudiante (permisos limitados)
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 1, 6);

-- Profesor (más permisos que estudiante)
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 2, 4);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 2, 6);

-- Personal (permisos operativos)
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 4, 2);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 4, 3);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 4, 6);
INSERT INTO RolPermiso (idRolPermiso, Roles_idRol, Permiso_idPermiso) VALUES (rolpermiso_seq.NEXTVAL, 4, 7);

-- ============================================
-- 14. BITÁCORA (Registros de actividad)
-- ============================================
INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'LOGIN', SYSDATE - 5, 'Usuario inició sesión', 'Usuario', 1);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'CREAR_PRESTAMO', SYSDATE - 5, 'Se creó préstamo para libro Database System Concepts', 'Prestamo', 8);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'LOGIN', SYSDATE - 3, 'Usuario inició sesión', 'Usuario', 2);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'CREAR_PRESTAMO', SYSDATE - 3, 'Se creó préstamo para libro Modern Operating Systems', 'Prestamo', 8);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'DEVOLVER_LIBRO', SYSDATE - 15, 'Se devolvió libro Clean Code', 'Prestamo', 8);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'CREAR_USUARIO', SYSDATE - 30, 'Se registró nuevo estudiante', 'Usuario', 1);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'LOGIN', SYSDATE - 2, 'Usuario inició sesión', 'Usuario', 6);

INSERT INTO Bitacora (idBitacora, accion, fechaHora, detalle, entidad, Usuario_idUsuario) 
VALUES (bitacora_seq.NEXTVAL, 'VER_ESTADISTICAS', SYSDATE - 1, 'Consultó estadísticas del sistema', 'Estadisticas', 1);

COMMIT;

-- ============================================
-- VERIFICACIÓN DE DATOS
-- ============================================
SELECT 'Editoriales' as Tabla, COUNT(*) as Cantidad FROM Editorial
UNION ALL
SELECT 'Autores', COUNT(*) FROM Autor
UNION ALL
SELECT 'Libros', COUNT(*) FROM Libro
UNION ALL
SELECT 'LibroAutor', COUNT(*) FROM LibroAutor
UNION ALL
SELECT 'Usuarios', COUNT(*) FROM Usuario
UNION ALL
SELECT 'Estudiantes', COUNT(*) FROM Estudiante
UNION ALL
SELECT 'Profesores', COUNT(*) FROM Profesor
UNION ALL
SELECT 'Personal', COUNT(*) FROM Personal
UNION ALL
SELECT 'UsuarioRol', COUNT(*) FROM UsuarioRol
UNION ALL
SELECT 'Préstamos', COUNT(*) FROM Prestamo
UNION ALL
SELECT 'Ejemplares', COUNT(*) FROM Ejemplar
UNION ALL
SELECT 'Permisos', COUNT(*) FROM Permiso
UNION ALL
SELECT 'RolPermiso', COUNT(*) FROM RolPermiso
UNION ALL
SELECT 'Bitácora', COUNT(*) FROM Bitacora;