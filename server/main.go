package main

import (
    "log"
    "net/http"
    "encoding/json"
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

    log.Println("Listing for requests at http://localhost:8000/hello")

    log.Fatal(http.ListenAndServe(":8000", nil))
}