package drivers

import (
	// userDomain "mini-project/businesses/users"
	// userDB "mini-project/drivers/mysql/users"

	studentDomain "mini-project/businesses/students"
	moduleDomain "mini-project/businesses/modules"
	teacherDomain "mini-project/businesses/teachers"
	assignmentDomain "mini-project/businesses/assignments"

	studentDB "mini-project/drivers/mysql/student"
	moduleDB "mini-project/drivers/mysql/module"
	teacherDB "mini-project/drivers/mysql/teacher"
	assignmentDB "mini-project/drivers/mysql/assignment"

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

func NewAssignmentRepository(conn *gorm.DB) assignmentDomain.Repository {
	return assignmentDB.NewMySQLRepository(conn)
}