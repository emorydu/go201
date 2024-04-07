# go201

⭐️ 偏好组合，正交解耦

1. 无类型体系，类型之间独立，没有子类型概念
2. 每个类型都可以具有自己的方法集合，类型定义与方法实现是正交独立的
3. 接口与其实现之间隐式关联
4. 包之间相对独立，没有子包概念
5. 垂直组合（类型嵌入）
```go
type poolLocal struct {
    private interface{}
    shared []interface{}
    Mutex
    pad [128]byte
}
...

type ReadWriter interface {
    Reader
    Writer
}
...
```
6. 水平组合（接口参数）
```go 
func ReadAll(r io.Reader) ([]byte, error)

func Copy(dst Writer, src Reader) (written int64, err error)
```

⭐️ 原生并发，轻量高效

goroutine scheduler -> goroutine(2KB) -> CPU

1. 并发不是并行
2. 并发有关结构，是一种将一个程序分解为多个小片段并且每个小片段可独立执行的程序设计方法（片段之间通过通信进行相互协作）
3. 并行有关执行，表示同时进行一些计算任务
4. 执行单元 goroutine
5. go+函数创建，函数退出goroutine退出
6. 使用channel进行并发goroutine的通信，通过select进行多路channel的并发控制

`Rob Pick`

7. 并发是大的组合概念，在程序设计层面对程序进行拆解组合，再映射到程序执行层面：goroutine各自执行特定工作，通过channel+select将goroutine组合连接起来
```go
package main

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)
	go Generate(ch)
	for i := 0; i < 10; i++ {
		prime := <-ch
		print(prime, "\n")
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
```

⭐️ Go项目目录结构（参考Kubernetes）（官方库）

⭐️ 使用gofmt、goimports进行微重构

⭐️ 包名简单一致

1. 给包命名时，考虑包自身名字，兼顾该包导出的标识符（变量、常量、类型、函数等）命名
```go
strings.Reader [good]
strings.StringReader [bad]
strings.NewReader [good]
strings.NewStringReader [bad]

bytes.Buffer [good]
bytes.ByteBuffer [bad]
bytes.NewBuffer [good]
bytes.NewByteBuffer [bad]
```

2. 保持变量声明与使用之间的距离越近越好，或者在第一次使用变量之前声明该变量

3. 常量多单词组合方式、系统错误码、系统信号名称全大写

⭐️ 接口

1. 接口类型优先以单个单词命名，拥有唯一方法或通过多个拥有唯一方法的组合接口 `FuncName + er`
```go
type Writer interface {
    Write(p []byte) (n int, err error)
}

type Reader interface {
    Read(p []byte) (n int, err error)
}

type Closer interface {
    Close() error
}

type ReadWriteCloser interface {
    Reader
    Writer
    Closer
}
```

2. 尽量定义小接口，通过接口组合方式构建程序

⭐️ 使用一致的变量声明方式

1. 使用变量之前需先进行变量声明

2. 在变量声明形式的选择上应尽量保持项目范围内一致

3. 包级变量：在package级别可见的变量，如果是导出的包级变量等价于全局变量

4. 局部变量：函数或方法体内声明的变量，仅函数或方法体内可见
```go
// 常见声明
var num int32
var s string = "hello"  // 声明并初始化
var i = 13  // 类型推断
n := 17 // 局部变量可用
var (   // 块声明并初始化
    crlf = []byte("\r\n")
    colonSpace = []byte(": ")
)

// 声明并同时显式初始化
var ErrClosedPipe = errors.New("io: read/write on closed pipe")

var EOF = errors.New("EOF")
var ErrShortWrite = errors.New("short write")
// 上述为声明变量同时显式初始化包级变量的最佳实践

// 尽量保持声明一致性
// ------------------------- good
var num = int32(17)
var fnum = float32(3.14)
// ------------------------- good

// ------------------------- bad
var num int32 = 7
var fnum float32 = 3.14
// ------------------------- bad
```

5. 声明延迟初始化
```go
// Go分配 zero value
// important：保证 zero value 可用
```

