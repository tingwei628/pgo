package main

import (
	_ "github.com/tingwei628/pgo/webapi/docs"
	"github.com/tingwei628/pgo/webapi/internal/database"
	"github.com/tingwei628/pgo/webapi/internal/service"
	"github.com/tingwei628/pgo/webapi/internal/transport"
	"log"
)

//	@title			web api
//	@version		1.0
//	@description	web api in pgo
//	@termsOfService
//	@contact.name	API Support
//	@contact.url
//	@contact.email
//
// @license.name
// @license.url
// @host		localhost:8080
// @BasePath	/
func main() {
	db, err := database.NewDB("postgres")
	if err != nil {
		log.Fatal(err)
	}
	svc := service.NewTodoService(db)
	server := transport.NewServer(svc)
	err = server.Serve()
	if err != nil {
		log.Fatal(err)
	}
}
