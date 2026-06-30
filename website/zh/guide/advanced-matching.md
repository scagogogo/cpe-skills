# 高级匹配

本示例演示CPE库中的高级匹配技术，包括模糊匹配、语义匹配和复杂的匹配策略。

## 概述

高级匹配超越了基本的字符串比较，提供了智能匹配算法，可以处理版本范围、同义词、模糊匹配和复杂的匹配条件。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "strings"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE高级匹配示例 ===")
    
    // 示例1：模糊匹配
    fmt.Println("\n1. 模糊匹配:")
    
    // 创建目标CPE
    targetCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    
    // 创建候选CPE（包含一些变体）
    candidateCPEs := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",      // 精确匹配
        "cpe:2.3:a:microsoft:win:10:*:*:*:*:*:*:*",         // 产品名缩写
        "cpe:2.3:a:microsoft:windows:10.0:*:*:*:*:*:*:*",   // 版本变体
        "cpe:2.3:a:microsoft:windows_10:10:*:*:*:*:*:*:*",  // 产品名变体
        "cpe:2.3:a:ms:windows:10:*:*:*:*:*:*:*",            // 供应商缩写
        "cpe:2.3:a:oracle:java:11:*:*:*:*:*:*:*",           // 完全不同
    }
    
    fmt.Printf("目标CPE: %s\n", targetCPE.GetURI())
    fmt.Println("模糊匹配结果:")
    
    for i, candidateStr := range candidateCPEs {
        candidateCPE, _ := cpeskills.ParseCpe23(candidateStr)
        
        // 计算相似度分数
        similarity := calculateSimilarity(targetCPE, candidateCPE)
        
        status := "❌"
        if similarity >= 0.7 { // 70%阈值
            status = "✅"
        }
        
        fmt.Printf("  %s %d. 相似度: %.2f - %s\n", 
            status, i+1, similarity, candidateStr)
    }
    
    // 示例2：版本范围匹配
    fmt.Println("\n2. 版本范围匹配:")
    
    // 定义漏洞影响的版本范围
    vulnerableRanges := []struct {
        product    string
        minVersion string
        maxVersion string
        cveID      string
    }{
        {"tomcat", "8.5.0", "8.5.4", "CVE-2021-25122"},
        {"java", "1.8.0", "1.8.0_291", "CVE-2021-2163"},
        {"nginx", "1.0.0", "1.18.0", "CVE-2021-23017"},
    }
    
    // 测试系统中的软件
    systemSoftware := []string{
        "cpe:2.3:a:apache:tomcat:8.5.3:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:8.5.5:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:1.8.0_281:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:1.8.0_301:*:*:*:*:*:*:*",
        "cpe:2.3:a:nginx:nginx:1.16.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:nginx:nginx:1.20.0:*:*:*:*:*:*:*",
    }
    
    fmt.Println("漏洞范围匹配:")
    for _, softwareStr := range systemSoftware {
        softwareCPE, _ := cpeskills.ParseCpe23(softwareStr)
        fmt.Printf("\n检查: %s %s\n", softwareCPE.ProductName, softwareCPE.Version)
        
        vulnerabilityFound := false
        for _, vulnRange := range vulnerableRanges {
            if softwareCPE.ProductName == vulnRange.product {
                if isVersionInRange(softwareCPE.Version, vulnRange.minVersion, vulnRange.maxVersion) {
                    fmt.Printf("  ⚠️  易受攻击: %s (版本 %s - %s)\n", 
                        vulnRange.cveID, vulnRange.minVersion, vulnRange.maxVersion)
                    vulnerabilityFound = true
                }
            }
        }
        
        if !vulnerabilityFound {
            fmt.Printf("  ✅ 未发现已知漏洞\n")
        }
    }
    
    // 示例3：语义匹配
    fmt.Println("\n3. 语义匹配:")
    
    // 定义同义词映射
    synonyms := map[string][]string{
        "microsoft": {"ms", "msft"},
        "windows": {"win", "windows_nt"},
        "internet_explorer": {"ie", "iexplore"},
        "apache": {"apache_software_foundation", "asf"},
        "tomcat": {"apache_tomcat", "catalina"},
    }
    
    // 测试语义匹配
    semanticTests := []struct {
        pattern string
        target  string
    }{
        {"microsoft", "ms"},
        {"windows", "win"},
        {"internet_explorer", "ie"},
        {"apache", "asf"},
        {"tomcat", "catalina"},
    }
    
    fmt.Println("语义匹配测试:")
    for i, test := range semanticTests {
        matches := semanticMatch(test.pattern, test.target, synonyms)
        
        status := "❌"
        if matches {
            status = "✅"
        }
        
        fmt.Printf("  %s %d. '%s' 匹配 '%s': %t\n", 
            status, i+1, test.pattern, test.target, matches)
    }
    
    // 示例4：复杂匹配条件
    fmt.Println("\n4. 复杂匹配条件:")
    
    // 定义复杂的匹配规则
    type MatchRule struct {
        Name        string
        Description string
        Condition   func(*cpeskills.CPE) bool
    }
    
    rules := []MatchRule{
        {
            "Web服务器",
            "Apache HTTP Server或Nginx",
            func(c *cpeskills.CPE) bool {
                return (c.Vendor == "apache" && c.ProductName == "http_server") ||
                       (c.Vendor == "nginx" && c.ProductName == "nginx")
            },
        },
        {
            "Microsoft产品",
            "任何Microsoft产品",
            func(c *cpeskills.CPE) bool {
                return c.Vendor == "microsoft"
            },
        },
        {
            "过时的Java",
            "Java版本低于11",
            func(c *cpeskills.CPE) bool {
                if c.Vendor == "oracle" && c.ProductName == "java" {
                    return isVersionLessThan(c.Version, "11.0.0")
                }
                return false
            },
        },
        {
            "关键基础设施",
            "操作系统或网络设备",
            func(c *cpeskills.CPE) bool {
                return c.Part.ShortName == "o" || c.Part.ShortName == "h"
            },
        },
    }
    
    // 测试软件清单
    inventory := []string{
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
        "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:1.8.0_291:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:h:cisco:catalyst_2960:*:*:*:*:*:*:*:*",
    }
    
    fmt.Println("复杂匹配规则应用:")
    for _, rule := range rules {
        fmt.Printf("\n规则: %s (%s)\n", rule.Name, rule.Description)
        
        matchCount := 0
        for _, itemStr := range inventory {
            itemCPE, _ := cpeskills.ParseCpe23(itemStr)
            if rule.Condition(itemCPE) {
                fmt.Printf("  ✅ %s %s %s\n", 
                    itemCPE.Vendor, itemCPE.ProductName, itemCPE.Version)
                matchCount++
            }
        }
        
        fmt.Printf("  匹配项: %d/%d\n", matchCount, len(inventory))
    }
    
    // 示例5：基于权重的匹配
    fmt.Println("\n5. 基于权重的匹配:")
    
    // 定义字段权重
    fieldWeights := map[string]float64{
        "vendor":  0.3,
        "product": 0.4,
        "version": 0.2,
        "part":    0.1,
    }
    
    referenceCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    
    testCPEs := []string{
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",   // 完全匹配
        "cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*",   // 版本不同
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*", // 产品不同
        "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",    // 供应商和产品不同
        "cpe:2.3:o:apache:tomcat:9.0.0:*:*:*:*:*:*:*",   // 部件类型不同
    }
    
    fmt.Printf("参考CPE: %s\n", referenceCPE.GetURI())
    fmt.Println("加权匹配分数:")
    
    for i, testStr := range testCPEs {
        testCPE, _ := cpeskills.ParseCpe23(testStr)
        score := calculateWeightedScore(referenceCPE, testCPE, fieldWeights)
        
        fmt.Printf("  %d. 分数: %.3f - %s\n", i+1, score, testStr)
    }
    
    // 示例6：上下文感知匹配
    fmt.Println("\n6. 上下文感知匹配:")
    
    // 定义不同的上下文
    contexts := map[string]func(*cpeskills.CPE, *cpeskills.CPE) bool{
        "安全扫描": func(pattern, target *cpeskills.CPE) bool {
            // 在安全上下文中，版本必须精确匹配
            return pattern.Vendor == target.Vendor &&
                   pattern.ProductName == target.ProductName &&
                   pattern.Version == target.Version
        },
        "资产清单": func(pattern, target *cpeskills.CPE) bool {
            // 在清单上下文中，版本可以是通配符
            return pattern.Vendor == target.Vendor &&
                   pattern.ProductName == target.ProductName &&
                   (pattern.Version == "*" || pattern.Version == target.Version)
        },
        "兼容性检查": func(pattern, target *cpeskills.CPE) bool {
            // 在兼容性上下文中，允许次版本差异
            if pattern.Vendor != target.Vendor || pattern.ProductName != target.ProductName {
                return false
            }
            return isVersionCompatible(pattern.Version, target.Version)
        },
    }
    
    patternCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.*:*:*:*:*:*:*:*")
    targetCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*")
    
    fmt.Printf("模式: %s\n", patternCPE.GetURI())
    fmt.Printf("目标: %s\n", targetCPE.GetURI())
    
    for contextName, matcher := range contexts {
        matches := matcher(patternCPE, targetCPE)
        
        status := "❌"
        if matches {
            status = "✅"
        }
        
        fmt.Printf("  %s %s上下文: %t\n", status, contextName, matches)
    }
}

