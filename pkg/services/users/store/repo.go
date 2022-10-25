package store

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"gorm.io/gorm"
)

type userRepo struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) users.Repo {
	return &userRepo{
		DB: db,
	}
}

func (u *userRepo) Get(ctx context.Context, id string) (users.User, error) {
	var user users.User
	err := u.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		// TODO: loggin to context
		// TODO: add custom error return
		return user, err
	}
	// loggin to context
	return user, nil
}

func (u *userRepo) GetAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	var users []users.User
	err := u.DB.Limit(limit).Offset(offset).Find(&users).Error
	if err != nil {
		// TODO: loggin to context
		// TODO: add custom error return
		return users, err
	}
	return users, nil
}

func (u *userRepo) Create(ctx context.Context, user users.User) (string, error) {
	err := u.DB.Create(&user).Error
	if err != nil {
		// TODO: loggin to context
		// TODO: add custom error return
		return "", err
	}
	return user.Id, nil
}
