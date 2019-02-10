$GOPATH 目录约定有三个子目录：

src 存放源代码（比如：.go .c .h .s等）
pkg 编译后生成的文件（比如：.a）
bin 编译后生成的可执行文件（为了方便，可以把此目录加入到 $PATH 变量中，如果有多个gopath，那么使用${GOPATH//://bin:}/bin添加所有的bin目录）

go语言有一个获取远程包的工具就是go get，目前go get支持多数开源社区(例如：github、googlecode、bitbucket、Launchpad)
