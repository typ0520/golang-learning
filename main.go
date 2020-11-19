package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hello World!")

	var name string = "Q1mi"
	var name2, age2 = "Q1mi", 20

	fmt.Println(name, name2, age2)

	m := 1
	fmt.Println(m)

	const (
		_      = iota
		FLAG_1 = 1 << (1 * iota)
		FLAG_2 = 1 << (1 * iota)
		FLAG_3 = 1 << (1 * iota)
		FLAG_4 = 1 << (1 * iota)
	)
	fmt.Println(FLAG_1, FLAG_2, FLAG_3, FLAG_4)

	// 十进制
	var a int = 10
	fmt.Printf("%d \n", a) // 10
	fmt.Printf("%b \n", a) // 1010  占位符%b表示二进制

	// 八进制  以0开头
	var b int = 077
	fmt.Printf("%o \n", b) // 77

	// 十六进制  以0x开头
	var c int = 0xff
	fmt.Printf("%x \n", c) // ff
	fmt.Printf("%X \n", c) // FF

	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)

	fmt.Println(len("a" + "b"))

	sqrtDemo()
	forDemo()
	arrayDemo()
	array2Demo()
}

func sqrtDemo() {
	var a, b = 3, 4
	var c int
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func forDemo() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func arrayDemo() {
	var a = [...]string{"北京", "上海", "深圳"}
	// 方法1：for循环遍历
	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}

	// 方法2：for range遍历
	for index, value := range a {
		fmt.Println(index, value)
	}

	a2 := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a2)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a2[2][1]) //支持索引取值:重庆
	for _, v1 := range a2 {
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}

	arr := [...]int{1, 3, 5, 7, 8}
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
	sum = 0
	for _, v := range arr {
		sum += v
	}
	fmt.Println(sum)
}

func array2Demo() {
	a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10
}
