# cpe-skills

> 全面的网络安全 CPE（通用平台枚举）工具包 —— 解析、匹配、生成、漏洞关联、SBOM 及更多能力。**AI 优先**：专为 AI Agent 直接消费而设计。

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cpe-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cpe-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cpe-skills)](https://goreportcard.com/report/github.com/scagogogo/cpe-skills)
[![Test Coverage](https://img.shields.io/badge/coverage-%E2%89%A5_91%25-brightgreen)](https://github.com/scagogogo/cpe-skills/actions)
[![Release](https://img.shields.io/github/v/release/scagogogo/cpe-skills?include_prereleases)](https://github.com/scagogogo/cpe-skills/releases)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Platforms](https://img.shields.io/badge/platforms-108-blue)](https://github.com/scagogogo/cpe-skills/releases)

**[官网](https://scagogogo.github.io/cpe-skills/) · [English](README.md) · [SKILLS](SKILLS.md) · [文档](https://scagogogo.github.io/cpe-skills/zh/) · [发布](https://github.com/scagogogo/cpe-skills/releases)**

</div>

---

<!-- AI-SUMMARY-START -->

> 本区块为机器可消费的结构化摘要，AI Agent 可直接提取项目元数据、集成方式与能力清单。

| 字段 | 值 |
|------|-----|
| **项目** | cpe-skills |
| **一句话定位** | 全面的 CPE（通用平台枚举）工具包 —— 解析、匹配、生成、漏洞关联、SBOM、VEX。 |
| **语言** | Go（`module github.com/scagogogo/cpe-skills`，要求 Go ≥ 1.18） |
| **覆盖率** | ≥ 91%（CI 对主包门槛 99.9%） |
| **平台** | 每次发布 108 个预编译二进制 —— 9 个操作系统 × 13 种架构 |
| **许可证** | MIT |
| **官网** | https://scagogogo.github.io/cpe-skills/ |
| **仓库** | https://github.com/scagogogo/cpe-skills |

### 集成方式（4 种使用途径）

| 路径 | 适用场景 | 安装 / 配置 |
|------|----------|-------------|
| **SKILLS** | AI / LLM Agent | `https://github.com/scagogogo/cpe-skills` |
| **Go SDK** | Go 应用 | `go get github.com/scagogogo/cpe-skills` |
| **CLI** | Shell / CI / 脚本 | `go install github.com/scagogogo/cpe-skills/cmd/cpe@latest`（或从 Releases 下载二进制） |
| **MCP** | 兼容 MCP 的 AI 客户端 | `command: cpe`, `args: ["mcp", "serve"]` |

### 能力（11 类）

`解析` · `匹配（NISTIR 7696）` · `生成与构建` · `WFN 绑定与转义` · `校验与归一化` · `存储与索引` · `漏洞关联（NVD/OSV/EPSS/KEV）` · `SBOM 与 PURL` · `风险评分与 VEX` · `导出（JSON/CSV/SARIF）` · `基础设施（集合/适用性/错误/日志）`

### 平台矩阵（108 个二进制）

| 操作系统 | 架构 |
|----------|------|
| Linux | 386, amd64, arm64, arm (5/6/7), mips, mips64, mipsle, mips64le, ppc64, ppc64le, riscv64, s390x, loong64 |
| macOS | amd64, arm64 (Apple Silicon) |
| Windows | 386, amd64, arm64 |
| FreeBSD / OpenBSD / NetBSD | 386, amd64, arm64, arm |
| Illumos / Solaris | amd64 |
| AIX | ppc64 |

<!-- AI-SUMMARY-END -->

---

## 快速开始（可直接复制）

### SKILLS —— 面向 AI / LLM

添加到你的 Claude Code skills 配置：

```
https://github.com/scagogogo/cpe-skills
```

### Go SDK

```bash
go get github.com/scagogogo/cpe-skills
```

```go
package main

import (
    "fmt"
    cpeskills "github.com/scagogogo/cpe-skills"
)

func main() {
    // 解析任意 CPE 格式（自动识别 2.2 / 2.3）
    c, _ := cpeskills.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Printf("Vendor: %s, Product: %s, Version: %s\n", c.Vendor, c.ProductName, c.Version)

    // NISTIR 7696 匹配
    matched, _ := cpeskills.QuickMatch(
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
    )
    fmt.Println("Matched:", matched)
}
```

### CLI

```bash
# 方式 A：通过 Go 安装
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest

# 方式 B：从 Releases 下载你平台的预编译二进制
#         → https://github.com/scagogogo/cpe-skills/releases（108 个平台）

# 方式 C：从源码编译
git clone https://github.com/scagogogo/cpe-skills.git
cd cpe-skills && go build -o cpe ./cmd/cpe

# 用法
cpe parse "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
cpe match "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*" \
          "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"
cpe search --vendor apache --product log4j
```

### MCP

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

## 解决了什么问题？

CPE（通用平台枚举）是 NIST 标准命名方案（NIST IR 7695/7696），用于标识 IT 系统、软件和包 —— 它是 CVE 漏洞匹配、SBOM 组件跟踪和供应链安全的基石。

CPE 很难用：两种不兼容格式（2.2 URI 与 2.3 Formatted String）、复杂的 WFN 绑定规则、多源漏洞数据（NVD、OSV、EPSS、KEV）、以及 SBOM ↔ PURL 桥接。**cpe-skills 解决了这一切** —— 单一工具包覆盖完整 CPE 生命周期，通过 4 条集成路径暴露。

![架构图](https://scagogogo.github.io/cpe-skills/architecture_zh.png)

![功能树](https://scagogogo.github.io/cpe-skills/feature_tree_zh.png)

---

## 文档

完整文档位于**[官网](https://scagogogo.github.io/cpe-skills/)**：

- **[使用指南](https://scagogogo.github.io/cpe-skills/zh/guide/)** — 实用使用示例（解析、匹配、WFN、NVD、SBOM……）
- **[API 参考](https://scagogogo.github.io/cpe-skills/zh/api/)** — 完整 API 文档
- **[SKILLS.md](SKILLS.md)** — AI skills 入口

涵盖每项能力（CPE 解析、高级匹配、漏洞关联、SBOM、VEX、导出等）的完整代码示例，请见官网指南。

---

## 贡献

欢迎贡献！请随时提交 Pull Request。

## 许可证

本项目基于 MIT 许可证 —— 详见 [LICENSE](LICENSE) 文件。
