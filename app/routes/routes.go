package routes

import (
	"mini-project/controllers/students"
	"mini-project/controllers/modules"
	"mini-project/controllers/teachers"
	// "mini-project/controllers/auth"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   	echo.MiddlewareFunc
	JWTMiddleware      	middleware.JWTConfig
	StudentController  	students.StudentController
	ModuleController  	modules.ModuleController
	TeacherController  	teachers.TeacherController
	// AuthController  	auth.AuthController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	students := e.Group("/api/v1/students")
	students.POST("/register", cl.StudentController.CreateStudent)

	teachers := e.Group("/api/v1/teachers")
	teachers.GET("", cl.TeacherController.GetTeachers)
	teachers.POST("/register", cl.TeacherController.CreateTeacher)

	modules := e.Group("/api/v1/modules", middleware.JWTWithConfig(cl.JWTMiddleware))
	modules.GET("", cl.ModuleController.GetModules)
	modules.POST("/create", cl.ModuleController.CreateModule)

	// auth := e.Group("/api/v1")
	// auth.POST("/login", cl.AuthController.Login)
	// auth.POST("/logout", cl.AuthController.Logout)
}