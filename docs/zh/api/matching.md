# 匹配算法

本页面描述了CPE库中用于比较和匹配CPE对象的各种算法和函数。

## 基本匹配函数

### Match

CPE对象的基本匹配方法。

```go
func (c *CPE) Match(other *CPE) bool
```

**参数：**
- `other`: 要匹配的另一个CPE对象

**返回值：**
- `bool`: 是否匹配

**示例：**
```go
cpe1, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
cpe2, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*")

if cpe1.Match(cpe2) {
    fmt.Println("CPE匹配成功")
}
```

### MatchCPE

高级CPE匹配函数，支持自定义选项。

```go
func MatchCPE(cpe1, cpe2 *CPE, options *MatchOptions) bool
```

**参数：**
- `cpe1`: 第一个CPE对象
- `cpe2`: 第二个CPE对象
- `options`: 匹配选项

**返回值：**
- `bool`: 是否匹配

**示例：**
```go
options := &cpe.MatchOptions{
    ExactMatch:     false,
    IgnoreCase:     true,
    AllowWildcards: true,
}

match := cpe.MatchCPE(cpe1, cpe2, options)
```

## 高级匹配算法

### FuzzyMatch

模糊匹配算法，返回相似度分数。

```go
func FuzzyMatch(cpe1, cpe2 *CPE) float64
```

**参数：**
- `cpe1`: 第一个CPE对象
- `cpe2`: 第二个CPE对象

**返回值：**
- `float64`: 相似度分数（0.0-1.0）

**示例：**
```go
score := cpe.FuzzyMatch(cpe1, cpe2)
fmt.Printf("相似度分数: %.2f\n", score)

if score >= 0.8 {
    fmt.Println("高度相似")
} else if score >= 0.6 {
    fmt.Println("中等相似")
} else {
    fmt.Println("相似度较低")
}
```

### WeightedMatch

加权匹配算法，允许为不同组件设置权重。

```go
func WeightedMatch(cpe1, cpe2 *CPE, weights MatchWeights) float64
```

**参数：**
- `cpe1`: 第一个CPE对象
- `cpe2`: 第二个CPE对象
- `weights`: 组件权重配置

**返回值：**
- `float64`: 加权匹配分数

**示例：**
```go
weights := cpe.MatchWeights{
    Part:    0.1,  // 组件类型权重较低
    Vendor:  0.3,  // 供应商权重中等
    Product: 0.4,  // 产品权重最高
    Version: 0.2,  // 版本权重中等
}

score := cpe.WeightedMatch(cpe1, cpe2, weights)
fmt.Printf("加权匹配分数: %.3f\n", score)
```

### SemanticMatch

语义匹配算法，理解同义词和缩写。

```go
func SemanticMatch(cpe1, cpe2 *CPE) bool
```

**示例：**
```go
// 这些CPE在语义上是等价的
ie1, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:internet_explorer:*:*:*:*:*:*:*:*")
ie2, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:ie:*:*:*:*:*:*:*:*")

if cpe.SemanticMatch(ie1, ie2) {
    fmt.Println("语义匹配成功")
}
```

## 模式匹配

### MatchPattern

模式匹配函数，支持通配符和正则表达式。

```go
func MatchPattern(target *CPE, pattern *CPE) bool
```

**参数：**
- `target`: 目标CPE对象
- `pattern`: 模式CPE对象

**返回值：**
- `bool`: 是否匹配模式

**示例：**
```go
// 创建匹配所有Microsoft产品的模式
pattern, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")

// 测试目标CPE
target, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*")

if cpe.MatchPattern(target, pattern) {
    fmt.Println("匹配Microsoft产品模式")
}
```

### MatchRegex

正则表达式匹配。

```go
func MatchRegex(cpe *CPE, field string, pattern string) bool
```

**参数：**
- `cpe`: CPE对象
- `field`: 要匹配的字段名
- `pattern`: 正则表达式模式

**示例：**
```go
// 匹配版本号模式
match := cpe.MatchRegex(cpeObj, "version", `^\d+\.\d+\.\d+$`)
if match {
    fmt.Println("版本号格式正确")
}
```

## 版本匹配

### CompareVersions

比较两个版本字符串。

