package main

import (
   //"encoding/json"
   "fmt"
   "net/http"

   "github.com/gorilla/mux"
)

func Hacker(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   hackerId := vars["id"]
   fmt.Fprintln(w, "Looking for hacker with id:", hackerId)
}

func Hackers(w http.ResponseWriter, r * http.Request) {
   fmt.Fprintln(w, "Looking for all hackers")
}
