package gocopy

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestIMapGet(t *testing.T) {
	str := `{
		"name": "Yasuo",
		"age": 18,
		"money": 18.2,
		"brother": {
			"name": "Yone",
			"age": 28,
			"money": 38.2
		}
	}`

	m := make(map[string]interface{})
	json.Unmarshal([]byte(str), &m)

	money, err := IMap(m).Get("brother").Get("money").Int()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("money", money)

	name, err := IMap(m).GetDeepKey("brother.name").String()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("name", name)
}
