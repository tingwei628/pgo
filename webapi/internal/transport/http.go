package transport

import (
	"github.com/tingwei628/pgo/webapi/internal/todo"
	"log"
	"net/http"
	"strings"
)

type TodoItem struct {
	Item string `json:"item"`
}

type Server struct {
	mut *http.ServeMux
}

var service *todo.TodoService

func NewServer(service *todo.TodoService) *Server {

	mux := http.NewServeMux()

	mux.HandleFunc("/todo", func(w http.ResponseWriter, r *http.Request) {
		todoHandler(w, r, service)
	})

	mux.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		_todoSearch(w, r, service)
	})

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

	return &Server{
		mut: mux,
	}
}

func (s *Server) Serve() error {
	return http.ListenAndServe(":8080", s.mut)
}

func todoHandler(w http.ResponseWriter, r *http.Request, service *todo.TodoService) {
	switch r.Method {
	case http.MethodGet:
		_todoGET(w, r, service)
	case http.MethodPost:
		_todoPOST(w, r, service)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func _todoPOST(w http.ResponseWriter, r *http.Request, service *todo.TodoService) {
	var t TodoItem
	ReadRequest(w, r, &t)
	err := service.Add(todo.Item{Task: t.Item})
	WriteResponseIfErrorExists(w, http.StatusBadRequest, err)
	w.WriteHeader(http.StatusCreated)
}

func _todoGET(w http.ResponseWriter, r *http.Request, service *todo.TodoService) {
	WriteJsonResponse(w, service.GetAll())
}

func _todoSearch(w http.ResponseWriter, r *http.Request, service *todo.TodoService) {
	keyword := r.URL.Query().Get("q")
	if strings.TrimSpace(keyword) == "" {
		errorResponse := ErrorResponse{
			Message: "q is required not empty",
		}
		errorResponse.Write(w, http.StatusBadRequest)
		return
	}
	WriteJsonResponse(w, service.Search(keyword))
}
