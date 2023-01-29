package exemplo6

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetHelloWorldURL(t *testing.T) {

	expected := "Hello world!"
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer srv.Close()

	res, err := getHelloWorldURL(srv.URL, "")
	if err != nil {
		t.Errorf("Error expected to be nil, got %q", err)
	}
	if string(res) != expected {
		t.Errorf("Expected %s, got %s", expected, string(res))
	}
}
