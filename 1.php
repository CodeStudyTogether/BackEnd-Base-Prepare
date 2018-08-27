1、双引号和单引号的区别
双引号解释变量，单引号不解释变量
双引号里插入单引号，其中单引号里如果有变量的话，变量解释
双引号的变量名后面必须要有一个非数字、字母、下划线的特殊字符，或者用{}讲变量括起来，否则会将变量名后面的部分当做一个整体，引起语法错误
双引号解释转义字符，单引号不解释转义字符，但是解释'\和\\
能使单引号字符尽量使用单引号，单引号的效率比双引号要高（因为双引号要先遍历一遍，判断里面有没有变量，然后再进行操作，而单引号则不需要判断）

2、常用的超全局变量(8个)
$_GET ----->get传送方式
$_POST ----->post传送方式
$_REQUEST ----->可以接收到get和post两种方式的值
$GLOBALS ----->所有的变量都放在里面
$_FILES ----->上传文件使用
$_SERVER ----->系统环境变量
$_SESSION ----->会话控制的时候会用到
$_COOKIE ----->会话控制的时候会用到

3、HTTP中POST、GET、PUT、DELETE方式的区别
HTTP定义了与服务器交互的不同的方法，最基本的是POST、GET、PUT、DELETE，与其比不可少的URL的全称是资源描述符，我们可以这样理解：url描述了一个网络上资源，而post、get、put、delegate就是对这个资源进行增、删、改、查的操作！

3.1表单中get和post提交方式的区别
get是把参数数据队列加到提交表单的action属性所指的url中，值和表单内各个字段一一对应，从url中可以看到；post是通过HTTPPOST机制，将表单内各个字段与其内容防止在HTML的head中一起传送到action属性所指的url地址，用户看不到这个过程
对于get方式，服务器端用Request.QueryString获取变量的值，对于post方式，服务器端用Request.Form获取提交的数据
get传送的数据量较小，post传送的数据量较大，一般被默认不受限制，但在理论上，IIS4中最大量为80kb，IIS5中为1000k，get安全性非常低，post安全性较高
3.2
GET请求会向数据库发索取数据的请求，从而来获取信息，该请求就像数据库的select操作一样，只是用来查询一下数据，不会修改、增加数据，不会影响资源的内容，即该请求不会产生副作用。无论进行多少次操作，结果都是一样的。

与GET不同的是，PUT请求是向服务器端发送数据的，从而改变信息，该请求就像数据库的update操作一样，用来修改数据的内容，但是不会增加数据的种类等，也就是说无论进行多少次PUT操作，其结果并没有不同。

POST请求同PUT请求类似，都是向服务器端发送数据的，但是该请求会改变数据的种类等资源，就像数据库的insert操作一样，会创建新的内容。几乎目前所有的提交操作都是用POST请求的。

DELETE请求顾名思义，就是用来删除某一个资源的，该请求就像数据库的delete操作。

4、HTTP状态码
点击这儿查看HTTP状态码详解

常见的HTTP状态码：

200 - 请求成功
301 - 资源(网页等)被永久转义到其他URL
404 - 请求的资源(网页等)不存在
505 - 内部服务器错误
HTTP状态码分类:

1** - 信息，服务器收到的请求，需要请求者继续执行操作
2** - 成功，操作被成功接收并处理
3** - 重定向，需要进一步的操作以完成请求
4** - 客户端错误，请求包含语法错误或者无法完成请求
5** 服务器错误，服务器在处理请求的过程
中发生了错误
