# Advanced Matching

This example demonstrates sophisticated CPE matching techniques including fuzzy matching, pattern matching, and complex matching algorithms.

## Overview

Advanced matching goes beyond simple string comparison to provide intelligent matching capabilities that can handle variations, patterns, and complex matching scenarios commonly encountered in real-world applications.

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== Advanced CPE Matching Examples ===")
    
    // Example 1: Fuzzy Matching
    fmt.Println("\n1. Fuzzy Matching:")
    
    // Target CPE with slight variations
    targetCPE, _ := cpe.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    
    // Test CPEs with variations
    testCPEs := []string{
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",     // Exact match
        "cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*",     // Version difference
        "cpe:2.3:a:apache:tomcat_server:9.0.0:*:*:*:*:*:*:*", // Product name variation
        "cpe:2.3:a:apache_software:tomcat:9.0.0:*:*:*:*:*:*:*", // Vendor variation
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",     // Completely different
    }
    
    fmt.Printf("Target: %s\n", targetCPE.GetURI())
    fmt.Println("Fuzzy matching results:")
    
    for i, testCPEStr := range testCPEs {
        testCPE, err := cpe.ParseCpe23(testCPEStr)
        if err != nil {
            log.Printf("Failed to parse %s: %v", testCPEStr, err)
            continue
        }
        
        // Calculate fuzzy match score (0.0 to 1.0)
        score := cpe.FuzzyMatch(targetCPE, testCPE)
        
        var status string
        switch {
        case score >= 0.9:
            status = "🟢 Excellent"
        case score >= 0.7:
            status = "🟡 Good"
        case score >= 0.5:
            status = "🟠 Fair"
        default:
            status = "🔴 Poor"
        }
        
        fmt.Printf("  %d. %s (Score: %.2f) %s\n", i+1, testCPE.ProductName, score, status)
    }
    
    // Example 2: Pattern Matching
    fmt.Println("\n2. Pattern Matching:")
    
    // Define patterns with wildcards and regex
    patterns := []struct {
        name    string
        pattern string
        desc    string
    }{
        {
            "Microsoft Products",
            "cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*",
            "Any Microsoft application",
        },
        {
            "Java Versions 8.x",
            "cpe:2.3:a:oracle:java:8.*:*:*:*:*:*:*:*",
            "Oracle Java version 8.x",
        },
        {
            "Windows Operating Systems",
            "cpe:2.3:o:microsoft:windows:*:*:*:*:*:*:*:*",
            "Any Windows OS version",
        },
        {
            "Apache Web Technologies",
            "cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*",
            "Any Apache application",
        },
    }
    
    // Test CPEs against patterns
    testTargets := []string{
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:8.0.291:*:*:*:*:*:*:*",
        "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
        "cpe:2.3:a:mozilla:firefox:95.0:*:*:*:*:*:*:*",
    }
    
    fmt.Println("Pattern matching results:")
    for _, target := range testTargets {
        targetCPE, _ := cpe.ParseCpe23(target)
        fmt.Printf("\nTarget: %s\n", target)
        
        for _, pattern := range patterns {
            patternCPE, _ := cpe.ParseCpe23(pattern.pattern)
            matches := cpe.MatchPattern(targetCPE, patternCPE)
            
            status := "❌"
            if matches {
                status = "✅"
            }
            
            fmt.Printf("  %s %s: %s\n", status, pattern.name, pattern.desc)
        }
    }
    
    // Example 3: Semantic Matching
    fmt.Println("\n3. Semantic Matching:")
    
    // Define semantic equivalences
    semanticRules := []struct {
        primary    string
        equivalent string
        reason     string
    }{
        {
            "cpe:2.3:a:microsoft:internet_explorer:*:*:*:*:*:*:*:*",
            "cpe:2.3:a:microsoft:ie:*:*:*:*:*:*:*:*",
            "IE is common abbreviation for Internet Explorer",
        },
        {
            "cpe:2.3:a:apache:http_server:*:*:*:*:*:*:*:*",
            "cpe:2.3:a:apache:httpd:*:*:*:*:*:*:*:*",
            "httpd is the daemon name for Apache HTTP Server",
        },
        {
            "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
            "cpe:2.3:o:microsoft:win10:*:*:*:*:*:*:*:*",
            "Win10 is common abbreviation for Windows 10",
        },
    }
    
    fmt.Println("Semantic matching examples:")
    for i, rule := range semanticRules {
        primaryCPE, _ := cpe.ParseCpe23(rule.primary)
        equivalentCPE, _ := cpe.ParseCpe23(rule.equivalent)
        
        // Test semantic matching
        matches := cpe.SemanticMatch(primaryCPE, equivalentCPE)
        
        status := "❌"
        if matches {
            status = "✅"
        }
        
        fmt.Printf("  %d. %s\n", i+1, status)
        fmt.Printf("     Primary: %s\n", primaryCPE.ProductName)
        fmt.Printf("     Equivalent: %s\n", equivalentCPE.ProductName)
        fmt.Printf("     Reason: %s\n", rule.reason)
    }
    
    // Example 4: Version Range Matching
    fmt.Println("\n4. Version Range Matching:")
    
    // Define version ranges
    versionRanges := []struct {
        name     string
        minVer   string
        maxVer   string
        inclusive bool
    }{
        {"Java 8 Updates", "8.0.0", "8.0.999", true},
        {"Tomcat 9.0.x", "9.0.0", "9.0.999", true},
        {"Windows 10 Builds", "10.0.10240", "10.0.19999", true},
        {"Office 2019 Versions", "16.0.0", "16.9.999", true},
    }
    
    // Test versions against ranges
    testVersions := []struct {
        cpe     string
        version string
    }{
        {"cpe:2.3:a:oracle:java:8.0.291:*:*:*:*:*:*:*", "8.0.291"},
        {"cpe:2.3:a:apache:tomcat:9.0.45:*:*:*:*:*:*:*", "9.0.45"},
        {"cpe:2.3:o:microsoft:windows:10.0.19041:*:*:*:*:*:*:*", "10.0.19041"},
        {"cpe:2.3:a:microsoft:office:16.0.13901:*:*:*:*:*:*:*", "16.0.13901"},
        {"cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*", "11.0.12"},
    }
    
    fmt.Println("Version range matching:")
    for _, test := range testVersions {
        testCPE, _ := cpe.ParseCpe23(test.cpe)
        fmt.Printf("\nTesting: %s %s (v%s)\n", testCPE.Vendor, testCPE.ProductName, test.version)
        
        for _, vrange := range versionRanges {
            inRange := cpe.IsVersionInRange(test.version, vrange.minVer, vrange.maxVer)
            
            status := "❌"
            if inRange {
                status = "✅"
            }
            
            fmt.Printf("  %s %s (%s - %s)\n", status, vrange.name, vrange.minVer, vrange.maxVer)
        }
    }
    
    // Example 5: Weighted Matching
    fmt.Println("\n5. Weighted Matching:")
    
    // Define matching weights for different components
    weights := cpe.MatchWeights{
        Part:    0.1,  // Part is less important
        Vendor:  0.3,  // Vendor is important
        Product: 0.4,  // Product is most important
        Version: 0.2,  // Version is moderately important
    }
    
    referenceCPE, _ := cpe.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    
    candidateCPEs := []string{
        "cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*",     // Same product, different version
        "cpe:2.3:a:apache:tomcat_server:9.0.0:*:*:*:*:*:*:*", // Similar product name
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*", // Same vendor, different product
        "cpe:2.3:o:apache:tomcat:9.0.0:*:*:*:*:*:*:*",     // Different part type
    }
    
    fmt.Printf("Reference: %s\n", referenceCPE.GetURI())
    fmt.Println("Weighted matching scores:")
    
    for i, candidateStr := range candidateCPEs {
        candidateCPE, _ := cpe.ParseCpe23(candidateStr)
        score := cpe.WeightedMatch(referenceCPE, candidateCPE, weights)
        
        fmt.Printf("  %d. Score: %.3f - %s\n", i+1, score, candidateCPE.GetURI())
    }
    
    // Example 6: Contextual Matching
    fmt.Println("\n6. Contextual Matching:")
    
    // Define context for matching
    context := cpe.MatchContext{
        Environment: "production",
        Platform:    "linux",
        Purpose:     "web_server",
    }
    
    // CPEs with context information
    contextualCPEs := []struct {
        cpe     string
        context cpe.MatchContext
    }{
        {
            "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
            cpe.MatchContext{Environment: "production", Platform: "linux", Purpose: "web_server"},
        },
        {
            "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
            cpe.MatchContext{Environment: "development", Platform: "linux", Purpose: "web_server"},
        },
        {
            "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",
            cpe.MatchContext{Environment: "production", Platform: "linux", Purpose: "web_server"},
        },
        {
            "cpe:2.3:a:microsoft:iis:10.0:*:*:*:*:*:*:*",
            cpe.MatchContext{Environment: "production", Platform: "windows", Purpose: "web_server"},
        },
    }
    
    fmt.Printf("Target context: %+v\n", context)
    fmt.Println("Contextual matching results:")
    
    for i, item := range contextualCPEs {
        itemCPE, _ := cpe.ParseCpe23(item.cpe)
        contextMatch := cpe.ContextualMatch(context, item.context)
        
        var status string
        switch {
        case contextMatch >= 0.9:
            status = "🟢 Perfect"
        case contextMatch >= 0.7:
            status = "🟡 Good"
        case contextMatch >= 0.5:
            status = "🟠 Partial"
        default:
            status = "🔴 Poor"
        }
        
        fmt.Printf("  %d. %s %s (Score: %.2f) %s\n", 
            i+1, itemCPE.Vendor, itemCPE.ProductName, contextMatch, status)
        fmt.Printf("     Context: %+v\n", item.context)
    }
    
    // Example 7: Machine Learning Enhanced Matching
    fmt.Println("\n7. ML-Enhanced Matching:")
    
    // Simulate ML model predictions
    mlModel := cpe.NewMLMatchingModel()
    
    // Train with sample data (in real implementation, this would use actual training data)
    trainingPairs := []cpe.MatchingPair{
        {
            CPE1: mustParseCPE("cpe:2.3:a:apache:tomcat:8.5.0:*:*:*:*:*:*:*"),
            CPE2: mustParseCPE("cpe:2.3:a:apache:tomcat:8.5.1:*:*:*:*:*:*:*"),
            Match: true,
            Score: 0.95,
        },
        {
            CPE1: mustParseCPE("cpe:2.3:a:apache:tomcat:8.5.0:*:*:*:*:*:*:*"),
            CPE2: mustParseCPE("cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*"),
            Match: false,
            Score: 0.1,
        },
    }
    
    err := mlModel.Train(trainingPairs)
    if err != nil {
        log.Printf("Failed to train ML model: %v", err)
    } else {
        fmt.Println("ML model trained successfully")
        
        // Test ML predictions
        testPairs := []struct {
            cpe1 string
            cpe2 string
        }{
            {
                "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
                "cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*",
            },
            {
                "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:word:2019:*:*:*:*:*:*:*",
            },
        }
        
        fmt.Println("ML matching predictions:")
        for i, pair := range testPairs {
            cpe1, _ := cpe.ParseCpe23(pair.cpe1)
            cpe2, _ := cpe.ParseCpe23(pair.cpe2)
            
            prediction := mlModel.Predict(cpe1, cpe2)
            
            fmt.Printf("  %d. Score: %.3f\n", i+1, prediction.Score)
            fmt.Printf("     CPE1: %s\n", cpe1.ProductName)
            fmt.Printf("     CPE2: %s\n", cpe2.ProductName)
            fmt.Printf("     Confidence: %.2f\n", prediction.Confidence)
        }
    }
}

