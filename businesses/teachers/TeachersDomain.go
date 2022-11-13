package teachers

import (
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	NIP 		string
	Email 		string
	Password 	string
	Name 		string
	RoleId 		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt	gorm.DeletedAt
}

type Usecase interface {
	GetTeachers() []Domain
	CreateTeacher(teacherDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
	Login(userDomain *Domain) string
}

type Repository interface {
	GetTeachers() []Domain
	CreateTeacher(teacherDomain *Domain) Domain
	GetByID(id string) Domain
	Update(id string, blogDomain *Domain) Domain
	Delete(id string) bool
	GetByEmail(studentDomain *Domain) Domain
}
