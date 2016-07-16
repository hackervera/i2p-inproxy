package main

import (
"fmt"
"net/http"
"strings"
"io/ioutil"
)

func handler(w http.ResponseWriter, r *http.Request) {
    domain := strings.Split(r.Host, ".")[0]
    url := "http://" + domain + ".i2p" + r.URL.Path
    resp, err := http.Get(url)
    fmt.Println(err)
    fmt.Println(resp)
    bs, _ := ioutil.ReadAll(resp.Body)
    fmt.Fprintf(w, string(bs))
}

func main() {
  http.HandleFunc("/", handler)
  fmt.Println("Running")
  http.ListenAndServe(":8080", nil)
}
