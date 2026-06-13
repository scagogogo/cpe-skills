# CPE - Common Platform Enumeration Library

<div align="center">

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)

**[English](README.md) | [简体中文](README_zh.md)**

</div>

### 📚 Documentation

**Complete API documentation and examples: [https://scagogogo.github.io/cpe/](https://scagogogo.github.io/cpe/)**

- [API Reference](https://scagogogo.github.io/cpe/api/) - Complete API documentation
- [Examples](https://scagogogo.github.io/cpe/examples/) - Practical code examples
- [Quick Start Guide](https://scagogogo.github.io/cpe/api/) - Getting started tutorial

### 📖 Introduction

The CPE (Common Platform Enumeration) library is a comprehensive Go implementation for processing, parsing, matching, and storing CPE (Common Platform Enumeration) data. CPE is a structured naming scheme for identifying classes of IT systems, software, and packages.

The library also includes integration with CVE (Common Vulnerabilities and Exposures), enabling developers to associate software components with known security vulnerabilities.

### ✨ Features

- **CPE Format Support**: Parse and generate CPE 2.2 and 2.3 formats
- **Advanced Matching**: CPE name matching with wildcards and special values
- **WFN Support**: Well-Formed Name format with bidirectional conversion
- **Applicability Language**: CPE Applicability Language support
- **Version Comparison**: Semantic version comparison and range matching
- **Dictionary Management**: CPE dictionary with XML import/export
- **CVE Integration**: Associate CPEs with Common Vulnerabilities and Exposures
- **Advanced Algorithms**: Fuzzy matching, subset/superset matching
- **Set Operations**: Union, intersection, difference operations on CPE collections
- **NVD Integration**: Built-in National Vulnerability Database feed integration
- **Error Handling**: Structured error handling with detailed error types
- **Storage Backends**: Multiple storage backends with persistence support
- **Caching**: Integrated caching mechanism for optimized performance

### 🚀 Installation

Install using Go modules:

```bash
go get github.com/scagogogo/cpe-skills
```

### 🔍 Quick Start

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse CPE 2.3 string
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Vendor: %s, Product: %s, Version: %s\n", 
        cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
    
    // Create matching pattern
    pattern, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")
    
    // Test matching
    if pattern.Match(cpeObj) {
        fmt.Println("CPE matches the pattern!")
    }
}
```

### 🏗️ Architecture

The library follows a modular design with the following core components:

1. **CPE Parser Engine**: Handles parsing and formatting of CPE strings
2. **Matching Engine**: Implements various CPE matching strategies
3. **Storage System**: Provides multiple storage backend options
4. **CVE Integration**: Connects CPE data with vulnerability information
5. **NVD Adapter**: Integrates with National Vulnerability Database

### 📝 Local Documentation Development

To run and develop documentation locally:

```bash
# Navigate to docs directory
cd docs

# Install dependencies
npm install

# Start development server
npm run docs:dev

# Build documentation
npm run docs:build

# Preview built documentation
npm run docs:preview
```

Documentation will be available at `http://localhost:5173` (dev mode) or `http://localhost:4173` (preview mode).

### 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
