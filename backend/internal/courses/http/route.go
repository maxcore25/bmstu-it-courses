package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/service"
)

// RegisterCourseRoutes registers the course routes with the provided gin router group.
func RegisterCourseRoutes(r *gin.RouterGroup, courseService service.CourseService) {
	courseHandler := NewCourseHandler(courseService)

	courseGroup := r.Group("/courses")
	{
		courseGroup.POST("", courseHandler.CreateCourse)
		courseGroup.GET("", courseHandler.GetAllCourses)
		courseGroup.GET("/:id", courseHandler.GetCourse)
		courseGroup.PATCH("/:id", courseHandler.UpdateCourseByID)
		courseGroup.DELETE("/:id", courseHandler.DeleteCourseByID)
	}
}
