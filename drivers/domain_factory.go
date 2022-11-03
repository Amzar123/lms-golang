package drivers

import (
	// userDomain "mini-project/businesses/users"
	// userDB "mini-project/drivers/mysql/users"

	studentDomain "mini-project/businesses/students"
	studentDB "mini-project/drivers/mysql/student"

	"gorm.io/gorm"
)

// func NewUserRepository(conn *gorm.DB) userDomain.Repository {
// 	return userDB.NewMySQLRepository(conn)
// }

func NewStudentRepository(conn *gorm.DB) studentDomain.Repository {
	return studentDB.NewMySQLRepository(conn)
}