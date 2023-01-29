package exemplo6

import (
	"io"
	"net/http"
)

func getHelloWorldURL(url string, nome string) (string, error) {
	res, err := http.Get(url + "/?nome=" + nome)
	if err != nil {
		return "", err
	}
	out, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(out), nil

}
