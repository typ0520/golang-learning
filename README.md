### 关键字
```
Go语言中有25个关键字：


break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

此外，Go语言中还有37个保留字。

Constants:    true  false  iota  nil

Types:    int  int8  int16  int32  int64  
            uint  uint8  uint16  uint32  uint64  uintptr
            float32  float64  complex128  complex64
            bool  byte  rune  string  error

Functions:   make  len  cap  new  append  copy  close  delete
                complex  real  imag
                panic  recover

```

### 变量

- 变量声明

```
var 变量名 类型 = 表达式

var name string = "Q1mi"
var age int = 18
var name, age = "Q1mi", 20

在函数内部，可以使用更简略的 := 方式声明并初始化变量。
m := 200 // 此处声明局部变量m
```

- 批量声明
```
var (
    a string
    b int
    c bool
    d float32
)
```

- 常量
```
const pi = 3.1415
const e = 2.7182

const (
    pi = 3.1415
    e = 2.7182
)

常量n1、n2、n3的值都是100
const (
    n1 = 100
    n2
    n3
)
```

- iota

go语言的常量计数器，只能在常量的表达式中使用
```
const (
    n1 = iota //0
    n2        //1
    n3        //2
    n4        //3
)

const (
    n1 = iota //0
    n2        //1
    _
    n4        //3
)

const (
    n1 = iota //0
    n2 = 100  //100
    n3 = iota //2
    n4        //3
)
const n5 = iota //0


const (
    _  = iota
    KB = 1 << (10 * iota)
    MB = 1 << (10 * iota)
    GB = 1 << (10 * iota)
    TB = 1 << (10 * iota)
    PB = 1 << (10 * iota)
)
```

### 基本数据类型

- 整型
```
uint8	无符号 8位整型 (0 到 255)
uint16	无符号 16位整型 (0 到 65535)
uint32	无符号 32位整型 (0 到 4294967295)
uint64	无符号 64位整型 (0 到 18446744073709551615)
int8	有符号 8位整型 (-128 到 127)
int16	有符号 16位整型 (-32768 到 32767)
int32	有符号 32位整型 (-2147483648 到 2147483647)
int64	有符号 64位整型 (-9223372036854775808 到 9223372036854775807)

uint	32位操作系统上就是uint32，64位操作系统上就是uint64
int	32位操作系统上就是int32，64位操作系统上就是int64
uintptr	无符号整型，用于存放一个指针
```

- 浮点型

float32、float64

```
float32 的浮点数的最大范围约为 3.4e38，可以使用常量定义：math.MaxFloat32。 float64 的浮点数的最大范围约为 1.8e308，可以使用一个常量定义：math.MaxFloat64
```

- 布尔值

true/false

- 字符串转义符

```
\r	回车符（返回行首）
\n	换行符（直接跳到下一行的同列位置）
\t	制表符
\'	单引号
\"	双引号
\\	反斜杠
```

- 字符串
```
多行字符串
s1 := `第一行
第二行
第三行
`
fmt.Println(s1)

字符串的常用操作
len(str)	求长度
+或fmt.Sprintf	拼接字符串
strings.Split	分割
strings.contains	判断是否包含
strings.HasPrefix,strings.HasSuffix	前缀/后缀判断
strings.Index(),strings.LastIndex()	子串出现的位置
strings.Join(a[]string, sep string)	join操作
```

### 类型转换

T(表达式)
```
var a, b = 3, 4
var c int
// math.Sqrt()接收的参数是float64类型，需要强制转换
c = int(math.Sqrt(float64(a*a + b*b)))
fmt.Println(c)
```

### if条件判断特殊写法

```
func ifDemo2() {
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}
}
```

### 循环for
break、goto、return、panic语句强制退出循环

```
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

无限循环
for {
    循环体语句
}
```

### 数组
var 数组变量名 [元素数量]T

