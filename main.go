package main

import (
  "net/http"
  "time"
)

func Get(url string) (resp *http.Response, err error)  {
  var netClient = &http.Client{
    Timeout: time.Second * 10,
  }
  resp, err = netClient.Get(url)
  defer resp.Body.Close()

  return resp, err
}

func main() {

}
