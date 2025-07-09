package service_test

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"twitter-challenge-exercise/internal/adapter/repository/database"
	"twitter-challenge-exercise/internal/core/domain"
	"twitter-challenge-exercise/internal/core/service"
	"twitter-challenge-exercise/pkg"
)

func TestUser_CreateUser_Success(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, pkg.NewEntityNotFoundError("user with email email@test.com not found"))
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))

	testUser.Password = "hashed_test_password"

	userRepository.On("CreateUser", mock.Anything, testUser).Return(testUser, nil)
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, nil)

	userService := service.NewUserService(userRepository)
	user, err := userService.CreateUser(context.Background(), testUser)

	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ExistingUserEmail(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(testUser, nil)

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "there is an existing user with email 'email@test.com'")

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ErrorLookingUserByEmail(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, errors.New("error looking user by email"))

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "error looking user by email")

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ExistingUsername(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, pkg.NewEntityNotFoundError("user with email email@test.com not found"))
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(testUser, nil)

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "there is an existing user with username 'jdoe'")

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ErrorLookingUserByUsername(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, pkg.NewEntityNotFoundError("user with email email@test.com not found"))
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, errors.New("error looking user by username"))

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "error looking user by username")

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ErrorCreatingUser(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, pkg.NewEntityNotFoundError("user with email email@test.com not found"))
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))

	testUser.Password = "hashed_test_password"

	userRepository.On("CreateUser", mock.Anything, testUser).Return(testUser, errors.New("error in database"))

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "error creating user")
	assert.True(t, pkg.IsServerError(err))

	userRepository.AssertExpectations(t)
}

func TestUser_CreateUser_ErrorReturningEntity(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(domain.User{}, pkg.NewEntityNotFoundError("user with email email@test.com not found"))
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))

	testUser.Password = "hashed_test_password"

	userRepository.On("CreateUser", mock.Anything, testUser).Return(testUser, nil)
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, errors.New("error in database"))

	userService := service.NewUserService(userRepository)
	_, err := userService.CreateUser(context.Background(), testUser)

	assert.Error(t, err, "error in database")

	userRepository.AssertExpectations(t)
}

func TestUser_UpdateUser_Success(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, nil).Once()

	testNewUser := getTestUser()
	testNewUser.Name = "another_name"

	userRepository.On("UpdateUser", mock.Anything, testNewUser).Return(testNewUser, nil)
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testNewUser, nil).Once()

	userService := service.NewUserService(userRepository)
	user, err := userService.UpdateUser(context.Background(), testNewUser)

	assert.NoError(t, err)
	assert.Equal(t, testNewUser, user)

	userRepository.AssertExpectations(t)
}

func TestUser_UpdateUser_ExistingUsername(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(testUser, nil)

	testNewUser := getTestUser()
	testNewUser.ID = 126
	testNewUser.Name = "another_name"

	userService := service.NewUserService(userRepository)
	_, err := userService.UpdateUser(context.Background(), testNewUser)

	assert.Error(t, err, "there is an existing user with username 'jdoe'")

	userRepository.AssertExpectations(t)
}

func TestUser_UpdateUser_ErrorLookingUserByUsername(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, errors.New("error looking user by username"))

	userService := service.NewUserService(userRepository)
	_, err := userService.UpdateUser(context.Background(), testUser)

	assert.Error(t, err, "error looking user by username")

	userRepository.AssertExpectations(t)
}

func TestUser_UpdateUser_ErrorReturningEntity(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(domain.User{}, errors.New("error in database")).Once()

	userService := service.NewUserService(userRepository)
	_, err := userService.UpdateUser(context.Background(), testUser)

	assert.Error(t, err, "error in database")

	userRepository.AssertExpectations(t)
}

func TestUser_UpdateUser_ErrorUpdatingUser(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByUsername", mock.Anything, testUser.Username).Return(domain.User{}, pkg.NewEntityNotFoundError("user with username jdoe not found"))
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, nil)

	testNewUser := getTestUser()
	testNewUser.Name = "another_name"

	userRepository.On("UpdateUser", mock.Anything, testNewUser).Return(domain.User{}, errors.New("error in database"))

	userService := service.NewUserService(userRepository)
	_, err := userService.UpdateUser(context.Background(), testNewUser)

	assert.Error(t, err, "error updating user")
	assert.True(t, pkg.IsServerError(err))

	userRepository.AssertExpectations(t)
}

func TestUser_GetUserByID_Success(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, nil)

	userService := service.NewUserService(userRepository)

	user, err := userService.GetUserByID(context.Background(), testUser.ID)

	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	userRepository.AssertExpectations(t)
}

func TestUser_GetUserByID_Error(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByID", mock.Anything, testUser.ID).Return(testUser, errors.New("error in database"))

	userService := service.NewUserService(userRepository)

	_, err := userService.GetUserByID(context.Background(), testUser.ID)

	assert.Error(t, err, "error in database")

	userRepository.AssertExpectations(t)
}

func TestUser_GetUserByEmail_Success(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(testUser, nil)

	userService := service.NewUserService(userRepository)

	user, err := userService.GetUserByEmail(context.Background(), testUser.Email)

	assert.NoError(t, err)
	assert.Equal(t, testUser, user)

	userRepository.AssertExpectations(t)
}

func TestUser_GetUserByEmail_Error(t *testing.T) {
	testUser := getTestUser()

	userRepository := database.NewUserMockRepository()
	userRepository.On("GetUserByEmail", mock.Anything, testUser.Email).Return(testUser, errors.New("error in database"))

	userService := service.NewUserService(userRepository)

	_, err := userService.GetUserByEmail(context.Background(), testUser.Email)

	assert.Error(t, err, "error in database")

	userRepository.AssertExpectations(t)
}

func TestUser_SearchUsers_Success(t *testing.T) {
	testUser := getTestUser()

	anotherTestUser := getTestUser()
	anotherTestUser.ID = 49
	anotherTestUser.Username = "user"
	anotherTestUser.Email = "test@email.com"

	testUsers := []domain.User{
		testUser,
		anotherTestUser,
	}

	userRepository := database.NewUserMockRepository()
	userRepository.On("SearchUsers", mock.Anything).Return(testUsers, nil)

	userService := service.NewUserService(userRepository)

	users, err := userService.SearchUsers(context.Background())

	assert.NoError(t, err)
	assert.Equal(t, testUsers, users)

	userRepository.AssertExpectations(t)
}

func TestUser_SearchUsers_Error(t *testing.T) {
	userRepository := database.NewUserMockRepository()
	userRepository.On("SearchUsers", mock.Anything).Return(nil, errors.New("error in database"))

	userService := service.NewUserService(userRepository)

	_, err := userService.SearchUsers(context.Background())

	assert.Error(t, err, "error in database")

	userRepository.AssertExpectations(t)
}

func getTestUser() domain.User {
	return domain.User{
		ID:        42,
		Name:      "John",
		Surname:   "Doe",
		Email:     "email@test.com",
		Password:  "super-strong-password",
		Username:  "jdoe",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}
}
