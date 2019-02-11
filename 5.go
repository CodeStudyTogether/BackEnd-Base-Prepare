Go从语言层面就支持了并行。

goroutine是Go并行设计的核心。goroutine说到底其实就是协程，但是它比线程更小，十几个goroutine可能体现在底层就是五六个线程，Go语言内部帮你实现了这些goroutine之间的内存共享。执行goroutine只需极少的栈内存(大概是4~5KB)，当然会根据相应的数据伸缩。也正因为如此，可同时运行成千上万个并发任务。goroutine比thread更易用、更高效、更轻便。

channel接收和发送数据都是阻塞的，除非另一端已经准备好，这样就使得Goroutines同步变的更加的简单，而不需要显式的lock。

var和const参考2.2Go语言基础里面的变量和常量申明
package和import已经有过短暂的接触
func 用于定义函数和方法
return 用于从函数返回
defer 用于类似析构函数
go 用于并发
select 用于选择不同类型的通讯
interface 用于定义接口，参考2.6小节
struct 用于定义抽象数据类型，参考2.5小节
break、case、continue、for、fallthrough、else、if、switch、goto、default这些参考2.3流程介绍里面
chan用于channel通讯
type用于声明自定义类型
map用于声明map类型数据
range用于读取slice、map、channel数据

