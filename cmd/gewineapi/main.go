package main

import (
	"net/http"

	"github.com/TRkizaki/gewineapi/controller"
	"github.com/TRkizaki/gewineapi/model/repository"
)

var wr = repository.NewWinelistRepository()
var wc = controller.NewWinelistController(wr)
var ro = controller.NewRouter(wc)

func main() {
	server := http.Server{
		Addr: ":8080", 
	}
	http.HandleFunc("/winelists/", ro.HandleWinelistsRequest)
	server.ListenAndServe()
}