对于普通的上网过程，系统其实是这样做的：浏览器本身是一个客户端，当你输入URL的时候，首先浏览器会去请求DNS服务器，通过DNS获取相应的域名对应的IP，然后通过IP地址找到IP对应的服务器后，要求建立TCP连接，等浏览器发送完HTTP Request（请求）包后，服务器接收到请求包之后才开始处理请求包，服务器调用自身服务，返回HTTP Response（响应）包；客户端收到来自服务器的响应后开始渲染这个Response包里的主体（body），等收到全部的内容随后断开与该服务器之间的TCP连接。

在浏览器中输入www.qq.com域名，操作系统会先检查自己本地的hosts文件是否有这个网址映射关系，如果有，就先调用这个IP地址映射，完成域名解析。
如果hosts里没有这个域名的映射，则查找本地DNS解析器缓存，是否有这个网址映射关系，如果有，直接返回，完成域名解析。
如果hosts与本地DNS解析器缓存都没有相应的网址映射关系，首先会找TCP/IP参数中设置的首选DNS服务器，在此我们叫它本地DNS服务器，此服务器收到查询时，如果要查询的域名，包含在本地配置区域资源中，则返回解析结果给客户机，完成域名解析，此解析具有权威性。

Golang开发api服务所遇到的一些基础问题

强类型语言
对于从弱类型语言转而来使用Golang的同学来说, 对于一些类型转换以及计算方面要多加留意

协程和channel的使用
这是golang的特点, 可以大幅度提高性能, 如果了解的不深入, 也请务必正确使用

性能分析工具
善用pprof工具来分析服务的性能. 
自带性能分析, 多么良心的语言呀, 但在线上环境请不要开启, 会影响性能.

配置以及多环境
这方面的可选方案有很多, 配置格式可以选择使用toml. 
多环境的区分可以通过命令行参数传参的方式, 来加载不同的配置文件, 比如生产环境, 测试环境, 正式环境.

golang服务的自动运维
自动重启, 当服务发布后, 会自动平滑重启
服务检测, 如果服务未能正常运行, 则尝试去启动它

golang 踩过的坑
1.闭包导致循环时取到的数据不正确
2.冒号形式初始化覆盖全局变量
3.污染了全局变量，导致不同请求间相互有影响
4.新开的协程内部有panic但是没有recover
5.普通map并发读写报错
6.map遍历顺序无法保证
7.http response body要读取完毕
8.http client 长连接参数
9.在一个值为 nil 的 channel 上发送和接收数据将永久阻塞

并发和并行：
你吃饭吃到一半，电话来了，你一直到吃完了以后才去接，这就说明你不支持并发也不支持并行。你吃饭吃到一半，电话来了，你停了下来接了电话，接完后继续吃饭，这说明你支持并发。你吃饭吃到一半，电话来了，你一边打电话一边吃饭，这说明你支持并行。并发的关键是你有处理多个任务的能力，不一定要同时。并行的关键是你有同时处理多个任务的能力。所以我认为它们最关键的点就是：是否是『同时』。
