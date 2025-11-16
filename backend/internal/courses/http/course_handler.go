package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/courses/service"
	httphelper "github.com/maxcore25/bmstu-it-courses/backend/internal/shared/http"
)

type CourseHandler struct {
	service service.CourseService
}

func NewCourseHandler(s service.CourseService) *CourseHandler {
	return &CourseHandler{service: s}
}

// CreateCourse godoc
// @Summary Create course
// @Tags Courses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param course body dto.CreateCourseRequest true "New course"
// @Success 201 {object} dto.CourseResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses [post]
func (h *CourseHandler) CreateCourse(c *gin.Context) {
	var req dto.CreateCourseRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	course, err := h.service.CreateCourse(&req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewCourseResponse(course)

	c.JSON(http.StatusCreated, resp)
}

// GetCourse godoc
// @Summary Get course by ID
// @Tags Courses
// @Produce json
// @Param id path string true "Course ID (uuid)"
// @Param expand query []string false "Relations to expand (author). Example: expand=author"
// @Success 200 {object} dto.CourseResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /courses/{id} [get]
func (h *CourseHandler) GetCourse(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	expand := httphelper.ParseExpand(c.QueryArray("expand"))

	course, err := h.service.GetCourse(id, expand)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "course not found"})
		return
	}

	resp := mapper.NewCourseResponse(course)
	c.JSON(http.StatusOK, resp)
}

// GetAllCourses godoc
// @Summary Get all courses
// @Tags Courses
// @Produce json
// @Param expand query []string false "Relations to expand (author). Example: expand=author"
// @Success 200 {array} dto.CourseResponse
// @Router /courses [get]
func (h *CourseHandler) GetAllCourses(c *gin.Context) {
	expand := httphelper.ParseExpand(c.QueryArray("expand"))

	courses, err := h.service.GetAllCourses(expand)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := make([]*dto.CourseResponse, len(courses))
	for i, course := range courses {
		resp[i] = mapper.NewCourseResponse(course)
	}

	c.JSON(http.StatusOK, resp)
}

// UpdateCourseByID godoc
// @Summary Update course by ID
// @Tags Courses
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Course ID (uuid)"
// @Param course body map[string]interface{} true "Course update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [patch]
func (h *CourseHandler) UpdateCourseByID(c *gin.Context) {
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
	if err := h.service.UpdateCourseByID(id, updateData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course updated successfully"})
}

// DeleteCourseByID godoc
// @Summary Delete course by ID
// @Tags Courses
// @Produce json
// @Security BearerAuth
// @Param id path string true "Course ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /courses/{id} [delete]
func (h *CourseHandler) DeleteCourseByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteCourseByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
