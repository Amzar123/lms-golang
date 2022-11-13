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
		NISN: "8090",
		Name: "Zapar",
		Email: "aji@gmail.com",
		Password: "TEST123",
		Class: "10",
	}

	m.Run()
}

// func TestGetAll(t *testing.T) {
// 	t.Run("Get All Teacher | Valid", func(t *testing.T) {
// 		teacherRepository.On("GetTeachers").Return([]teachers.Domain{teacherDomain}).Once()

// 		result := teacherService.GetTeachers()

// 		assert.Equal(t, 1, len(result))
// 	})

// 	t.Run("Get All Teacher | InValid", func(t *testing.T) {
// 		teacherRepository.On("GetTeachers").Return([]teachers.Domain{}).Once()

// 		result := teacherService.GetTeachers()

// 		assert.Equal(t, 0, len(result))
// 	})
// }

// func TestGetByID(t *testing.T) {
// 	t.Run("Get By ID | Valid", func(t *testing.T) {
// 		teacherRepository.On("GetByID", "900000").Return(teacherDomain).Once()

// 		result := teacherService.GetByID("900000")

// 		assert.NotNil(t, result)
// 	})

// 	t.Run("Get By ID | InValid", func(t *testing.T) {
// 		teacherRepository.On("GetByID", "-900000").Return(teachers.Domain{}).Once()

// 		result := teacherService.GetByID("-900000")

// 		assert.NotNil(t, result)
// 	})
// }

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

// func TestUpdate(t *testing.T) {
// 	t.Run("Update | Valid", func(t *testing.T) {
// 		teacherRepository.On("Update", "900000", &teacherDomain).Return(teacherDomain).Once()

// 		result := teacherService.Update("900000", &teacherDomain)

// 		assert.NotNil(t, result)
// 	})

// 	t.Run("Update | InValid", func(t *testing.T) {
// 		teacherRepository.On("Update", "900000", &teachers.Domain{}).Return(teachers.Domain{}).Once()

// 		result := teacherService.Update("900000", &teachers.Domain{})

// 		assert.NotNil(t, result)
// 	})
// }

// func TestDelete(t *testing.T) {
// 	t.Run("Delete | Valid", func(t *testing.T) {
// 		teacherRepository.On("Delete", "900000").Return(true).Once()

// 		result := teacherService.Delete("900000")

// 		assert.True(t, result)
// 	})

// 	t.Run("Delete | InValid", func(t *testing.T) {
// 		teacherRepository.On("Delete", "-900000").Return(false).Once()

// 		result := teacherService.Delete("-900000")

// 		assert.False(t, result)
// 	})
// }

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