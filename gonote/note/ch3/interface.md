# 接口

## 接口：接口即契约

**认识接口类型**

接口类型是由`type`和`interface`关键字定义的一组方法集合，其中方法集合唯一确定了这个接口类型所表示的接口：
```go
type MyInterface interface {
    M1(int) error
    M2(io.Writer ...string) 
}
```
方法和函数签名一样，方法的参数列表中形参名字与返回值列表中的具名返回值，都不作为区分两个方法的依据。

Go1.14版本以后，Go接口类型允许嵌入不同类型的接口类型存在交集，但前提是交集中的方法不仅名字相同，它的函数签名相同，否则会出错：
```go
type Interface1 interface {
    M1()
}

type Interface2 interface {
    M1(string)
    M2(int)
}

type Interace3 interface {
    Interface1
    Interace2 //编译出错
}
```

接口类型声明：
```go
var i error //接口零值为nil
```

***Go规定：如果一个类型T的方法集合是某接口类型I的方法集合等价的集合或超集，那么T实现接口I，那么T类型变量可以赋值给接口I。***

空接口类型：`interface{}`，可以表示任意类型。
```go
var i interface{} = 22 //ok
i = "ying" //ok
type T struct{}
var t T
i = t //ok
i = &t //ok
```

**类型断言(type assertion)**：通过接口类型变量“还原”它的右值类型,右值类型可以是接口类型。
```go
v, ok := i.(T)
v := i.(T)
```
i是接口类型，T是非接口类型且是想还原的类型。

- ok是`true`，则变量v的类型为T，v的值为i的右值；
- ok是`false`，变量V还是要还原的类型T，值为T类型的零值。

```go
var a int64 = 13
var i interface{} = a
v1, ok := i.(int64)  //13 true
fmt.Printf("v1=%d, the type of v1 is %T, ok=%t\n", v1, v1, ok) // v1=13, the type of v1 is int64, ok=true
v2, ok := i.(string)
fmt.Printf("v2=%s, the type of v2 is %T, ok=%t\n", v2, v2, ok) // v2=, the type of v2 is string, ok=false
v3 := i.(int64) 
fmt.Printf("v3=%d, the type of v3 is %T\n", v3, v3) // v3=13, the type of v3 is int64
v4 := i.([]int) // panic: interface conversion: interface {} is int64, not []int
fmt.Printf("the type of v4 is %T\n", v4) //[]int
```

如果断言类型T是接口类型，类型断言语义为：断言i的值实现了接口类型T.

- OK是`true`，v为i的值类型，而非接口类型T;
- ok是`false`，v类型信息为接口类型T，值为`nil`。
```go
type MyInterface interface {
    M1()
}

type T int
               
func (T) M1() {
    println("T's M1")
}              
               
func main() {  
    var t T    
    var i interface{} = t
    v1, ok := i.(MyInterface)
    if !ok {   
        panic("the value of i is not MyInterface")
    }          
    v1.M1()    
    fmt.Printf("the type of v1 is %T\n", v1) // the type of v1 is main.T
               
    i = int64(13)
    v2, ok := i.(MyInterface) //没有ok判断,触发panic
    if !ok {
        fmt.Println("error", v2) //<nil>
    }
    fmt.Printf("the type of v2 is %T\n", v2)  //the type of v2 is <nil>
}
```

**尽量定义“小接口”**
接口的背后，是通过把类型的行为抽象成**契约**，建立双方共同遵守的约定，这种契将双方的耦合度降到了最低。Go接口的特点：

- **隐式契约，无需签署，自动生效**
Go接口和它的接口实现者之间的关系是隐式，不需要关键修饰。实现者只要实现接口方法中的全部方法便遵守了契约。

- **更倾向“小契约”**
“小契约”表现在**尽量定义小接口，方法个数定义在1-3个之间。**

**接口有哪些优势**
1. **接口越小，抽象程度越高**

计算机程序本身就是对真实世界的抽象与再构建。抽象就是对同类事物去除它具体的、次要的方面，抽取它相同的、主要的方面。不同的抽象程度，会导致抽象出的概念对应事物的集合不同。抽象程度越高，对应集合空间就越大；低则越具象化，更具体接近事物真实面貌，对应集合空间越小。

2. **小接口易于实现测试**
3. **小接口表示的“契约”职责单一，易于复用组合**

**定义小接口，可以遵循的几点**

