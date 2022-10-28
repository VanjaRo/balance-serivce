package users

import "context"

// Repo defines the DB level interaction of users
type Repo interface {
	Get(ctx context.Context, id string) (User, error)
	GetUserBalance(ctx context.Context, id string) (float64, error)
	UpdateUserBalance(ctx context.Context, userId string, balanceDiff float64) error
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
}

// Service defines the business logic of users
type Service interface {
	Get(ctx context.Context, userId string) (User, error)
	GetBalance(ctx context.Context, userId string) (float64, error)
	UpdateUserBalance(ctx context.Context, userId string, balance float64) error
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, userId string, balance float64) (string, error)
}

type user struct {
	repo Repo
}

func NewUserService(repo Repo) Service {
	return &user{
		repo: repo,
	}
}

func (u *user) Get(ctx context.Context, userId string) (User, error) {
	return u.repo.Get(ctx, userId)
}

func (u *user) GetBalance(ctx context.Context, userId string) (float64, error) {
	return u.repo.GetUserBalance(ctx, userId)
}

func (u *user) GetAll(ctx context.Context, limit, offset int) ([]User, error) {
	return u.repo.GetAll(ctx, limit, offset)
}

func (u *user) Create(ctx context.Context, userId string, balance float64) (string, error) {
	// check empty string for userId
	if userId == "" {
		return "", ErrEmptyUserId
	}
	// check if user already exists
	_, err := u.repo.Get(ctx, userId)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	// check negative balance
	if balance < 0 {
		return "", ErrNegativeBalance
	}
	newUser := User{
		Id:      userId,
		Balance: balance,
		Version: 1,
	}

	return u.repo.Create(ctx, newUser)
}

func (u *user) UpdateUserBalance(ctx context.Context, userId string, balanceDiff float64) error {
	return u.repo.UpdateUserBalance(ctx, userId, balanceDiff)
}
