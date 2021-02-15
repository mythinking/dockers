package main

import (
  "fmt"
  "net/http"
)

func NotifyHandler(w http.ResponseWriter, r *http.Request) {
  r.ParseForm()
  len := r.ContentLength
  body := make([]byte, len)
  r.Body.Read(body)
  fmt.Println("url: ", r.URL)
  fmt.Println("body: ", string(body))
  fmt.Fprintln(w, string(body))
}

func main() {
  http.HandleFunc("/sms", NotifyHandler)
  http.ListenAndServe("0.0.0.0:19000", nil)
}
