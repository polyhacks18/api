package main

type Person struct {
   FirstName string            `json:"firstName"`
   LastName string             `json:"lastName"`
   id []byte
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
   Person{
      "Jessie",
      "Pullaro",
      []byte("451C3D487AEDA41E70A44977D69571DF991E3BC2EC1E4A0C8B25C427C85C57F0"),
      4,
      make([]PointTransaction, 0),
   },
}
