package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"
)

func RegisterBranchRoutes(r *gin.RouterGroup, branchService service.BranchService) {
	branchHandler := NewBranchHandler(branchService)

	branchGroup := r.Group("/branches")
	{
		branchGroup.POST("", branchHandler.CreateBranch)
		branchGroup.GET("", branchHandler.GetAllBranches)
		branchGroup.GET("/:id", branchHandler.GetBranch)
		branchGroup.PATCH("/:id", branchHandler.UpdateBranchByID)
		branchGroup.DELETE("/:id", branchHandler.DeleteBranchByID)
	}
}