func mustParseCPE(cpeStr string) *cpe.CPE {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
    if err != nil {
        panic(err)
    }
    return cpeObj
}
```

## Key Concepts

### 1. Matching Types

- **Exact**: Perfect string match
- **Fuzzy**: Similarity-based matching
- **Pattern**: Wildcard and regex matching
- **Semantic**: Meaning-based matching
- **Contextual**: Environment-aware matching

### 2. Scoring Systems

- **Binary**: Match/no-match
- **Similarity**: 0.0 to 1.0 score
- **Weighted**: Component-based scoring
- **Confidence**: Prediction certainty

### 3. Advanced Techniques

- **Machine Learning**: Trained models for prediction
- **Context Awareness**: Environment-specific matching
- **Version Ranges**: Flexible version matching
- **Semantic Equivalence**: Alias and abbreviation handling

## Best Practices

1. **Choose Appropriate Method**: Select matching technique based on use case
2. **Tune Thresholds**: Adjust score thresholds for your domain
3. **Validate Results**: Test matching accuracy with known datasets
4. **Consider Performance**: Balance accuracy with computational cost
5. **Handle Edge Cases**: Account for unusual CPE formats and values

## Performance Considerations

1. **Caching**: Cache expensive matching computations
2. **Indexing**: Use appropriate data structures for large datasets
3. **Parallel Processing**: Distribute matching across multiple cores
4. **Early Termination**: Stop processing when confidence is sufficient

## Next Steps

- Learn about [NVD Integration](./nvd-integration.md) for real-world matching
- Explore [CVE Mapping](./cve-mapping.md) for vulnerability correlation
- Check out [Storage](./storage.md) for persisting matching results
