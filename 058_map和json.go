package main

import (
	"encoding/json"
	"fmt"
)

func JsonToMapDemo() {
	jsonStr := `
        {
                "name": "jqw",
                "age": 18
        }
        `
	var mapResult map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
	}
	fmt.Println(mapResult)
}
func MapToJsonDemo1() {
	mapInstances := []map[string]interface{}{}
	instance_1 := map[string]interface{}{"name": "John", "age": 10}
	instance_2 := map[string]interface{}{"name": "Alex", "age": 12}
	mapInstances = append(mapInstances, instance_1, instance_2)

	jsonStr, err := json.Marshal(mapInstances)

	if err != nil {
		fmt.Println("MapToJsonDemo err: ", err)
	}
	fmt.Println(string(jsonStr))
}

func main() {
	fmt.Println("-----1.json转map----")
	JsonToMapDemo()
	fmt.Println("-----2.map转json----")
	MapToJsonDemo1()

}
