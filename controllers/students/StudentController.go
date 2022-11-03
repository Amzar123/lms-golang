package students

import (
	// "mini-project/app/middlewares"
	"mini-project/businesses/students"
	// controller "mini-project/controllers"
	"mini-project/controllers/students/request"
	"mini-project/controllers/students/response"
	"net/http"

	// "github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type StudentController struct {
	studentUseCase students.Usecase
}

func NewAuthController(studentUC students.Usecase) *StudentController {
	return &StudentController{
		studentUseCase: studentUC,
	}
}

func (ctrl *StudentController) CreateStudent(c echo.Context) error {
	studentInput := request.Student{}

	if err := c.Bind(&studentInput); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid request",
		})
	}

	err := studentInput.Validate()

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"message": "validation failed",
		})
	}

	student := ctrl.studentUseCase.Create(studentInput.ToDomain())

	return c.JSON(http.StatusCreated, response.FromDomain(student))
}

// func (ctrl *AuthController) Login(c echo.Context) error {
// 	userInput := request.User{}

// 	if err := c.Bind(&userInput); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "invalid request",
// 		})
// 	}

// 	err := userInput.Validate()

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"message": "validation failed",
// 		})
// 	}

// 	token := ctrl.authUseCase.Login(userInput.ToDomain())

// 	if token == "" {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "invalid email or password",
// 		})
// 	}

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"token": token,
// 	})
// }

// func (ctrl *AuthController) Logout(c echo.Context) error {
// 	user := c.Get("user").(*jwt.Token)

// 	isListed := middlewares.CheckToken(user.Raw)

// 	if !isListed {
// 		return c.JSON(http.StatusUnauthorized, map[string]string{
// 			"message": "invalid token",
// 		})
// 	}

// 	middlewares.Logout(user.Raw)

// 	return c.JSON(http.StatusOK, map[string]string{
// 		"message": "logout success",
// 	})
// }

// func (ctrl *AuthController) GetAllUser(c echo.Context) error {
// 	usersData := ctrl.authUseCase.GetAllUser()

// 	users := []response.User{}

// 	for _, user := range usersData {
// 		users = append(users, response.FromDomain(user))
// 	}

// 	return controller.NewResponse(c, http.StatusOK, "success", "all user", users)
// }