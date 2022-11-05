package student

import (
	"mini-project/businesses/students"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type studentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) students.Repository {
	return &studentRepository{
		conn: conn,
	}
}

func (ur *studentRepository) Create(studentDomain *students.Domain) students.Domain {
	password, _ := bcrypt.GenerateFromPassword([]byte(studentDomain.Password), bcrypt.DefaultCost)

	rec := FromDomain(studentDomain)

	rec.Password = string(password)

	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (ur *studentRepository) GetByEmail(studentDomain *students.Domain) students.Domain {
	var student Student

	ur.conn.First(&student, "email = ?", studentDomain.Email)

	if student.NISN == "" {
		fmt.Println("Student not found")
		return students.Domain{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(studentDomain.Password))

	if err != nil {
		fmt.Println("password failed!")
		return students.Domain{}
	}

	return student.ToDomain()
}

// func (ur *userRepository) GetModules() []users.Domain {
// 	var rec []User

// 	ur.conn.Find(&rec)

// 	userDomain := []users.Domain{}

// 	for _, user := range rec {
// 		userDomain = append(userDomain, user.ToDomain())
// 	}

// 	return userDomain
// }
