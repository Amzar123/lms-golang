package modules

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	IdModule      	string          
	Module     		string         	
	TeacherId		string			
	CreatedAt 		time.Time      	
	UpdatedAt 		time.Time      	
	DeletedAt 		gorm.DeletedAt 	
}

type Usecase interface {
	GetModules() []Domain
	Create(moduleDomain *Domain) Domain
	// Login(studentDomain *Domain) string
}

type Repository interface {
	GetModules() []Domain
	Create(moduleDomain *Domain) Domain
	// GetByEmail(studentDomain *Domain) Domain
}