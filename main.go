package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func init() {
	fmt.Println("main.init")
}

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
	structDemo()
	interfaceDemo()
	reflectDemo()
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

//类型定义
type NewInt int

//类型别名
type MyInt = int

type Person struct {
	name string
	age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{
		name,
		age,
	}
}

func (p Person) Dream() {
	fmt.Println("I have a dream")
}

func structDemo() {
	var a NewInt
	var b MyInt
	fmt.Printf("type of a: %T\n", a)
	fmt.Printf("type of b: %T\n", b)

	type person struct {
		name string
		city string
		age  int8
	}

	type person2 struct {
		name, city string
		age        int8
	}

	var p1 person
	p1.name = "dp"
	p1.city = "shanghai"
	p1.age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)
	fmt.Printf("p1: %T\n", &p1) //p1: *main.person

	var user struct {
		Name string
		Age  int
	}
	user.Name = "xx"
	user.Age = 18
	fmt.Printf("%#v\n", user)
	fmt.Printf("user: %T\n", &user) //user: *struct { Name string; Age int }

	var p2 = new(person)
	p2.name = "dp2"
	p2.city = "shanghai2"
	fmt.Printf("p2: %T\n", p2) //p2: *main.person

	p3 := &person{}
	fmt.Printf("%T\n", p3)     //*main.person
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"", city:"", age:0}
	p3.name = "七米"
	p3.age = 30
	p3.city = "成都"
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"七米", city:"成都", age:30}

	var p4 person
	fmt.Printf("p4=%#v\n", p4) //p4=main.person{name:"", city:"", age:0}

	p5 := person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p5=%#v\n", p5) //p5=main.person{name:"小王子", city:"北京", age:18}

	p6 := &person{
		name: "小王子",
		city: "北京",
		age:  18,
	}
	fmt.Printf("p6=%#v\n", p6) //p6=&main.person{name:"小王子", city:"北京", age:18}

	p8 := &person{
		"沙河娜扎",
		"北京",
		28,
	}
	fmt.Printf("p8=%#v\n", p8) //p8=&main.person{name:"沙河娜扎", city:"北京", age:28}

	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	fmt.Printf("stus %p\n", &stus)

	for i, stu := range stus {
		fmt.Printf("stu,index: %d %p\n", i, &stu)
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}

	p9 := NewPerson("dp", 18)
	p9.Dream()
	fmt.Printf("p9=%#v\n", p9)
	fmt.Println("structDemo")
}

// Sayer 接口
type Sayer interface {
	say()
}

// Mover 接口
type Mover interface {
	move()
}

// 接口嵌套
type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("猫会动")
}

func showType(a interface{}) {
	_, ok := a.(string)
	if ok {
		fmt.Println("string")
	}
	fmt.Printf("type:%T value:%v\n", a, a)

	switch v := a.(type) {
	case string:
		fmt.Printf("x is a string，value is %v\n", v)
	case int:
		fmt.Printf("x is a int is %v\n", v)
	case bool:
		fmt.Printf("x is a bool is %v\n", v)
	default:
		fmt.Println("unsupport type！")
	}
}

func interfaceDemo() {
	var x Mover
	var wangcai = cat{} // 旺财是dog类型
	x = wangcai         // x可以接收dog类型
	var fugui = &cat{}  // 富贵是*dog类型
	x = fugui           // x可以接收*dog类型
	x.move()

	var x2 animal
	x2 = cat{name: "花花"}
	x2.move()
	x2.say()

	showType("str")
	showType(0)

	fmt.Println("interfaceDemo")
}

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", t)
	fmt.Printf("type: %v, kind: %v\n", t.Name(), t.Kind())
	//fmt.Printf("type: %T\n", x)
}

func reflectValue(x interface{}) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

func reflectSetValue(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
	}
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}

type student struct {
	Name  string
	Score int
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func reflectDemo() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)

	type person struct {
		name string
		age  int
	}

	var d = person{
		name: "zwz",
		age:  18,
	}
	reflectType(d)
	reflectValue(a)
	reflectValue(b)

	reflectSetValue(&b)
	fmt.Printf("b: %d\n", b)

	// *int类型空指针
	var a1 *int
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a1).IsNil())
	// nil值
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	// 实例化一个匿名结构体
	b1 := struct{}{}
	// 尝试从结构体中查找"abc"字段
	fmt.Println("存在的结构体成员:", reflect.ValueOf(b1).FieldByName("abc").IsValid())
	// 尝试从结构体中查找"abc"方法
	fmt.Println("存在的结构体方法:", reflect.ValueOf(b1).MethodByName("abc").IsValid())
	// map
	c := map[string]int{}
	// 尝试从map中查找一个不存在的键
	fmt.Println("map中不存在的键：", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())

	stu1 := student{
		Name:  "zs",
		Score: 90,
	}
	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
	}
	if scoreField, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s index:%d type:%v json tag:%v\n", scoreField.Name, scoreField.Index, scoreField.Type, scoreField.Tag.Get("json"))
	}

	v := reflect.ValueOf(stu1)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name: %s\n", t.Method(i).Name)
		fmt.Printf("method %s\n", methodType)

		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
