开发中...
## 介绍
cqhttp的消息处理服务端,负责解析接收到的qq消息,并处理
目前实现消息交互,正在拓展具体的消息处理

## 功能
* 消息处理框架,接收go-cqhttp的消息
* 随机获取pixiv排行版图片

## 使用
1. 配置并启用go-cqhttp
2. [下载](https://docs.go-cqhttp.org/)
3. 首次运行选择 `3` 反向Websocket通信
4. 修改运行后生成的config.yml
5. uin: qq号
6. ws-reverse中的 universal: 改为 universal: ws://127.0.0.1:9999/ws (端口号9999)
7. 然后再次启动,最后启动cqhttp-server

如果无法爬取pivix图片,需要本地挂梯子并且代理端口 7890(clashX默认代理端口)