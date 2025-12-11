package admin

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*Admin, error)
	FindByID(id uint) (*Admin, error)
	FindAll() ([]Admin, error)
	Create(admin *Admin) error
	Update(admin *Admin) error
	Delete(id uint) error
	UpdatePartial(id uint, updates map[string]interface{}) error
}
type adminRepository struct {
	db *gorm.DB
}

var (
	ErrDBInternal = errors.New("error interno de la base de datos")
	ErrNotFound = errors.New("admin no encontrado")
	ErrPasswordWrong = errors.New("contraseña incorrecta")
)

// Create implements AdminRepository.
func (r *adminRepository) Create(admin *Admin) error {
	result := r.db.Create(admin)
	if result.Error != nil {
		return fmt.Errorf("%w: no se pudo crear el admin: %s", ErrDBInternal, result.Error)
	}


	return nil
}

// Delete implements AdminRepository.
func (r *adminRepository) Delete(id uint) error {
	result := r.db.Delete(&Admin{}, id)
	if result.Error != nil {
		return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: admin con id %d", ErrNotFound, id)
	}
	
	return nil
}

// FindAll implements AdminRepository.
func (r *adminRepository) FindAll() ([]Admin, error) {
	var admins []Admin
	result := r.db.Find(&admins)
	if result.Error != nil {
		return nil,fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}
	return admins, nil
}

func (r *adminRepository) FindByUsername(username string) (*Admin, error) {
	var admin Admin
	result := r.db.Where("username = ?", username).First(&admin)
	
	if result.Error != nil {
		return nil, fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%w: no se encontro el admin con nombre %s",ErrNotFound, username)
	}

	return &admin, nil
}

// Update implements AdminRepository.
func (r *adminRepository) Update(admin *Admin) error {
	result := r.db.Model(&Admin{}).Where("ID = ?", admin.ID).Updates(admin)

	if result.Error != nil {
		return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("%w: cliente con documento %d", ErrNotFound,admin.ID)
	}
	return nil
}

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepository{db: db}
}

func (r *adminRepository) FindByID(id uint) (*Admin, error) {	
	var admin *Admin

	result := r.db.First(&admin,id)
	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%w: admin con id %d no encontrado", ErrNotFound, id)
	}

	if result.Error != nil {
		return nil, fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
	}
	return admin,nil
}

func (r *adminRepository) UpdatePartial(id uint, updates map[string]interface{}) error {
	fmt.Println("llegas aqui??? %d")
	fmt.Printf("Updates ENVIADOS: %#v\n", updates)
    result := r.db.Model(&Admin{}).Where("id = ?", id).Updates(updates)

	fmt.Println("SQL:", result.Statement.SQL.String())

    if result.Error != nil {
			fmt.Println("fallo actualizando")
        return fmt.Errorf("%w: %s", ErrDBInternal, result.Error)
    }

    if result.RowsAffected == 0 {
			fmt.Println("fallo actualizando")
        return fmt.Errorf("%w: admin con id %d no encontrado", ErrNotFound, id)
    }
		fmt.Println("se supone que actualizamos")
    return nil
}

