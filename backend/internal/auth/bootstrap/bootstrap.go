package bootstrap

import (
	"fmt"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
)

// SeedDefaultAdmin checks if an admin user exists and creates it if missing.
func SeedDefaultAdmin(userRepo repository.UserRepository) error {
	const defaultEmail = "admin@mail.ru"
	const defaultPassword = "qwe123"

	admin, _ := userRepo.GetByEmail(defaultEmail)
	if admin != nil {
		fmt.Println("✅ Admin user already exists:", admin.Email)
		return nil
	}

	hashedPassword, err := utils.HashPassword(defaultPassword)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	adminUser := &model.User{
		FirstName:      "Admin",
		LastName:       "User",
		Email:          defaultEmail,
		Password:       hashedPassword,
		Role:           model.RoleAdmin,
		KnowledgeLevel: model.KnowledgeLevelAdvanced,
	}

	if err := userRepo.Create(adminUser); err != nil {
		return fmt.Errorf("failed to create default admin: %w", err)
	}

	fmt.Println("✅ Default admin created:", adminUser.Email)
	return nil
}
