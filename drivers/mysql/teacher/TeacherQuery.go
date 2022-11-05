package teacher

import (
	"mini-project/businesses/teachers"
	// "fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type teacherRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) teachers.Repository {
	return &teacherRepository{
		conn: conn,
	}
}

func (ur *teacherRepository) CreateTeacher(teacherDomain *teachers.Domain) teachers.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(teacherDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(teacherDomain)

	rec.Password = string(password)

	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}