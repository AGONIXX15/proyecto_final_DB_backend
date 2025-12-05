package admin

import (
	"github.com/AGONIXX15/db_proyecto_final/internal/utils"
)

type AdminService struct {
	Repo AdminRepository
}

func NewAdminService(repo AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) CreateAdmin(admin *Admin) error {
	_, err := s.Repo.FindByUsername(admin.Username)
	if err != nil {
		return err
	}

	hashedPassword := utils.HashPassword(admin.Password)
	admin.Password = string(hashedPassword)
	return s.Repo.Create(admin)
}

func (s *AdminService) GetAllAdmins() ([]Admin, error) {
	return s.Repo.FindAll()
}

func (s *AdminService) GetAdmin(id uint) (*Admin, error) {
	return s.Repo.FindByID(id)
}
func (s *AdminService) GetByUsernameAdmin(username string) (*Admin, error) {
	return s.Repo.FindByUsername(username)
}

func (s *AdminService) DeleteAdmin(id uint) error {
	return s.Repo.Delete(id)
}

func (s *AdminService) UpdateAdmin(admin *Admin) error {

	hashedPassword := utils.HashPassword(admin.Password)
	admin.Password = string(hashedPassword)
	return s.Repo.Update(admin)
}

func (s *AdminService) LoginAdmin(username string, password string) error {
	admin, err := s.Repo.FindByUsername(username)
	if err != nil {
		return err
	}

	if admin.Password == utils.HashPassword(password) {
		return nil
	}

	return ErrPasswordWrong
}

