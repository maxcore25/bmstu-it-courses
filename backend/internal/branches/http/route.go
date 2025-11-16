package http

import (
	"github.com/gin-gonic/gin"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/middleware"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/branches/service"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

func RegisterBranchRoutes(r *gin.RouterGroup, branchService service.BranchService, jwtManager *utils.JWTManager) {
	branchHandler := NewBranchHandler(branchService)

	branchGroup := r.Group("/branches")
	{
		branchGroup.POST("", middleware.AuthMiddleware(jwtManager), branchHandler.CreateBranch)
		branchGroup.GET("", branchHandler.GetAllBranches)
		branchGroup.GET("/:id", branchHandler.GetBranch)
		branchGroup.PATCH("/:id", middleware.AuthMiddleware(jwtManager), branchHandler.UpdateBranchByID)
		branchGroup.DELETE("/:id", middleware.AuthMiddleware(jwtManager), branchHandler.DeleteBranchByID)
	}
}