6. 声明类聚与就近原则
```go
// 声明类聚
var (
    bufioReaderPool sync.Pool
    bufioWriter2kPool sync.Pool
    bufioWriter4kPoll sync.Pool
)

var copyBufPool = sync.Pool {
    New: func() interface {
        b := make([]byte, 32*1024)
        return &b
    },
}
...

var (
    aLongTimeAge = time.Unix(1, 0)
    noDeadline = time.Time{}
    noCancel = (chan struct{})(nil)
)

var threadLimit chan struct{}
...

// 就近原则
// 如果一个包级变量在包内部被多处使用，那么这个变量放在源文件头部声明
var ErrNoCookie = errors.New("http: named cookie not present")

func (r *Request) Cookie(name string) (*Cookie, error) {
    for _, c := range readCookies(r.Header, name) {
        return c, nil
    }

    return nil, ErrNoCookie
}
```

⭐️ 局部变量声明形式

1. 延迟初始化局部变量，采用var
```go
func (r *byteReplacer) Replace(s string) string {
    var buf []byte
    for i := 0; i < len(s); i++ {
        b := s[i]
        if r[b] != b {
            if buf == nil {
                if buf = []byte(s)
            }
            buf[i] = r[b]
        }
    }
    if buf == nil {
        return s
    }

    return string(buf)
}

func Bar() {
    var err error
    defer func() {
        if err != nil {
            ...
        }
    }()

    err = Bar()
    ...
}
```
2. 声明显式初始化局部变量，短变量
```go
num := 17
fnum := 3.14
s := "hello, world!"


num := int32(17)
fnum := float32(3.14)
s := []byte("hello, world")
```

3. 在分支控制时声明变量，短变量
```go
func (v *Buffers) WriteTo(w io.Writer) (n int64, err error) {
    // tips
    if wv, ok := w.(buffersWriter); ok {
        return wv.writeBuffers(v)
    }

    // tips
    for _, b := range *v {
        nb, err := w.Write(b)
        n += int64(nb)
        if err != nil {
            v.consume(n)
            return n, err
        }
    }
    v.consume(n)

    return n, nil
}
```

4. 函数内尽量类聚/就近
```go
// 函数设计 “单一职责”
func (r *Resolver) resolveAddrList(ctx context.Context, op, network, addr string, hint Addr) (addrList, error) {
    ...
    var (
        tcp *TCPAddr
        udp *UDPAddr
        io *IPAddr
        wildcard bool
    )
    ...
}
```

⭐️ 正确使用无类型常量简化代码

1. 声明常量指定类型
```go
const (
    O_RDONLY int = syscall.O_RDONLY
    O_WRONLY int = syscall.O_WRONLY
    O_RDWR int = syscall.O_RDWR
    O_APPEND int = syscall.O_APPEND
    ...
)
```

2. 无类型常量
```go
const (
    SeekStart = 0
    SeekCurrent = 1
    SeekEnd = 2
)
```

3. Go中，两个类型即便拥有相同底层类型，也仍然是不同数据类型，不可以彼此运算
```go
type CInt int

func main() {
    var n int = 5
    var m CInt = 6
    fmt.Println(n + m)  // compile error: invalid operation: n + m (mismatched types int and CInt)
    fmt.Println(n + int(m)) // ok
}


type CInt int
const n CInt = 13
const m int = n + 5 // compile error: cannot use n + 5 (type CInt) as type int in const initializer

func main() {
    var number int = 5
    fmt.Println(number + n) // compile error: invalid operation: number + n (mismatched types int and CInt)
    fmt.Println(number + int(n))    // ok
}


const (
    five = 5
    pi = 3.1415
    s = "hello, world!"
    char = 'c'
    t = false
)

type CInt int
type CFloat float32
type CString string

func main() {
    var (
        cf CInt = five
        cpi CFloat = pi
        cs CString = s
    )

    ce := float64(five + pi)    // magical
    ...
}
```

⭐️ iota

1. 使用iota定义枚举
```go
const (
    Apple, Orange = 11, 22
    Strawberry, Grape   // 11, 22
    Pear, Watermelon    // 11, 22
)


const (
    mutexLocked = 1 << iota
    mutexWoken
    mutexStarving
    mutexWaiterShift = iota
    starvationThresholdNs = 1e6
)


const (
    _ = iota
    One
    Two
    Three
    _
    Five
    ...
)
```

2. 使用有类型枚举常量保证类型安全（Kubernetes做法）
```go
type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)
```

⭐️ 尽量定义零值可用类型（最佳实践）

1. Go提供默认值（当通过`声明`或`new`为变量分配内存、或通过`复合文字字面量`或调用`make`，且不提供显式初始化）
```go
// 具有递归性
number: 0
floating: 0.0
boolean: false
strings: ""
pointer, interface, slice, channel, map, function: nil
```

