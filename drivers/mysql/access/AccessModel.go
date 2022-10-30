package access

import (
	"mini-project/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Access struct {
	IdModule        	string          `json:"idModule"`
	Nisn     			string         	`json:"nisn" faker:"role"`
	CreatedAt 			time.Time      	`json:"created_at"`
	UpdatedAt 			time.Time      	`json:"updated_at"`
	DeletedAt 			gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain *access.Domain) *Access {
	return &Access{
		IdModule:       domain.IdModule,
		Nisn:			domain.Nisn,
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}

func (rec *Access) ToDomain() access.Domain {
	return access.Domain{
		IdModule:       rec.IdModule,
		Nisn:			rec.Nisn,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
		DeletedAt: 		rec.DeletedAt,
	}
}
