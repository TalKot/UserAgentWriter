package main

import (
	"fmt"
	"net/http"
	"os"

	ghandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//const - webserver port
const (
	WEBSERVERPORT = ":8000"
)

func main() {
	//setup server
	printLogo()
	setupConsumer()

	//setup routes
	r := mux.NewRouter()
	r.HandleFunc("/health-check", responseOK).Methods("GET")
	r.HandleFunc("/load", loadHandler).Methods("GET")

	//ghandlers.LoggingHandler - applying logging routes with middelware
	http.Handle("/", ghandlers.LoggingHandler(os.Stdout, r))

	//setup http server listener
	http.ListenAndServe(WEBSERVERPORT, nil)
}

func printLogo() {
		fmt.Println("")
  	fmt.Println("______      __   __ __      __ ")
 		fmt.Println("/_  __/___ _/ /  / //_/___  / /_")
  	fmt.Println("/ / / __ `/ /  / ,< / __ \/ __/")
 		fmt.Println("/ / / /_/ / /  / /| / /_/ / /_  ")
		fmt.Println("/_/  \__,_/_/  /_/ |_\____/\__/  ")
		fmt.Println("")
}
