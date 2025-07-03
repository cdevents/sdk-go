package examples

import (
	"log"
	"net/http"
)

// Copied from https://github.com/eswdd/go-smee/blob/33b0bac1f1ef3abef04c518ddf7552b04edbadd2/smee.go#L54C1-L67C2
func CreateSmeeChannel() (*string, error) {
	httpClient := http.Client{
		CheckRedirect: func(_ *http.Request, _ []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	resp, err := httpClient.Head("https://smee.io/new")
	if err != nil {
		return nil, err
	}

	loc := resp.Header.Get("Location")
	return &loc, nil
}

func PanicOnError(e error, message string) {
	if e != nil {
		log.Fatalf(message+", %v", e)
	}
}