⭐️ 零值可用

```go
var numbers []int   // nil
nubmers = append(numbers, 1)
nubmers = append(numbers, 2)
nubmers = append(numbers, 3)
fmt.Println(numbers)


func main() {
    var p *net.TCPAddr  // nil
    fmt.Println(p)  // <nil>
}

func (a *TCPAddr) String() string {
    if a == nil {
        return "<nil>"
    }
    ip := ipEmptyString(a.IP)
    if a.Zone != "" {
        return JoinHostPort(ip+"%"+a.Zone, itoa(a.Port))
    }

    return JoinHostPort(ip, itoa(a.Port))
}


var mu sync.Mutex   // nil
mu.Lock()
mu.Unlock()


func main() {
    var b bytes.Buffer  // nil
    b.Write([]byte("Effective Go"))
    fmt.Println(b.String())
}
...
```

1. 切片类型只有在append场景下零值可用
```go
var numbers []int
numbers[0] = 1 // error
numbers = append(numbers, 1)    // ok
```

2. map类型不提供零值可用，必须初始化后才可使用
```go
var m map[string]string
m["go"] = "Go"  // error

m := make(map[string]string)
m["go"] = "Go"  // ok
```

3. 零值可用的类型尽量避免值复制
```go
var mu sync.Mutex
mcpy := mu  // error
foo(mu) // error
bar(&mu)    // ok
```

4. 践行Go理念，自定义类型尽量零值可用

⭐️ 优雅赋值

1. 逐个赋值（不推荐，特定场合下必须使用）
```go
var ceo CEO
ceo.name = "Emory.Du"
s.age = 24

var array [5]int
a[0] = 0
...

numbers := make([]int, 5, 5)
numbers[0] = 0
numbers[1] = 1
numbers[2] = 2
...

m := make(map[string]string)
m["go"] = "Go"
...
```

2. 构造器方式赋值
```go
ceo := CEO{
    Name: "Emory.Du",
    Age: 24,
}
array := [5]int{0, 1, 2, 3, 4}
numbers := []int{0, 1, 2, 3, 4}
m := map[stirng]string{"go": "Go"}
```

3. 结构体复合字面值
```go
// 1. 推荐filed:value对struct变量进行值构造（降低结构体类型使用者与结构体类型设计者之间的耦合度）最佳实践

// 2. field:value任意次序出现，但是不可重复出现，未出现赋对应类型零值，未导出字段其他包不可见

// 3. 增加&符号，获得对应类型的指针类型变量
```

4. map复合字面值
```go
var unitMap = map[string]int64 {
    "ns": int64(Nanosecond),
    "us": int64(Microsecond),
    "ms": int64(Millisecond),
    ...
}

var stateName = map[ConnState]string {
    StateNew: "new",
    StateActive: "active",
    ...
}
```

⭐️⭐️ 切片实现原理

1. 数组是`固定长度、同类型连续序列`，传递数组属于值拷贝行为，性能损耗大

2. 切片是数组的描述符，Go runtime 层面内部表示切片是一个结构体
```go
type slice struct {
    array unsafe.Pointer    // 底层数组`某`元素指针，代表切片的启始
    len int // 切片长度，当前切片中元素个数
    cap int // 切片最大容量，cap>=len，取决于底层数组的长度
}
```

3. 针对同一个底层数组的多个切片，经过修改值操作会产生相互影响

4. 切片作为函数参数传递，传递的是runtime.slice实例，性能损耗小

5. slice的动态扩容：当前底层数组容量无法满足的情况下，动态分配新的底层数组，新数组长度按照一定算法扩展。新数组建立以后，会将旧数组中的数据复制到新数组中，runtime slice指向新数组，旧数组GC

6. 可预估容量的切片避免过多的内存分配和复制代价


⭐️⭐️ map实现原理

1. map是无序key:value pair，key必须是可比较的，使用map必须初始化，map是引用类型（意味着函数内可以更改值）
```go
var statusText = map[int]string {
    StatusOK: "OK",
    StatusCreated: "Created",
    StatusAccepted: "Accepted",
    ...
}

icookies := make(map[string][]*Cookie)

http2commonLowerHeader := make(map[string]string, len(common))
```

2. [map基本操作](./ch3/sources/map_op.go)

3. [map遍历](./ch3/sources/tmap.go)

4. [map固序遍历](./ch3/sources/fixed_map.go)

