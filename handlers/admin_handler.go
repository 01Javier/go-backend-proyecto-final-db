package handlers

import (
	"net/http"
	"proyecto-bd-final/database"
	"proyecto-bd-final/internal/repository"
	"proyecto-bd-final/models"
	"proyecto-bd-final/pkg/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	usuarioRepo  *repository.UsuarioRepository
	bitacoraRepo *repository.BitacoraRepository
}

func NewAdminHandler() *AdminHandler {
	return &AdminHandler{
		usuarioRepo:  repository.NewUsuarioRepository(database.DB),
		bitacoraRepo: repository.NewBitacoraRepository(database.DB),
	}
}

func (h *AdminHandler) GetAllUsuarios(c *gin.Context) {
	usuarios, err := h.usuarioRepo.GetAllUsuarios()
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener usuarios", http.StatusInternalServerError, err)
		return
	}

	for i := range usuarios {
		usuarios[i].Contrasenia = ""
	}

	utils.SuccessResponse(c, "Usuarios obtenidos exitosamente", usuarios)
}

func (h *AdminHandler) GetUsuarioByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "ID de usuario inválido", http.StatusBadRequest, err)
		return
	}

	usuario, err := h.usuarioRepo.GetUsuarioByID(id)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Usuario no encontrado", http.StatusNotFound, err)
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

	// Limpiar contraseña
	usuarioConRoles.Contrasenia = ""

	utils.SuccessResponse(c, "Usuario obtenido exitosamente", usuarioConRoles)
}

func (h *AdminHandler) UpdateUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "ID de usuario inválido", http.StatusBadRequest, err)
		return
	}

	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		utils.ErrorResponseWithDetail(c, "Datos inválidos", http.StatusBadRequest, err)
		return
	}

	usuario.IDUsuario = id

	if err := h.usuarioRepo.UpdateUsuario(&usuario); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al actualizar usuario", http.StatusInternalServerError, err)
		return
	}

	adminID, _ := c.Get("usuarioId")
	bitacora := &models.Bitacora{
		Accion:    "ACTUALIZAR",
		Detalle:   "Usuario actualizado: " + usuario.Correo,
		Entidad:   "Usuario",
		UsuarioID: adminID.(int),
	}
	h.bitacoraRepo.CreateRegistro(bitacora)

	utils.SuccessResponse(c, "Usuario actualizado exitosamente", usuario)
}

func (h *AdminHandler) DeleteUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "ID de usuario inválido", http.StatusBadRequest, err)
		return
	}

	usuario, err := h.usuarioRepo.GetUsuarioByID(id)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Usuario no encontrado", http.StatusNotFound, err)
		return
	}

	if err := h.usuarioRepo.DeleteUsuario(id); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al eliminar usuario", http.StatusInternalServerError, err)
		return
	}

	adminID, _ := c.Get("usuarioId")
	bitacora := &models.Bitacora{
		Accion:    "ELIMINAR",
		Detalle:   "Usuario eliminado: " + usuario.Correo,
		Entidad:   "Usuario",
		UsuarioID: adminID.(int),
	}
	h.bitacoraRepo.CreateRegistro(bitacora)

	utils.SuccessResponse(c, "Usuario eliminado exitosamente", nil)
}

func (h *AdminHandler) BuscarUsuarios(c *gin.Context) {
	termino := c.Query("q")
	if termino == "" {
		utils.ErrorResponse(c, "Parámetro de búsqueda vacío", http.StatusBadRequest)
		return
	}

	resultados, err := h.usuarioRepo.BuscarUsuariosPorNombreOCorreo(termino)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error en búsqueda", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Usuarios encontrados", resultados)
}
