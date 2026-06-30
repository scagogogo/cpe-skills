# 集合操作

本页面描述了CPE库中用于处理CPE集合的功能，包括集合创建、操作、过滤和分析。

## CPE集合

### CPESet

CPE集合的主要结构体。

```go
type CPESet struct {
    items map[string]*CPE // CPE项目映射
    mutex sync.RWMutex    // 读写锁
}
```

### NewCPESet

创建新的空CPE集合。

```go
func NewCPESet() *CPESet
```

### NewCPESetFromSlice

从CPE切片创建集合。

```go
func NewCPESetFromSlice(cpes []*CPE) *CPESet
```

### NewCPESetFromStrings

从CPE字符串列表创建集合。

```go
func NewCPESetFromStrings(cpeStrings []string) *CPESet
```

**示例：**
```go
// 创建空集合
set1 := cpeskills.NewCPESet()

// 从切片创建
cpes := []*cpeskills.CPE{cpe1, cpe2, cpe3}
set2 := cpeskills.NewCPESetFromSlice(cpes)

// 从字符串创建
cpeStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
}
set3 := cpeskills.NewCPESetFromStrings(cpeStrings)

fmt.Printf("集合大小: %d, %d, %d\n", set1.Size(), set2.Size(), set3.Size())
```

## 基本操作

### Add

向集合添加CPE。

```go
func (s *CPESet) Add(cpe *CPE) bool
```

### Remove

从集合移除CPE。

```go
func (s *CPESet) Remove(cpe *CPE) bool
```

### Contains

检查集合是否包含CPE。

```go
func (s *CPESet) Contains(cpe *CPE) bool
```

### Size

获取集合大小。

```go
func (s *CPESet) Size() int
```

### Clear

清空集合。

```go
func (s *CPESet) Clear()
```

**示例：**
```go
set := cpeskills.NewCPESet()

// 添加CPE
cpe1, _ := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
cpe2, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")

added1 := set.Add(cpe1)
added2 := set.Add(cpe2)
added3 := set.Add(cpe1) // 重复添加

fmt.Printf("添加结果: %t, %t, %t\n", added1, added2, added3)
fmt.Printf("集合大小: %d\n", set.Size())

// 检查包含
if set.Contains(cpe1) {
    fmt.Println("集合包含Windows CPE")
}

// 移除CPE
removed := set.Remove(cpe1)
fmt.Printf("移除结果: %t, 新大小: %d\n", removed, set.Size())
```

## 集合运算

### Union

并集操作。

```go
func (s *CPESet) Union(other *CPESet) *CPESet
```

### Intersection

交集操作。

```go
func (s *CPESet) Intersection(other *CPESet) *CPESet
```

### Difference

差集操作。

```go
func (s *CPESet) Difference(other *CPESet) *CPESet
```

### SymmetricDifference

对称差集操作。

```go
func (s *CPESet) SymmetricDifference(other *CPESet) *CPESet
```

**示例：**
```go
// 创建两个集合
set1 := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
})

set2 := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
})

fmt.Printf("集合1大小: %d\n", set1.Size())
fmt.Printf("集合2大小: %d\n", set2.Size())

// 并集
union := set1.Union(set2)
fmt.Printf("并集大小: %d\n", union.Size())

// 交集
intersection := set1.Intersection(set2)
fmt.Printf("交集大小: %d\n", intersection.Size())

// 差集
difference := set1.Difference(set2)
fmt.Printf("差集大小: %d\n", difference.Size())

// 对称差集
symDiff := set1.SymmetricDifference(set2)
fmt.Printf("对称差集大小: %d\n", symDiff.Size())
```

## 集合比较

### Equals

检查两个集合是否相等。

```go
func (s *CPESet) Equals(other *CPESet) bool
```

### IsSubsetOf

检查是否为另一个集合的子集。

```go
func (s *CPESet) IsSubsetOf(other *CPESet) bool
```

