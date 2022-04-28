package controller

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"

	"github.com/TRkizaki/gewineapi/controller/dto"
	"github.com/TRkizaki/gewineapi/model/entity"
	"github.com/TRkizaki/gewineapi/model/repository"
)

type WinelistController interface {
	GetWinelist(w http.ResponseWriter, r *http.Request)
	PostWinelist(w http.ResponseWriter, r *http.Request)
	PutWinelist(w http.ResponseWriter, r *http.Request)
	DeleteWinelist(w http.ResponseWriter, r *http.Request)
}

type winelistController struct {
	wr repository.WinelistRepository
}

func NewWinelistController(wr repository.WinelistRepository) WinelistController {
	return &winelistController{wr}
}

func (wc *winelistController) GetTodos(w http.ResponseWriter, r *http.Request) {
	winelists, err := wc.wr.GetWinelists()
	if err != nil {
		w.WriteHeader(500)
		return
	}

	var winelistResponses []dto.WinelistResponse
	for _, v := range winelists {
		winelistResponses = append(winelistResponses, dto.WinelistResponse{Id: v.Id, Title: v.Title, Content: v.Content})
	}

	var winelistsResponse dto.WinelistsResponse
	winelistsResponse.Winelists = winelistResponses

	output, _ := json.MarshalIndent(winelistsResponse.Winelists, "", "\t\t")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
}

func (wc *winelistController) PostTodo(w http.ResponseWriter, r *http.Request) {
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var winelistRequest dto.WinelistRequest
	json.Unmarshal(body, &winelistRequest)

	winelist := entity.WinelistEntity{Title: winelistRequest.Title, Brand: winelistRequest.Brand}
	id, err := wc.wr.InsertTodo(winelist)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Location", r.Host+r.URL.Path+strconv.Itoa(id))
	w.WriteHeader(201)
}

func (tc *winelistController) PutTodo(w http.ResponseWriter, r *http.Request) {
	winelistId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	var winelistRequest dto.WinelistRequest
	json.Unmarshal(body, &winelistRequest)

	winelist := entity.WinelistEntity{Id: winelistId, Title: winelistRequest.Title, Brand: winelistRequest.Brand}
	err = wc.wr.UpdateWinelist(winelist)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}

func (wc *winelistController) DeleteWinelist(w http.ResponseWriter, r *http.Request) {
	winelistId, err := strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		w.WriteHeader(400)
		return
	}

	err = wc.wr.DeleteWinelist(winelistId)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(204)
}