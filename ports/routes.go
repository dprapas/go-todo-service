package ports

import (
	"github.com/dprapas/go-todo-service/app"
	log "github.com/sirupsen/logrus"
	"net/http"

)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func CreateRoutes(todoSvc app.TodoSvc, l *log.Logger) Routes {

	todoHandler := NewTodoHandler(todoSvc, l)

	var routes = Routes{
		Route{
			"Create",
			"POST",
			"/api/todos",
			todoHandler.Create,
		}, {
			"Health",
			"GET",
			"/health",
			todoHandler.Health,
		},
	}
	return routes
}
