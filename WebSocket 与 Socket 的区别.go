WebSocket 是一种在单个 TCP 连接上进行全双工通信的协议。WebSocket 使得客户端和服务器之间的数据交换变得更加简单，允许服务端主动向客户端推送数据。在 WebSocket API 中，浏览器和服务器只需要完成一次 HTTP 握手，两者之间就直接可以创建持久性的连接，并进行双向数据传输。

一般情况下我们使用 HTTP 有一个很大的缺陷，就是 HTTP 只能由客户端来主动发起，如果有需要服务端主动通知的业务，就需要轮训。轮询的效率低，非常浪费资源。为了解决 Web 端即时通讯的需求就出现了 WebSocket。

WebSocket 最早是在 HTML5 标准中的一部分，基本现代浏览器都是支持的，都可以直接使用。虽然是 HTML5 的一部分，但是在浏览器之外也有相应的实现。比如 Socket.io 这个库支持 Java、C++、Swift、Datr ，Python 的 ws4py，C++ 的WebSocket++ 等等。
服务端主流的 Web 容器如 Apache、Nginx、Tomcat 也都支持 WebSocket，还有 WebSocket 服务端框架如 websocketd 等。

网络中的 Socket 并不是什么协议，而是为了使用 TCP，UDP 而抽象出来的一层 API，它是位于应用层和传输层之间的一个抽象层。Socket 是对 TCP/IP 的封装；HTTP 是轿车，提供了封装或者显示数据的具体形式；Socket 是发动机，提供了网络通信的能力。

Socket 是传输控制层的接口。用户可以通过 Socket 来操作底层 TCP/IP 协议族通信。
WebSocket 是一个完整应用层协议。
Socket 更灵活，WebSocket 更易用。
两者都能做即时通讯
