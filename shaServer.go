package main

import (
  "crypto/sha256"
  "fmt"
  "log"
  "net/http"
  "strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
  keys, ok := r.URL.Query()["key"]
  if !ok || len(keys) < 1 {
    log.Println("Url Param missing")
    return
  }
  counts, ok := r.URL.Query()["count"]
  if !ok || len(counts) < 1 {
    log.Println("Url Param missing")
    return
  }
  key := keys[0]
  count, err := strconv.Atoi(counts[0])
  if err != nil {
    log.Println("count should be an integer")
  }
  sum := sha256.Sum256([]byte(key))
  for i := 1; i < count; i++ {
    sum = sha256.Sum256([]byte(key))
    log.Printf("Calculated SHA for %v: %v/n",key, sum)
  }
  fmt.Fprintf(w, "<html><body>Calculated SHA for %v: %v\n</body></html>",key, sum)
}
 
// We'll calculate shas to busy the cpu.  We can vary the filesize and
// the number of times we calculate sha given url params
func main() {
  http.HandleFunc("/", handler)
  log.Fatal(http.ListenAndServe(":80", nil))
}
