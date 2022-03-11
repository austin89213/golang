package main

import (
	"bufio"
	"fmt"
	"os"
	"encoding/json"
)

type user struct{
	Name string `json:"username`
	Permissions map[string]bool 
}

func main(){
	var input []byte
	for in := bufio.NewScanner(os.Stdin); in.Scan();{
		input = append(input, in.Bytes()...)
	}
	var users []user
	err := json.Unmarshal(input, &users)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, user := range users{
		fmt.Print("+ ", user.Name)
		switch p := user.Permissions;{
		case p == nil:
			fmt.Print(" has no power.")
		case p["admin"]:
			fmt.Print(" is an admin.")
		case p["write"]:
			fmt.Print(" can write")
		}
	}
}