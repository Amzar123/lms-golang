package response

import (
	"mini-project/businesses/assignments"
	"time"

	"gorm.io/gorm"
)

type Assignment struct {
	IdAssignment        string          `json:"idAssignment" gorm:"primaryKey"`
	Name     			string         	`json:"name" faker:"name"`
	Deadline			string      	`json:"deadline"`
	Class				string			`json:"class" faker:"class"`
	CreatedAt 			time.Time      	`json:"created_at"`
	UpdatedAt 			time.Time      	`json:"updated_at"`
	DeletedAt 			gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain assignments.Domain) Assignment {
	return Assignment{
		IdAssignment:		domain.IdAssignment,     
		Name:				domain.Name,
		Deadline:			domain.Deadline,
		Class:				domain.Class,
		CreatedAt:			domain.CreatedAt,
		UpdatedAt:			domain.UpdatedAt,
		DeletedAt:			domain.DeletedAt,
	}
}
