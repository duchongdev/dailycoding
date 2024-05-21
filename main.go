package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Address string
	Email   string
}

func main() {
	u := User{
		Name:    "zhangsan",
		Age:     18,
		Address: "beijing",
		Email:   "test@qq.com",
	}

	str, _ := json.Marshal(u)
	fmt.Printf("str = %s\n", str)
}
