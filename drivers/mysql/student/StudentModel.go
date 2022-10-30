package student

import (
	"mini-project/businesses/users"
	"time"

	"gorm.io/gorm"
)

type Student struct {
	NISN        string          `json:"nisn" gorm:"primaryKey"`
	Email     	string         	`json:"email" gorm:"unique" faker:"email"`
	Password  	string         	`json:"password" faker:"password"`
	Name		string			`json:"name" faker:"name"`
	Class		string			`json:"class"`
	CreatedAt 	time.Time      	`json:"created_at"`
	UpdatedAt 	time.Time      	`json:"updated_at"`
	DeletedAt 	gorm.DeletedAt 	`json:"deleted_at"`
}

func FromDomain(domain *students.Domain) *Student {
	return &Student{
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

func (rec *Student) ToDomain() students.Domain {
	return students.Domain{
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
