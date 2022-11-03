package request

import (
	"mini-project/businesses/students"

	"github.com/go-playground/validator/v10"
)

type Student struct {
	NISN		string `json:"nisn" validate:"required"`
	Name	string `json: "name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Class	string `json:"class" validate:"required"`
}

func (req *Student) ToDomain() *students.Domain {
	return &students.Domain{
		NISN:		req.NISN,
		Name:		req.Name,
		Email:    	req.Email,
		Password: 	req.Password,
		Class:		req.Class,
	}
}

func (req *Student) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}

type StudentLogin struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (req *StudentLogin) ToDomainLogin() *students.Domain {
	return &students.Domain{
		Email:    	req.Email,
		Password: 	req.Password,
	}
}

func (req *StudentLogin) ValidateLogin() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}