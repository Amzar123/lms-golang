package assignment

import (
	"mini-project/businesses/assignments"
	// "fmt"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type assignmentRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) assignments.Repository {
	return &assignmentRepository{
		conn: conn,
	}
}

func (ur *assignmentRepository) GetAssignments() []assignments.Domain {
	var rec []Assignment

	ur.conn.Find(&rec)

	assignmentDomain := []assignments.Domain{}

	for _, assignment := range rec {
		assignmentDomain = append(assignmentDomain, assignment.ToDomain())
	}

	return assignmentDomain
}

func (ur *assignmentRepository) Create(assignmentDomain *assignments.Domain) assignments.Domain {

	rec := FromDomain(assignmentDomain)
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}