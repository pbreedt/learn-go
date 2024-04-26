package generics

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

func TestRequests(t *testing.T) {

	go RunServer()

	url := "http://" + getServerAddress()

	expect := `{"string":"","int":0,"float":0,"bool":false,"object":{"name":""},"type":"response","out":"pong"}`
	if _, body := testRequest(t, "GET", url+"/ping", `{"type":"request"}`); strings.TrimSpace(body) != strings.TrimSpace(expect) {
		t.Fatalf("expected:'%s', got:'%s'", expect, body)
	}

	expect = `{"type":"response","out":"blah"}`
	if _, body := testRequest(t, "GET", url+"/echo", `{"type":"request", "in":"blah"}`); strings.TrimSpace(body) != strings.TrimSpace(expect) {
		t.Fatalf("expected:'%s', got:'%s'", expect, body)
	}

}

func RunServer() {

	svr := NewServer()

	complexHandler := Handler[ComplexRequest, ComplexResponse]{
		Pattern: "/ping",
		HandleRequest: func(req ComplexRequest) ComplexResponse {
			res := ComplexResponse{Type: "response", Out: "pong"}
			fmt.Printf("complexHandler:\nrequest:%+v\nresponse:%+v\n", req, res)
			return res
		},
	}
	// Can use either:
	// complexHandler.AddToServerMux(svr)
	AddHandler(svr, &complexHandler)

	simpleHandler := Handler[SimpleRequest, SimpleResponse]{
		Pattern: "/echo",
		HandleRequest: func(req SimpleRequest) (res SimpleResponse) {
			res = SimpleResponse{Type: "response", Out: req.In}
			fmt.Printf("simpleHandler:\nrequest:%+v\nresponse:%+v\n", req, res)
			return res
		},
	}
	// Can use either:
	simpleHandler.AddToServerMux(svr)
	// AddHandler(svr, &simpleHandler)

	fmt.Printf("Server listening on %s\n", getServerAddress())
	svr.Run(getServerAddress())
}

type ComplexPayload struct {
	String string  `json:"string"`
	Int    int     `json:"int"`
	Float  float64 `json:"float"`
	Bool   bool    `json:"bool"`
	Obj    struct {
		Name string `json:"name"`
	} `json:"object"`
}

type ComplexRequest struct {
	ComplexPayload
	Type string `json:"type"`
	In   string `json:"in"`
}

type ComplexResponse struct {
	ComplexPayload
	Type string `json:"type"`
	Out  string `json:"out"`
}

type SimpleRequest struct {
	Type string `json:"type"`
	In   string `json:"in"`
}

type SimpleResponse struct {
	Type string `json:"type"`
	Out  string `json:"out"`
}

func getServerAddress() string {
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	if len(host) == 0 {
		host = "localhost"
	}
	if len(port) == 0 {
		port = "8080"
	}
	return fmt.Sprintf("%s:%s", host, port)
}

func testRequest(t *testing.T, method, path string, body string) (*http.Response, string) {
	req, err := http.NewRequest(method, path, bytes.NewBufferString(body))
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
		return nil, ""
	}
	defer resp.Body.Close()

	return resp, string(respBody)
}
