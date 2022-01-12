package app

import (
	"context"
	"github.com/dprapas/go-todo-service/domain"
	log "github.com/sirupsen/logrus"
)

type TodoSvc struct {
	l   *log.Logger
	ctx context.Context
}

func NewTodoSvc(l *log.Logger, ctx context.Context) TodoSvc {
	return TodoSvc{
		l:   l,
		ctx: ctx,
	}
}

func (svc TodoSvc) Create(description string, owner string, finished bool) error {
	todo := domain.Todo{
		Description: description,
		Owner:       owner,
		Finished:    finished,
	}
	log.Infof("todo item created with description %s and finished status %t", todo.Description, todo.Finished)
	return nil
}
