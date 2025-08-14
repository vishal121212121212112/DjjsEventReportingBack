package admin

import (
	"errors"
	"event-reporting/app/models"
	"event-reporting/app/database/pgsql/repository"
	"event-reporting/app/utils/hashing"
)

type AdminService interface {
	RegisterAdmin(admin *models.Admin) (*models.Admin, error)
	AdminLogin(username, password string) (*models.Admin, error)
	UpdateAdminPassword(oldPassword, newPassword string) error
}

type adminService struct {
	repo repository.AdminRepository
}

func NewAdminService(r repository.AdminRepository) AdminService {
	return &adminService{repo: r}
}

func (s *adminService) RegisterAdmin(admin *models.Admin) (*models.Admin, error) {
	// Hash the password
	hashedPassword, err := hashing.Hash(admin.Password)
	if err != nil {
		return nil, err
	}
	admin.Password = hashedPassword

	// Save admin to the database
	err = s.repo.CreateAdmin(admin)
	if err != nil {
		return nil, err
	}

	return admin, nil
}

func (s *adminService) AdminLogin(username, password string) (*models.Admin, error) {
	// Find admin by username
	admin, err := s.repo.FindAdminByUsername(username)
	if err != nil {
		return nil, err
	}

	// Verify password
	if err := hashing.Compare(admin.Password, password); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return admin, nil
}

func (s *adminService) UpdateAdminPassword(oldPassword, newPassword string) error {
	// Here you would implement logic to update password for admin.
	// For example, you can get the current admin, check the old password, and update with new one.
	return nil
}
