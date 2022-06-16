package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offset"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offset"`
}

type ConsumeResponse struct {
	Record Record `json:"record"`
}

func NewHttpServer(address string) *http.Server {
	server := newHttpServer()
	router := mux.NewRouter()
	router.HandleFunc("/", server.handleProduce).Methods("POST")
	router.HandleFunc("/", server.handleConsume).Methods("GET")

	return &http.Server{
		Addr:    address,
		Handler: router,
	}
}

func newHttpServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}
