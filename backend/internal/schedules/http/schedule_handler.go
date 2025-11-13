package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"
	httphelper "github.com/maxcore25/bmstu-it-courses/backend/internal/shared/http"
)

type ScheduleHandler struct {
	service service.ScheduleService
}

func NewScheduleHandler(s service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{service: s}
}

// CreateSchedule godoc
// @Summary Create schedule
// @Tags Schedules
// @Accept json
// @Produce json
// @Param schedule body dto.CreateScheduleRequest true "New schedule"
// @Success 201 {object} dto.ScheduleResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules [post]
func (h *ScheduleHandler) CreateSchedule(c *gin.Context) {
	var req dto.CreateScheduleRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	schedule, err := h.service.CreateSchedule(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewScheduleResponse(schedule)

	c.JSON(http.StatusCreated, resp)
}

// GetSchedule godoc
// @Summary Get schedule by ID
// @Tags Schedules
// @Produce json
// @Param id path string true "Schedule ID (uuid)"
// @Success 200 {object} dto.ScheduleResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /schedules/{id} [get]
func (h *ScheduleHandler) GetSchedule(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	schedule, err := h.service.GetSchedule(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "schedule not found"})
		return
	}

	resp := mapper.NewScheduleResponse(schedule)

	c.JSON(http.StatusOK, resp)
}

// GetAllSchedules godoc
// @Summary Get all schedules
// @Tags Schedules
// @Produce json
// @Success 200 {array} dto.ScheduleResponse
// @Router /schedules [get]
func (h *ScheduleHandler) GetAllSchedules(c *gin.Context) {
	schedules, err := h.service.GetAllSchedules()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]*dto.ScheduleResponse, len(schedules))
	for i, schedule := range schedules {
		resp[i] = mapper.NewScheduleResponse(schedule)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateScheduleByID godoc
// @Summary Update schedule by ID
// @Tags Schedules
// @Accept json
// @Produce json
// @Param id path string true "Schedule ID (uuid)"
// @Param schedule body map[string]interface{} true "Schedule update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [patch]
func (h *ScheduleHandler) UpdateScheduleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	var updateData map[string]any
	if !httphelper.BindJSON(c, &updateData) {
		return
	}
	if err := h.service.UpdateScheduleByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "schedule updated successfully"})
}

// DeleteScheduleByID godoc
// @Summary Delete schedule by ID
// @Tags Schedules
// @Produce json
// @Param id path string true "Schedule ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /schedules/{id} [delete]
func (h *ScheduleHandler) DeleteScheduleByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteScheduleByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
