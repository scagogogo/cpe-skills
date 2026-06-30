# WFN格式

本页面描述了Well-Formed Name (WFN)格式的处理功能，这是CPE的内部标准表示格式。

## WFN概述

Well-Formed Name (WFN) 是CPE的规范内部表示格式，提供了一种标准化的方式来表示CPE组件，使匹配和比较操作更加高效和可靠。

### WFN结构

```go
type WFN struct {
    Part            string // 组件类型 (a, h, o)
    Vendor          string // 供应商
    Product         string // 产品
    Version         string // 版本
    Update          string // 更新
    Edition         string // 版本
    Language        string // 语言
    SoftwareEdition string // 软件版本
    TargetSoftware  string // 目标软件
    TargetHardware  string // 目标硬件
    Other           string // 其他
}
```

### WFN特殊值

```go
const (
    WFNAny           = "*"  // 任意值
    WFNNotApplicable = "-"  // 不适用
)
```

## WFN创建

### NewWFN

创建新的WFN对象。

```go
func NewWFN() *WFN
```

### NewWFNFromCPE

从CPE对象创建WFN。

```go
func NewWFNFromCPE(cpe *CPE) (*WFN, error)
```

**示例：**
```go
// 创建空WFN
wfn := cpeskills.NewWFN()
wfn.Part = "a"
wfn.Vendor = "microsoft"
wfn.Product = "windows"
wfn.Version = "10"

// 从CPE创建WFN
cpeObj, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
wfn, err := cpeskills.NewWFNFromCPE(cpeObj)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("WFN: %s\n", wfn.String())
```

## 格式转换

### CPEToWFN

将CPE对象转换为WFN。

```go
func CPEToWFN(cpe *CPE) (*WFN, error)
```

### WFNToCPE23

将WFN转换为CPE 2.3格式。

```go
func WFNToCPE23(wfn *WFN) (string, error)
```

### WFNToCPE22

将WFN转换为CPE 2.2格式。

```go
func WFNToCPE22(wfn *WFN) (string, error)
```

**示例：**
```go
// CPE到WFN转换
cpeObj, _ := cpeskills.ParseCpe23("cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*")
wfn, err := cpeskills.CPEToWFN(cpeObj)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("原始CPE: %s\n", cpeObj.GetURI())
fmt.Printf("WFN格式: %s\n", wfn.String())

// WFN到CPE转换
cpe23, err := cpeskills.WFNToCPE23(wfn)
if err != nil {
    log.Fatal(err)
}

cpe22, err := cpeskills.WFNToCPE22(wfn)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("转换为CPE 2.3: %s\n", cpe23)
fmt.Printf("转换为CPE 2.2: %s\n", cpe22)
```

## WFN操作

### String

获取WFN的字符串表示。

```go
func (w *WFN) String() string
```

### Clone

克隆WFN对象。

```go
func (w *WFN) Clone() *WFN
```

### Equals

比较两个WFN是否相等。

```go
func (w *WFN) Equals(other *WFN) bool
```

**示例：**
```go
wfn1 := &cpeskills.WFN{
    Part:    "a",
    Vendor:  "apache",
    Product: "tomcat",
    Version: "9.0.0",
}

wfn2 := wfn1.Clone()
wfn2.Version = "9.0.1"

fmt.Printf("WFN1: %s\n", wfn1.String())
fmt.Printf("WFN2: %s\n", wfn2.String())
fmt.Printf("相等: %t\n", wfn1.Equals(wfn2))
```

## WFN验证

### Validate

验证WFN的有效性。

```go
func (w *WFN) Validate() error
```

### ValidateWFN

验证WFN对象。

```go
func ValidateWFN(wfn *WFN) error
```

### IsValidWFNValue

检查值是否为有效的WFN值。

```go
func IsValidWFNValue(value string) bool
```

**示例：**
```go
wfn := &cpeskills.WFN{
    Part:    "a",
    Vendor:  "microsoft",
    Product: "windows",
    Version: "10",
}

// 验证WFN
err := wfn.Validate()
if err != nil {
    fmt.Printf("WFN验证失败: %v\n", err)
} else {
    fmt.Println("WFN验证通过")
}

// 验证单个值
if cpeskills.IsValidWFNValue("microsoft") {
    fmt.Println("'microsoft'是有效的WFN值")
}
```

## WFN匹配

### Match

WFN匹配。

```go
func (w *WFN) Match(other *WFN) bool
```

### MatchWFN

WFN匹配函数。

```go
func MatchWFN(wfn1, wfn2 *WFN) bool
```

### CompareWFN

比较两个WFN。

```go
func CompareWFN(wfn1, wfn2 *WFN) int
```

