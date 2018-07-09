package main

import (
  "fmt"
  "github.com/gorilla/mux"
  "log"
  "net/http"
  "net/http/httputil"
  "os"
  "bufio"
  "time"
  "github.com/tomasen/realip"
)

var logfile, _ = os.Create("/root/httpdebug/debug.log")
var logwriter = bufio.NewWriter(logfile)

func DumpRequest(w http.ResponseWriter, req *http.Request) {
  fmt.Fprint(logwriter, "===========\n" + time.Now().String() + "\n")
  fmt.Fprint(logwriter, realip.FromRequest(req) + "\n----------\n")

  requestDump, err := httputil.DumpRequest(req, true)
  if err != nil {
    fmt.Fprint(w, err.Error())
    fmt.Fprint(logwriter, err.Error())
  } else {
    fmt.Fprint(w, string(requestDump))
    fmt.Fprint(logwriter, string(requestDump))
  }

  logwriter.Flush()
}

func main() {
  router := mux.NewRouter()
  router.PathPrefix("/").HandlerFunc(DumpRequest)
  log.Fatal(http.ListenAndServe(":80", router))
  log.Fatal(http.ListenAndServeTLS(":443", "fullchain.pem", "privkey.pem",  router))
}
