package users

import "context"

// Repo defines the DB level interaction of users
type Repo interface {
	Get(ctx context.Context, id string) (User, error)
	GetUserBalance(ctx context.Context, id string) (float64, error)
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
}

// Service defines the business logic of users
type Service interface {
	Get(ctx context.Context, id string) (User, error)
	GetBalance(ctx context.Context, id string) (float64, error)
	GetAll(ctx context.Context, limit, offset int) ([]User, error)
	Create(ctx context.Context, user User) (string, error)
}

type user struct {
	repo Repo
}

func NewUserService(repo Repo) Service {
	return &user{
		repo: repo,
	}
}

func (u *user) Get(ctx context.Context, id string) (User, error) {
	return u.repo.Get(ctx, id)
}

func (u *user) GetBalance(ctx context.Context, id string) (float64, error) {
	return u.repo.GetUserBalance(ctx, id)
}

func (u *user) GetAll(ctx context.Context, limit, offset int) ([]User, error) {
	return u.repo.GetAll(ctx, limit, offset)
}

func (u *user) Create(ctx context.Context, user User) (string, error) {
	// check if user already exists
	_, err := u.repo.Get(ctx, user.Id)
	if err == nil {
		return "", ErrUserAlreadyExists
	}
	// check negative balance
	if user.Balance < 0 {
		return "", ErrNegativeBalance
	}
	// set user version to 1
	user.Version = 1
	return u.repo.Create(ctx, user)
}
