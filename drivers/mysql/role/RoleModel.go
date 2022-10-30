package role

import (
	"mini-project/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Role struct {
	IdRole        	string          `json:"idRole" gorm:"primaryKey"`
	Role     		string         	`json:"role" gorm:"unique" faker:"role"`
	CreatedAt 		time.Time      	`json:"created_at"`
	UpdatedAt 		time.Time      	`json:"updated_at"`
	DeletedAt 		gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain *roles.Domain) *Role {
	return &Role{
		IdRole:        	domain.IdRole,
		Role:			domain.Role
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}

func (rec *User) ToDomain() roles.Domain {
	return roles.Domain{
		IdRole:        	rec.IdRole,
		Role:			rec.Role
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
		DeletedAt: 		rec.DeletedAt,
	}
}
