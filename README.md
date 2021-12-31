# 1.命名规范

1.1 Go是一门区分大小写的语言。

命名规则涉及变量、常量、全局函数、结构、接口、方法等的命名。 Go语言从语法层面进行了以下限定：任何需要对外暴露的名字必须以大写字母开头，不需要对外暴露的则应该以小写字母开头。

    当命名（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Analysize，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；
    命名如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 private ）

1.2 包名称

保持package的名字和目录保持一致，尽量采取有意义的包名，简短，有意义，尽量和标准库不要冲突。包名应该为小写单词，不要使用下划线或者混合大小写。

package domain package main

1.3 文件命名

尽量采取有意义的文件名，简短，有意义，应该为小写单词，使用下划线分隔各个单词。

approve_service.go

1.4 结构体命名

    采用驼峰命名法，首字母根据访问控制大写或者小写

    struct 申明和初始化格式采用多行，例如下面：

    type MainConfig struct {
        Port string `json:"port"`
        Address string `json:"address"`
    }
    config := MainConfig{"1234", "123.221.134"}

1.5 接口命名

    命名规则基本和上面的结构体类型

    单个函数的结构名以 “er” 作为后缀，例如 Reader , Writer 。

    type Reader interface {
            Read(p []byte) (n int, err error)
    }

1.6 变量命名

和结构体类似，变量名称一般遵循驼峰法，首字母根据访问控制原则大写或者小写，但遇到特有名词时，需要遵循以下规则：

    如果变量为私有，且特有名词为首个单词，则使用小写，如 appService
    若变量类型为 bool 类型，则名称应以 Has, Is, Can 或 Allow 开头

var isExist bool var hasConflict bool var canManage bool var allowGitHook bool

1.7常量命名

常量均需使用全部大写字母组成，并使用下划线分词

const APP_URL = "https://www.baidu.com"

如果是枚举类型的常量，需要先创建相应类型：

type Scheme string

const (
HTTP Scheme = "http"
HTTPS Scheme = "https"
)

2. 错误处理#

   错误处理的原则就是不能丢弃任何有返回err的调用，不要使用 _ 丢弃，必须全部处理。接收到错误，要么返回err，或者使用log记录下来 尽早return：一旦有错误发生，马上返回 尽量不要使用panic，除非你知道你在做什么
   错误描述如果是英文必须为小写，不需要标点结尾 采用独立的错误流进行处理

// 错误写法 if err != nil { // error handling } else { // normal code }

// 正确写法 if err != nil { // error handling return // or continue, etc. } // normal code

3. 单元测试#

单元测试文件名命名规范为 example_test.go 测试用例的函数名称必须以 Test 开头，例如：TestExample 每个重要的函数都要首先编写测试用例，测试用例和正规代码一起提交方便进行回归测试 。


# fmt
```text
Print系列函数会将内容输出到系统的标准输出，区别在于
Print函数直接输出内容，
Printf函数支持格式化输出字符串，
Println函数会在输出内容的结尾添加一个换行符。
Fprint系列函数会将内容输出到一个io.Writer接口类型的变量w中，我们通常用这个函数往文件中写入内容。
Sprint系列函数会把传入的数据生成并返回一个字符串。
Errorf函数根据format参数生成格式化字符串并返回一个包含该字符串的错误。

```
# 枚举

```text
go语言并没有提供enum的定义，我们可以使用const来模拟枚举类型。
```

```go

type PolicyType int32

const (
Policy_MIN      PolicyType = 0
Policy_MAX      PolicyType = 1
Policy_MID      PolicyType = 2
Policy_AVG      PolicyType = 3
)

func foo(p PolicyType) {
fmt.Printf("enum value: %v\n", p)
}

```

# 切片和数组

```text
切片没有自己的任何数据。它只是底层数组的一个引用。
对切片所做的任何修改都将反映在底层数组中。
数组是值类型，而切片是引用类型。

append(slice,3,4,5) 修改变地址值
append()会改变切片所引用的数组的内容，从而影响到引用同一数组的其他切片。当使用append()追加元素到切片时，
如果容量不够（也就是(cap-len) == 0），Go就会创建一个新的内存地址来储存元素。
```

# goroutine

```text
go+函数:开启一个单独的goroutine执行函数
多个 goroutine ,逆序执行
	 
```

```go
func rand2() {
// 随机种子,防止 随机数 每次都一样
rand.Seed(time.Now().UnixNano())
for i := 0; i < 1000; i++ {
num := rand.Intn(100)
fmt.Println("随机数 = ", num)
}
}
```

## sync.WaitGroup

## GMP 调度模型

```text
GMP
M:N,把m个goroutine分配给n个系统线程执行

goroutine: 初始栈大小2K

go并发模型是  csp,通过通信共享内存,而不是通过内存共享实现

```

## channel 通道
```text
1.一定要初始化
make(chan int)
2. 结构体 比较大,一般使用 指针

3. 含有缓冲区 make(chan int,16)
    无缓冲区 make(chan int)
```

#NSQ 分布式的队列
```text
nsqlookupd 管理多个nsqd
```

## context 
```text
控制子goroutine退出
```

# 第三方库
```text
tail: 日志读取库
sarama: 写日志
etcd: 类似 zookeeper,consul,用于配置共享和服务的注册和发现
es: 搜索引擎
Kibana: 搭配es使用的UI
prometheus: 采集监控数据,主动拉取
grafana: prometheus的第三方UI
```