package store

import (
	"context"

	"github.com/VanjaRo/balance-serivce/pkg/services/users"
	"github.com/VanjaRo/balance-serivce/pkg/utils/log"
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
		log.Info(ctx, "user not found")
		// TODO: add custom error return
		return users.User{}, users.ErrUserNotFound
	}
	// loggin to context

	return user, nil
}

// gets balance from user
func (u *userRepo) GetUserBalance(ctx context.Context, id string) (float64, error) {
	var userBalance float64
	err := u.DB.Model(&users.User{}).Where("id = ?", id).Select("balance").Scan(&userBalance).Error
	if err != nil {
		log.Info(ctx, "user not found")
		return 0, users.ErrUserNotFound
	}
	return userBalance, nil
}

func (u *userRepo) GetAll(ctx context.Context, limit, offset int) ([]users.User, error) {
	var usrs []users.User
	err := u.DB.Limit(limit).Offset(offset).Find(&usrs).Error
	if err != nil {
		log.Error(ctx, "error while getting all users")
		return []users.User{}, users.ErrUserQuery
	}
	return usrs, nil
}

func (u *userRepo) Create(ctx context.Context, user users.User) (string, error) {
	err := u.DB.Create(&user).Error
	if err != nil {
		log.Error(ctx, "error while creating user")
		return "", users.ErrUserCreate
	}
	log.Info(ctx, "user with id=%s created", user.Id)
	return user.Id, nil
}
