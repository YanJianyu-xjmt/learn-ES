# Intro

## 0 简单概念

### 0.1 基本概念

- 专业术语
  
  - es 可以相当于一个mySQL
  
  - index(索引)->database(数据库)
  
  - Type(类型)->table(表)
  
  - Documents->row 行
  
  - Fields(字段)-> column 列
  
  - Type 已经删除了，从6开始 一个索引只包含一个type

### 0.2 倒排索引

- 就是通过一些词语 来定位位置

## 1 基本操作

### 1.1 安装

- 使用elastic search 6.4.2 版本这个 直接bin  运行 elasticsearch.bat文件

## 2 Resultful/JSON Interface

### 2.1 查询Index信息

第二个查查询所有的索引的详细信息

```
GET http://127.0.0.1:9200/shopping
```

```
GET http://127.0.0.1:9200/_cat/indices?v
```

### 2.2 创建索引

```
PUT http://127.0.0.1:9200/shopping
```

```
{
    "acknowledged": true,
    "shards_acknowledged": true,
    "index": "shopping"
}
```

### 2.3 删除索引

就是DELTE其他的一致

#### 2.4 插入文档

使用方法用于插入文档 

post 选了body raw json 

```
http://127.0.0.1:9200/_cat/indices

{
    "title":"小米手机",
    "category":"小米",
    "images":"http://www.gulixueyuan.com/xm.jpg",
    "price":3999.00
}
```

```
{
    "_index": "shopping",
    "_type": "_doc",
    "_id": "FnBY64ABVdvUy_zzKGwk",
    "_version": 1,
    "result": "created",
    "_shards": {
        "total": 2,
        "successful": 1,
        "failed": 0
    },
    "_seq_no": 0,
    "_primary_term": 1
}
```

这个_id 是随机生成的唯一id，因为这个id是随机的，所以多次同样请求返回是不一样的，这就不满足幂等性的原则所以不能用GET 和 PUT ，只能用POST

然后如果想自己生成id post 在后面路径直接加定义的id，这就可以用PUT了，因为id不是随机生成的了，返回就是一样的是幂等的了

```
http://127.0.0.1:9200/shopping/_doc/1001
http://127.0.0.1:9200/shoppint/_create/1001
```

### 2.5 主键搜索

就是路径GET直接写在后面

```
http://127.0.0.1:9200/shopping/_doc/1001
```

```
{
    "_index": "shopping",
    "_type": "_doc",
    "_id": "1001",
    "_version": 2,
    "found": true,
    "_source": {
        "title": "小米手机",
        "category": "小米",
        "images": "http://www.gulixueyuan.com/xm.jpg",
        "price": 3999.00
    }
}
```

### 2.6 全查询

```
http://127.0.0.1:9200/shopping/_search
```

GET 全部如上

## 3 DSL 结构化查询(query dsl)/结构化过滤(filter dsl)

### 3.1 覆盖写

这就是就是全部修改，那么是幂等性的不会出现中间状态，那么直接使用PUT /POST 也可以。就和上面插入一摸一样

### 3.2 局部修改

不是更新全部，那么只能拿使用POST

```
http://127.0.0.1:9200/shopping/_update/1001{
    "doc":{
        "price":5999
    }
}
```

这里要加一个 doc ，这样，然后这里不用_doc 用 _UPDATe 明确告知是修改

### 3.3 删除

就是和之前一样 用 /_doc

### 3.4 条件查询

#### 3.4.1 查询条件放在path中

GET 这里用 _search 

```
http://127.0.0.1:9200/shopping/_search?q=category:小米
```

#### 3.4.2 查询条件放在body 里面

这里这里用 query 然后使用match 匹配模式

```''
http://127.0.0.1:9200/shopping/_search
{
    "query":{
        "match":{
            "category":"小米"
        }
    }
}
```

这里就要讲

query：查询

- term： 主要用于精准匹配那些值，比如数字，日期，布尔值或者not_analyzed 字符串

- match全文搜索：会对内容进行分词 再进行匹配

- match_all：所有的文档都会被查出来

- match_phrate 对内容短语不景行拆分，

- 复杂的搜索：filter query的嵌套

### 3.4.3 分页查询

- 如果查询的结果非常多，那么要进行分页，

```
http://127.0.0.1:9200/shopping/_search
{   
    "query":{
        "match":{
            "category":"小米"
        }
    },
    "from":0,
    "size":1
}
```

### 3.4.4  查询控制

就是减少网络开销 只返回部分字段

然后排序

```
{
    "query":{
        "match_all":{}
    },
    "from":0,
    "size":3,
    "sort":{
        "price":{
            "order":"asc"
        {
    "query":{
        "match_all":{}
    },
    "from":0,
    "size":3,
    "_source":[
        "title"
    ],
    "sort":{
        "price":{
            "order":"asc"
        }
    }
}}
    }
}
```

这里通过 _source 来控制字， sort 排序，order定顺序，asc正序 desc 降序

### 3.4.4 多条件查询

#### 3.4.4.1 must查询

多条件

must 表示的是 每个条件都要符合

```
{
    "query":{
        "bool":{
            "must":[
                {
                    "match":{
                        "category":"小米"
                    }
                },
                {
                    "match":{
                        "title":"小米手机"
                    }
                }
            ]
        }
    }
}
```

#### 3.4.4.2 should 查询

should 相当于

```

```

#### 3.4.4.3 range查询

range 效果是用于范围查询

## 4 聚合

用于结果分组聚合

## 5 分布式介绍

- 索引分片
