package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

var (
	masterdata = `{ "Hello": "world" }`
)

type Message struct {
	Hello string
}

func main() {
	fmt.Println("start server")

	decoder := json.NewDecoder(strings.NewReader(masterdata))

	var message Message
	err := decoder.Decode(&message)
	if err != nil {
		fmt.Printf("%s", err)
	}
	fmt.Printf("%v", message.Hello)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("Hello browser"))
	})

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("world"))
	})

	http.ListenAndServe("localhost:11181", nil)
}
