# 利用go的反射机制 对map和struct进行转化

- MapToStruct
- StructToMap
- IMap (map[sting]interface{}) 解决取值断言的问题，使用场景如下
```go
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

// 注意 数值类型 int float都可以取到值，int会对float做截断
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
```