# 解析功能

本页面描述了CPE库中用于解析CPE字符串的函数和方法，包括CPE 2.2和2.3格式的解析功能。

## 主要解析函数

### ParseCpe23

解析CPE 2.3格式字符串。

```go
func ParseCpe23(cpe23 string) (*CPE, error)
```

**参数：**
- `cpe23`: CPE 2.3格式字符串

**返回值：**
- `*CPE`: 解析后的CPE对象
- `error`: 解析错误（如果有）

**示例：**
```go
cpeObj, err := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("供应商: %s, 产品: %s\n", cpeObj.Vendor, cpeObj.ProductName)
```

### ParseCpe22

解析CPE 2.2格式字符串。

```go
func ParseCpe22(cpe22 string) (*CPE, error)
```

**参数：**
- `cpe22`: CPE 2.2格式字符串

**返回值：**
- `*CPE`: 解析后的CPE对象
- `error`: 解析错误（如果有）

**示例：**
```go
cpeObj, err := cpeskills.ParseCpe22("cpe:/a:apache:tomcat:9.0.0")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("产品: %s, 版本: %s\n", cpeObj.ProductName, cpeObj.Version)
```

### ParseCPE

通用CPE解析函数，自动检测格式。

```go
func ParseCPE(cpeString string) (*CPE, error)
```

**参数：**
- `cpeString`: CPE字符串（2.2或2.3格式）

**返回值：**
- `*CPE`: 解析后的CPE对象
- `error`: 解析错误（如果有）

**示例：**
```go
// 自动检测格式
cpe23, _ := cpeskills.ParseCPE("cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*")
cpe22, _ := cpeskills.ParseCPE("cpe:/a:oracle:java:11.0.12")

fmt.Printf("两个CPE是否相同: %t\n", cpe23.Match(cpe22))
```

## 批量解析函数

### ParseCPEList

批量解析CPE字符串列表。

```go
func ParseCPEList(cpeStrings []string) ([]*CPE, []error)
```

**参数：**
- `cpeStrings`: CPE字符串列表

**返回值：**
- `[]*CPE`: 成功解析的CPE对象列表
- `[]error`: 解析错误列表

**示例：**
```go
cpeStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:/a:apache:tomcat:9.0.0",
    "invalid-cpe-string",
}

cpes, errors := cpeskills.ParseCPEList(cpeStrings)
fmt.Printf("成功解析: %d, 错误: %d\n", len(cpes), len(errors))
```

### ParseCPEFromFile

从文件中解析CPE字符串。

```go
func ParseCPEFromFile(filename string) ([]*CPE, error)
```

**参数：**
- `filename`: 包含CPE字符串的文件路径

**返回值：**
- `[]*CPE`: 解析后的CPE对象列表
- `error`: 文件读取或解析错误

**示例：**
```go
cpes, err := cpeskills.ParseCPEFromFile("cpe_list.txt")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("从文件解析了 %d 个CPE\n", len(cpes))
```

## 解析选项

### ParseOptions

解析选项配置。

```go
type ParseOptions struct {
    StrictMode      bool // 严格模式
    AllowEmpty      bool // 允许空字段
    NormalizeCase   bool // 规范化大小写
    ValidateFormat  bool // 验证格式
    IgnoreErrors    bool // 忽略非致命错误
}
```

### ParseWithOptions

使用自定义选项解析CPE。

```go
func ParseWithOptions(cpeString string, options *ParseOptions) (*CPE, error)
```

**示例：**
```go
options := &cpeskills.ParseOptions{
    StrictMode:     true,
    NormalizeCase:  true,
    ValidateFormat: true,
}

cpeObj, err := cpeskills.ParseWithOptions(cpeString, options)
if err != nil {
    log.Printf("严格模式解析失败: %v", err)
}
```

## 格式验证

### ValidateCPEFormat

验证CPE字符串格式。

```go
func ValidateCPEFormat(cpeString string) error
```

**参数：**
- `cpeString`: 要验证的CPE字符串

**返回值：**
- `error`: 验证错误（如果格式无效）

**示例：**
```go
err := cpeskills.ValidateCPEFormat("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
if err != nil {
    fmt.Printf("格式无效: %v\n", err)
} else {
    fmt.Println("格式有效")
}
```

### IsCPE23Format

检查字符串是否为CPE 2.3格式。

```go
func IsCPE23Format(cpeString string) bool
```

### IsCPE22Format

检查字符串是否为CPE 2.2格式。

```go
func IsCPE22Format(cpeString string) bool
```

**示例：**
```go
cpeString := "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"

if cpeskills.IsCPE23Format(cpeString) {
    fmt.Println("这是CPE 2.3格式")
} else if cpeskills.IsCPE22Format(cpeString) {
    fmt.Println("这是CPE 2.2格式")
} else {
    fmt.Println("未知格式")
}
```

## 组件解析

### ParsePart

解析组件类型。

```go
func ParsePart(partString string) (Part, error)
```

