package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterCourseRoutes(r *gin.RouterGroup, courseService service.CourseService, jwtManager *utils.JWTManager) {
	courseHandler := NewCourseHandler(courseService)

	courseGroup := r.Group("/courses")
	courseGroup.GET("", courseHandler.GetAllCourses)
	courseGroup.GET("/:id", courseHandler.GetCourse)

	protected := courseGroup.Group("")
	protected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		protected.POST("", courseHandler.CreateCourse)
		protected.PATCH("/:id", courseHandler.UpdateCourseByID)
		protected.DELETE("/:id", courseHandler.DeleteCourseByID)
	}
}
