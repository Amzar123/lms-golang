package teachers_test

import (
	"mini-project/app/middlewares"
	"mini-project/businesses/teachers"
	_teacherMock "mini-project/businesses/teachers/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	teacherRepository _teacherMock.Repository
	teacherService    teachers.Usecase

	teacherDomain teachers.Domain
)

func TestMain(m *testing.M) {
	teacherService = teachers.NewTeacherUsecase(&teacherRepository, &middlewares.ConfigJWT{})

	teacherDomain = teachers.Domain{
		NIP:      "900000",
		Email:    "guru@gmail.com",
		Password: "$2a$10$bZws9ZXgAyD/MUCM5gNsVOlx8GytFphoKL6y7uAjM/QwyHk7NKV1W",
		Name:     "Aji Muhammad Zapar",
		RoleId:   "2",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All Teacher | Valid", func(t *testing.T) {
		teacherRepository.On("GetTeachers").Return([]teachers.Domain{teacherDomain}).Once()

		result := teacherService.GetTeachers()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All Teacher | InValid", func(t *testing.T) {
		teacherRepository.On("GetTeachers").Return([]teachers.Domain{}).Once()

		result := teacherService.GetTeachers()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		teacherRepository.On("GetByID", "900000").Return(teacherDomain).Once()

		result := teacherService.GetByID("900000")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		teacherRepository.On("GetByID", "-900000").Return(teachers.Domain{}).Once()

		result := teacherService.GetByID("-900000")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		teacherRepository.On("CreateTeacher", &teacherDomain).Return(teacherDomain).Once()

		result := teacherService.CreateTeacher(&teacherDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		teacherRepository.On("CreateTeacher", &teachers.Domain{}).Return(teachers.Domain{}).Once()

		result := teacherService.CreateTeacher(&teachers.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		teacherRepository.On("Update", "900000", &teacherDomain).Return(teacherDomain).Once()

		result := teacherService.Update("900000", &teacherDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		teacherRepository.On("Update", "900000", &teachers.Domain{}).Return(teachers.Domain{}).Once()

		result := teacherService.Update("900000", &teachers.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		teacherRepository.On("Delete", "900000").Return(true).Once()

		result := teacherService.Delete("900000")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		teacherRepository.On("Delete", "-900000").Return(false).Once()

		result := teacherService.Delete("-900000")

		assert.False(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		teacherRepository.On("GetByEmail", &teacherDomain).Return(teachers.Domain{}).Once()

		result := teacherService.Login(&teacherDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		teacherRepository.On("GetByEmail", &teachers.Domain{}).Return(teachers.Domain{}).Once()

		result := teacherService.Login(&teachers.Domain{})

		assert.Empty(t, result)
	})
}
