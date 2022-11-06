package auth

import (
	// "mini-project/app/middlewares"
	// "mini-project/businesses/auth"
	// controller "mini-project/controllers"
	// "mini-project/controllers/auth/request"
	// "mini-project/controllers/auth/response"
	// "net/http"

	// "github.com/golang-jwt/jwt"
	// "github.com/labstack/echo/v4"
)

// type AuthController struct {
// 	authUseCase auth.Usecase
// }

// func NewAuthController(authUC auth.Usecase) *AuthController {
// 	return &AuthController{
// 		authUseCase: authUC,
// 	}
// }

// func (ctrl *AuthController) Login(c echo.Context) error {
// 	userInput := request.Auth{}

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