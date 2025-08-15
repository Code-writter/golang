# Maps

Maps are a data structure which allow us to store data values in key:value pairs.
 > More like a object in the Javescript

 > map [ key_type ] value_type { key : value, }

A map is an unordered and changeable collection that does not allow duplicates.

 > The default value of a map is nil.

```go
userInfo := map[string]int{
		"huxn": 17,
		"alex": 18,
		"john": 27,
	}

	fmt.Println(userInfo)
	userInfo["jordan"] = 15
	fmt.Println(userInfo["huxn"])
	fmt.Println(userInfo["alex"])
	fmt.Println(userInfo["john"])
```

# Loop

```go
	for key, value := range users {
		fmt.Println(key, value)
	}
```