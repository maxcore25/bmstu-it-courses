package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

// RegisterScheduleRoutes registers the schedule routes with the provided gin router group.
func RegisterScheduleRoutes(r *gin.RouterGroup, scheduleService service.ScheduleService, jwtManager *utils.JWTManager) {
	scheduleHandler := NewScheduleHandler(scheduleService)

	scheduleGroup := r.Group("/schedules")
	{
		scheduleGroup.POST("", middleware.AuthMiddleware(jwtManager), scheduleHandler.CreateSchedule)
		scheduleGroup.GET("", scheduleHandler.GetAllSchedules)
		scheduleGroup.GET("/:id", scheduleHandler.GetSchedule)
		scheduleGroup.PATCH("/:id", middleware.AuthMiddleware(jwtManager), scheduleHandler.UpdateScheduleByID)
		scheduleGroup.DELETE("/:id", middleware.AuthMiddleware(jwtManager), scheduleHandler.DeleteScheduleByID)
	}
}
