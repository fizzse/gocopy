package gocopy

import (
	"fmt"
	"reflect"
)

func indirect(v reflect.Value) reflect.Value {
	// Issue #24153 indicates that it is generally not a guaranteed property
	// that you may round-trip a reflect.Value by calling Value.Addr().Elem()
	// and expect the value to still be settable for values derived from
	// unexported embedded struct fields.
	//
	// The logic below effectively does this when it first addresses the value
	// (to satisfy possible pointer methods) and continues to dereference
	// subsequent pointers as necessary.
	//
	// After the first round-trip, we set v back to the original value to
	// preserve the original RW flags contained in reflect.Value.
	v0 := v
	haveAddr := false

	// If v is a named type and is addressable,
	// start with its address, so that if the type has pointer methods,
	// we find them.
	if v.Kind() != reflect.Pointer && v.Type().Name() != "" && v.CanAddr() {
		haveAddr = true
		v = v.Addr()
	}
	for {
		// Load value from interface, but only if the result will be
		// usefully addressable.
		if v.Kind() == reflect.Interface && !v.IsNil() {
			e := v.Elem()
			if e.Kind() == reflect.Pointer && !e.IsNil() && (e.Elem().Kind() == reflect.Pointer) {
				haveAddr = false
				v = e
				continue
			}
		}

		if v.Kind() != reflect.Pointer {
			break
		}

		// Prevent infinite loop if v is an interface pointing to its own address:
		//     var v interface{}
		//     v = &v
		if v.Elem().Kind() == reflect.Interface && v.Elem().Elem() == v {
			v = v.Elem()
			break
		}
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}

		if haveAddr {
			v = v0 // restore original value after round-trip Value.Addr().Elem()
			haveAddr = false
		} else {
			v = v.Elem()
		}
	}
	return v
}

func MapToStruct(m map[string]interface{}, container interface{}, tagKey string) (err error) {
	rv := reflect.ValueOf(container)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		err = fmt.Errorf("container must be pointer")
		return
	}

	// 如果是空地址 则赋值
	rv = indirect(rv)
	tv := rv.Type()

	for i := 0; i < rv.NumField(); i++ {
		tag := tv.Field(i).Tag.Get(tagKey)
		if tag == "" {
			continue
		}

		value, ok := m[tag]
		if !ok {
			continue
		}

		fieldValue := rv.Field(i)
		if !fieldValue.CanSet() {
			continue
		}

		// 如果是空值 分配空间
		if fieldValue.Kind() == reflect.Pointer {
			if fieldValue.IsNil() {
				fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
			}
			fieldValue = fieldValue.Elem()
		}

		if fieldValue.Kind() == reflect.Ptr { // 如果字段是指针类型，则需要先获取其指向的值
			fieldValue = fieldValue.Elem()
		}

		newValue := reflect.ValueOf(value).Convert(fieldValue.Type())
		fieldValue.Set(newValue)
	}

	return
}
