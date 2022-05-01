# DSL

## 0 查询类别

### 0.1 叶子查询  

- 在特定的字段上查找特定的值 match term range 查询。这些查询可以自己使用

### 0.2 复合查询子句

- 包含其他叶查询或复合查询子句，以合理的方式结合多条查询(bool 或 dis_max 查询)，或者改变查询行为(not 或constant_score查询)

### 0.3 查询 query

- 用于检查内容与条件是否匹配，并且计算_score 元字段表示匹配度，查询的结构中以query参数开始执行内容查询

### 0.4 过滤 filter

- filter 不计算匹配得分，知识确定文档是否匹配，一般就回被缓存提高性能



### 1 布尔查询

- 分析文本并组成一个布尔查询，operate 设置为“or” 或者“and” 。 可以用于匹配should子句 最小数量可以使用minimum_should_match 参数来设置
- 可以设置 analyzer 来控制在文本上执行分析过程的分词器，默认是字段映射中明确定义或者默认搜索分词器
- lenient 参数可以设置为true 忽略数据类型匹配出错造成的异常，可以尝试通过文本查询字符串来查询数字类型字段，默认false



- 模糊匹配 fuzziness 可以请求字段类型进行模糊匹配

```http
{
	"match":{
		"message" : {"query":"this is a test","operator":"and"}
	}
}
```

- 零索引词查询 如果查询使用的分词器移除所有词元，默认行为是不匹配任何文档，使用 zero_terms_query 选项进行修改，接受none(默认)和all 相当于 match_all 查询

```json
{
	"match":{
		"message":{
			"query": "to be or not to be",
			"operateor":"and",
			"zero_terms_query":"all"
		}
	}
}
```



- 短语查询
  - 短语查询 分析文本且创建短语查询

```
{
	"match_phrase":{
		"message"： "this is a test"
	}
}
```

​	因为短语查询只是标准查询的一个类型， 可以用于一下方式使用

```json
{
	"match":{
		"message":{"query":"this is a test","type":"phrase"}
	}
}

{
	"match":{
		"message":{"query":"this is a test","analyzer":"my_analyzer"}
	}
}
```

slop 可配置的slop 匹配索引词 

- 短语前缀匹配

  可以对文本最后一个字段进行 前缀匹配 例如：

  ```json
  {
  	"match_phrase_prefix":{"message":"this is a test"}
  }
  ```

  

### 2 多字段查询

在标准查询

```
{
	"multi_match":{
		"query":"this is a test",
		"fields": ["subject","message"]
	}
}
```

字段可以通过通配符指定， 个别字段可以用caret(^）进行加权，比如

```
{
	"multi_match":{
		"query":"this is a test",
		"fields": ["subject^3","message"]
	}
}
```



说明 subject 重要3倍



然后这里 must 相当于 and

shold 相当于 or

must_not 相当于 and not