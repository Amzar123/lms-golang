package response

import (
	"mini-project/businesses/modules"
	"time"

	"gorm.io/gorm"
)

type Module struct {
	IdModule        	string          `json:"idModule" gorm:"primaryKey"`
	Module     			string         	`json:"module" gorm:"unique" faker:"module"`
	TeacherId			string			`json:"teacherId"`
	CreatedAt 			time.Time      	`json:"created_at"`
	UpdatedAt 			time.Time      	`json:"updated_at"`
	DeletedAt 			gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain modules.Domain) Module {
	return Module{
		IdModule:       domain.IdModule,
		Module:			domain.Module,
		TeacherId:		domain.TeacherId,
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}
