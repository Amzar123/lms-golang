package students

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	NISN        string
	Email     	string
	Password  	string
	Name		string
	Class		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	gorm.DeletedAt
}

type Usecase interface {
	// GetAllUser() []Domain
	Create(userDomain *Domain) Domain
	// Login(userDomain *Domain) string
}

type Repository interface {
	// GetAllUser() []Domain
	Create(userDomain *Domain) Domain
	// GetByEmail(userDomain *Domain) Domain
}