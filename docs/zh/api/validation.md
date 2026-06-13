# 验证功能

本页面描述了CPE库中用于验证CPE数据完整性和正确性的各种验证功能。

## 基本验证

### ValidateCPE

验证CPE对象的完整性。

```go
func ValidateCPE(cpe *CPE) error
```

### ValidateCPEString

验证CPE字符串格式。

```go
func ValidateCPEString(cpeString string) error
```

### ValidateCPEFormat

验证CPE格式是否正确。

```go
func ValidateCPEFormat(cpeString string) error
```

**示例：**
```go
// 验证CPE对象
cpeObj, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
err := cpe.ValidateCPE(cpeObj)
if err != nil {
    fmt.Printf("CPE验证失败: %v\n", err)
} else {
    fmt.Println("✅ CPE验证通过")
}

// 验证CPE字符串
testStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",  // 有效
    "cpe:/a:apache:tomcat:9.0.0",                    // 有效
    "invalid-cpe-format",                            // 无效
    "cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*",   // 无效部件
}

for _, testStr := range testStrings {
    err := cpe.ValidateCPEString(testStr)
    status := "✅"
    if err != nil {
        status = "❌"
    }
    fmt.Printf("%s %s\n", status, testStr)
    if err != nil {
        fmt.Printf("   错误: %v\n", err)
    }
}
```

## 组件验证

### ValidatePart

验证组件类型。

```go
func ValidatePart(part string) error
```

### ValidateVendor

验证供应商名称。

```go
func ValidateVendor(vendor string) error
```

### ValidateProduct

验证产品名称。

```go
func ValidateProduct(product string) error
```

### ValidateVersion

验证版本信息。

```go
func ValidateVersion(version string) error
```

**示例：**
```go
// 验证各个组件
components := map[string]string{
    "part":    "a",
    "vendor":  "microsoft",
    "product": "windows",
    "version": "10.0.19041",
}

validators := map[string]func(string) error{
    "part":    cpe.ValidatePart,
    "vendor":  cpe.ValidateVendor,
    "product": cpe.ValidateProduct,
    "version": cpe.ValidateVersion,
}

fmt.Println("组件验证结果:")
for component, value := range components {
    validator := validators[component]
    err := validator(value)
    
    status := "✅"
    if err != nil {
        status = "❌"
    }
    
    fmt.Printf("  %s %s: %s\n", status, component, value)
    if err != nil {
        fmt.Printf("     错误: %v\n", err)
    }
}
```

## 格式验证

### IsCPE23Format

检查是否为CPE 2.3格式。

```go
func IsCPE23Format(cpeString string) bool
```

### IsCPE22Format

检查是否为CPE 2.2格式。

```go
func IsCPE22Format(cpeString string) bool
```

### IsValidCPEFormat

检查是否为有效的CPE格式。

```go
func IsValidCPEFormat(cpeString string) bool
```

**示例：**
```go
testFormats := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:/a:apache:tomcat:9.0.0",
    "invalid-format",
    "cpe:2.3:incomplete",
}

fmt.Println("格式检查结果:")
for _, testStr := range testFormats {
    is23 := cpe.IsCPE23Format(testStr)
    is22 := cpe.IsCPE22Format(testStr)
    isValid := cpe.IsValidCPEFormat(testStr)
    
    fmt.Printf("字符串: %s\n", testStr)
    fmt.Printf("  CPE 2.3: %t\n", is23)
    fmt.Printf("  CPE 2.2: %t\n", is22)
    fmt.Printf("  有效格式: %t\n\n", isValid)
}
```

## 语义验证

### ValidateSemantics

验证CPE的语义正确性。

```go
func ValidateSemantics(cpe *CPE) error
```

### ValidateConsistency

验证CPE组件之间的一致性。

```go
func ValidateConsistency(cpe *CPE) error
```

### ValidateLogicalConstraints

验证逻辑约束。

```go
func ValidateLogicalConstraints(cpe *CPE) error
```

