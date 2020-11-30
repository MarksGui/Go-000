package main

type UserEntity struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	Age  uint32 `json:"age"`
}
