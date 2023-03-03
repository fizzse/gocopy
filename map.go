package gocopy

import (
	"reflect"
	"strings"
)

type AnyValue struct {
	Value   interface{}
	InValid bool
}

type IMap map[string]interface{}

func (m IMap) Get(key string) (value *AnyValue) {
	value = &AnyValue{InValid: true}
	iv, ok := m[key]
	if !ok {
		return
	}

	value.Value = iv
	value.InValid = false
	return
}

func (m IMap) GetDeepKey(key string) (value *AnyValue) {
	keys := strings.Split(key, ".")
	for i, k := range keys {
		if i == 0 {
			value = m.Get(k)
			continue
		}

		value = value.Get(k)
	}

	return
}

func (i *AnyValue) Get(key string) (value *AnyValue) {
	if i.InValid {
		return i
	}

	m, ok := i.Value.(map[string]interface{})
	if !ok {
		return
	}

	return IMap(m).Get(key)
}

func (i *AnyValue) Int() (value int, err error) {
	if i.InValid {
		return
	}

	rv := reflect.ValueOf(i.Value)
	switch rv.Kind() {
	case reflect.Int:
		v := rv.Int()
		value = int(v)

	case reflect.Float64:
		v := rv.Float()
		value = int(v)
	}

	return
}

func (i *AnyValue) Float() (value float64, err error) {
	if i.InValid {
		return
	}

	rv := reflect.ValueOf(i.Value)
	switch rv.Kind() {
	case reflect.Int:
		v := rv.Int()
		value = float64(v)

	case reflect.Float64:
		v := rv.Float()
		value = v
	}

	return
}

func (i *AnyValue) String() (value string, err error) {
	if i.InValid {
		return
	}

	rv := reflect.ValueOf(i.Value)
	switch rv.Kind() {
	case reflect.String:
		v := rv.String()
		value = v

	case reflect.Array:
		v := rv.Bytes()
		value = string(v)
	}

	return
}
