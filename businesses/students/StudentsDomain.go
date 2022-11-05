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
	// GetModules() []Domain
	Create(studentDomain *Domain) Domain
	Login(studentDomain *Domain) string
}

type Repository interface {
	// GetModules() []Domain
	Create(studentDomain *Domain) Domain
	GetByEmail(studentDomain *Domain) Domain
}