**示例：**
```go
// 创建测试CPE
testCPEs := []*cpe.CPE{
    // 语义正确的CPE
    {
        Part:        cpe.PartApplication,
        Vendor:      "microsoft",
        ProductName: "windows",
        Version:     "10",
    },
    // 语义不一致的CPE（操作系统标记为应用程序）
    {
        Part:        cpe.PartApplication,
        Vendor:      "microsoft",
        ProductName: "windows_server", // 这应该是操作系统
        Version:     "2019",
    },
}

fmt.Println("语义验证结果:")
for i, testCPE := range testCPEs {
    fmt.Printf("\n测试CPE %d: %s %s %s\n", i+1, 
        testCPE.Vendor, testCPE.ProductName, testCPE.Version)
    
    // 语义验证
    err := cpe.ValidateSemantics(testCPE)
    if err != nil {
        fmt.Printf("  ❌ 语义验证失败: %v\n", err)
    } else {
        fmt.Printf("  ✅ 语义验证通过\n")
    }
    
    // 一致性验证
    err = cpe.ValidateConsistency(testCPE)
    if err != nil {
        fmt.Printf("  ❌ 一致性验证失败: %v\n", err)
    } else {
        fmt.Printf("  ✅ 一致性验证通过\n")
    }
}
```

## 批量验证

### ValidateCPEList

批量验证CPE列表。

```go
func ValidateCPEList(cpes []*CPE) []ValidationResult
```

### ValidationResult

验证结果结构。

```go
type ValidationResult struct {
    CPE     *CPE          // 被验证的CPE
    Valid   bool          // 是否有效
    Errors  []error       // 错误列表
    Warnings []string     // 警告列表
    Score   float64       // 质量分数 (0-1)
}
```

### ValidateCPESet

验证CPE集合。

```go
func ValidateCPESet(cpeSet *CPESet) *SetValidationResult
```

**示例：**
```go
// 创建测试CPE列表
cpeStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    "invalid-cpe-format",
    "cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*", // 无效部件
}

cpes := []*cpe.CPE{}
for _, cpeStr := range cpeStrings {
    cpeObj, err := cpe.ParseCPE(cpeStr)
    if err == nil {
        cpes = append(cpes, cpeObj)
    } else {
        // 为无效CPE创建占位符
        cpes = append(cpes, &cpe.CPE{Cpe23: cpeStr})
    }
}

// 批量验证
results := cpe.ValidateCPEList(cpes)

fmt.Println("批量验证结果:")
for i, result := range results {
    fmt.Printf("\nCPE %d: %s\n", i+1, cpeStrings[i])
    fmt.Printf("  有效: %t\n", result.Valid)
    fmt.Printf("  质量分数: %.2f\n", result.Score)
    
    if len(result.Errors) > 0 {
        fmt.Printf("  错误:\n")
        for _, err := range result.Errors {
            fmt.Printf("    - %v\n", err)
        }
    }
    
    if len(result.Warnings) > 0 {
        fmt.Printf("  警告:\n")
        for _, warning := range result.Warnings {
            fmt.Printf("    - %s\n", warning)
        }
    }
}
```

## 自定义验证

### ValidationRule

验证规则接口。

```go
type ValidationRule interface {
    Validate(cpe *CPE) error
    GetName() string
    GetDescription() string
}
```

### Validator

验证器结构。

```go
type Validator struct {
    Rules   []ValidationRule // 验证规则列表
    Strict  bool            // 严格模式
    Options *ValidationOptions // 验证选项
}
```

### ValidationOptions

验证选项。

```go
type ValidationOptions struct {
    CheckSemantics    bool // 检查语义
    CheckConsistency  bool // 检查一致性
    CheckFormat       bool // 检查格式
    AllowDeprecated   bool // 允许已弃用的CPE
    RequireComplete   bool // 要求完整的CPE
}
```

