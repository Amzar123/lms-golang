package teachers

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	NIP 		string
	Email 		string
	Password 	string
	Name 		string
	RoleId 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt	gorm.DeletedAt
}

type Usecase interface {
	// GetAllUser() []Domain
	CreateTeacher(teacherDomain *Domain) Domain
	// Login(userDomain *Domain) string
}

type Repository interface {
	// GetAllUser() []Domain
	CreateTeacher(teacherDomain *Domain) Domain
	// GetByEmail(userDomain *Domain) Domain
}