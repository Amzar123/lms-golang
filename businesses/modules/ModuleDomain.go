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
	GetByID(id string) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}

type Repository interface {
	GetModules() []Domain
	Create(moduleDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
}