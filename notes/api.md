# API 



## 0 API协议

### 0.0 官方文档

https://www.elastic.co/guide/en/elasticsearch/reference/6.5/api-conventions.html

### 0.1 通过HTTP协议发送JSON结构的命令

一般基于Restful api

ES服务器使用的 Async IO asynchronous 是能够很好的解决 C10K问题，单机10000个链接。

### 0.2 设置文件

http 模块允许通过http 公开Es的API

```
http.compression 如果可以accept-Encoding 可以进行压缩
http.compression_level 确认http 回复 ，范围在1-9
http.cors.enabled 启用或者关闭跨源资源共享，就是从另外一个源上的浏览器请求
http.cors.allow-origin 那些源可以被允许
http.cors.allow-method 允许的metthod GET HEAD POST PUT　DELETE
http.pipelining　支持ｈｔｔｐ管线化
http.pipelining.max_events　最大支持的管线化事务
```

内部比如ｊａｖａ客户端和服务器，用ｊｓｏｎ内部通信。