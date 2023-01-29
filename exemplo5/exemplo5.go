package exemplo5

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func handlerHello(w http.ResponseWriter, r *http.Request) {
	query, err := url.ParseQuery(r.URL.RawQuery)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid request")
		return
	}
	name := query.Get("name")
	if name == "" {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello world!")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello %s", strings.ToUpper(name))
}
