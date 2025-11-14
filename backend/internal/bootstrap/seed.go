package bootstrap

import (
	"fmt"
	"os"

	"github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"
	"gorm.io/gorm"
)

func RunSeeds(db *gorm.DB) error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Println("Running seeds for environment:", env)

	// Minimal production-safe seeds
	if err := SeedDefaultAdmin(repository.NewUserRepository(db)); err != nil {
		return err
	}

	// Local / Dev seeders
	if env == "development" || env == "local" {
		if err := SeedSandboxData(db); err != nil {
			return err
		}
	}

	return nil
}
