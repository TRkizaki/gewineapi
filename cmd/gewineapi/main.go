package main

import (
	"net/http"

	"github.com/trkizaki/gewineapi/controller"
	"github.com/trkizaki/gewineapi/model/repository"
)

var tr = repository.NewWinelistRepository()
var tc = controller.NewWinelistController(tr)
var ro = controller.NewRouter(tc)

func main() {
	server := http.Server{
		Addr: ":8080", 
	}
	http.HandleFunc("/winelists/", ro.HandleWinelistsRequest)
	server.ListenAndServe()
}