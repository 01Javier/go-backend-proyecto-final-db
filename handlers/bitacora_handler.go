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

type BitacoraHandler struct {
	bitacoraRepo *repository.BitacoraRepository
}

func NewBitacoraHandler() *BitacoraHandler {
	return &BitacoraHandler{
		bitacoraRepo: repository.NewBitacoraRepository(database.DB),
	}
}

func (h *BitacoraHandler) GetAllRegistros(c *gin.Context) {
	limitStr := c.DefaultQuery("limit", "100")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 100
	}

	registros, err := h.bitacoraRepo.GetAllRegistros(limit)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener registros de bitácora", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Registros obtenidos exitosamente", registros)
}

func (h *BitacoraHandler) GetRegistrosByUsuario(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "ID de usuario inválido", http.StatusBadRequest, err)
		return
	}

	registros, err := h.bitacoraRepo.GetRegistrosByUsuario(id)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener registros", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Registros obtenidos exitosamente", registros)
}

func (h *BitacoraHandler) GetRegistrosByEntidad(c *gin.Context) {
	entidad := c.Param("entidad")

	registros, err := h.bitacoraRepo.GetRegistrosByEntidad(entidad)
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener registros", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Registros obtenidos exitosamente", registros)
}

func (h *BitacoraHandler) CreateRegistro(c *gin.Context) {
	var req models.RegistroBitacoraRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		utils.ErrorResponseWithDetail(c, "Datos inválidos", http.StatusBadRequest, err)
		return
	}

	usuarioID, exists := c.Get("usuarioId")
	if !exists {
		utils.ErrorResponse(c, "Usuario no autenticado", http.StatusUnauthorized)
		return
	}

	bitacora := &models.Bitacora{
		Accion:    req.Accion,
		Detalle:   req.Detalle,
		Entidad:   req.Entidad,
		UsuarioID: usuarioID.(int),
	}

	if err := h.bitacoraRepo.CreateRegistro(bitacora); err != nil {
		utils.ErrorResponseWithDetail(c, "Error al crear registro", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Registro creado exitosamente", bitacora)
}
