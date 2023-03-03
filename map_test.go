package gocopy

import (
	"encoding/json"
	"errors"
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

	v := NewIMap(m)
	money, err := v.Get("brother").Get("money").Int()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("money", money)

	name, err := v.GetDeep("brother.name").String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name", name)
}

func TestIMapGet2(t *testing.T) {
	m := map[string]string{
		"name": "Yasuo",
		"age":  "18",
	}

	v := NewIMap(m)

	name, err := v.GetDeep("name").String()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("name", name)
}

func TestIMapGet_error(t *testing.T) {
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

	v := NewIMap(m)
	_, err := v.Get("brother").Get("money").Get("name").Int()
	if errors.Is(err, FieldNotExistsError) {
		fmt.Println(err)
	}

	_, err = v.GetDeep("brother.name").Int()
	if errors.Is(err, FieldTypeError) {
		fmt.Println(err)
	}
}
