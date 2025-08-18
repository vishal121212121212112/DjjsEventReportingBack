package userServiceHandler

import (
	"errors"
	database "event-reporting/app/database/pgsql/repository"
	"event-reporting/app/dtos"
	"event-reporting/app/models"
	branchServiceHandler "event-reporting/app/services/branch"
	"event-reporting/app/utils/constants"
	"event-reporting/app/utils/hashing"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type UserService struct {
	repo *database.Repository
}

func (s *UserService) CreateUser(creator models.User, req models.UserCreateRequest) (string, error) {
	var newUser models.User

	// Permission check
	switch creator.Type {
	case constants.UserTypeAdmin:
		if req.Type != constants.UserTypeBranchCoordinator && req.Type != constants.UserTypeBranchAssitant &&
			req.Type != constants.UserTypeAdmin && req.Type != constants.UserTypeITAssistant {
			return "", errors.New("admin can only create branch coordinators, assistants, or IT assistants")
		}
	case constants.UserTypeBranchCoordinator:
		if req.Type != constants.UserTypeBranchAssitant {
			return "", errors.New("branch coordinator can only create branch assistants")
		}
	default:
		return "", errors.New("permission denied")
	}

	// Check if email already exists
	conditions := map[string]interface{}{"email": req.Email}
	var existingUser models.User
	if err := s.repo.Find(&existingUser, conditions); err == nil {
		return "", errors.New("email already exists")
	}

	// Check if username already exists
	if req.Username != "" {
		conditions = map[string]interface{}{"username": req.Username}
		if err := s.repo.Find(&existingUser, conditions); err == nil {
			return "", errors.New("username already exists")
		}
	}

	newUser.ID = uuid.New()
	newUser.Email = req.Email
	newUser.Username = req.Username
	newUser.Type = req.Type
	newUser.CreatedOn = time.Now().Format(time.RFC3339)
	newUser.UpdatedOn = time.Now().Format(time.RFC3339)
	newUser.CreatedBy = creator.ID
	newUser.UpdatedBy = creator.ID

	passwordToUse := req.Password
	if passwordToUse == "" {
		return "", errors.New("password is required")
	}

	hashedPassword, err := hashing.HashData(passwordToUse)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}
	newUser.Password = string(hashedPassword)
	if err := s.repo.Create(&newUser); err != nil {
		return "", err
	}

	if req.Type == constants.UserTypeBranchCoordinator {
		branchReq := dtos.CreateBranchWithUserEmail{
			Email:     req.Email,
			CreatedBy: creator.ID.String(),
			UpdatedBy: newUser.ID.String(),
		}

		branchService := branchServiceHandler.NewBranchService(s.repo)
		if err := branchService.CreateBranchWithUserEmail(branchReq); err != nil {
			return "", fmt.Errorf("failed to create branch for branch coordinator: %w", err)
		}
	}

	return newUser.ID.String(), nil
}

func (s *UserService) GetUserByEmail(email string, user *models.User) error {
	if email == "" {
		return errors.New("email is required")
	}
	conditions := map[string]interface{}{"email": email}
	err := s.repo.Find(user, conditions)
	if err != nil {
		return err
	}
	return nil
}

// Add new method to get user by username
func (s *UserService) GetUserByUsername(username string, user *models.User) error {
	if username == "" {
		return errors.New("username is required")
	}
	conditions := map[string]interface{}{"username": username}
	err := s.repo.Find(user, conditions)
	if err != nil {
		return err
	}
	return nil
}

// Add new method to get user by email OR username
func (s *UserService) GetUserByIdentifier(identifier string, user *models.User) error {
	if identifier == "" {
		return errors.New("identifier is required")
	}

	// Try to find by email first
	conditions := map[string]interface{}{"email": identifier}
	err := s.repo.Find(user, conditions)
	if err == nil {
		return nil // Found by email
	}

	// If not found by email, try username
	conditions = map[string]interface{}{"username": identifier}
	err = s.repo.Find(user, conditions)
	if err != nil {
		return errors.New("user not found")
	}

	return nil
}

func (s *UserService) GetUserByID(userID string, user *models.User) error {
	conditions := map[string]interface{}{"id": userID}
	return s.repo.Find(user, conditions)
}

func (s *UserService) UpdateUserToken(userID uuid.UUID, token string, additionalUpdates map[string]interface{}) error {
	conditions := map[string]interface{}{"id": userID}
	updates := map[string]interface{}{"token": token}

	// Merge additional updates
	for key, value := range additionalUpdates {
		updates[key] = value
	}

	return s.repo.UpdateFields(&models.User{}, conditions, updates)
}

func (s *UserService) UpdateUserLoginInfo(userID uuid.UUID, updates map[string]interface{}) error {
	conditions := map[string]interface{}{"id": userID}
	return s.repo.UpdateFields(&models.User{}, conditions, updates)
}
