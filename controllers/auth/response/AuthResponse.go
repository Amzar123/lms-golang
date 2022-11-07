package response

import (
	"mini-project/businesses/auth"
)

type Auth struct {
	Email    		string       `json:"email" gorm:"primaryKey"`
	Password     	string       `json:"password" gorm:"unique" faker:"module"`
}

func FromDomain(domain auth.Domain) Auth {
	return Auth{
		Email:       	domain.Email,
		Password:		domain.Password,
	}
}
