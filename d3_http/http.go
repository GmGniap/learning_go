package main

import (
	"fmt"
	"io"
	"net/http"
)

/*
http.Get uses the http.DefaultClient to make a request to the given url
res is the HTTP response that comes back from the server
defer res.Body.Close() ensures that the response body is properly closed after reading. Not doing so can cause memory issues.
io.ReadAll reads the response body into a slice of bytes []byte called data
*/
func getIssueData(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("request error: %w", err)
	}
	defer res.Body.Close()
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("reading resp error: %w", err)
	}
	return data, nil
}
