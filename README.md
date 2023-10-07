开发中... 屎一样的文档,有空再改
## 介绍
cqhttp的消息处理服务端,负责解析接收到的qq消息,并处理
目前实现消息交互,正在拓展具体的消息处理

## 功能
* 封装go-cqhttp消息处理框架,根据消息内容与发送类型进行处理并返回
* 爬取pixiv排行版,wallhaven图片,发送固定消息后随机返回
* 支持通过配置文件修改白名单(私聊、群聊)
* todo 图片更新定时任务｜通过消息更新图片

## 环境
1. 配置并启用[go-cqhttp](https://docs.go-cqhttp.org/)
2. 首次运行选择 `3` 反向Websocket通信
3. 修改运行后生成的config.yml
4. uin: qq号
5. ws-reverse中的 universal: 改为 universal: ws://127.0.0.1:9999/ws (端口号9999)
6. 然后再次启动,最后启动cqhttp-server

## 使用
* 下载源码使用go run启动/使用Goland启动
* 编译文件后直接启动
  1. 安装go环境,cmake环境
  2. 进入build文件夹`make build linux`或者`make build macos`编译
  3. 将build文件夹下的`cqhttp-server`与根目录下的`config.yml`放到同目录中
  4. 直接运行二进制文件`./cqhttp-server`

如果无法爬取pivix图片,需要本地挂梯子并且代理端口 7890(clashX默认代理端口)

## 技术栈
* 后端: [gin](github.com/gin-gonic/gin)｜[websocket](github.com/gorilla/websocket)
* 爬虫: [goquery](github.com/PuerkitoBio/goquery)

## 程序结构预想
主要划分模块【websocket】【消息处理】，需要对功能进行分层重构
