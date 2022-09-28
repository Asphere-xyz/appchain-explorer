package container

import (
	"github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/Ankr-network/ankr-protocol/shared/gateway"
	"github.com/Ankr-network/ankr-protocol/shared/staking"
	"github.com/Ankr-network/ankr-protocol/shared/websocket"
	log "github.com/sirupsen/logrus"
	"go.uber.org/dig"
)

type IContainer interface {
	Provide(constructor interface{}, opts ...dig.ProvideOption) error
	Invoke(function interface{}, opts ...dig.InvokeOption) error
	MustProvide(constructor interface{}, opts ...dig.ProvideOption) interface{}
	MustInvoke(function interface{}, opts ...dig.InvokeOption)
}

type defaultDigContainer struct {
	di *dig.Container
}

func (c *defaultDigContainer) Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return c.di.Provide(constructor, opts...)
}

func (c *defaultDigContainer) Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return c.di.Invoke(function, opts...)
}

func (c *defaultDigContainer) MustProvide(constructor interface{}, opts ...dig.ProvideOption) interface{} {
	var result interface{}
	must(c.di.Provide(constructor, opts...))
	return result
}

func (c *defaultDigContainer) MustInvoke(function interface{}, opts ...dig.InvokeOption) {
	must(c.di.Invoke(function, opts...))
}

func newDefaultEmptyContainer() IContainer {
	di := &defaultDigContainer{di: dig.New()}
	di.MustProvide(func() IContainer {
		return di
	})
	di.MustProvide(common.NewLogger)
	di.MustProvide(common.NewViper)
	di.MustProvide(common.NewViperConfigProvider)
	di.MustProvide(common.NewEmitter)
	return di
}

func NewContainer() IContainer {
	di := newDefaultEmptyContainer()
	di.MustProvide(database.NewService)
	di.MustProvide(database.NewStateDb)
	di.MustProvide(gateway.NewServer)
	di.MustProvide(websocket.NewServer)
	di.MustProvide(staking.NewService)
	return di
}

func MustStartDefault(di IContainer) {
	di.MustInvoke(func(s *database.Service) error { return di.Invoke(s.Start) })
	di.MustInvoke(func(s *database.StateDb) error { return di.Invoke(s.Start) })
	di.MustInvoke(func(s *gateway.Server) error { return di.Invoke(s.Start) })
	di.MustInvoke(func(s *websocket.Server) error { return di.Invoke(s.Start) })
	di.MustInvoke(func(s *staking.Service) error { return di.Invoke(s.Start) })
}

func must(err error) {
	if err != nil {
		log.Fatalf("failed to initialize DI: %s", err)
	}
}
