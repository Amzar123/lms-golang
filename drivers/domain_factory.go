package drivers

import (
	userDomain "clean-code/businesses/users"
	userDB "clean-code/drivers/mysql/users"

	"gorm.io/gorm"
)

func NewUserRepository(conn *gorm.DB) userDomain.Repository {
	return userDB.NewMySQLRepository(conn)
}