package teachers

import (
	"mini-project/app/middlewares"
	"mini-project/businesses/teachers"
	controller "mini-project/controllers"
	"mini-project/controllers/teachers/request"
	"mini-project/controllers/teachers/response"
	"net/http"

	"github.com/golang-jwt/jwt"
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

func (ctrl *TeacherController) GetByID(c echo.Context) error {
	var id string = c.Param("id")

	teacher := ctrl.teacherUseCase.GetByID(id)

	if teacher.NIP == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "teacher not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "note found", response.FromDomain(teacher))
}

func (ctrl *TeacherController) Update(c echo.Context) error {
	input := request.Teacher{}

	if err := c.Bind(&input); err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	var teacherId string = c.Param("id")

	err := input.Validate()

	if err != nil {
		return controller.NewResponse(c, http.StatusBadRequest, "failed", "validation failed", "")
	}

	teacher := ctrl.teacherUseCase.Update(teacherId, input.ToDomain())

	if teacher.NIP == "" {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "Teacher not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Teacher updated", response.FromDomain(teacher))
}

func (ctrl *TeacherController) Delete(c echo.Context) error {
	var teacherId string = c.Param("id")

	isSuccess := ctrl.teacherUseCase.Delete(teacherId)

	if !isSuccess {
		return controller.NewResponse(c, http.StatusNotFound, "failed", "Teacher not found", "")
	}

	return controller.NewResponse(c, http.StatusOK, "success", "Teacher deleted", "")
}

func (ctrl *TeacherController) Login(c echo.Context) error {
	teacherInput := request.TeacherLogin{}

	if err := c.Bind(&teacherInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := teacherInput.ValidateLogin()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	token := ctrl.teacherUseCase.Login(teacherInput.ToDomainLogin())

	if token == "" {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid email or password",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ctrl *TeacherController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	isListed := middlewares.CheckToken(user.Raw)

	if !isListed {
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"message": "invalid token",
		})
	}

	middlewares.Logout(user.Raw)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout success",
	})
}