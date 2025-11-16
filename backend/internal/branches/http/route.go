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
	branchGroup.GET("", branchHandler.GetAllBranches)
	branchGroup.GET("/:id", branchHandler.GetBranch)

	protected := branchGroup.Group("")
	protected.Use(
		middleware.AuthMiddleware(jwtManager),
		middleware.RoleMiddleware("admin"),
	)
	{
		protected.POST("", branchHandler.CreateBranch)
		protected.PATCH("/:id", branchHandler.UpdateBranchByID)
		protected.DELETE("/:id", branchHandler.DeleteBranchByID)
	}
}
