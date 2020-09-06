package main

import "fmt"
import _ "go-demo/demo1"

func main() {
	println("hello world")
	fmt.Println("hello world")

	// 测试导入包时会执行 init
}
