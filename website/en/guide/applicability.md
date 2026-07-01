# Applicability Language

This example demonstrates how to use CPE Applicability Language for expressing complex matching conditions and logical relationships between CPE names.

## Overview

CPE Applicability Language allows you to create sophisticated expressions that define when a particular piece of information (like a vulnerability) applies to a system. It supports logical operators (AND, OR, NOT) and complex nested conditions.

An applicability expression is parsed into a logical expression tree. The diagram below shows how a rule like `(CPE_A AND CPE_B) OR (NOT CPE_C)` is represented: an `OR` root combines an `AND` branch and a `NOT` branch, with concrete CPE names as the leaves.

```mermaid
flowchart TD
    ROOT{"OR"}
    AND_NODE{"AND"}
    NOT_NODE{"NOT"}
    A["CPE_A: windows 10"]
    B["CPE_B: internet_explorer"]
    C["CPE_C: windows_update kb5005565"]
    ROOT --> AND_NODE
    ROOT --> NOT_NODE
    AND_NODE --> A
    AND_NODE --> B
    NOT_NODE --> C
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE Applicability Language Examples ===")
    
    // Example 1: Basic Applicability Expressions
    fmt.Println("\n1. Basic Applicability Expressions:")
    
    // Simple expression: applies to Windows 10
    expr1 := "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*"
    
    // OR expression: applies to Windows 10 OR Windows 11
    expr2 := `(cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR 
               cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*)`
    
    // AND expression: applies to Windows 10 AND specific update
    expr3 := `(cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* AND 
               cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*)`
    
    expressions := []struct {
        name string
        expr string
        desc string
    }{
        {"Simple", expr1, "Single CPE match"},
        {"OR Logic", expr2, "Multiple alternatives"},
        {"AND Logic", expr3, "Multiple requirements"},
    }
    
    for _, e := range expressions {
        fmt.Printf("\n%s Expression:\n", e.name)
        fmt.Printf("  Description: %s\n", e.desc)
        fmt.Printf("  Expression: %s\n", e.expr)
        
        // Parse the expression
        parsedExpr, err := cpeskills.ParseApplicabilityExpression(e.expr)
        if err != nil {
            log.Printf("Failed to parse expression: %v", err)
            continue
        }
        
        fmt.Printf("  Parsed successfully: %t\n", parsedExpr != nil)
    }
    
    // Example 2: Complex Nested Expressions
    fmt.Println("\n2. Complex Nested Expressions:")
    
    // Complex vulnerability applicability
    complexExpr := `
    (
        (cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR 
         cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*) 
        AND 
        (cpe:2.3:a:microsoft:internet_explorer:*:*:*:*:*:*:*:* OR 
         cpe:2.3:a:microsoft:edge:*:*:*:*:*:*:*:*)
        AND NOT 
        cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*
    )`
    
    fmt.Printf("Complex Expression:\n%s\n", complexExpr)
    
    parsedComplex, err := cpeskills.ParseApplicabilityExpression(complexExpr)
    if err != nil {
        log.Printf("Failed to parse complex expression: %v", err)
    } else {
        fmt.Printf("Successfully parsed complex expression\n")
        fmt.Printf("Expression type: %s\n", parsedComplex.Type())
        fmt.Printf("Number of operands: %d\n", len(parsedComplex.Operands()))
    }
    
    // Example 3: Testing Applicability
    fmt.Println("\n3. Testing Applicability:")
    
    // Define test systems
    testSystems := []struct {
        name string
        cpes []string
    }{
        {
            "Windows 10 with IE",
            []string{
                "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:internet_explorer:11:*:*:*:*:*:*:*",
            },
        },
        {
            "Windows 11 with Edge",
            []string{
                "cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:edge:95.0.1020.44:*:*:*:*:*:*:*",
            },
        },
        {
            "Windows 10 with patch",
            []string{
                "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:internet_explorer:11:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*",
            },
        },
        {
            "Linux system",
            []string{
                "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
                "cpe:2.3:a:mozilla:firefox:95.0:*:*:*:*:*:*:*",
            },
        },
    }
    
    // Test each system against the complex expression
    for _, system := range testSystems {
        fmt.Printf("\nTesting system: %s\n", system.name)
        
        // Convert CPE strings to objects
        systemCPEs := make([]*cpeskills.CPE, 0, len(system.cpes))
        for _, cpeStr := range system.cpes {
            cpeObj, err := cpeskills.ParseCpe23(cpeStr)
            if err != nil {
                log.Printf("Failed to parse CPE %s: %v", cpeStr, err)
                continue
            }
            systemCPEs = append(systemCPEs, cpeObj)
        }
        
        // Test applicability
        applies := cpeskills.EvaluateApplicability(parsedComplex, systemCPEs)
        
        status := "❌ Not applicable"
        if applies {
            status = "✅ Applicable"
        }
        
        fmt.Printf("  Result: %s\n", status)
        fmt.Printf("  System CPEs:\n")
        for _, cpeStr := range system.cpes {
            fmt.Printf("    - %s\n", cpeStr)
        }
    }
    
    // Example 4: Version Range Applicability
    fmt.Println("\n4. Version Range Applicability:")
    
    // Expression for Java versions 8.0 to 11.0 (exclusive)
    javaRangeExpr := `
    (cpe:2.3:a:oracle:java:8.*:*:*:*:*:*:*:* OR
     cpe:2.3:a:oracle:java:9.*:*:*:*:*:*:*:* OR
     cpe:2.3:a:oracle:java:10.*:*:*:*:*:*:*:*)
    `
    
    fmt.Printf("Java Version Range Expression:\n%s\n", javaRangeExpr)
    
    javaExpr, err := cpeskills.ParseApplicabilityExpression(javaRangeExpr)
    if err != nil {
        log.Printf("Failed to parse Java expression: %v", err)
    } else {
        // Test different Java versions
        javaVersions := []string{
            "cpe:2.3:a:oracle:java:7.0.80:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:8.0.291:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:9.0.4:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:17.0.1:*:*:*:*:*:*:*",
        }
        
        fmt.Println("\nTesting Java versions:")
        for _, javaVer := range javaVersions {
            javaCPE, _ := cpeskills.ParseCpe23(javaVer)
            applies := cpeskills.EvaluateApplicability(javaExpr, []*cpeskills.CPE{javaCPE})
            
            status := "❌"
            if applies {
                status = "✅"
            }
            
            fmt.Printf("  %s %s\n", status, javaVer)
        }
    }
    
    // Example 5: Platform-Specific Applicability
    fmt.Println("\n5. Platform-Specific Applicability:")
    
    // Expression for web servers on Linux
    webServerLinuxExpr := `
    (cpe:2.3:o:*:linux:*:*:*:*:*:*:*:* OR
     cpe:2.3:o:canonical:ubuntu:*:*:*:*:*:*:*:* OR
     cpe:2.3:o:redhat:enterprise_linux:*:*:*:*:*:*:*:*)
    AND
    (cpe:2.3:a:apache:http_server:*:*:*:*:*:*:*:* OR
     cpe:2.3:a:nginx:nginx:*:*:*:*:*:*:*:*)
    `
    
    fmt.Printf("Web Server on Linux Expression:\n%s\n", webServerLinuxExpr)
    
    webServerExpr, err := cpeskills.ParseApplicabilityExpression(webServerLinuxExpr)
    if err != nil {
        log.Printf("Failed to parse web server expression: %v", err)
    } else {
        // Test different server configurations
        serverConfigs := []struct {
            name string
            cpes []string
        }{
            {
                "Apache on Ubuntu",
                []string{
                    "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
                    "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
                },
            },
            {
                "Nginx on RHEL",
                []string{
                    "cpe:2.3:o:redhat:enterprise_linux:8:*:*:*:*:*:*:*",
                    "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",
                },
            },
            {
                "IIS on Windows",
                []string{
                    "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                    "cpe:2.3:a:microsoft:internet_information_services:10.0:*:*:*:*:*:*:*",
                },
            },
        }
        
        fmt.Println("\nTesting server configurations:")
        for _, config := range serverConfigs {
            configCPEs := make([]*cpeskills.CPE, 0, len(config.cpes))
            for _, cpeStr := range config.cpes {
                cpeObj, _ := cpeskills.ParseCpe23(cpeStr)
                configCPEs = append(configCPEs, cpeObj)
            }
            
            applies := cpeskills.EvaluateApplicability(webServerExpr, configCPEs)
            
            status := "❌"
            if applies {
                status = "✅"
            }
            
            fmt.Printf("  %s %s\n", status, config.name)
        }
    }
    
    // Example 6: Expression Optimization
    fmt.Println("\n6. Expression Optimization:")
    
    // Original verbose expression
    verboseExpr := `
    (cpe:2.3:a:microsoft:office:2016:*:*:*:*:*:*:* OR
     cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:* OR
     cpe:2.3:a:microsoft:office:365:*:*:*:*:*:*:*)
    AND
    (cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR
     cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*)
    `
    
    // Optimized expression using wildcards
    optimizedExpr := `
    cpe:2.3:a:microsoft:office:*:*:*:*:*:*:*:*
    AND
    cpe:2.3:o:microsoft:windows:*:*:*:*:*:*:*:*
    `
    
    fmt.Printf("Verbose Expression:\n%s\n", verboseExpr)
    fmt.Printf("Optimized Expression:\n%s\n", optimizedExpr)
    
    verboseParsed, _ := cpeskills.ParseApplicabilityExpression(verboseExpr)
    optimizedParsed, _ := cpeskills.ParseApplicabilityExpression(optimizedExpr)
    
    // Test both expressions
    testCPEs := []*cpeskills.CPE{
        mustParse("cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*"),
        mustParse("cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*"),
    }
    
    verboseResult := cpeskills.EvaluateApplicability(verboseParsed, testCPEs)
    optimizedResult := cpeskills.EvaluateApplicability(optimizedParsed, testCPEs)
    
    fmt.Printf("Verbose result: %t\n", verboseResult)
    fmt.Printf("Optimized result: %t\n", optimizedResult)
    fmt.Printf("Results match: %t\n", verboseResult == optimizedResult)
}

func mustParse(cpeStr string) *cpeskills.CPE {
    cpeObj, err := cpeskills.ParseCpe23(cpeStr)
    if err != nil {
        panic(err)
    }
    return cpeObj
}
```

## Key Concepts

### 1. Logical Operators

- **AND**: All conditions must be true
- **OR**: At least one condition must be true  
- **NOT**: Condition must be false

### 2. Expression Structure

- **Simple**: Single CPE match
- **Compound**: Multiple CPEs with operators
- **Nested**: Complex hierarchical conditions

### 3. Use Cases

- **Vulnerability Applicability**: Define affected systems
- **Policy Compliance**: Specify required configurations
- **Asset Classification**: Group similar systems
- **Patch Management**: Identify update targets

## Best Practices

1. **Use Wildcards**: Simplify expressions with wildcards when appropriate
2. **Group Logically**: Group related conditions together
3. **Test Thoroughly**: Validate expressions against known systems
4. **Document Intent**: Comment complex expressions
5. **Optimize Performance**: Prefer simpler expressions when possible

## Next Steps

- Learn about [Advanced Matching](./advanced-matching.md) for complex scenarios
- Explore [CPE Sets](./sets.md) for bulk operations
- Check out [NVD Integration](./nvd-integration.md) for real-world applicability
