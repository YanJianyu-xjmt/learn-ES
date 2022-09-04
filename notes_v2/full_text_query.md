# Full text queries 全文搜索

## 2 match

### 2.1 简单介绍

返回匹配的文本text,number,date和boolean。

```
GET /_search
{
  "query": {
    "match": {
      "message": {
        "query": "this is a test"
      }
    }
  }
}
```

     这里非常简单

### 2.2 参数

    只有field 就是字段名 比如“message”

     参数如下：

- query

- analyzer

- auto_generate_synonyms_phrase_query

- fuzziness

- max_expansions
  
  - 整个terms 的大小，默认50。估计就是最大是 长度里面的符合条件。

- prefix_length 
  
  - 最长长度

- fuzzy_transpositions
  
  - 改顺序

- fuzzy_rewrite
  
  - 用于重写query，

- operator
  
  - boolean 逻辑 看什么逻辑 orand

- minmun_should_match
  
  - 最小匹配的数目

- zero_TERMS_QUERY
  
  - 不中任何一个关键词可能出

## 3 match boolean prefix

    match_query_prefix query 分析

    这里非常简单

```

GET /_search
{
  "query": {
    "match_bool_prefix" : {
      "message" : "quick brown f"
    }
  }
}
```

```
GET /_search
{
  "query": {
    "bool" : {
      "should": [
        { "term": { "message": "quick" }},
        { "term": { "message": "brown" }},
        { "prefix": { "message": "f"}}
      ]
    }
  }
}
```



    还有should语句也非常好用

```
GET /_search
{
  "query": {
    "match_bool_prefix": {
      "message": {
        "query": "quick brown f",
        "analyzer": "keyword"
      }
    }
  }
}
```

    也能指定analyzer

## 4 match_phrase 短语匹配

非常简单，主要是解决match 会切词的问题

```
GET /_search
{
  "query": {
    "match_phrase": {
      "message": {
        "query": "this is a test",
        "analyzer": "my_analyzer"
      }
    }
  }
}
```

## 5 match_phrase_prefix

前缀匹配

## 6 combine fields

    这个感觉是多个字段综合

```
GET /_search
{
  "query": {
    "combined_fields" : {
      "query":      "database systems",
      "fields":     [ "title", "abstract", "body"],
      "operator":   "and"
    }
  }
}
```

非常简单

## 7 multi-match query

    多filed 匹配，感觉非常实用

```
GET /_search
{
  "query": {
    "multi_match" : {
      "query":    "this is a test", 
      "fields": [ "subject", "message" ] 
    }
  }
}
```

## 8 query

```
GET /_search
{
  "query": {
    "query_string": {
      "query": "(new york city) OR (big apple)",
      "default_field": "content"
    }
  }
}
```

非常简单就收or should

## 9 SImplpe query string

这个非常牛逼

```
GET /_search
{
  "query": {
    "simple_query_string" : {
        "query": "\"fried eggs\" +(eggplant | potato) -frittata",
        "fields": ["title^5", "body"],
        "default_operator": "and"
    }
  }
}
```

```
+ signifies AND operation
| signifies OR operation
- negates a single token
" wraps a number of tokens to signify a phrase for searching
* at the end of a term signifies a prefix query
( and ) signify precedence
~N after a word signifies edit distance (fuzziness)
~N after a phrase signifies slop amount
```

这些简单的符号
