package cmd

import (
	"github.com/Mansub/CRUD_SERVER/configs"
	"github.com/Mansub/CRUD_SERVER/network"
	"github.com/Mansub/CRUD_SERVER/repository"
	"github.com/Mansub/CRUD_SERVER/service"
)

type Cmd struct {
	config     *configs.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filepath string) *Cmd {
	c := &Cmd{
		config: configs.NewConfig(filepath),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
