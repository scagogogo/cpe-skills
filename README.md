# cpe-skills

> A comprehensive CPE (Common Platform Enumeration) toolkit for cybersecurity — parsing, matching, generation, vulnerability correlation, SBOM, and beyond. **AI-first**: designed for AI agents to consume directly.

<div align="center">

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cpe-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cpe-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cpe-skills)](https://goreportcard.com/report/github.com/scagogogo/cpe-skills)
[![Test Coverage](https://img.shields.io/badge/coverage-%E2%89%A5_91%25-brightgreen)](https://github.com/scagogogo/cpe-skills/actions)
[![Release](https://img.shields.io/github/v/release/scagogogo/cpe-skills?include_prereleases)](https://github.com/scagogogo/cpe-skills/releases)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Platforms](https://img.shields.io/badge/platforms-108-blue)](https://github.com/scagogogo/cpe-skills/releases)

**[Website](https://scagogogo.github.io/cpe-skills/) · [简体中文](README_zh.md) · [SKILLS](SKILLS.md) · [Docs](https://scagogogo.github.io/cpe-skills/en/) · [Releases](https://github.com/scagogogo/cpe-skills/releases)**

</div>

---

<!-- AI-SUMMARY-START -->

> This block is structured for machine consumption. AI agents can extract project metadata, integration paths, and capabilities directly from it.

| Field | Value |
|-------|-------|
| **Project** | cpe-skills |
| **One-liner** | Comprehensive CPE (Common Platform Enumeration) toolkit for cybersecurity — parsing, matching, generation, vulnerability correlation, SBOM, VEX. |
| **Language** | Go (`module github.com/scagogogo/cpe-skills`, requires Go ≥ 1.18) |
| **Coverage** | ≥ 91% (CI gate at 99.9% on main package) |
| **Platforms** | 108 prebuilt binaries per release — 9 OSes × 13 architectures |
| **License** | MIT |
| **Website** | https://scagogogo.github.io/cpe-skills/ |
| **Repo** | https://github.com/scagogogo/cpe-skills |

### Integration Paths (4 ways to use)

| Path | Best for | Install / Config |
|------|----------|------------------|
| **SKILLS** | AI / LLM agents | `https://github.com/scagogogo/cpe-skills` |
| **Go SDK** | Go applications | `go get github.com/scagogogo/cpe-skills` |
| **CLI** | Shell / CI / scripts | `go install github.com/scagogogo/cpe-skills/cmd/cpe@latest` (or download binary from Releases) |
| **MCP** | MCP-compatible AI clients | `command: cpe`, `args: ["mcp", "serve"]` |

### Capabilities (11 categories)

`Parsing` · `Matching (NISTIR 7696)` · `Generation & Builder` · `WFN Binding & Escaping` · `Validation & Normalization` · `Storage & Index` · `Vulnerability Correlation (NVD/OSV/EPSS/KEV)` · `SBOM & PURL` · `Risk Scoring & VEX` · `Export (JSON/CSV/SARIF)` · `Infrastructure (Sets/Applicability/Errors/Logging)`

### Platform Matrix (108 binaries)

| OS | Architectures |
|----|---------------|
| Linux | 386, amd64, arm64, arm (5/6/7), mips, mips64, mipsle, mips64le, ppc64, ppc64le, riscv64, s390x, loong64 |
| macOS | amd64, arm64 (Apple Silicon) |
| Windows | 386, amd64, arm64 |
| FreeBSD / OpenBSD / NetBSD | 386, amd64, arm64, arm |
| Illumos / Solaris | amd64 |
| AIX | ppc64 |

<!-- AI-SUMMARY-END -->

---

## Quick Start (copy-paste ready)

### SKILLS — for AI / LLM

Add to your Claude Code skills configuration:

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
    // Parse any CPE format (auto-detect 2.2 / 2.3)
    c, _ := cpeskills.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Printf("Vendor: %s, Product: %s, Version: %s\n", c.Vendor, c.ProductName, c.Version)

    // NISTIR 7696 matching
    matched, _ := cpeskills.QuickMatch(
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
    )
    fmt.Println("Matched:", matched)
}
```

### CLI

```bash
# Option A: install via Go
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest

# Option B: download a prebuilt binary for your platform from Releases
#           → https://github.com/scagogogo/cpe-skills/releases (108 platforms)

# Option C: build from source
git clone https://github.com/scagogogo/cpe-skills.git
cd cpe-skills && go build -o cpe ./cmd/cpe

# Usage
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

## What Problem Does It Solve?

CPE (Common Platform Enumeration) is the NIST-standard naming scheme (NIST IR 7695/7696) for identifying IT systems, software, and packages — it's the backbone of CVE vulnerability matching, SBOM component tracking, and supply chain security.

Working with CPE is hard: two incompatible formats (2.2 URI vs 2.3 Formatted String), complex WFN binding rules, multi-source vulnerability data (NVD, OSV, EPSS, KEV), and SBOM ↔ PURL bridging. **cpe-skills solves all of this** with a single toolkit covering the full CPE lifecycle, exposed through 4 integration paths.

![Architecture](https://scagogogo.github.io/cpe-skills/architecture_en.png)

![Feature Tree](https://scagogogo.github.io/cpe-skills/feature_tree_en.png)

---

## Documentation

Full documentation lives on the **[website](https://scagogogo.github.io/cpe-skills/)**:

- **[Guide](https://scagogogo.github.io/cpe-skills/en/guide/)** — practical usage examples (parsing, matching, WFN, NVD, SBOM, …)
- **[API Reference](https://scagogogo.github.io/cpe-skills/en/api/)** — complete API documentation
- **[SKILLS.md](SKILLS.md)** — AI skills entry point

For comprehensive code examples covering every capability (CPE parsing, advanced matching, vulnerability correlation, SBOM, VEX, export, etc.), see the website guide.

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.
