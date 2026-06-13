# 核心类型

本页面描述了CPE库中的核心数据类型和结构体，这些是使用库时需要了解的基本构建块。

## CPE 结构体

### CPE

主要的CPE结构体，表示一个完整的通用平台枚举对象。

```go
type CPE struct {
    Cpe23           string    // CPE 2.3格式字符串
    Part            Part      // 组件类型（应用程序、硬件、操作系统）
    Vendor          Vendor    // 供应商信息
    ProductName     Product   // 产品名称
    Version         Version   // 版本信息
    Update          Update    // 更新信息
    Edition         Edition   // 版本信息
    Language        Language  // 语言代码
    SoftwareEdition string    // 软件版本
    TargetSoftware  string    // 目标软件
    TargetHardware  string    // 目标硬件
    Other           string    // 其他信息
    Cve             string    // 关联的CVE
    Url             string    // 参考URL
}
```

#### 方法

```go
// 获取CPE的URI表示
func (c *CPE) GetURI() string

// 获取CPE 2.3格式字符串
func (c *CPE) GetCPE23() string

// 获取CPE 2.2格式字符串
func (c *CPE) GetCPE22() string

// 检查CPE是否匹配另一个CPE
func (c *CPE) Match(other *CPE) bool

// 验证CPE的有效性
func (c *CPE) Validate() error

// 克隆CPE对象
func (c *CPE) Clone() *CPE
```

## 组件类型

### Part

表示CPE的组件类型。

```go
type Part struct {
    ShortName   string // 短名称（a, h, o）
    LongName    string // 长名称（Application, Hardware, Operating System）
    Description string // 描述
}
```

#### 预定义常量

```go
var (
    PartApplication     = Part{ShortName: "a", LongName: "Application", Description: "应用程序"}
    PartHardware        = Part{ShortName: "h", LongName: "Hardware", Description: "硬件设备"}
    PartOperatingSystem = Part{ShortName: "o", LongName: "Operating System", Description: "操作系统"}
)
```

### Vendor

供应商信息类型。

```go
type Vendor string

// 常用供应商常量
const (
    VendorMicrosoft = Vendor("microsoft")
    VendorApache    = Vendor("apache")
    VendorOracle    = Vendor("oracle")
    VendorGoogle    = Vendor("google")
    VendorCisco     = Vendor("cisco")
)
```

### Product

产品名称类型。

```go
type Product string

// 常用产品常量
const (
    ProductWindows     = Product("windows")
    ProductTomcat      = Product("tomcat")
    ProductJava        = Product("java")
    ProductChrome      = Product("chrome")
    ProductFirefox     = Product("firefox")
)
```

### Version

版本信息类型。

```go
type Version string

// 特殊版本值
const (
    VersionAny = Version("*")  // 任意版本
    VersionNA  = Version("-")  // 不适用
)
```

## WFN 类型

### WFN

Well-Formed Name结构体，CPE的内部表示格式。

```go
type WFN struct {
    Part            string // 组件类型
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

#### WFN 特殊值

```go
const (
    WFNAny           = "*"  // 任意值
    WFNNotApplicable = "-"  // 不适用
)
```

#### 方法

```go
// 获取WFN的字符串表示
func (w *WFN) String() string

// 验证WFN的有效性
func (w *WFN) Validate() error

// 比较两个WFN
func (w *WFN) Compare(other *WFN) int

// 检查WFN是否匹配另一个WFN
func (w *WFN) Match(other *WFN) bool
```

## 匹配类型

### MatchOptions

匹配选项配置。

```go
type MatchOptions struct {
    ExactMatch      bool    // 是否精确匹配
    IgnoreCase      bool    // 是否忽略大小写
    AllowWildcards  bool    // 是否允许通配符
    FuzzyThreshold  float64 // 模糊匹配阈值
}
```

### MatchResult

匹配结果。

```go
type MatchResult struct {
    Match      bool    // 是否匹配
    Score      float64 // 匹配分数（0.0-1.0）
    Confidence float64 // 置信度
    Details    string  // 详细信息
}
```

### MatchWeights

匹配权重配置。

```go
type MatchWeights struct {
    Part    float64 // 组件类型权重
    Vendor  float64 // 供应商权重
    Product float64 // 产品权重
    Version float64 // 版本权重
    Update  float64 // 更新权重
    Edition float64 // 版本权重
}
```

## 存储类型

### Storage

存储接口定义。

```go
type Storage interface {
    // 初始化存储
    Initialize() error
    
    // 存储CPE
    Store(cpe *CPE) error
    
    // 检索CPE
    Retrieve(id string) (*CPE, error)
    
    // 删除CPE
    Delete(id string) error
    
    // 列出所有CPE
    List() ([]*CPE, error)
    
    // 搜索CPE
    Search(query string) ([]*CPE, error)
    
    // 关闭存储
    Close() error
}
```

### FileStorage

基于文件的存储实现。

```go
type FileStorage struct {
    BaseDir     string // 基础目录
    EnableCache bool   // 是否启用缓存
    CacheSize   int    // 缓存大小
}
```

### MemoryStorage

基于内存的存储实现。

```go
type MemoryStorage struct {
    data map[string]*CPE // 内存数据
    mu   sync.RWMutex    // 读写锁
}
```

## 集合类型

### CPESet

CPE集合类型。

```go
type CPESet struct {
    items map[string]*CPE // CPE项目
    mu    sync.RWMutex    // 读写锁
}
```

#### 方法

```go
// 添加CPE到集合
func (s *CPESet) Add(cpe *CPE) bool

