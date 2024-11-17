package transport

import (
	"github.com/swaggo/http-swagger/v2"
	"github.com/tingwei628/pgo/webapi/internal/entity"
	"github.com/tingwei628/pgo/webapi/internal/service"
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

func NewServer(svc *service.TodoService) *Server {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))
	mux.HandleFunc("GET /search", TodoSearch(svc))
	mux.HandleFunc("GET /todo", TodoGET(svc))
	mux.HandleFunc("POST /todo", TodoPOST(svc))

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

// TodoPOST
//
//	@Summary	add todo
//	@Description
//	@Tags		todo
//	@Accept		json
//	@Produce	json
//	@Param		todo	body	TodoItem	true	"todo item"
//	@Success	200
//	@Failure	400	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/todo [post]
func TodoPOST(service *service.TodoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var t TodoItem
		ReadRequest(w, r, &t)
		err := service.Add(entity.Item{Task: t.Item})
		WriteResponseIfErrorExists(w, http.StatusBadRequest, err)
		w.WriteHeader(http.StatusCreated)
	}
}

// TodoGET
//
//	@Summary	get all todos
//	@Description
//	@Tags		todo
//	@Accept		json
//	@Produce	json
//	@Success	200	{array}		entity.Item
//	@Failure	400	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/todo [get]
func TodoGET(service *service.TodoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items, err := service.GetAll()
		WriteResponseIfErrorExists(w, http.StatusInternalServerError, err)
		WriteJsonResponse(w, items)
	}
}

// TodoSearch
//
//	@Summary	search todo by keyword
//	@Description
//	@Tags		todo
//	@Accept		json
//	@Produce	json
//	@Param		q	query		string	true	"keyword"
//	@Success	200	{array}		entity.Item
//	@Failure	400	{object}	ErrorResponse
//	@Failure	500	{object}	ErrorResponse
//	@Router		/search [get]
func TodoSearch(service *service.TodoService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keyword := r.URL.Query().Get("q")
		if strings.TrimSpace(keyword) == "" {
			errorResponse := ErrorResponse{
				Message: "q is required not empty",
			}
			errorResponse.Write(w, http.StatusBadRequest)
			return
		}
		items, err := service.Search(keyword)
		WriteResponseIfErrorExists(w, http.StatusInternalServerError, err)
		WriteJsonResponse(w, items)
	}
}
