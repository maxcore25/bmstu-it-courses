package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

// RegisterCourseRoutes registers the course routes with the provided gin router group.
func RegisterCourseRoutes(r *gin.RouterGroup, courseService service.CourseService, jwtManager *utils.JWTManager) {
	courseHandler := NewCourseHandler(courseService)

	courseGroup := r.Group("/courses")
	{
		courseGroup.POST("", middleware.AuthMiddleware(jwtManager), courseHandler.CreateCourse)
		courseGroup.GET("", courseHandler.GetAllCourses)
		courseGroup.GET("/:id", courseHandler.GetCourse)
		courseGroup.PATCH("/:id", middleware.AuthMiddleware(jwtManager), courseHandler.UpdateCourseByID)
		courseGroup.DELETE("/:id", middleware.AuthMiddleware(jwtManager), courseHandler.DeleteCourseByID)
	}
}
