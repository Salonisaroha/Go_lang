package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	start := time.Now()
	ctx := context.WithValue(context.Background(), "username", "Saloni")
	userID, err := fetchUserID(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The response took %v -> %+v\n", time.Since(start), userID)

}

func fetchUserID(ctx context.Context) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Millisecond*100)
	defer cancel()
	val := ctx.Value("username")
	fmt.Println("The value is ", val)
	type result struct {
		userID string
		err    error
	}
	resultCh := make(chan result, 1)
	go func() {
		res, err := thirdpartyHTTPCall()
		resultCh <- result{
			userID: res,
			err:    err,
		}
	}()
	select {
	// Done()
	// 1 -> the context timeout is exceed
	// 2. -> the context has been manually cancelled -> cancel()
	case <-ctx.Done():
		return " ", ctx.Err()
	case res := <-resultCh:
		return res.userID, res.err
	}

}
func thirdpartyHTTPCall() (string, error) {
	time.Sleep(time.Millisecond * 500)
	return "user id 1", nil
}
