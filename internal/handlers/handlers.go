package handlers

import (
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"leadgen-test-task/internal/model"
	"leadgen-test-task/internal/repository"
)

type BuildingHandler struct {
	logger *zap.SugaredLogger
	repo   *repository.BuildingRepository
}

func NewBuildingHandler(repo *repository.BuildingRepository, logger *zap.SugaredLogger) *BuildingHandler {
	return &BuildingHandler{repo: repo, logger: logger}
}

// @Summary Create a new building
// @Description Add a new building to the database
// @Tags buildings
// @Accept json
// @Produce json
// @Param building body model.Building true "Building Info"
// @Success 201 {object} model.Building
// @Router /buildings [post]
func (h *BuildingHandler) CreateBuilding(c *gin.Context) {
	var building model.Building
	if err := c.ShouldBindJSON(&building); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Errorf("Failed to parse request body: %s", err)
		return
	}

	if err := h.repo.Create(&building); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Errorf("Failed to create building: %s", err)
		return
	}

	c.JSON(http.StatusCreated, building)
}

// @Summary List buildings
// @Description Get list of buildings with optional filters
// @Tags buildings
// @Accept json
// @Produce json
// @Param city query string false "City filter"
// @Param year query int false "Year filter"
// @Param floor_count query int false "Floor count filter"
// @Success 200 {array} model.Building
// @Router /buildings [get]
func (h *BuildingHandler) ListBuildings(c *gin.Context) {
	city := c.Query("city")
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Errorf("Failed to parse year: %s", err)
		return
	}
	floorCount, err := strconv.Atoi(c.Query("floor_count"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		h.logger.Errorf("Failed to parse request floor_count: %s", err)
		return
	}

	buildings, err := h.repo.FindAll(city, year, floorCount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		h.logger.Errorf("Failed to find buildings: %s", err)
		return
	}

	c.JSON(http.StatusOK, buildings)
}
