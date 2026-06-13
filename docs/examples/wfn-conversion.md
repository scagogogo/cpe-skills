# WFN Conversion

This example demonstrates how to work with Well-Formed Names (WFN), the internal representation format used by the CPE library for processing and matching.

## Overview

Well-Formed Names (WFN) are the canonical internal representation of CPE names. They provide a standardized way to represent CPE components that makes matching and comparison operations more efficient and reliable.

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    fmt.Println("=== WFN Conversion Examples ===")
    
    // Example 1: CPE to WFN Conversion
    fmt.Println("\n1. CPE to WFN Conversion:")
    
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:/a:oracle:java:1.8.0_291",
        "cpe:2.3:o:linux:kernel:5.4.0:*:*:*:*:*:*:*",
    }
    
    for i, cpeStr := range cpeStrings {
        fmt.Printf("\nExample %d: %s\n", i+1, cpeStr)
        
        // Parse CPE
        cpeObj, err := cpe.ParseCPE(cpeStr)
        if err != nil {
            log.Printf("Failed to parse CPE: %v", err)
            continue
        }
        
        // Convert to WFN
        wfn, err := cpe.CPEToWFN(cpeObj)
        if err != nil {
            log.Printf("Failed to convert to WFN: %v", err)
            continue
        }
        
        fmt.Printf("  Original CPE: %s\n", cpeStr)
        fmt.Printf("  WFN Format:   %s\n", wfn.String())
        fmt.Printf("  Part:         %s\n", wfn.Part)
        fmt.Printf("  Vendor:       %s\n", wfn.Vendor)
        fmt.Printf("  Product:      %s\n", wfn.Product)
        fmt.Printf("  Version:      %s\n", wfn.Version)
    }
    
    // Example 2: WFN to CPE Conversion
    fmt.Println("\n2. WFN to CPE Conversion:")
    
    // Create WFN manually
    wfn := &cpe.WFN{
        Part:           "a",
        Vendor:         "adobe",
        Product:        "reader",
        Version:        "2021.001.20150",
        Update:         cpe.WFNAny,
        Edition:        cpe.WFNAny,
        Language:       cpe.WFNAny,
        SoftwareEdition: cpe.WFNAny,
        TargetSoftware: cpe.WFNAny,
        TargetHardware: cpe.WFNAny,
        Other:          cpe.WFNAny,
    }
    
    fmt.Printf("WFN: %s\n", wfn.String())
    
    // Convert to CPE 2.3
    cpe23, err := cpe.WFNToCPE23(wfn)
    if err != nil {
        log.Printf("Failed to convert to CPE 2.3: %v", err)
    } else {
        fmt.Printf("CPE 2.3: %s\n", cpe23)
    }
    
    // Convert to CPE 2.2
    cpe22, err := cpe.WFNToCPE22(wfn)
    if err != nil {
        log.Printf("Failed to convert to CPE 2.2: %v", err)
    } else {
        fmt.Printf("CPE 2.2: %s\n", cpe22)
    }
    
    // Example 3: WFN Attribute Values
    fmt.Println("\n3. WFN Attribute Values:")
    
    // Demonstrate different WFN attribute values
    examples := []struct {
        name  string
        value string
        desc  string
    }{
        {"ANY", cpe.WFNAny, "Matches any value"},
        {"NA", cpe.WFNNotApplicable, "Not applicable"},
        {"Literal", "windows", "Literal string value"},
        {"Quoted", cpe.QuoteWFNValue("special~chars"), "Quoted special characters"},
    }
    
    for _, example := range examples {
        fmt.Printf("  %s: '%s' - %s\n", example.name, example.value, example.desc)
    }
    
    // Example 4: WFN Matching
    fmt.Println("\n4. WFN Matching:")
    
    // Create source and target WFNs
    sourceWFN := &cpe.WFN{
        Part:    "a",
        Vendor:  "microsoft",
        Product: cpe.WFNAny, // Any product
        Version: cpe.WFNAny, // Any version
    }
    
    targetWFNs := []*cpe.WFN{
        {Part: "a", Vendor: "microsoft", Product: "windows", Version: "10"},
        {Part: "a", Vendor: "microsoft", Product: "office", Version: "2019"},
        {Part: "a", Vendor: "oracle", Product: "java", Version: "11"},
        {Part: "o", Vendor: "microsoft", Product: "windows", Version: "10"},
    }
    
    fmt.Printf("Source WFN: %s\n", sourceWFN.String())
    fmt.Println("Matching against targets:")
    
    for i, targetWFN := range targetWFNs {
        match := cpe.MatchWFN(sourceWFN, targetWFN)
        status := "❌"
        if match {
            status = "✅"
        }
        fmt.Printf("  %s Target %d: %s\n", status, i+1, targetWFN.String())
    }
    
    // Example 5: WFN Validation
    fmt.Println("\n5. WFN Validation:")
    
    validationTests := []struct {
        wfn   *cpe.WFN
        desc  string
        valid bool
    }{
        {
            &cpe.WFN{Part: "a", Vendor: "microsoft", Product: "windows"},
            "Valid application WFN",
            true,
        },
        {
            &cpe.WFN{Part: "x", Vendor: "microsoft", Product: "windows"},
            "Invalid part value",
            false,
        },
        {
            &cpe.WFN{Part: "a", Vendor: "", Product: "windows"},
            "Empty vendor",
            false,
        },
        {
            &cpe.WFN{Part: "a", Vendor: "microsoft", Product: ""},
            "Empty product",
            false,
        },
    }
    
    for i, test := range validationTests {
        err := cpe.ValidateWFN(test.wfn)
        isValid := err == nil
        
        status := "❌"
        if isValid == test.valid {
            status = "✅"
        }
        
        fmt.Printf("  %s Test %d: %s\n", status, i+1, test.desc)
        fmt.Printf("    WFN: %s\n", test.wfn.String())
        if err != nil {
            fmt.Printf("    Error: %v\n", err)
        }
    }
    
    // Example 6: WFN Normalization
    fmt.Println("\n6. WFN Normalization:")
    
    unnormalizedWFN := &cpe.WFN{
        Part:    "A", // Should be lowercase
        Vendor:  "Microsoft", // Should be lowercase
        Product: "Windows~10", // Special characters
        Version: "10.0.19041.1234",
    }
    
    fmt.Printf("Before normalization: %s\n", unnormalizedWFN.String())
    
    normalizedWFN := cpe.NormalizeWFN(unnormalizedWFN)
    fmt.Printf("After normalization:  %s\n", normalizedWFN.String())
    
    // Example 7: WFN Comparison
    fmt.Println("\n7. WFN Comparison:")
    
    wfn1 := &cpe.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.0",
    }
    wfn2 := &cpe.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.1",
    }
    wfn3 := &cpe.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.0",
    }
    
    fmt.Printf("WFN1: %s\n", wfn1.String())
    fmt.Printf("WFN2: %s\n", wfn2.String())
    fmt.Printf("WFN3: %s\n", wfn3.String())
    
    fmt.Printf("WFN1 == WFN2: %t\n", cpe.CompareWFN(wfn1, wfn2) == 0)
    fmt.Printf("WFN1 == WFN3: %t\n", cpe.CompareWFN(wfn1, wfn3) == 0)
    fmt.Printf("WFN1 < WFN2:  %t\n", cpe.CompareWFN(wfn1, wfn2) < 0)
}
```

## Expected Output

```
=== WFN Conversion Examples ===

