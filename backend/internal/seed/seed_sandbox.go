package seed

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

	"github.com/maxcore25/bmstu-it-courses/backend/internal/shared/utils"
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
		fmt.Println("ℹ️  Sandbox data already seeded — skipping")
		return nil
	}

	// Shared password
	hashedPass, err := utils.HashPassword("qwe123")
	if err != nil {
		return fmt.Errorf("❌ failed to hash password: %w", err)
	}

	// ---------- FACTORIES ----------

	tutorNames := []struct{ First, Last string }{
		{"Иван", "Петров"},
		{"Алексей", "Сидоров"},
		{"Максим", "Андреев"},
		{"Дмитрий", "Козлов"},
		{"Сергей", "Федоров"},
		{"Николай", "Ильин"},
	}

	branchAddresses := []string{
		"Москва, ул. Ленина, 10",
		"Москва, пр-т Мира, 45",
		"Москва, ул. Тверская, 22",
	}

	courseNames := []string{
		"Основы Go",
		"Веб-разработка",
		"Продвинутый Python",
		"Основы Linux",
		"Основы баз данных",
		"Алгоритмы и структуры данных",
	}

	// Factory: tutor user
	newTutor := func(i int) *authModel.User {
		return &authModel.User{
			ID:             uuid.New(),
			FirstName:      tutorNames[i].First,
			LastName:       tutorNames[i].Last,
			Email:          fmt.Sprintf("tutor%d@mail.ru", i+1),
			Password:       hashedPass,
			Role:           authModel.RoleTutor,
			KnowledgeLevel: authModel.KnowledgeLevelAdvanced,
			Rating:         floatPtr(4.5 + 0.1*float64(i)),
			Portfolio:      strPtr("Опыт преподавания и коммерческой разработки более 5 лет. Занимался построением систем от проектирования до деплоя в продакшн."),
		}
	}

	// Factory: client user
	newClient := func(i int) *authModel.User {
		return &authModel.User{
			ID:             uuid.New(),
			FirstName:      "Клиент",
			LastName:       fmt.Sprintf("Номер%d", i+1),
			Email:          fmt.Sprintf("client%d@mail.ru", i+1),
			Password:       hashedPass,
			Role:           authModel.RoleClient,
			KnowledgeLevel: authModel.KnowledgeLevelBeginner,
		}
	}

	// Factory: branch
	newBranch := func(i int) *branchModel.Branch {
		return &branchModel.Branch{
			ID:      uuid.New(),
			Address: branchAddresses[i],
			Rooms:   3 + i,
		}
	}

	// Factory: course (1 tutor → 1 course)
	newCourse := func(i int, tutorID uuid.UUID) *courseModel.Course {
		return &courseModel.Course{
			ID:         uuid.New(),
			Name:       courseNames[i],
			Difficulty: authModel.KnowledgeLevelBeginner,
			Duration:   "8 недель",
			Price:      35000 + int64(i)*5000,
			Format:     courseModel.CourseFormatGroup,
			AuthorID:   tutorID,
		}
	}

	// Factory: schedule (1 course → 1 schedule)
	newSchedule := func(courseID uuid.UUID, branchID uuid.UUID) *scheduleModel.Schedule {
		return &scheduleModel.Schedule{
			CourseID: courseID,
			BranchID: &branchID,
			StartAt:  time.Now().AddDate(0, 0, 3),
			EndAt:    time.Now().AddDate(0, 0, 3+60),
			Capacity: 15,
			Reserved: 0,
		}
	}

	// ---------- INSERT DATA ----------

	// Branches (3)
	branchIDs := make([]uuid.UUID, 3)
	for i := range 3 {
		b := newBranch(i)
		if err := brRepo.Create(b); err != nil {
			return fmt.Errorf("❌ failed to seed branch: %w", err)
		}
		branchIDs[i] = b.ID
	}

	// Tutors (6)
	tutorIDs := make([]uuid.UUID, 6)
	for i := range 6 {
		t := newTutor(i)
		if err := usrRepo.Create(t); err != nil {
			return fmt.Errorf("❌ failed to seed tutor: %w", err)
		}
		tutorIDs[i] = t.ID
	}

	// One Client
	client := newClient(0)
	if err := usrRepo.Create(client); err != nil {
		return fmt.Errorf("❌ failed to seed client: %w", err)
	}

	// Courses (6)
	courseIDs := make([]uuid.UUID, 6)
	for i := range 6 {
		c := newCourse(i, tutorIDs[i])
		if err := crsRepo.Create(c); err != nil {
			return fmt.Errorf("❌ failed to seed course: %w", err)
		}
		courseIDs[i] = c.ID
	}

	// Schedules (6)
	for i := range 6 {
		// distribute across 3 branches
		branchID := branchIDs[i%3]

		s := newSchedule(courseIDs[i], branchID)
		if err := schRepo.Create(s); err != nil {
			return fmt.Errorf("❌ failed to seed schedule: %w", err)
		}
	}

	fmt.Println("✅ Sandbox data seeded successfully")
	return nil
}

func floatPtr(v float64) *float64 { return &v }
func strPtr(v string) *string     { return &v }
