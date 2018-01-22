package main

import (
   "encoding/json"
   "fmt"
   "net/http"

   "github.com/gorilla/mux"
)

type Message struct {
   Message string `json:"message"`
}

func Hacker(w http.ResponseWriter, r *http.Request) {
   vars := mux.Vars(r)
   hackerId := vars["id"]
   mess := Message{"Searching for infor about hacker with id: " + string(hackerId)}
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   if err := json.NewEncoder(w).Encode(mess); err != nil {
      panic(err)
   }
}

func Hackers(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   mess := Message{"This will be an array of hackers"}
   if err := json.NewEncoder(w).Encode(mess); err != nil {
      panic(err)
   }
}

func Schedule(w http.ResponseWriter, r *http.Request) {
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusOK)
   mess := Message{"This will be the schedule"}
   if err := json.NewEncoder(w).Encode(mess); err != nil {
      panic(err)
   }

}

func Signin(w http.ResponseWriter, r *http.Request) {
   fmt.Fprintln(w, "This is where an okey message will be sent when signin is successful")
}
