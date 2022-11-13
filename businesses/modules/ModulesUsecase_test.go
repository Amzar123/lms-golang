package modules_test

import (
	"mini-project/app/middlewares"
	"mini-project/businesses/modules"
	_moduleMock "mini-project/businesses/modules/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	moduleRepository _moduleMock.Repository
	moduleService    modules.Usecase

	moduleDomain modules.Domain
)

func TestMain(m *testing.M) {
	moduleService = modules.NewModuleUsecase(&moduleRepository, &middlewares.ConfigJWT{})

	moduleDomain = modules.Domain{
		IdModule		: "",
		Module			: "ini adalah module",
		TeacherId		: "",
	}

	m.Run()
}

func TestGetAll(t *testing.T) {
	t.Run("Get All module | Valid", func(t *testing.T) {
		moduleRepository.On("GetModules").Return([]modules.Domain{moduleDomain}).Once()

		result := moduleService.GetModules()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All module | InValid", func(t *testing.T) {
		moduleRepository.On("GetModules").Return([]modules.Domain{}).Once()

		result := moduleService.GetModules()

		assert.Equal(t, 0, len(result))
	})
}

func TestGetByID(t *testing.T) {
	t.Run("Get By ID | Valid", func(t *testing.T) {
		moduleRepository.On("GetByID", "7111a840-3099-4c62-85b6-00d665d42cba").Return(moduleDomain).Once()

		result := moduleService.GetByID("7111a840-3099-4c62-85b6-00d665d42cba")

		assert.NotNil(t, result)
	})

	t.Run("Get By ID | InValid", func(t *testing.T) {
		moduleRepository.On("GetByID", "-7111a840-3099-4c62-85b6-00d665d42cba").Return(modules.Domain{}).Once()

		result := moduleService.GetByID("-7111a840-3099-4c62-85b6-00d665d42cba")

		assert.NotNil(t, result)
	})
}

func TestCreate(t *testing.T) {
	t.Run("Create | Valid", func(t *testing.T) {
		moduleRepository.On("Create", &moduleDomain).Return(moduleDomain).Once()

		result := moduleService.Create(&moduleDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create | InValid", func(t *testing.T) {
		moduleRepository.On("Create", &modules.Domain{}).Return(modules.Domain{}).Once()

		result := moduleService.Create(&modules.Domain{})

		assert.NotNil(t, result)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update | Valid", func(t *testing.T) {
		moduleRepository.On("Update", "7111a840-3099-4c62-85b6-00d665d42cba", &moduleDomain).Return(moduleDomain).Once()

		result := moduleService.Update("7111a840-3099-4c62-85b6-00d665d42cba", &moduleDomain)

		assert.NotNil(t, result)
	})

	t.Run("Update | InValid", func(t *testing.T) {
		moduleRepository.On("Update", "7111a840-3099-4c62-85b6-00d665d42cba", &modules.Domain{}).Return(modules.Domain{}).Once()

		result := moduleService.Update("7111a840-3099-4c62-85b6-00d665d42cba", &modules.Domain{})

		assert.NotNil(t, result)
	})
}

func TestDelete(t *testing.T) {
	t.Run("Delete | Valid", func(t *testing.T) {
		moduleRepository.On("Delete", "7111a840-3099-4c62-85b6-00d665d42cba").Return(true).Once()

		result := moduleService.Delete("7111a840-3099-4c62-85b6-00d665d42cba")

		assert.True(t, result)
	})

	t.Run("Delete | InValid", func(t *testing.T) {
		moduleRepository.On("Delete", "-7111a840-3099-4c62-85b6-00d665d42cba").Return(false).Once()

		result := moduleService.Delete("-7111a840-3099-4c62-85b6-00d665d42cba")

		assert.False(t, result)
	})
}
