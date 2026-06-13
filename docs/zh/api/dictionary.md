# 字典管理

本页面描述了CPE库中用于管理CPE字典的功能，包括字典的加载、搜索、验证和维护。

## CPE字典结构

### CPEDictionary

CPE字典的主要结构体。

```go
type CPEDictionary struct {
    Entries      []*CPEDictionaryEntry // 字典条目列表
    LastModified time.Time             // 最后修改时间
    Version      string                // 字典版本
    Generator    string                // 生成器信息
    Timestamp    time.Time             // 时间戳
}
```

### CPEDictionaryEntry

字典中的单个条目。

```go
type CPEDictionaryEntry struct {
    CPE23        string    // CPE 2.3格式字符串
    CPE22        string    // CPE 2.2格式字符串（如果有）
    Title        string    // 条目标题
    References   []string  // 参考链接
    LastModified time.Time // 最后修改时间
    Deprecated   bool      // 是否已弃用
    DeprecatedBy string    // 被什么替代
}
```

## 字典创建和加载

### NewCPEDictionary

创建新的CPE字典。

```go
func NewCPEDictionary() *CPEDictionary
```

### LoadDictionaryFromFile

从文件加载CPE字典。

```go
func LoadDictionaryFromFile(filename string) (*CPEDictionary, error)
```

**支持的文件格式：**
- XML (NVD官方格式)
- JSON
- CSV

**示例：**
```go
// 从XML文件加载
dict, err := cpe.LoadDictionaryFromFile("official-cpe-dictionary_v2.3.xml")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("加载了 %d 个CPE条目\n", len(dict.Entries))
```

### LoadDictionaryFromURL

从URL加载CPE字典。

```go
func LoadDictionaryFromURL(url string) (*CPEDictionary, error)
```

**示例：**
```go
// 从NVD官方URL加载
nvdURL := "https://nvd.nist.gov/feeds/xml/cpe/dictionary/official-cpe-dictionary_v2.3.xml.gz"
dict, err := cpe.LoadDictionaryFromURL(nvdURL)
if err != nil {
    log.Printf("从URL加载失败: %v", err)
}
```

## 字典操作

### AddEntry

向字典添加条目。

```go
func (d *CPEDictionary) AddEntry(entry *CPEDictionaryEntry) error
```

### RemoveEntry

从字典移除条目。

```go
func (d *CPEDictionary) RemoveEntry(cpe23 string) error
```

### UpdateEntry

更新字典条目。

```go
func (d *CPEDictionary) UpdateEntry(cpe23 string, entry *CPEDictionaryEntry) error
```

**示例：**
```go
// 创建新条目
entry := &cpe.CPEDictionaryEntry{
    CPE23:        "cpe:2.3:a:example:product:1.0:*:*:*:*:*:*:*",
    Title:        "Example Product 1.0",
    References:   []string{"https://example.com/product"},
    LastModified: time.Now(),
}

// 添加到字典
err := dict.AddEntry(entry)
if err != nil {
    log.Printf("添加条目失败: %v", err)
}
```

## 字典搜索

### Search

在字典中搜索CPE条目。

```go
func (d *CPEDictionary) Search(query string, limit int) []*CPEDictionaryEntry
```

**参数：**
- `query`: 搜索查询字符串
- `limit`: 返回结果的最大数量

**示例：**
```go
// 搜索Microsoft相关的CPE
results := dict.Search("microsoft", 10)
fmt.Printf("找到 %d 个Microsoft相关的CPE:\n", len(results))

for i, entry := range results {
    fmt.Printf("  %d. %s\n", i+1, entry.Title)
    fmt.Printf("     %s\n", entry.CPE23)
}
```

### SearchByVendor

按供应商搜索。

```go
func (d *CPEDictionary) SearchByVendor(vendor string) []*CPEDictionaryEntry
```

### SearchByProduct

按产品搜索。

```go
func (d *CPEDictionary) SearchByProduct(product string) []*CPEDictionaryEntry
```

### SearchByPattern

使用模式搜索。

```go
func (d *CPEDictionary) SearchByPattern(pattern string) []*CPEDictionaryEntry
```

**示例：**
```go
// 搜索所有Apache产品
apacheEntries := dict.SearchByVendor("apache")
fmt.Printf("Apache产品数量: %d\n", len(apacheEntries))

// 搜索Tomcat产品
tomcatEntries := dict.SearchByProduct("tomcat")
fmt.Printf("Tomcat相关条目: %d\n", len(tomcatEntries))

// 使用通配符模式搜索
javaEntries := dict.SearchByPattern("cpe:2.3:a:*:*java*:*:*:*:*:*:*:*:*")
fmt.Printf("Java相关条目: %d\n", len(javaEntries))
```

## 字典验证

### ValidateCPE

验证CPE是否在字典中。

```go
func (d *CPEDictionary) ValidateCPE(cpe23 string) bool
```

### GetEntry

获取特定的字典条目。

```go
func (d *CPEDictionary) GetEntry(cpe23 string) *CPEDictionaryEntry
```

