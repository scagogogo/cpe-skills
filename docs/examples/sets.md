# CPE Sets

This example demonstrates how to work with collections of CPE objects using the CPE Sets functionality for efficient bulk operations.

## Overview

CPE Sets provide a powerful way to manage collections of CPE objects, perform set operations (union, intersection, difference), and apply bulk transformations and filters.

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE Sets Examples ===")
    
    // Example 1: Creating CPE Sets
    fmt.Println("\n1. Creating CPE Sets:")
    
    // Create individual CPE objects
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
    }
    
    // Method 1: Create set from strings
    set1 := cpe.NewCPESetFromStrings(cpeStrings)
    fmt.Printf("Set 1 size: %d\n", set1.Size())
    
    // Method 2: Create empty set and add items
    set2 := cpe.NewCPESet()
    for _, cpeStr := range cpeStrings[:3] { // Add first 3 items
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("Failed to parse %s: %v", cpeStr, err)
            continue
        }
        set2.Add(cpeObj)
    }
    fmt.Printf("Set 2 size: %d\n", set2.Size())
    
    // Method 3: Create from slice of CPE objects
    cpeObjects := make([]*cpe.CPE, 0, len(cpeStrings))
    for _, cpeStr := range cpeStrings[2:] { // Add last 3 items
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            continue
        }
        cpeObjects = append(cpeObjects, cpeObj)
    }
    set3 := cpe.NewCPESetFromSlice(cpeObjects)
    fmt.Printf("Set 3 size: %d\n", set3.Size())
    
    // Example 2: Set Operations
    fmt.Println("\n2. Set Operations:")
    
    fmt.Println("Set 1 contents:")
    set1.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    fmt.Println("Set 2 contents:")
    set2.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    fmt.Println("Set 3 contents:")
    set3.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Union: All unique items from both sets
    unionSet := set2.Union(set3)
    fmt.Printf("\nUnion of Set 2 and Set 3 (size: %d):\n", unionSet.Size())
    unionSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Intersection: Items present in both sets
    intersectionSet := set1.Intersection(set2)
    fmt.Printf("\nIntersection of Set 1 and Set 2 (size: %d):\n", intersectionSet.Size())
    intersectionSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Difference: Items in first set but not in second
    differenceSet := set1.Difference(set2)
    fmt.Printf("\nDifference of Set 1 - Set 2 (size: %d):\n", differenceSet.Size())
    differenceSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Example 3: Filtering Sets
    fmt.Println("\n3. Filtering Sets:")
    
    // Create a larger set for filtering examples
    largeSetStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:edge:95.0.1020.44:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:mysql:8.0.26:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
        "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:h:cisco:catalyst_2960:*:*:*:*:*:*:*:*",
    }
    
    largeSet := cpe.NewCPESetFromStrings(largeSetStrings)
    fmt.Printf("Large set size: %d\n", largeSet.Size())
    
    // Filter by vendor
    microsoftCPEs := largeSet.FilterByVendor("microsoft")
    fmt.Printf("\nMicrosoft CPEs (size: %d):\n", microsoftCPEs.Size())
    microsoftCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Filter by part (applications only)
    applicationCPEs := largeSet.FilterByPart("a")
    fmt.Printf("\nApplication CPEs (size: %d):\n", applicationCPEs.Size())
    applicationCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Filter by product pattern
    apacheCPEs := largeSet.FilterByProduct("apache")
    fmt.Printf("\nApache CPEs (size: %d):\n", apacheCPEs.Size())
    apacheCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // Custom filter function
    customFilter := func(cpe *cpe.CPE) bool {
        // Filter for applications with version information
        return cpe.Part.ShortName == "a" && cpe.Version != "*" && cpe.Version != ""
    }
    
    versionedApps := largeSet.Filter(customFilter)
    fmt.Printf("\nVersioned Applications (size: %d):\n", versionedApps.Size())
    versionedApps.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s (v%s)\n", cpe.ProductName, cpe.Version)
    })
    
    // Example 4: Set Transformations
    fmt.Println("\n4. Set Transformations:")
    
    // Transform to extract vendor information
    vendors := largeSet.Map(func(cpe *cpe.CPE) string {
        return cpe.Vendor
    })
    
    uniqueVendors := removeDuplicateStrings(vendors)
    fmt.Printf("Unique vendors: %v\n", uniqueVendors)
    
    // Transform to create summary information
    summaries := largeSet.Map(func(cpe *cpe.CPE) string {
        return fmt.Sprintf("%s %s %s", cpe.Vendor, cpe.ProductName, cpe.Version)
    })
    
    fmt.Println("\nCPE Summaries:")
    for i, summary := range summaries {
        fmt.Printf("  %d. %s\n", i+1, summary)
    }
    
    // Example 5: Set Aggregation
    fmt.Println("\n5. Set Aggregation:")
    
    // Group by vendor
    vendorGroups := largeSet.GroupBy(func(cpe *cpe.CPE) string {
        return cpe.Vendor
    })
    
    fmt.Println("CPEs grouped by vendor:")
    for vendor, cpes := range vendorGroups {
        fmt.Printf("  %s (%d items):\n", vendor, len(cpes))
        for _, cpe := range cpes {
            fmt.Printf("    - %s\n", cpe.ProductName)
        }
    }
    
    // Group by part type
    partGroups := largeSet.GroupBy(func(cpe *cpe.CPE) string {
        return cpe.Part.LongName
    })
    
    fmt.Println("\nCPEs grouped by part type:")
    for partType, cpes := range partGroups {
        fmt.Printf("  %s: %d items\n", partType, len(cpes))
    }
    
    // Example 6: Set Statistics
    fmt.Println("\n6. Set Statistics:")
    
    stats := largeSet.GetStatistics()
    fmt.Printf("Set Statistics:\n")
    fmt.Printf("  Total CPEs: %d\n", stats.TotalCount)
    fmt.Printf("  Applications: %d\n", stats.ApplicationCount)
    fmt.Printf("  Operating Systems: %d\n", stats.OperatingSystemCount)
    fmt.Printf("  Hardware: %d\n", stats.HardwareCount)
    fmt.Printf("  Unique Vendors: %d\n", stats.UniqueVendors)
    fmt.Printf("  Unique Products: %d\n", stats.UniqueProducts)
    
    // Example 7: Set Persistence
    fmt.Println("\n7. Set Persistence:")
    
    // Save set to file
    filename := "cpe_set_export.json"
    err := largeSet.SaveToFile(filename)
    if err != nil {
        log.Printf("Failed to save set: %v", err)
    } else {
        fmt.Printf("Set saved to %s\n", filename)
    }
    
    // Load set from file
    loadedSet, err := cpe.LoadCPESetFromFile(filename)
    if err != nil {
        log.Printf("Failed to load set: %v", err)
    } else {
        fmt.Printf("Set loaded from %s (size: %d)\n", filename, loadedSet.Size())
        
        // Verify loaded set matches original
        if loadedSet.Size() == largeSet.Size() {
            fmt.Println("✅ Loaded set size matches original")
        } else {
            fmt.Println("❌ Loaded set size differs from original")
        }
    }
    
    // Example 8: Set Comparison
    fmt.Println("\n8. Set Comparison:")
    
    // Create two similar sets
    setA := cpe.NewCPESetFromStrings([]string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    })
    
    setB := cpe.NewCPESetFromStrings([]string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
    })
    
    fmt.Printf("Set A size: %d\n", setA.Size())
    fmt.Printf("Set B size: %d\n", setB.Size())
    
    // Check equality
    areEqual := setA.Equals(setB)
    fmt.Printf("Sets are equal: %t\n", areEqual)
    
    // Check if one is subset of another
    isSubset := setA.IsSubsetOf(setB)
    fmt.Printf("Set A is subset of Set B: %t\n", isSubset)
    
    // Find common elements
    common := setA.Intersection(setB)
    fmt.Printf("Common elements: %d\n", common.Size())
    
    // Find unique elements in each set
    uniqueA := setA.Difference(setB)
    uniqueB := setB.Difference(setA)
    
    fmt.Printf("Unique to Set A: %d\n", uniqueA.Size())
    fmt.Printf("Unique to Set B: %d\n", uniqueB.Size())
}

