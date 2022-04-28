package controller

import (
	"encoding/json"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/TRkizaki/gewineapi/controller/dto"
	"github.com/TRkizaki/gewineapi/test"
)

func TestGetWinelists_NotFound(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/winelists/", nil)

	target := NewWinelistController(&test.MockTodoRepository{})
	target.GetWinelists(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var winelistsResponse dto.WinelistsResponse
	json.Unmarshal(body, &winelistsResponse)
	if len(winelistsResponse.Winelists) != 0 {
		t.Errorf("Response is %v", winelistsResponse.Winelists)
	}
}

func TestGetWinelists_ExistWinelist(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/winelist/", nil)

	target := NewWinelistController(&test.MockWInelistRepositoryGetWinelistsExist{})
	target.GetWinelists(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	body := make([]byte, w.Body.Len())
	w.Body.Read(body)
	var winelistsResponse dto.WinelistsResponse
	json.Unmarshal(body, &winelistsResponse.Winelists)
	if len(winelistsResponse.Winelists) != 2 {
		t.Errorf("Response is %v", winelistsResponse.Winelists)
	}
}

func TestGetWinelists_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/winelists/", nil)

	target := NewWinelistController(&test.MockWinelistRepositoryError{})
	target.GetWinelists(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Content-Type") != "" {
		t.Errorf("Content-Type is %v", w.Header().Get("Content-Type"))
	}

	if w.Body.Len() != 0 {
		t.Errorf("body is %v", w.Body.Len())
	}
}

func TestPostWinelist_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","content":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/winelists/", json)

	target := NewWinelistController(&test.MockWinelistRepository{})
	target.PostWinelist(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != r.Host+r.URL.Path+"2" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPostWinelist_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/winelists/", json)

	target := NewWinelistController(&test.MockWinelistRepositoryError{})
	target.PostWinelist(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
	if w.Header().Get("Location") != "" {
		t.Errorf("Location is %v", w.Header().Get("Location"))
	}
}

func TestPutWInelist_OK(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/winelists/2", json)

	target := NewWinelistController(&test.MockWinelistRepository{})
	target.PutWinelist(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutWinelist_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/winelists/", nil)

	target := NewWinelistController(&test.MockWinelistRepository{})
	target.PutWinelist(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutWinelist_Error(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","contents":"test-content"}`)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/winelists/2", json)

	target := NewWinelistController(&test.MockWinelistRepositoryError{})
	target.PutWinelist(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteWinelist_OK(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/winelists/2", nil)

	target := NewWinelistController(&test.MockWinelistRepository{})
	target.DeleteWinelist(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteWinelist_InvalidPath(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/winelists/", nil)

	target := NewWinelistController(&test.MockWinelistRepositoryError{})
	target.DeleteWinelist(w, r)

	if w.Code != 400 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteWinelist_Error(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/winelists/2", nil)

	target := NewWinelistController(&test.MockWinelistRepositoryError{})
	target.DeleteWinelist(w, r)

	if w.Code != 500 {
		t.Errorf("Response cod is %v", w.Code)
	}
}