5. ⭐️ runtime map
```go
m := make(map[keyType]valType, capcityhint) -> m := runtime.makemap(maptype, capacityhint, m)
v:= m["key"] -> v := runtime.mapaccess1(maptype, m, "key")
v, ok := m["key"] -> v, ok := runtime.mapaccess2(maptype, m, "key)
m["key"] = "value" -> v := runtime.mapassign(maptype, m, "key")
delete(m, "key") -> runtime.mapdelete(maptype, m, "key")
...
```

6. ⭐️ [`并发不安全map`](./ch3/sources/unsafemap.go) 不支持并发写

7. 尽量预估map的容量

⭐️⭐️ string类型

1. string类型的数据是`不变`的
```go
func main() {
    s := "hello world!"
    fmt.Println("original string:", s)

    sl := []byte(s)
    sl[0] = 'w'
    fmt.Println("slice:", string(sl))
    fmt.Println("after reslice, the original string is:", s)
}
```

2. `零值可用`
```go
var s string
fmt.Println(s)  // ""
fmt.Println(len(s)) // 0
```

3. string基本操作
```go
len(s)

s := "hello,"
s = s + " world!"
s += ":)"
fmt.Println(s)  // hello, world:)

== != >= <= > <
```

4. string比较规则（`长度、数据指针、内容`）

5. string内部表示
```go
type stringStruct struct {
    str unsafe.Pointer
    len int
}

func rawstring(size int) (s string, b []byte) {
    p := mallocgc(uintptr(size), nil, false))
    stringStructOf(&s).str = p
    stringStructOf(&s).len = size

    *(*slice)(unsafe.Pointer(&b)) = slice {p, size, zise}   // 该slice写入数据后则GC

    return
}
```

6. [`string的高效构造`](./ch3/sources/str_test.go)
```go
预初始化strings.Builder
预初始化bytes.Buffer & strings.Join
未预初始化 strings.Builder & bytes.Buffer
fmt.Sprintf
```

7. string、[]byte、[]rune双向转换

⭐️⭐️ package

package是go的基本单元，用于组织源代码

1. 在每个源文件显式列出所有依赖的包导入

2. Go包之间不能存在循环依赖，是一张有向无环图，包可以单独编译也可以并行编译

3. [包变量声明语句中的表达式求值顺序](./ch3/sources/evaluation_order_1.go)
```go
// ready for initialization：未初始化的且不含有对应初始化表达式或初始化表达式不依赖任何未初始化变量的变量
// 包级变量的初始化按照变量声明的先后顺序进行
// 包级变量的初始化过程就是按照声明顺序递归寻找下一个read for initialization
```

4. [Go规定表达式操作数中的所有函数、方法以及channel操作按照从左到右的次序进行求值](./ch3/sources/evaluation_order_4.go)

5. [赋值语句求值](./ch3/sources/evaluation_order_6.go)

6. [惰性求值](./ch3/sources/evaluation_order_7.go)

⭐️ 代码块与作用域

1. 显示代码块：使用{}包含的

2. Universe代码块（隐式代码块）：所有Go源码都在此处

3. 包代码块（隐式代码块）：每个包都有一个，放置该包的所有Go源码

4. 文件代码块（隐式代码块）：每个文件具有一个，包含着该文件中所有Go源码

5. switch、select语句中每个子句都被视为一个隐式代码块

6. 预定义标识符(make new cap len...)的作用域是Universe代码块

7. 函数外声明的常量、变量、类型或函数[`除了方法`]对应的标识符的作用域是Package代码块

8. Go源文件中导入的包名称的作用域是Flle代码块

