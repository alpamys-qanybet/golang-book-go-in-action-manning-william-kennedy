package handlers

import (
  "encoding/json"
  "net/http"
)

func Routes() {
  http.HandleFunc("/sendjson", SendJSON)
}

func SendJSON(rw http.ResponseWriter, r *http.Request) {
  u := struct {
    Name string
    Email string
  } {
    Name: "Bill",
    Email: "bill@gmail.com",
  }

  rw.Header().Set("Content-Type", "application/json")
  rw.WriteHeader(http.StatusOK)

  json.NewEncoder(rw).Encode(&u)
}
