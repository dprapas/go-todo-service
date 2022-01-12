package ports

import (
	"context"
	"github.com/dprapas/go-todo-service/app"
	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func InitRouter(l *log.Logger) *mux.Router {
	ctx := context.Background()

	todoSvc := app.NewTodoSvc(l, ctx)
	return NewRouter(todoSvc, l)
}
