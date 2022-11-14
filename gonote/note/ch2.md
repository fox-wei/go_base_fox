# Go 语言基础语法


## 变量声明

在编程语言中，为了方便操作内存特定位置的数据， 我们使用一个特定的名字与特定位置的内存块绑定，这个名字称为**变量**。
**变量所绑定的内存区域要有一个明确的边界的**，即通过这个变量，我们要明确知道是4个字节或8字节内存，编译器或解释器要明确知道。

动态语言的解释器在运行的时候通过变量赋值分析，自动确定变量的边界，并且在动态语言中，一个变量可以在运行时被赋予大小不同的边界。静态语言的编译器则需要使用者明确提供变量的边界。

Go语言变量声明方式：
```go
//在函数外声明变量
var name type = value
var a int //a=0 自动赋予类型的零值
var a = 10 //自动推断变量类型

//在函数内声明变量的以一种方式
func main() {
    b := "fox" //局部变量
}

//多变量声明
var (
    x = 10
    y string = "fox"
)

//官方推荐使用
var (
    x = int32(10)
    y = int8(24)
)

//类型推断
var (
    xx = "fox"
    yy = 24
)


```

Go语言类型的零值，即相应类型的默认值：
|   内置类型    |   零值（默认值）  |
|----|----|
|   所有整型类型    | 0 |
|   浮点类型    |   0.0 |
|   布尔类型    | false |
|   字符串类型  |   ""  |
|   指针、接口、切片、channel、map和函数类型    |   nil |

数组、结构体类型变量的零值就是他们组成元素都为零值的结果。

Go语言的变量类型分为两类：一类是包级变量(package varible)，变量在函数外声明(var声明)，首字母大写变量为导出变量，视为全局变量；首字母为小写的变量则包外不可用；函数内定义的变量为局部变量。


## 代码块与作用域

Go代码块是包裹在一对大括号内部的声明和语句序列。如果大括号内部无语句则叫做空代码块。

