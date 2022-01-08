package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name      string   `json:"name"`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name
	return err
}

func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr " + p.Name,
	})
	return v, err
}

func main() {
	b := []byte(`{"name":"mike","age":20,"nicknames":["a","b","c"]}`)
	var p Person
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)

	v, _ := json.Marshal(p)
	fmt.Println(string(v))
}
