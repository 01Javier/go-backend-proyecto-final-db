package handlers

import (
	"net/http"
	"proyecto-bd-final/database"
	"proyecto-bd-final/internal/repository"
	"proyecto-bd-final/models"
	"proyecto-bd-final/pkg/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	usuarioRepo  *repository.UsuarioRepository
	bitacoraRepo *repository.BitacoraRepository
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		usuarioRepo:  repository.NewUsuarioRepository(database.DB),
		bitacoraRepo: repository.NewBitacoraRepository(database.DB),
	}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithDetail(c, "Datos de registro inválidos", http.StatusBadRequest, err)
		return
	}

	existingUser, _ := h.usuarioRepo.GetUsuarioByCorreo(req.Correo)
	if existingUser != nil {
		utils.ErrorResponse(c, "El correo ya está registrado", http.StatusConflict)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Contrasenia), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al procesar la contraseña", http.StatusInternalServerError, err)
		return
	}

	usuario := &models.Usuario{
		Nombre:      req.Nombre,
		Apellido:    req.Apellido,
		Correo:      req.Correo,
		Contrasenia: string(hashedPassword),
		Telefono:    req.Telefono,
	}

	if err := h.usuarioRepo.CreateUsuario(usuario); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al crear el usuario", http.StatusInternalServerError, err)
		return
	}

	if err := h.usuarioRepo.AssignRoleToUsuario(usuario.IDUsuario, req.RolID); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al asignar rol", http.StatusInternalServerError, err)
		return
	}

	switch req.RolID {
	case 1: // Estudiante
		if req.Carnet != nil && req.Carrera != nil && req.Semestre != nil {
			estudiante := &models.Estudiante{
				Carrera:   *req.Carrera,
				Semestre:  *req.Semestre,
				UsuarioID: usuario.IDUsuario,
			}
			if err := h.usuarioRepo.CreateEstudiante(estudiante); err != nil {
				utils.ErrorResponseWithDetail(c, "Error al crear registro de estudiante", http.StatusInternalServerError, err)
				return
			}
		}
	case 2: // Profesor
		if req.Facultad != nil {
			profesor := &models.Profesor{
				Facultad:  *req.Facultad,
				UsuarioID: usuario.IDUsuario,
			}
			if err := h.usuarioRepo.CreateProfesor(profesor); err != nil {
				utils.ErrorResponseWithDetail(c, "Error al crear registro de profesor", http.StatusInternalServerError, err)
				return
			}
		}
	case 3: // Personal/Admin
		if req.Puesto != nil {
			personal := &models.Personal{
				Puesto:    *req.Puesto,
				UsuarioID: usuario.IDUsuario,
			}
			if err := h.usuarioRepo.CreatePersonal(personal); err != nil {
				utils.ErrorResponseWithDetail(c, "Error al crear registro de personal", http.StatusInternalServerError, err)
				return
			}
		}
	}

	bitacora := &models.Bitacora{
		Accion:    "REGISTRO",
		Detalle:   "Nuevo usuario registrado: " + usuario.Correo,
		Entidad:   "Usuario",
		UsuarioID: usuario.IDUsuario,
	}
	h.bitacoraRepo.CreateRegistro(bitacora)

	roles, _ := h.usuarioRepo.GetUsuarioRoles(usuario.IDUsuario)

	usuarioConRoles := models.UsuarioConRoles{
		Usuario: *usuario,
		Roles:   roles,
	}

	token, err := utils.GenerateToken(usuarioConRoles)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al generar token", http.StatusInternalServerError, err)
		return
	}

	usuarioConRoles.Contrasenia = ""

	response := models.AuthResponse{
		Token:   token,
		Usuario: usuarioConRoles,
	}

	utils.SuccessResponse(c, "Usuario registrado exitosamente", response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req models.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithDetail(c, "Datos de login inválidos", http.StatusBadRequest, err)
		return
	}

	usuario, err := h.usuarioRepo.GetUsuarioByCorreo(req.Correo)
	if err != nil {
		utils.ErrorResponse(c, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usuario.Contrasenia), []byte(req.Contrasenia)); err != nil {
		utils.ErrorResponse(c, "Credenciales inválidas", http.StatusUnauthorized)
		return
	}

	roles, err := h.usuarioRepo.GetUsuarioRoles(usuario.IDUsuario)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener roles", http.StatusInternalServerError, err)
		return
	}

	usuarioConRoles := models.UsuarioConRoles{
		Usuario: *usuario,
		Roles:   roles,
	}

	token, err := utils.GenerateToken(usuarioConRoles)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al generar token", http.StatusInternalServerError, err)
		return
	}

	bitacora := &models.Bitacora{
		Accion:    "LOGIN",
		Detalle:   "Usuario inició sesión: " + usuario.Correo,
		Entidad:   "Usuario",
		UsuarioID: usuario.IDUsuario,
	}
	h.bitacoraRepo.CreateRegistro(bitacora)

	usuarioConRoles.Contrasenia = ""

	response := models.AuthResponse{
		Token:   token,
		Usuario: usuarioConRoles,
	}

	utils.SuccessResponse(c, "Login exitoso", response)
}

func (h *AuthHandler) GetProfile(c *gin.Context) {
	usuarioID, exists := c.Get("usuarioId")
	if !exists {
		utils.ErrorResponse(c, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	usuario, err := h.usuarioRepo.GetUsuarioByID(usuarioID.(int))
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener perfil", http.StatusInternalServerError, err)
		return
	}

	roles, err := h.usuarioRepo.GetUsuarioRoles(usuario.IDUsuario)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener roles", http.StatusInternalServerError, err)
		return
	}

	usuarioConRoles := models.UsuarioConRoles{
		Usuario: *usuario,
		Roles:   roles,
	}

	usuarioConRoles.Contrasenia = ""

	utils.SuccessResponse(c, "Perfil obtenido exitosamente", usuarioConRoles)
}