**示例：**
```go
// 创建自定义验证规则
type VendorWhitelistRule struct {
    AllowedVendors []string
}

func (r *VendorWhitelistRule) Validate(cpe *cpe.CPE) error {
    for _, allowed := range r.AllowedVendors {
        if cpe.Vendor == allowed {
            return nil
        }
    }
    return fmt.Errorf("供应商 '%s' 不在白名单中", cpe.Vendor)
}

func (r *VendorWhitelistRule) GetName() string {
    return "VendorWhitelist"
}

func (r *VendorWhitelistRule) GetDescription() string {
    return "检查供应商是否在允许的白名单中"
}

// 使用自定义验证器
validator := &cpe.Validator{
    Rules: []cpe.ValidationRule{
        &VendorWhitelistRule{
            AllowedVendors: []string{"microsoft", "apache", "oracle"},
        },
    },
    Strict: true,
    Options: &cpe.ValidationOptions{
        CheckSemantics:   true,
        CheckConsistency: true,
        CheckFormat:      true,
    },
}

// 验证CPE
testCPE, _ := cpe.ParseCpe23("cpe:2.3:a:unknown_vendor:product:1.0:*:*:*:*:*:*:*")
err := validator.Validate(testCPE)
if err != nil {
    fmt.Printf("自定义验证失败: %v\n", err)
}
```

## 质量评估

### AssessQuality

评估CPE质量。

```go
func AssessQuality(cpe *CPE) *QualityAssessment
```

### QualityAssessment

质量评估结果。

```go
type QualityAssessment struct {
    OverallScore    float64            // 总体分数 (0-1)
    ComponentScores map[string]float64 // 各组件分数
    Issues          []QualityIssue     // 质量问题
    Recommendations []string           // 改进建议
}
```

### QualityIssue

质量问题。

```go
type QualityIssue struct {
    Type        IssueType // 问题类型
    Severity    Severity  // 严重程度
    Component   string    // 相关组件
    Description string    // 问题描述
    Suggestion  string    // 改进建议
}
```

