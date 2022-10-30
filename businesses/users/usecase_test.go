package users_test

import (
	"clean-code/app/middlewares"
	"clean-code/businesses/users"
	_userMock "clean-code/businesses/users/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	usersRepository _userMock.Repository
	usersService    users.Usecase

	usersDomain users.Domain
)

func TestMain(m *testing.M) {
	usersService = users.NewUserUsecase(&usersRepository, &middlewares.ConfigJWT{})

	usersDomain = users.Domain{
		Email:    "test@test.com",
		Password: "123123",
	}

	m.Run()
}

func TestCreate(t *testing.T) {
	t.Run("Create User | Valid", func(t *testing.T) {
		usersRepository.On("Create User", &usersDomain).Return(usersDomain).Once()

		result := usersService.Create(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Create User | InValid", func(t *testing.T) {
		usersRepository.On("Create User", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Create(&users.Domain{})

		assert.NotNil(t, result)
	})
}

func TestLogin(t *testing.T) {
	t.Run("Login | Valid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &usersDomain).Return(users.Domain{}).Once()

		result := usersService.Login(&usersDomain)

		assert.NotNil(t, result)
	})

	t.Run("Login | InValid", func(t *testing.T) {
		usersRepository.On("GetByEmail", &users.Domain{}).Return(users.Domain{}).Once()

		result := usersService.Login(&users.Domain{})

		assert.Empty(t, result)
	})
}

func TestGetAllUser(t *testing.T) {
	t.Run("Get All User| Valid", func(t *testing.T) {
		userRepository.On("GetAllUser").Return([]users.Domain{userDomain}).Once()

		result := userService.GetAllUser()

		assert.Equal(t, 1, len(result))
	})

	t.Run("Get All User| InValid", func(t *testing.T) {
		userRepository.On("GetAllUser").Return([]users.Domain{}).Once()

		result := userService.GetAllUser()

		assert.Equal(t, 0, len(result))
	})
}