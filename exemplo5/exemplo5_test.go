package exemplo5

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlerHello(t *testing.T) {
	testcases := []struct {
		name      string
		inputName string
		expected  string
	}{
		{name: "when input is provided", inputName: "Eduardo", expected: "Hello EDUARDO"},
		{name: "when input is empty", inputName: "", expected: "Hello world!"},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {

			req := httptest.NewRequest(http.MethodGet, "/?name="+tc.inputName, nil)
			w := httptest.NewRecorder()
			handlerHello(w, req)
			res := w.Result()
			defer res.Body.Close()

			data, err := io.ReadAll(res.Body)
			if err != nil {
				t.Errorf("expected to be nil, got %q", err)
			}
			if string(data) != tc.expected {
				t.Errorf("Expected to be %s, got %s", tc.expected, string(data))
			}
		})
	}
}
