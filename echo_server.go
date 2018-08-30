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

	fmt.Printf("Headers: %+v\n", r.Header)

	if len(bodyBytes) > 0 {
		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			fmt.Printf("JSON parse error: %v", err)
			return
		}
		fmt.Println(string(prettyJSON.Bytes()))
	} else {
		fmt.Printf("Body: No Body Supplied\n")
	}

	// CORS headers
	//w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	//w.Header().Set("Access-Control-Allow-Methods", "GET,PUT,POST,DELETE,OPTIONS")
	//w.Header().Set("Access-Control-Allow-Credentials", "true")
	//w.Header().Set("Access-Control-Allow-Headers", "Accept-Encoding,Authorization,X-Forwarded-For,Content-Type,Origin,Server")
	
	fmt.Println(string(prettyJSON.Bytes()))
}

func main() {
	fmt.Printf("Starting request echo server on port %v\n", port)
	err := http.ListenAndServe(fmt.Sprintf("127.0.0.1:%v", port), requestLogger{})
	fmt.Printf("Server error: %v\n", err)
}
