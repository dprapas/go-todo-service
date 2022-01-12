package ports

import (
	"encoding/json"
	"fmt"
	"github.com/dprapas/go-todo-service/app"
	"github.com/dprapas/go-todo-service/domain"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type TodoHandler struct {
	todoSvc app.TodoSvc
	l       *log.Logger
}

func NewTodoHandler(todoSvc app.TodoSvc, l *log.Logger) TodoHandler {
	return TodoHandler{
		todoSvc: todoSvc,
		l:       l,
	}
}

func (h *TodoHandler) Create(w http.ResponseWriter, r *http.Request) {

	var req domain.CreateTodoRequest

	body, _ := ioutil.ReadAll(r.Body)

	if err := json.Unmarshal(body, &req); err != nil {
		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(JsonErr{
			Code: "BAD_REQUEST",
			Text: "Request body cannot be processed",
		})
		return
	}

	if len(req.Description) == 0 {
		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(JsonErr{
			Code: "BAD_REQUEST",
			Text: "All the mandatory fields username, password must be provided",
		})
		return
	}

	err := h.todoSvc.Create(req.Description, req.Owner, req.Finished)

	if err != nil {
		w.Header().Set(ContentType, ApplicationJson)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(JsonErr{
			Code: err.Error(),
			Text: "Error",
		})
		return
	}

	w.Header().Set(ContentType, ApplicationJson)
	w.WriteHeader(http.StatusOK)
}

func (h *TodoHandler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Healthy")
}
