package server

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Request struct {
	Input string `json:"inputString"`
}

type Response struct {
	Output string `json:"outputString"`
}

const VERSION = "version: 1.0.0"

func VersionHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			return
		}

		return
	}
	fmt.Println(VERSION)
	_, err := w.Write([]byte(VERSION))
	if err != nil {
		return
	}
}

func DecodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			return
		}
		return
	}
	d, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "ReadAll - Error", http.StatusBadRequest)
		return
	}
	var req Request
	err = json.Unmarshal(d, &req)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unmarshal - Error", http.StatusBadRequest)
		return
	}
	decodeString, err := base64.StdEncoding.DecodeString(req.Input)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid base64", http.StatusBadRequest)
	}
	var res Response
	res.Output = string(decodeString)
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		return
	}

	fmt.Println(req)

}

func HardopHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Error:", http.StatusMethodNotAllowed)
		_, err := fmt.Fprintf(w, "%s\n", "Method not allowed!")
		if err != nil {
			return
		}
		return
	}
	var mas []int
	for i := 500; i < 527; i++ {
		mas = append(mas, i)
	}
	var ind = rand.Intn(27)
	duration := time.Duration(rand.Intn(11)+10) * time.Second
	time.Sleep(duration)
	var num = rand.Intn(2)
	if num == 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(mas[ind])
	}
}