1. CPE to WFN Conversion:

Example 1: cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*
  Original CPE: cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*
  WFN Format:   wfn:[part="a",vendor="microsoft",product="windows",version="10",update=ANY,edition=ANY,language=ANY,sw_edition=ANY,target_sw=ANY,target_hw=ANY,other=ANY]
  Part:         a
  Vendor:       microsoft
  Product:      windows
  Version:      10

Example 2: cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*
  Original CPE: cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*
  WFN Format:   wfn:[part="a",vendor="apache",product="tomcat",version="9.0.0",update=ANY,edition=ANY,language=ANY,sw_edition=ANY,target_sw=ANY,target_hw=ANY,other=ANY]
  Part:         a
  Vendor:       apache
  Product:      tomcat
  Version:      9.0.0

2. WFN to CPE Conversion:
WFN: wfn:[part="a",vendor="adobe",product="reader",version="2021.001.20150",update=ANY,edition=ANY,language=ANY,sw_edition=ANY,target_sw=ANY,target_hw=ANY,other=ANY]
CPE 2.3: cpe:2.3:a:adobe:reader:2021.001.20150:*:*:*:*:*:*:*
CPE 2.2: cpe:/a:adobe:reader:2021.001.20150

3. WFN Attribute Values:
  ANY: '*' - Matches any value
  NA: '-' - Not applicable
  Literal: 'windows' - Literal string value
  Quoted: 'special\~chars' - Quoted special characters

