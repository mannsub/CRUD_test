package service

import (
	"sync"

	"github.com/Mansub/CRUD_SERVER/repository"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

type Service struct {
	// repository
	repository *repository.Repository

	User *User
}

func NewService(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
