package usecases

import (
	"errors"
	"workshop4-backend/domain/entities"
	"workshop4-backend/domain/repositories"
)

type UserUseCase struct {
	userRepo repositories.UserRepository
}

func NewUserUseCase(userRepo repositories.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) CreateUser(req *entities.UserCreateRequest) (*entities.User, error) {
	// Check if user with email already exists
	existingUser, _ := uc.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("user with this email already exists")
	}

	user := &entities.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Phone:       req.Phone,
		DateOfBirth: req.DateOfBirth,
		Address:     req.Address,
		City:        req.City,
		Country:     req.Country,
		PostalCode:  req.PostalCode,
		Avatar:      req.Avatar,
	}

	err := uc.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserByID(id uint) (*entities.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}

	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) GetAllUsers() ([]*entities.User, error) {
	users, err := uc.userRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (uc *UserUseCase) UpdateUser(id uint, req *entities.UserUpdateRequest) (*entities.User, error) {
	if id == 0 {
		return nil, errors.New("invalid user ID")
	}

	// Check if user exists
	existingUser, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	// Update only provided fields
	if req.FirstName != nil {
		existingUser.FirstName = *req.FirstName
	}
	if req.LastName != nil {
		existingUser.LastName = *req.LastName
	}
	if req.Email != nil {
		// Check if email is already taken by another user
		userWithEmail, _ := uc.userRepo.GetByEmail(*req.Email)
		if userWithEmail != nil && userWithEmail.ID != id {
			return nil, errors.New("email already taken by another user")
		}
		existingUser.Email = *req.Email
	}
	if req.Phone != nil {
		existingUser.Phone = *req.Phone
	}
	if req.DateOfBirth != nil {
		existingUser.DateOfBirth = *req.DateOfBirth
	}
	if req.Address != nil {
		existingUser.Address = *req.Address
	}
	if req.City != nil {
		existingUser.City = *req.City
	}
	if req.Country != nil {
		existingUser.Country = *req.Country
	}
	if req.PostalCode != nil {
		existingUser.PostalCode = *req.PostalCode
	}
	if req.Avatar != nil {
		existingUser.Avatar = *req.Avatar
	}

	err = uc.userRepo.Update(id, existingUser)
	if err != nil {
		return nil, err
	}

	return existingUser, nil
}

func (uc *UserUseCase) DeleteUser(id uint) error {
	if id == 0 {
		return errors.New("invalid user ID")
	}

	// Check if user exists
	_, err := uc.userRepo.GetByID(id)
	if err != nil {
		return err
	}

	return uc.userRepo.Delete(id)
}
