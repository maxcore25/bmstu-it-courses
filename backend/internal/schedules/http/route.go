package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"
)

// RegisterScheduleRoutes registers the schedule routes with the provided gin router group.
func RegisterScheduleRoutes(r *gin.RouterGroup, scheduleService service.ScheduleService) {
	scheduleHandler := NewScheduleHandler(scheduleService)

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.POST("", scheduleHandler.CreateSchedule)
		scheduleGroup.GET("", scheduleHandler.GetAllSchedules)
		scheduleGroup.GET("/:id", scheduleHandler.GetSchedule)
		scheduleGroup.PATCH("/:id", scheduleHandler.UpdateScheduleByID)
		scheduleGroup.DELETE("/:id", scheduleHandler.DeleteScheduleByID)
	}
}
