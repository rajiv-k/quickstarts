package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"time"
)

const (
	defaultPort = "9000"
)

var (
	hostID         string
	unixSocketPath string
)

type Resp struct {
	StatusCode     int                 `json:"status_code"`
	Method         string              `json:"method"`
	URI            string              `json:"uri"`
	Client         string              `jsont:"client"`
	HostID         string              `json:"host_id"`
	RequestHeaders map[string][]string `json:"request_headers"`
}

func main() {
	flag.StringVar(&unixSocketPath, "unix", "", "unix socket path")
	flag.Parse()

	hostID = os.Getenv("HOST_ID")

	router := &http.ServeMux{}
	router.HandleFunc("/echo", echoHandler)

	listener, err := createListener(unixSocketPath)
	if err != nil {
		return
	}

	fmt.Printf("listening on %v\n", listener.Addr())
	if err := http.Serve(listener, router); err != nil {
		fmt.Printf("ERROR: could not start server: %v\n", err)
	}
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%v %v %v %v\n", time.Now().UTC(), r.Method, r.RequestURI, r.RemoteAddr)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resp := Resp{
		StatusCode:     http.StatusOK,
		Method:         r.Method,
		URI:            r.RequestURI,
		Client:         r.RemoteAddr,
		HostID:         hostID,
		RequestHeaders: r.Header,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		fmt.Printf("ERROR: could not encode response: %v\n", err)
	}
}

func createListener(unixSocketPath string) (net.Listener, error) {
	if unixSocketPath != "" {
		unixListener, err := net.Listen("unix", unixSocketPath)
		if err != nil {
			fmt.Printf("ERROR: could not create unix socket: %v\n", err)
			return nil, err
		}
		return unixListener, nil
	} else {
		port := os.Getenv("PORT")
		if port == "" {
			port = defaultPort
		}
		tcpListener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
		if err != nil {
			fmt.Printf("ERROR: could not create tcp listener: %v\n", err)
			return nil, err
		}
		return tcpListener, err
	}
}
