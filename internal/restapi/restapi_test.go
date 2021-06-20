package restapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerCreateMethodFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/create", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Create)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// func TestHandlerCreateDecodeBody(t *testing.T) {
// 	req, err := http.NewRequest("POST", "/create", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(Create)

// 	handler.ServeHTTP(rr, req)
// }

func TestHandlerGetOneMethodFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-one", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetOne)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestHandlerGetAllMethodFail(t *testing.T) {
	req, err := http.NewRequest("GET", "/get-all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAll)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
