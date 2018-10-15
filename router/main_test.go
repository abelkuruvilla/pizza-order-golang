package router

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"pizza-delivery/controller"
	"testing"
)

func TestRouter(t *testing.T) {
	var jsonStr = []byte(`{"name":"Abel","email":"abel@sd.vom","mobile":8089352216,"veg":1,"meat":2}`)
	req, err := http.NewRequest("POST", "http://localhost:8010/buy_pizza", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	h := controller.Handler{}
	handler := http.Handler(responseHandler(h.GetStatus))
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusAccepted {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusAccepted)
	}

	expected := `{
		"data": "Successfully placed order"
	}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
