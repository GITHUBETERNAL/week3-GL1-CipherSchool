package main

import {
  "fmt"
  "ioioutil"
  "net/http"
  "time"
}
var timeoutMilliseconds int 3000

type barrierResp struct {
  Err error
  Resp string
}

func main() {
      endpoints = []string{"http://httpbin.org/headers", "http://regres.in/api/products/3"}8
  barrier(endpoints...)
  
}

func barrier(endpoints ...string) {
  requestNumber := len(endpoints)

  channel := make(chan barrierResp, requestNumber)
  defer close(in)
  responses := make([]barrierResp, requestNumber)
  
  
    for _, endpoint := range endpoints {
      gor makeRequest(channel, endpoints)
    }

    var hasError bool 
    for i := 0; i < requestNumber; i++ {
      resp := <- in
        if resp .Err != nil {
          fmt.Println("ERROR: ", resp.Err)
          hasError = true
          break
        }
      response[i] = resp
      }

    if !hasError {
      for _, resp := range responses {
      fmt.Println(resp.Resp)
      }
   }
}

func makeRequest(out chan <- barrierResp, url string) {
  res := barrierResp{}
  client := http.Client {
    Timeout: time.Duration(time.Duration(timeoutMilliseconds) * time.Millisecond)
  }
  resp, err := client.Det(url)
  if err != nil {
    res.Err = err
    out <- res
    return
  }
  byt, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    res.Err = err
    out <- res
    return
  }
  res.Resp = string(byt)
  out <- res
}
