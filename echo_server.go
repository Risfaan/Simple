package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var port string = "8000"

type requestLogger struct{}

func (rl requestLogger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var bodyBytes []byte
	var err error

	if r.Body != nil {
		bodyBytes, err = ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Printf("Body reading error: %v", err)
			return
		}
		defer r.Body.Close()
	}

	var prettyJSON bytes.Buffer
	if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
		fmt.Printf("JSON parse error: %v", err)
		return
	}

	fmt.Println(string(prettyJSON.Bytes()))
}

func main() {
	fmt.Printf("Starting request echo server on port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", port), requestLogger{})
	fmt.Printf("Server error: %v\n", err)
}
