package main

type People struct {
   fName string                `json:"firstName"`
   lName string                `json:"lastName"`
   id []byte                   `json:"id"`
   position int                `json:"pos"`
   pointLog []PointTransaction `json:"pointLog"`
}

type PointTransaction struct {
   pointValue int
   transactionType string
   tag string
}