```
// 定义一个长度为3元素类型为int的数组a
var a [3]int

var a [3]int
var b [4]int
a = b //不可以这样做，因为此时a和b是不同的类型

//编译器根据初始值的个数自行推断数组的长度
var numArray = [...]int{1, 2}

a := [...]int{1: 1, 3: 5}
fmt.Println(a)                  // [0 1 0 5]

//使用指定的初始值完成初始化
var cityArray = [3]string{"北京", "上海", "深圳"} 

var a = [...]string{"北京", "上海", "深圳"}
// 方法1：for循环遍历
for i := 0; i < len(a); i++ {
    fmt.Println(a[i])
}

// 方法2：for range遍历
for index, value := range a {
    fmt.Println(index, value)
}

a := [3][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a) //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆
}

定长的是值类型、非定长是引用类型
func main() {
	a := [3]int{10, 20, 30}
	modifyArray(a)
	fmt.Println(a)//[10 20 30]
}

func modifyArray(x [3]int) {
	x[0] = 100
}

func main() {
	a := []int{10, 20, 30}
	modifyArray(a)
	fmt.Println(a)//[100 20 30]
}

func modifyArray(x []int) {
	x[0] = 100
}
```

### 切片

make([]T, size, cap)

```
a := make([]int, 2, 10)
fmt.Println(a)      //[0 0]
fmt.Println(len(a)) //2
fmt.Println(cap(a)) //10

var a []string              //声明一个字符串切片
```

### map

map是一种无序的基于key-value的数据结构，Go语言中的map是引用类型

```
scoreMap := make(map[string]int, 8)
scoreMap["张三"] = 90
scoreMap["小明"] = 100
fmt.Println(scoreMap)
fmt.Println(scoreMap["小明"])
fmt.Printf("type of a:%T\n", scoreMap)

//判断key是否存在
v, ok := scoreMap["张三"]
if ok {
    fmt.Println(v)
} else {
    fmt.Println("not found")
}

//遍历
for k, v := range scoreMap {
    fmt.Println(k)
    fmt.Println(v)
}
for k := range scoreMap {
    fmt.Println(k)
}
for _, v := range scoreMap {
    fmt.Println(v)
}

//删除键值对
delete(scoreMap, "张三")

for k, v := range scoreMap {
    fmt.Println(k)
    fmt.Println(v)
}
```

### 函数

```
func 函数名(参数)(返回值){
    函数体
}
```

```
func intSum(x int, y int) int {
	return x + y
}

//可变参数
func intSum2(x ...int) int {
	fmt.Println(x) //x是一个切片
	sum := 0
	for _, v := range x {
		sum = sum + v
	}
	return sum
}

//定义函数类型
type calculation func(int, int) int
func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
var c calculation
c = add


//匿名函数
// 将匿名函数保存到变量
add := func(x, y int) {
    fmt.Println(x + y)
}
add(10, 20) // 通过变量调用匿名函数

//自执行函数：匿名函数定义完加()直接执行
func(x, y int) {
    fmt.Println(x + y)
}(10, 20)
```

### defer语句

```
fmt.Println("start")
defer fmt.Println(1)
defer fmt.Println(2)
defer fmt.Println(3)
fmt.Println("end")

start
end
3
2
1
```

```
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func main() {
	x := 1
	y := 2
	defer calc("AA", x, calc("A", x, y))
	x = 10
	defer calc("BB", x, calc("B", x, y))
	y = 20
}

A 1 2 3
B 10 2 12
BB 10 12 22
AA 1 3 4
```

### 内置函数介绍

```
close	主要用来关闭channel
len	用来求长度，比如string、array、slice、map、channel
new	用来分配内存，主要用来分配值类型，比如int、struct。返回的是指针
make	用来分配内存，主要用来分配引用类型，比如chan、map、slice
append	用来追加元素到数组、slice中
panic和recover	用来做错误处理
```

### 指针

```
//指针取值
a := 10
b := &a // 取变量a的地址，将指针保存到b中
fmt.Printf("type of b:%T\n", b)
c := *b // 指针取值（根据指针去内存取值）
fmt.Printf("type of c:%T\n", c)
fmt.Printf("value of c:%v\n", c)

type of b:*int
type of c:int
value of c:10

make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身；
而new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
```

### 结构体

```
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

```
//类型定义
type NewInt int
//类型别名
type MyInt = int

func structDemo() {
	var a NewInt
	var b MyInt
    fmt.Printf("type of a:%T\n", a) //type of a:main.NewInt
	fmt.Printf("type of b:%T\n", b) //type of b:int
}
```

```
type person struct {
    name string
    city string
    age int8
}

type person2 struct {
    name, city string
    age int8
}

var p1 person
p1.name = "dp"
p1.city = "shanghai"
p1.age = 18
fmt.Printf("p1=%v\n", p1)
fmt.Printf("p1=%#v\n", p1)
fmt.Printf("p1: %T\n", &p1)//p1: *main.person

