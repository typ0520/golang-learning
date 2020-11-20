package main

import (
	"fmt"
	"math"
	"strings"
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
	mapDemo()
	funcDemo()
	deferDemo()
	pointDemo()
	fmt.Println("end")
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

	a2 := [5]int{1, 2, 3, 4, 5}
	s := a2[1:3] // s := a[low:high]
	fmt.Printf("s:%v len(s):%v cap(s):%v\n", s, len(s), cap(s))

	var s2 []int
	s2 = append(s, 1)       // [1]
	s2 = append(s, 2, 3, 4) // [1 2 3 4]
	fmt.Println(s2)
}

func mapDemo() {
	scoreMap := make(map[string]int, 8)
	scoreMap["张三"] = 90
	scoreMap["小明"] = 100
	fmt.Println(scoreMap)
	fmt.Println(scoreMap["小明"])
	fmt.Printf("type of a:%T\n", scoreMap)

	userInfo := map[string]string{
		"username": "typ0520",
		"pwd":      "111111",
	}
	fmt.Println(userInfo)

	v, ok := scoreMap["张三"]
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("not found")
	}

	for k, v := range scoreMap {
		fmt.Println(k)
		fmt.Println(v)
	}
	for _, v := range scoreMap {
		fmt.Println(v)
	}
	delete(scoreMap, "张三")

	for k, v := range scoreMap {
		fmt.Println(k)
		fmt.Println(v)
	}

	//写一个程序，统计一个字符串中每个单词出现的次数。比如：”how do you do”中how=1 do=2 you=1。
	str := "how do you do"
	m := make(map[string]int, 10)
	for _, v := range strings.Split(str, " ") {
		_, ok := m[v]
		if ok {
			m[v] = m[v] + 1
		} else {
			m[v] = 1
		}
	}

	for k, v := range m {
		fmt.Printf("%s : %d", k, v)
		fmt.Println("")
	}
}

func add(x, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

func calc(x, y int, op func(int, int) int) int {
	return op(x, y)
}

func getOp(s string) func(int, int) int {
	switch s {
	case "+":
		return add
	case "-":
		return sub
	default:
		return nil
	}
}

func adder(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func calc2(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func funcDemo() {
	fmt.Println(calc(1, 2, add))
	fmt.Println(getOp("+")(1, 2))

	// 将匿名函数保存到变量
	add := func(x, y int) {
		fmt.Println(x + y)
	}
	add(10, 20) // 通过变量调用匿名函数

	//自执行函数：匿名函数定义完加()直接执行
	func(x, y int) {
		fmt.Println(x + y)
	}(10, 20)

	var f = adder(10)
	fmt.Println(f(10)) //20
	fmt.Println(f(20)) //40
	fmt.Println(f(30)) //70

	f1 := adder(20)
	fmt.Println(f1(40)) //60
	fmt.Println(f1(50)) //110

	f1, f2 := calc2(10)
	fmt.Println(f1(1), f2(2)) //11 9

	fmt.Println("start")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("end")
}

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func deferDemo() {
	fmt.Println(f1())
	fmt.Println(f2())
}

func pointDemo() {
	fmt.Println("pointDemo")
	a := 10
	b := &a // 取变量a的地址，将指针保存到b中
	fmt.Printf("type of b:%T\n", b)
	c := *b // 指针取值（根据指针去内存取值）
	fmt.Printf("type of c:%T\n", c)
	fmt.Printf("value of c:%v\n", c)

	a1 := new(int)
	b1 := new(bool)
	fmt.Printf("%T\n", a1) // *int
	fmt.Printf("%T\n", b1) // *bool
	fmt.Println(*a1)       // 0
	fmt.Println(*b1)       // false
}
