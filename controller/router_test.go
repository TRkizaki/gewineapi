package controller

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/TRkizaki/gewineapi/test"
)

var mux *http.ServeMux

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	os.Exit(code)
}

func setUp() {
	target := NewRouter(&test.MockWinelistController{})
	mux = http.NewServeMux()
	mux.HandleFunc("/winelists/", target.HandleWinelistsRequest)

}

func TestGetWinelists(t *testing.T) {
	r, _ := http.NewRequest("GET", "/winelists/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPostWinelist(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","Brand":"test-brand","Price":"test-price"}`)
	r, _ := http.NewRequest("POST", "/winelists/", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 201 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestPutWinelist(t *testing.T) {
	json := strings.NewReader(`{"title":"test-title","Brand":"test-brand","Price":"test-price"}`)
	r, _ := http.NewRequest("PUT", "/winelists/2", json)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestDeleteWinelist(t *testing.T) {
	r, _ := http.NewRequest("DELETE", "/winelists/2", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 204 {
		t.Errorf("Response cod is %v", w.Code)
	}
}

func TestInvalidMethod(t *testing.T) {
	r, _ := http.NewRequest("PATCH", "/winelists/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, r)

	if w.Code != 405 {
		t.Errorf("Response cod is %v", w.Code)
	}
}