//指针类型结构体
var p2 = new(person)
p2.name = "dp2"
p2.city = "shanghai2"
fmt.Printf("p2: %T\n", p2)//p2: *main.person

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
```

- 匿名结构体

```
var user struct{Name string; Age int}
user.Name = "xx"
user.Age = 18
fmt.Printf("%#v\n", user)
fmt.Printf("user: %T\n", &user)//user: *struct { Name string; Age int }
```

- 结构体函数
```
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
```

```
//Person 结构体
type Person struct {
	name string
	age  int8
}

//NewPerson 构造函数
func NewPerson(name string, age int8) *Person {
	return &Person{
		name: name,
		age:  age,
	}
}

//Dream Person做梦的方法
func (p Person) Dream() {
	fmt.Printf("%s的梦想是学好Go语言！\n", p.name)
}

func main() {
	p1 := NewPerson("小王子", 25)
	p1.Dream()
}
```

- 结构体的“继承”

```
//Animal 动物
type Animal struct {
	name string
}

func (a *Animal) move() {
	fmt.Printf("%s会动！\n", a.name)
}

//Dog 狗
type Dog struct {
	Feet    int8
	*Animal //通过嵌套匿名结构体实现继承
}

func (d *Dog) wang() {
	fmt.Printf("%s会汪汪汪~\n", d.name)
}

func main() {
	d1 := &Dog{
		Feet: 4,
		Animal: &Animal{ //注意嵌套的是结构体指针
			name: "乐乐",
		},
	}
	d1.wang() //乐乐会汪汪汪~
	d1.move() //乐乐会动！
}
```

- 序列化

```
//Student 学生
type Student struct {
	ID     int
	Gender string
	Name   string
}

//Class 班级
type Class struct {
	Title    string
	Students []*Student
}

