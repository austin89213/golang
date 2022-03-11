package main

import ("encoding/json" 
		"fmt")

type permissions map[string]bool

type user struct{
	Name string  `json:"username"`
	Password string `json:"-"`
	Permissions permissions `json:"perms,omitempty"`
}

func main(){
	users := []user{
		{"austin", "1234", nil},
		{"ivan", "5678", nil},
		{"god", "fly", map[string]bool{"admin":true}},
	}
	out, err := json.MarshalIndent(users, "", "\t")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(out))
}