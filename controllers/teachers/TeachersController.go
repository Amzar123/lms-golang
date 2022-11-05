package teachers

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/teachers"
	controller "mini-project/controllers"
	"mini-project/controllers/teachers/request"
	"mini-project/controllers/teachers/response"
	"net/http"

	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type TeacherController struct {
	teacherUseCase teachers.Usecase
}

func NewTeacherController(teacherUC teachers.Usecase) *TeacherController {
	return &TeacherController{
		teacherUseCase: teacherUC,
	}
}

func (ctrl *TeacherController) CreateTeacher(c echo.Context) error {
	teacherInput := request.Teacher{}

	if err := c.Bind(&teacherInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := teacherInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	teacher := ctrl.teacherUseCase.CreateTeacher(teacherInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomain(teacher))
}

func (ctrl *TeacherController) GetTeachers(c echo.Context) error {
	teachersData := ctrl.teacherUseCase.GetTeachers()

	teachers := []response.Teacher{}

	for _, teacher := range teachersData {
		teachers = append(teachers, response.FromDomain(teacher))
	}

	return controller.NewResponse(c, http.StatusOK, "success", "all module", teachers)
}