**示例：**
```go
// 创建匹配模式
pattern := &cpeskills.WFN{
    Part:    "a",
    Vendor:  "microsoft",
    Product: cpeskills.WFNAny, // 任意产品
    Version: cpeskills.WFNAny, // 任意版本
}

// 测试目标
targets := []*cpeskills.WFN{
    {Part: "a", Vendor: "microsoft", Product: "windows", Version: "10"},
    {Part: "a", Vendor: "microsoft", Product: "office", Version: "2019"},
    {Part: "a", Vendor: "oracle", Product: "java", Version: "11"},
}

fmt.Printf("模式: %s\n", pattern.String())
fmt.Println("匹配结果:")

for i, target := range targets {
    match := pattern.Match(target)
    status := "❌"
    if match {
        status = "✅"
    }
    fmt.Printf("  %s 目标%d: %s\n", status, i+1, target.String())
}
```

## WFN规范化

### Normalize

规范化WFN。

```go
func (w *WFN) Normalize() *WFN
```

### NormalizeWFN

规范化WFN对象。

```go
func NormalizeWFN(wfn *WFN) *WFN
```

### NormalizeWFNValue

规范化WFN值。

```go
func NormalizeWFNValue(value string) string
```

**示例：**
```go
// 创建需要规范化的WFN
unnormalized := &cpeskills.WFN{
    Part:    "A",           // 应该是小写
    Vendor:  "Microsoft",   // 应该是小写
    Product: "Windows~10",  // 特殊字符需要转义
    Version: "10.0.19041.1234",
}

fmt.Printf("规范化前: %s\n", unnormalized.String())

// 规范化
normalized := unnormalized.Normalize()
fmt.Printf("规范化后: %s\n", normalized.String())

// 规范化单个值
value := "Product~Name"
normalizedValue := cpeskills.NormalizeWFNValue(value)
fmt.Printf("值规范化: %s -> %s\n", value, normalizedValue)
```

## WFN转义

### QuoteWFNValue

转义WFN值中的特殊字符。

```go
func QuoteWFNValue(value string) string
```

### UnquoteWFNValue

反转义WFN值。

```go
func UnquoteWFNValue(value string) string
```

### EscapeWFNValue

转义WFN值。

```go
func EscapeWFNValue(value string) string
```

**示例：**
```go
// 包含特殊字符的值
specialValue := "product~name*with?chars"

// 转义
quoted := cpeskills.QuoteWFNValue(specialValue)
fmt.Printf("原始值: %s\n", specialValue)
fmt.Printf("转义后: %s\n", quoted)

// 反转义
unquoted := cpeskills.UnquoteWFNValue(quoted)
fmt.Printf("反转义: %s\n", unquoted)

// 验证往返转换
if specialValue == unquoted {
    fmt.Println("✅ 往返转换成功")
} else {
    fmt.Println("❌ 往返转换失败")
}
```

## WFN集合操作

### WFNSet

WFN集合类型。

```go
type WFNSet struct {
    items map[string]*WFN
    mutex sync.RWMutex
}
```

### NewWFNSet

创建新的WFN集合。

```go
func NewWFNSet() *WFNSet
```

### WFN集合方法

```go
// 添加WFN
func (s *WFNSet) Add(wfn *WFN) bool

// 移除WFN
func (s *WFNSet) Remove(wfn *WFN) bool

// 检查是否包含
func (s *WFNSet) Contains(wfn *WFN) bool

// 获取大小
func (s *WFNSet) Size() int

// 转换为切片
func (s *WFNSet) ToSlice() []*WFN
```