9. receiver、parameters、return var对应的标识符作用域是函数体
```go
func Foo() {
    if a := 1; true {
        fmt.Println(a)
    }
}
=>
func Foo() {
    {
        a := 1
        if true {
            fmt.Println(a)
        }
    }
}

func Foo() {
    if a, b := 1, 2; false {
        fmt.Println(a)
    } else {
        fmt.Println(b)
    }
}
=>
func Foo() {
    {
        a, b := 1, 2
        if false {
            fmt.Println(a)
        } else {
            fmt.Println(b)
        }
    }
}

func main() {
    if a := 1; false {
    } else if b := 2; false {
    } else if c := 3; false {
    } else {
        println(a, b, c)
    }
}
=>
func main() {
    {
        a := 1
        if false {

        } else {
            { 
                b := 2
                if false {

                } else {
                    {
                        c := 3
                        if false {

                        } else {
                            println(a, b, c)
                        }
                    }
                }
            }
        }
    }
}
```
```go
for a, b := 1, 10; a < b; a++ {
    ...
}
=>
{
    a, b := 1, 10
    for ; a < b; a++ {
        ...
    }
}

var numbers = []int {1, 2, 3}
for i, n := range numbers {
    ...
}
=>
var numbers = []int {1, 2, 3}
{
    i, n := 0, 0
    for i, n = range numbers {
        ...
    }
}
```
```go
switch x, y := 1, 2; x + y {
case 3:
    a := 1
    fmt.Println("case1: a = ", a)
    fallthrough
case 10:
    a := 5
    fmt.Println("case2: a = ", a)
    fallthrough
default:
    a := 7
    fmt.Println("default case: a = ", a)
}
=>
{
    x, y := 1, 2
    switch x + y {
    case 3:
        {
            a := 1
            fmt.Println("case1: a = ", a)
        }   
    fallthrough 
    case 10:
        {
            a := 5
            fmt.Println("case2: a = ", a)
        } 
    fallthrough 
    default:
        {
            a := 7
            fmt.Println("default case: a = ", a)
        } 
    }
}
```
```go
c1 := make(chan int)
c2 := make(chan int, 1)
c2 <- 11

select {
case c1 <- 1:
    fmt.Println("SendStmt case has been chosen")
case i := <-c2:
    _ = i
    fmt.Println("RecvStmt case has been chosen")
default:
    fmt.Println("defualt case has been chosen")
}
=>
c1 := make(chan int)
c2 := make(chan int, 1)
c2 <- 11

select {
case c1 <- 1:
    {
        fmt.Println("SendStmt case has been chosen")
    }
case "when":
    {
        i := <-c2
        _ = i
        fmt.Println("RecvStmt case has been chosen")
    }
default:
    {
        fmt.Println("defualt case has been chosen")
    }
}
```

⭐️⭐️ `快乐路径`

1. 出现错误，快速返回，成功逻辑不要嵌入if-else语句中

⭐️⭐️ For-Range

1. [迭代变量重用](./ch3/sources/range1.go)

2. [迭代的是副本](./ch3/sources/range4.go)`slice, string, map, array, channel`

3. [string迭代的是rune](./ch3/sources/string_range.go)

4. [⭐️map-range](./ch3/sources/map_range.go)

5. [⭐️channel-range](./ch3/sources/channel_range.go)

6. [break](./ch3/sources/break.go)

7. [break-label](./ch3/sources/break-label.go)

⭐️⭐️ `init函数`

函数与方法是Go程序的逻辑块基本单元

1. init函数：包初始化时调用，多个init函数在初始化Go包时，按照一定次序逐一调用，每个init只执行一次，常用于包级数据初始化以及初始化状态检查

2. 绝不依赖init函数的执行次序：先被传递给Go编译器的源文件中的init函数先被执行

3. Go程序初始化顺序：import->const->var->init（深度优先查找）

4. init函数使用场景
```go
// 重置包级变量值
func init() {
    CommandLine.Usage = commandLineUsage
}


var closedchan = make(chan struct{})

func init() {
    close(closedchan)
}

// 包级变量初始化
var specialBytes [16]byte

func special(b byte) bool {
    return b < utf8.RuneSelf && specialBytes[b%16]&(1<<(b/16)) != 0
}

func init() {
    for _, b := range []byte(`\.+*?()|[]{}^$`) {
        specialBytes[b%16] | 1 << (b / 16)
    }
}

func init() {
    sort.Sort(sort.Reverse(byMaskLength(rfc6742policyTable)))
}

var (
    http2VerboseLogs bool
    http2logFrameWrites bool
    http2logFrameReads bool
    http2inTests bool
)

func init() {
    e := os.Getenv("GODEBUG")
    if strings.Contains(e, "http2debug=1") {
        http2VerboseLogs = true
    }
    if strings.Contains(e, "http2debug=2") {
        http2VerboseLogs = true
        http2logFrameWrites = true
        http2logFrameReads = true
    }
}


// init函数中注册模式
import (
    "database/sql"
    _ "github.com/lib/pq" // func init() {sql.Register("postgres", &Driver{})}
)

func main() {
    db, err := sql.Open("postgres", "protocol")
    if err != nil {
        panic(err)
    }
    ...
}
```

