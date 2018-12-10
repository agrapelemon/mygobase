package main

import (
	"encoding/json"
	"fmt"
)

type NameModel struct {
	Name1 string
	Name2 string
}
type HostModel struct {
	IP   string
	Name []NameModel
}
type Host struct {
	IP   string
	Name string
}

func main() {
	fmt.Println("----1.结构体转json")

	m := Host{Name: "Sky", IP: "192.168.23.92"}

	b, err := json.Marshal(m)
	if err != nil {

		fmt.Println("Umarshal failed:", err)
		return
	}

	fmt.Println("json:", string(b))

	fmt.Println("----2.json转结构体")
	var movies2 Host
	data := `{"IP":"192.168.23.92","Name":"Sky"}`

	// movies2 := make([]Movie, 10)
	if err3 := json.Unmarshal([]byte(data), &movies2); err3 != nil {
		fmt.Println("反序列化 failed:", err)
	}
	fmt.Printf("%+v\r\n", movies2)

}