// 辅助函数：计算CPE相似度
func calculateSimilarity(cpe1, cpe2 *cpeskills.CPE) float64 {
    var score float64
    
    // 供应商匹配 (权重: 30%)
    if cpe1.Vendor == cpe2.Vendor {
        score += 0.3
    } else if strings.Contains(cpe1.Vendor, cpe2.Vendor) || strings.Contains(cpe2.Vendor, cpe1.Vendor) {
        score += 0.15
    }
    
    // 产品匹配 (权重: 40%)
    if cpe1.ProductName == cpe2.ProductName {
        score += 0.4
    } else if strings.Contains(cpe1.ProductName, cpe2.ProductName) || strings.Contains(cpe2.ProductName, cpe1.ProductName) {
        score += 0.2
    }
    
    // 版本匹配 (权重: 20%)
    if cpe1.Version == cpe2.Version {
        score += 0.2
    } else if strings.HasPrefix(cpe1.Version, cpe2.Version) || strings.HasPrefix(cpe2.Version, cpe1.Version) {
        score += 0.1
    }
    
    // 部件匹配 (权重: 10%)
    if cpe1.Part.ShortName == cpe2.Part.ShortName {
        score += 0.1
    }
    
    return score
}

// 辅助函数：语义匹配
func semanticMatch(pattern, target string, synonyms map[string][]string) bool {
    if pattern == target {
        return true
    }
    
    // 检查同义词
    if syns, exists := synonyms[pattern]; exists {
        for _, syn := range syns {
            if syn == target {
                return true
            }
        }
    }
    
    // 反向检查
    if syns, exists := synonyms[target]; exists {
        for _, syn := range syns {
            if syn == pattern {
                return true
            }
        }
    }
    
    return false
}