**示例：**
```go
part, err := cpeskills.ParsePart("a")
if err != nil {
    log.Fatal(err)
}
fmt.Printf("组件类型: %s (%s)\n", part.ShortName, part.LongName)
```

### ParseVendor

解析供应商信息。

```go
func ParseVendor(vendorString string) (Vendor, error)
```

### ParseProduct

解析产品信息。

```go
func ParseProduct(productString string) (Product, error)
```

### ParseVersion

解析版本信息。

```go
func ParseVersion(versionString string) (Version, error)
```

## 特殊值处理

### 通配符处理

CPE解析器能够正确处理特殊值：

- `*` (星号): 表示"任意值"
- `-` (连字符): 表示"不适用"

```go
// 解析包含通配符的CPE
cpeObj, _ := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")
fmt.Printf("产品: %s\n", cpeObj.ProductName) // 输出: *
```

### 转义字符处理

解析器能够处理CPE中的转义字符：

```go
// 包含特殊字符的CPE
cpeObj, _ := cpeskills.ParseCpe23("cpe:2.3:a:vendor:product\\~name:1.0:*:*:*:*:*:*:*")
fmt.Printf("产品: %s\n", cpeObj.ProductName) // 输出: product~name
```

## 错误处理

### 解析错误类型

解析过程中可能遇到的错误类型：

```go
const (
    ErrorInvalidFormat   = "无效的CPE格式"
    ErrorInvalidPart     = "无效的组件类型"
    ErrorInvalidVendor   = "无效的供应商"
    ErrorInvalidProduct  = "无效的产品名称"
    ErrorInvalidVersion  = "无效的版本信息"
    ErrorTooFewComponents = "组件数量不足"
    ErrorTooManyComponents = "组件数量过多"
)
```

### 错误检查函数

```go
// 检查是否为格式错误
func IsFormatError(err error) bool

// 检查是否为组件错误
func IsComponentError(err error) bool

// 检查是否为验证错误
func IsValidationError(err error) bool
```

**示例：**
```go
_, err := cpeskills.ParseCpe23("invalid-cpe")
if cpeskills.IsFormatError(err) {
    fmt.Println("这是一个格式错误")
}
```

## 性能优化

### 解析缓存

对于频繁解析的CPE字符串，可以使用缓存：

```go
// 启用解析缓存
cpeskills.EnableParseCache(1000) // 缓存1000个解析结果

// 解析CPE（会被缓存）
cpeObj1, _ := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
cpeObj2, _ := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*") // 从缓存获取

// 清除缓存
cpeskills.ClearParseCache()
```

### 批量解析优化

```go
// 并行解析大量CPE
func ParseCPEListParallel(cpeStrings []string, workers int) ([]*CPE, []error)
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
    // 基本解析示例
    fmt.Println("=== 基本解析示例 ===")
    
    // 解析CPE 2.3
    cpe23, err := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("CPE 2.3解析结果:\n")
    fmt.Printf("  组件类型: %s\n", cpe23.Part.LongName)
    fmt.Printf("  供应商: %s\n", cpe23.Vendor)
    fmt.Printf("  产品: %s\n", cpe23.ProductName)
    fmt.Printf("  版本: %s\n", cpe23.Version)
    
    // 解析CPE 2.2
    cpe22, err := cpeskills.ParseCpe22("cpe:/a:apache:tomcat:9.0.0")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("\nCPE 2.2解析结果:\n")
    fmt.Printf("  供应商: %s\n", cpe22.Vendor)
    fmt.Printf("  产品: %s\n", cpe22.ProductName)
    fmt.Printf("  版本: %s\n", cpe22.Version)
    
    // 批量解析示例
    fmt.Println("\n=== 批量解析示例 ===")
    
    cpeStrings := []string{
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:/a:nginx:nginx:1.18.0",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
        "invalid-cpe-format", // 这个会产生错误
    }
    
    cpes, errors := cpeskills.ParseCPEList(cpeStrings)
    
    fmt.Printf("成功解析: %d 个CPE\n", len(cpes))
    fmt.Printf("解析错误: %d 个\n", len(errors))
    
    for i, cpeObj := range cpes {
        fmt.Printf("  %d. %s %s %s\n", i+1, cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
    }
    
    for i, err := range errors {
        fmt.Printf("  错误 %d: %v\n", i+1, err)
    }
    
    // 格式验证示例
    fmt.Println("\n=== 格式验证示例 ===")
    
    testStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:/a:apache:tomcat:9.0.0",
        "invalid-format",
        "cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*", // 无效的组件类型
    }
    
    for _, testStr := range testStrings {
        err := cpeskills.ValidateCPEFormat(testStr)
        if err != nil {
            fmt.Printf("❌ %s - %v\n", testStr, err)
        } else {
            fmt.Printf("✅ %s - 格式有效\n", testStr)
        }
    }
}
```

## 下一步

- 学习[匹配算法](./matching.md)来比较解析后的CPE对象
- 了解[WFN格式](./wfn.md)来处理内部表示
- 探索[验证功能](./validation.md)来确保CPE数据质量
