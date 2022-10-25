package users

import "context"

// Repo defines the DB level interaction of users
type Repo interface {
	Get(ctx context.Context, id string) (User, error)
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
	userDb, err := u.repo.Get(ctx, id)
	if err != nil {
		return 0, err
	}
	return userDb.Balance, nil
}

func (u *user) GetAll(ctx context.Context, limit, offset int) ([]User, error) {
	return u.repo.GetAll(ctx, limit, offset)
}

func (u *user) Create(ctx context.Context, user User) (string, error) {
	return u.repo.Create(ctx, user)
}
