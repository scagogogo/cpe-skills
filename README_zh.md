# CPE - 通用平台枚举库

<div align="center">

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)

**[English](README.md) | [简体中文](README_zh.md)**

</div>

## 📚 文档

**完整的API文档和使用示例请访问：[https://scagogogo.github.io/cpe/zh/](https://scagogogo.github.io/cpe/zh/)**

- [API 参考文档](https://scagogogo.github.io/cpe/zh/api/) - 完整的API文档
- [使用示例](https://scagogogo.github.io/cpe/zh/examples/) - 实用的代码示例
- [快速开始指南](https://scagogogo.github.io/cpe/zh/api/) - 快速上手教程

## 📖 简介

CPE (Common Platform Enumeration) 库是一个完整的Go语言实现，用于处理、解析、匹配和存储CPE (通用平台枚举)。CPE是一种结构化命名方案，用于标识IT系统、软件和软件包的类别。

该库还包括与CVE (Common Vulnerabilities and Exposures) 集成的功能，使开发者能够将软件组件与已知的安全漏洞关联起来。

## ✨ 功能特性

- **CPE格式支持**：解析和生成CPE 2.2和2.3格式
- **高级匹配**：支持通配符和特殊值的CPE名称匹配
- **WFN支持**：Well-Formed Name格式及双向转换
- **适用性语言**：CPE适用性语言支持
- **版本比较**：语义化版本比较和范围匹配
- **字典管理**：CPE字典及XML导入导出
- **CVE集成**：将CPE与通用漏洞披露关联
- **高级算法**：模糊匹配、子集/超集匹配
- **集合操作**：CPE集合的并集、交集、差集操作
- **NVD集成**：内置国家漏洞数据库集成
- **错误处理**：结构化错误处理和详细错误类型
- **存储后端**：多种存储后端及持久化支持
- **缓存机制**：集成缓存机制优化性能

## 🚀 安装

使用Go模块安装：

```bash
go get github.com/scagogogo/cpe-skills
```

## 🔍 快速开始

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // 解析CPE 2.3字符串
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("供应商: %s, 产品: %s, 版本: %s\n", 
        cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
    
    // 创建匹配模式
    pattern, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")
    
    // 测试匹配
    if pattern.Match(cpeObj) {
        fmt.Println("CPE匹配成功！")
    }
}
```

## 🏗️ 系统架构

该库采用模块化设计，包含以下核心组件：

1. **CPE解析引擎**：处理CPE字符串的解析和格式化
2. **匹配引擎**：实现各种CPE匹配策略
3. **存储系统**：提供多种存储后端选项
4. **CVE集成**：连接CPE数据与漏洞信息
5. **NVD适配器**：与国家漏洞数据库集成

## 📝 本地文档开发

在本地运行和开发文档：

```bash
# 进入文档目录
cd docs

# 安装依赖
npm install

# 启动开发服务器
npm run docs:dev

# 构建文档
npm run docs:build

# 预览构建后的文档
npm run docs:preview
```

文档将在 `http://localhost:5173`（开发模式）或 `http://localhost:4173`（预览模式）上运行。

## 🤝 贡献

欢迎贡献！请随时提交Pull Request。

## 📄 许可证

本项目采用MIT许可证 - 详见 [LICENSE](LICENSE) 文件。