4. WFN Matching:
Source WFN: wfn:[part="a",vendor="microsoft",product=ANY,version=ANY]
Matching against targets:
  ✅ Target 1: wfn:[part="a",vendor="microsoft",product="windows",version="10"]
  ✅ Target 2: wfn:[part="a",vendor="microsoft",product="office",version="2019"]
  ❌ Target 3: wfn:[part="a",vendor="oracle",product="java",version="11"]
  ❌ Target 4: wfn:[part="o",vendor="microsoft",product="windows",version="10"]

5. WFN Validation:
  ✅ Test 1: Valid application WFN
    WFN: wfn:[part="a",vendor="microsoft",product="windows"]
  ✅ Test 2: Invalid part value
    WFN: wfn:[part="x",vendor="microsoft",product="windows"]
    Error: invalid part value: x
  ✅ Test 3: Empty vendor
    WFN: wfn:[part="a",vendor="",product="windows"]
    Error: vendor cannot be empty
  ✅ Test 4: Empty product
    WFN: wfn:[part="a",vendor="microsoft",product=""]
    Error: product cannot be empty

6. WFN Normalization:
Before normalization: wfn:[part="A",vendor="Microsoft",product="Windows~10",version="10.0.19041.1234"]
After normalization:  wfn:[part="a",vendor="microsoft",product="windows\~10",version="10.0.19041.1234"]

7. WFN Comparison:
WFN1: wfn:[part="a",vendor="apache",product="tomcat",version="9.0.0"]
WFN2: wfn:[part="a",vendor="apache",product="tomcat",version="9.0.1"]
WFN3: wfn:[part="a",vendor="apache",product="tomcat",version="9.0.0"]
WFN1 == WFN2: false
WFN1 == WFN3: true
WFN1 < WFN2:  true
```

## Key Concepts

### 1. WFN Structure

A WFN consists of 11 attributes:
- **part**: Component type (a, h, o)
- **vendor**: Vendor name
- **product**: Product name
- **version**: Version string
- **update**: Update identifier
- **edition**: Edition information
- **language**: Language code
- **sw_edition**: Software edition
- **target_sw**: Target software
- **target_hw**: Target hardware
- **other**: Other information

### 2. Special Values

- **ANY (*)**: Matches any value
- **NA (-)**: Not applicable/undefined
- **Literal**: Exact string match

### 3. WFN Benefits

- **Canonical Form**: Standardized representation
- **Efficient Matching**: Optimized for comparison operations
- **Validation**: Built-in validation rules
- **Normalization**: Consistent formatting

## Best Practices

1. **Use WFN for Internal Processing**: Convert CPE strings to WFN for operations
2. **Validate WFNs**: Always validate WFN objects before use
3. **Normalize Input**: Normalize WFNs for consistent comparison
4. **Handle Special Values**: Properly handle ANY and NA values
5. **Convert Back**: Convert WFN back to CPE format for output

## Next Steps

- Learn about [Advanced Matching](./advanced-matching.md) using WFN
- Explore [CPE Sets](./sets.md) for bulk WFN operations
- Check out [Storage](./storage.md) for persisting WFN data
