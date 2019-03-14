Go是一门类似C的编译型语言，但是它的编译速度非常快。

break    default      func    interface    select
case     defer        go      map          struct
chan     else         goto    package      switch
const    fallthrough  if      range        type
continue for          import  return       var


package main

import "fmt"

func main() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
}

main.main()函数(这个函数位于主包）是每一个独立的可运行程序的入口点。Go使用UTF-8字符串和标识符(因为UTF-8的发明者也就是Go的发明者之一)，所以它天生支持多语言。

函数可以返回函数类型
defer定义延迟调用，无论函数是否出错都确保结束前被调用
只有所有成员都支持==操作时，结构才支持相等操作
