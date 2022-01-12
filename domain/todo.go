package domain

type CreateTodoRequest struct {
	Description string `json:"description"`
	Owner       string `json:"owner"`
	Finished    bool   `json:"finished"`
}

type CreateTodoResponse struct {
	Message string `json:"message"`
}

type Todo struct {
	ID          string
	Description string
	Owner       string
	Finished    bool
}
