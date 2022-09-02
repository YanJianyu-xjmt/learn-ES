# Advanced

## 0 aims

    为二次学习es，这次的目标是能够达到工业级的水平。

[Boolean query | Elasticsearch Guide [8.4] | Elastic](https://www.elastic.co/guide/en/elasticsearch/reference/current/query-dsl-bool-query.html)

## 1 复合查询

### 1.1 boolean query 布尔查询

#### 1.1.1 主要逻辑查询

- must

- filter

- should

- must_not

eg例子

```
POST _search
{
  "query": {
    "bool" : {
      "must" : {
        "term" : { "user.id" : "kimchy" }
      },
      "filter": {
        "term" : { "tags" : "production" }
      },
      "must_not" : {
        "range" : {
          "age" : { "gte" : 10, "lte" : 20 }
        }
      },
      "should" : [
        { "term" : { "tags" : "env1" } },
        { "term" : { "tags" : "deployed" } }
      ],
      "minimum_should_match" : 1,
      "boost" : 1.0
    }
  }
}
```

### 1.1.2 minmun_should_match

    能用minimum_should_match参数可以控制should 必须要匹配的数目，如果bool只有一个should字句，木有must和filter，默认是1。否则是0.

## 、、这里要测试验证一下

### 1.1.3  bool.filter 参与评分

一般情况下，filter 的条件不会影响评分

```
GET _search
{
  "query": {
    "bool": {
      "filter": {
        "term": {
          "status": "active"
        }
      }
    }
  }
}
```

这个例子 中会符合active的文档，但是所有人评分是0

```
GET _search
{
  "query": {
    "bool": {
      "must": {
        "match_all": {}
      },
      "filter": {
        "term": {
          "status": "active"
        }
      }
    }
  }
}
```

这个例子，由于match all 大家都是1.0

```

```

这个例子，由于用了constant_socre大家都是1

## 2 Full text query

full text query 用于分析 text 字段 比如email

的正文。如果分词是用的同一个分词器就ok。



### 2.1 intervals

间隔

按照mtching rules返回文档