![alt](https://static001.geekbang.org/resource/image/3d/85/3d02138cf8f8a7a85fe0cfe5c29a6585.jpg?wh=1920x1047)


宇宙代码块(Universe Block)，囊跨范围最大，所有Go源码都在这个隐式代码块中。
宇宙代码块嵌套包代码块(Package Block)，每个Go包含有所有Go源码。包代码块嵌套若干文件代码块(File Block)，每个文件代码块对应着一个.go文件。
```go
... ...
 var a int = 2020
  
 func checkYear() error {
     err := errors.New("wrong year")
 
     switch a, err := getYear(); a { //屏蔽变量a=2020,外部err
     case 2020:
         fmt.Println("it is", a, err)
     case 2021:
         fmt.Println("it is", a)
         err = nil //不是外部的err
     }

     fmt.Println("after check, it is", a)
     return err //err != nil
 }
 
 type new int //屏蔽预定义标识符new
 
 func getYear() (new, error) {
     var b int16 = 2021
     return new(b), nil
 }

 func main() {
     err := checkYear()
     if err != nil {
         fmt.Println("call checkYear error:", err)
         return
     }
     fmt.Println("call checkYear ok")
 }
```
 在上述代码存在三个问题：
 1. 函数外部定义了a，内部又定义a,导致外部a=2020被屏蔽；
 2. swith代码块外部定义了err，内部定义了err，内部err值的改变不能影响外部的err，err无法设置为nil;
 3. int定义为预定义标识符new，预定义标识符的含义发生改变。

 ## 基本数据类型
 Go语言类型分为三大类型：基本数据类型，复合数据类型（结构体）和接口类型。

整型分为平台无关型和平台相关型，主要区别在不同平台下长度是否一致。例如int8长度为一个字节。平台无关型整型分为有符号整型(int8-int64)和无符号整型(uint8-uint64)。两者本质差别是在于**最高位二进制**是否被解释为符号位，并且负数是用2的补码表示，先取反最后加1.

平台相关型int, uint, uintptr(大到足以存储任意一个指针的值)。
字面值与格式化输出：
```go
d1 := 0b10000001 // 二进制，以"0b"为前缀
d2 := 0B10000001 // 二进制，以"0B"为前缀
e1 := 0o700      // 八进制，以"0o"为前缀
e2 := 0O700      // 八进制，以"0O"为前缀
c2 := 0Xddeeff // 十六进制，以"0X"为前缀
```
1.13以后增加数字分隔符"`_`",分隔数字提供可读性。

整型变量输出为不同进制形式：
`%b`:二进制；`%d`:十进制；`%o|O`:八进制；`%x|X`:十六进制。

**浮点数**
IEEE 754规范
|   符号位(S)   |   阶码(E) |   尾数(M) |
|----|----|----|
|   sign    |   exponent    |   maintissa   |
浮点数的值：(-1)^s x 1.M x  2^(E-offset)

单精度(float32)和双精度(float64)浮点数在阶码和尾数上不同：
|  浮点类型   |   符号位(bit位数) |   阶码(bit) |  阶码偏移值| 尾数(bit)|
|----|----|----|----|----|
|   float32    |   1   |  8   | 127 |23|
|   float64    |   1   |  11| 1023  |   52  |
十进制浮点数转换成二进制浮点数(IEEE 754):
1. 整数和小数转换为二进制；
2. 小数点右移，直到整数仅有一个1,移动位数为指数；
3. 计算阶码：阶码=指数+偏移值

## 字符串
c 语言中非原生字符串带来的坏处：
1. 无法进行类型校验，存在不安全因素
2. 结尾要使用"\0",防止溢出
3. 采用字符数组方式定义，字符串可变，高并发场景存在隐患
4. 获取长度代价较大
5. 内部无对非ACCII字符的支持

Go原生字符串的好处：
1. 字符串数据不可变，提高字符串并发安全性和存储利用率
2. 结尾没有"\0",消除获取长度的开销
3. 原生支持所见即所得字符串，即非解释字符串，用反引号表示
4. 采用UTF-8编码方式，消除源码在不同环境的乱码情况
字符串的本质：
```go
type StringHeader struct {
    Data uintptr
    Len int
}
```

Go语言的组成从字节的角度看：Go字符串由一个可空的字节序列组成，字节的个数为字符串的长度，Go采用UTF-8编码，英文为一个字节，汉字为3个字节；从字符角度看：字符串由一个可空的字符组成。
所以在使用`len`方法的时候，返回的是字符串的字节个数而不是字符个数；
字符串迭代，普通for和for_range；普通for方式迭代字符串的每一个字节而不是字符；而for_range方式则每轮迭代字符串的Unicode的码点及该字符串的偏移值；要获取非ACCII字符串的长度可以使用utf-8包中的RuneCountInStringn方法。
```go
    s := "我是韦宗富" //非ASCII
	fmt.Println("for-test")
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i]) //字节；3*5
	}

	fmt.Printf("\nfor range test\n")
	for _, a := range s {
		fmt.Printf("%s", string(a)) //汉字字节值转化；5
	}
```
## Go常量
常量是一种在源码编译期间被创建的语法元素，它的值在程序的生命周期内保持不变。Go常量的“创新”：
1. 支持无类型常量，即可以不需要类型声明；
```go
const m = "黄莹"
```

2. 支持隐式自动转型；
```go
type myInt int
const n  = 8 
func main() {
    var s myInt = 4
    fmt.Println(s+n) //n自动转换为myInt类型
}
```
但是要注意基本数据类型的溢出问题。

3. 可以用于实现枚举

 枚举类型的本质上就是一个有限数量常量构成的集合。
 Go使用const块的两个特性，自动重复上上一行，引入const块中的行偏移值指示器iota;
 第一特性重复上一行：
 ```go
 const (
     Apple, Banana = 11, 22
     Strawberry, Grape //11, 22
     A, B //11, 22
 )
 ```
 第二个特性：**iota行偏移值**<br>
 iota是Go语言的一个预定义标识符，它表示的是const声明块（包括单行声明）中，每个常量所处位置在块中的偏移值（从零开始）。
```go
const (
    a = iota //iota = 0
    b //1
    c //2
)

//另一种形成
const (
    x = 1 + iota //1
    y  //1+1=2
    z //1+2=3
    m = 7 //iota=3
    n //iota=4
)
```
 ## 同构复合类型：从定长数组到变长切片

Go语言的数组是一个长度固定的、由同构类型元素组成的连续序列。
数组声明方式：var s [N]type;or var s [...]type = {s1, s2, ...};多维数组理解方式，可以进行拆分最后还是一个一维数组：例如数组var mArr\[2\]\[3\]\[4\]int
![alt](https://static001.geekbang.org/resource/image/27/d3/274f3fc9e753b416f5c0615d256a99d3.jpg?wh=1920x1047)

数组类型变量是一个整体，这就意味着一个数组变量表示的是整个数组。这点与 C 语言完全不同，在 C 语言中，数组变量可视为指向数组第一个元素的指针。在go中数组是按值传递。

Go的切片是一个三元结构：
```go
type slice struct {
    array unsafe.Pointer
    len   int
    cap   int
}
```
array是指向底层数组的指针，另外两个分别表示长度和容量，切片可以简单理解为是对数组的封装。
切片创建的三种方式；

```go
func main() {
    var s1 []int //未初始化，nil;make初始化
    var s2 []int = {1, 2, 3, 4}
    var s3 := []int{1, 2, 3}
    s4 := make([]int, 4, 5) //len, cap, 默认len==cap
}
```
初值为零值nil的切片类型变量，可以借助内置的append的函数进行操作，这种在Go语言中被称为“零值可用”。
还有一种方式采用array[low:high:max];获取的切片长度为high-low;容量为max-low,array是数组或者切片。

```go
arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
sl := arr[3:7:9]
```
![alt](https://static001.geekbang.org/resource/image/96/34/96407488137f185d860c6c3624072f34.jpg?wh=1920x1047)
通过这种方式创建切片，他们是公用同一底层数组，所以两者公用区域数组值的改变，另一个也会改变，例如：
```go
s1[0] = 13
fmt.Println(arr[3]) //4-->13
```
切片是可以扩容，使用append方法实现：
```go
s := []int{1, 2}
s = append(s, 5)
```
使用append扩容，当容量和长度相同时会长度*2增加容量，这样这个切片会与原先底层数组“关系解绑"。

## 复合数据类型：map类型
map是一组无序的键值对，用key和value表示,表示代码：
```go
v := map[type]type{}
```
对于key的类型要求，**要求支持"=="和"!="两种比较运算符，所以key类型不能是函数类型，切片类型，map自身类型，对于复合类型struct，则根据其字段是否支持类型比较确定。**

map不支持**零值可用**，切片和函数类型也是如此，只支持nil比较而不支持同类型比较；使用map必须要先初始化，初始化有两种方式：
```go
s := make(map[type][type] [,cap]) //容量为可选方式，map不受容量限制，可动态扩容
x := map[type]type
```
Go允许忽略字面值中的元素类型：
```go
type person struct {
    name string
    age int
}

func main() {
    s := map[person]int {
        {"fox", 21}:1,
        {"ying", 20}:2,
    }
}
```
**map基本操作**

插入键值对：把value赋值给map对应的key：
```go
m := make(map[string]int)
m["1"] = 1
```
某个key已经存在，我们插入一个新值，则覆盖原先key的值。<br>
获取键值对数量：len方法，不能使用cap获取容量。<br>
查找和读取数据：

查找时，我们无法判断这个key是否存在map中；因为map的机制中，key不存在map中 ，我们会获取这个value的零值。所以要确定key是否存在map，使用"comma ok“方法确定：

```go
s := make(map[string]int)
v, ok := s["1"]
if !ok {
    //不存在
}
```
在 Go 语言中，请使用“comma ok”惯用法对 map 进行键查找和键值读取操作。

删除数据：delete(map, key);遍历map，使用for range方法：
```go
s := map[string]int{"fox":1, "yi":2}
for k, v := range s {
    //
}
```
对map进行多次遍历，每次元素的顺序都不相同。所以对map，程序逻辑千万不要依赖遍历map所得到的元素次序。

## 复合数据类型：用结构体建立对真实世界的抽象
编写程序是为解决真实世界的问题，对真实世界事物体的重要属性进行提炼，并映射到程序世界中，这就是所谓的对真实世界的抽象。

Go如何定义一个新类型：
1. 类型定义，type name 基本类类型
```go 
type myInt int
```
其中int可以称为底层类型，即采用这种方法定义新类型，go语言的原生类型称为**底层类型**。**底层类型用来判断两个类型本质上是否相同。**
```go
type t1 int
type t2 t1

func main() {
    var n t1 = 5
    var m t2 = t2(n) //ok
}
```
上述两个类型，本质上是相同的类型，他们的变量可以**显示转型**进行相互赋值；相反，两个本质上是不同的类型，不能进行相互转换。<br>
除了基于已有原生类型，符合类型也可以进行定义新类型：

```go
type M map[string]int
type N []int
```
2. 类型别名,type T = S
T和S本质上是同一个类型，例如：
```go
type T = string //别名
func main() {
    var s string = "ying"
    var s1 T = s //ok
}
```

结构体定义方式：
```go
type name struct {
    filed1 type
    filed2 type
    //...
}
```
字段首字母大小写决定这个字段是否包外可导，即包是否可以通过结构体名直接访问该字段。如果设置为小写，但是包外需要访问，我们可以设置一个方法去访问该字段。
我们还可以用**空标识符“_”作为结构体类型定义中的字段名称**。这样以空标识符为名称的字段，不能被外部包引用，甚至无法被结构体所在的包使用。
特殊结构体：

1. 空结构体
```go
type Empty struct {} //不包含任何字段
var s Empty
fmt.Println(unsafe.Sizeof(s)) //0
```
**空结构体变量内存占用为0，基于空结构体内存零开销的特性，Go经常使用空结构体作为“事件”信息进行Goroutine之间的通信。**

```go
var c = make(chan struct{})
c <-struct{}
```
2. 其他结构体作为自定结构体的类型
```go
type person struct {
    name string 
    age int
}

type book struct {
    title string
    author person
}

type student struct {
    person //匿名字段
    score float32
}
```
访问结构体中的字段，`对象名.字段名`即可；内部结构体字段访问，我们有时候不用对象名也直接访问。如果内部又同名字段，还是要用内部对象名访问，防止访问错误。
**结构体类型T不能在内部定义自己，但是\*T,[]T, map[type]T是可以在内部定义，因为这些都是引用类型，内存空间确定所以可以在内部定义。**

Go 结构体类型由若干个字段组成，当这个结构体类型变量的各个字段的值都是零值时，我们就说这个结构体类型变量处于零值状态。**采用零值初始化得到的零值变量，是有意义的，而且是直接可用的，我称这种类型为“零值可用”类型。**

map和slice零值状态不能直接使用，需要使用初始化(make)。

```go 
//“零值可用”的运用
var mu sync.Mutex
mu.Lock()
mu.Unlock()
```
字面值初始化结构体，可以直接赋值结构体：
```go
type book struct {
    name string
    page int
}

func main() {
    s := book{"自学是门手艺", 200}
}
```
但是这种方式赋值对于包外非导出字段初始化会出现错误；官方推荐方式是file:value的方式赋值，这样可以不考虑字段顺序：
```go
s := book{name:"go语言程序设计", page:200}
```

**内存布局**



## 控制结构

程序 = 数据结构 + 算法
数据机构对应go语言中的基本类型和复合类型；算法是对真实世界运作规律的抽象，是解决真实世界中问题的步骤。

**if语言使用**
基本使用方式：
```go
//二元分支
if expression1 {
    //
} else {
    //
}

//多分支表示,不推荐
if expression1 {
    //
} else if expression2 {
    //
} else if .. {

} else {

}
```
if语句的自用变量：
```go
if a := f(); a > 0 {
    //
} else {
    fmt.Println(a)
}
```
因为变量a在if作用域内所以内部可以使用，但需要注意变量遮蔽问题(if内部声明同名变量)。

if的快乐路径原则：
在日常编码中要减少多分支结构，甚至是二分支结构的使用，这会有助于我们编写出优雅、简洁、易读易维护且不易错的代码。
```go
func doSomething() error { 
    if errorCondition1 { 
        // some error logic ... ... 
        return err1 
    } 
    // some success logic ... ... 
    if errorCondition2 { 
        // some error logic ... ... return err2 
    } 
    // some success logic ... ... 
    return nil
}
```
- 仅使用单分支结构
- 判断为false，在单分支快速返回
- 正常逻辑在代码布局靠“左”
- 函数到最后表示成功
我们尽量将命中概率高的语句写在前面。

**人们常说时间会改变一切，但事实上你得自己来**

**for循环**
for循环使用形式：
```go
//一般形式
for i := 0(前置语句); i< count(判断条件); i++(后置语句) {
    //循环体
} 

//while
for ;判断条件; {
    //循环体
}

//无限循环
for {
    //
}

//声明多循环变量
for i,j:=0, 0; i<10 && j <5; i+=1, j+=1 {
    //
}
```

另一种方式：for range<br>

**这种迭代使用于数组，切片，map，channel，string类型，而且迭代对象是对象的副本，所以对于非引用类型来说，迭代过程修改值是无效的。**

for 语句的"坑”与避坑方法
1. for range循环变量重用：
```go
func main() {
    var m = []int{1, 2, 3, 4, 5}  
             
    for i, v := range m {
        go func() {
            time.Sleep(time.Second * 3)
            fmt.Println(i, v) //5*(4, 5)
        }() //优化(i, v)作为参数
    }

    time.Sleep(time.Second * 10)
}
```
理想结果是0,1; 1, 2;...
但是因为goroutine的存在，结果是4, 5; 4, 5...
i,v变量只声明一次，实际上被不断重用；所以这是一片临界区，goroutine竞争导致，只输出最后的结果。

2. 参与循环的是 range 表达式的副本

  ```go
  func main() {
      var a = [5]int{1, 2, 3, 4, 5}
      var r [5]int
  
      fmt.Println("original a =", a)
  
      for i, v := range a {
          if i == 0 {
              a[1] = 12
              a[2] = 13
          }
          r[i] = v
      }
  
      fmt.Println("after for range loop, r =", r) //1, 2, 3, 4, 5
      fmt.Println("after for range loop, a =", a) //1, 12, 13, 4, 5
  }
  ```

  因为for-range是a的副本，虽然改变了外部a的值，但是副本不受影响；所以r的值与改变之前相同；如果a是切片，则改变相同，因为切片是引用类型，副本是地址。

3. 遍历map中元素的随机性
```go
var m = map[string]int{
    "tony": 21,
    "tom":  22,
    "jim":  23,
}

counter := 0
for k, v := range m {
    if counter == 0 {
        delete(m, "tony")
    }
    counter++
    fmt.Println(k, v)
}
fmt.Println("counter is ", counter)
```
结果可能会输出tony的信息，因为如果tony第一个随机访问到则输出，否则不会有。同理添加新的值，可能不会输出；这取决于添加元素的访问位置。

**switch语句**
一般形式：
```go
switch initStmt; expr {
    case expr1:
        // 执行分支1
    case expr2:
        // 执行分支2
    case expr3_1, expr3_2, expr3_3:
        // 执行分支3
    case expr4:
        // 执行分支4
    ... ...
    case exprN:
        // 执行分支N
    default: 
        // 执行默认分支
}
```
switch 语句各表达式的求值结果可以为各种类型值，只要它的类型支持比较操作就可以了。

**type switch**
```go
func main() {
    var x interface{} = 13 //空接口，表示任意类型
    switch x.(type) { //x必须是接口变量
    case nil:
        println("x is nil")
    case int:
        println("the type of x is int")
    case string:
        println("the type of x is string")
    case bool:
        println("the type of x is string")
    default:
        println("don't support the type")
    }
}
```
## 函数
### 函数基础
函数是唯一一种基于特定输入，实现特定任务并可返回任务执行结果的代码块（go方法本质上是函数)。
go函数组成：关键字`func`，函数名，参数列表，返回值列表以及函数体：

```go
func hello(name string) string {
    return "hello " + name 
}

//函数另一种变换，匿名函数的一种形式
var hello = func(name string) string {
     return "hello " + name 
}
```
函数声明中的函数名其实就是变量名，函数声明中的`func`关键字、参数列表和返回值列表共同构成了函数类型。
**而参数列表与返回值列表的组合也被称为函数签名，它是两个函数类型是否相同的决定因素。**
我们可以忽略参数列表和返回值列表的参数名称，函数签名相同则两个函数类型是相同类型。

**每个函数声明所定义的函数，仅仅是对应的函数类型的一个实例。**

Go语言中，函数参数的传递采用**值传递**，将实际参数在内存中的表示逐位拷贝到形式参数中。基本数据类型，数组，结构体在内存中表示就是自身数据，因此值传递就是他们自身，传递开销与自身大小成正比。
切片，字符串,map，它们的表示是它们数据内容的“描述符”，它们大小固定，开销较小，这种方式叫做“浅拷贝”。

变长参数表示：`...type`，本质上是切片；可以传入多个同类型参数；所以可直接传入切片,方法`name...`。<br>
Go支持多返回值，返回值参数有名字的称为具名返回值，根据特殊场景的需求使用，例如函数使用`defer`；大部分情况时非具名返回值，即无参数名。

**函数是一等公民**

*编程语言一等公民：*

> 如果一门编程语言对某种语言元素的创建和使用没有限制，我们可以像对待值（value）一样对待这种语法元素，那么我们就称这种语法元素是这门编程语言的“一等公民”。拥有“一等公民”待遇的语法元素可以存储在变量中，可以作为参数传递给函数，可以在函数内部创建并可以作为返回值从函数返回。

1. 函数可以存储在变量中；
2. 支持在函数内创建并返回；
```go
func setup(task string) func() {
    println("do some setup stuff for", task)
    return func() {
        println("do some teardown stuff for", task)
    }
}

func main() {
    teardown := setup("demo")
    defer teardown()
    println("do some bussiness stuff")
}
```


3. 作为参数传入函数；

```go
type myFunc func(int, int) int //函数类型

func times(a, b int, fn myFunc) int	 {
    fmt.Println("start...")
    return fn(a, b)
}
```
设计：输入两个数字，满足某种运算需求
```go
type operater func(int, int) int

func calculate(a, b int, op operater) int {
    return op(a, b)
}

func main() {
    //乘法运算
    s := calculate(3, 5, func(a, b int) int {
        return a * b
    })

    //加法
    s1 := calculate(5, 6, func(a, b int) int {
        return a * b
    })

    fmt.Println(s, s1)
}
```
4. 拥有自己的类型，函数在go语言中是类型。

    ```go
    // $GOROOT/src/net/http/server.go
    type HandlerFunc func(ResponseWriter, *Request)
    
    // $GOROOT/src/sort/genzfunc.go
    type visitFunc func(ast.Node) ast.Visitor
    ```

    

**函数“一等公民”特性的妙用**

应用一：函数类型妙用

```go

func greeting(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome, Gopher!\n")
}                    

func main() {
    http.ListenAndServe(":8080", http.HandlerFunc(greeting))
}
```

应用二：利用闭包简化函数调用

<u>Go闭包是在函数内部创建的匿名函数，这个匿名函数可以访问创建它的函数参数与局部变量。</u>

一个应用就是柯里化，把接收多个参数的函数转换接收一个单一参数的函数。

```go
func times(x, y int) int {
	return x * y
}

func partitalTimes(x int) func(int) int {
	return func(i int) int {
		return times(x, i)
	}
}

func main() {
    twoTimes := partitalTimes(2) //func(i int) int {return times(2, i)}
	fmt.Println(twoTimes(5)) //2*5
	fmt.Println(twoTimes(6)) //2*6
}
```



### 函数：结合多返回值处理错误

Go利用多返回值处理错误，错误信息通过返回值携带(n int, err error);
error是一个接口，同一使用这个接口表示错误，定义如下：

```go
// $GOROOT/src/builtin/builtin.go
type interface error {
    Error() string
}
```
Go提供构造错误的两种方法以实现error接口：errors.New("")和fmt.Errorf("", argument):
```go
err := errors.New("your first demo error")
errWithCtx = fmt.Errorf("index %d is out of bounds", i)

//使用方式
result, err := dosomething()
if err != nil { //错误处理
    err //处理错误
}
```

使用error类型，而不是传统的整型或其他类型作为错误类型的好处：

- 统一了错误类型；
- 错误是值，可以使用==和!=逻辑比较；
- 易扩展，支持自定义错误上下文。

错误处理使用策略：

**透明错误处理策略**：根据函数/方法返回的error类型变量中携带的错误值信息做决策，选择后序代码执行路径的过程。

```go
//这是最常见的一种使用方式
err := doSomething()
if err != nil {
    // 不关心err变量底层错误值所携带的具体上下文信息
    // 执行简单错误处理逻辑并返回
    ... ...
    return err
}
```

**“哨兵”处理错误策略**：	

### 怎样让函数更简洁健壮

**健壮性“三不要”原则**

原则一：不要相信外部输入的参数；需要对输入参数进行检查，返回预设厕屋。
原则二：不要忽略任何一个错误；调用函数不一定成功，要检查函数返回的错误值。
原则三：不要假定异常不会发生；异常不是错误，错误是可预测的，也是经常发生的；但异常是少见、意料之外的。<br>
通常意义上的异常是硬件异常、操作系统异常、语言运行时异常，代码中潜在的bug导致的异常，比如以分母为0情况。

**Go语言异常处理机制-panic**
panic主要来源：一是来自**Go运行时**，二是通过调用`panic`函数主动触发。
```go
func foo() {
    println("call foo")
    bar()
    println("exit foo")
}

func bar() {
    println("call bar")
    panic("panic occurs in bar")
    zoo()
    println("exit bar")
}

func zoo() {
    println("call zoo")
}

func main() {
    println("call main")
    foo()
    println("exit main")
}

//resutl
/*
call main
call foo
call bar 
panic: panic occurs in bar
*/
```
程序在执行过程中，在bar函数中调用panic引发异常，程序从这里开始paincking，自身执行停止。因为bar没有捕捉到这个panic，这个panic继续沿着函数栈向上走，一直到main，这个程序停止运行。

`recover`函数可以捕捉panic并恢复运行，修改bar函数：
```go

func bar() {
    defer func() {
        if e := recover(); e != nil {
            fmt.Println("recover the panic:", e)
        }
    }()

    println("call bar")
    panic("panic occurs in bar")
    zoo()
    println("exit bar")
}

/*
call main
call foo
call bar
recover the panic: panic occurs in bar
exit foo
exit main
*/
```
bar函数发生异常后，执行被中断；因为recover，panic没有继续蔓延，bar()其他代码可以正常执行。

如何应对panic?
1. 评估程序对panic的忍受度，不同应用对程序异常引起的程序崩溃退出的忍受度是不一样。例如单次运行于控制台的命令行交互程序，和常住内存的http服务器，前者引发的异常仅是再次重新运行，而后者崩溃会引发整个网站停止服务。
局部影响整体，我们可以将捕捉与恢复放在goroutine起始处。

2. 提示潜在bug
```go

// $GOROOT/src/encoding/json/encode.go

func (w *reflectWithString) resolve() error {
    ... ...
    switch w.k.Kind() {
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        w.ks = strconv.FormatInt(w.k.Int(), 10)
        return nil
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        w.ks = strconv.FormatUint(w.k.Uint(), 10)
        return nil
    }
    panic("unexpected map key type")
}
```

3. 不要混淆异常与错误

**使用defer函数简化**
defer 是 Go 语言提供的一种延迟调用机制，defer 的运作离不开函数。怎么理解呢？这句话至少有以下两点含义：
- 在 Go 中，只有在函数（和方法）内部才能使用 defer；
- defer 关键字后面只能接函数（或方法），这些函数被称为 deferred 函数。defer 将它们注册到其所在 Goroutine 中，用于存放 deferred 函数的栈数据结构中，这些 deferred 函数将在执行 defer 的函数退出前，按后进先出（LIFO）的顺序被程序调度执行。

defer注意事项：
1. 内置函数中， close、copy、delete、print、recover可以被设置为deferred函数。
2. **defer后面的表达式，是将deferred函数注册到deferred函数栈的时候进行求值。**
```go
func foo1() {
    for i := 0; i <= 3; i++ {
        defer fmt.Println(i)
    }
}

func foo2() {
    for i := 0; i <= 3; i++ {
        defer func(n int) {
            fmt.Println(n)
        }(i)
    }
}

func foo3() {
    for i := 0; i <= 3; i++ {
        defer func() {
            fmt.Println(i)
        }()
    }
}

func main() {
    fmt.Println("foo1 result:")
    foo1() //deferred函数顺序：println(0), println(1),...;3 2 1 0
    fmt.Println("\nfoo2 result:")
    foo2() //func(0), func(1), ...; 3 2 1 0
    fmt.Println("\nfoo3 result:")
    foo3() //func(), func(), ...; 4 4 4 4
}
```
3. 知晓defer性能消耗，go1.17版本defer性能开销较小，之前版本的defer性能开销相对比较大。

## 方法

### 一、方法的本质

方法的形式：
```go
func (r T or *T) methodName(arguements type) type {
    //body
}
```
与函数类似，不过多出了一个接收者。方法接收器（receiver）参数、函数 / 方法参数，以及返回值变量对应的作用域范围，都是函数 / 方法体对应的显式代码块。<br>
**receiver 参数的基类型本身不能为指针类型或接口类型**
```go
type myInt *int
//错误
func (my myInt) hi() {}

type reader io.Reader //接口类型
//错误
func (r reader) hi() {}
```
Go对方法声明位置为：方法声明与receiver参数的基类声明放在同一包内。基于此结论得出：
1. **不能为原生类型（int, float, map等）添加方法**
2. **不能跨越Go包为其他包的类型声明新方法**

Go 语言中的方法的本质就是，一个以方法的 receiver 参数作为第一个参数的普通函数。

### 二、方法集合与如何选择receiver类型

receive参数r类型是指针，则可以修改结构体，普通类型不能修改。<br>
**选择recceiver参数类型第一原则**
方法对receiver参数代表的实例的修改，选择*T(指针类型，T代表类型)作为参数的类型。
```go

  type T struct {
      a int
  }
  
  func (t T) M1() {
      t.a = 10
  }
 
 func (t *T) M2() {
     t.a = 11
 }
 
 func main() {
     var t1 T
     println(t1.a) // 0
     t1.M1()
     println(t1.a) // 0
     t1.M2() //(&t1).M2()
     println(t1.a) // 11
 
     var t2 = &T{}
     println(t2.a) // 0
     t2.M1() //(*t2).M1()
     println(t2.a) // 0
     t2.M2()
     println(t2.a) // 11
 }
```
 从中我们可以知道，无论是T类型实例，还是\*T类型实例，都既可以调用recevier为T的方法，也可以调用为\*T的方法。这是因为Go编译器进行了类型转换。<br>

 **选择receiver参数类型的第二原则**
 receiver参数类型size较大的时候，选择指针类型会更好。

 **方法集合**
 ```go
 
type Interface interface {
    M1()
    M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
    var t T
    var pt *T
    var i Interface

    i = pt //*T参数类型包含的方法为*T和T类型的方法，反之不同
    i = t // cannot use t (type T) as type Interface in assignment: T does not implement Interface (M2 method has pointer receiver)
}
 ```
方法集合是用来判断一个类型是否实现了某个接口类型，或者说方法集合决定接口实现。

```go
//输出非接口类型的方法集合
func dumpMethodSet(i interface{}) {
    dynTyp := reflect.TypeOf(i)

    if dynTyp == nil {
        fmt.Printf("there is no dynamic type\n")
        return
    }

    n := dynTyp.NumMethod()
    if n == 0 {
        fmt.Printf("%s's method set is empty!\n", dynTyp)
        return
    }

    fmt.Printf("%s's method set:\n", dynTyp)
    for j := 0; j < n; j++ {
        fmt.Println("-", dynTyp.Method(j).Name)
    }
    fmt.Printf("\n")
}


type Interface interface {
    M1()
    M2()
}

type T struct{}

func (t T) M1()  {}
func (t *T) M2() {}

func main() {
    var t T
    var pt *T
    dumpMethodSet(t) //-M1
    dumpMethodSet(pt) //-M1, M2
}
```
T类型的方法集合只有M1，*T类型的方法集合包含了M1,M2。所以这就是为什么pt可以赋值为接口的含义，它实现了接口的所有方法，而t没有实现，不能赋值。

方法集合决定接口实现的含义：如果某类型T的方法集合与接口类型的的方法集合相同，或者T方法集合的接口类型的超集，那么可以说T类型实现了该接口。

**选择receicer参数类型的第三个原则**
T是否需要实现某个接口，即是否存在T类型的变量赋值给某个接口的情况。
> 我的理解就是，T需要实现某个接口，优先使用T参数类型，因为*T类型包含了T类型的方法，同样也会实现这个接口。

思考题：
```go

type T struct{}

func (T) M1()
func (T) M2()

type S T
```
S类型是否包含了T的方法：结论是不包含。因为S是一个新的类型，所以两者不相同。但是type S = T，S是T的别名，两个是同一类型所以S包含T的方法。

### 类型嵌套模拟实现“继承”

**独立的自定义类型**：这个类型所有方法都是自己显示实现的。

**非独立的自定义类型**：某种方法不是自己显示实现，通过类型嵌入(Type Embedding)“继承”其他类型的方法。

**类型嵌入**：在一个类型定义中嵌入其他类型，Go支持两种类型嵌入，分别是结构体类型嵌入和接口类型嵌入。

**接口类型嵌入**
```go
type E interface {
    M1()
    M2()
}

//接口嵌入
type I interface {
    E 
    M3()
}
```
接口类型嵌入的语义就是新接口类型（如接口I）将嵌入的接口类型（接口E）的方法集合，并入到自己的方法集合中。

Go的接口通常只含少量的方法（1-2两个方法），通过接口类型嵌入其他接口类型实现接口组合，**这也是Go语言中基于已有接口类型构建新接口类型的惯用法。**

例如go的io包中定义的一系列接口：
```go

// $GOROOT/src/io/io.go

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

//接口嵌入
type ReadWriter interface {
    Reader
    Writer
}

type ReadCloser interface {
    Reader
    Closer
}

type WriteCloser interface {
    Writer
    Closer
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```
这种类型嵌入在go1.14版本之前有约束，就是嵌入接口类型的方法集合不能有交集，同时新嵌入接口的方法集合名字，不能与新接口中的其他方法同名。之后的版本没有这个约束条件。
例如go1.14版本之前：
```go
type Interface1 interface {
    M1()
}

type Interface2 interface {
    M1()
    M2()
}

type Interface3 interface {
    Interface1
    Interface2 // Error: duplicate method M1
}

type Interface4 interface {
    Interface2
    M2() // Error: duplicate method M2
}

func main() {
}
```

**结构体类型嵌入**

**带嵌入类型的结构体定义**
```go

type T1 int
type t2 struct{
    n int
    m int
}

type I interface {
    M1()
}

type S1 struct {
    T1 //嵌入字段
    *t2 //嵌入字段
    I   //嵌入字段
    a int
    b string
}
```
这种以某个类型名、类型的指针类型名或接口类型名，直接作为结构体字段的方式就叫做结构体的类型嵌入，这些字段也被叫做嵌入字段（Embedded Field）。
**嵌入字段：它们即代表字段的名字，也代表字段的类型**。

带嵌入字段的结构体和普通结构体的不同：
```go
type MyInt int

func (n *MyInt) Add(m int) {
    *n = *n + MyInt(m)
}

type t struct {
    a int
    b int
}

type S struct {
    *MyInt
    t
    io.Reader
    s string
    n int
}

func main() {
    m := MyInt(17)
    r := strings.NewReader("hello, go")
    s := S{
        MyInt: &m,
        t: t{
            a: 1,
            b: 2,
        },
        Reader: r,
        s:      "demo",
    }

    var sl = make([]byte, len("hello, go"))
    s.Reader.Read(sl) 
    fmt.Println(string(sl)) // hello, go
    s.MyInt.Add(5) 
    fmt.Println(*(s.MyInt)) // 22
}
```
嵌入字段的可见性的嵌入类型的可见性。如果这个类型首字母大写，说明这个嵌入字段是可导出的。嵌入字段的底层类型不能是指针类型，嵌入字段在结构体中是唯一的。

**“继承实现”原理**
更改上述一部分代码：
```go
var sl = make([]byte, len("hello, go"))
s.Read(sl) //结构体s没有实现Read方法
fmt.Println(string(sl))
s.Add(5) //没有实现Add方法
fmt.Println(*(s.MyInt))
//程序正常运行，返回正确结果
```
这段代码似乎告诉我们：Read方法与Add方法是S方法集合，但是S并没有显示实现这个两个方法。

其中这两个方法来自两个嵌入字段Readr和MyInt，S“继承”它们的方法实现。实际工作机制：go先检查S是否Read方法，于是查看S嵌入字段是否定义Read方法，所以Reader字段被找出来了，s.Read调用转换为s.Reader.Read。Add方法同理。

这种看似“继承”的机制，实际上是组合的思想，更具体点是组合中的代理（delegate）模式：
![alt](https://static001.geekbang.org/resource/image/a2/fc/a236306ea461e2ca90505ca9819c94fc.jpg?wh=1980x1080)

接口类型嵌入本质上是将嵌入类型的方法集合并入新类型接口的方法集合，结构体类型对类型嵌入比较宽泛，可以是任意自定义类型或接口类型。

1. **结构体类型嵌入接口类型**
```go

type I interface {
    M1()
    M2()
}

type T struct {
    I
}

func (T) M3() {}

func main() {
    var t T
    var p *T
    dumpMethodSet(t) //M1, M2, M3
    dumpMethodSet(p) //同上
}
```
结构体的方法集合，包含嵌入的接口类型的方法集合。

不过要注意的是，嵌入接口类型方法存在交集的时候会出错。这是因为嵌入其他类型的结构体本身是一个代理，在调用其实例方法时，go先查看结构本身是否实现了这个方法，没有实现则查找嵌入字段的方法，多个嵌入字段包含相同的方法，编译器无法确定选择哪一个。

解决方法是：结构体自己实现这些方法或者消除两个接口方法集合存在交集的情况。

结构体类型嵌套的妙用：**简化单元测试的编写**

由于嵌入接口类型的结构体包含这个接口类型的方法集合，这意味着，这个结构体类型也是嵌入接口类型的一个实现。
```go
package employee
  
type Result struct {
    Count int
}

func (r Result) Int() int { return r.Count }

type Rows []struct{}

type Stmt interface {
    Close() error
    NumInput() int
    Exec(stmt string, args ...string) (Result, error)
    Query(args []string) (Rows, error)
}

// 返回男性员工总数
func MaleCount(s Stmt) (int, error) {
    result, err := s.Exec("select count(*) from employee_tab where gender=?", "1")
    if err != nil {
        return 0, err
    }

    return result.Int(), nil
}
```
由于MaleCount函数只是用Stmt接口的一个方法，为了测试要Result全部实现这个接口，开销过大。如何快速建立伪对象：
```go
package employee
  
import "testing"

type fakeStmtForMaleCount struct {
    Stmt
}

func (fakeStmtForMaleCount) Exec(stmt string, args ...string) (Result, error) {
    return Result{Count: 5}, nil
}

func TestEmployeeMaleCount(t *testing.T) {
    f := fakeStmtForMaleCount{}
    c, _ := MaleCount(f)
    if c != 5 {
        t.Errorf("want: %d, actual: %d", 5, c)
        return
    }
}
```
我建立了一个为为对象，嵌入接口类型相当于实现了这个接口，只要自己实现需要测试的方法即可。

**结构体类型中嵌入结构体类型**
```go

type T1 struct{}

func (T1) T1M1()   { println("T1's M1") }
func (*T1) PT1M2() { println("PT1's M2") }

type T2 struct{}

func (T2) T2M1()   { println("T2's M1") }
func (*T2) PT2M2() { println("PT2's M2") }

type T struct {
    T1
    *T2
}

func main() {
    t := T{
        T1: T1{},
        T2: &T2{},
    }

    dumpMethodSet(t)
    dumpMethodSet(&t)
}
```
- T1包含方法：M1()
- *T1包含方法：M1(), PT1M2()
- T2类型同理
上述代码输出结果为：
```shell
main.T's method set:
- PT2M2
- T1M1
- T2M1

*main.T's method set:
- PT1M2
- PT2M2
- T1M1
- T2M1
```
我们可以知道：
- T包含方法：T1包含的方法+*T2包含的方法
- *T方法包含：\*T1包含的方法+\*T2包含的方法

**defined类型与alias类型的方法集合**
Go 语言中，凡通过类型声明语法声明的类型都被称为 defined 类型。比如:
```go
type person struct {
    name string 
    age int
}

type Myint int
```
新定义的类型与原defined类型是不同类型。对基于接口类型创建的defined的接口类型，它们的方法集合与原接口类型的方法集合是一致的。对基于非接口类型创建的defined类型，我们看一下例子：
```go
package main

type T struct{}

func (T) M1()  {}
func (*T) M2() {}

type T1 T

func main() {
  var t T
  var pt *T
  var t1 T1
  var pt1 *T1

  dumpMethodSet(t) //M1
  dumpMethodSet(t1) //empty

  dumpMethodSet(pt) //M1, M2
  dumpMethodSet(pt1) //empty
}
```
从输出结果中可知道，这两个类型没有“继承”的关系，从逻辑上讲，两者是不同的类型的语义。

对于alias定义的类型:type alias = T；这是一种类型别名，本质上是相同的类型，所以可以使用原来的方法。

## 实践：跟踪函数调用链，理解代码更直观

用defer跟踪函数的执行过程：
```go
package main

import "fmt"

func Trace(name string) func() {
    fmt.Println("enter:", name)
    return func() {
        fmt.Println("exit:", name)
    }
}

func foo() {
    defer Trace("foo")()
    bar()
}

func bar() {
    defer Trace("bar")()
}

func main() {
    defer Trace("main")()
    foo()
}
```
程序输出结果：
```shell
enter: main
enter: foo
enter: bar
exit: bar
exit: foo
exit: main
```
函数调用链的不足：
- 调用Trace时需要手动显示传入跟踪的函数名；
- 如果是并发，不同Goroutine中的函数跟踪链混在一起无法分辨；
- 输出的跟踪结果缺少层次感，调用关系不易识别；
- 对要跟踪函数，需要手动调用Trace函数

现在逐一解决问题：

**自动获取跟踪函数名**
Treace函数进阶自动获取函数调用者的信息：
```go

// trace1/trace.go

func Trace() func() {
    pc, _, _, ok := runtime.Caller(1)
    if !ok {
        panic("not found caller")
    }

    fn := runtime.FuncForPC(pc)
    name := fn.Name()

    println("enter:", name)
    return func() { println("exit:", name) }
}

func foo() {
    defer Trace()()
    bar()
}

func bar() {
    defer Trace()()
}

func main() {
    defer Trace()()
    foo()
}
```
runtime.caller获取当前Gorotine的函数调用栈，参数标识的是哪一个栈帧的信息。当参数为0时，返回的是Caller函数的调用者的函数信息；为1返回Trace函数调用者信息。caller函数的四个返回值分别代表：程序计数器(pc)，第二和第三个代表对应源文件名和所在行数；最后一个参数调用是否成功获取这些信息。

通过 runtime.FuncForPC 函数和程序计数器（PC）得到被跟踪函数的函数名称。

新版trace实现：
```go
// trace1/trace.go

func Trace() func() {
    pc, _, _, ok := runtime.Caller(1)
    if !ok {
        panic("not found caller")
    }

    fn := runtime.FuncForPC(pc) //获取函数
    name := fn.Name()

    println("enter:", name)
    return func() { println("exit:", name) }
}

func foo() {
    defer Trace()()
    bar()
}

func bar() {
    defer Trace()()
}

func main() {
    defer Trace()()
    foo()
}
```
结果：
```shell
enter: main.main
enter: main.foo
enter: main.bar
exit: main.bar
exit: main.foo
exit: main.main
```
程序结果返回了函数名及函数对应的包名。

**增加Goroutine标识**

继续对Trace函数进行改造，让它支持多Goroutine函数调用链的跟踪。解决方案：**在输出的函数出入信息时，带上一个在程序每次执行时能唯一区分Goroutine的Goroutine ID**。