### IsSupersetOf

检查是否为另一个集合的超集。

```go
func (s *CPESet) IsSupersetOf(other *CPESet) bool
```

### IsDisjoint

检查两个集合是否不相交。

```go
func (s *CPESet) IsDisjoint(other *CPESet) bool
```

**示例：**
```go
setA := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
})

setB := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
})

setC := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
})

fmt.Printf("A == B: %t\n", setA.Equals(setB))
fmt.Printf("B ⊆ A: %t\n", setB.IsSubsetOf(setA))
fmt.Printf("A ⊇ B: %t\n", setA.IsSupersetOf(setB))
fmt.Printf("A ∩ C = ∅: %t\n", setA.IsDisjoint(setC))
```

## 过滤操作

### Filter

使用自定义函数过滤集合。

```go
func (s *CPESet) Filter(predicate func(*CPE) bool) *CPESet
```

### FilterByVendor

按供应商过滤。

```go
func (s *CPESet) FilterByVendor(vendor string) *CPESet
```

### FilterByProduct

按产品过滤。

```go
func (s *CPESet) FilterByProduct(product string) *CPESet
```

### FilterByPart

按组件类型过滤。

```go
func (s *CPESet) FilterByPart(part string) *CPESet
```

### FilterByVersion

按版本过滤。

```go
func (s *CPESet) FilterByVersion(version string) *CPESet
```

**示例：**
```go
// 创建大集合
largeSet := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
    "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
})

fmt.Printf("原始集合大小: %d\n", largeSet.Size())

// 按供应商过滤
microsoftCPEs := largeSet.FilterByVendor("microsoft")
fmt.Printf("Microsoft CPE数量: %d\n", microsoftCPEs.Size())

// 按组件类型过滤
applications := largeSet.FilterByPart("a")
fmt.Printf("应用程序数量: %d\n", applications.Size())

// 自定义过滤
versionedApps := largeSet.Filter(func(cpe *cpeskills.CPE) bool {
    return cpeskills.Part.ShortName == "a" && cpeskills.Version != "*"
})
fmt.Printf("有版本号的应用程序: %d\n", versionedApps.Size())
```

## 转换操作

### ToSlice

转换为CPE切片。

```go
func (s *CPESet) ToSlice() []*CPE
```

### ToStringSlice

转换为字符串切片。

```go
func (s *CPESet) ToStringSlice() []string
```

### Map

映射转换。

```go
func (s *CPESet) Map(mapper func(*CPE) interface{}) []interface{}
```

**示例：**
```go
set := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
})

// 转换为切片
cpeSlice := set.ToSlice()
fmt.Printf("CPE切片长度: %d\n", len(cpeSlice))

// 转换为字符串切片
stringSlice := set.ToStringSlice()
fmt.Printf("字符串切片: %v\n", stringSlice)

// 映射转换 - 提取供应商名称
vendors := set.Map(func(cpe *cpeskills.CPE) interface{} {
    return cpeskills.Vendor
})
fmt.Printf("供应商列表: %v\n", vendors)
```

## 聚合操作

### GroupBy

按指定字段分组。

```go
func (s *CPESet) GroupBy(keyFunc func(*CPE) string) map[string]*CPESet
```

### CountBy

按指定字段计数。

```go
func (s *CPESet) CountBy(keyFunc func(*CPE) string) map[string]int
```

### ForEach

遍历集合中的每个元素。

```go
func (s *CPESet) ForEach(action func(*CPE))
```

**示例：**
```go
set := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
})

// 按供应商分组
vendorGroups := set.GroupBy(func(cpe *cpeskills.CPE) string {
    return cpeskills.Vendor
})

fmt.Println("按供应商分组:")
for vendor, group := range vendorGroups {
    fmt.Printf("  %s: %d 个CPE\n", vendor, group.Size())
}

// 按组件类型计数
partCounts := set.CountBy(func(cpe *cpeskills.CPE) string {
    return cpeskills.Part.LongName
})

fmt.Println("按组件类型计数:")
for part, count := range partCounts {
    fmt.Printf("  %s: %d\n", part, count)
}

// 遍历操作
fmt.Println("所有CPE:")
set.ForEach(func(cpe *cpeskills.CPE) {
    fmt.Printf("  %s %s\n", cpeskills.Vendor, cpeskills.ProductName)
})
```

