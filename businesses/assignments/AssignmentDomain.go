package assignments

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	IdAssignment        string       
	Name     			string         	
	Deadline			string     	
	Class				string			
	CreatedAt 			time.Time      	
	UpdatedAt 			time.Time      	
	DeletedAt 			gorm.DeletedAt 	
}

type Usecase interface {
	GetAssignments() []Domain
	Create(assignmentDomain *Domain) Domain
}

type Repository interface {
	GetAssignments() []Domain
	Create(assignmentDomain *Domain) Domain
}