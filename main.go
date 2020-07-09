package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Path string `json:"path"`
	Body string `json:"body"`
}

func main() {
	fmt.Println("====== read masterdata =====")

	fp, err := os.Open("messages.json")
	if err != nil {
		fmt.Printf("%s", err)
	}

	defer fp.Close()

	bytes := make([]byte, 256)
	stringBytes := make([]byte, 0)

	for {
		index, err := fp.Read(bytes)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		if err != nil {
			fmt.Printf("%s", err)
			break
		}
		fmt.Printf("%d", index)
		fmt.Printf("%v", string(bytes))

		stringBytes = append(stringBytes, bytes...)
	}
	fmt.Println("===== start server =====")

	decoder := json.NewDecoder(strings.NewReader(string(stringBytes)))

	var messages []Message
	err = decoder.Decode(&messages)
	if err != nil {
		fmt.Printf("%s", err)
	}
	for _, m := range messages {
		fmt.Println(m.Path)

		path := m.Path
		body := m.Body
		http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Write([]byte(body))
		})
	}

	err = http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Printf("%s", err)
	}
}
