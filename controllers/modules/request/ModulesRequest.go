package request

import (
	"mini-project/businesses/modules"

	"github.com/go-playground/validator/v10"
)

type Module struct {
	Module			string	`json:"module" gorm:"unique" faker:"module"`
	TeacherId		string	`json:"teacherId"`
}

func (req *Module) ToDomain() *modules.Domain {
	return &modules.Domain{
		Module:			req.Module,
		TeacherId:		req.TeacherId,
	}
}

func (req *Module) Validate() error {
	validate := validator.New()

	err := validate.Struct(req)

	return err
}