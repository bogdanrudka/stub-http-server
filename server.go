package main

import (
	"net/http"
	"log"
	"fmt"
	"io"
)

type StubServer struct {
	response Response
}

func (s StubServer) ServeHTTP(rw http.ResponseWriter, res *http.Request) {
	switch v := s.response.Body.(type) {
	case string:
		io.WriteString(rw, v)
	default:
		fmt.Printf("Type %t not supported", v)
	}
}

func main() {
	conf := InitConfig()
	for _, m := range conf.Mappings {
		http.Handle("/"+m.Request.Path, StubServer{response: m.Response})
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf("localhost:%d", conf.Settings.Port), nil))
}
