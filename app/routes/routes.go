package routes

import (
	"mini-project/controllers/students"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   echo.MiddlewareFunc
	JWTMiddleware      middleware.JWTConfig
	StudentController  students.StudentController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	students := e.Group("/api/v1/students")
	students.POST("/register", cl.StudentController.CreateStudent)
	students.POST("/login", cl.StudentController.Login)
	// students.GET("/modules", middleware.JWTWithConfig(cl.JWTMiddleware), cl.StudentController.Modules)

	// e.POST("/logout", cl.AuthController.Logout)

}