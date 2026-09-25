//go:build wireinject
// +build wireinject

package container

import (
	"github.com/google/wire"
)

func InitializeContainer(opt persistence.AOFProviderOption) (*Container, func(), error) {
	wire.Build(
		storage.NewStore,
		persistence.NewAOFProvider,
		protocol.NewParser,
		usecase.NewStats,
		usecase.NewCommandHandler,
		handler.NewTCPHandler,
		NewContainer,
	)
	return nil, nil, nil
}
