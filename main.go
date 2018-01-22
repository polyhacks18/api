package main 

import (
   "fmt"
   "net/http"
   "log"
)

func main() {
   var hackers []People

   hackers = append(hackers, People{fName: "Greg", lName: "Willard", id: []byte("1234")})
   fmt.Println(hackers)

   router := NewRouter()

   log.Fatal(http.ListenAndServe(":8080", router))
}
