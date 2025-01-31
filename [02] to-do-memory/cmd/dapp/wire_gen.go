// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Mugen-Builders/to-do-memory/configs"
	"github.com/Mugen-Builders/to-do-memory/internal/domain"
	"github.com/Mugen-Builders/to-do-memory/internal/infra/cartesi/advance_handler"
	"github.com/Mugen-Builders/to-do-memory/internal/infra/cartesi/inspect_handler"
	"github.com/Mugen-Builders/to-do-memory/internal/infra/repository"
	"github.com/google/wire"
)

// Injectors from wire.go:

func NewAdvanceHandlers(db *configs.InMemoryDB) (*AdvanceHandlers, error) {
	toDoRepositoryInMemory := repository.NewToDoRepositoryInMemory(db)
	toDoAdvanceHandlers := advance_handler.NewToDoAdvanceHandlers(toDoRepositoryInMemory)
	advanceHandlers := &AdvanceHandlers{
		ToDoAdvanceHandlers: toDoAdvanceHandlers,
	}
	return advanceHandlers, nil
}

func NewInspectHandlers(db *configs.InMemoryDB) (*InspectHandlers, error) {
	toDoRepositoryInMemory := repository.NewToDoRepositoryInMemory(db)
	toDoInspectHandlers := inspect_handler.NewToDoInspectHandlers(toDoRepositoryInMemory)
	inspectHandlers := &InspectHandlers{
		ToDoInspectHandlers: toDoInspectHandlers,
	}
	return inspectHandlers, nil
}

// wire.go:

var setToDoRepositoryDependency = wire.NewSet(repository.NewToDoRepositoryInMemory, wire.Bind(new(domain.ToDoRepository), new(*repository.ToDoRepositoryInMemory)))

var setAdvanceHandlers = wire.NewSet(advance_handler.NewToDoAdvanceHandlers)

var setInspectHandlers = wire.NewSet(inspect_handler.NewToDoInspectHandlers)

type AdvanceHandlers struct {
	ToDoAdvanceHandlers *advance_handler.ToDoAdvanceHandlers
}

type InspectHandlers struct {
	ToDoInspectHandlers *inspect_handler.ToDoInspectHandlers
}
