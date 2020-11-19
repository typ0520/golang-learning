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

### 数组(值类型)
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
```

### 变长数组

make([]T, size, cap)

```
a := make([]int, 2, 10)
	fmt.Println(a)      //[0 0]
	fmt.Println(len(a)) //2
	fmt.Println(cap(a)) //10

var a []string              //声明一个字符串切片
```

