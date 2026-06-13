---
layout: home

hero:
  name: "CPE 库"
  text: "Go语言通用平台枚举库"
  tagline: "一个用于解析、匹配和管理CPE（通用平台枚举）信息的综合性Go语言库"
  actions:
    - theme: brand
      text: 开始使用
      link: /zh/api/
    - theme: alt
      text: 查看GitHub
      link: https://github.com/scagogogo/cpe-skills

features:
  - title: CPE 2.2 & 2.3 支持
    details: 完全支持CPE 2.2和2.3格式的解析和生成功能
  - title: 高级匹配算法
    details: 复杂的匹配算法，包括模糊匹配、正则表达式支持和版本比较
  - title: WFN 支持
    details: 完整的Well-Formed Name（WFN）格式支持及双向转换
  - title: NVD 集成
    details: 内置国家漏洞数据库集成，用于漏洞映射
  - title: 存储后端
    details: 多种存储后端，包括基于文件和内存存储，支持缓存
  - title: CPE 集合
    details: CPE集合的集合操作，包括并集、交集和差集
---

## 快速开始

安装库：

```bash
go get github.com/scagogogo/cpe-skills
```

解析CPE字符串：

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // 解析CPE 2.3格式
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("供应商: %s\n", cpeObj.Vendor)
    fmt.Printf("产品: %s\n", cpeObj.ProductName)
    fmt.Printf("版本: %s\n", cpeObj.Version)
}
```

## 功能特性

### 🔍 解析和格式化
- 解析CPE 2.2和2.3格式字符串
- 从结构化数据生成CPE字符串
- 验证CPE格式和组件

### 🎯 匹配和比较
- 支持通配符的基本CPE匹配
- 带有模糊逻辑的高级匹配
- 版本比较和范围匹配
- 正则表达式匹配

### 📚 字典支持
- 解析NVD CPE字典XML
- 存储和检索CPE字典
- 搜索和过滤字典条目

### 🔗 NVD集成
- 下载和解析NVD CPE数据源
- 将CPE映射到CVE漏洞
- 自动数据更新和缓存

### 💾 存储
- 基于JSON格式的文件存储
- 用于测试的内存存储
- 性能缓存层
- 可插拔存储接口

### 🧮 集合操作
- 创建和管理CPE集合
- 并集、交集和差集操作
- 使用高级条件过滤集合

## 文档

- [API 参考文档](/zh/api/) - 完整的API文档
- [使用示例](/zh/examples/) - 实用的代码示例
- [GitHub 仓库](https://github.com/scagogogo/cpe-skills) - 源代码和问题反馈
