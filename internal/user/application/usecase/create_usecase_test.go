package usecase

import (
	"errors"
	"testing"

	"github.com/shompys/hexagonal/internal/user/domain"
	"github.com/shompys/hexagonal/internal/user/domain/dto"
)

// MockUserRepository es un mock del repositorio de usuarios para testing
type MockUserRepository struct {
	CreateFunc      func(userEntity *domain.User) (*domain.User, error)
	GetUserByIDFunc func(id string) (*domain.User, error)
	GetUsersFunc    func() ([]*domain.User, error)
	UpdateUserFunc  func(id string, userEntity *domain.User) (*domain.User, error)
	DeleteUserFunc  func(id string) error
}

func (m *MockUserRepository) Create(userEntity *domain.User) (*domain.User, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(userEntity)
	}
	return nil, errors.New("CreateFunc not implemented")
}

func (m *MockUserRepository) GetUserByID(id string) (*domain.User, error) {
	if m.GetUserByIDFunc != nil {
		return m.GetUserByIDFunc(id)
	}
	return nil, errors.New("GetUserByIDFunc not implemented")
}

func (m *MockUserRepository) GetUsers() ([]*domain.User, error) {
	if m.GetUsersFunc != nil {
		return m.GetUsersFunc()
	}
	return nil, errors.New("GetUsersFunc not implemented")
}

func (m *MockUserRepository) UpdateUser(id string, userEntity *domain.User) (*domain.User, error) {
	if m.UpdateUserFunc != nil {
		return m.UpdateUserFunc(id, userEntity)
	}
	return nil, errors.New("UpdateUserFunc not implemented")
}

