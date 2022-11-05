package drivers

import (
	// userDomain "mini-project/businesses/users"
	// userDB "mini-project/drivers/mysql/users"

	studentDomain "mini-project/businesses/students"
	moduleDomain "mini-project/businesses/modules"
	teacherDomain "mini-project/businesses/teachers"

	studentDB "mini-project/drivers/mysql/student"
	moduleDB "mini-project/drivers/mysql/module"
	teacherDB "mini-project/drivers/mysql/teacher"

	"gorm.io/gorm"
)

func NewStudentRepository(conn *gorm.DB) studentDomain.Repository {
	return studentDB.NewMySQLRepository(conn)
}

func NewModuleRepository(conn *gorm.DB) moduleDomain.Repository {
	return moduleDB.NewMySQLRepository(conn)
}

func NewTeacherRepository(conn *gorm.DB) teacherDomain.Repository {
	return teacherDB.NewMySQLRepository(conn)
}