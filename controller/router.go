package controller

import (
	"net/http"
)

//  public interface for Extrinsics package
type Router interface {
	HandleWinelistsRequest(w http.ResponseWriter, r *http.Request)
}

// private Router struct
type router struct {
	wc WinelistController
}

// Router constructor. The argument passed WinelistController,return Router struct pointer
func NewRouter(tc WinelistController) Router {
	return &router{wc}
}

func (ro *router) HandleWinelistsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.wc.GetWinelists(w, r)
	case "POST":
		ro.wc.PostWinelists(w, r)
	case "PUT":
		ro.wc.PutWinelist(w, r)
	case "DELETE":
		ro.wc.DeleteWinelist(w, r)
	default:
		w.WriteHeader(405)
	}
}