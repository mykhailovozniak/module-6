package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIsJSON(t *testing.T) {
	validJSONString := `{ "userId" : 1, "id" : 1, "title" : "sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto" }`

	result := IsJSON(validJSONString)

	if result != true {
		t.Errorf("Should return true as string valid json")
	}

	invalidJSONString := `some string`

	invalidJSONResult := IsJSON(invalidJSONString)

	if invalidJSONResult != false {
		t.Errorf("Should return false as string not valid json")
	}
}

func TestPostControllerBadInput(t *testing.T) {
	req, err := http.NewRequest("GET", "/post", nil)

	if err != nil {
		t.Errorf("Error during making call to /post")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("error")
	}
}

func TestPostControllerCachePost(t *testing.T) {
	req, err := http.NewRequest("GET", "/post?postId=1", nil)

	if err != nil {
		t.Errorf("Error during making call to /post?postId=1")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error because status is not code 200")
	}

	cached := rr.HeaderMap.Get("cached")

	if cached != "true" {
		t.Errorf("error because cached header is not true")
	}

	correctBody := `{"userId":1,"id":1,"title":"sunt aut facere repellat provident occaecati excepturi optio reprehenderit","body":"quia et suscipit\\nsuscipit recusandae consequuntur expedita et cum\\nreprehenderit molestiae ut ut quas totam\\nnostrum rerum est autem sunt rem eveniet architecto"}`
	body := rr.Body.String()

	if body != correctBody {
		t.Errorf("Body in response is not correct")
	}
}

func TestPostControllerExternalPost(t *testing.T) {
	req, err := http.NewRequest("GET", "/post?postId=4", nil)

	if err != nil {
		t.Errorf("Error during making call to /post?postId=4")
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostController)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("error because status is not code 200")
	}

	cached := rr.HeaderMap.Get("cached")

	if cached != "false" {
		t.Errorf("error because cached header is not false")
	}

	correctBody := `{"userId":1,"id":4,"title":"eum et est occaecati","body":"ullam et saepe reiciendis voluptatem adipisci\nsit amet autem assumenda provident rerum culpa\nquis hic commodi nesciunt rem tenetur doloremque ipsam iure\nquis sunt voluptatem rerum illo velit"}`
	body := rr.Body.String()

	if body != correctBody {
		t.Errorf("Body in response is not correct")
	}
}
