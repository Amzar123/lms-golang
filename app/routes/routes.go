package routes

import (
	"mini-project/controllers/students"
	"mini-project/controllers/modules"
	"mini-project/controllers/teachers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   	echo.MiddlewareFunc
	JWTMiddleware      	middleware.JWTConfig
	StudentController  	students.StudentController
	ModuleController  	modules.ModuleController
	TeacherController  	teachers.TeacherController
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	students := e.Group("/api/v1/students")
	students.POST("/register", cl.StudentController.CreateStudent)
	students.POST("/login", cl.StudentController.Login)

	teachers := e.Group("/api/v1/teacher")
	teachers.POST("/register", cl.TeacherController.CreateTeacher)
	teachers.GET("", cl.TeacherController.GetTeachers)

	modules := e.Group("/api/v1/modules", middleware.JWTWithConfig(cl.JWTMiddleware),)
	modules.GET("", cl.ModuleController.GetModules)
	modules.POST("/create", cl.ModuleController.CreateModule)

	// e.POST("/logout", cl.AuthController.Logout)
}