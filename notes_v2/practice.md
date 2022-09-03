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

按照mtching rules返回文```

```
POST _search
{
  "query": {
    "intervals" : {
      "my_text" : {
        "all_of" : {
          "ordered" : true,
          "intervals" : [
            {
              "match" : {
                "query" : "my favorite food",
                "max_gaps" : 0,
                "ordered" : true
              }
            },
            {
              "any_of" : {
                "intervals" : [
                  { "match" : { "query" : "hot water" } },
                  { "match" : { "query" : "cold porridge" } }
                ]
              }
            }
          ]
        }
      }
    }
  }
}
```

这里interval  my favorite food 这里不能有间隔 且顺序要对

### 2.1.1 Top level Params

这里是top level 参数

- query
  
  - 这就很简单是 写目标query

- max_gaps
  
  - 最大的间隔

- ordered

    匹配的词必须要是按照顺序的

- analyzer
  
  - 指定分析器

- filter
  
  - 过滤器

- use_FIELD
  
  - 这里可以用里面的field 而不是外面的field

## 2.2 prefix

    前缀terms 匹配

-   prefix 就是指定的query

- analyzer 大概就是指定分析器

- use_field 同上

## 2.3 wildcard

## 2.4 fuzzy

- 这个很吊 模糊匹配

- term 匹配的词

- prefix_length 必须要匹配的前缀长度，默认0

- transpositions 显性的指定 是否可以改变位置 比如 ab-》ba 默认是true

## 这里是试试

- fuzziness 
  
  - matching能够允许的最长的编辑距离

- anaaaaalyzer
  
  - 分析器

- use—field 字段

## 2.5 any_OF

就是任一就可以

- intervals

- max_gaps

- ordred

- filter

## 2.6 同上all_OF

## 2.7 FILTE

- FILTER返回intervals
  
  - after
  
  - before
  
  - contained_by 感觉这个包含
  
  - containing
  
  - not_contained_by
  
  - not_ontaining 不包非常有用
  
  - not_overlapping
  
  - overlapping
  
  - script
