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
http.port 
http.publish_port 
```

