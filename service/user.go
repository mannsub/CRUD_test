package service

import (
	"github.com/Mansub/CRUD_SERVER/repository"
	"github.com/Mansub/CRUD_SERVER/types"
)

type User struct {
	UserRepository *repository.UserRepository
}

func newUserService(userRepository *repository.UserRepository) *User {
	return &User{
		UserRepository: userRepository,
	}
}

func (u *User) Create(newUser *types.User) error {
	return u.UserRepository.Create(newUser)
}

func (u *User) Update(name string, newAge int64) error {
	return u.UserRepository.Update(name, newAge)
}

func (u *User) Delete(user *types.User) error {
	return u.UserRepository.Delete(user.Name)
}

func (u *User) Get() []*types.User {
	return u.UserRepository.Get()
}
