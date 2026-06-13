# CPE集合

本示例演示如何使用CPE集合功能处理CPE对象集合，进行高效的批量操作。

## 概述

CPE集合提供了一种强大的方式来管理CPE对象集合，执行集合操作（并集、交集、差集），并应用批量转换和过滤器。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    fmt.Println("=== CPE集合示例 ===")
    
    // 示例1：创建CPE集合
    fmt.Println("\n1. 创建CPE集合:")
    
    // 创建单个CPE对象
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
    }
    
    // 方法1：从字符串创建集合
    set1 := cpe.NewCPESetFromStrings(cpeStrings)
    fmt.Printf("集合1大小: %d\n", set1.Size())
    
    // 方法2：创建空集合并添加项目
    set2 := cpe.NewCPESet()
    for _, cpeStr := range cpeStrings[:3] { // 添加前3个项目
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("解析%s失败: %v", cpeStr, err)
            continue
        }
        set2.Add(cpeObj)
    }
    fmt.Printf("集合2大小: %d\n", set2.Size())
    
    // 方法3：从CPE对象切片创建
    cpeObjects := make([]*cpe.CPE, 0, len(cpeStrings))
    for _, cpeStr := range cpeStrings[2:] { // 添加后3个项目
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            continue
        }
        cpeObjects = append(cpeObjects, cpeObj)
    }
    set3 := cpe.NewCPESetFromSlice(cpeObjects)
    fmt.Printf("集合3大小: %d\n", set3.Size())
    
    // 示例2：集合操作
    fmt.Println("\n2. 集合操作:")
    
    fmt.Println("集合1内容:")
    set1.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    fmt.Println("集合2内容:")
    set2.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    fmt.Println("集合3内容:")
    set3.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 并集：两个集合中的所有唯一项目
    unionSet := set2.Union(set3)
    fmt.Printf("\n集合2和集合3的并集 (大小: %d):\n", unionSet.Size())
    unionSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 交集：两个集合中都存在的项目
    intersectionSet := set1.Intersection(set2)
    fmt.Printf("\n集合1和集合2的交集 (大小: %d):\n", intersectionSet.Size())
    intersectionSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 差集：第一个集合中有但第二个集合中没有的项目
    differenceSet := set1.Difference(set2)
    fmt.Printf("\n集合1 - 集合2的差集 (大小: %d):\n", differenceSet.Size())
    differenceSet.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 示例3：过滤集合
    fmt.Println("\n3. 过滤集合:")
    
    // 为过滤示例创建更大的集合
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
    fmt.Printf("大集合大小: %d\n", largeSet.Size())
    
    // 按供应商过滤
    microsoftCPEs := largeSet.FilterByVendor("microsoft")
    fmt.Printf("\nMicrosoft CPE (大小: %d):\n", microsoftCPEs.Size())
    microsoftCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 按部件过滤（仅应用程序）
    applicationCPEs := largeSet.FilterByPart("a")
    fmt.Printf("\n应用程序CPE (大小: %d):\n", applicationCPEs.Size())
    applicationCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 按产品模式过滤
    apacheCPEs := largeSet.FilterByProduct("apache")
    fmt.Printf("\nApache CPE (大小: %d):\n", apacheCPEs.Size())
    apacheCPEs.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s\n", cpe.GetURI())
    })
    
    // 自定义过滤函数
    customFilter := func(cpe *cpe.CPE) bool {
        // 过滤有版本信息的应用程序
        return cpe.Part.ShortName == "a" && cpe.Version != "*" && cpe.Version != ""
    }
    
    versionedApps := largeSet.Filter(customFilter)
    fmt.Printf("\n有版本的应用程序 (大小: %d):\n", versionedApps.Size())
    versionedApps.ForEach(func(cpe *cpe.CPE) {
        fmt.Printf("  - %s (v%s)\n", cpe.ProductName, cpe.Version)
    })
    
    // 示例4：集合转换
    fmt.Println("\n4. 集合转换:")
    
    // 转换以提取供应商信息
    vendors := largeSet.Map(func(cpe *cpe.CPE) string {
        return cpe.Vendor
    })
    
    uniqueVendors := removeDuplicateStrings(vendors)
    fmt.Printf("唯一供应商: %v\n", uniqueVendors)
    
    // 转换以创建摘要信息
    summaries := largeSet.Map(func(cpe *cpe.CPE) string {
        return fmt.Sprintf("%s %s %s", cpe.Vendor, cpe.ProductName, cpe.Version)
    })
    
    fmt.Println("\nCPE摘要:")
    for i, summary := range summaries {
        fmt.Printf("  %d. %s\n", i+1, summary)
    }
    
    // 示例5：集合聚合
    fmt.Println("\n5. 集合聚合:")
    
    // 按供应商分组
    vendorGroups := largeSet.GroupBy(func(cpe *cpe.CPE) string {
        return cpe.Vendor
    })
    
    fmt.Println("按供应商分组的CPE:")
    for vendor, cpes := range vendorGroups {
        fmt.Printf("  %s (%d项):\n", vendor, len(cpes))
        for _, cpe := range cpes {
            fmt.Printf("    - %s\n", cpe.ProductName)
        }
    }
    
    // 按部件类型分组
    partGroups := largeSet.GroupBy(func(cpe *cpe.CPE) string {
        return cpe.Part.LongName
    })
    
    fmt.Println("\n按部件类型分组的CPE:")
    for partType, cpes := range partGroups {
        fmt.Printf("  %s: %d项\n", partType, len(cpes))
    }
    
    // 示例6：集合统计
    fmt.Println("\n6. 集合统计:")
    
    stats := largeSet.GetStatistics()
    fmt.Printf("集合统计:\n")
    fmt.Printf("  总CPE数: %d\n", stats.TotalCount)
    fmt.Printf("  应用程序: %d\n", stats.ApplicationCount)
    fmt.Printf("  操作系统: %d\n", stats.OperatingSystemCount)
    fmt.Printf("  硬件: %d\n", stats.HardwareCount)
    fmt.Printf("  唯一供应商: %d\n", stats.UniqueVendors)
    fmt.Printf("  唯一产品: %d\n", stats.UniqueProducts)
    
    // 示例7：集合持久化
    fmt.Println("\n7. 集合持久化:")
    
    // 保存集合到文件
    filename := "cpe_set_export.json"
    err := largeSet.SaveToFile(filename)
    if err != nil {
        log.Printf("保存集合失败: %v", err)
    } else {
        fmt.Printf("集合已保存到 %s\n", filename)
    }
    
    // 从文件加载集合
    loadedSet, err := cpe.LoadCPESetFromFile(filename)
    if err != nil {
        log.Printf("加载集合失败: %v", err)
    } else {
        fmt.Printf("从 %s 加载集合 (大小: %d)\n", filename, loadedSet.Size())
        
        // 验证加载的集合与原始集合匹配
        if loadedSet.Size() == largeSet.Size() {
            fmt.Println("✅ 加载的集合大小与原始匹配")
        } else {
            fmt.Println("❌ 加载的集合大小与原始不同")
        }
    }
    
    // 示例8：集合比较
    fmt.Println("\n8. 集合比较:")
    
    // 创建两个相似的集合
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
    
    fmt.Printf("集合A大小: %d\n", setA.Size())
    fmt.Printf("集合B大小: %d\n", setB.Size())
    
    // 检查相等性
    areEqual := setA.Equals(setB)
    fmt.Printf("集合相等: %t\n", areEqual)
    
    // 检查一个是否是另一个的子集
    isSubset := setA.IsSubsetOf(setB)
    fmt.Printf("集合A是集合B的子集: %t\n", isSubset)
    
    // 查找公共元素
    common := setA.Intersection(setB)
    fmt.Printf("公共元素: %d\n", common.Size())
    
    // 查找每个集合中的唯一元素
    uniqueA := setA.Difference(setB)
    uniqueB := setB.Difference(setA)
    
    fmt.Printf("集合A独有: %d\n", uniqueA.Size())
    fmt.Printf("集合B独有: %d\n", uniqueB.Size())
}

