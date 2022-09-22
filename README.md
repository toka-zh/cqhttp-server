开发中... 屎一样的文档,有空再改
## 介绍
cqhttp的消息处理服务端,负责解析接收到的qq消息,并处理
目前实现消息交互,正在拓展具体的消息处理

## 功能
* 消息处理框架,接收go-cqhttp的消息
* 随机获取pixiv排行版图片
* 随机获取wallhaven图片
* 支持通过配置文件修改白名单(私聊、群聊)
* todo 通过消息更新图片

## 环境
1. 配置并启用go-cqhttp
2. [下载](https://docs.go-cqhttp.org/)
3. 首次运行选择 `3` 反向Websocket通信
4. 修改运行后生成的config.yml
5. uin: qq号
6. ws-reverse中的 universal: 改为 universal: ws://127.0.0.1:9999/ws (端口号9999)
7. 然后再次启动,最后启动cqhttp-server

## 使用
* 下载源码使用go run启动/使用Goland启动
* 编译文件后直接启动
  1. 安装go环境,cmake环境
  2. 进入build文件夹`make build linux`或者`make build macos`编译
  3. 将build文件夹下的`cqhttp-server`与根目录下的`config.yml`放到同目录中
  4. 直接运行二进制文件`./cqhttp-server`

如果无法爬取pivix图片,需要本地挂梯子并且代理端口 7890(clashX默认代理端口)