func main() {
	c := &Class{
		Title:    "101",
		Students: make([]*Student, 0, 200),
	}
	for i := 0; i < 10; i++ {
		stu := &Student{
			Name:   fmt.Sprintf("stu%02d", i),
			Gender: "男",
			ID:     i,
		}
		c.Students = append(c.Students, stu)
	}
	//JSON序列化：结构体-->JSON格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed")
		return
	}
	fmt.Printf("json:%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}

```

### package

- 首字母小写，外部包不可见，只能在当前包内使用
- 首字母大写，外部包可见，可在其他包中使用(变量、函数、结构体、接口)
- 单行导入
```
import "包1"
import "包2"
```

- 多行导入
```
import (
    "包1"
    "包2"
)
```

- 自定义包名

import 别名 "包的路径"

```
import "fmt"
import m "github.com/Q1mi/studygo/pkg_test"

func main() {
	fmt.Println(m.Add(100, 200))
	fmt.Println(m.Mode)
}
```

```
import (
    "fmt"
    m "github.com/Q1mi/studygo/pkg_test"
 )

func main() {
	fmt.Println(m.Add(100, 200))
	fmt.Println(m.Mode)
}
```

- init()函数

![](https://www.liwenzhou.com/images/Go/package/init01.png)
![](https://www.liwenzhou.com/images/Go/package/init02.png)


### 接口

接口在命名时，一般会在单词后面添加er

当方法名首字母是大写且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包（package）之外的代码访问。

```
type 接口类型名 interface{
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

```
type Mover interface {
	move()
}

type dog struct {}

func (d dog) move() {
	fmt.Println("狗会动")
}

var x Mover
var wangcai = dog{} // 旺财是dog类型
x = wangcai         // x可以接收dog类型
var fugui = &dog{}  // 富贵是*dog类型
x = fugui           // x可以接收*dog类型
x.move()
```

- 空接口

```
// 定义一个空接口x
var x interface{}
s := "Hello 沙河"
x = s
fmt.Printf("type:%T value:%v\n", x, x)
i := 100
x = i
fmt.Printf("type:%T value:%v\n", x, x)
b := true
x = b
fmt.Printf("type:%T value:%v\n", x, x)

// 空接口作为函数参数
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
```

### 反射

```
func reflectX(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Printf("type: %v, kind: %v\n", t.Name(), t.Kind())
	v := reflect.ValueOf(x)

	if v.Kind() == reflect.Int64 {
		v.SetInt(100) //修改的是副本，reflect包会引发panic
	}
	// 反射中使用 Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200)
	}
}
```

- isNil()

报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic

- isValid()

返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic

- 结构体反射

```
Field(i int) StructField	根据索引，返回索引对应的结构体字段的信息。
NumField() int	返回结构体成员字段数量。
FieldByName(name string) (StructField, bool)	根据给定字符串返回字符串对应的结构体字段的信息。
FieldByIndex(index []int) StructField	多层成员访问时，根据 []int 提供的每个结构体的字段索引，返回字段的信息。
FieldByNameFunc(match func(string) bool) (StructField,bool)	根据传入的匹配函数匹配需要的字段。
NumMethod() int	返回该类型的方法集中方法的数目
Method(int) Method	返回该类型方法集中的第i个方法
MethodByName(string)(Method, bool)	根据方法名返回该类型方法集中的方法
```

```
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

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod())
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}
```

### 并发 

```
var wg sync.WaitGroup

func hello(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func main() {

	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}
```

### 网络编程

### 单元测试

```
类型	格式	作用
测试函数	函数名前缀为Test	测试程序的一些逻辑行为是否正确
基准函数	函数名前缀为Benchmark	测试函数的性能
示例函数	函数名前缀为Example	为文档提供示例文档
```

```
//Demo_test.go

package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	excepted := 3
	got := 1 + 2
	if got != excepted {
		t.Errorf("excepted:%v, got:%v", excepted, got) // 测试失败输出错误提示
	}
}
```

### 常用标准库

- math

```
//开方
math.Sqrt
```

- fmat与格式化占位符

```
//Print 标准输出
//Printf 格式化输出
//Println 输出的内容结尾添加一个换行符

//输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容
//Fprint
//Fprintf
//Fprintln

//输出到字符串
//Sprint
//Sprintf
//Sprintln

//Errorf 格式化字符串对应的错误

//格式化占位符
%v	值的默认格式表示
%+v	类似%v，但输出结构体时会添加字段名
%#v	值的Go语法表示
%T	打印值的类型
%%	百分号
%t 布尔值 true/false
整型
%b	表示为二进制
%c	该值对应的unicode码值
%d	表示为十进制
%o	表示为八进制
%x	表示为十六进制，使用a-f
%X	表示为十六进制，使用A-F
%U	表示为Unicode格式：U+1234，等价于”U+%04X”
%q	该值对应的单引号括起来的go语法字符字面值，必要时会采用安全的转义表示
浮点数与复数
%b	无小数部分、二进制指数的科学计数法，如-123456p-78
%e	科学计数法，如-1234.456e+78
%E	科学计数法，如-1234.456E+78
%f	有小数部分但无指数部分，如123.456
%F	等价于%f
%g	根据实际情况采用%e或%f格式（以获得更简洁、准确的输出）
%G	根据实际情况采用%E或%F格式（以获得更简洁、准确的输出）
字符串和[]byte
%s	直接输出字符串或者[]byte
%q	该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
%x	每个字节用两字符十六进制数表示（使用a-f
%X	每个字节用两字符十六进制数表示（使用A-F）
指针
%p	表示为十六进制，并加上前导的0x
宽度标识符
%f	默认宽度，默认精度
%9f	宽度9，默认精度
%.2f	默认宽度，精度2
%9.2f	宽度9，精度2
%9.f	宽度9，精度0
```

```

bufio.NewReader//从终端获取完整内容

reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
fmt.Print("请输入内容：")
text, _ := reader.ReadString('\n') // 读到换行
text = strings.TrimSpace(text)
fmt.Printf("%#v\n", text)
```

- time

```
//获取当前时间
now := time.Now()
fmt.Printf("current time: %v\n", now)

year := now.Year()     //年
month := now.Month()   //月
day := now.Day()       //日
hour := now.Hour()     //小时
minute := now.Minute() //分钟
second := now.Second() //秒
fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)

//获取当前时间戳
timestamp1 := now.Unix()     //时间戳
timestamp2 := now.UnixNano() //纳秒时间戳
fmt.Printf("current timestamp1:%v\n", timestamp1)
fmt.Printf("current timestamp2:%v\n", timestamp2)

//时间间隔(纳秒为单位)
// const (
// 	Nanosecond  Duration = 1
// 	Microsecond          = 1000 * Nanosecond
// 	Millisecond          = 1000 * Microsecond
// 	Second               = 1000 * Millisecond
// 	Minute               = 60 * Second
// 	Hour                 = 60 * Minute
// )

//时间操作
later := now.Add(time.Hour) // 当前时间加1小时后的时间
fmt.Println(later)

d := now.Sub(later)
fmt.Println(d)

fmt.Printf("时间是否相等: %t\n", later.Equal(now))
fmt.Printf("时间是否在xx之前: %t\n", later.Before(now))
fmt.Printf("时间是否在xx之后: %t\n", later.After(now))

//定时器
ticker := time.Tick(time.Second) //定义一个1秒间隔的定时器
for i := range ticker {
	fmt.Println(i) //每秒都会执行的任务
	break
}

//时间格式化(格式化的模板为Go的出生时间2006年1月2号15点04分 Mon Jan)
fmt.Println(now.Format("2006-01-02 15:04:05"))
```

- 获取命令行参数

```
package main

import (
	"fmt"
	"os"
)

//os.Args demo
func main() {
	//os.Args是一个[]string
	if len(os.Args) > 0 {
		for index, arg := range os.Args {
			fmt.Printf("args[%d]=%v\n", index, arg)
		}
	}
}
```

- 文件操作

```
//打开关闭文件
file, err := os.Open("./main.go")
if err != nil {
	fmt.Println("open file failed!, err:", err)
	return
}
defer file.Close()
//读取文件
var tmp = make([]byte, 128)
n, err := file.Read(tmp)
if err == io.EOF {
	fmt.Println("文件读完了")
	return
}
if err != nil {
	fmt.Println("read file failed, err:", err)
	return
}
fmt.Printf("读取了%d字节数据\n", n)
fmt.Println(string(tmp[:n]))

//循环读取文件所有内容
var content []byte
for {
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了")
		break
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	content = append(content, tmp[:n]...)
}
fmt.Println(string(content))

// bufio按行读取示例
reader := bufio.NewReader(file)
for {
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		if len(line) != 0 {
			fmt.Println(line)
		}
		break
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Print(line)
}

//ioutil读取整个文件
buf, err := ioutil.ReadFile("./main.go")
if err != nil {
	fmt.Println("read file failed, err: ", err)
	return
}
fmt.Println(string(buf))

//file2, err := os.OpenFile("output.txt", os.O_CREATE | os.O_TRUNC | os.O_WRONLY, 0666)
//写入
// file.Write([]byte(str))       //写入字节切片数据
// file.WriteString("hello 小王子") //直接写入字符串数据

// writer := bufio.NewWriter(file)
// writer.WriteString("hello沙河\n") //将数据先写入缓存
// writer.Flush() //将缓存中的内容写入文件

//err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
```

- strconv包

```
//string -> int
s1 := "100"
i1, err := strconv.Atoi(s1)
if err != nil {
	fmt.Println("can't convert to int")
} else {
	fmt.Printf("type: %T value: %#v\n", i1, i1)
}

//int -> string
i2 := 200
s2 := strconv.Itoa(i2)
fmt.Printf("type:%T value:%#v\n", s2, s2)

//parse函数
//string -> bool
s3 := "false"
b1, _ := strconv.ParseBool(s3)
fmt.Printf("type:%T value:%#v\n", b1, b1)

//string -> int
s4 := "100"
i3, err := strconv.Atoi(s4)
if err != nil {
	fmt.Println("can't convert to int")
} else {
	fmt.Printf("type: %T value: %#v\n", i3, i3)
}
//ParseUint 类似ParseInt但不接受正负号，用于无符号整型

//string -> float
s5 := "3.14"
f1, _ := strconv.ParseFloat(s5, 2)
fmt.Printf("type:%T value:%#v\n", f1, f1)

//Format系列函数
//FormatBool
//FormatInt
var i4 int64 = 0xff
fmt.Printf("%#v %#v %#v\n", strconv.FormatInt(i4, 2), strconv.FormatInt(i4, 10), strconv.FormatInt(i4, 16))
//FormatUint
//FormatFloat
```

### 问题

```
type student struct {
	name string
	age  int
}

func main() {
	m := make(map[string]*student)
	stus := []student{
		{name: "小王子", age: 18},
		{name: "娜扎", age: 23},
		{name: "大王八", age: 9000},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}
大王八 => 大王八
小王子 => 大王八
娜扎 => 大王八
???
https://segmentfault.com/a/1190000017527311?utm_campaign=studygolang.com&utm_medium=studygolang.com&utm_source=studygolang.com
```