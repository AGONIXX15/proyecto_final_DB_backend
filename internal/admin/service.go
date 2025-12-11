package admin

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type AdminService struct {
	Repo AdminRepository
}

func NewAdminService(repo AdminRepository) *AdminService {
	return &AdminService{Repo: repo}
}

func (s *AdminService) CreateAdmin(admin *Admin) error {
	_, err := s.Repo.FindByUsername(admin.Username)
	if err == nil {
		return fmt.Errorf("usuario ya existia")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
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
    existingAdmin, err := s.GetByUsernameAdmin(admin.Username) // Devuelve todo el admin
    if err != nil {
        return fmt.Errorf("no se encontro el usuario a actualizar")
    }

    if admin.Password != "" && admin.Password != existingAdmin.Password {
        hashed, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
        if err != nil {
            return fmt.Errorf("error al hashear la contraseña")
        }
        admin.Password = string(hashed)
    } else {
        admin.Password = existingAdmin.Password
    }

    return s.Repo.Update(admin)
}


func (s *AdminService) LoginAdmin(username, password string) (*Admin, error) {
    admin, err := s.Repo.FindByUsername(username)
    if err != nil {
			fmt.Println("usuario con ese nombre no encontrao")
        return nil, err
    }

    // Verificar contraseña usando bcrypt
    if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(password)); err != nil {
			fmt.Println("contraseñas diferentes")
        return nil, ErrPasswordWrong
    }

    return admin, nil
}

func (s *AdminService) UpdateAdminPartial(id uint, updates map[string]interface{}) error {
    return s.Repo.UpdatePartial(id, updates)
}


