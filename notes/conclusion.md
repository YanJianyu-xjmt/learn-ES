# 总结

## 1 添加数据

- 每个doc送到doc数据流， 都要添加一个 timestamp 字段

### 1.1 例子

- 发送以下字段请求到log， log-my_app_default这个数据流，因为不存在，会自动用内置logs- * - *索引模板    

```
POST logs-my_app-default/_doc
{
  "@timestamp": "2099-05-06T16:21:15.000Z",
  "event": {
    "original": "192.0.2.42 - - [06/May/2099:16:21:15 +0000] \"GET /images/bg.jpg HTTP/1.0\" 200 24736"
  }
}
```

- 
