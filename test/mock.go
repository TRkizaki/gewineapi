package test

import (
	"errors"
	"net/http"

	"github.com/TRKizaki/gewineapi/model/entity"
)

type MockWinelistController struct {
}

func (mtc *MockWinelistController) GetWinelists(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (mtc *MockWinelistController) PostWinelist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(201)
}

func (mtc *MockWinelistController) PutWinelist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func (mtc *MockWinelistController) DeleteWinelist(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

type MockWinelistRepository struct {
}

func (mtr *MockWinelistRepository) GetWinelists() (winelists []entity.WinelistEntity, err error) {
	winelists = []entity.WinelistEntity{}
	return
}

func (mtr *MockWinelistRepository) InsertWinelist(winelist entity.WinelistEntity) (id int, err error) {
	id = 2
	return
}

func (mtr *MockWinelistRepository) UpdateWinelist(winelist entity.WinelistEntity) (err error) {
	return
}

func (mtr *MockWinelistRepository) DeleteWinelist(id int) (err error) {
	return
}

type MockWinelistRepositoryGetWinelistsExist struct {
}

func (mtrgex *MockWinelistRepositoryGetWinelistsExist) GetWinelists() (winelists []entity.WinelistEntity, err error) {
	winelists = []entity.TodoEntity{}
	winelists = append(winelists, entity.WinelistEntity{Id: 1, Title: "title1", Brand: "brands1", Price: "price1"})
	winelists = append(winelists, entity.WinelistEntity{Id: 2, Title: "title2", Brand: "brands2", Price: "price2"})
	return
}

func (mtrgex *MockWinelistRepositoryGetWinelistsExist) InsertWinelist(winelist entity.WinelistEntity) (id int, err error) {
	return
}

func (mtrgex *MockWinelistRepositoryGetWinelistsExist) UpdateWinelist(winelist entity.WinelistEntity) (err error) {
	return
}

func (mtrgex *MockWinelistRepositoryGetWinelistsExist) DeleteWinelist(id int) (err error) {
	return
}

type MockWinelistRepositoryError struct {
}

func (mtrgtn *MockWinelistRepositoryError) GetWinelists() (winelists []entity.WinelistEntity, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgie *MockWinelistRepositoryError) InsertWinelist(winelist entity.WinelistEntity) (id int, err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgue *MockWinelistRepositoryError) UpdateWinelist(winelist entity.WinelistEntity) (err error) {
	err = errors.New("unexpected error occurred")
	return
}

func (mtrgde *MockWinelistRepositoryError) DeleteWinelist(id int) (err error) {
	err = errors.New("unexpected error occurred")
	return
}