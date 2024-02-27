package usecase

import (
	"bank-ina-test/domain"
	"bank-ina-test/repository"
	"context"
)

type UserUsecase struct {
	UserRepository repository.UserRepository
}

func (uc *UserUsecase) RegisterUser(ctx context.Context, user *domain.User) error {
	// Call repository method to create user
	return uc.UserRepository.Create(ctx, user)
}

func (uc *UserUsecase) LoginUser(ctx context.Context, user *domain.Login) (string, error) {
	// Call repository method to login user
	return uc.UserRepository.Login(user)
}

func (uc *UserUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	// Call repository method to get all users
	return uc.UserRepository.GetAll(ctx)
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	// Call repository method to get user by ID
	return uc.UserRepository.GetByID(ctx, id)
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, user *domain.User) error {
	// Call repository method to update user
	return uc.UserRepository.Update(ctx, user)
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, id uint) error {
	// Call repository method to delete user
	return uc.UserRepository.Delete(ctx, id)
}
