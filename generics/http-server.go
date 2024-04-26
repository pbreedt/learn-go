package generics

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Superfluous Server for handling requests with payloads defined as regular Go structs
// Intented to illustrate the use of generics with methods and functions
type Server struct {
}

// Handler instantiated with T1 (request type) and T2 (response type)
// If special parsing is needed to convert the request to T1, RequestParser can be set, otherwise json.NewDecoder will be used
type Handler[T1, T2 any] struct {
	Pattern       string
	HandleRequest func(T1) T2 // Function to handle the request. T1 is the request type, T2 is the response type
}

func NewServer() *Server {
	return &Server{}
}

func (svr *Server) Run(addr string) {
	http.ListenAndServe(addr, nil)
}

// Cannot define method on Server since types T1 & T2 cannot be defined on input params of a method, below is invalid
// func (s *Server) AddHandler[T1, T2 any](h *Handler[T1, T2]) {}
// func (s *Server)[T1, T2 any] AddHandler(h *Handler[T1, T2]) {}

// Options are 1) define a function:
// AddHandler adds the route Handler to the provided Server
func AddHandler[T1, T2 any](s *Server, h *Handler[T1, T2]) {
	http.HandleFunc(h.Pattern, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := new(T1)

		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			JSON(w, err.Error(), 400)
			return
		}
		res := h.HandleRequest(*req)

		JSON(w, res, 200)
	}))
}

// Options are 2) define the method on Handler:
// AddToServerMux adds the route Handler to the provided Server
func (h *Handler[T1, T2]) AddToServerMux(s *Server) {
	AddHandler(s, h)
}

func JSON(w http.ResponseWriter, v interface{}, status int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(buf.Bytes()) //nolint:errcheck
}
