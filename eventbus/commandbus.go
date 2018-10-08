package eventbus

import (
	"VaScanGo/domain"
	"fmt"
	"sync"
)

type CommandBus struct {}

type CommandHandler interface {
	HandleCommand(cmd domain.Command) error
}

type CommandHandlerRegistry struct {
	sync.RWMutex
	registry map[string]CommandHandler
}

func NewCommandHandlerRegistry() *CommandHandlerRegistry {
	return &CommandHandlerRegistry{
		registry: make(map[string]CommandHandler),
	}
}

func (c *CommandHandlerRegistry) RegisterCommand(cmdType string, cmd CommandHandler) {
	c.Lock()
	defer c.Unlock()
	c.registry[cmdType] = cmd
}

func (c *CommandHandlerRegistry) GetCommand(cmdType string) (CommandHandler, error) {
	handler, ok := c.registry[cmdType]
	if !ok {
		return nil, fmt.Errorf("can't find %s in registry", cmdType)
	}
	return handler, nil
}

func (c *CommandBus) Handle(cmd domain.Command, cmdRegistry *CommandHandlerRegistry) error {
	cmdType := cmd.GetType()
	commandHandler, _ := cmdRegistry.GetCommand(cmdType)
	err := commandHandler.HandleCommand(cmd)
	if err != nil {
		return fmt.Errorf("error command execution: %s", err.Error())
	}
	return nil
}