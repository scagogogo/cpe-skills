# CPE - Common Platform Enumeration Library

<div align="center">

![Version](https://img.shields.io/badge/version-1.0.0-blue.svg)
![License](https://img.shields.io/badge/license-MIT-green.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.18-blue.svg)

**[English](#english) | [简体中文](#简体中文)**

</div>

---

## English

### 📚 Documentation

**Complete API documentation and examples: [https://scagogogo.github.io/cpe/](https://scagogogo.github.io/cpe/)**

- [API Reference](https://scagogogo.github.io/cpe/api/) - Complete API documentation
- [Examples](https://scagogogo.github.io/cpe/examples/) - Practical code examples
- [Quick Start Guide](https://scagogogo.github.io/cpe/api/) - Getting started tutorial

### 📖 Introduction

The CPE (Common Platform Enumeration) library is a comprehensive Go implementation for processing, parsing, matching, and storing CPE (Common Platform Enumeration) data. CPE is a structured naming scheme for identifying classes of IT systems, software, and packages.

The library also includes integration with CVE (Common Vulnerabilities and Exposures), enabling developers to associate software components with known security vulnerabilities.

### ✨ Features

- **CPE Format Support**: Parse and generate CPE 2.2 and 2.3 formats
- **Advanced Matching**: CPE name matching with wildcards and special values
- **WFN Support**: Well-Formed Name format with bidirectional conversion
- **Applicability Language**: CPE Applicability Language support
- **Version Comparison**: Semantic version comparison and range matching
- **Dictionary Management**: CPE dictionary with XML import/export
- **CVE Integration**: Associate CPEs with Common Vulnerabilities and Exposures
- **Advanced Algorithms**: Fuzzy matching, subset/superset matching
- **Set Operations**: Union, intersection, difference operations on CPE collections
- **NVD Integration**: Built-in National Vulnerability Database feed integration
- **Error Handling**: Structured error handling with detailed error types
- **Storage Backends**: Multiple storage backends with persistence support
- **Caching**: Integrated caching mechanism for optimized performance

## 🏗️ 系统架构

### 核心组件

<div align="center">
  <img src="https://via.placeholder.com/800x400?text=CPE+Library+Architecture" alt="CPE库架构图" width="80%"/>
</div>

CPE库采用模块化设计，主要由以下几个核心组件构成：

1. **CPE解析引擎**：负责解析和格式化CPE字符串，支持CPE 2.2和2.3标准
   - 字符串解析器：将CPE URI转换为内部数据结构
   - 格式化器：将内部数据结构转换为标准CPE字符串

2. **匹配引擎**：实现CPE匹配逻辑，支持多种匹配策略
   - 基础匹配：精确匹配和通配符匹配
   - 高级匹配：正则表达式、模糊匹配和距离计算
   - 版本比较：语义化版本比较和版本范围检查

3. **存储系统**：提供多种存储后端选项
   - 内存存储：适用于临时数据和高性能场景
   - 文件存储：持久化数据到本地文件系统
   - 可扩展接口：允许实现自定义存储后端

4. **CVE集成模块**：连接CPE和漏洞信息
   - CVE引用管理：创建和维护CVE与CPE的关联
   - 漏洞查询：根据产品信息查询相关漏洞
   - 文本分析：从非结构化文本中提取CVE标识符

5. **数据源适配器**：连接外部数据源
   - NVD适配器：与美国国家漏洞数据库集成
   - 厂商适配器：与软件供应商漏洞数据源集成
   - 通用REST API适配器：连接自定义漏洞数据源

### 数据流

CPE库中的数据流遵循以下路径：

1. **输入处理**：通过解析器将外部CPE字符串转换为内部数据结构
2. **数据操作**：使用匹配引擎和表达式语言处理CPE数据
3. **持久化**：通过存储系统保存和检索CPE数据
4. **漏洞关联**：利用CVE集成模块关联漏洞信息
5. **数据聚合**：通过多源搜索功能整合来自不同数据源的信息

### 接口设计

库的接口设计遵循以下原则：

- **一致性**：所有组件使用一致的接口约定
- **模块化**：每个组件都是独立的，可以单独使用
- **可扩展性**：核心接口支持自定义实现
- **简单性**：公共API简洁明了，易于使用

## ✨ 特性

- 完整支持CPE 2.2和CPE 2.3格式
- 高级匹配功能，包括正则表达式和模糊匹配
- 内置版本比较功能
- 表达式语言用于复杂的适用性语句
- 多种存储选项（内存、文件）
- 与NVD数据源集成
- CVE关联和查询功能
- 可扩展的数据源架构

### 🚀 Installation

Install using Go modules:

```bash
go get github.com/scagogogo/cpe
```

### 🔍 Quick Start

### 基本使用

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cpe"
)

func main() {
    // 解析CPE 2.3字符串
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
if err != nil {
        panic(err)
    }
    
    fmt.Printf("CPE详情: 供应商=%s, 产品=%s, 版本=%s\n", 
               cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
               
    // 创建匹配条件
    criteria := &cpe.CPE{
        Vendor: "microsoft",
        ProductName: "windows",
    }
    
    // 执行匹配
    if cpeObj.Match(criteria) {
        fmt.Println("匹配成功!")
    }
}
```

### 使用CVE功能

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cpe"
)

func main() {
    // 从文本中提取CVE ID
    text := "系统受到CVE-2021-44228和CVE-2022-22965漏洞的影响"
    cveIDs := cpe.ExtractCVEsFromText(text)
    fmt.Printf("发现CVE: %v\n", cveIDs)
    
    // 按年份分组
    grouped := cpe.GroupCVEsByYear(cveIDs)
    fmt.Printf("按年份分组: %v\n", grouped)
    
    // 创建CVE引用
    cveRef := cpe.NewCVEReference("CVE-2021-44228")
    cveRef.Description = "Log4j远程代码执行漏洞"
    cveRef.SetSeverity(10.0) // Critical
    
    // 添加受影响的CPE
    cveRef.AddAffectedCPE("cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*")
}
```

## 📚 API 文档

<details open>
<summary><b>CPE 相关功能</b></summary>

### 解析与格式化

#### `ParseCpe23(cpe23 string) (*CPE, error)`

解析CPE 2.3格式字符串并转换为CPE结构体。

```go
cpe, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
```

#### `ParseCpe22(cpe22 string) (*CPE, error)`

解析CPE 2.2格式字符串并转换为CPE结构体。

```go
cpe, err := cpe.ParseCpe22("cpe:/a:microsoft:windows:10")
```

#### `FormatCpe23(cpe *CPE) string`

将CPE对象格式化为CPE 2.3字符串。

```go
cpeString := cpe.FormatCpe23(cpeObj)
```

#### `FormatCpe22(cpe *CPE) string`

将CPE对象格式化为CPE 2.2字符串。

```go
cpeString := cpe.FormatCpe22(cpeObj)
```

### 匹配功能

#### `Match(other *CPE) bool`

检查CPE是否与给定的CPE匹配。

```go
if cpe1.Match(cpe2) {
    fmt.Println("匹配成功")
}
```

#### `MatchCPE(criteria *CPE, target *CPE, options *MatchOptions) bool`

高级CPE匹配功能，支持自定义匹配选项。

```go
options := cpe.DefaultMatchOptions()
options.IgnoreVersion = true
if cpe.MatchCPE(criteria, target, options) {
    fmt.Println("匹配成功")
}
```

#### `AdvancedMatchCPE(criteria *CPE, target *CPE, options *AdvancedMatchOptions) bool`

最灵活的CPE匹配功能，支持高级选项如正则表达式、模糊匹配等。

```go
options := cpe.NewAdvancedMatchOptions()
options.UseRegex = true
options.IgnoreCase = true
if cpe.AdvancedMatchCPE(criteria, target, options) {
    fmt.Println("匹配成功")
}
```

### 版本比较

#### `compareVersions(criteria *CPE, target *CPE, options *AdvancedMatchOptions) bool`

比较两个CPE的版本。

```go
options := cpe.NewAdvancedMatchOptions()
options.VersionCompareMode = "greater"
options.VersionLower = "2.0"
result := cpe.compareVersions(cpe1, cpe2, options)
```

#### `compareVersionStrings(v1, v2 string) int`

比较两个版本字符串，返回-1 (v1 < v2)、0 (v1 == v2) 或 1 (v1 > v2)。

```go
result := cpe.compareVersionStrings("1.2.3", "1.3.0")
if result < 0 {
    fmt.Println("v1 < v2")
}
```

</details>

<details open>
<summary><b>CVE 相关功能</b></summary>

### CVE引用

#### `NewCVEReference(cveID string) *CVEReference`

创建一个新的CVE引用。

```go
cveRef := cpe.NewCVEReference("CVE-2021-44228")
```

#### `AddAffectedCPE(cpeURI string)`

向CVE引用添加受影响的CPE。

```go
cveRef.AddAffectedCPE("cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*")
```

#### `RemoveAffectedCPE(cpeURI string) bool`

从CVE引用中移除受影响的CPE。

```go
removed := cveRef.RemoveAffectedCPE("cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*")
```

#### `AddReference(reference string)`

添加参考链接到CVE引用。

```go
cveRef.AddReference("https://nvd.nist.gov/vuln/detail/CVE-2021-44228")
```

#### `SetSeverity(cvssScore float64)`

设置CVE的CVSS评分和对应的严重性级别。

```go
cveRef.SetSeverity(9.8) // 设置为Critical级别
```

#### `SetMetadata(key string, value interface{})`

设置CVE的元数据。

```go
cveRef.SetMetadata("exploitAvailable", true)
```

#### `GetMetadata(key string) (interface{}, bool)`

获取CVE的元数据。

```go
value, exists := cveRef.GetMetadata("exploitAvailable")
```

#### `RemoveMetadata(key string) bool`

移除CVE的元数据。

```go
removed := cveRef.RemoveMetadata("exploitAvailable")
```

### CVE查询与处理

#### `QueryByCVE(cves []*CVEReference, cveID string) []*CPE`

根据CVE ID查询关联的CPE。

```go
cpes := cpe.QueryByCVE(cveList, "CVE-2021-44228")
```

#### `GetCVEInfo(cves []*CVEReference, cveID string) *CVEReference`

获取CVE的详细信息。

```go
cveInfo := cpe.GetCVEInfo(cveList, "CVE-2021-44228")
```

#### `ExtractCVEsFromText(text string) []string`

从文本中提取CVE ID。

```go
cveIDs := cpe.ExtractCVEsFromText("系统受到CVE-2021-44228影响")
```

#### `GroupCVEsByYear(cveIDs []string) map[string][]string`

按年份对CVE ID进行分组。

```go
grouped := cpe.GroupCVEsByYear(cveIDs)
```

#### `SortCVEs(cveIDs []string) []string`

对CVE ID列表进行排序。

```go
sorted := cpe.SortCVEs(cveIDs)
```

#### `RemoveDuplicateCVEs(cveIDs []string) []string`

去除CVE ID列表中的重复项。

```go
unique := cpe.RemoveDuplicateCVEs(cveIDs)
```

#### `GetRecentCVEs(cveIDs []string, years int) []string`

获取最近N年的CVE ID。

```go
recent := cpe.GetRecentCVEs(cveIDs, 2) // 获取最近2年的CVE
```

#### `ValidateCVE(cveID string) bool`

验证CVE ID是否有效。

```go
isValid := cpe.ValidateCVE("CVE-2021-44228")
```

#### `QueryByProduct(cves []*CVEReference, vendor, product, version string) []*CVEReference`

根据产品信息查询相关CVE。

```go
results := cpe.QueryByProduct(cveList, "apache", "log4j", "2.0")
```

</details>

<details open>
<summary><b>存储相关功能</b></summary>

### 内存存储

#### `NewMemoryStorage() *MemoryStorage`

创建一个新的内存存储实例。

```go
storage := cpe.NewMemoryStorage()
err := storage.Initialize()
```

#### `StoreCPE(cpe *CPE) error`

存储CPE到内存。

```go
err := storage.StoreCPE(cpeObj)
```

#### `RetrieveCPE(id string) (*CPE, error)`

从内存检索CPE。

```go
cpe, err := storage.RetrieveCPE("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
```

### 文件存储

#### `NewFileStorage(baseDir string, useCache bool) (*FileStorage, error)`

创建一个新的文件存储实例。

```go
storage, err := cpe.NewFileStorage("./cpe_data", true)
err = storage.Initialize()
```

#### `StoreCPE(cpe *CPE) error`

存储CPE到文件系统。

```go
err := storage.StoreCPE(cpeObj)
```

#### `RetrieveCPE(id string) (*CPE, error)`

从文件系统检索CPE。

```go
cpe, err := storage.RetrieveCPE("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
```

### 通用存储接口

所有存储实现都兼容Storage接口，可以互换使用。

```go
var storage cpe.Storage
storage = cpe.NewMemoryStorage()
// 或
storage, _ = cpe.NewFileStorage("./cpe_data", true)

// 使用通用接口操作
err := storage.Initialize()
err = storage.StoreCPE(cpeObj)
cpe, err := storage.RetrieveCPE(cpeID)
```

</details>

<details open>
<summary><b>集合与过滤</b></summary>

### CPE集合

#### `NewCPESet(name string, description string) *CPESet`

创建一个新的CPE集合。

```go
set := cpe.NewCPESet("Windows产品", "微软Windows系列产品")
```

#### `Add(cpe *CPE)`

向集合中添加CPE。

```go
set.Add(cpeObj)
```

#### `Remove(cpe *CPE) bool`

从集合中移除CPE。

```go
removed := set.Remove(cpeObj)
```

#### `Contains(cpe *CPE) bool`

检查集合是否包含指定CPE。

```go
if set.Contains(cpeObj) {
    fmt.Println("集合包含该CPE")
}
```

#### `Size() int`

返回集合大小。

```go
count := set.Size()
```

#### `Filter(criteria *CPE, options *MatchOptions) *CPESet`

根据条件过滤集合。

```go
criteria := &cpe.CPE{Vendor: "microsoft"}
options := cpe.DefaultMatchOptions()
filteredSet := set.Filter(criteria, options)
```

#### `Union(other *CPESet) *CPESet`

计算两个集合的并集。

```go
unionSet := set1.Union(set2)
```

#### `Intersection(other *CPESet) *CPESet`

计算两个集合的交集。

```go
intersectionSet := set1.Intersection(set2)
```

#### `Difference(other *CPESet) *CPESet`

计算两个集合的差集。

```go
differenceSet := set1.Difference(set2)
```

</details>

<details open>
<summary><b>适用性语言</b></summary>

### 表达式

#### `ParseExpression(expr string) (Expression, error)`

解析适用性表达式。

```go
expr, err := cpe.ParseExpression("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
```

#### `FilterCPEs(cpes []*CPE, expr Expression) []*CPE`

使用表达式过滤CPE列表。

```go
filteredCPEs := cpe.FilterCPEs(cpeList, expr)
```

### 表达式类型

- `CPEExpression` - 匹配单个CPE
- `ANDExpression` - 匹配所有子表达式
- `ORExpression` - 匹配任一子表达式
- `NOTExpression` - 反转子表达式的匹配结果

```go
// AND表达式示例
expr, _ := cpe.ParseExpression("AND(cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*, cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*)")

// OR表达式示例
expr, _ := cpe.ParseExpression("OR(cpe:2.3:a:microsoft:edge:*:*:*:*:*:*:*:*, cpe:2.3:a:google:chrome:*:*:*:*:*:*:*:*)")

// NOT表达式示例
expr, _ := cpe.ParseExpression("NOT(cpe:2.3:a:microsoft:edge:*:*:*:*:*:*:*:*)")
```

</details>

<details open>
<summary><b>NVD集成</b></summary>

### NVD数据源

#### `DefaultNVDFeedOptions() *NVDFeedOptions`

创建默认的NVD Feed下载选项。

```go
options := cpe.DefaultNVDFeedOptions()
options.CacheDir = "/tmp/nvd-cache"
```

#### `DownloadAndParseCPEDict(options *NVDFeedOptions) (*CPEDictionary, error)`

下载并解析NVD CPE字典。

```go
dict, err := cpe.DownloadAndParseCPEDict(options)
```

#### `DownloadAndParseCPEMatch(options *NVDFeedOptions) (*CPEMatchData, error)`

下载并解析NVD CPE Match数据。

```go
match, err := cpe.DownloadAndParseCPEMatch(options)
```

#### `DownloadAllNVDData(options *NVDFeedOptions) (*NVDCPEData, error)`

下载所有NVD数据。

```go
data, err := cpe.DownloadAllNVDData(options)
```

### NVD数据查询

#### `FindCVEsForCPE(cpe *CPE) []string`

查找与特定CPE相关的所有CVE。

```go
cves := nvdData.FindCVEsForCPE(cpeObj)
```

#### `FindCPEsForCVE(cveID string) []*CPE`

查找与特定CVE相关的所有CPE。

```go
cpes := nvdData.FindCPEsForCVE("CVE-2021-44228")
```

</details>

<details open>
<summary><b>数据源集成</b></summary>

### 数据源

#### `NewDataSource(sourceType DataSourceType, name, description, url string) *DataSource`

创建新的数据源。

```go
ds := cpe.NewDataSource(cpe.DataSourceNVD, "NVD", "National Vulnerability Database", "https://services.nvd.nist.gov/rest/json/")
```

#### `CreateNVDDataSource(apiKey string) *DataSource`

创建NVD数据源。

```go
nvd := cpe.CreateNVDDataSource("YOUR_API_KEY")
```

#### `CreateGitHubDataSource(token string) *DataSource`

创建GitHub数据源。

```go
github := cpe.CreateGitHubDataSource("YOUR_GITHUB_TOKEN")
```

#### `CreateRedHatDataSource() *DataSource`

创建RedHat数据源。

```go
redhat := cpe.CreateRedHatDataSource()
```

### 多源搜索

#### `NewMultiSourceSearch(sources []*DataSource) *MultiSourceVulnerabilitySearch`

创建新的多数据源搜索。

```go
sources := []*cpe.DataSource{nvd, github, redhat}
search := cpe.NewMultiSourceSearch(sources)
```

#### `SearchByCVE(cveID string) ([]*CVEReference, error)`

根据CVE ID在多个数据源中搜索。

```go
results, err := search.SearchByCVE("CVE-2021-44228")
```

#### `SearchByCPE(cpe *CPE) ([]*CVEReference, error)`

根据CPE在多个数据源中搜索。

```go
results, err := search.SearchByCPE(cpeObj)
```

</details>

## 🔆 高级使用示例

以下是一些完整的示例，展示如何在实际场景中结合使用库的多个功能。

<details open>
<summary><b>示例1: 漏洞扫描系统</b></summary>

该示例展示如何创建一个简单的漏洞扫描系统，用于检测给定软件组件列表中的潜在安全漏洞。

```go
package main

import (
    "fmt"
    "log"
    "time"
    
    "github.com/scagogogo/cpe"
)

func main() {
    // 1. 初始化NVD数据源
    options := cpe.DefaultNVDFeedOptions()
    options.CacheDir = "./nvd-cache"
    options.MaxAge = 24 * time.Hour // 每24小时更新一次
    
    log.Println("正在下载NVD数据...")
    nvdData, err := cpe.DownloadAllNVDData(options)
if err != nil {
        log.Fatalf("无法获取NVD数据: %v", err)
    }
    log.Println("NVD数据下载完成")
    
    // 2. 定义要扫描的软件清单
    softwareInventory := []string{
        "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:openssl:openssl:1.0.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:google:chrome:90.0.4430.85:*:*:*:*:*:*:*",
    }
    
    // 3. 扫描每个软件组件的漏洞
    log.Println("开始扫描漏洞...")
    var criticalVulnerabilities []*cpe.CVEReference
    
    for _, cpeStr := range softwareInventory {
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("无法解析CPE: %s, 错误: %v", cpeStr, err)
            continue
        }
        
        // 查找相关的CVE
        relatedCVEs := nvdData.FindCVEsForCPE(cpeObj)
        if len(relatedCVEs) > 0 {
            fmt.Printf("\n发现组件 '%s: %s %s' 存在 %d 个潜在漏洞\n", 
                cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version, len(relatedCVEs))
            
            // 获取每个CVE的详细信息
            for _, cveID := range relatedCVEs {
                cveInfo := nvdData.GetCVEDetails(cveID)
                
                // 评估风险级别
                if cveInfo.CVSS >= 7.0 {
                    criticalVulnerabilities = append(criticalVulnerabilities, cveInfo)
                    fmt.Printf("  [高危] %s (CVSS: %.1f) - %s\n", 
                        cveInfo.ID, cveInfo.CVSS, cveInfo.Description)
                } else if cveInfo.CVSS >= 4.0 {
                    fmt.Printf("  [中危] %s (CVSS: %.1f) - %s\n", 
                        cveInfo.ID, cveInfo.CVSS, cveInfo.Description)
                }
            }
        } else {
            fmt.Printf("组件 '%s: %s %s' 未发现已知漏洞\n", 
                cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
        }
    }
    
    // 4. 生成总结报告
    fmt.Printf("\n========== 漏洞扫描总结 ==========\n")
    fmt.Printf("扫描组件总数: %d\n", len(softwareInventory))
    fmt.Printf("发现高危漏洞: %d\n", len(criticalVulnerabilities))
    
    if len(criticalVulnerabilities) > 0 {
        fmt.Println("\n建议优先修复以下组件:")
        for _, cve := range criticalVulnerabilities {
            fmt.Printf("  - %s (影响: %s)\n", cve.ID, cve.AffectedProducts)
        }
    }
}
```
</details>

<details open>
<summary><b>示例2: 软件资产管理</b></summary>

该示例展示如何使用库实现软件资产清单管理。

```go
package main

import (
    "fmt"
    "log"
    "os"
    "time"
    
    "github.com/scagogogo/cpe"
)

// 资产类型枚举
const (
    AssetTypeServer     = "SERVER"
    AssetTypeWorkstation = "WORKSTATION"
    AssetTypeNetwork    = "NETWORK"
    AssetTypeApplication = "APPLICATION"
)

// 资产信息
type Asset struct {
    CPE         *cpe.CPE
    AssetType   string
    Location    string
    Owner       string
    InstallDate time.Time
    Notes       string
}

// 资产管理器
type AssetManager struct {
    assets      map[string]*Asset
    storage     cpe.Storage
}

// 创建资产管理器
func NewAssetManager(storageDir string) (*AssetManager, error) {
// 初始化文件存储
    storage, err := cpe.NewFileStorage(storageDir, true)
if err != nil {
        return nil, fmt.Errorf("初始化存储失败: %v", err)
    }
    
    if err := storage.Initialize(); err != nil {
        return nil, fmt.Errorf("存储初始化失败: %v", err)
    }
    
    return &AssetManager{
        assets:  make(map[string]*Asset),
        storage: storage,
    }, nil
}

// 添加新资产
func (am *AssetManager) AddAsset(cpeStr, assetType, location, owner, notes string) error {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
if err != nil {
        return fmt.Errorf("解析CPE失败: %v", err)
    }
    
    // 创建资产对象
    asset := &Asset{
        CPE:         cpeObj,
        AssetType:   assetType,
        Location:    location,
        Owner:       owner,
        InstallDate: time.Now(),
        Notes:       notes,
    }
    
    // 存储CPE信息
    if err := am.storage.StoreCPE(cpeObj); err != nil {
        return fmt.Errorf("存储CPE失败: %v", err)
    }
    
    // 保存资产信息
    assetID := cpeObj.GetURI()
    am.assets[assetID] = asset
    
    return nil
}

// 查找特定供应商的所有资产
func (am *AssetManager) FindAssetsByVendor(vendor string) []*Asset {
    var results []*Asset
    
    for _, asset := range am.assets {
        if asset.CPE.Vendor == vendor {
            results = append(results, asset)
        }
    }
    
    return results
}

// 生成资产报告
func (am *AssetManager) GenerateReport() {
    fmt.Println("=========== 软件资产清单报告 ===========")
    fmt.Printf("总资产数量: %d\n\n", len(am.assets))
    
    // 按资产类型分组
    assetsByType := make(map[string][]*Asset)
    for _, asset := range am.assets {
        assetsByType[asset.AssetType] = append(assetsByType[asset.AssetType], asset)
    }
    
    // 打印分组信息
    for assetType, assets := range assetsByType {
        fmt.Printf("== %s (%d) ==\n", assetType, len(assets))
        
        for _, asset := range assets {
            cpe := asset.CPE
            fmt.Printf("  - %s %s %s\n", cpe.Vendor, cpe.ProductName, cpe.Version)
            fmt.Printf("    位置: %s, 负责人: %s\n", asset.Location, asset.Owner)
            if asset.Notes != "" {
                fmt.Printf("    备注: %s\n", asset.Notes)
            }
            fmt.Println()
        }
    }
}

func main() {
    // 创建资产管理器
    assetManager, err := NewAssetManager("./asset-data")
if err != nil {
        log.Fatalf("创建资产管理器失败: %v", err)
    }
    
    // 添加资产
    if err := assetManager.AddAsset(
        "cpe:2.3:a:microsoft:windows_server:2019:*:*:*:*:*:*:*",
        AssetTypeServer,
        "北京数据中心",
        "系统运维组",
        "主域控制器",
    ); err != nil {
        log.Printf("添加资产失败: %v", err)
    }
    
    if err := assetManager.AddAsset(
        "cpe:2.3:a:apache:tomcat:9.0.50:*:*:*:*:*:*:*",
        AssetTypeApplication,
        "应用服务器01",
        "应用运维组",
        "电商网站后端",
    ); err != nil {
        log.Printf("添加资产失败: %v", err)
    }
    
    if err := assetManager.AddAsset(
        "cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*",
        AssetTypeWorkstation,
        "财务部",
        "IT支持组",
        "标准办公软件",
    ); err != nil {
        log.Printf("添加资产失败: %v", err)
    }
    
    // 生成资产报告
    assetManager.GenerateReport()
    
    // 查找特定供应商的资产
    microsoftAssets := assetManager.FindAssetsByVendor("microsoft")
    fmt.Printf("\n发现 %d 个微软资产:\n", len(microsoftAssets))
    for _, asset := range microsoftAssets {
        fmt.Printf("  - %s %s (%s)\n", 
            asset.CPE.ProductName, asset.CPE.Version, asset.AssetType)
    }
}
```
</details>

<details open>
<summary><b>示例3: CVE 文本分析</b></summary>

该示例展示如何从非结构化文本中提取并分析CVE信息。

```go
package main

import (
    "fmt"
    "log"
    "strings"
    "time"
    
    "github.com/scagogogo/cpe"
)

// 安全公告结构
type SecurityBulletin struct {
    Title     string
    Content   string
    Published time.Time
    Source    string
}

// CVE分析器
type CVEAnalyzer struct {
    nvdData    *cpe.NVDCPEData
    cveDetails map[string]*cpe.CVEReference
}

// 创建CVE分析器
func NewCVEAnalyzer() (*CVEAnalyzer, error) {
    options := cpe.DefaultNVDFeedOptions()
    options.CacheDir = "./nvd-cache"
    
    // 下载NVD数据
    nvdData, err := cpe.DownloadAllNVDData(options)
if err != nil {
        return nil, fmt.Errorf("下载NVD数据失败: %v", err)
    }
    
    return &CVEAnalyzer{
        nvdData:    nvdData,
        cveDetails: make(map[string]*cpe.CVEReference),
    }, nil
}

// 分析安全公告
func (ca *CVEAnalyzer) AnalyzeBulletin(bulletin SecurityBulletin) map[string]interface{} {
    result := make(map[string]interface{})
    
    // 提取CVE ID
    cveIDs := cpe.ExtractCVEsFromText(bulletin.Title + " " + bulletin.Content)
    
    // 如果没有找到CVE ID，返回空结果
    if len(cveIDs) == 0 {
        result["found_cves"] = false
        return result
    }
    
    // 排序并去重
    uniqueCVEs := cpe.SortAndRemoveDuplicateCVEs(cveIDs)
    
    result["found_cves"] = true
    result["cve_count"] = len(uniqueCVEs)
    result["cve_ids"] = uniqueCVEs
    
    // 按年份分组
    cvesByYear := cpe.GroupCVEsByYear(uniqueCVEs)
    result["cves_by_year"] = cvesByYear
    
    // 提取每个CVE的详细信息
    cveDetails := make(map[string]map[string]interface{})
    var highRiskCVEs []string
    
    for _, cveID := range uniqueCVEs {
        // 获取CVE详情
        cveInfo := ca.nvdData.GetCVEDetails(cveID)
        if cveInfo == nil {
            // 如果NVD数据中没有，创建一个基本引用
            cveInfo = cpe.NewCVEReference(cveID)
        }
        
        // 保存详情以供后续使用
        ca.cveDetails[cveID] = cveInfo
        
        // 提取关键信息
        cveDetail := make(map[string]interface{})
        cveDetail["description"] = cveInfo.Description
        cveDetail["cvss_score"] = cveInfo.CVSS
        cveDetail["severity"] = cveInfo.Severity
        cveDetail["affected_cpes"] = cveInfo.AffectedCPEs
        
        // 检查是否高风险
        if cveInfo.CVSS >= 7.0 {
            highRiskCVEs = append(highRiskCVEs, cveID)
        }
        
        cveDetails[cveID] = cveDetail
    }
    
    result["cve_details"] = cveDetails
    result["high_risk_cves"] = highRiskCVEs
    
    return result
}

// 打印分析结果
func printAnalysisResult(bulletin SecurityBulletin, result map[string]interface{}) {
    fmt.Printf("======= 安全公告分析 =======\n")
    fmt.Printf("标题: %s\n", bulletin.Title)
    fmt.Printf("来源: %s\n", bulletin.Source)
    fmt.Printf("发布时间: %s\n\n", bulletin.Published.Format("2006-01-02"))
    
    if !result["found_cves"].(bool) {
        fmt.Println("未发现CVE标识符")
        return
    }
    
    cveCount := result["cve_count"].(int)
    cveIDs := result["cve_ids"].([]string)
    
    fmt.Printf("发现 %d 个CVE:\n", cveCount)
    for _, id := range cveIDs {
        detail := result["cve_details"].(map[string]map[string]interface{})[id]
        
        severity := "未知"
        if s, ok := detail["severity"]; ok && s != nil {
            severity = s.(string)
        }
        
        cvssScore := 0.0
        if s, ok := detail["cvss_score"]; ok && s != nil {
            cvssScore = s.(float64)
        }
        
        description := "无描述"
        if d, ok := detail["description"]; ok && d != nil {
            description = d.(string)
            if len(description) > 100 {
                description = description[:97] + "..."
            }
        }
        
        fmt.Printf("  - %s [%s, CVSS: %.1f]\n    %s\n", 
            id, severity, cvssScore, description)
    }
    
    // 显示高风险CVE
    if highRisk, ok := result["high_risk_cves"].([]string); ok && len(highRisk) > 0 {
        fmt.Printf("\n高风险漏洞 (%d):\n", len(highRisk))
        for _, id := range highRisk {
            fmt.Printf("  - %s\n", id)
        }
    }
    
    // 显示年份分布
    if yearData, ok := result["cves_by_year"].(map[string][]string); ok {
        fmt.Println("\nCVE年份分布:")
        for year, cves := range yearData {
            fmt.Printf("  %s: %d个\n", year, len(cves))
        }
    }
}

func main() {
    // 初始化CVE分析器
    analyzer, err := NewCVEAnalyzer()
if err != nil {
        log.Fatalf("初始化CVE分析器失败: %v", err)
    }
    
    // 模拟一些安全公告
    bulletins := []SecurityBulletin{
        {
            Title:     "Microsoft发布6月安全更新修复多个高危漏洞",
            Content:   "微软在最新的安全更新中修复了多个严重漏洞，包括CVE-2023-35311和CVE-2023-32046等。这些漏洞可能导致远程代码执行。建议用户尽快更新系统。",
            Published: time.Date(2023, 6, 14, 0, 0, 0, 0, time.UTC),
            Source:    "Microsoft Security",
        },
        {
            Title:     "Apache Log4j远程代码执行漏洞警告",
            Content:   "Log4j存在严重的远程代码执行漏洞(CVE-2021-44228),影响版本2.0到2.14.1。攻击者可以利用该漏洞执行任意代码。建议立即升级到2.15.0或更高版本。",
            Published: time.Date(2021, 12, 10, 0, 0, 0, 0, time.UTC),
            Source:    "Apache Foundation",
        },
    }
    
    // 分析所有公告
    for _, bulletin := range bulletins {
        result := analyzer.AnalyzeBulletin(bulletin)
        printAnalysisResult(bulletin, result)
        fmt.Println("\n-------------------------------\n")
    }
}
```
</details>

## 📊 使用场景

- 软件组件分析 (SCA)
- 漏洞管理系统
- 供应链安全
- 合规检查
- 资产清单管理
- 安全产品集成

## 🛠️ 最佳实践

<details open>
<summary><b>性能优化</b></summary>

### 缓存管理

* **合理设置缓存过期时间**：NVD和CPE数据量较大，合理设置缓存可以显著提高性能。
```go
  options := cpe.DefaultNVDFeedOptions()
  options.CacheDir = "/app/cache"
  options.MaxAge = 24 * time.Hour // 数据每天更新一次
  ```

* **使用内存缓存**：对于频繁访问的数据，优先使用内存存储。
  ```go
  // 创建带缓存的文件存储
  storage, _ := cpe.NewFileStorage("./data", true) // 第二个参数启用缓存
  ```

* **预加载常用数据**：对于频繁使用的CPE数据，可以在应用启动时预加载。
  ```go
  // 应用启动时预加载
  commonCPEs := []string{
      "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
      "cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*",
  }
  
  for _, cpeStr := range commonCPEs {
      cpe, _ := cpe.ParseCpe23(cpeStr)
      storage.StoreCPE(cpe) // 预先存储到缓存
  }
  ```

### 查询优化

* **使用精确查询**：在可能的情况下，使用更精确的查询条件减少结果集大小。
```go
  // 不推荐
  criteria := &cpe.CPE{Vendor: "microsoft"}
  
  // 推荐
criteria := &cpe.CPE{
      Vendor: "microsoft",
    ProductName: "windows",
}
  ```

* **批量处理**：处理大量CPE数据时，使用批处理而非单个处理。
  ```go
  // 批量处理示例
  processBatch := func(cpes []*cpe.CPE, batchSize int) {
      total := len(cpes)
      for i := 0; i < total; i += batchSize {
          end := i + batchSize
          if end > total {
              end = total
          }
          batch := cpes[i:end]
          // 处理批次
      }
  }
  ```

* **使用索引**：如果实现自定义存储，为常查询字段添加索引。
  ```go
  // 自定义存储时添加索引示例
  type CustomStorage struct {
      data        map[string]*cpe.CPE
      vendorIndex map[string][]string // 厂商 -> CPE ID 列表
  }
  ```

### 内存管理

* **限制结果集大小**：处理大量数据时设置合理的结果集上限。
  ```go
  // 设置最大结果数
  const maxResults = 1000
  if len(results) > maxResults {
      results = results[:maxResults]
  }
  ```

* **流式处理**：对于非常大的数据集，使用流式处理避免一次性加载全部内容。
```go
  // 使用回调函数处理大量结果
  searchWithCallback := func(criteria *cpe.CPE, callback func(*cpe.CPE) bool) {
      // 搜索实现
      // 对每个结果调用callback
      // 如果callback返回false则停止处理
  }
  ```

* **资源释放**：确保正确关闭存储和释放资源。
  ```go
  storage, _ := cpe.NewFileStorage("./data", true)
  defer storage.Close()
  ```
</details>

<details open>
<summary><b>安全建议</b></summary>

### 数据验证

* **验证外部输入**：处理用户输入的CPE或CVE字符串时进行验证。
```go
  // 验证CVE ID
  if !cpe.ValidateCVE(userInput) {
      return errors.New("无效的CVE ID")
  }
  
  // 验证CPE字符串
  _, err := cpe.ParseCpe23(userInput)
if err != nil {
      return fmt.Errorf("无效的CPE: %v", err)
  }
  ```

* **强制类型检查**：使用类型断言时添加安全检查。
  ```go
  value, ok := metadata["key"].(string)
  if !ok {
      return errors.New("类型错误")
  }
  ```

### 错误处理

* **详细记录错误**：记录详细的错误信息便于调试和审计。
```go
if err != nil {
      log.Printf("解析CPE失败: %v, 输入: %s", err, input)
      return nil, err
  }
  ```

* **有意义的错误返回**：返回描述性错误信息。
  ```go
  if len(cveID) < 13 {
      return fmt.Errorf("CVE ID '%s' 格式无效: 长度不足", cveID)
  }
  ```

### 数据源安全

* **控制API密钥**：安全存储和管理NVD API密钥。
  ```go
  // 从环境变量获取API密钥
  apiKey := os.Getenv("NVD_API_KEY")
  dataSource := cpe.CreateNVDDataSource(apiKey)
  ```

* **验证数据源**：验证下载的NVD数据的完整性。
  ```go
  // 验证数据哈希
  if !cpe.VerifyFeedIntegrity(data, expectedHash) {
      return errors.New("数据完整性检查失败")
  }
  ```
</details>

<details open>
<summary><b>集成建议</b></summary>

### 与现有系统集成

* **使用适配器模式**：创建适配器连接第三方系统。
```go
  // CMDB适配器示例
  type CMDBAdapter struct {
      client CMDBClient
  }
  
  func (a *CMDBAdapter) ImportFromCMDB() ([]*cpe.CPE, error) {
      // 从CMDB导入资产并转换为CPE
  }
  
  func (a *CMDBAdapter) ExportToCMDB(cpes []*cpe.CPE) error {
      // 将CPE导出到CMDB
  }
  ```

* **实现标准接口**：确保自定义组件实现库定义的接口。
  ```go
  // 实现Storage接口
  type DatabaseStorage struct {
      // ...
  }
  
  func (ds *DatabaseStorage) Initialize() error { /* ... */ }
  func (ds *DatabaseStorage) StoreCPE(cpe *CPE) error { /* ... */ }
  func (ds *DatabaseStorage) RetrieveCPE(id string) (*CPE, error) { /* ... */ }
  // 实现其他接口方法...
  ```

### 分布式部署

* **无状态设计**：设计组件时避免依赖共享状态。
```go
  // 创建独立服务
  type CPEService struct {
      storage cpe.Storage
  }
  
  // 方法可以独立调用，不依赖服务状态
  func (s *CPEService) MatchCPE(cpeStr1, cpeStr2 string) (bool, error) {
      cpe1, err := cpe.ParseCpe23(cpeStr1)
      if err != nil {
          return false, err
      }
      
      cpe2, err := cpe.ParseCpe23(cpeStr2)
if err != nil {
          return false, err
      }
      
      return cpe1.Match(cpe2), nil
  }
  ```

* **考虑共享缓存**：在微服务架构中使用共享缓存提高性能。
  ```go
  // 使用Redis作为共享缓存
  type RedisStorage struct {
      client RedisClient
      ttl    time.Duration
  }
  ```
</details>

## ❓ 常见问题 (FAQ)

<details open>
<summary><b>CPE格式问题</b></summary>

### 如何选择使用CPE 2.2还是CPE 2.3格式？

**答**: CPE 2.3是较新的标准，提供更丰富的表示能力，建议优先使用。但如果需要与只支持2.2格式的系统集成，库提供了转换功能：

```go
// 2.2转为2.3
cpe22, _ := cpe.ParseCpe22("cpe:/a:microsoft:windows:10")
cpe23Str := cpe.FormatCpe23(cpe22)

// 2.3转为2.2
cpe23, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
cpe22Str := cpe.FormatCpe22(cpe23)
```

### CPE中的特殊值(如*和-)有什么区别？

**答**: 在CPE中，特殊值有不同的含义：
- `*` (任意值): 表示该属性可以是任何值，用于模糊匹配
- `-` (NA值): 表示该属性不适用于当前组件
- `?` (未知值): 表示该属性的值未知

```go
// *表示任意Windows版本
cpeAny, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*")

// -表示该产品没有版本概念
cpeNA, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:defender:-:*:*:*:*:*:*:*")
```

### 如何处理包含特殊字符的产品名称？

**答**: CPE标准对特殊字符有转义规则，库会自动处理这些字符：

```go
// 包含特殊字符的产品名
cpe, _ := cpe.ParseCpe23("cpe:2.3:a:jquery:jquery\\.ui:1.12.1:*:*:*:*:*:*:*")
fmt.Println(cpe.ProductName) // 输出: jquery.ui (已自动转义)

// 创建对象时自动处理特殊字符
newCpe := &cpe.CPE{
    Part: *cpe.PartApplication,
    Vendor: "node.js",  // 包含点号
    ProductName: "express/connect", // 包含斜杠
}
// 转为URI时会自动转义：cpe:2.3:a:node\.js:express\/connect:*:*:*:*:*:*:*:*
```
</details>

<details open>
<summary><b>匹配与比较问题</b></summary>

### 如何实现版本范围匹配？

**答**: 使用高级匹配选项可以实现版本范围匹配：

```go
options := cpe.NewAdvancedMatchOptions()
options.VersionCompareMode = "range"
options.VersionLower = "1.0.0"
options.VersionUpper = "2.0.0"

criteria := &cpe.CPE{
    Vendor: "apache",
    ProductName: "log4j",
}

target, _ := cpe.ParseCpe23("cpe:2.3:a:apache:log4j:1.5.0:*:*:*:*:*:*:*")
isMatch := cpe.AdvancedMatchCPE(criteria, target, options) // 返回true，1.5.0在范围内
```

### 为什么我的正则表达式匹配不工作？

**答**: 确保在匹配选项中启用了正则表达式，并使用正确的正则语法：

```go
options := cpe.NewAdvancedMatchOptions()
options.UseRegex = true // 必须启用正则

criteria := &cpe.CPE{
    Vendor: "apache",
    ProductName: "log[0-9]j", // 正则表达式
}

// 将匹配log4j
target, _ := cpe.ParseCpe23("cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*")
isMatch := cpe.AdvancedMatchCPE(criteria, target, options)
```

### 如何实现忽略某些字段的匹配？

**答**: 使用匹配选项可以配置忽略特定字段：

```go
options := cpe.NewAdvancedMatchOptions()
options.IgnoreFields = []string{"version", "update"}

criteria := &cpe.CPE{
    Vendor: "microsoft",
    ProductName: "windows",
    Version: "10", // 会被忽略
}

target, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:11:*:*:*:*:*:*:*")
isMatch := cpe.AdvancedMatchCPE(criteria, target, options) // 返回true，因为忽略版本
```
</details>

<details open>
<summary><b>性能相关问题</b></summary>

### 如何高效处理大量CPE数据？

**答**: 处理大量CPE数据时的建议：

1. 使用批处理：每次处理一批CPE，避免一次加载全部数据
2. 使用并发处理：利用Go的并发特性分批并行处理
3. 启用缓存：对频繁访问的CPE启用缓存
4. 使用索引：在存储实现中为常查询字段建立索引

```go
// 并发处理示例
func ProcessCPEsConcurrently(cpes []*cpe.CPE, concurrency int, processor func(*cpe.CPE) error) error {
    total := len(cpes)
    if total == 0 {
        return nil
    }
    
    // 控制并发数
    semaphore := make(chan struct{}, concurrency)
    errChan := make(chan error, total)
    
    // 并发处理每个CPE
    for _, c := range cpes {
        semaphore <- struct{}{} // 获取槽位
        go func(cpe *cpe.CPE) {
            defer func() { <-semaphore }() // 释放槽位
            err := processor(cpe)
            if err != nil {
                errChan <- err
            } else {
                errChan <- nil
            }
        }(c)
    }
    
    // 收集错误
    for i := 0; i < total; i++ {
        if err := <-errChan; err != nil {
            return err
        }
    }
    
    return nil
}
```

### NVD数据下载很慢，有什么优化方法？

**答**: 优化NVD数据下载的建议：

1. 使用缓存：设置合理的缓存过期时间，避免频繁下载
2. 增量更新：只下载自上次更新以来的新数据
3. 考虑代理：如果网络环境限制，可以使用代理服务器
4. 本地镜像：对于大型部署，考虑建立NVD数据的本地镜像

```go
// 设置下载选项
options := cpe.DefaultNVDFeedOptions()
options.CacheDir = "./nvd-cache"
options.MaxAge = 24 * time.Hour
options.Proxy = "http://your-proxy:8080" // 如果需要代理
options.UserAgent = "YourApp/1.0"
options.Timeout = 5 * time.Minute // 较长的超时时间

// 检查是否需要更新
if !cpe.NeedsUpdate("nvdcpematch", options) {
    // 使用缓存数据
    data, _ := cpe.LoadCachedFeed("nvdcpematch", options)
    // ...
} else {
    // 需要更新，下载新数据
    nvdData, _ := cpe.DownloadAllNVDData(options)
    // ...
}
```
</details>

<details open>
<summary><b>集成问题</b></summary>

### 如何将库集成到现有资产管理系统？

**答**: 集成到现有系统的步骤：

1. 创建适配器：实现将系统资产数据转换为CPE的适配器
2. 映射字段：将系统中的厂商、产品、版本等字段映射到CPE属性
3. 实现双向同步：确保CPE变更可以反映到系统，反之亦然
4. 使用事件机制：为重要操作实现事件通知

```go
// 资产管理系统适配器示例
type AssetSystemAdapter struct {
    client    AssetSystemClient
    converter FieldConverter
}

// 转换系统资产为CPE
func (a *AssetSystemAdapter) ConvertToCPE(asset Asset) (*cpe.CPE, error) {
    return &cpe.CPE{
        Part:        *cpe.PartApplication,
        Vendor:      a.converter.MapVendor(asset.Manufacturer),
        ProductName: a.converter.MapProduct(asset.ProductName),
        Version:     a.converter.MapVersion(asset.Version),
        // 映射其他字段...
    }, nil
}

// 转换CPE为系统资产
func (a *AssetSystemAdapter) ConvertToAsset(c *cpe.CPE) (Asset, error) {
    return Asset{
        Manufacturer: a.converter.ReverseMapVendor(c.Vendor),
        ProductName:  a.converter.ReverseMapProduct(c.ProductName),
        Version:      c.Version,
        // 映射其他字段...
    }, nil
}
```

### 我需要实现自定义存储，有什么建议？

**答**: 实现自定义存储的建议：

1. 实现Storage接口：确保实现所有required接口方法
2. 考虑性能：针对查询模式优化存储结构
3. 添加错误处理：所有操作都返回明确的错误
4. 支持批处理：实现批量操作以提高性能
5. 实现事务支持：支持原子操作和回滚

```go
// 自定义数据库存储示例
type DatabaseStorage struct {
    db        *sql.DB
    tableName string
}

func NewDatabaseStorage(dsn string, tableName string) (*DatabaseStorage, error) {
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }
    
    return &DatabaseStorage{
        db:        db,
        tableName: tableName,
    }, nil
}

// 实现Storage接口方法
func (ds *DatabaseStorage) Initialize() error {
    // 创建表和索引
    query := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s (
        id VARCHAR(255) PRIMARY KEY,
        vendor VARCHAR(100) NOT NULL,
        product VARCHAR(100) NOT NULL,
        version VARCHAR(50),
        data TEXT NOT NULL,
        INDEX idx_vendor (vendor),
        INDEX idx_product (product),
        INDEX idx_version (version)
    )`, ds.tableName)
    
    _, err := ds.db.Exec(query)
    return err
}

func (ds *DatabaseStorage) StoreCPE(c *cpe.CPE) error {
    // 实现存储逻辑
    // ...
}

func (ds *DatabaseStorage) RetrieveCPE(id string) (*cpe.CPE, error) {
    // 实现检索逻辑
    // ...
}

// 实现其他接口方法...
```
</details>

### 代码贡献

1. **核心功能改进**: 优化匹配算法，提高性能
2. **新存储实现**: 添加更多存储后端支持
3. **缺陷修复**: 修复已知问题和改进错误处理

### 文档贡献

1. **使用案例**: 贡献更多的实际使用场景和案例
2. **教程**: 编写入门教程和深入指南
3. **API文档**: 改进和扩展API文档

### 测试贡献

1. **单元测试**: 增加测试覆盖率
2. **基准测试**: 创建性能基准和比较
3. **集成测试**: 添加与外部系统集成的测试

### 贡献流程

1. 查看 [Issues](https://github.com/scagogogo/cpe/issues) 中的待处理任务
2. Fork仓库并创建您的特性分支
3. 提交更改并确保测试通过
4. 推送到您的分支并提交Pull Request
5. 等待代码审查和合并

</details>

## 📄 开源协议

本项目采用 [MIT 协议](https://github.com/scagogogo/cpe/blob/main/LICENSE) 进行许可。

## 🤝 贡献指南

欢迎贡献代码、文档和反馈。请通过GitHub Issues和Pull Requests提交您的贡献。

## 📦 相关项目

- [scagogogo/cve](https://github.com/scagogogo/cve) - CVE处理工具库





