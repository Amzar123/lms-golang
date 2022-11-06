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

func (ur *moduleRepository) Create(moduleDomain *modules.Domain) modules.Domain {

	rec := FromDomain(moduleDomain)
	result := ur.conn.Create(&rec)

	result.Last(&rec)

	return rec.ToDomain()
}

func (nr *moduleRepository) GetByID(id string) modules.Domain {
	var module Module

	nr.conn.First(&module, "id_module = ?", id)

	return module.ToDomain()
}

func (nr *moduleRepository) Update(id string, moduleDomain *modules.Domain) modules.Domain {
	var module modules.Domain = nr.GetByID(id)

	updatedModule := FromDomain(&module)

	updatedModule.Module = moduleDomain.Module
	nr.conn.Save(&updatedModule)

	return updatedModule.ToDomain()
}

func (nr *moduleRepository) Delete(id string) bool {
	var module modules.Domain = nr.GetByID(id)

	deletedModule := FromDomain(&module)

	result := nr.conn.Delete(&deletedModule)

	if result.RowsAffected == 0 {
		return false
	}

	return true
}
