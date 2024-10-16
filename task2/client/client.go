package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"task2/task2/server"
	"time"
)

func VersionRequest() {
	request, err := http.NewRequest(http.MethodGet, "http://localhost:8080/version", nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	httpData, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	d, err := io.ReadAll(httpData.Body)
	if err != nil {
		fmt.Println("Error in ReadAll():", err)
	}
	fmt.Println(string(d))
}

func DecodeRequest(base64 string) {
	var req = server.Request{Input: base64}
	jsonBody, _ := json.Marshal(req)
	request, err := http.NewRequest(http.MethodPost, "http://localhost:8080/decode", bytes.NewBuffer(jsonBody))
	if err != nil {
		fmt.Println("Get:", err)
		return
	}

	httpData, err := http.DefaultClient.Do(request)

	if err != nil {
		fmt.Println("Error in Do():", err)
		return
	}

	d, err := io.ReadAll(httpData.Body)
	if err != nil {
		fmt.Println("Error in ReadAll():", err)
	}

	var res server.Response
	err = json.Unmarshal(d, &res)

	fmt.Println(res.Output)
}

func HardopRequest() {
	ctx, can := context.WithTimeout(context.Background(), 15*time.Second)
	defer can()
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8080/hard-op", nil)
	if err != nil {
		fmt.Println("Get:", err)
		return
	}
	httpData, err := http.DefaultClient.Do(request)

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			fmt.Println("false")
			return
		}
		fmt.Println("Error in Do():", err)

	}

	fmt.Println("true, " + strconv.Itoa(httpData.StatusCode))

}
func main() {
	VersionRequest()
	DecodeRequest("aGVsbG8gZ29sYW5nIQ==")
	HardopRequest()

}