// 从集合中移除CPE
func (s *CPESet) Remove(cpe *CPE) bool

// 检查CPE是否在集合中
func (s *CPESet) Contains(cpe *CPE) bool

// 获取集合大小
func (s *CPESet) Size() int

// 清空集合
func (s *CPESet) Clear()

// 转换为切片
func (s *CPESet) ToSlice() []*CPE

// 并集操作
func (s *CPESet) Union(other *CPESet) *CPESet

// 交集操作
func (s *CPESet) Intersection(other *CPESet) *CPESet

// 差集操作
func (s *CPESet) Difference(other *CPESet) *CPESet
```

## 错误类型

### CPEError

CPE相关错误的基础类型。

```go
type CPEError struct {
    Type    ErrorType // 错误类型
    Message string    // 错误消息
    Field   string    // 相关字段
    Value   string    // 错误值
}
```

### ErrorType

错误类型枚举。

```go
type ErrorType int

const (
    ErrorTypeInvalidFormat ErrorType = iota // 无效格式
    ErrorTypeInvalidPart                    // 无效组件类型
    ErrorTypeInvalidVendor                  // 无效供应商
    ErrorTypeInvalidProduct                 // 无效产品
    ErrorTypeInvalidVersion                 // 无效版本
    ErrorTypeParsingError                   // 解析错误
    ErrorTypeValidationError                // 验证错误
    ErrorTypeStorageError                   // 存储错误
)
```

#### 方法

```go
// 实现error接口
func (e *CPEError) Error() string

// 获取错误类型
func (e *CPEError) GetType() ErrorType

// 检查是否为特定类型的错误
func (e *CPEError) Is(errorType ErrorType) bool
```

## 字典类型

### CPEDictionary

CPE字典结构。

```go
type CPEDictionary struct {
    Entries      []*CPEDictionaryEntry // 字典条目
    LastModified time.Time             // 最后修改时间
    Version      string                // 字典版本
}
```

### CPEDictionaryEntry

字典条目。

```go
type CPEDictionaryEntry struct {
    CPE23        string    // CPE 2.3格式
    Title        string    // 标题
    References   []string  // 参考链接
    LastModified time.Time // 最后修改时间
}
```

## NVD 类型

### NVDClient

NVD客户端。

```go
type NVDClient struct {
    APIKey      string        // API密钥
    BaseURL     string        // 基础URL
    Timeout     time.Duration // 超时时间
    RateLimit   int           // 速率限制
}
```

### CVEEntry

CVE条目。

```go
type CVEEntry struct {
    ID            string    // CVE ID
    Description   string    // 描述
    BaseScore     float64   // CVSS基础分数
    PublishedDate time.Time // 发布日期
    LastModified  time.Time // 最后修改时间
    AffectedCPEs  []string  // 受影响的CPE
}
```

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/scagogogo/cpe"
)

func main() {
    // 创建CPE对象
    cpeObj := &cpe.CPE{
        Part:        cpe.PartApplication,
        Vendor:      cpe.VendorMicrosoft,
        ProductName: cpe.ProductWindows,
        Version:     cpe.Version("10"),
    }
    
    // 验证CPE
    if err := cpeObj.Validate(); err != nil {
        fmt.Printf("CPE验证失败: %v\n", err)
        return
    }
    
    // 获取CPE字符串
    fmt.Printf("CPE 2.3: %s\n", cpeObj.GetCPE23())
    fmt.Printf("CPE 2.2: %s\n", cpeObj.GetCPE22())
    
    // 创建CPE集合
    cpeSet := cpe.NewCPESet()
    cpeSet.Add(cpeObj)
    
    fmt.Printf("集合大小: %d\n", cpeSet.Size())
}
```

## 下一步

- 了解[解析功能](./parsing.md)来处理CPE字符串
- 学习[匹配算法](./matching.md)来比较CPE对象
- 探索[存储接口](./storage.md)来持久化CPE数据
