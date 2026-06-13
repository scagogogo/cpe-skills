# cpe-skills

A comprehensive CPE (Common Platform Enumeration) toolkit — supporting **SKILLS**, **Go SDK**, **CLI**, and **MCP** integration for all cybersecurity products.

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cpe-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cpe-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cpe-skills)](https://goreportcard.com/report/github.com/scagogogo/cpe-skills)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/scagogogo/cpe-skills)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Release](https://img.shields.io/github/v/release/scagogogo/cpe-skills?include_prereleases)](https://github.com/scagogogo/cpe-skills/releases)

**[English](README.md) | [简体中文](README_zh.md) | [SKILLS Documentation](SKILLS.md)**

</div>

---

## 🚀 Quick Integration

### SKILLS (One-Click)

Add to your Claude Code skills configuration:

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

Use as an MCP server for AI-powered CPE operations:

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

## 📖 Introduction

**cpe-skills** is a comprehensive CPE (Common Platform Enumeration) toolkit that provides full lifecycle support for CPE operations — parsing, matching, generation, storage, and NVD integration. It is designed as a foundational SDK that powers cybersecurity products at every layer.

CPE is a standardized naming scheme (NIST IR 7695/7696) for identifying IT systems, software, and packages. This library implements the complete CPE specification, including WFN binding, name matching, applicability language, and CVE correlation.

## ✨ Features

| Category | Description |
|----------|-------------|
| **Parsing** | CPE 2.2 & 2.3 URI parsing with auto-detection |
| **Formatting** | CPE string generation for 2.2 and 2.3 formats |
| **Matching** | NISTIR 7696 name matching (exact, subset, superset, disjoint) |
| **WFN Binding** | Well-Formed Name format with bidirectional conversion |
| **Generation** | CPE creation, fuzzy generation, merging, and random generation |
| **Builder** | Fluent builder pattern for CPE construction |
| **Escaping** | NISTIR 7695 character escaping system |
| **Validation** | CPE and component validation |
| **Version Compare** | Semantic version comparison and range matching |
| **Applicability** | CPE applicability language (AND/OR expressions) |
| **Storage** | In-memory and file-based storage with caching |
| **NVD Integration** | National Vulnerability Database feed integration |
| **CVE Mapping** | CVE-CPE relationship querying |
| **Set Operations** | Union, intersection, difference on CPE collections |
| **Advanced Matching** | Fuzzy, partial, regex, subset, distance-based matching |
| **Convenience API** | MustParse, QuickMatch, Clone, FilterByPart, etc. |

## 📦 Integration Methods

### 1. SKILLS (Recommended for AI/LLM)

SKILLS provides a natural language interface for CPE operations. Add to your AI skills configuration:

```
https://github.com/scagogogo/cpe-skills
```

Once configured, you can ask your AI assistant to:
- Parse and validate CPE strings
- Match CPEs against patterns
- Generate CPE from product information
- Query CVE-CPE relationships

### 2. Go SDK

```go
package main

import (
    "fmt"
    cpe "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse any CPE format
    c, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Printf("Vendor: %s, Product: %s\n", c.Vendor, c.ProductName)

    // Quick match two CPEs
    matched, _ := cpe.QuickMatch(
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
    )
    fmt.Println("Matched:", matched)

    // Builder pattern
    built := cpe.NewBuilder().
        PartApplication().
        Vendor("apache").
        Product("log4j").
        Version("2.14.1").
        Build()

    // Convenience functions
    c2 := cpe.MustParse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    apps := cpe.FilterByPart(allCPEs, cpe.PartApplication)
}
```

### 3. CLI

```bash
# Install
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest

# Or download binary from https://github.com/scagogogo/cpe-skills/releases

# Parse a CPE
cpe parse "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"

# Match two CPEs
cpe match "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*" \
          "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"

# Search CPEs
cpe search --vendor apache --product log4j
```

### 4. MCP (Model Context Protocol)

Use cpe-skills as an MCP server for AI-powered workflows:

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

This enables AI assistants to perform CPE operations through the standardized MCP protocol.

## 🔍 API Reference

### Parsing & Formatting

```go
c, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")  // auto-detect
c, _ := cpe.ParseCpe22("cpe:/a:microsoft:windows:10")                 // CPE 2.2
c, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")  // CPE 2.3

str := cpe.FormatCpe23(c)                    // → "cpe:2.3:a:..."
str, _ := cpe.FormatCPE(c, "2.2")            // → "cpe:/a:..."
```

### Matching

```go
matched := cpe1.Match(cpe2)
matched, _ := cpe.QuickMatch(str1, str2)
matched := cpe.AdvancedMatchCPE(criteria, target, opts)
```

### Generation

```go
c := cpe.GenerateCPE("a", "apache", "log4j", "2.14.1")
c := cpe.FuzzyGenerateCPE("a", "apache", "log4j", "2.x")
c := cpe.NewBuilder().PartApplication().Vendor("apache").Product("log4j").Build()
c := cpe.RandomCPE()
```

### Storage

```go
ms := cpe.NewMemoryStorage()
fs, _ := cpe.NewFileStorage("/data/cpes", true)
```

### Convenience

```go
c := cpe.MustParse(str)                              // panic on error
c := cpe.ParseOr(str, defaultCPE)                    // fallback on error
apps := cpe.FilterByPart(cpes, cpe.PartApplication)  // filter by part
strs := cpe.CPEsToStrings(cpes)                      // batch convert
```

## 🌍 Supported Platforms

| OS | Architectures |
|----|---------------|
| Linux | 386, amd64, arm64, arm (5/6/7), mips, mips64, mipsle, mips64le, ppc64, ppc64le, riscv64, s390x, loong64 |
| macOS | amd64, arm64 (Apple Silicon) |
| Windows | 386, amd64, arm64 |
| FreeBSD | 386, amd64, arm64, arm |
| OpenBSD | 386, amd64, arm64, arm |
| NetBSD | 386, amd64, arm64, arm |
| Illumos | amd64 |
| Solaris | amd64 |
| AIX | ppc64 |

## 📊 Project Stats

- **327+** exported functions
- **976+** test cases
- **100%** test coverage
- **44** platform binaries per release

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