5. [init注册模式(工厂模式)](./ch4/sources/get_image_size.go)

6. init中检查失败一般直接panic

⭐️⭐️ 函数是一等公民

可以像对待值一样对待这种语法元素，这个语法元素就被称为一等公民

1. 创建函数
```go
// 普通创建
func newPrinter() *pp {
    p := ppFree.Get().(*pp)
    p.panicking = false
    p.erroring = false
    p.wrapErrs = false
    p.fmt.init(&p.buf)
    return p
}

// 函数内创建
func hexdumpWords(p, end uintptr, mark func(uintptr) byte) {
    p1 := func(x uintptr) {
        var buf [2 * sys.PtrSize]byte
        for i := len(buf) - 1; i >= 0; i-- {
            if x&0xF < 10 {
                buf[i] = byte(x&0xF) + '0'
            } else {
                buf[i] = byte(x&0xF) - 10 + 'a'
            }
            x >>= 4
        }
        gwrite(buf[:])
    }
    ...
}

// 作为类型
type HandlerFunc func(ResponseWriter, *Request)

type visitFunc func(ast.Node) ast.Visitor

type action func(current score) (result score, turnIsOver bool)

// 存储到变量
func vdsoParseSymbols(info *vdsoInfo, version int32) {
    ...
    apply := func(symIndex uint32, k, vdsoSysmbolKey) bool {
        ...
        return true
    }
    ...
}

// 作为函数入参
func AfterFunc(d Duration, f func()) *Timer {
    t := &Timer {
        r: runtimeTimer{
            when: when(d),
            f: goFunc,
            arg: f,
        },
    }
    startTimer(&t.r)
    return t
}

// 作为函数返回值
func makeCutsetFunc(cutset string) func(rune) bool {
    ...
    return func(r rune) bool {
        return IndexRune(cutset, r) >= 0
    }
}
```

2. 函数可以放入数组、切片、map等结构中，可以赋值给interface{}，[建立元素为函数的channel](./ch4/sources/func_channel.go)

3. [函数显式类型转换](./ch4/sources/conv_func.go)

4. 闭包：在函数内部定义的匿名函数，并且允许该匿名函数访问定义它的外部函数的作用域

5. [柯里化函数](./ch4/sources/currying.go)：接受多个参数的函数变换成接受一个单一参数的函数，并返回接受余下的参数和返回结果的新函数

6. [函子](./ch4/sources/functor.go)：是一个容器类型，该容器类型需要实现一个方法，接受一个函数类型参数，并在容器的每个元素上应用那个函数，得到一个新函子，原函子容器内部的元素值不受影响）

7. `延续传递式` 不推荐使用
```go
func Max(n int, m int) int {
    if n > m {
        return n
    } else {
        return m
    }
}

=>

func Max(n int, m int, f func(int)) {
    if n > m {
        f(n)
    } else {
        f(m)
    }
}

func main() {
    Max(5, 6, func(y int) {fmt.Printf("%d\n", y)})
}


func factorial(n int, f func(int)) {
    if n == 1 {
        f(1)
    } else {
        factorial(n-1, func(y int) { f(n * y) })
    }
}

func main() {
    factorial(5, func(y int) { fmt.Printf("%d\n", y) })
}
```

⭐️⭐️ defer

defer后只能接函数或者方法，执行方式LIFO，即使遇到panic

1. 释放资源
```go
func WriteToFile(fnmae string, data []byte, mu *sync.Mutex) error {
    mu.Lock()
    defer mu.Unlock()
    f, err := os.OpenFile(fname, os.O_RDWR, 0666)
    if err != nil {
        return err
    }
    defer f.Close()

    _, err = f.Seek(0, 2)
    if err != nil {
        return err
    }

    _, err = f.Write(data)
    if err != nil {
        return err
    }

    return f.Sync()
}
```

2. [拦截panic](./ch4/sources/in_panic.go)

3. [修改具名返回值](./ch4/sources/mod_ret.go)

4. [defer调试](./ch4/sources/ddebug.go)

5. 还原旧变量值
```go
func init() {
    oldFsinit := fsinit
    defer func() { fsinit = oldFsinit }()
    fsinit = func() {}
    Mkdir("/dev", 0555)
    ...
}
```

6. 支持defer的内置函数
```go
support: close, copy, delete, print, recover
`unsupport: append, cap, len, make, new`
```

