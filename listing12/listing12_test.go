package listing12

import (
  // "encoding/xml"
  "fmt"
  "net/http"
  "net/http/httptest"
  // "testing"
)

const checkMark = "\u2713"
const ballotX = "\u2717"

var feed = `<?xml version="1.0" encoding="UTF-8"?>
<rss>
  <channel>
    <title>Going Go Programming</title>
    <description>Golang : https://github.com/goinggo</description>
    <link>http://www.goinggo.net/</link>
    <item>
      <pubDate>Sun, 15 Mar 2015 15:04:00 +0000</pubDate>
      <title>Object Oriented Programming Mechanics</title>
      <description>Go is an object oriented language.</description>
      <link>http://www.goinggo.net/2015/03/object-oriented</link>
    </item>
  </channel>
</rss>`

func mockServer() *httptest.Server {
  f := func (w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/xml")
    fmt.Fprintln(w, feed)
  }

  return httptest.NewServer(http.HandlerFunc(f))
}
