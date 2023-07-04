package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "9000"
)

var hostID string

type Resp struct {
	StatusCode int    `json:"status_code"`
	Method     string `json:"method"`
	URI        string `json:"uri"`
	Client     string `jsont:"client"`
	HostID     string `json:"host_id"`
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	hostID = os.Getenv("HOST_ID")

	router := &http.ServeMux{}
	router.HandleFunc("/echo", echoHandler)
	fmt.Printf("listening on %v\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", port), router); err != nil {
		fmt.Printf("ERROR: could not start server: %v\n", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v %v %v\n", time.Now().UTC(), r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Resp{
		StatusCode: http.StatusOK,
		Method:     r.Method,
		URI:        r.RequestURI,
		Client:     r.RemoteAddr,
		HostID:     hostID,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("ERROR: could not encode response: %v\n", err)
	}
}
