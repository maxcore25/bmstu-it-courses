package bootstrap

import (
	"fmt"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/dto"
	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/service"
)

// SeedDefaultAdmin checks if an admin user exists and creates it if missing.
func SeedDefaultAdmin(userService service.UserService) error {
	const defaultEmail = "admin@mail.ru"
	const defaultPassword = "qwe123"

	// Check if the admin already exists
	admin, err := userService.GetByEmail(defaultEmail)
	if err == nil && admin != nil {
		fmt.Println("✅ Admin user already exists:", admin.Email)
		return nil
	}

	// Create the admin user if not found
	req := dto.CreateUserRequest{
		FirstName:      "Admin",
		LastName:       "User",
		Email:          defaultEmail,
		Password:       defaultPassword,
		KnowledgeLevel: "advanced",
	}

	newAdmin, err := userService.CreateUser(req)
	if err != nil {
		return fmt.Errorf("failed to create default admin: %w", err)
	}

	fmt.Println("✅ Default admin created:", newAdmin.Email)
	return nil
}
