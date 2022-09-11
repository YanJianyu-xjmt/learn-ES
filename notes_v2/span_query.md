# Span query

## 0 introduction

    span queries 是 低级位置query，用于专业的特定的terms上进行专业的顺序和优先级控制。这是经典实现特定的legal文档或者parents



    允许使设置boost 在一个outer span query上。复合span querie 比如span near ，只用list匹配spans 去找到 主要的spans，并产生分数。分数不是符合inner span queries计算，仅仅应用与分数计算，不是spans。



span queries 不能和 non-span queries 混合       



## 1 span_containing

    接受一系列的span queries，但是只有返回span能够匹配span query

## 2 span_field_making query

    允许query 像是 span-near 或者 span-or 交叉不同的领域

## 3 span_first query

    接受一个span query 能够匹配 头N个位置的字段

## 4 span multi

    能够统配 term range prefix wildcard regexp fuzzy

## 5 span not

    匹配一个    span query 并互斥能匹配上的

## 6 span or

    组合span or

## 7 span witin

    会返回一个系列别的span queries
