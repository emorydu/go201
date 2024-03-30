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
```



