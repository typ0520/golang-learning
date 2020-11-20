package main

import (
	"fmt"
	"mypackage" // 导入同一项目下的mypackage包
)

func init() {
	fmt.Println("moduledemo.init")
}

func main() {
	mypackage.New()
	fmt.Println("main2")
}
