package main

import (
	"fmt"
	"time"

	"github.com/skatsuta/ffjson-sample/user"
)

func main() {
	u := &user.User{
		ID:      1,
		Name:    "John",
		Age:     20,
		Created: time.Now(),
	}

	b, e := u.MarshalJSON()
	if e != nil {
		panic(e)
	}
	fmt.Println(string(b))

	j := []byte(`{"Name": "Ken", "Age": 30, "ID": 10}`)
	if e := u.UnmarshalJSON(j); e != nil {
		panic(e)
	}
	fmt.Printf("%#v", u)
}
