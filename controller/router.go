package controller

import (
	"net/http"
)

// 外部パッケージ用public interface
type Router interface {
	HandleWinelistsRequest(w http.ResponseWriter, r *http.Request)
}

// 非公開のRouter struct
type router struct {
	tc WinelistController
}

// Routerのconstructor。argumentにWinelistControllerを受け取り、Router structのpointerをreturnする。
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