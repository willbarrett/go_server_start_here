package main

import (
  "github.com/gorilla/mux"
  "net/http"
)

func HomeHandler(w http.ResponseWriter, request *http.Request) {
  w.Write([]byte("Hello World"))
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}
