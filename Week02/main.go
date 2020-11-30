package main

import "net/http"

func main() {
	http.HandleFunc("/user", GetUserInfo)
	http.ListenAndServe(":8080", nil)
}
