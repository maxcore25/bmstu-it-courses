package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/mapper"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"
	httphelper "github.com/maxcore25/bmstu-it-courses/backend/internal/shared/http"
)

type BranchHandler struct {
	service service.BranchService
}

func NewBranchHandler(s service.BranchService) *BranchHandler {
	return &BranchHandler{service: s}
}

// CreateBranch godoc
// @Summary Create branch
// @Tags Branches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param branch body dto.CreateBranchRequest true "New branch"
// @Success 201 {object} dto.BranchResponse
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branches [post]
func (h *BranchHandler) CreateBranch(c *gin.Context) {
	var req dto.CreateBranchRequest

	if !httphelper.BindJSON(c, &req) {
		return
	}

	// Manually map CreateBranchRequest to the model/dto expected by service
	newBranch := dto.CreateBranchRequest{
		Address: req.Address,
		Rooms:   req.Rooms,
	}

	branch, err := h.service.CreateBranch(newBranch)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	resp := mapper.NewBranchResponse(branch)

	c.JSON(http.StatusCreated, resp)
}

// GetBranch godoc
// @Summary Get branch by ID
// @Tags Branches
// @Produce json
// @Param id path string true "Branch ID (uuid)"
// @Success 200 {object} dto.BranchResponse
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Router /branches/{id} [get]
func (h *BranchHandler) GetBranch(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	branch, err := h.service.GetBranch(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "branch not found"})
		return
	}

	resp := mapper.NewBranchResponse(branch)

	c.JSON(http.StatusOK, resp)
}

// GetAllBranches godoc
// @Summary Get all branches
// @Tags Branches
// @Produce json
// @Success 200 {array} dto.BranchResponse
// @Router /branches [get]
func (h *BranchHandler) GetAllBranches(c *gin.Context) {
	branches, err := h.service.GetAllBranches()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	resp := make([]*dto.BranchResponse, len(branches))
	for i, branch := range branches {
		resp[i] = mapper.NewBranchResponse(branch)
	}
	c.JSON(http.StatusOK, resp)
}

// UpdateBranchByID godoc
// @Summary Update branch by ID
// @Tags Branches
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path string true "Branch ID (uuid)"
// @Param branch body dto.UpdateBranchRequest true "Branch update data"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branches/{id} [patch]
func (h *BranchHandler) UpdateBranchByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	var req dto.UpdateBranchRequest
	if !httphelper.BindJSON(c, &req) {
		return
	}

	if err := h.service.UpdateBranchByID(id, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "branch updated successfully"})
}

// DeleteBranchByID godoc
// @Summary Delete branch by ID
// @Tags Branches
// @Produce json
// @Security BearerAuth
// @Param id path string true "Branch ID (uuid)"
// @Success 204 {object} nil
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /branches/{id} [delete]
func (h *BranchHandler) DeleteBranchByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}
	if err := h.service.DeleteBranchByID(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
