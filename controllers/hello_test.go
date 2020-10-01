package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)
func TestHelloController(t *testing.T) {
	  req, err := http.NewRequest("GET", "/hello", nil)

	  if err != nil {
	  	t.Errorf("Error during call to /hello")
	  }

	  rr := httptest.NewRecorder()
	  handler := http.HandlerFunc(HelloController)

	  handler.ServeHTTP(rr, req)

	  if status := rr.Code; status != http.StatusOK {
	  	t.Errorf("error because status code not equal 200")
	  }

	  body := rr.Body.String()
	  correctBody := "Hello world"

	  if body != correctBody {
	  	t.Errorf("Body is not correct")
	  }
}
