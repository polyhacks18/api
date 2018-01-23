package main

type Person struct {
   FirstName string                `json:"firstName,omitempty"`
   LastName string                `json:"lastName,omitempty"`
   ID []byte                   `json:"id,omitempty"`
   Position int                `json:"pos,omitempty"`
   PointLog []PointTransaction `json:"pointLog,omitempty"`
}

type PointTransaction struct {
   PointValue int             `json:"pointValue,omitempty"`
   TransactionType string     `json:"transactionType,omitempty"`
   Tag string                 `json:"tag,omitempty"`
}

type People []Person

var Database People = People{
   Person{
      "Greg",
      "Willard",
      []byte("9999"),
      1,
      make([]PointTransaction, 0),
   },
   Person{
      "Vamsi",
      "Hanumanthu",
      []byte("1234"),
      2,
      make([]PointTransaction, 0),
   },
   Person{
      "Caleb",
      "Long",
      []byte("1111"),
      3,
      make([]PointTransaction, 0),
   },
}
