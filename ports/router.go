package ports

import (
	"github.com/dprapas/go-todo-service/app"
	log "github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter(todoSvc app.TodoSvc, l *log.Logger) *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	routes := CreateRoutes(todoSvc, l)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	return router
}
