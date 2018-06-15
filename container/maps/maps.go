package main

import "fmt"

func main() {

	m := map[string]string{
		"name":    "myway",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int) // m2 == empty map

	var m3 map[string]int // m3 == nil

	fmt.Println(m, m2, m3)

	//遍历map
	fmt.Println("Traversing map")
	for k, v := range m {
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for _, v := range m {
		fmt.Println(v)
	}

	//取值
	fmt.Println("Getting values")
	courseName, ok := m["course"]
	fmt.Println(courseName, ok)

	//判断值是否存在
	if causeName, ok := m["cause"]; ok {
		fmt.Println(causeName, ok)
	} else {
		fmt.Println("key does not exist")
	}

	//删除elements
	fmt.Println("Deleting values")
	name, ok := m["name"]
	fmt.Println(name, ok)

	delete(m, "name")
	name, ok = m["name"]
	fmt.Println(name, ok)

}
