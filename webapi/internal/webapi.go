package internal

import (
	"github.com/tingwei628/pgo/webapi/internal/todo"
	"github.com/tingwei628/pgo/webapi/internal/transport"
	"log"
)

type API struct {
}

func (api *API) Run() {

	service := todo.NewTodoService()
	server := transport.NewServer(service)
	err := server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
