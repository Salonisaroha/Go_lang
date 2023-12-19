package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// create http request
	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/users", nil)
	if err != nil {
		fmt.Printf("error fetching users")
	}

	//perform the http request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("error processing request")
	}

	defer res.Body.Close()

	//get body from the response
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error reading response body")
	}

	fmt.Printf("data of size %d\n", len(data))
}