**示例：**
```go
// 创建WFN集合
wfnSet := cpeskills.NewWFNSet()

// 添加WFN
wfn1 := &cpeskills.WFN{Part: "a", Vendor: "microsoft", Product: "windows"}
wfn2 := &cpeskills.WFN{Part: "a", Vendor: "apache", Product: "tomcat"}

wfnSet.Add(wfn1)
wfnSet.Add(wfn2)

fmt.Printf("集合大小: %d\n", wfnSet.Size())

// 检查包含
if wfnSet.Contains(wfn1) {
    fmt.Println("集合包含Windows WFN")
}

// 转换为切片
wfnSlice := wfnSet.ToSlice()
fmt.Printf("集合内容:\n")
for i, wfn := range wfnSlice {
    fmt.Printf("  %d. %s\n", i+1, wfn.String())
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
    fmt.Println("=== WFN格式示例 ===")
    
    // 示例1：CPE到WFN转换
    fmt.Println("\n1. CPE到WFN转换:")
    
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:/a:oracle:java:11.0.12",
    }
    
    for i, cpeStr := range cpeStrings {
        fmt.Printf("\n示例 %d: %s\n", i+1, cpeStr)
        
        cpeObj, err := cpeskills.ParseCPE(cpeStr)
        if err != nil {
            log.Printf("解析失败: %v", err)
            continue
        }
        
        wfn, err := cpeskills.CPEToWFN(cpeObj)
        if err != nil {
            log.Printf("转换失败: %v", err)
            continue
        }
        
        fmt.Printf("  WFN: %s\n", wfn.String())
        fmt.Printf("  部件: %s\n", wfn.Part)
        fmt.Printf("  供应商: %s\n", wfn.Vendor)
        fmt.Printf("  产品: %s\n", wfn.Product)
        fmt.Printf("  版本: %s\n", wfn.Version)
    }
    
    // 示例2：WFN到CPE转换
    fmt.Println("\n2. WFN到CPE转换:")
    
    wfn := &cpeskills.WFN{
        Part:    "a",
        Vendor:  "adobe",
        Product: "reader",
        Version: "2021.001.20150",
        Update:  cpeskills.WFNAny,
        Edition: cpeskills.WFNAny,
    }
    
    fmt.Printf("WFN: %s\n", wfn.String())
    
    cpe23, err := cpeskills.WFNToCPE23(wfn)
    if err != nil {
        log.Printf("转换为CPE 2.3失败: %v", err)
    } else {
        fmt.Printf("CPE 2.3: %s\n", cpe23)
    }
    
    cpe22, err := cpeskills.WFNToCPE22(wfn)
    if err != nil {
        log.Printf("转换为CPE 2.2失败: %v", err)
    } else {
        fmt.Printf("CPE 2.2: %s\n", cpe22)
    }
    
    // 示例3：WFN匹配
    fmt.Println("\n3. WFN匹配:")
    
    pattern := &cpeskills.WFN{
        Part:    "a",
        Vendor:  "microsoft",
        Product: cpeskills.WFNAny,
        Version: cpeskills.WFNAny,
    }
    
    targets := []*cpeskills.WFN{
        {Part: "a", Vendor: "microsoft", Product: "windows", Version: "10"},
        {Part: "a", Vendor: "microsoft", Product: "office", Version: "2019"},
        {Part: "a", Vendor: "oracle", Product: "java", Version: "11"},
        {Part: "o", Vendor: "microsoft", Product: "windows", Version: "10"},
    }
    
    fmt.Printf("匹配模式: %s\n", pattern.String())
    fmt.Println("匹配结果:")
    
    for i, target := range targets {
        match := pattern.Match(target)
        status := "❌"
        if match {
            status = "✅"
        }
        fmt.Printf("  %s 目标%d: %s\n", status, i+1, target.String())
    }
    
    // 示例4：WFN验证
    fmt.Println("\n4. WFN验证:")
    
    testWFNs := []*cpeskills.WFN{
        {Part: "a", Vendor: "microsoft", Product: "windows"},     // 有效
        {Part: "x", Vendor: "microsoft", Product: "windows"},     // 无效部件
        {Part: "a", Vendor: "", Product: "windows"},              // 空供应商
        {Part: "a", Vendor: "microsoft", Product: ""},            // 空产品
    }
    
    for i, testWFN := range testWFNs {
        err := testWFN.Validate()
        status := "✅"
        if err != nil {
            status = "❌"
        }
        
        fmt.Printf("  %s 测试%d: %s\n", status, i+1, testWFN.String())
        if err != nil {
            fmt.Printf("     错误: %v\n", err)
        }
    }
    
    // 示例5：WFN规范化
    fmt.Println("\n5. WFN规范化:")
    
    unnormalized := &cpeskills.WFN{
        Part:    "A",
        Vendor:  "Microsoft",
        Product: "Windows~10",
        Version: "10.0.19041.1234",
    }
    
    fmt.Printf("规范化前: %s\n", unnormalized.String())
    
    normalized := unnormalized.Normalize()
    fmt.Printf("规范化后: %s\n", normalized.String())
    
    // 示例6：特殊字符处理
    fmt.Println("\n6. 特殊字符处理:")
    
    specialValues := []string{
        "product~name",
        "version*with?wildcards",
        "name:with:colons",
        "path\\with\\backslashes",
    }
    
    for _, value := range specialValues {
        quoted := cpeskills.QuoteWFNValue(value)
        unquoted := cpeskills.UnquoteWFNValue(quoted)
        
        fmt.Printf("  原始: %s\n", value)
        fmt.Printf("  转义: %s\n", quoted)
        fmt.Printf("  还原: %s\n", unquoted)
        fmt.Printf("  正确: %t\n\n", value == unquoted)
    }
}
```

## 下一步

- 了解[匹配算法](./matching.md)来使用WFN进行高效匹配
- 学习[集合操作](./sets.md)来批量处理WFN对象
- 探索[验证功能](./validation.md)来确保WFN数据质量
