package gocopy

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	m := map[string]interface{}{
		"name":  "张三",
		"age":   18,
		"money": 18.2,
	}

	type Person struct {
		Name  string   `json:"name"`
		Age   *int     `json:"age"`
		Money *float64 `json:"money"`
	}

	//var p *Person
	var p = new(Person)
	p.Age = new(int)
	p.Money = new(float64)

	_ = MapToStruct(m, &p, "json")
	c, _ := json.Marshal(p)
	fmt.Println(string(c))

	//json.Unmarshal()
}
