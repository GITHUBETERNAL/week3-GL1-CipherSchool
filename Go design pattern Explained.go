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

  in := make(chan barrierResp, requestNumber)
  defer close(in)
  responses := make([]barrierResp, requestNumber)
  
  
    for _, endpoint := range endpoints {
      gor makeRequest(in, endpoints)
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