### FindSimilar

查找相似的CPE条目。

```go
func (d *CPEDictionary) FindSimilar(cpe23 string, threshold float64) []*CPEDictionaryEntry
```

**示例：**
```go
// 验证CPE是否存在
cpeString := "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
if dict.ValidateCPE(cpeString) {
    fmt.Println("CPE在官方字典中")
    
    // 获取详细信息
    entry := dict.GetEntry(cpeString)
    if entry != nil {
        fmt.Printf("标题: %s\n", entry.Title)
        fmt.Printf("最后修改: %s\n", entry.LastModified.Format("2006-01-02"))
    }
} else {
    fmt.Println("CPE不在官方字典中")
    
    // 查找相似的条目
    similar := dict.FindSimilar(cpeString, 0.8)
    if len(similar) > 0 {
        fmt.Printf("找到 %d 个相似条目:\n", len(similar))
        for i, entry := range similar {
            fmt.Printf("  %d. %s\n", i+1, entry.CPE23)
        }
    }
}
```

## 字典统计

### GetStatistics

获取字典统计信息。

```go
func (d *CPEDictionary) GetStatistics() *DictionaryStats
```

### DictionaryStats

字典统计信息结构。

```go
type DictionaryStats struct {
    TotalEntries      int                    // 总条目数
    ApplicationCount  int                    // 应用程序数量
    OperatingSystemCount int                 // 操作系统数量
    HardwareCount     int                    // 硬件数量
    VendorCount       int                    // 供应商数量
    DeprecatedCount   int                    // 已弃用条目数量
    VendorDistribution map[string]int        // 供应商分布
    LastUpdate        time.Time              // 最后更新时间
}
```

**示例：**
```go
stats := dict.GetStatistics()

fmt.Printf("字典统计信息:\n")
fmt.Printf("  总条目数: %d\n", stats.TotalEntries)
fmt.Printf("  应用程序: %d\n", stats.ApplicationCount)
fmt.Printf("  操作系统: %d\n", stats.OperatingSystemCount)
fmt.Printf("  硬件设备: %d\n", stats.HardwareCount)
fmt.Printf("  供应商数量: %d\n", stats.VendorCount)
fmt.Printf("  已弃用条目: %d\n", stats.DeprecatedCount)

fmt.Println("\n主要供应商分布:")
for vendor, count := range stats.VendorDistribution {
    if count > 100 { // 只显示条目数超过100的供应商
        fmt.Printf("  %s: %d\n", vendor, count)
    }
}
```

## 字典导出

### SaveToFile

将字典保存到文件。

```go
func (d *CPEDictionary) SaveToFile(filename string, format string) error
```

**支持的格式：**
- "xml": XML格式
- "json": JSON格式
- "csv": CSV格式

### ExportToXML

导出为XML格式。

```go
func (d *CPEDictionary) ExportToXML() ([]byte, error)
```

### ExportToJSON

导出为JSON格式。

```go
func (d *CPEDictionary) ExportToJSON() ([]byte, error)
```

**示例：**
```go
// 保存为JSON格式
err := dict.SaveToFile("cpe_dictionary.json", "json")
if err != nil {
    log.Printf("保存失败: %v", err)
} else {
    fmt.Println("字典已保存为JSON格式")
}

// 导出为XML字节数组
xmlData, err := dict.ExportToXML()
if err != nil {
    log.Printf("导出XML失败: %v", err)
} else {
    fmt.Printf("XML数据大小: %d 字节\n", len(xmlData))
}
```

## 字典合并

### Merge

合并两个字典。

```go
func (d *CPEDictionary) Merge(other *CPEDictionary) error
```

### MergeWithConflictResolution

带冲突解决的字典合并。

```go
func (d *CPEDictionary) MergeWithConflictResolution(
    other *CPEDictionary, 
    resolver ConflictResolver,
) error
```

### ConflictResolver

冲突解决器接口。

```go
type ConflictResolver interface {
    Resolve(existing, incoming *CPEDictionaryEntry) *CPEDictionaryEntry
}
```

**示例：**
```go
// 简单合并
dict1, _ := cpe.LoadDictionaryFromFile("dict1.xml")
dict2, _ := cpe.LoadDictionaryFromFile("dict2.xml")

err := dict1.Merge(dict2)
if err != nil {
    log.Printf("合并失败: %v", err)
} else {
    fmt.Printf("合并后条目数: %d\n", len(dict1.Entries))
}

// 带冲突解决的合并
resolver := &cpe.LatestWinsResolver{} // 使用最新的条目
err = dict1.MergeWithConflictResolution(dict2, resolver)
```

## 字典更新

### UpdateFromNVD

从NVD更新字典。

```go
func (d *CPEDictionary) UpdateFromNVD() error
```

### CheckForUpdates

检查是否有可用更新。

```go
func (d *CPEDictionary) CheckForUpdates() (bool, error)
```

### GetUpdateInfo

获取更新信息。

```go
func (d *CPEDictionary) GetUpdateInfo() (*UpdateInfo, error)
```

