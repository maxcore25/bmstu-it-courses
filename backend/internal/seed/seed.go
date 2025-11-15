package seed

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

func RunSeeds(db *gorm.DB) error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	fmt.Println("Running seeds for environment:", env)

	// Minimal production-safe seeds
	if err := SeedDefaultAdmin(db); err != nil {
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
