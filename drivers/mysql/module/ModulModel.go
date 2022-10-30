package module

import (
	"mini-project/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Module struct {
	IdModule        	string          `json:"idModule" gorm:"primaryKey"`
	Module     			string         	`json:"role" gorm:"unique" faker:"role"`
	TeacherId			string			`json:"teacherId"`
	CreatedAt 			time.Time      	`json:"created_at"`
	UpdatedAt 			time.Time      	`json:"updated_at"`
	DeletedAt 			gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain *roles.Domain) *Role {
	return &Role{
		IdRole:        	domain.IdRole,
		Module:			domain.Module,
		TeacherId:		domain.TeacherId,
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}

func (rec *User) ToDomain() roles.Domain {
	return roles.Domain{
		IdRole:        	rec.IdRole,
		Module:			rec.Module,
		TeacherId:		rec.TeacherId,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
		DeletedAt: 		rec.DeletedAt,
	}
}
