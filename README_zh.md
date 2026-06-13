# cpe-skills

全面的 CPE（通用平台枚举）工具包 — 支持 **SKILLS**、**Go SDK**、**CLI** 和 **MCP** 接入，覆盖所有网络安全产品需求。

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cpe-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cpe-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cpe-skills)](https://goreportcard.com/report/github.com/scagogogo/cpe-skills)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/scagogogo/cpe-skills)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/scagogogo/cpe-skills?include_prereleases)](https://github.com/scagogogo/cpe-skills/releases)

**[English](README.md) | [简体中文](README_zh.md) | [SKILLS 文档](SKILLS.md)**

</div>

---

## 🚀 快速接入

### SKILLS（一键接入）

添加到你的 Claude Code skills 配置中：

```
https://github.com/scagogogo/cpe-skills
```

### Go SDK

```bash
go get github.com/scagogogo/cpe-skills
```

### CLI

```bash
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest
```

### MCP

作为 MCP 服务器用于 AI 驱动的 CPE 操作：

```json
{
  "mcpServers": {
    "cpe-skills": {
      "command": "cpe",
      "args": ["mcp", "serve"]
    }
  }
}
```

---

## 📖 简介

**cpe-skills** 是一个全面的 CPE（通用平台枚举）工具包，提供 CPE 全生命周期支持 — 解析、匹配、生成、存储和 NVD 集成。它作为底层支撑 SDK，为各层网络安全产品提供能力。

CPE 是一种标准化命名方案（NIST IR 7695/7696），用于标识 IT 系统、软件和软件包。本库实现了完整的 CPE 规范，包括 WFN 绑定、名称匹配、适用性语言和 CVE 关联。

## ✨ 功能特性

| 类别 | 描述 |
|------|------|
| **解析** | CPE 2.2 & 2.3 URI 解析，自动检测格式 |
| **格式化** | CPE 2.2 和 2.3 格式的字符串生成 |
| **匹配** | NISTIR 7696 名称匹配（精确、子集、超集、不相交） |
| **WFN 绑定** | Well-Formed Name 格式双向转换 |
| **生成** | CPE 创建、模糊生成、合并和随机生成 |
| **构建器** | 流式 Builder 模式构建 CPE |
| **转义** | NISTIR 7695 字符转义系统 |
| **校验** | CPE 和组件校验 |
| **版本比较** | 语义化版本比较和范围匹配 |
| **适用性语言** | CPE 适用性语言（AND/OR 表达式） |
| **存储** | 内存和文件存储，支持缓存 |
| **NVD 集成** | 国家漏洞数据库集成 |
| **CVE 映射** | CVE-CPE 关系查询 |
| **集合操作** | CPE 集合的并集、交集、差集 |
| **高级匹配** | 模糊、部分、正则、子集、距离匹配 |
| **便捷 API** | MustParse、QuickMatch、Clone、FilterByPart 等 |

## 📦 接入方式

### 1. SKILLS（推荐用于 AI/LLM）

SKILLS 提供自然语言接口用于 CPE 操作。添加到你的 AI skills 配置中：

```
https://github.com/scagogogo/cpe-skills
```

配置完成后，你可以让 AI 助手：
- 解析和校验 CPE 字符串
- 将 CPE 与模式匹配
- 从产品信息生成 CPE
- 查询 CVE-CPE 关系

### 2. Go SDK

```go
package main

import (
    "fmt"
    cpe "github.com/scagogogo/cpe-skills"
)

func main() {
    // 解析任意 CPE 格式
    c, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Printf("供应商: %s, 产品: %s\n", c.Vendor, c.ProductName)

    // 快速匹配两个 CPE
    matched, _ := cpe.QuickMatch(
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
    )
    fmt.Println("匹配结果:", matched)

    // Builder 模式
    built := cpe.NewBuilder().
        PartApplication().
        Vendor("apache").
        Product("log4j").
        Version("2.14.1").
        Build()

    // 便捷函数
    c2 := cpe.MustParse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    apps := cpe.FilterByPart(allCPEs, cpe.PartApplication)
}
```

### 3. CLI

```bash
# 安装
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest

# 或者从 https://github.com/scagogogo/cpe-skills/releases 下载二进制文件

# 解析 CPE
cpe parse "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"

# 匹配两个 CPE
cpe match "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*" \
          "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"

# 搜索 CPE
cpe search --vendor apache --product log4j
```

### 4. MCP（模型上下文协议）

将 cpe-skills 作为 MCP 服务器用于 AI 驱动的工作流：

```json
{
  "mcpServers": {
    "cpe-skills": {
      "command": "cpe",
      "args": ["mcp", "serve"]
    }
  }
}
```

这使 AI 助手能够通过标准化的 MCP 协议执行 CPE 操作。

## 🔍 API 参考

### 解析与格式化

```go
c, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")  // 自动检测
c, _ := cpe.ParseCpe22("cpe:/a:microsoft:windows:10")                 // CPE 2.2
c, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")  // CPE 2.3

str := cpe.FormatCpe23(c)                    // → "cpe:2.3:a:..."
str, _ := cpe.FormatCPE(c, "2.2")            // → "cpe:/a:..."
```

### 匹配

```go
matched := cpe1.Match(cpe2)
matched, _ := cpe.QuickMatch(str1, str2)
matched := cpe.AdvancedMatchCPE(criteria, target, opts)
```

### 生成

```go
c := cpe.GenerateCPE("a", "apache", "log4j", "2.14.1")
c := cpe.FuzzyGenerateCPE("a", "apache", "log4j", "2.x")
c := cpe.NewBuilder().PartApplication().Vendor("apache").Product("log4j").Build()
c := cpe.RandomCPE()
```

### 存储

```go
ms := cpe.NewMemoryStorage()
fs, _ := cpe.NewFileStorage("/data/cpes", true)
```

### 便捷函数

```go
c := cpe.MustParse(str)                              // 出错 panic
c := cpe.ParseOr(str, defaultCPE)                    // 出错返回默认值
apps := cpe.FilterByPart(cpes, cpe.PartApplication)  // 按 Part 筛选
strs := cpe.CPEsToStrings(cpes)                      // 批量转换
```

## 🌍 支持平台

| 操作系统 | 架构 |
|----------|------|
| Linux | 386, amd64, arm64, arm (5/6/7), mips, mips64, mipsle, mips64le, ppc64, ppc64le, riscv64, s390x, loong64 |
| macOS | amd64, arm64 (Apple Silicon) |
| Windows | 386, amd64, arm64 |
| FreeBSD | 386, amd64, arm64, arm |
| OpenBSD | 386, amd64, arm64, arm |
| NetBSD | 386, amd64, arm64, arm |
| Illumos | amd64 |
| Solaris | amd64 |
| AIX | ppc64 |

## 📊 项目统计

- **327+** 导出函数
- **976+** 测试用例
- **100%** 测试覆盖率
- **44** 平台二进制文件（每次发布）

## 🤝 贡献

欢迎贡献！请随时提交 Pull Request。

## 📄 许可证

本项目采用 MIT 许可证 — 详见 [LICENSE](LICENSE) 文件。
