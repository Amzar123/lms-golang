package teacher

import (
	"mini-project/businesses/teachers"
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	NIP        	string          `json:"nip" gorm:"primaryKey"`
	Email     	string         	`json:"email" gorm:"unique" faker:"email"`
	Password  	string         	`json:"password" faker:"password"`
	Name		string			`json:"name" faker:"name"`
	RoleId		string			`json:"roleId"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain *teachers.Domain) *Teacher {
	return &Teacher{
		NIP:        	domain.NIP,
		Email:     		domain.Email,
		Password:  		domain.Password,
		Name:			domain.Name,
		RoleId:			domain.RoleId,
		CreatedAt: 		domain.CreatedAt,
		UpdatedAt: 		domain.UpdatedAt,
		DeletedAt: 		domain.DeletedAt,
	}
}

func (rec *Teacher) ToDomain() teachers.Domain {
	return teachers.Domain{
		NIP:        	rec.NIP,
		Email:     		rec.Email,
		Password:  		rec.Password,
		Name:			rec.Name,
		RoleId:			rec.RoleId,
		CreatedAt: 		rec.CreatedAt,
		UpdatedAt: 		rec.UpdatedAt,
		DeletedAt: 		rec.DeletedAt,
	}
}
