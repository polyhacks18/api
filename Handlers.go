package main

import (
   "encoding/json"
   "fmt"
   "net/http"

   "github.com/gorilla/mux"
)

func Hacker(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   hackerId := vars["id"]
   fmt.Fprintln(w, "Looking for hacker with id:", hackerId)
}

func Hackers(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "Looking for all hackers")
}

func Schedule(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   mess := `{"message":"This is where event items will show up"}`
   encode,_ := json.Marshal(mess)

   if err := json.NewEncoder(w).Encode(encode); err != nil {
      panic(err)
   }

}

func Signin(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "This is where an okey message will be sent when signin is successful")
}
