package module

import (
	"mini-project/businesses/modules"
	// "fmt"

	// "golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type moduleRepository struct {
	conn *gorm.DB
}

func NewMySQLRepository(conn *gorm.DB) modules.Repository {
	return &moduleRepository{
		conn: conn,
	}
}

func (ur *moduleRepository) GetModules() []modules.Domain {
	var rec []Module

	ur.conn.Find(&rec)

	moduleDomain := []modules.Domain{}

	for _, module := range rec {
		moduleDomain = append(moduleDomain, module.ToDomain())
	}

	return moduleDomain
}
