package request

import (
	"mini-project/businesses/assignments"

	"github.com/go-playground/validator/v10"
)

type Assignment struct {
	IdAssignment		string			`json:"idAssignment"`
	Name     			string         	`json:"name" faker:"name"`
	Deadline			string      	`json:"deadline"`
	Class				string			`json:"class" faker:"class"`
}

func (req *Assignment) ToDomain() *assignments.Domain {
	return &assignments.Domain{
		IdAssignment:		req.IdAssignment,	
		Name:				req.Name,
		Deadline:			req.Deadline,
		Class:				req.Class,
	}
}

func (req *Assignment) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}