# Golang 第五次作业

在上一次作业中，部分同学进行了分流

- 选择做 pr 的同学：联系负责人提供新的需求
- 选择实现接口的同学：请完成下列内容

## 目的

- 掌握Http协议和Web工作原理
- 掌握 WebSocket 原理与实践
- 掌握关系型数据库的基本操作
- 完成项目结构优化
- 完成一定的自动化流程
- 完善上一个项目的文档

## 背景

众所周知，FanOne是个家喻户晓的**Aquaman**，她经常在社交软件上找小哥哥们聊天，以至于被多个平台封杀，请你写一个IM即时通信系统，让FanOne能聊天自由吧！

## 任务

请遵照以下接口文档完成功能

[https://apifox.com/apidoc/shared-dcb1a5ef-5c75-4a0e-9486-e3bd748379a0](https://apifox.com/apidoc/shared-dcb1a5ef-5c75-4a0e-9486-e3bd748379a0)

本次新增的需求

| 模块名 | 最低需要完成的接口                           | 数量 |
| ------ | -------------------------------------------- | ---- |
| 用户   | 获取 MFA qrcode、绑定 MFA               | 2    |
| 视频   | 视频流         | 1    |
| 互动   |  | 0    |
| 社交   | 聊天       | 1    |

这次作业只新增了 4 个接口，但是按照作业守恒定律，你会在其他地方付出时间

### 接口
1. 互动模块：评论接口需要实现**对评论进行评论**（即支持 comment_id 请求字段）
2. 互动模块：点赞接口需要处理对评论的点赞
3. 社交模块：完成基于 websocket 的聊天功能，考虑到聊天的实时性，请使用 Redis + MySQL 方式实现

Hertz 框架内置 WebSocket 实现，请使用 Hertz 内置的 WebSocket（[文档](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/protocol/websocket/)）

### 目录结构

目录结构一定程度上决定了其他人理解你项目的难易程度，如果你的项目具有目录结构的提升可能性，请优化你的目录结构（**请重点关注这一点**）

如果你上一个项目的目录树不是用`tree`生成的，请使用`tree`命令来生成目录树，windows 下也有相应的解决方案（不需要精确到每一个文件，只需要到目录，以及一些关键文件）

### 代码复用性

上一次作业没有复用性要求，这一次有了。
对于多次、重复出现的代码，请考虑整合、抽象、提取。

**你需要在文档中说明你这部分的修改情况**，也可以留空（比如你认为已经没什么地方可以提升复用性了，不强制要求）

### 源代码管理

首先，请修改你这个项目的仓库权限——**要求所有人不能直接推送到 main 分支**（请注意，如果你仍然在使用 master 分支，立即改为 main 分支作为主分支）

接下来，**在你这个月，以及未来的所有代码变更中，使用 Pull Request（pr）完成**，即使这个项目目前仍然只有你一个人维护

在这过程中，**请注意 pr 的规范性**，你可以自己给自己拟定一套规范，也可以参考一些开源社区的规范

尤其是需要注意 pr 的标题，**尽可能的保证可以通过标题直接知道你这个 pr 做了什么**，但 pr 的标题不宜过长

**我们会检查你的 pr 记录**（不会检查时间，放心，可以在合适的时间范围内赶 deadline）

你还需要保证你的项目具备下述文件：

1. `.gitignore`：如果你项目有一些无关数据，请使用 gitignore 忽略掉，下一次答辩会检查各位的仓库干净程度
2. `.dockerignoer`：与上一个类似，但是目的是为了减少打包过程中的无关数据
3. `.editorconfig`：EditorConfig有助于为跨各种编辑器和 IDE 处理同一项目的多个开发人员维护一致的编码风格。
4. `.gitattributes`：用于指定 Git 应该如何对待特定文件或路径中的文件

你需要自行利用搜索引擎完成这几个文件的简单学习

### 持续集成（CI）

你的项目需要引入 Github Action 工作流（[Github Action 快速入门](https://docs.github.com/zh/enterprise-cloud@latest/actions/quickstart)）

要求至少实现以下几点：

1. 漏洞扫描：CodeQL（[关于使用 Code QL 进行代码扫描](https://docs.github.com/zh/code-security/code-scanning/introduction-to-code-scanning/about-code-scanning-with-codeql)）
2. 代码规范：golangci-lint（[Introduction - golangci-lint](https://golangci-lint.run/)）

其中，golangci-lint 是一个静态代码扫描检查工具，它有本地的 cli（命令行）帮助你快速找到哪些地方的代码不合规范，规范是一个合格的软件工程师必备的技能，因此你需要

1. 较为熟悉的使用 golangci-lint
2. 在你项目的根目录中添加一个`.golangci.yml`文件，这个文件将会指定静态检查的严格程度（[使用教程](https://golangci-lint.run/usage/configuration/)）
3. 在下一次答辩中，你需要说明你的这份配置文件开启了哪些检查，**开启过少的检查是不合适的**

不要直接套用现成的配置，你需要知道workflow的配置内容，答辩时会对你简单了解一下

### 文档编写

我们鼓励你**使用飞书文档**来提升文档编写的效率

- 对于所有使用了缓存（如 Redis ）的接口，请在文档中**绘制一份流程图**（这很简单）来描述接口运作原理
- 对于所有使用了消息队列的接口，同上
- 将这份飞书文档粘贴到你项目的`README.md`上

自己编写接口的流程图，有助于我们快速了解整个流程，**同时也有助于你自己发现这个流程中的问题**（如有）

同时，请检查你的`README.md`文件，如果可能，可以做一些文档拆分（在根目录建立一个`docs`文档目录，里面存放子文档），`README.md`应该是尽可能的简单描述这个项目

最后，**请提供一个部署文档在仓库中**（可以写在 `README.md`），告诉用户你的项目应当如何部署到服务器（请注意这个需求，这隐含着一个要求：**你的项目自己部署到服务器上过**）

### 细节优化

**错误处理**

请[参照这篇文章](https://juejin.cn/post/7246777406387306553)来检查你的项目对于错误的处理是否合适，如果不合适，请修改
***
**参数校验**

Hertz 支持参数校验（[绑定与校验 | CloudWeGo](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/binding-and-validate/)），请添加这个功能
***
**流量治理**

Hertz 集成了 Sentinel（[Sentinel | CloudWeGo](https://www.cloudwego.io/zh/docs/hertz/tutorials/service-governance/sentinel/)），请集成这个功能。要求做到自定义配置（具体请参照 Hertz example 与 Sentinel 官方文档，这是阿里巴巴的项目，中文支持良好）
***
**代码生成**

如果你在上一次作业中没有使用 hz，请使用 hz 代码生成工具自动生成项目
***
**常量规范**

将常量统一进行管理，这部分可以参考github上的一些项目（请不要只对着一个项目借鉴），提升代码复用性，降低代码修改难度。常量不只是数字，还包含文本，对于一些大量重复的文本，也使用常量管理
***
**配置**

请独立一个 `config` 文件夹，并内置一个`config.yaml`（请勿使用 ini），该文件负责一些常量的配置，要求支持配置**热更新**

可以使用 Viper 库（[spf13/viper](https://github.com/spf13/viper)）

同时，为了方便我们**检查你的数据库结构**，请在 `config` 文件夹内新建一个 `sql` 文件夹，在该文件夹内储存你数据库的建表语句（如`init.sql`）

### 单元测试
单测是很重要的（这里省略 20000 个字），虽然这和性能优化无关，但是你仍然需要为你的项目添加一定的单元测试

你可以简单的进行单元测试入门（[Golang 单元测试指引 | 腾讯技术工程](https://zhuanlan.zhihu.com/p/267341653)、[Golang 单元测试合集整理](https://zhuanlan.zhihu.com/p/656105651)）

Hertz 同样也提供了单元测试能力（[单测 | CloudWeGo](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/unit-test/)）

你必须对你的项目添加一定量的单元测试（考虑大家能力、时间不同，不要求全部完成）

请在报告中提供：

1. 单元测试覆盖率（可以使用 go 自带的 `go test`命令行工具获取单元测试覆盖率）
2. 哪些部分使用了单元测试
3. 你的项目该如何进行单元测试

**你需要在报告的结尾添加你的单元测试学习笔记**.

我们很少硬性规定一定要写笔记，但是这部分请认真对待，你可以写自己对单元测试的理解、`go test` 命令行工具的了解等。

字数不限，不需要贴很多字，**不需要套话（写的笔记人能看得懂就行）**，但是请保证是自己的产出。

### 性能优化

**提示：请在完成以上所有内容后开始着手完成这部分的内容**

这部分主要是在已有基础上进行修改、更迭
***
**缓存**

有一些接口是可以利用缓存提升接口相应效率的，请自行选择你认为需要完成优化的接口，并逐个进行优化（不要求全部完成）

***
**数据库**

众所周知，在八股文中，有超级多涉及到数据库的内容。

但这里不是要你去背八股文，而是发挥你尽可能的努力，按照接口需求，对数据库**表的结构**进行优化
***
**并发**

对于一些接口（例如上传头像），实际上可以并发操作的，请对你认为可以进行优化（使用 goroutine）的地方进行优化

tips：实际上，对于一些非敏感性内容（例如点赞），**可以提前返回响应**，然后在服务端再进行点赞的落库
***
**请在报告中展示你的优化内容**（需要提供优化前和优化后，不需要说明提升了多少效率）

这部分没有硬性工作量要求，但并不是说你可以直接略过或者就没做什么事（你敢略过试试！），因为**前面的文字看似多、实际工作量并不多**

## 报告

你需要编写一份报告用于答辩（使用飞书文档），在项目提交时提交（可以先提交文档链接，后续继续优化文档），**不限制报告格式**、内容，但需要拥有以下内容

1. （Problem Restatement）问题重述：用**最简短**的话复述你这次需要完成的内容
2. （問題が解決しました）问题解决：使用打勾（复选框）来示意你**全部完成的内容**，对于部分完成的内容，请不要打勾，而是描述你目前已经完成的内容
3. （如有）（Spotlight）项目亮点：这部分不是必须的，如果你认为你的项目**有一些巧思**，请写上
4. （如有）（Advancement）进阶：超出文档需求的完成量，比如实现了部分Bonus 内容，则在这部分描述
5. （如有）（Argument）抱怨：你对这个文档存在的不足的抱怨，请尽量写，**不要害羞**，最好不要写个无


本次作业不要求全部完成，**但是会衡量你的工作量**，请酌情注意任务需求


## Bonus

1. 项目使用 Kitex（不会用不要乱上哈）
2. 请考虑你的聊天系统的性能（例如使用Benchmark测试）
3. 考虑聊天传输的安全性（可以学习一下Telegram是如何保证传输安全性的，但是现阶段是做不到的，可以尝试做一些小的安全性优化）
4. 使用消息队列（RabbitMQ、RocketMQ、Kafka等）

## 预告

在未来的作业中，你需要完成一个以图搜图的接口，这个功能的实现基于一个非关系型数据库——向量数据库（Vector Database），可以先了解[Milvus](https://milvus.io/)

**答辩时会询问你对上下文（Context）这个概念的理解**


可以了解一定的可观测性和治理特性，学有余力的可以开始学习链路追踪（Jaeger）、监控（Prometheus）等内容

## 参考

有啥好参考的？难道你现在还不会谷歌？

- [WebSocket | CloudWeGo]([github.com/gorilla/websocket](https://www.cloudwego.io/zh/docs/hertz/tutorials/basic-feature/protocol/websocket/))
- [慢聊Go之GoLang中使用Gorilla Websocket](https://juejin.cn/post/6946952376825675812)
- [RabbitMQ Go语言客户端教程2——工作队列](https://www.liwenzhou.com/posts/Go/rabbitmq-2/)