7. [注意defer求值时机](./ch4/sources/deferv.go)

⭐️⭐️ 方法

类型的函数

1. 方法名首字母是否大写决定了该方法是不是导出方法

2. 方法定义要与类型定义放在同一个包内。故此不可以为原生类型(int map ...)添加自定义方法

3. `不能横跨Go包为其他包内的自定义类型定义方法`

4. receiver参数的基类型本身不能是指针类型或接口类型
```go
// 以下代码编译报错
type MyInt *int

func (r MyInt) String string (
    ...
    return ""
)

type MyReader io.Reader

func (r MyReader) Read(p []byte) (int, error) {
    ...
    return 0, nil
}
```

方法的本质：首参数为receiver的函数
```go
var t T
t.Get()
t.Set(1)
=>
// 方法表达式
var t T
T.Get(t)
(*T).Set(&t, 1)
```

1. `类型T只能调用T的方法集合中的方法，*T只能调用*T的方法集合中的方法`

2. 方法变量
```go
var t T
f := (*T).Set
f(&t, 3)
f := T.Get
f(t)
```

[选择正确receiver类型](./ch4/sources/receiver.go)
```go
// 遇事不决用指针（类型较大比较节省内存）
func (t T) M1() <=> M1(t T)
func (t *T) M2() <=> M2(t *T)

var t T
t.M1() // ok
t.M2() // <=> (&t).M2()

var t = &T{}
t.M1()  // <=> (*t).M1()
t.M2()  // ok
```

⭐️⭐️⭐️ 决定接口实现的重要因素
```go
// 方法集合决定接口实现

// 如果某个自定义类型T的方法集合是某个接口类型的方法集合的超集，则该类型T实现了该接口
// 且类型T的变量可以被赋值给该接口类型的变量
```

[实用工具](./ch4/sources/msutil.go)

1. `非接口类型的自定义类型T，其方法集合由所有receiver为T类型的方法组成，类型*T的方法集合则包含所有receiver为T和*T类型的方法`

⭐️⭐️⭐️ 类型嵌入

1. [接口嵌入接口](./ch4/sources/ici.go)

2. [结构体嵌入接口](./ch4/sources/sci.go) `结构体嵌入某接口的同时，也实现了该接口` [惯用法](./ch4/sources/fake_test.go)
```go
// 结构体嵌入多个接口类型且这些接口类型的方法集合存在交集

// 1. 优先选择结构体自身实现的方法

// 2. 如果结构体自身并未实现，那么查找结构体中的嵌入接口类型的方法集合中是否具有该方法，如果有，则提升为结构体的方法

// 3. 如果结构体嵌入多个接口类型且这些接口类型的方法集合存在交集，Go编译器报错，除非结构体自己实现了交集中的所有方法
```

3. [结构体嵌入结构体](./ch4/sources/scs.go)，类似于`继承`

方法集合

1. [defined类型的方法集合](./ch4/sources/defined.go)

2. `类型别名的方法集合`：与原类型拥有完全相同的方法集合，无论原类型是接口还是非接口类型


[可变长参数应用模式](./ch4/sources/args.go)

[模拟函数重载](./ch4/sources/mockoverload.go)`Go不支持函数重载!`

[模拟可选参数与默认参数的实现](./ch4/sources/mockoption.go)

⭐️⭐️⭐ [`Optional模式`](./ch4/sources/optionalmod.go)

⭐️⭐️⭐⭐⭐ 接口内部表示
```go
// runtime
// Go中每种类型都有唯一的_type信息，无论是内置原生类型还是自定义类型
// Go runtime会为程序内的全部类型建立只读的共享_type信息表，因此
// 拥有相同动态类型的同类接口类型变量的_type/tab信息是相同的，而接口类型
// 变量的data部分执行一个动态分配的内存空间，该内存空间存储的是赋值给接口
// 类型变量的动态类型变量的值。
type iface struct { // 拥有方法的接口类型变量
    /*
    type itab struct {
        // type interfacetype struct {
        //     typ _type    // 类型信息
        //     pkgpath data // 包路径名
        //     mhdr []imethod   // 接口方法集合切片
        // } 
        inter *interfacetype    // 该接口类型自身信息
        _type *_type    // 接口类型变量的动态类型信息
        hash uint32 
        _ [4]byte
        fun [1]uintptr  // 动态类型已实现的接口方法的调用地址数组
    }
    */
    tab *itab   // 接口本身的信息（类型信息、方法列表信息、动态类型所实现的方法的信息...）
    data unsafe.Pointer // 指向当前赋值给该接口类型变量的动态类型
}

type eface struct { // 没有方法的空接口类型变量 (interface{})
    /*
    type _type struct {
        size uintptr
        ptrdata uintptr
        tflag tflag
        align uint8
        fieldalign uint8
        kind uint8
        alg *typeAlg
        gcdata *byte
        str nameOff
        ptrToThis typeOff
    }
    */
    _type *_type    // 指向一个_type类型结构，该结构为该接口类型变量的动态类型信息
    data unsafe.Pointer // 指向当前赋值给该接口类型变量的动态类型
}
```

