package drivers

import (
	// userDomain "mini-project/businesses/users"
	// userDB "mini-project/drivers/mysql/users"

	studentDomain "mini-project/businesses/students"
	moduleDomain "mini-project/businesses/modules"

	studentDB "mini-project/drivers/mysql/student"
	moduleDB "mini-project/drivers/mysql/module"

	"gorm.io/gorm"
)

func NewStudentRepository(conn *gorm.DB) studentDomain.Repository {
	return studentDB.NewMySQLRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewMySQLRepository(conn)
}