```go
func CompareVersions(version1, version2 string) int
```

**返回值：**
- `-1`: version1 < version2
- `0`: version1 == version2
- `1`: version1 > version2

**示例：**
```go
result := cpe.CompareVersions("1.2.3", "1.2.4")
switch result {
case -1:
    fmt.Println("版本1较旧")
case 0:
    fmt.Println("版本相同")
case 1:
    fmt.Println("版本1较新")
}
```

### IsVersionInRange

检查版本是否在指定范围内。

```go
func IsVersionInRange(version, minVersion, maxVersion string) bool
```

**示例：**
```go
inRange := cpe.IsVersionInRange("1.2.5", "1.2.0", "1.3.0")
if inRange {
    fmt.Println("版本在范围内")
}
```

### MatchVersionPattern

版本模式匹配。

```go
func MatchVersionPattern(version, pattern string) bool
```

**示例：**
```go
// 匹配1.x.x版本
match := cpe.MatchVersionPattern("1.2.3", "1.*.*")
if match {
    fmt.Println("匹配版本模式")
}
```

## 集合匹配

### MatchAny

检查CPE是否匹配集合中的任一项。

```go
func MatchAny(target *CPE, cpeSet *CPESet) bool
```

### MatchAll

检查CPE是否匹配集合中的所有项。

```go
func MatchAll(target *CPE, cpeSet *CPESet) bool
```

### FindMatches

在集合中查找所有匹配项。

```go
func FindMatches(target *CPE, cpeSet *CPESet) []*CPE
```

**示例：**
```go
// 创建CPE集合
cpeSet := cpe.NewCPESet()
cpeSet.Add(cpe1)
cpeSet.Add(cpe2)
cpeSet.Add(cpe3)

// 查找匹配项
matches := cpe.FindMatches(targetCPE, cpeSet)
fmt.Printf("找到 %d 个匹配项\n", len(matches))
```

## 匹配选项配置

### MatchOptions

匹配选项结构体。

```go
type MatchOptions struct {
    ExactMatch      bool    // 精确匹配
    IgnoreCase      bool    // 忽略大小写
    AllowWildcards  bool    // 允许通配符
    FuzzyThreshold  float64 // 模糊匹配阈值
    UseSemantics    bool    // 使用语义匹配
    VersionTolerance string // 版本容差
}
```

### DefaultMatchOptions

获取默认匹配选项。

```go
func DefaultMatchOptions() *MatchOptions
```

### StrictMatchOptions

获取严格匹配选项。

```go
func StrictMatchOptions() *MatchOptions
```

### FuzzyMatchOptions

获取模糊匹配选项。

```go
func FuzzyMatchOptions(threshold float64) *MatchOptions
```

## 匹配结果

### MatchResult

详细的匹配结果。

```go
type MatchResult struct {
    Match       bool              // 是否匹配
    Score       float64           // 匹配分数
    Confidence  float64           // 置信度
    Details     map[string]float64 // 各组件匹配详情
    Explanation string            // 匹配说明
}
```

### DetailedMatch

获取详细匹配结果。

```go
func DetailedMatch(cpe1, cpe2 *CPE, options *MatchOptions) *MatchResult
```

**示例：**
```go
result := cpe.DetailedMatch(cpe1, cpe2, options)

fmt.Printf("匹配结果: %t\n", result.Match)
fmt.Printf("匹配分数: %.3f\n", result.Score)
fmt.Printf("置信度: %.3f\n", result.Confidence)
fmt.Printf("说明: %s\n", result.Explanation)

for component, score := range result.Details {
    fmt.Printf("  %s: %.3f\n", component, score)
}
```

## 性能优化

### 匹配缓存

```go
// 启用匹配缓存
cpe.EnableMatchCache(5000)

// 匹配操作会被缓存
match1 := cpe1.Match(cpe2) // 计算并缓存
match2 := cpe1.Match(cpe2) // 从缓存获取

// 清除缓存
cpe.ClearMatchCache()
```

### 批量匹配

```go
// 批量匹配优化
func BatchMatch(targets []*CPE, patterns []*CPE) [][]bool
```

### 并行匹配

