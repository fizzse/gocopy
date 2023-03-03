package gocopy

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

var (
	FieldTypeError      = errors.New("field type error")
	FieldNotExistsError = errors.New("field not exists")
)

type IMap struct {
	Value   interface{}
	InValid bool
	getKey  string
}

func NewIMap(m interface{}) (v *IMap) {
	v = &IMap{Value: m}
	rv := reflect.ValueOf(m)
	if rv.Kind() != reflect.Map {
		v.InValid = true
	}
	return
}

func (i *IMap) Get(key string) (v *IMap) {
	v = &IMap{getKey: i.getKey}
	if i.InValid {
		return i
	}

	if v.getKey == "" {
		v.getKey = key
	} else {
		v.getKey = v.getKey + "." + key
	}

	rv := reflect.ValueOf(i.Value)
	if rv.Kind() != reflect.Map {
		v.InValid = true
		return
	}

	iv := rv.MapIndex(reflect.ValueOf(key)).Interface()
	v.Value = iv
	return
}

func (i *IMap) GetDeep(key string) (v *IMap) {
	v = &IMap{Value: i.Value}
	if i.InValid {
		return v
	}

	rv := reflect.ValueOf(i.Value)
	if rv.Kind() != reflect.Map {
		v.InValid = true
		return
	}

	keys := strings.Split(key, ".")
	for _, key := range keys {
		v = v.Get(key)
		if v.InValid {
			return
		}
	}

	return
}

func (i *IMap) Valid() (err error) {
	if i.InValid {
		err = fmt.Errorf("%w: key %s not exists", FieldNotExistsError, i.getKey)
	}

	return
}

func (i *IMap) Int() (value int, err error) {
	if err = i.Valid(); err != nil {
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
	default:
		err = fmt.Errorf("%w: %s is %s not int", FieldTypeError, i.getKey, rv.Kind().String())
		return
	}

	return
}

func (i *IMap) Float() (value float64, err error) {
	if err = i.Valid(); err != nil {
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
	default:
		err = fmt.Errorf("%w: %s is %s not float", FieldTypeError, i.getKey, rv.Kind().String())
		return
	}

	return
}

func (i *IMap) String() (value string, err error) {
	if err = i.Valid(); err != nil {
		return
	}

	rv := reflect.ValueOf(i.Value)
	switch rv.Kind() {
	case reflect.String:
		v := rv.String()
		value = v

	default:
		err = fmt.Errorf("%w: %s is %s not string", FieldTypeError, i.getKey, rv.Kind().String())
		return
	}

	return
}