// Helper function to remove duplicate strings
func removeDuplicateStrings(slice []string) []string {
    seen := make(map[string]bool)
    result := []string{}
    
    for _, item := range slice {
        if !seen[item] {
            seen[item] = true
            result = append(result, item)
        }
    }
    
    return result
}
```

## Key Concepts

### 1. Set Creation

- **From Strings**: Parse CPE strings directly into a set
- **From Objects**: Create set from existing CPE objects
- **Empty Set**: Start with empty set and add items

### 2. Set Operations

- **Union**: Combine two sets (A ∪ B)
- **Intersection**: Common elements (A ∩ B)
- **Difference**: Elements in A but not B (A - B)

### 3. Filtering and Transformation

- **Filter**: Select subset based on criteria
- **Map**: Transform each element
- **GroupBy**: Organize elements by key

### 4. Set Analysis

- **Statistics**: Count elements by type
- **Comparison**: Check equality and subset relationships
- **Aggregation**: Summarize set contents

## Best Practices

1. **Use Sets for Bulk Operations**: More efficient than individual operations
2. **Filter Early**: Apply filters to reduce set size before expensive operations
3. **Cache Results**: Store frequently used filtered sets
4. **Validate Input**: Check CPE validity before adding to sets
5. **Monitor Memory**: Large sets can consume significant memory

## Performance Tips

1. **Batch Operations**: Group multiple operations together
2. **Use Appropriate Data Structures**: Sets are optimized for uniqueness
3. **Parallel Processing**: Use goroutines for independent set operations
4. **Lazy Evaluation**: Defer expensive operations until needed

## Next Steps

- Learn about [Advanced Matching](./advanced-matching.md) with sets
- Explore [Storage](./storage.md) for persisting large sets
- Check out [NVD Integration](./nvd-integration.md) for real-world datasets