// 辅助函数去除重复字符串
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

## 关键概念

### 1. 集合创建

- **从字符串**: 直接将CPE字符串解析为集合
- **从对象**: 从现有CPE对象创建集合
- **空集合**: 从空集合开始并添加项目

### 2. 集合操作

- **并集**: 合并两个集合 (A ∪ B)
- **交集**: 公共元素 (A ∩ B)
- **差集**: A中有但B中没有的元素 (A - B)

### 3. 过滤和转换

- **过滤**: 根据条件选择子集
- **映射**: 转换每个元素
- **分组**: 按键组织元素

### 4. 集合分析

- **统计**: 按类型计算元素
- **比较**: 检查相等性和子集关系
- **聚合**: 汇总集合内容

## 最佳实践

1. **批量操作使用集合**: 比单个操作更高效
2. **早期过滤**: 在昂贵操作前应用过滤器以减少集合大小
3. **缓存结果**: 存储经常使用的过滤集合
4. **验证输入**: 添加到集合前检查CPE有效性
5. **监控内存**: 大集合可能消耗大量内存

## 性能提示

1. **批量操作**: 将多个操作组合在一起
2. **使用适当的数据结构**: 集合针对唯一性进行了优化
3. **并行处理**: 对独立的集合操作使用goroutine
4. **延迟评估**: 延迟昂贵操作直到需要时

## 下一步

- 学习[高级匹配](./advanced-matching.md)与集合
- 探索[存储](./storage.md)来持久化大集合
- 查看[NVD集成](./nvd-integration.md)了解实际数据集
