package handlers

import (
	"net/http"
	"proyecto-bd-final/database"
	"proyecto-bd-final/internal/repository"
	"proyecto-bd-final/pkg/utils"

	"github.com/gin-gonic/gin"
)

type EstadisticasHandler struct {
	prestamoRepo *repository.PrestamoRepository
	bitacoraRepo *repository.BitacoraRepository
}

func NewEstadisticasHandler() *EstadisticasHandler {
	return &EstadisticasHandler{
		prestamoRepo: repository.NewPrestamoRepository(database.DB),
		bitacoraRepo: repository.NewBitacoraRepository(database.DB),
	}
}

func (h *EstadisticasHandler) GetEstadisticas(c *gin.Context) {
	stats, err := h.prestamoRepo.GetEstadisticas()
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener estadísticas", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Estadísticas obtenidas exitosamente", stats)
}

func (h *EstadisticasHandler) GetPrestamosVencidos(c *gin.Context) {
	prestamos, err := h.prestamoRepo.GetPrestamosVencidos()
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener préstamos vencidos", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Préstamos vencidos obtenidos exitosamente", prestamos)
}

func (h *EstadisticasHandler) GetAllPrestamos(c *gin.Context) {
	prestamos, err := h.prestamoRepo.GetAllPrestamos()
	if err != nil {
		utils.ErrorResponseWithDetail(c, "Error al obtener préstamos", http.StatusInternalServerError, err)
		return
	}

	utils.SuccessResponse(c, "Préstamos obtenidos exitosamente", prestamos)
}
