package repository

import (
	"bank-ina-test/domain"
	"context"
	"fmt"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) Login(user *domain.Login) (string, error) {
	var getData *domain.User

	result := r.DB.First(&getData, "email = ?", user.Email)
	err := result.Error
	if err != nil {
		return "", fmt.Errorf("email not found")
	}

	// verifPass = VerifyPassword(getData.Password, user.Password)
	if !VerifyPassword(getData.Password, user.Password) {
		return "", fmt.Errorf("email and password does not match")
	}

	tokens, err := CreateToken(getData)
	if err != nil {
		return "", err
	}

	return tokens, nil
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	var userData = user
	userData.Password = HashPassword(user.Password)
	return r.DB.Create(userData).Error
}

func (r *UserRepository) GetAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	if err := r.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uint) (*domain.User, error) {
	var user domain.User
	if err := r.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	result := r.DB.Model(&user).Updates(domain.User{Email: user.Email, Name: user.Name})
	return result.Error
}

func (r *UserRepository) Delete(ctx context.Context, id uint) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
