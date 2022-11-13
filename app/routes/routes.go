package routes

import (
	"mini-project/controllers/students"
	"mini-project/controllers/modules"
	"mini-project/controllers/teachers"
	"mini-project/controllers/assignments"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	LoggerMiddleware   		echo.MiddlewareFunc
	JWTMiddleware      		middleware.JWTConfig
	StudentController  		students.StudentController
	ModuleController  		modules.ModuleController
	TeacherController  		teachers.TeacherController
	AssignmentController  	assignments.AssignmentController
	
}

func (cl *ControllerList) RouteRegister(e *echo.Echo) {
	e.Use(cl.LoggerMiddleware)

	students := e.Group("/api/v1/students")
	students.POST("/register", cl.StudentController.CreateStudent)

	teachers := e.Group("/api/v1/teachers")
	teachers.GET("", cl.TeacherController.GetTeachers)
	teachers.POST("/register", cl.TeacherController.CreateTeacher)
	teachers.DELETE("/delete/:id", cl.TeacherController.Delete)
	teachers.PUT("/update/:id", cl.TeacherController.Update)
	teachers.GET("/:id", cl.TeacherController.GetByID)

	modules := e.Group("/api/v1/modules", middleware.JWTWithConfig(cl.JWTMiddleware))
	modules.GET("", cl.ModuleController.GetModules)
	modules.POST("/create", cl.ModuleController.CreateModule)
	modules.DELETE("/delete/:id", cl.ModuleController.Delete)
	modules.PUT("/update/:id", cl.ModuleController.Update)
	modules.GET("/:id", cl.ModuleController.GetByID)

	assignments := e.Group("/api/v1/assignments")
	assignments.GET("", cl.AssignmentController.GetAssignments)
	assignments.POST("/create", cl.AssignmentController.CreateAssignment)
	assignments.DELETE("/delete/:id", cl.AssignmentController.Delete)
	assignments.PUT("/update/:id", cl.AssignmentController.Update)
	assignments.GET("/:id", cl.AssignmentController.GetByID)

	students.POST("/login", cl.StudentController.Login)
	students.POST("/logout", cl.StudentController.Logout)

	teachers.POST("/login", cl.TeacherController.Login)
	teachers.POST("/logout", cl.TeacherController.Logout)
}

