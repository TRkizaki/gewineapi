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
	tc WinelistController
}

// Router constructor. The argument passed WinelistController,return Router struct pointer
func NewRouter(tc WinelistController) Router {
	return &router{tc}
}

func (ro *router) HandleWinelistsRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ro.tc.GetWinelists(w, r)
	case "POST":
		ro.tc.PostWinelists(w, r)
	case "PUT":
		ro.tc.PutWinelist(w, r)
	case "DELETE":
		ro.tc.DeleteWinelist(w, r)
	default:
		w.WriteHeader(405)
	}
}