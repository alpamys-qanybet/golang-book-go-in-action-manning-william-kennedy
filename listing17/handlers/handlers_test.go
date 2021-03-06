package handlers_test

import (
  "encoding/json"
  "net/http"
  "net/http/httptest"
  "testing"

  "books/go-in-action/listing17/handlers"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

// init function is declared to initialize the routes.
// If the routes aren’t initialized before the unit tests are run, then the tests will fail with an http.StatusNotFound error
func init() {
  handlers.Routes()
}

func TestSendJson(t *testing.T) {
  t.Log("Given the need to test the SendJSON endpoint.")
  {
    req, err := http.NewRequest("GET", "/sendjson", nil)
    if err != nil {
      t.Fatal("\tShould be able to create a request.", ballotX, err)
    }
    t.Log("\tShould be able to create a request.", checkMark)

    rw := httptest.NewRecorder()
    http.DefaultServeMux.ServeHTTP(rw, req)

    if rw.Code != http.StatusOK {
      t.Fatal("\tShould receive a \"200\"", ballotX, rw.Code)
    }
    t.Log("\tShould receive a \"200\"", checkMark)

    u := struct {
      Name string
      Email string
    } {}

    if err := json.NewDecoder(rw.Body).Decode(&u); err != nil {
      t.Fatal("\tShould decode the response.", ballotX)
    }
    t.Log("\tShould decode the response.", checkMark)

    if u.Name == "Bill" {
      t.Log("\tShould have a Name.", checkMark)
    } else {
      t.Error("\tShould have a Name.", ballotX, u.Name)
    }

    if u.Email == "bill@gmail.com" {
      t.Log("\tShould have an Email.", checkMark)
    } else {
      t.Error("\tShould have an Email.", ballotX, u.Email)
    }
  }
}
