package teacher

import (
	"mini-project/businesses/teachers"
	"fmt"

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

func (ur *teacherRepository) GetTeachers() []teachers.Domain {
	var rec []Teacher

	ur.conn.Find(&rec)

	teacherDomain := []teachers.Domain{}

	for _, teacher := range rec {
		teacherDomain = append(teacherDomain, teacher.ToDomain())
	}

	return teacherDomain
}

func (ur *teacherRepository) GetByEmail(teacherDomain *teachers.Domain) teachers.Domain {
	var teacher Teacher

	ur.conn.First(&teacher, "email = ?", teacherDomain.Email)

	if teacher.NIP == "" {
		fmt.Println("Teacher not found")
		return teachers.Domain{}
	}

	err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(teacherDomain.Password))

	if err != nil {
		fmt.Println("password failed!")
		return teachers.Domain{}
	}

	return teacher.ToDomain()
}

func (nr *teacherRepository) GetByID(id string) teachers.Domain {
	var teacher Teacher

	nr.conn.First(&teacher, "n_ip = ?", id)

	return teacher.ToDomain()
}

func (nr *teacherRepository) Update(id string, teacherDomain *teachers.Domain) teachers.Domain {
	var teacher teachers.Domain = nr.GetByID(id)

	updatedTeacher := FromDomain(&teacher)

	updatedTeacher.Email = teacherDomain.Email
	updatedTeacher.Name = teacherDomain.Name
	// updatedTeacher.Email = teacherDomain.Email
	nr.conn.Save(&updatedTeacher)

	return updatedTeacher.ToDomain()
}

func (nr *teacherRepository) Delete(id string) bool {
	var teacher teachers.Domain = nr.GetByID(id)

	deletedTeacher := FromDomain(&teacher)

	result := nr.conn.Delete(&deletedTeacher)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