// 辅助函数：版本范围检查
func isVersionInRange(version, minVersion, maxVersion string) bool {
    return compareVersions(version, minVersion) >= 0 && compareVersions(version, maxVersion) <= 0
}

// 辅助函数：版本比较
func compareVersions(v1, v2 string) int {
    // 简化的版本比较实现
    if v1 == v2 {
        return 0
    }
    if v1 < v2 {
        return -1
    }
    return 1
}

// 辅助函数：版本小于比较
func isVersionLessThan(version, threshold string) bool {
    return compareVersions(version, threshold) < 0
}

// 辅助函数：版本兼容性检查
func isVersionCompatible(required, available string) bool {
    // 简化的兼容性检查
    return compareVersions(available, required) >= 0
}

// 辅助函数：计算加权分数
func calculateWeightedScore(ref, test *cpeskills.CPE, weights map[string]float64) float64 {
    var score float64
    
    if ref.Vendor == test.Vendor {
        score += weights["vendor"]
    }
    
    if ref.ProductName == test.ProductName {
        score += weights["product"]
    }
    
    if ref.Version == test.Version {
        score += weights["version"]
    }
    
    if ref.Part.ShortName == test.Part.ShortName {
        score += weights["part"]
    }
    
    return score
}
```

## 关键概念

### 1. 匹配策略

- **精确匹配**: 所有字段必须完全相同
- **模糊匹配**: 基于相似度阈值
- **语义匹配**: 理解同义词和缩写
- **上下文匹配**: 根据使用场景调整规则

### 2. 相似度计算

- **字段权重**: 不同字段的重要性不同
- **字符串距离**: 使用编辑距离等算法
- **部分匹配**: 子字符串和前缀匹配
- **模式匹配**: 正则表达式和通配符

### 3. 版本处理

- **范围匹配**: 检查版本是否在范围内
- **兼容性**: 向后兼容性检查
- **语义版本**: 理解major.minor.patch结构
- **特殊格式**: 处理构建号和日期版本

## 最佳实践

1. **选择合适的匹配策略**: 根据用例选择精确或模糊匹配
2. **调整阈值**: 根据数据质量调整相似度阈值
3. **使用上下文**: 在不同场景中应用不同的匹配规则
4. **验证结果**: 始终验证匹配结果的准确性
5. **性能优化**: 对大数据集使用索引和缓存

## 下一步

- 学习[NVD集成](./nvd-integration.md)获取实际漏洞数据
- 探索[CPE集合](./sets.md)进行批量高级匹配
- 查看[存储](./storage.md)来持久化匹配结果
