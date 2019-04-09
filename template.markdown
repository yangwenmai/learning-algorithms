# 算法学习(learning-algorithms) 及 [LeetCode](https://leetcode.com) 的题解（Go 版本）{{- /* 本文件是用来生成 README.md 的模板 */}}

[![LeetCode badge](https://leetcode-badge.chyroc.cn/?name=maiyang)](https://leetcode-badge.chyroc.cn/?name=maiyang)
[![LeetCode Ranking](https://img.shields.io/badge/{{.Username}}-{{.Ranking}}-blue.svg)](https://leetcode.com/{{.Username}}/)
[![Build Status](https://travis-ci.org/yangwenmai/learning-algorithms.svg?branch=master)](https://travis-ci.org/yangwenmai/learning-algorithms)
[![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/learning-algorithms)](https://goreportcard.com/report/github.com/yangwenmai/learning-algorithms)
[![Documentation](https://godoc.org/github.com/yangwenmai/learning-algorithms?status.svg)](http://godoc.org/github.com/yangwenmai/learning-algorithms)
[![Coverage Status](https://coveralls.io/repos/github/yangwenmai/learning-algorithms/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/learning-algorithms?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/yangwenmai/learning-algorithms.svg)](https://github.com/yangwenmai/learning-algorithms/issues)
[![license](https://img.shields.io/github/license/yangwenmai/learning-algorithms.svg?maxAge=2592000)](https://github.com/yangwenmai/learning-algorithms/LICENSE)
[![codecov](https://codecov.io/gh/yangwenmai/learning-algorithms/branch/master/graph/badge.svg)](https://codecov.io/gh/yangwenmai/learning-algorithms)

如果你正在算法之门外彷徨，如果你正在发愁不知道该怎么进行算法学习，或许本项目能够适合你。

本项目是我准备的算法学习历程，包括我对算法书籍的整理，算法实践（Go 实现），以及各种算法的理解，也还会有一些有关方面的经验总结。

>希望有更多的人参与进来，咱们一起前行，一起精进。

Inspire by [aQuaYi/LeetCode-in-Go](https://github.com/aQuaYi/LeetCode-in-Go), [WindomZ/leetcode-graphql](https://github.com/WindomZ/leetcode-graphql)

## LeetCode 生成器

> 可以生成 LeetCode 模板，还可以更新 README ，还可以添加待办任务等部分琐碎的工作。

### 配置方法

1. 在 `.gitignore` 中，添加 `*.toml`（避免用户名密码暴露）
1. 在 `learning-algorithms` 目录下，添加文本文件`config.toml`。
1. 把以下内容复制到 `config.toml` 中。

```toml
Username="your leetcode username"
Password="your leetcode password"
```

## 进度

>统计范围：能提交 Go 解答的免费算法题。

{{.ProgressTable}}

### 开发环境

- macOS 10.14.3
- go version go1.12.3 darwin/amd64
- GoLand

## 完成清单

{{.CompletedTable}}

## 未完成清单

{{.AvailableTable}}

以下免费的算法题，暂时不能提交 Go 解答

{{.UnavailableList}}

## 题解

[notes](practice/notes) LeetCode 题解，也会记录我答题过程中对知识点的总结。

----

## 学习方法

1. 仔细研读经典算法书籍（详细书单见后文）
2. 把所有经典数据结构和算法都写一遍
3. 刷题（LeetCode等）
4. 找志同道合的人一起交流学习
5. 参加比赛
6. 。。。

----

## 推荐阅读

### 基础（<b style="color:red">进行中...</b>）

- ☆☆☆☆☆《算法基础-打开算法之门》科尔曼[算法导论作者之一]（<b style="color:blue">阅读进度20%</b>）
- ☆☆☆☆☆《算法导论》 是一本对算法介绍比较全面的经典书籍
- ☆☆☆☆《编程珠玑I》
- ☆☆☆☆《编程珠玑II》
- ☆☆☆《算法-Algorithms》第四版，普林斯顿大学，这本近千页的书只有6章,其中四章分别是排序，查找，图，字符串，足见介绍的有多么的细致。（**暂时不阅读**）

### 算法演示

- [David Galles 的可视化的数据结构和算法](http://www.cs.usfca.edu/~galles/visualization/)
- [酷壳-可视化的数据结构和算法](https://coolshell.cn/articles/4671.html)

### 编程网站

1. [LeetCode](http://leetcode.com/) 全球最大的在线程序评测系统。
2. [OpenJudge](http://openjudge.cn/) OpenJudge是开放的在线程序评测系统 您可以创建自己的OJ小组。

### 笔试刷题必备

- 《剑指offer》
- 《编程之法:面试和算法心得》
- 《算法谜题》
- 《编程之美》

### 延伸阅读

- 《推荐系统-技术、评估及高效算法》
- 《深入理解计算机系统》
- 《TCP/IP详解一二三卷》

## 其他

### 1. [awesome-algorithms](https://github.com/tayllan/awesome-algorithms)

### 2. [July 博客](http://blog.csdn.net/v_july_v)

>比方说：《数据挖掘十大算法系列》、《经典算法研究》、《BAT机器学习面试1000题系列》、《程序员编程艺术》等系列。

- [The-Art-Of-Programming-By-July](https://github.com/julycoding/The-Art-Of-Programming-By-July)
- [程序员编程艺术](http://blog.csdn.net/v_JULY_v/article/details/6460494)

## 如何参与/贡献？

1. Fork 此项目
2. 克隆你自己的项目到你本地 （git clone https://github.com/your_github_name/learning-algorithms.git）
2. 创建你新的 feature 分支 (git checkout -b my_feature)
3. 添加并提交你的修改内容 (git commit -am 'Add some feature')
4. 推送到你自己项目的远端 feature 分支 (git push origin my_feature)
5. 创建一个新的 PR（Pull Request）

## 参考资料

- [阻碍你使用 GraphQL 的十个问题](https://leetcode-cn.com/articles/%E9%98%BB%E7%A2%8D%E4%BD%A0%E4%BD%BF%E7%94%A8-graphql-%E7%9A%84%E5%8D%81%E4%B8%AA%E9%97%AE%E9%A2%98/)

## License

License: BSD 3-Clause License
