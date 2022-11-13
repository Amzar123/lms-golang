package students_test

import (
	"mini-project/app/middlewares"
	"mini-project/businesses/students"
	_studentMock "mini-project/businesses/students/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	studentRepository _studentMock.Repository
	studentService    students.Usecase

	studentDomain students.Domain
)

func TestMain(m *testing.M) {
	studentService = students.NewStudentUsecase(&studentRepository, &middlewares.ConfigJWT{})

	studentDomain = students.Domain{
		NISN:     "8090",
		Name:     "Zapar",
		Email:    "aji@gmail.com",
		Password: "TEST123",
		Class:    "10",
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		studentRepository.On("Create", &studentDomain).Return(studentDomain).Once()

		result := studentService.Create(&studentDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		studentRepository.On("Create", &students.Domain{}).Return(students.Domain{}).Once()

		result := studentService.Create(&students.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		studentRepository.On("GetByEmail", &studentDomain).Return(students.Domain{}).Once()

		result := studentService.Login(&studentDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		studentRepository.On("GetByEmail", &students.Domain{}).Return(students.Domain{}).Once()

		result := studentService.Login(&students.Domain{})

		assert.Empty(t, result)
	})
}
