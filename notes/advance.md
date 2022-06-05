# Advance

## reindex

用于

- 分片数变更

- mapping 字段变更

- 分词规则修改

## redinx

提供redinex这个API，相比重新导入快不少，实测相比bulk导入速度的5-10倍

redinex 不会尝试设置目标索引，不会复制索引的设置 运行_redinx 操作之前设置目标索引，包括设置映射 分片计数 副本等

根据复制源 创建目标索引，然后执行redinex命令

```
Post _redinex
{
    "source": {
        "index":"old_index"    
    },
    "dest":{
        "index":"new_index"
    }
},
```

## 实战

覆盖更新 说明 ”version_type“ ：”internal“ 会盲目的将文档转储到目标中，覆盖任何相同类型和ID文件

```
Post _reindex{
    "source":{
    "index":"twitter"
},
    "dest":{
    "index": "new_twitter"
    "version_type":"internal"
}
}
```

### 创建丢失的文档并更新旧版本的文档

要创建 op_type 设置将导致 _reindex 仅在目标索引中创建丢失的文档，所有存在文档都会引起版本冲突。只要两个索引中存在id相同的记录，就会引起版本冲突。

```
POST _reindex
{
  "source": {
    "index": "twitter"
  },
  "dest": {
    "index": "new_twitter",
    "op_type": "create"
  }
}
```

### 解决冲突

默认情况下，版本冲突会终止 _redinex进程 冲突请求正文参数可用于 执行_redinex 继续处理有关版本冲突的下一个文档，需要注意的是，其他错误类型，不受参数的影响 

当设置”conflicts“：”proceed“在请求正文中设置的时候，_reindex 进程将继续处理版本冲突并返回遇到的版本冲突计数

```
POST _reindex
{
  "conflicts": "proceed",
  "source": {
    "index": "twitter"
  },
  "dest": {
    "index": "new_twitter",
    "op_type": "create"
  }
}
```

### source 中添加查询条件

```
POST _reindex
{
  "source": {
    "index": "twitter",
    "query": {
      "term": {
        "user": "kimchy"
      }
    }
  },
  "dest": {
    "index": "new_twitter"
  }
}
```



[ES索引重建reindex详解_斗者_2013的博客-CSDN博客_es reindex](https://blog.csdn.net/w1014074794/article/details/120483334)

[ES索引重建reindex详解_斗者_2013的博客-CSDN博客_es reindex](https://blog.csdn.net/w1014074794/article/details/120483334)



大概就是





## 性能优化

- 常规的如果是进行 少量数据迁移，reindex就能达到很好的要求，如果数据量太大，reindex速度会变得很慢

- reindex 的核心做跨索引 跨集群的数据迁移
  
  慢的原因及优化思路无非包括：
  
  - 批量大小值可能太小，需要结合堆内存 线程池调整大小
  
  - redinex 的底层实现是scroll实现，借助scroll并行优化 提升效率
  
  - 跨索引 跨集群的核心是写入数据 考虑写入优化的角度提升效率

- 批量提升写入大小值

- 提高scroll的并行度

```
POST _reindex
{
  "source": {
    "index": "source",
    "size": 5000
  },
  "dest": {
    "index": "dest"
  }
}
```

写入size 

```
POST _reindex?slices=5&refresh
{
  "source": {
    "index": "twitter"
  },
  "dest": {
    "index": "new_twitter"
  }
```

这里是提供 slice  slice scroll 来并行化 重新索引 过程， 这种sliciing 可以设定分为两种方式，手动设置分片， 自动设置分片



如果是自动设置分片 slices 设置位auto    

当slices 数量等于索引中的分片数量时， 查询性能最高效，slices 大小大于分数目，非但不会提升效率，反而增加性能

如果slices 数字很大，例如500 建议选择一个较低的数字，因为过大slices 会影响性能



实践证明 比默认设置redinex 速度可以增加10倍



es 中 scroll 技术，如果一次性要查询出一大批的数据，那么性能很差，那么性能会很差，一般采用scroll滚动查询，一批一批查询，知道所有的查询完成使用scroll 滚动搜索，可以搜索出一批数据  采用基于_doc 排序，性能比较高





https://blog.csdn.net/paicMis/article/details/84113057[ES中scroll技术_Xlucas的博客-CSDN博客_es scroll](https://blog.csdn.net/paicMis/article/details/84113057)
















