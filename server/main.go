package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
    // Hello world, the web server

    helloHandler := func(w http.ResponseWriter, req *http.Request) {
        w.Header().Set("Content-Type", "application/json")

        resp := make(map[string]bool)
        resp["status"] = true
        jsonResp, err := json.Marshal(resp)
        if err != nil {
            log.Fatalf("Error happened in JSON marshal. Err: %s", err)
        }
        w.Write(jsonResp)
        return
    }

    http.HandleFunc("/", helloHandler)

    log.Println("Listing for requests at http://localhost:3030/hello")

    log.Fatal(http.ListenAndServe(":3030", nil))
}