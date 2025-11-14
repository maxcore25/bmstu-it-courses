package bootstrap

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"

	authModel "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/model"
	userRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/auth/repository"

	branchModel "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/model"
	branchRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/branches/repository"

	courseModel "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/model"
	courseRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/courses/repository"

	scheduleModel "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/model"
	scheduleRepo "github.com/maxcore25/bmstu-it-courses/backend/internal/schedules/repository"
)

func SeedSandboxData(db *gorm.DB) error {
	fmt.Println("🔄 Seeding sandbox development data...")

	// Instantiate repositories
	brRepo := branchRepo.NewBranchRepository(db)
	usrRepo := userRepo.NewUserRepository(db)
	crsRepo := courseRepo.NewCourseRepository(db)
	schRepo := scheduleRepo.NewScheduleRepository(db)

	// Check if sandbox data already exists to avoid reseeding
	branches, err := brRepo.GetAll()
	if err != nil {
		return fmt.Errorf("❌ failed to check existing branches: %w", err)
	}
	if len(branches) > 0 {
		fmt.Println("ℹ️  Sandbox data already seeded - skipping")
		return nil
	}

	// --- Branch ---
	branchID := uuid.New()
	branch := &branchModel.Branch{
		ID:      branchID,
		Address: "Main Campus, 1st Floor",
		Rooms:   5,
	}
	if err := brRepo.Create(branch); err != nil {
		return fmt.Errorf("failed to seed branch: %w", err)
	}

	// --- Tutor ---
	tutorID := uuid.New()
	tutor := &authModel.User{
		ID:             tutorID,
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "tutor@example.com",
		Password:       "$2a$10$xxxxxx", // bcrypt placeholder
		Role:           authModel.RoleTutor,
		KnowledgeLevel: authModel.KnowledgeLevelAdvanced,
		Rating:         floatPtr(4.9),
		Portfolio:      strPtr("10 years in backend engineering..."),
	}
	if err := usrRepo.Create(tutor); err != nil {
		return fmt.Errorf("failed to seed tutor: %w", err)
	}

	// --- Client ---
	clientID := uuid.New()
	client := &authModel.User{
		ID:             clientID,
		FirstName:      "Alice",
		LastName:       "Smith",
		Email:          "client@example.com",
		Password:       "$2a$10$xxxxxx",
		Role:           authModel.RoleClient,
		KnowledgeLevel: authModel.KnowledgeLevelBeginner,
	}
	if err := usrRepo.Create(client); err != nil {
		return fmt.Errorf("failed to seed client: %w", err)
	}

	// --- Course ---
	courseID := uuid.New()
	course := &courseModel.Course{
		ID:         courseID,
		Name:       "Go Backend Development",
		Difficulty: authModel.KnowledgeLevelBeginner,
		Duration:   "10 weeks",
		Price:      49000,
		Format:     courseModel.CourseFormatGroup,
		AuthorID:   tutorID,
	}
	if err := crsRepo.Create(course); err != nil {
		return fmt.Errorf("failed to seed course: %w", err)
	}

	// --- Schedule ---
	schedule := &scheduleModel.Schedule{
		CourseID: courseID,
		BranchID: &branchID,
		StartAt:  time.Now().AddDate(0, 0, 7),
		EndAt:    time.Now().AddDate(0, 0, 7+84),
		Capacity: 20,
		Reserved: 2,
	}
	if err := schRepo.Create(schedule); err != nil {
		return fmt.Errorf("failed to seed schedule: %w", err)
	}

	fmt.Println("✅ Sandbox data seeded")
	return nil
}

func floatPtr(v float64) *float64 { return &v }
func strPtr(v string) *string     { return &v }
