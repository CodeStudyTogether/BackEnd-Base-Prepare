Socket有两种：TCP Socket和UDP Socket

WebSocket是HTML5的重要特性，它实现了基于浏览器的远程socket，它使浏览器和服务器可以进行全双工通信，许多浏览器（Firefox、Google Chrome和Safari）都已对此做了支持。

Go语言标准包里面没有提供对WebSocket的支持，但是在由官方维护的go.net子包中有对这个的支持

RPC就是想实现函数调用模式的网络化。客户端就像调用本地函数一样，然后客户端把这些参数打包之后通过网络传递到服务端，服务端解包到处理过程中执行，然后执行的结果反馈给客户端。
RPC（Remote Procedure Call Protocol）——远程过程调用协议，是一种通过网络从远程计算机程序上请求服务，而不需要了解底层网络技术的协议。它假定某些传输协议的存在，如TCP或UDP，以便为通信程序之间携带信息数据。通过它可以使函数调用模式网络化。在OSI网络通信模型中，RPC跨越了传输层和应用层。RPC使得开发包括网络分布式多程序在内的应用程序更加容易。

Go标准包中已经提供了对RPC的支持，而且支持三个级别的RPC：TCP、HTTP、JSONRPC

HTTP是一种协议,RPC可以通过HTTP来实现,也可以通过Socket自己实现一套协议来实现
如果是一个大型的网站，内部子系统较多、接口非常多的情况下，RPC框架的好处就显示出来了

通过GDB我们可以单步调试、可以查看变量、修改变量、打印执行过程等

go基础知识：
func struct是关键字

通过指针变量 p 访问其成员变量 name
p.name
或者
(*p).name

协程和线程都可以实现程序的并发执行
通过channel来进行协程间的通信

一个包中，可以包含多个init函数
程序编译时，先执行导入包的init函数，再执行本包内的init函数

关于局部变量的初始化
var i int = 10 或
var i = 10 或
i := 10

defer关键字：
defer 调用的函数参数的值在 defer 定义时就确定了, 而 defer 函数内部所使用的变量的值需要在这个函数运行时才确定
如果有多个defer 调用, 则调用的顺序是先进后出的顺序, 类似于入栈出栈一样
defer 函数调用的执行时机是外层函数设置返回值之后, 并且在即将返回之前

golang中的引用类型包括
A. 数组切片
B. map
C. channel
D. interface

关于main函数（可执行程序的执行起点）
A. main函数不能带参数
B. main函数不能定义返回值
C. main函数所在的包必须为main包
D. main函数中可以使用flag包来获取和解析命令行参数

panic 是用来表示非常严重的不可恢复的错误的。在Go语言中这是一个内置函数，接收一个interface{}类型的值（也就是任何值了）作为参数。panic的作用就像我们平常接触的异常。不过Go可没有try…catch，所以，panic一般会导致程序挂掉（除非recover）。所以，Go语言中的异常，那真的是异常了。你可以试试，调用panic看看，程序立马挂掉，然后Go运行时会打印出调用栈。
但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。
panic就是这么简单。抛出个真正意义上的异常。
