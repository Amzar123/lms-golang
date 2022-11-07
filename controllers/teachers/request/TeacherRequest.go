package request

import (
	"mini-project/businesses/teachers"

	"github.com/go-playground/validator/v10"
)

type Teacher struct {
	NIP        	string          `json:"nip" gorm:"primaryKey"`
	Email     	string         	`json:"email" gorm:"unique" faker:"email"`
	Password  	string         	`json:"password" faker:"password"`
	Name		string			`json:"name" faker:"name"`
	RoleId		string			`json:"roleId"`
}

func (req *Teacher) ToDomain() *teachers.Domain {
	return &teachers.Domain{
		NIP:        	req.NIP,
		Email:     		req.Email,
		Password:  		req.Password,
		Name:			req.Name,
		RoleId:			req.RoleId,
	}
}

func (req *Teacher) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

type TeacherLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *TeacherLogin) ToDomainLogin() *teachers.Domain {
	return &teachers.Domain{
		Email:    	req.Email,
		Password: 	req.Password,
	}
}

func (req *TeacherLogin) ValidateLogin() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}