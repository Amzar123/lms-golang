package users

import (
	"clean-code/businesses/users"
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
	Email     string         `json:"email" gorm:"unique" faker:"email"`
	Password  string         `json:"password" faker:"password"`
}

func FromDomain(domain *users.Domain) *User {
	return &User{
		ID:        domain.ID,
		Email:     domain.Email,
		Password:  domain.Password,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *User) ToDomain() users.Domain {
	return users.Domain{
		ID:        rec.ID,
		Email:     rec.Email,
		Password:  rec.Password,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
