package response

import (
	"mini-project/businesses/students"
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

func FromDomain(domain students.Domain) Student {
	return Student{
		NISN: 		domain.NISN,
		Email: 		domain.Email,
		Password: 	domain.Password,
		Name: 		domain.Name,
		Class: 		domain.Class,
		CreatedAt: 	domain.CreatedAt,
		UpdatedAt:	domain.UpdatedAt,
		DeletedAt: 	domain.DeletedAt,
	}
}