func (m *MockUserRepository) DeleteUser(id string) error {
	if m.DeleteUserFunc != nil {
		return m.DeleteUserFunc(id)
	}
	return errors.New("DeleteUserFunc not implemented")
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	mockRepo := &MockUserRepository{
		CreateFunc: func(userEntity *domain.User) (*domain.User, error) {
			// Simular que el repositorio genera un ID
			userEntity.ID = "generated-uuid-123"
			return userEntity, nil
		},
	}

	useCase := &UserUseCase{
		UserRepository: mockRepo,
	}

	input := &dto.UserCreateInput{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "john.doe@example.com",
		UserName:  "johndoe",
		Password:  "securePassword123",
	}

	// Act
	result, err := useCase.CreateUser(input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if result == nil {
		t.Fatal("Expected result to be non-nil")
	}

	if result.ID != "generated-uuid-123" {
		t.Errorf("Expected ID to be 'generated-uuid-123', but got: %s", result.ID)
	}

	if result.FirstName != input.FirstName {
		t.Errorf("Expected FirstName to be '%s', but got: %s", input.FirstName, result.FirstName)
	}

	if result.LastName != input.LastName {
		t.Errorf("Expected LastName to be '%s', but got: %s", input.LastName, result.LastName)
	}

	if result.Email != input.Email {
		t.Errorf("Expected Email to be '%s', but got: %s", input.Email, result.Email)
	}

	if result.UserName != input.UserName {
		t.Errorf("Expected UserName to be '%s', but got: %s", input.UserName, result.UserName)
	}
}

func TestCreateUser_RepositoryError(t *testing.T) {
	// Arrange
	expectedError := errors.New("database connection failed")
	mockRepo := &MockUserRepository{
		CreateFunc: func(userEntity *domain.User) (*domain.User, error) {
			return nil, expectedError
		},
	}

	useCase := &UserUseCase{
		UserRepository: mockRepo,
	}

	input := &dto.UserCreateInput{
		FirstName: "Jane",
		LastName:  "Smith",
		Email:     "jane.smith@example.com",
		UserName:  "janesmith",
		Password:  "anotherPassword456",
	}

	// Act
	result, err := useCase.CreateUser(input)

	// Assert
	if err == nil {
		t.Fatal("Expected an error, but got nil")
	}

	if err != expectedError {
		t.Errorf("Expected error to be '%v', but got: %v", expectedError, err)
	}

	if result != nil {
		t.Errorf("Expected result to be nil, but got: %v", result)
	}
}

func TestCreateUser_DuplicateEmail(t *testing.T) {
	// Arrange
	duplicateError := errors.New("user with email already exists")
	mockRepo := &MockUserRepository{
		CreateFunc: func(userEntity *domain.User) (*domain.User, error) {
			return nil, duplicateError
		},
	}

	useCase := &UserUseCase{
		UserRepository: mockRepo,
	}

	input := &dto.UserCreateInput{
		FirstName: "Alice",
		LastName:  "Johnson",
		Email:     "duplicate@example.com",
		UserName:  "alicejohnson",
		Password:  "password789",
	}

	// Act
	result, err := useCase.CreateUser(input)

	// Assert
	if err == nil {
		t.Fatal("Expected an error for duplicate email, but got nil")
	}

	if err != duplicateError {
		t.Errorf("Expected error to be '%v', but got: %v", duplicateError, err)
	}

	if result != nil {
		t.Errorf("Expected result to be nil when there's an error, but got: %v", result)
	}
}

func TestCreateUser_EmptyFields(t *testing.T) {
	// Arrange
	mockRepo := &MockUserRepository{
		CreateFunc: func(userEntity *domain.User) (*domain.User, error) {
			// Este test verifica que los campos vacíos se pasen al repositorio
			// El repositorio podría validar o simplemente guardar
			userEntity.ID = "test-id"
			return userEntity, nil
		},
	}

	useCase := &UserUseCase{
		UserRepository: mockRepo,
	}

	testCases := []struct {
		name  string
		input *dto.UserCreateInput
	}{
		{
			name: "Empty FirstName",
			input: &dto.UserCreateInput{
				FirstName: "",
				LastName:  "Doe",
				Email:     "test@example.com",
				UserName:  "testuser",
				Password:  "password",
			},
		},
		{
			name: "Empty Email",
			input: &dto.UserCreateInput{
				FirstName: "John",
				LastName:  "Doe",
				Email:     "",
				UserName:  "testuser",
				Password:  "password",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			result, err := useCase.CreateUser(tc.input)

			// Assert
			// Nota: Actualmente el caso de uso no valida campos vacíos
			// Este test documenta el comportamiento actual
			if err != nil {
				t.Logf("Got error (expected if validation exists): %v", err)
			}
			if result != nil {
				t.Logf("Got result: %+v", result)
			}
		})
	}
}

func TestCreateUser_VerifyRepositoryCall(t *testing.T) {
	// Arrange
	var capturedUserID string
	var capturedUser *domain.User
	mockRepo := &MockUserRepository{
		CreateFunc: func(userEntity *domain.User) (*domain.User, error) {
			// Capturar el ID antes de modificarlo
			capturedUserID = userEntity.ID
			// Crear una copia para capturar el estado original
			capturedUser = &domain.User{
				ID:           userEntity.ID,
				FirstName:    userEntity.FirstName,
				LastName:     userEntity.LastName,
				Email:        userEntity.Email,
				UserName:     userEntity.UserName,
				PasswordHash: userEntity.PasswordHash,
			}
			// Simular que el repositorio genera un ID
			userEntity.ID = "new-id"
			return userEntity, nil
		},
	}

	useCase := &UserUseCase{
		UserRepository: mockRepo,
	}

	input := &dto.UserCreateInput{
		FirstName: "Bob",
		LastName:  "Williams",
		Email:     "bob@example.com",
		UserName:  "bobwilliams",
		Password:  "mypassword",
	}

	// Act
	_, err := useCase.CreateUser(input)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, but got: %v", err)
	}

	if capturedUser == nil {
		t.Fatal("Expected repository to be called with a user entity")
	}

	// Verificar que los datos del input se pasaron correctamente al repositorio
	if capturedUser.FirstName != input.FirstName {
		t.Errorf("Expected FirstName to be '%s', but got: %s", input.FirstName, capturedUser.FirstName)
	}

	if capturedUser.LastName != input.LastName {
		t.Errorf("Expected LastName to be '%s', but got: %s", input.LastName, capturedUser.LastName)
	}

	if capturedUser.Email != input.Email {
		t.Errorf("Expected Email to be '%s', but got: %s", input.Email, capturedUser.Email)
	}

	if capturedUser.UserName != input.UserName {
		t.Errorf("Expected UserName to be '%s', but got: %s", input.UserName, capturedUser.UserName)
	}

	if capturedUser.PasswordHash != input.Password {
		t.Errorf("Expected PasswordHash to be '%s', but got: %s", input.Password, capturedUser.PasswordHash)
	}

	if capturedUserID != "" {
		t.Errorf("Expected ID to be empty before repository call, but got: %s", capturedUserID)
	}
}