## 统计分析

### GetStatistics

获取集合统计信息。

```go
func (s *CPESet) GetStatistics() *SetStatistics
```

### SetStatistics

集合统计信息结构。

```go
type SetStatistics struct {
    TotalCount           int                    // 总数量
    ApplicationCount     int                    // 应用程序数量
    OperatingSystemCount int                    // 操作系统数量
    HardwareCount        int                    // 硬件数量
    UniqueVendors        int                    // 唯一供应商数量
    UniqueProducts       int                    // 唯一产品数量
    VendorDistribution   map[string]int         // 供应商分布
    ProductDistribution  map[string]int         // 产品分布
    VersionDistribution  map[string]int         // 版本分布
}
```

**示例：**
```go
set := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
    "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
})

stats := set.GetStatistics()

fmt.Printf("集合统计信息:\n")
fmt.Printf("  总数量: %d\n", stats.TotalCount)
fmt.Printf("  应用程序: %d\n", stats.ApplicationCount)
fmt.Printf("  操作系统: %d\n", stats.OperatingSystemCount)
fmt.Printf("  硬件: %d\n", stats.HardwareCount)
fmt.Printf("  唯一供应商: %d\n", stats.UniqueVendors)
fmt.Printf("  唯一产品: %d\n", stats.UniqueProducts)

fmt.Println("供应商分布:")
for vendor, count := range stats.VendorDistribution {
    fmt.Printf("  %s: %d\n", vendor, count)
}
```

## 持久化操作

### SaveToFile

保存集合到文件。

```go
func (s *CPESet) SaveToFile(filename string) error
```

### LoadFromFile

从文件加载集合。

```go
func LoadCPESetFromFile(filename string) (*CPESet, error)
```

### ExportToCSV

导出为CSV格式。

```go
func (s *CPESet) ExportToCSV(filename string) error
```

### ExportToJSON

导出为JSON格式。

```go
func (s *CPESet) ExportToJSON(filename string) error
```

