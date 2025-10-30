package handlers

import (
	"net/http"
	"proyecto-bd-final/database"
	"proyecto-bd-final/internal/repository"
	"proyecto-bd-final/models"
	"proyecto-bd-final/pkg/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UsuarioHandler struct {
	usuarioRepo  *repository.UsuarioRepository
	bitacoraRepo *repository.BitacoraRepository
}

func NewUsuarioHandler() *UsuarioHandler {
	return &UsuarioHandler{
		usuarioRepo:  repository.NewUsuarioRepository(database.DB),
		bitacoraRepo: repository.NewBitacoraRepository(database.DB),
	}
}

func (h *UsuarioHandler) RegisterUsuario(c *gin.Context) {
	var req models.RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithDetail(c, "Datos inválidos", http.StatusBadRequest, err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Contrasenia), bcrypt.DefaultCost)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al encriptar contraseña", http.StatusInternalServerError, err)
		return
	}

	usuario := &models.Usuario{
		Nombre:        req.Nombre,
		Apellido:      req.Apellido,
		Correo:        req.Correo,
		Telefono:      req.Telefono,
		Contrasenia:   string(hashedPassword),
		FechaRegistro: time.Now(),
	}

	existingUser, _ := h.usuarioRepo.GetUsuarioByCorreo(req.Correo)
	if existingUser != nil {
		utils.ErrorResponse(c, "El correo ya está registrado", http.StatusConflict)
		return
	}

	if err := h.usuarioRepo.CreateUsuario(usuario); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al crear usuario", http.StatusInternalServerError, err)
		return
	}

	if err := h.usuarioRepo.AssignRoleToUsuario(usuario.IDUsuario, req.RolID); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al asignar rol al usuario", http.StatusInternalServerError, err)
		return
	}

	switch req.RolID {
	case 2:
		if req.Carrera == nil || req.Semestre == nil {
			utils.ErrorResponse(c, "Datos de estudiante incompletos", http.StatusBadRequest)
			return
		}
		est := &models.Estudiante{
			Carrera:   *req.Carrera,
			Semestre:  *req.Semestre,
			UsuarioID: usuario.IDUsuario,
		}
		h.usuarioRepo.CreateEstudiante(est)

	case 3:
		if req.Facultad == nil {
			utils.ErrorResponse(c, "Datos de profesor incompletos", http.StatusBadRequest)
			return
		}
		prof := &models.Profesor{
			Facultad:  *req.Facultad,
			UsuarioID: usuario.IDUsuario,
		}
		h.usuarioRepo.CreateProfesor(prof)

	case 4:
		if req.Puesto == nil {
			utils.ErrorResponse(c, "Datos de personal incompletos", http.StatusBadRequest)
			return
		}
		per := &models.Personal{
			Puesto:    *req.Puesto,
			UsuarioID: usuario.IDUsuario,
		}
		h.usuarioRepo.CreatePersonal(per)
	}

	bitacora := &models.Bitacora{
		Accion:    "CREAR",
		Detalle:   "Usuario registrado: " + usuario.Correo,
		Entidad:   "Usuario",
		UsuarioID: usuario.IDUsuario,
	}
	h.bitacoraRepo.CreateRegistro(bitacora)

	utils.SuccessResponse(c, "Usuario registrado exitosamente", gin.H{
		"idUsuario": usuario.IDUsuario,
		"correo":    usuario.Correo,
	})
}

func (h *UsuarioHandler) GetPerfilUsuario(c *gin.Context) {
	usuarioID, exists := c.Get("usuarioId")
	if !exists {
		utils.ErrorResponse(c, "No se encontró el usuario en el contexto", http.StatusUnauthorized)
		return
	}

	id := usuarioID.(int)

	usuario, err := h.usuarioRepo.GetUsuarioByID(id)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Usuario no encontrado", http.StatusNotFound, err)
		return
	}

	roles, err := h.usuarioRepo.GetUsuarioRoles(usuario.IDUsuario)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener roles del usuario", http.StatusInternalServerError, err)
		return
	}

	usuarioConRoles := models.UsuarioConRoles{
		Usuario: *usuario,
		Roles:   roles,
	}

	usuarioConRoles.Contrasenia = ""

	utils.SuccessResponse(c, "Perfil obtenido correctamente", usuarioConRoles)
}

func (h *UsuarioHandler) UpdatePerfilUsuario(c *gin.Context) {
    usuarioID, exists := c.Get("usuarioId")
    if !exists {
        utils.ErrorResponse(c, "No autorizado", http.StatusUnauthorized)
        return
    }

    var req models.UpdatePerfilRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponseWithDetail(c, "Datos inválidos", http.StatusBadRequest, err)
        return
    }

    usuario := &models.Usuario{
        IDUsuario: usuarioID.(int),
        Nombre:    req.Nombre,
        Apellido:  req.Apellido,
        Correo:    req.Correo,
        Telefono:  req.Telefono,
    }

    if err := h.usuarioRepo.UpdateUsuario(usuario); err != nil {
        utils.ErrorResponseWithDetail(c, "Error al actualizar perfil", http.StatusInternalServerError, err)
        return
    }

    utils.SuccessResponse(c, "Perfil actualizado correctamente", nil)
}

