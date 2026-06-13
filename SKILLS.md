# cpe-skills

A comprehensive Go SDK for CPE (Common Platform Enumeration) — providing parsing, matching, generation, storage, and NVD integration for all cybersecurity products.

[![Go Reference](https://pkg.go.dev/badge/github.com/scagogogo/cpe-skills.svg)](https://pkg.go.dev/github.com/scagogogo/cpe-skills)
[![Go Report Card](https://goreportcard.com/badge/github.com/scagogogo/cpe-skills)](https://goreportcard.com/report/github.com/scagogogo/cpe-skills)
[![Test Coverage](https://img.shields.io/badge/coverage-100%25-brightgreen)](https://github.com/scagogogo/cpe-skills)

## Installation

### As a Go Library

```bash
go get github.com/scagogogo/cpe-skills
```

### As a CLI Tool

You can install the `cpe` command-line tool in several ways:

#### Option 1: Install via Go

```bash
go install github.com/scagogogo/cpe-skills/cmd/cpe@latest
```

#### Option 2: Download from GitHub Releases

1. Go to the [Releases page](https://github.com/scagogogo/cpe-skills/releases)
2. Find the latest release
3. Download the archive for your platform:
   - **Linux (amd64)**: `cpe-skills_<version>_linux_x86_64.tar.gz`
   - **Linux (arm64)**: `cpe-skills_<version>_linux_aarch64.tar.gz`
   - **macOS (amd64)**: `cpe-skills_<version>_darwin_x86_64.tar.gz`
   - **macOS (arm64/Apple Silicon)**: `cpe-skills_<version>_darwin_aarch64.tar.gz`
   - **Windows (amd64)**: `cpe-skills_<version>_windows_x86_64.zip`
   - **FreeBSD (amd64)**: `cpe-skills_<version>_freebsd_x86_64.tar.gz`
   - And many more platforms (see the full list in releases)
4. Extract the archive and move the binary to your PATH:

```bash
# Linux/macOS example
tar xzf cpe-skills_*_linux_x86_64.tar.gz
chmod +x cpe
sudo mv cpe /usr/local/bin/
```

```powershell
# Windows example (PowerShell)
Expand-Archive cpe-skills_*_windows_x86_64.zip
Move-Item cpe.exe C:\Windows\
```

#### Option 3: Install via Homebrew (macOS/Linux)

```bash
brew tap scagogogo/tap
brew install cpe-skills
```

#### Verify Installation

```bash
cpe version
```

## Quick Start

```go
package main

import (
    "fmt"
    cpe "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse a CPE string
    c, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Println(c.Vendor)     // microsoft
    fmt.Println(c.ProductName) // windows
    fmt.Println(c.Version)    // 10

    // Match two CPEs
    c2, _ := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    fmt.Println(c.Match(c2)) // true

    // Build a CPE using the builder pattern
    built := cpe.NewBuilder().
        PartApplication().
        Vendor("apache").
        Product("log4j").
        Version("2.14.1").
        Build()

    // Quick match without creating CPE objects
    matched, _ := cpe.QuickMatch(
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
    )
    fmt.Println(matched) // true
}
```

## Features

| Feature | Description | Key Functions |
|---------|-------------|---------------|
| **Parsing** | CPE 2.2 & 2.3 URI parsing | `Parse`, `ParseCpe22`, `ParseCpe23`, `ParseURI` |
| **Formatting** | CPE string generation | `FormatCpe22`, `FormatCpe23`, `FormatURI`, `FormatCPE` |
| **Matching** | NISTIR 7696 name matching | `Match`, `MatchCPE`, `QuickMatch`, `AdvancedMatchCPE` |
| **WFN Binding** | Well-Formed Name operations | `BindToFS`, `BindToURI`, `UnbindFS`, `UnbindURI` |
| **Generation** | CPE creation & composition | `GenerateCPE`, `FuzzyGenerateCPE`, `MergeCPEs`, `RandomCPE` |
| **Builder** | Fluent builder pattern | `NewBuilder().Part().Vendor().Product().Version().Build()` |
| **Escaping** | NISTIR 7695 character escaping | `escapeForFS`, `escapeForURI`, `quoteForWFN` |
| **Validation** | CPE validation | `ValidateCPE`, `ValidateComponent`, `IsCPE23String` |
| **Storage** | In-memory & file-based storage | `MemoryStorage`, `FileStorage`, `StorageManager` |
| **NVD Integration** | NVD data source & CPE dictionary | `DownloadAndParseCPEDict`, `FindCVEsForCPE` |
| **Applicability** | CPE applicability language | `ANDExpression`, `ORExpression`, `ParseExpression` |
| **Version Compare** | Version range comparison | `CompareVersions`, `IsVersionInRange`, `IsSubVersion` |
| **CVE Mapping** | CVE-CPE relationship | `FindVulnerableCPEs`, `QueryByCVE`, `QueryByCPE` |
| **Convenience** | High-level helpers | `MustParse`, `ParseOr`, `Clone`, `FilterByPart` |

## API Reference

### Parsing & Formatting

```go
// Parse any CPE format (auto-detects 2.2 or 2.3)
cpe, err := cpe.Parse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")

// Parse specific format
cpe, err := cpe.ParseCpe22("cpe:/a:microsoft:windows:10")
cpe, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")

// Format to string
str23 := cpe.FormatCpe23(cpe)  // CPE 2.3 format
str22 := cpe.FormatCpe22(cpe)  // CPE 2.2 format
str, _ := cpe.FormatCPE(cpe, "2.3") // Choose version
```

### Matching

```go
// Basic matching
matched := cpe1.Match(cpe2)

// With options
matched := cpe.MatchCPE(criteria, target, &cpe.MatchOptions{
    IgnoreVersion: true,
    AllowSubVersions: true,
    VersionRange: true,
    MinVersion: "2.0",
    MaxVersion: "3.5",
})

// Advanced matching
matched := cpe.AdvancedMatchCPE(criteria, target, &cpe.AdvancedMatchOptions{
    MatchMode: "subset",
    UseFuzzyMatch: true,
    UseRegex: true,
})

// Quick match (strings only)
matched, err := cpe.QuickMatch(cpeStr1, cpeStr2)
```

### Builder Pattern

```go
cpe := cpe.NewBuilder().
    PartApplication().
    Vendor("apache").
    Product("log4j").
    Version("2.14.1").
    Update("sp1").
    Edition("pro").
    Language("en-us").
    Build()
```

### Storage

```go
// In-memory storage
ms := cpe.NewMemoryStorage()
ms.Initialize()
ms.StoreCPE(cpe)
results, _ := ms.SearchCPE(criteria, options)

// File-based storage
fs, _ := cpe.NewFileStorage("/path/to/storage", true)
fs.Initialize()
fs.StoreCPE(cpe)
results, _ := fs.SearchCPE(criteria, options)

// Storage manager with cache
sm := cpe.NewStorageManager(ms)
sm.SetCache(fs)
```

### NVD Integration

```go
// Download and parse NVD CPE dictionary
opts := &cpe.NVDFeedOptions{CacheDir: "/tmp/nvd-cache"}
dict, _ := cpe.DownloadAndParseCPEDict(opts)

// Find CVEs for a CPE
cves, _ := cpe.FindCVEsForCPE(dict, "cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*")
```

### Convenience Functions

```go
// MustParse panics on error (for initialization)
cpe := cpe.MustParse("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")

// ParseOr returns a default on error
cpe := cpe.ParseOr(input, defaultCPE)

// String validation
cpe.IsCPE23String(str) // true/false
cpe.IsCPE22String(str) // true/false

// Clone a CPE
copy := cpe.Clone(original)

// Batch conversions
strs := cpe.CPEsToStrings(cpes)
cpes := cpe.StringsToCPEs(strs)

// Filtering
apps := cpe.FilterByPart(allCPEs, cpe.PartApplication)
msCPEs := cpe.FilterByVendor(allCPEs, "microsoft")
winCPEs := cpe.FilterByProduct(allCPEs, "windows")
```

## CLI Usage

```bash
# Parse a CPE
cpe parse "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"

# Match two CPEs
cpe match "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*" \
          "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*"

# Search CPEs
cpe search --vendor apache --product log4j

# Look up CPE dictionary
cpe dict --nvd
```

## Supported Platforms

The CLI tool (`cpe`) is available for the following platforms via GitHub Releases:

| OS | Architectures |
|----|---------------|
| Linux | 386, amd64, arm64, arm (5/6/7), mips, mips64, mipsle, mips64le, ppc64, ppc64le, riscv64, s390x, loong64 |
| macOS (Darwin) | amd64, arm64 (Apple Silicon) |
| Windows | 386, amd64, arm64 |
| FreeBSD | 386, amd64, arm64, arm |
| OpenBSD | 386, amd64, arm64, arm |
| NetBSD | 386, amd64, arm64, arm |
| Illumos | amd64 |
| Solaris | amd64 |
| AIX | ppc64 |

## License

MIT License
