package repository

import (
	entity "main.go/bisnes_internal/entitise/domen/user"
)

type UserRepository interface {
	Save(user *entity.User) (*entity.User, error)
}

type InMemoryUserRepository struct {
	users []*entity.User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{users: []*entity.User{}}
}

func (r *InMemoryUserRepository) Save(user *entity.User) (*entity.User, error) {
	user.ID = len(r.users) + 1
	r.users = append(r.users, user)
	return user, nil
}