```go
// 并行匹配大型数据集
func ParallelMatch(targets []*CPE, pattern *CPE, workers int) []bool
```

## 匹配策略

### 精确匹配策略

```go
func ExactMatchStrategy() MatchStrategy
```

### 模糊匹配策略

```go
func FuzzyMatchStrategy(threshold float64) MatchStrategy
```

### 语义匹配策略

```go
func SemanticMatchStrategy() MatchStrategy
```

### 自定义匹配策略

```go
type MatchStrategy interface {
    Match(cpe1, cpe2 *CPE) bool
    Score(cpe1, cpe2 *CPE) float64
}

// 实现自定义策略
type CustomMatchStrategy struct {
    // 自定义字段
}

func (s *CustomMatchStrategy) Match(cpe1, cpe2 *CPE) bool {
    // 自定义匹配逻辑
    return true
}

func (s *CustomMatchStrategy) Score(cpe1, cpe2 *CPE) float64 {
    // 自定义评分逻辑
    return 0.8
}
```

## 完整示例

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // 创建测试CPE
    cpe1, _ := cpe.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    cpe2, _ := cpe.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.1:*:*:*:*:*:*:*")
    cpe3, _ := cpe.ParseCpe23("cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*")
    
    fmt.Println("=== 基本匹配测试 ===")
    
    // 基本匹配
    fmt.Printf("cpe1 匹配 cpe2: %t\n", cpe1.Match(cpe2))
    fmt.Printf("cpe1 匹配 cpe3: %t\n", cpe1.Match(cpe3))
    
    fmt.Println("\n=== 模糊匹配测试 ===")
    
    // 模糊匹配
    score12 := cpe.FuzzyMatch(cpe1, cpe2)
    score13 := cpe.FuzzyMatch(cpe1, cpe3)
    
    fmt.Printf("cpe1 与 cpe2 相似度: %.3f\n", score12)
    fmt.Printf("cpe1 与 cpe3 相似度: %.3f\n", score13)
    
    fmt.Println("\n=== 加权匹配测试 ===")
    
    // 加权匹配
    weights := cpe.MatchWeights{
        Part:    0.1,
        Vendor:  0.3,
        Product: 0.4,
        Version: 0.2,
    }
    
    weightedScore := cpe.WeightedMatch(cpe1, cpe2, weights)
    fmt.Printf("加权匹配分数: %.3f\n", weightedScore)
    
    fmt.Println("\n=== 详细匹配结果 ===")
    
    // 详细匹配
    options := cpe.DefaultMatchOptions()
    result := cpe.DetailedMatch(cpe1, cpe2, options)
    
    fmt.Printf("匹配: %t\n", result.Match)
    fmt.Printf("分数: %.3f\n", result.Score)
    fmt.Printf("置信度: %.3f\n", result.Confidence)
    fmt.Printf("说明: %s\n", result.Explanation)
    
    fmt.Println("组件详情:")
    for component, score := range result.Details {
        fmt.Printf("  %s: %.3f\n", component, score)
    }
    
    fmt.Println("\n=== 版本比较测试 ===")
    
    // 版本比较
    versions := []string{"9.0.0", "9.0.1", "9.1.0", "10.0.0"}
    baseVersion := "9.0.0"
    
    for _, version := range versions {
        result := cpe.CompareVersions(baseVersion, version)
        var relation string
        switch result {
        case -1:
            relation = "较旧"
        case 0:
            relation = "相同"
        case 1:
            relation = "较新"
        }
        fmt.Printf("%s 相对于 %s: %s\n", baseVersion, version, relation)
    }
    
    fmt.Println("\n=== 模式匹配测试 ===")
    
    // 模式匹配
    pattern, _ := cpe.ParseCpe23("cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*")
    targets := []*cpe.CPE{cpe1, cpe2}
    
    for i, target := range targets {
        match := cpe.MatchPattern(target, pattern)
        fmt.Printf("目标 %d 匹配Apache模式: %t\n", i+1, match)
    }
}
```

## 下一步

- 了解[存储接口](./storage.md)来持久化匹配结果
- 学习[集合操作](./sets.md)来处理大量CPE匹配
- 探索[NVD集成](./nvd.md)来进行漏洞匹配
