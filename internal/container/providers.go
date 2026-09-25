package container

import (
	"github.com/lucasmilhoranca/kvlike/internal/adapter/handler"
	"github.com/lucasmilhoranca/kvlike/internal/adapter/protocol"
	"github.com/lucasmilhoranca/kvlike/internal/domain/repository"
	"github.com/lucasmilhoranca/kvlike/internal/usecase"
)

type Container struct {
	Store          repository.KeyValueRepository
	Persistence    repository.PersistenceRepository
	CommandHandler *usecase.CommandHandler
	TCPHandler     *handler.TCPHandler
	Parser         *protocol.Parser
}

func NewContainer(
	store repository.KeyValueRepository,
	persist repository.PersistenceRepository,
	parser *protocol.Parser,
	commandHandler *usecase.CommandHandler,
	tcpHandler *handler.TCPHandler,
) *Container {
	return &Container{
		Store:          store,
		Persistence:   persist,
		CommandHandler: commandHandler,
		TCPHandler:     tcpHandler,
		Parser:         parser,
	}
}

func (c *Container) Close() error {
	if c.Persistence != nil {
		return c.Persistence.Close()
	}
	return nil
}