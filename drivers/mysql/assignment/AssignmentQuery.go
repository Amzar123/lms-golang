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

func (nr *assignmentRepository) GetByID(id string) assignments.Domain {
	var assignment Assignment

	nr.conn.First(&assignment, "id_assignment = ?", id)

	return assignment.ToDomain()
}

func (nr *assignmentRepository) Update(id string, assignmentDomain *assignments.Domain) assignments.Domain {
	var assignment assignments.Domain = nr.GetByID(id)

	updatedAssignment := FromDomain(&assignment)

	updatedAssignment.Name = assignmentDomain.Name
	updatedAssignment.Deadline = assignmentDomain.Deadline
	updatedAssignment.Class = assignmentDomain.Class

	nr.conn.Save(&updatedAssignment)

	return updatedAssignment.ToDomain()
}

func (nr *assignmentRepository) Delete(id string) bool {
	var assignment assignments.Domain = nr.GetByID(id)

	deletedAssignment := FromDomain(&assignment)

	result := nr.conn.Delete(&deletedAssignment)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
