package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"errors"
)

func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	info, err := NewUserService().GetInfo()
	if err != nil {
		log.Printf("%+v\n", err)
		if errors.Is(err, ErrNotFound) {
			fmt.Fprintf(os.Stdout, "%s", "查无记录")
			return
		}
		fmt.Fprintf(os.Stdout, "%s", "查询失败")
		return
	}

	fmt.Fprintf(os.Stdout, "%+v\n", info)
}
