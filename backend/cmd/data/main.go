package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"custom_age"`
	}

	u := User{
		ID:   1,
		Name: "gopher",
		Age:  14,
	}

	b, err := json.Marshal(u)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))

	u = User{}
	if err := json.Unmarshal(b, &u); err != nil {
		panic(err)
	}

	fmt.Println(u)

}
