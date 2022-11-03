package student

import (
	"mini-project/businesses/students"
	// "fmt"

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

// func (ur *userRepository) GetByEmail(userDomain *users.Domain) users.Domain {
// 	var user User

// 	ur.conn.First(&user, "email = ?", userDomain.Email)

// 	if user.ID == 0 {
// 		fmt.Println("user not found")
// 		return users.Domain{}
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDomain.Password))

// 	if err != nil {
// 		fmt.Println("password failed!")
// 		return users.Domain{}
// 	}

// 	return user.ToDomain()
// }

// func (ur *userRepository) GetAllUser() []users.Domain {
// 	var rec []User

// 	ur.conn.Find(&rec)

// 	userDomain := []users.Domain{}

// 	for _, user := range rec {
// 		userDomain = append(userDomain, user.ToDomain())
// 	}

// 	return userDomain
// }