**示例：**
```go
// 评估CPE质量
testCPEs := []string{
    "cpe:2.3:a:microsoft:windows:10.0.19041.1234:*:*:*:*:*:*:*", // 高质量
    "cpe:2.3:a:vendor:product:1.0:*:*:*:*:*:*:*",                // 中等质量
    "cpe:2.3:a:*:*:*:*:*:*:*:*:*:*",                            // 低质量
}

fmt.Println("CPE质量评估:")
for i, cpeStr := range testCPEs {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
    if err != nil {
        continue
    }
    
    assessment := cpe.AssessQuality(cpeObj)
    
    fmt.Printf("\nCPE %d: %s\n", i+1, cpeStr)
    fmt.Printf("  总体分数: %.2f\n", assessment.OverallScore)
    
    fmt.Printf("  组件分数:\n")
    for component, score := range assessment.ComponentScores {
        fmt.Printf("    %s: %.2f\n", component, score)
    }
    
    if len(assessment.Issues) > 0 {
        fmt.Printf("  质量问题:\n")
        for _, issue := range assessment.Issues {
            fmt.Printf("    - %s: %s\n", issue.Component, issue.Description)
        }
    }
    
    if len(assessment.Recommendations) > 0 {
        fmt.Printf("  改进建议:\n")
        for _, rec := range assessment.Recommendations {
            fmt.Printf("    - %s\n", rec)
        }
    }
}
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE验证功能示例 ===")
    
    // 测试CPE字符串
    testCPEs := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",  // 有效
        "cpe:/a:apache:tomcat:9.0.0",                    // 有效
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",  // 有效
        "invalid-cpe-format",                            // 格式错误
        "cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*",   // 无效部件
        "cpe:2.3:a::product:1.0:*:*:*:*:*:*:*",         // 空供应商
    }
    
    fmt.Println("1. 基本验证:")
    for i, cpeStr := range testCPEs {
        fmt.Printf("\n测试 %d: %s\n", i+1, cpeStr)
        
        // 格式验证
        err := cpe.ValidateCPEFormat(cpeStr)
        if err != nil {
            fmt.Printf("  ❌ 格式验证失败: %v\n", err)
            continue
        } else {
            fmt.Printf("  ✅ 格式验证通过\n")
        }
        
        // 解析并验证CPE对象
        cpeObj, err := cpe.ParseCPE(cpeStr)
        if err != nil {
            fmt.Printf("  ❌ 解析失败: %v\n", err)
            continue
        }
        
        err = cpe.ValidateCPE(cpeObj)
        if err != nil {
            fmt.Printf("  ❌ CPE验证失败: %v\n", err)
        } else {
            fmt.Printf("  ✅ CPE验证通过\n")
        }
        
        // 质量评估
        assessment := cpe.AssessQuality(cpeObj)
        fmt.Printf("  质量分数: %.2f\n", assessment.OverallScore)
    }
    
    fmt.Println("\n2. 组件验证:")
    components := map[string][]string{
        "part": {"a", "h", "o", "x"},           // x是无效的
        "vendor": {"microsoft", "apache", ""},   // 空字符串无效
        "product": {"windows", "tomcat", ""},    // 空字符串无效
        "version": {"10", "9.0.0", "*", "-"},   // 都是有效的
    }
    
    validators := map[string]func(string) error{
        "part":    cpe.ValidatePart,
        "vendor":  cpe.ValidateVendor,
        "product": cpe.ValidateProduct,
        "version": cpe.ValidateVersion,
    }
    
    for component, values := range components {
        fmt.Printf("\n%s 验证:\n", component)
        validator := validators[component]
        
        for _, value := range values {
            err := validator(value)
            status := "✅"
            if err != nil {
                status = "❌"
            }
            
            displayValue := value
            if displayValue == "" {
                displayValue = "(空字符串)"
            }
            
            fmt.Printf("  %s %s", status, displayValue)
            if err != nil {
                fmt.Printf(" - %v", err)
            }
            fmt.Println()
        }
    }
    
    fmt.Println("\n3. 批量验证:")
    
    // 创建CPE列表进行批量验证
    cpeList := []*cpe.CPE{}
    for _, cpeStr := range testCPEs[:3] { // 只使用前3个有效的
        cpeObj, err := cpe.ParseCPE(cpeStr)
        if err == nil {
            cpeList = append(cpeList, cpeObj)
        }
    }
    
    results := cpe.ValidateCPEList(cpeList)
    
    fmt.Printf("批量验证了 %d 个CPE:\n", len(results))
    for i, result := range results {
        fmt.Printf("  CPE %d: 有效=%t, 分数=%.2f\n", 
            i+1, result.Valid, result.Score)
        
        if len(result.Warnings) > 0 {
            fmt.Printf("    警告: %v\n", result.Warnings)
        }
    }
    
    fmt.Println("\n4. 自定义验证规则:")
    
    // 创建自定义验证器
    validator := &cpe.Validator{
        Strict: true,
        Options: &cpe.ValidationOptions{
            CheckSemantics:   true,
            CheckConsistency: true,
            CheckFormat:      true,
            RequireComplete:  true,
        },
    }
    
    testCPE, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    err := validator.Validate(testCPE)
    if err != nil {
        fmt.Printf("自定义验证失败: %v\n", err)
    } else {
        fmt.Printf("✅ 自定义验证通过\n")
    }
    
    fmt.Println("\n5. 质量评估详情:")
    
    highQualityCPE, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10.0.19041.1234:*:*:*:*:*:*:*")
    assessment := cpe.AssessQuality(highQualityCPE)
    
    fmt.Printf("高质量CPE评估:\n")
    fmt.Printf("  总体分数: %.2f\n", assessment.OverallScore)
    fmt.Printf("  组件分数:\n")
    for component, score := range assessment.ComponentScores {
        fmt.Printf("    %s: %.2f\n", component, score)
    }
    
    if len(assessment.Recommendations) > 0 {
        fmt.Printf("  改进建议:\n")
        for _, rec := range assessment.Recommendations {
            fmt.Printf("    - %s\n", rec)
        }
    }
}
```

## 下一步

- 了解[错误处理](./errors.md)来处理验证错误
- 学习[存储接口](./storage.md)来持久化验证结果
- 探索[集合操作](./sets.md)来批量验证CPE数据
