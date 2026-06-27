---
layout: home

hero:
  name: "CPE Library"
  text: "Common Platform Enumeration for Go"
  tagline: "A comprehensive Go library for parsing, matching, and managing CPE (Common Platform Enumeration) information"
  actions:
    - theme: brand
      text: Get Started
      link: /api/
    - theme: alt
      text: View on GitHub
      link: https://github.com/scagogogo/cpe-skills

features:
  - title: CPE 2.2 & 2.3 Support
    details: Full support for both CPE 2.2 and 2.3 formats with parsing and generation capabilities
  - title: Advanced Matching
    details: Sophisticated matching algorithms including fuzzy matching, regex support, and version comparison
  - title: WFN Support
    details: Complete Well-Formed Name (WFN) format support with bidirectional conversion
  - title: NVD Integration
    details: Built-in integration with National Vulnerability Database for vulnerability mapping
  - title: Storage Backends
    details: Multiple storage backends including file-based and memory storage with caching
  - title: CPE Sets
    details: Set operations for CPE collections including union, intersection, and difference
---

## Quick Start

Install the library:

```bash
go get github.com/scagogogo/cpe-skills
```

Parse a CPE string:

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse CPE 2.3 format
    cpeObj, err := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Vendor: %s\n", cpeObj.Vendor)
    fmt.Printf("Product: %s\n", cpeObj.ProductName)
    fmt.Printf("Version: %s\n", cpeObj.Version)
}
```

## Features

### 🔍 Parsing & Formatting
- Parse CPE 2.2 and 2.3 format strings
- Generate CPE strings from structured data
- Validate CPE format and components

### 🎯 Matching & Comparison
- Basic CPE matching with wildcard support
- Advanced matching with fuzzy logic
- Version comparison and range matching
- Regular expression matching

### 📚 Dictionary Support
- Parse NVD CPE Dictionary XML
- Store and retrieve CPE dictionaries
- Search and filter dictionary entries

### 🔗 NVD Integration
- Download and parse NVD CPE feeds
- Map CPEs to CVE vulnerabilities
- Automatic data updates and caching

### 💾 Storage
- File-based storage with JSON format
- Memory storage for testing
- Caching layer for performance
- Pluggable storage interface

### 🧮 Set Operations
- Create and manage CPE sets
- Union, intersection, and difference operations
- Filter sets with advanced criteria

## Documentation

- [API Reference](/api/) - Complete API documentation
- [Examples](/examples/) - Practical usage examples
- [GitHub Repository](https://github.com/scagogogo/cpe-skills) - Source code and issues

---

*Last updated: $(date)*