**示例：**
```go
// 检查更新
hasUpdates, err := dict.CheckForUpdates()
if err != nil {
    log.Printf("检查更新失败: %v", err)
} else if hasUpdates {
    fmt.Println("发现可用更新")
    
    // 获取更新信息
    updateInfo, err := dict.GetUpdateInfo()
    if err == nil {
        fmt.Printf("新版本: %s\n", updateInfo.NewVersion)
        fmt.Printf("发布日期: %s\n", updateInfo.ReleaseDate.Format("2006-01-02"))
    }
    
    // 执行更新
    err = dict.UpdateFromNVD()
    if err != nil {
        log.Printf("更新失败: %v", err)
    } else {
        fmt.Println("字典更新成功")
    }
} else {
    fmt.Println("字典已是最新版本")
}
```

## 字典索引

### BuildIndex

构建字典索引以提高搜索性能。

```go
func (d *CPEDictionary) BuildIndex() error
```

### RebuildIndex

重建索引。

```go
func (d *CPEDictionary) RebuildIndex() error
```

### IndexStats

获取索引统计信息。

```go
func (d *CPEDictionary) IndexStats() *IndexStatistics
```

**示例：**
```go
// 构建索引
fmt.Println("构建字典索引...")
err := dict.BuildIndex()
if err != nil {
    log.Printf("构建索引失败: %v", err)
} else {
    fmt.Println("索引构建完成")
    
    // 获取索引统计
    indexStats := dict.IndexStats()
    fmt.Printf("索引大小: %d KB\n", indexStats.SizeKB)
    fmt.Printf("索引条目: %d\n", indexStats.EntryCount)
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
    fmt.Println("=== CPE字典管理示例 ===")
    
    // 创建新字典
    dict := cpe.NewCPEDictionary()
    
    // 添加示例条目
    entries := []*cpe.CPEDictionaryEntry{
        {
            CPE23:        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
            Title:        "Microsoft Windows 10",
            References:   []string{"https://www.microsoft.com/windows"},
            LastModified: time.Now(),
        },
        {
            CPE23:        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
            Title:        "Apache Tomcat 9.0.0",
            References:   []string{"https://tomcat.apache.org/"},
            LastModified: time.Now(),
        },
        {
            CPE23:        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
            Title:        "Oracle Java SE 11.0.12",
            References:   []string{"https://www.oracle.com/java/"},
            LastModified: time.Now(),
        },
    }
    
    // 批量添加条目
    for _, entry := range entries {
        err := dict.AddEntry(entry)
        if err != nil {
            log.Printf("添加条目失败: %v", err)
        }
    }
    
    fmt.Printf("字典中有 %d 个条目\n", len(dict.Entries))
    
    // 搜索示例
    fmt.Println("\n=== 搜索示例 ===")
    
    // 按关键词搜索
    results := dict.Search("microsoft", 5)
    fmt.Printf("搜索'microsoft'找到 %d 个结果:\n", len(results))
    for i, entry := range results {
        fmt.Printf("  %d. %s\n", i+1, entry.Title)
    }
    
    // 按供应商搜索
    apacheEntries := dict.SearchByVendor("apache")
    fmt.Printf("\nApache产品数量: %d\n", len(apacheEntries))
    
    // 验证CPE
    fmt.Println("\n=== 验证示例 ===")
    testCPE := "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
    if dict.ValidateCPE(testCPE) {
        fmt.Printf("✅ %s 在字典中\n", testCPE)
        
        entry := dict.GetEntry(testCPE)
        if entry != nil {
            fmt.Printf("   标题: %s\n", entry.Title)
            fmt.Printf("   参考: %v\n", entry.References)
        }
    } else {
        fmt.Printf("❌ %s 不在字典中\n", testCPE)
    }
    
    // 获取统计信息
    fmt.Println("\n=== 统计信息 ===")
    stats := dict.GetStatistics()
    fmt.Printf("总条目数: %d\n", stats.TotalEntries)
    fmt.Printf("应用程序: %d\n", stats.ApplicationCount)
    fmt.Printf("供应商数量: %d\n", stats.VendorCount)
    
    // 导出字典
    fmt.Println("\n=== 导出示例 ===")
    err := dict.SaveToFile("example_dictionary.json", "json")
    if err != nil {
        log.Printf("保存失败: %v", err)
    } else {
        fmt.Println("✅ 字典已保存为JSON格式")
    }
    
    // 构建索引
    fmt.Println("\n=== 索引示例 ===")
    err = dict.BuildIndex()
    if err != nil {
        log.Printf("构建索引失败: %v", err)
    } else {
        fmt.Println("✅ 索引构建完成")
        
        indexStats := dict.IndexStats()
        fmt.Printf("索引条目数: %d\n", indexStats.EntryCount)
    }
}
```

## 下一步

- 了解[NVD集成](./nvd.md)来获取官方CPE字典
- 学习[验证功能](./validation.md)来确保字典数据质量
- 探索[存储接口](./storage.md)来持久化字典数据
