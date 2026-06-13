# Examples

This section provides practical examples demonstrating how to use the CPE library in real-world scenarios. Each example includes complete, runnable code with explanations.

## Available Examples

### Basic Usage
- **[Basic Parsing](./basic-parsing.md)** - Parse CPE strings and access components
- **[CPE Matching](./matching.md)** - Compare and match CPE objects
- **[WFN Conversion](./wfn-conversion.md)** - Convert between CPE and WFN formats

### Advanced Features
- **[Version Comparison](./version-comparison.md)** - Compare version strings and ranges
- **[Applicability Language](./applicability.md)** - Use CPE applicability expressions
- **[CPE Sets](./sets.md)** - Work with collections of CPEs
- **[Advanced Matching](./advanced-matching.md)** - Use sophisticated matching algorithms

### Integration
- **[Storage](./storage.md)** - Persist CPE data with different backends
- **[NVD Integration](./nvd-integration.md)** - Download and use NVD data
- **[CVE Mapping](./cve-mapping.md)** - Map CPEs to vulnerabilities

## Quick Start Example

Here's a simple example to get you started:

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    // Parse a CPE string
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    // Access components
    fmt.Printf("Vendor: %s\n", cpeObj.Vendor)
    fmt.Printf("Product: %s\n", cpeObj.ProductName)
    fmt.Printf("Version: %s\n", cpeObj.Version)
    
    // Create a pattern for matching
    pattern, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")
    
    // Test matching
    if pattern.Match(cpeObj) {
        fmt.Println("CPE matches the Microsoft pattern!")
    }
}
```

## Example Categories

### 🔍 Parsing and Formatting
Learn how to parse CPE strings, handle different formats, and convert between CPE 2.2 and 2.3.

### 🎯 Matching and Comparison
Discover various matching techniques from basic wildcard matching to advanced fuzzy matching with scoring.

### 📊 Data Management
Explore how to store, retrieve, and manage large collections of CPE data efficiently.

### 🔗 External Integration
See how to integrate with external data sources like the National Vulnerability Database.

### 🛡️ Security Applications
Learn how to use CPE for vulnerability management, asset inventory, and security scanning.

## Running the Examples

All examples are complete, standalone programs. To run them:

1. **Install the library:**
   ```bash
   go get github.com/scagogogo/cpe
   ```

2. **Create a new Go file** with the example code

3. **Run the example:**
   ```bash
   go run example.go
   ```

## Example Structure

Each example follows a consistent structure:

- **Overview** - What the example demonstrates
- **Complete Code** - Full, runnable program
- **Explanation** - Step-by-step breakdown
- **Output** - Expected results
- **Variations** - Alternative approaches or extensions

## Common Patterns

### Error Handling
```go
cpeObj, err := cpe.ParseCpe23(cpeString)
if err != nil {
    if cpe.IsInvalidFormatError(err) {
        fmt.Printf("Invalid format: %s\n", cpeString)
        return
    }
    log.Fatal(err)
}
```

### Resource Cleanup
```go
storage, err := cpe.NewFileStorage("./data", true)
if err != nil {
    log.Fatal(err)
}
defer storage.Close()

err = storage.Initialize()
if err != nil {
    log.Fatal(err)
}
```

### Batch Processing
```go
cpeStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0:*:*:*:*:*:*:*",
}

for _, cpeStr := range cpeStrings {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
    if err != nil {
        log.Printf("Failed to parse %s: %v", cpeStr, err)
        continue
    }
    
    // Process the CPE
    fmt.Printf("Processed: %s\n", cpeObj.GetURI())
}
```

## Best Practices

### 1. Always Handle Errors
```go
// Good
cpeObj, err := cpe.ParseCpe23(input)
if err != nil {
    return fmt.Errorf("failed to parse CPE: %w", err)
}

// Bad
cpeObj, _ := cpe.ParseCpe23(input) // Ignoring errors
```

### 2. Use Appropriate Storage
```go
// For testing
storage := cpe.NewMemoryStorage()

// For production
storage, err := cpe.NewFileStorage("./cpe-data", true)
```

### 3. Validate Input
```go
err := cpe.ValidateCPEString(userInput)
if err != nil {
    return fmt.Errorf("invalid CPE format: %w", err)
}
```

### 4. Use Sets for Collections
```go
// Efficient for large collections
cpeSet := cpe.NewCPESet()
cpeSet.Add(cpe1, cpe2, cpe3)

// Filter efficiently
microsoftCPEs := cpeSet.FilterByVendor("microsoft")
```

## Performance Tips

### 1. Enable Caching
```go
storage, _ := cpe.NewFileStorage("./data", true) // Enable cache
```

### 2. Use Batch Operations
```go
// Better than individual operations
cpeSet := cpe.FromArray(cpeArray)
results := cpeSet.FilterByVendor("microsoft")
```

### 3. Reuse Match Options
```go
options := cpe.DefaultMatchOptions()
// Reuse options for multiple matches
for _, cpe := range cpes {
    if cpe.MatchCPE(pattern, cpe, options) {
        // Handle match
    }
}
```

## Getting Help

If you need help with the examples:

1. Check the [API Reference](/api/) for detailed function documentation
2. Look at the complete example code for context
3. Review the error handling patterns
4. Check the [GitHub repository](https://github.com/scagogogo/cpe) for issues and discussions

## Contributing Examples

We welcome contributions of new examples! If you have a useful example that demonstrates a particular use case, please consider contributing it to the project.

## Next Steps

Start with the [Basic Parsing](./basic-parsing.md) example to learn the fundamentals, then explore the more advanced examples based on your specific needs.