- **首先，别管接口大小，先抽象出接口**
**专注于接口是编写强大而灵活的Go代码的关键**，先针对问题领域进行深入理解，聚焦抽象并发现接口。例如下图展示：
![alt](https://static001.geekbang.org/resource/image/d1/yy/d1234a47990959faabb2f28d566dc5yy.jpg?wh=1980x1080)

初期不介意这个接口方法的数量，对问题领域的理解是循序渐进的。

**第二点，将大接口分拆成小接口**
有了接口后，接口会被用到代码的各个地方。一段时间后，分析哪些场合使用了接口的哪些方法，是否将这些场合使用的接口方法提出来放入新的小接口中：
![alt](https://static001.geekbang.org/resource/image/c9/51/c9a7e97533477e4293ba2yy1e0a56451.jpg?wh=1980x1080)

**最后，要注意接口的单一契约职责**
小接口是否需要进一步拆分成只有一个方法，这个没有标准答案。可以考量现有接口是否需要满足单一契约职责。如果需要进一步拆分，提升抽象程度。


## 接口：为什么nil接口不等于nil

**接口是Go这门静态语言中唯一“动静兼备”的语法特性**。

**接口的静态特性与动态特性**

**静态特性**体现在接口变量具有静态类型，那意味着在编译器的编译阶段对所有接口类型变量赋值进行类型检查。
```go
var err error = 1 //错误
```
接口的**动态类型**，就体现在接口类型变量在运行时还存储了右值的真实类型信息。这个右值真实类型被称为接口类型的动态类型。
```go
var err error
err = errors.New("error1")
fmt.Printf("%T\n", err) //*errors.errorString
```

接口类型变量在程序运行时被赋值为不同的动态类型变量，每次赋值接口类型存储的动态类型信息都会变化。这让Go语言拥有动态语言那样使用Duck Type(鸭子类型)的灵活性。鸭子类型是指某类型表现出的特性（比如是否可以作为接口类型的右值，即这个类型实现了这个接口），不是由其基因（如c++的父类）决定的，而是由类型表现出来的行为（比如类型拥有的方法）决定的。
```go
import "fmt"

//!Duck type example

//*Duack method
type QuackAble interface {
	Quack()
}

type Duck struct {
	name string
}

func (d Duck) Quack() {
	fmt.Printf("Duck:%s is quack\n", d.name)
}

type Dog struct {
	name string
	age  uint
}

func (d Dog) Quack() {
	fmt.Printf("Dog:%s is quack and age is %d\n", d.name, d.age)
}

type Bird struct {
	name string
	high int16
}

func (b Bird) Quack() {
	fmt.Printf("Bird:%s is fly and the max high is %d\n", b.name, b.high)
}

func QuackInForest(an QuackAble) {
	an.Quack()
}

func main() {
	animal := []QuackAble{
		Duck{"Tom"},
		Dog{"Wei", 3},
		Bird{"ying", 8844},
	}

	for _, an := range animal {
		QuackInForest(an)
	}
}
```
使用接口类型QuackAble来代表Quack这一特征，三个类型都具有这个特征(实现了这个接口中的所有方法)，可以赋值给函数`QuackInForest`的参数接口变量an。三个类型都是鸭子类型，它们之间没有必然联系，可以赋值给接口变量是因为它们表现出QuackAble所要的特征。

与动态语言不同，Go接口保证了“动态特征”的使用安全性，比如如果int类型赋值给接口类型QuackAble，则编译器就能捕捉到这个错误。

**nil error 值 != nil**

我的信念是：希望自己可以变得更好，变得更强大，更加自信，有自己的能力；不再是过去的胆小鬼，我希望通过自己的改变，让自己更加自律。


## 接口：Go中最强大的魔法

**一切皆组合**

**正交**从几何看，两条直线以相交，那么两条线是正交的。从向量术语说，两个直线相互不依赖，沿着某一条直线移动，你投影到直线的位置不变。

在计算机技术中的正交性表示某种不依赖或解耦性。如果两个或多个事物中的一个发生变化，不会影响其他事物，那么这些事物就是正交性。

Go提供正交性语法元素：

- 无类型体系，没有父子概念；
- 方法与类型正交，每个类型可以有自己的方法集合，方法本质上是一个将receiver作为第一个参数的函数；
- 接口与它的实现者无“显示关联”，接口作为Go语言提供的具有天然正交的语法元素。

Go静态骨架结构两种组合方式：垂直组合和水平组合。
![alt](https://static001.geekbang.org/resource/image/2f/44/2f07b12yyea031bdc38fc3bbc316dc44.jpg?wh=1980x1080)

**垂直组合**应用在类型嵌入，通过组合的方式构建新的类型，主要表现这个几方面：
1. 接口内嵌入接口类型
2. 嵌入接口构建结构体类型
3. 嵌入结构体类型构建新类型

**水平组合**是为了连接各个水平组合，接口是连接的关键。
例如：
```go
func Save(f *os.File, data []byte) error
```
`*os.File`是来表述数据写入地址，功能实现看似很好。但是测试中，我们需要打开或创建真实磁盘文件才能获得结构体的实例。测试过程，函数写入文件后，还需要再次操作文件，读取文件写入的内容判断写入是否正确等；其次违反了ISP原则(接口分离原则)，`os.File`包含了写入操作不相关的方法；最后与`os.File`的依赖失去了扩展性。

新版的Save函数：
```go
func Save(w io.Writer, data []byte) error{
    _, err := w.Write(data)
    return err
}
```
使用接口作为参数，实现了ISP原则，io.Writer接口只有一个Write方法，这是Save函数需要的。不仅可以实现本地存储写入，而且也可以网络存储写入，让函数扩展得到提升。一个简单的例子：
```go
package main

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

func Save(w io.Writer, data []byte) error {
	_, err := w.Write(data)
	return err
}

func main() {
	b := make([]byte, 0, 128)

	buf := bytes.NewBuffer(b) //*bytes.Buffer类型，实现io.Writer接口
	data := []byte("Hello ying, I'll become the better people!")

	err := Save(buf, data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	saved := buf.Bytes()
	if !reflect.DeepEqual(saved, data) {
		fmt.Printf("want:%s \nactual:%s\n", string(data), string(saved))
	} else {
		fmt.Println("save!!!")
	}
}
```
接口承担了应用骨架的“关节”角色，通过接口进行水平组合的方式是：使用接口类型作为参数的函数或方法。

**接口的应用模式**

**基本模式**
接受接口类型作为参数的函数或方法。
```go
func YourFuncName(param YourInterfaceType)
```
![alt](https://static001.geekbang.org/resource/image/7b/d6/7be0727d1eda688dbyye480fd8d869d6.jpg?wh=1980x1080)

函数或方法中的接口作为“关节”，支持包中的多个类型与函数或方法连接到一起，实现某一特性。

**创建模式**

"接受接口，返回结构体"，这是一种将接口作为“关节”的应用模式。常见方式是以NewXXX开头的函数，返回一个对象。标准库的一段代码：
```go
func New(out io.Writer, prefix string, flag int) *Logger { 
    return &Logger{out: out, prefix: prefix, flag: flag}
}
```

**装饰器模式**

在基本模式的基础上，当返回值类型的参数类型相同时，可以得到下面函数类型：
```go
func YourWrapperFunc(param YourInterfaceType) YourInterfaceType
```
通过这个函数，可以对输入参数类型的包装，不改变被包装类型（输入参数类型）的情况下，返回具备新功能的类型。

```go
func main() {
	r := strings.NewReader("Hello ying!")
	lr := io.LimitReader(r, 5) //装饰函数，限制读取内容
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

//装饰器的设计
// $GOROOT/src/io/io.go
func LimitReader(r Reader, n int64) Reader { return &LimitedReader{r, n} }

type LimitedReader struct {
    R Reader // underlying reader
    N int64  // max bytes remaining
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
    // ... ...
}
```
由于包装器模式下的包装函数（如上面的 LimitReader）的返回值类型与参数类型相同，因此我们可以将多个接受同一接口类型参数的包装函数组合成一条链来调用，形式是这样的：
```go
YourWrapperFunc1(YourWrapperFunc2(YourWrapperFunc3(...)))
```
例如，函数截取字符读取内容并将结果转化为大写：
```go
import (
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

type capitalzedReader struct {
	r io.Reader
}

//*实现大写转换
func (r *capitalzedReader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)

	if err != nil {
		return 0, err
	}

	q := bytes.ToUpper(p) //?大写
	copy(p, q)
	return n, err
}

//*装饰函数
func CapReader(r io.Reader) io.Reader {
	return &capitalzedReader{r: r}
}

func main() {
	r := strings.NewReader("Hello ying!")
	lr := CapReader(io.LimitReader(r, 5)) //*先截取，再转换读取后的字符
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}
```

这只是针对接口的实现，也可以针对其他类型包装，例如对函数的封装：
```go
type Func func(string)

func Hello(name string) {
	fmt.Printf("Hello, %s\n", name)
}

//*装饰器
func MyDecorator(fn Func) Func {
	return func(s string) {
		fmt.Println("好久不见了，我想你!")
		fn(s)
	}
}

func main() {
	s := MyDecorator(Hello)
	s("ying")
}
```

**适配器模式**

适配器模式的核心是适配器函数**类型**(Adapter Function Type)。适配器函数**类型**是一个辅助水平组合实现的“工具”**类型**。它可以将一个满足特定函数签名的普通函数，显式转换成自身类型实例，转换后同时也是实现了某个接口。最典型的例子是http包的HandlerFunc：
```go

func greetings(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome!")
}

//HandlerFunc类型实现了Handler接口
func main() {
    http.ListenAndServe(":8080", http.HandlerFunc(greetings))
}

//Handler源代码
// $GOROOT/src/net/http/server.go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}

type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```
将普通函数转换成了实现Handler接口的类型，经过转换后我们将它作为函数参数，从而实现基于接口的组合。

**中间件(Middeware)**

中间间是装饰器模式和适配器模式结合的产物。
```go
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func validateAuth(s string) error {
	if s != "123456" {
		return fmt.Errorf("%s", "bad auth token")
	}

	return nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	s := time.Now()
	t := fmt.Sprint(s)
	fmt.Fprintf(w, "Hello, ying!!! "+t)
}

func logHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		h.ServeHTTP(w, r)
	})
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := validateAuth(r.Header.Get("auth"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {
	http.ListenAndServe(":9090", logHandler(authHandler(http.HandlerFunc(greetings))))
}

//使用curl 工具测试结果：
/*
$curl http://localhost:9090
bad auth param

$curl -H "auth:123456" localhost:9090/ 
Hello, ying!!! 2022-04-19 23:50:30.3681058 +0800 CST m=+1484.087173901
*/
```

**尽量使用空接口作为函数参数类型**，一旦使用空接口作为函数参数类型，你将失去编译器为你提供的类型安全保护屏障。