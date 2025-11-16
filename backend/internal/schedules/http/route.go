package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterScheduleRoutes(r *gin.RouterGroup, scheduleService service.ScheduleService, jwtManager *utils.JWTManager) {
	scheduleHandler := NewScheduleHandler(scheduleService)

	scheduleGroup := r.Group("/schedules")
	scheduleGroup.GET("", scheduleHandler.GetAllSchedules)
	scheduleGroup.GET("/:id", scheduleHandler.GetSchedule)

	protected := scheduleGroup.Group("")
	protected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		protected.POST("", scheduleHandler.CreateSchedule)
		protected.PATCH("/:id", scheduleHandler.UpdateScheduleByID)
		protected.DELETE("/:id", scheduleHandler.DeleteScheduleByID)
	}
}
