package webapi

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Building a web API
Testing a web API
JSON encoding/decoding
Exploring godoc
Extending existing packages or
Shelling out

add a web API in front of godoc so its data can be queried via HTTP/JSON
*/
type API struct {
}

const (
	//ADDR = "127.0.0.1:9999"
	ADDR = ":9999"
)

var server *http.Server = &http.Server{
	Addr: ADDR,
}

//	func init() {
//		//httpHandlerSetup()
//	}
func (api *API) Run() {
	httpHandlerSetup()
}

func httpHandlerSetup() {
	http.HandleFunc("/", handler_start)

	idleConnectionClosed := make(chan struct{})

	go func() {
		fmt.Printf("\n[%v] listen to os signal", time.Now().Local().Format("2006-01-02 15:04:05"))
		sigint := make(chan os.Signal, 1)
		///signal.Notify(sigint, os.Interrupt)
		signal.Notify(sigint, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-sigint // block here until here OS signal
		fmt.Printf("\n[%v] received an interrupt siginal", time.Now().Local().Format("2006-01-02 15:04:05"))

		// wait 5 sec then shutdown server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := server.Shutdown(ctx); err != nil {
			log.Fatalf("\n[%v] server shutdonw error %v", time.Now().Local().Format("2006-01-02 15:04:05"), err)
		}

		// directly shutdown server
		//if err := server.Shutdown(context.Background()); err != nil {
		//	log.Fatalf("server shutdonw error %v", err)
		//}

		// after close all connection then close idleConnectionClosed
		fmt.Printf("\n[%v] idle connections closed", time.Now().Local().Format("2006-01-02 15:04:05"))
		close(idleConnectionClosed)
	}()

	if serverError := server.ListenAndServe(); serverError != nil && serverError != http.ErrServerClosed {
		log.Fatalf("ListenAndServe %v", serverError)
	}

	// block until idleConnectionClosed closed
	<-idleConnectionClosed
	fmt.Printf("\n[%v] server closed", time.Now().Local().Format("2006-01-02 15:04:05"))
	//if serverError := http.ListenAndServe(":9999", nil); serverError != nil && serverError != http.ErrServerClosed {
	//	log.Fatal(serverError)
	//}

	//log.Fatal(http.ListenAndServe(":9999", nil))
}
func handler_start(w http.ResponseWriter, r *http.Request) {

	fmt.Printf("\n[%v] start", time.Now().Local().Format("2006-01-02 15:04:05"))
	time.Sleep(time.Second * 2)
	fmt.Fprintf(w, "Hello world")
}
