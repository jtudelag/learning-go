package main

import (
    "strings"
    "os"
    "time"
    "fmt"
    "net/http"
    "log"
    "github.com/gorilla/mux"
    "strconv"
)
var TIMEOUT int = 0

func RootHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Welcome to RH Services Meeting 2017 Hackaton Day!\n"))
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
    time.Sleep(time.Second * time.Duration(TIMEOUT))
    w.WriteHeader(http.StatusOK)
}

func HealthHandlerPost(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  TIMEOUT,_ = strconv.Atoi(vars["timeout"])
  res := fmt.Sprintf("Health Timeout set to %v seconds!\n", TIMEOUT)
  w.Write([]byte(res))
}

func CustomValueHandler(w http.ResponseWriter, r *http.Request) {
    var res string
    myvar := os.Getenv("MYVAR")
    if len(strings.TrimSpace(myvar)) == 0 {
      res = "MYVAR no esta definida!!\n"
    } else {
      res = fmt.Sprintf("El valor de MYVAR es: %v\n", myvar)
    }
    w.Write([]byte(res))

}

func main() {
    r := mux.NewRouter()
    // Routes consist of a path and a handler function.
    r.HandleFunc("/", RootHandler).Methods("GET")
    r.HandleFunc("/health", HealthHandler).Methods("GET")
    r.HandleFunc("/health/{timeout}", HealthHandlerPost).Methods("POST")
    r.HandleFunc("/customvalue", CustomValueHandler).Methods("GET")
    // Static files
    r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("/tmp/myweb/"))))

    //TODO: Reads the port as an argument!

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", r))
}