1. 接口类型变量具有静态类型，在编译阶段进行类型检查，接口类型变量兼具动态类型

2. 接口类型变量在程序运行时可以被赋值为不同的动态类型变量，从而支持运行时多态

3. 判断接口类型变量是否相同`只需判断_type/tab是否相同以及data指针所指向的内存空间所存储的数据值是否相同`

4. [`nil error值 != nil`](./ch5/sources/inter1.go)

5. [nil接口变量](./ch5/sources/nilinterface.go)

6. [空接口类型变量](./ch5/sources/emptyinterface.go)

7. [非空接口类型变量](./ch5/sources/nonemptyinterface.go)

8. [空接口类型变量与非空接口类型变量的等值比较](./ch5/sources/typevs.go)

⭐️⭐️⭐⭐⭐ 接口使用原则

1. 尽量定义小接口
```go
// 接口越小，抽象程度越高，被接纳度越高
// 易于实现和测试
// 契约职责单一，易于复用组合（尝试通过嵌入其他已有接口类型构建新接口类型）
```

2. 尽量不使用空接口作为函数参数
```go
// 空接口不提供任何信息
// 使用空接口作为函数参数会失去静态类型语言类型安全检查的保护屏障
```

3. [使用接口作为程序水平组合的连接点](./ch5/sources/horizontal.go) [中间件](./ch5/sources/md.go)
```go
// 一切都是组合
// 垂直组合（类型嵌入）-> 进而实现方法实现的复用、接口定义重用
// 水平组合（函数参数）-> 作为程序水平组合的连接点

// 包裹函数：接受接口类型参数，并返回与其参数类型相同的返回值
func LimitReader(r Reader, n int64) Reader { return &LimiterdReader{r, n} }

type LimitedReader struct {
    R Reader
    N int64
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
    ...
}
```

4. [使用接口提高代码可测性](./ch5/sources/v2/mail_test.go)


⭐️⭐️⭐⭐⭐ 并发编程

1. 并行（并行是启动多个单线程应用的实例，每个实例运行在一个核上，尽可能利用多核计算资源）

2. 并发（将应用分解为多个基本执行单元，可独立运行的模块，每个模块运行在一个单独的操作系统线程中）

3. Go原生并发，轻量高效，不是使用传统操作系统线程作为承载分解后的代码片段的基本执行单元，使用goroutine为并发程序设计提供原生支持
```go
// goroutine：由Go运行时负责调度的用户层轻量级线程
// goroutine优势
// 1. 占用资源小，每个goroutine初始栈2KB
// 2. go runtime调度而不是操作系统调度，上下文切换代价小
// 3. 语言原生支持
// 4. 内置channel作为goroutine间通信原语
```

并发是一种能够让程序由若干个代码片段独立组合而成，并且每个片段都是独立运行的能力

⭐️⭐️⭐⭐⭐ goroutine调度原理

1. GPM模型
```go
G：goroutine，存储了goroutine的执行栈信息、goroutine状态以及goroutine的任务函数等，G对象是可重用的，是跨M调度的
P：processor，逻辑上的，P的数量决定了系统内最大可并行的G的数量（CPU核数>=P的数量），P中最有用的是其拥有的各种G对象的队列、链表、一些缓存和状态
M：Machine，表示真正的执行计算资源，在绑定了有效的P后，进入一个调度循环

调度循环机制：从各种队列、P的本地运行队列中获取G，切换到G的执行栈上并执行G的函数，调用goexit做清理工作并回到M，如此反复，M并不保留G的状态
```

2. 抢占式调度（解决局部“饿死”问题`一个G中出现死循环的代码逻辑，那么G将永久占用分配给他的P和M，而位于同一个P中的其他G得不到调度`）

3. CSP并发模型