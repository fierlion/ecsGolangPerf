package main

import (
  "crypto/sha256"
  "fmt"
  "log"
  "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  sum := sha256.Sum256([]byte("this is a hash\n"))
  fmt.Fprintf(w, "Calculated SHA: %v\n", sum)
}

func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":8080", nil))
}