**示例：**
```go
set := cpeskills.NewCPESetFromStrings([]string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
})

// 保存到文件
err := set.SaveToFile("cpe_set.json")
if err != nil {
    log.Printf("保存失败: %v", err)
} else {
    fmt.Println("✅ 集合已保存")
}

// 从文件加载
loadedSet, err := cpeskills.LoadCPESetFromFile("cpe_set.json")
if err != nil {
    log.Printf("加载失败: %v", err)
} else {
    fmt.Printf("✅ 加载了 %d 个CPE\n", loadedSet.Size())
}

// 导出为CSV
err = set.ExportToCSV("cpe_set.csv")
if err != nil {
    log.Printf("CSV导出失败: %v", err)
} else {
    fmt.Println("✅ CSV导出成功")
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
    fmt.Println("=== CPE集合操作示例 ===")
    
    // 创建示例数据
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
    }
    
    // 创建集合
    allCPEs := cpeskills.NewCPESetFromStrings(cpeStrings)
    fmt.Printf("创建了包含 %d 个CPE的集合\n", allCPEs.Size())
    
    // 过滤操作
    fmt.Println("\n=== 过滤操作 ===")
    
    microsoftCPEs := allCPEs.FilterByVendor("microsoft")
    fmt.Printf("Microsoft产品: %d 个\n", microsoftCPEs.Size())
    
    applications := allCPEs.FilterByPart("a")
    fmt.Printf("应用程序: %d 个\n", applications.Size())
    
    // 自定义过滤
    webServers := allCPEs.Filter(func(cpe *cpeskills.CPE) bool {
        return cpeskills.ProductName == "tomcat" || cpeskills.ProductName == "http_server"
    })
    fmt.Printf("Web服务器: %d 个\n", webServers.Size())
    
    // 集合运算
    fmt.Println("\n=== 集合运算 ===")
    
    set1 := allCPEs.FilterByVendor("microsoft")
    set2 := allCPEs.FilterByVendor("apache")
    
    fmt.Printf("Microsoft集合大小: %d\n", set1.Size())
    fmt.Printf("Apache集合大小: %d\n", set2.Size())
    
    union := set1.Union(set2)
    fmt.Printf("并集大小: %d\n", union.Size())
    
    intersection := set1.Intersection(set2)
    fmt.Printf("交集大小: %d\n", intersection.Size())
    
    // 分组统计
    fmt.Println("\n=== 分组统计 ===")
    
    vendorGroups := allCPEs.GroupBy(func(cpe *cpeskills.CPE) string {
        return cpeskills.Vendor
    })
    
    fmt.Println("按供应商分组:")
    for vendor, group := range vendorGroups {
        fmt.Printf("  %s: %d 个产品\n", vendor, group.Size())
        
        // 显示该供应商的产品
        group.ForEach(func(cpe *cpeskills.CPE) {
            fmt.Printf("    - %s %s\n", cpeskills.ProductName, cpeskills.Version)
        })
    }
    
    // 统计分析
    fmt.Println("\n=== 统计分析 ===")
    
    stats := allCPEs.GetStatistics()
    fmt.Printf("统计信息:\n")
    fmt.Printf("  总数量: %d\n", stats.TotalCount)
    fmt.Printf("  应用程序: %d\n", stats.ApplicationCount)
    fmt.Printf("  操作系统: %d\n", stats.OperatingSystemCount)
    fmt.Printf("  唯一供应商: %d\n", stats.UniqueVendors)
    fmt.Printf("  唯一产品: %d\n", stats.UniqueProducts)
    
    // 转换操作
    fmt.Println("\n=== 转换操作 ===")
    
    // 提取供应商列表
    vendors := allCPEs.Map(func(cpe *cpeskills.CPE) interface{} {
        return cpeskills.Vendor
    })
    
    // 去重
    uniqueVendors := make(map[string]bool)
    for _, vendor := range vendors {
        uniqueVendors[vendor.(string)] = true
    }
    
    fmt.Printf("唯一供应商列表: ")
    for vendor := range uniqueVendors {
        fmt.Printf("%s ", vendor)
    }
    fmt.Println()
    
    // 持久化
    fmt.Println("\n=== 持久化操作 ===")
    
    // 保存为JSON
    err := allCPEs.SaveToFile("all_cpes.json")
    if err != nil {
        log.Printf("保存失败: %v", err)
    } else {
        fmt.Println("✅ 集合已保存为JSON")
    }
    
    // 导出为CSV
    err = allCPEs.ExportToCSV("all_cpes.csv")
    if err != nil {
        log.Printf("CSV导出失败: %v", err)
    } else {
        fmt.Println("✅ 集合已导出为CSV")
    }
    
    // 从文件加载验证
    loadedSet, err := cpeskills.LoadCPESetFromFile("all_cpes.json")
    if err != nil {
        log.Printf("加载失败: %v", err)
    } else {
        fmt.Printf("✅ 从文件加载了 %d 个CPE\n", loadedSet.Size())
        
        // 验证数据一致性
        if allCPEs.Equals(loadedSet) {
            fmt.Println("✅ 数据一致性验证通过")
        } else {
            fmt.Println("❌ 数据一致性验证失败")
        }
    }
}
```

## 下一步

- 了解[存储接口](./storage.md)来持久化大型CPE集合
- 学习[匹配算法](./matching.md)来在集合中查找匹配项
- 探索[NVD集成](./nvd.md)来处理大规模CPE数据集
