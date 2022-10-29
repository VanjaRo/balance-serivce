package store

import (
	"context"
	"os/user"

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
		log.Info(ctx, "user with id=%s not found", id)
		// TODO: add custom error return
		return users.User{}, users.ErrUserNotFound
	}
	// loggin to context

	return user, nil
}

// gets balance from user
func (u *userRepo) GetUserBalance(ctx context.Context, id string) (int, error) {

	var userBalance int
	err := u.DB.Model(&users.User{}).Select("balance").Where("id = ?", id).First(&userBalance).Error
	if err != nil {
		log.Info(ctx, "user with id=%s not found", id)
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

func (u *userRepo) UpdateUserBalance(ctx context.Context, userId string, balanceDiff int) error {
	//  update balance using optimistic locking
	for i := 0; i < 100; i++ {
		// get user
		user, err := u.Get(ctx, userId)
		if err != nil {
			return err
		}
		// check if balance is negative
		if user.Balance+balanceDiff < 0.0 {
			return users.ErrNegativeBalance
		}
		// update balance
		user.Balance += balanceDiff
		user.Version++
		result := u.DB.Model(&users.User{}).Where("id = ? AND version = ?", userId, user.Version-1).Updates(&user)
		if result.Error != nil {
			log.Error(ctx, "error while updating balance")
			return result.Error
		}
		// check if user was updated
		if result.RowsAffected == 0 {
			continue
		}
		return nil
	}
	return nil
}

func (u *userRepo) Migrate() error {
	return u.DB.AutoMigrate(&user.User{})
}
