https://github.com/dyc87112/SpringBoot-Learning/tree/2.x

Spring Boot让我们的Spring应用变的更轻量化。我们不必像以前那样繁琐的构建项目、打包应用、部署到Tomcat等应用服务器中来运行我们的业务服务。通过Spring Boot实现的服务，只需要依靠一个Java类，把它打包成jar，并通过java -jar命令就可以运行起来。这一切相较于传统Spring应用来说，已经变得非常的轻便、简单。

总结一下Spring Boot的主要优点：
为所有Spring开发者更快的入门
开箱即用，提供各种默认配置来简化项目配置
内嵌式容器简化Web项目
没有冗余代码生成和XML配置的要求

Project：使用什么构建工具，Maven还是Gradle；本教程将采用大部分Java人员都熟悉的Maven，以方便更多读者入门学习。
Language：使用什么编程语言，Java、Kotlin还是Groovy；本教程将采用Java为主编写，以方便更多读者入门学习。
Spring Boot：选用的Spring Boot版本；这里将使用当前最新的2.1.3版本。
Project Metadata：项目的元数据；其实就是Maven项目的基本元素，点开More options可以看到更多设置，根据自己组织的情况输入相关数据，比如：

若你的环境有多个版本的JDK，注意到选择Java SDK的时候请选择Java 8（具体根据你在第一步中选择的Java版本为准）

Spring几乎是每一位Java开发人员都耳熟能详的开发框架，不论您是一名初出茅庐的程序员还是经验丰富的老司机，都会对其有一定的了解或使用经验。在现代企业级应用架构中，Spring技术栈几乎成为了Java语言的代名词，那么Spring为什么能够在众多开源框架中脱颖而出，成为业内一致认可的技术解决方案呢？我们不妨从最初的Spring Framework开始，看看它为什么能够横扫千军，一统江湖！

2004年3月，Spring的第一个版本以及其创始人Rod Johnson的经典力作《Expert one-on-one J2EE Development without EJB》发布，打破了当时Java开发领域的传统思考模式，企业级应用开始走向“轻量化”发展的步伐。

最初的Spring Framework 1.0并不像如今的Spring那么复杂，但是在该版本中已经包含了Spring中最为核心的两大要素：依赖注入和面向切面编程，这两个功能是Spring区别于其他优秀框架，并在企业级应用中建立核心地位的关键所在。很多开发者在初涉Java应用的时候很可能会觉得这两个功能的意义并不大，因为不用它们我们依然可以很好的实现业务功能，事实也确实如此，但是随着业务的迭代和开发的深入，复杂多变的需求开始慢慢侵蚀原本“完美”的架构，开发与测试的难度逐步增大，往往在这个时候，我们才体会到了Spring的价值。所以，即便在Spring的最初版本中也封装了诸多偏业务型的功能封装，如：邮件发送、事务管理等，但我们要知道真正让企业级应用离不开Spring的理由并不是这些与业务直接相关的功能，而是上面所提及的与业务实现毫不相关的两大核心。

由于在初期版本中Spring对很多功能性封装并没有今天的Spring那么强大，所以很长一段时间，我们都采用了Spring做工程管理来整合其他更优秀的功能型框架来完成系统开发的架构模式，比如曾经风靡一时的Spring + Struts + Hibernate架构，相信可以勾起一代人的回忆。

2014年4月1日，Spring Boot发布了第一个正式版本。该项目旨在帮助开发者更容易地创建基于Spring的应用程序和服务，使得现有的和新的Spring开发者能够最快速地获得所需要的Spring功能。一直到今天发布2.x版本，共经历了近4年的发展，Spring Boot已经是一个拥有了21000多Star，15000多次Commits，贡献者超过400多名的超热门开源项目。

支持最新的Java 9
基于Spring 5构建，Spring的新特性均可以在Spring Boot 2.0中使用
为各种组件的响应式编程提供了自动化配置，如：Reactive Spring Data、Reactive Spring Security等
支持Spring MVC的非阻塞式替代方案WebFlux以及嵌入式Netty Server
Spring Boot 2.0的发布，Spring Cloud Finchley还会远吗？

在Spring Boot的版本计划中明确说明了2.0版本开始才对Java 9进行支持，而1.x版本暂时没有对Java 9的支持计划，所以如果我们要使用Java 9，就必须将Spring Boot版本升级至2.0。

Spring Boot 2.0 要求Java 版本必须8以上， Java 6 和 7 不再支持。

Spring Boot 2.0 发布了Kotlin runApplication扩展：
