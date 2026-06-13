# 使用示例

本节提供了实用的示例，演示如何在实际场景中使用CPE库。每个示例都包含完整的、可运行的代码和解释。

## 可用示例

### 基础用法
- **[基础解析](./basic-parsing.md)** - 解析CPE字符串并访问组件
- **[CPE匹配](./matching.md)** - 比较和匹配CPE对象
- **[WFN转换](./wfn-conversion.md)** - CPE和WFN格式之间的转换

### 高级功能
- **[版本比较](./version-comparison.md)** - 比较版本字符串和范围
- **[适用性语言](./applicability.md)** - 使用CPE适用性表达式
- **[CPE集合](./sets.md)** - 处理CPE集合
- **[高级匹配](./advanced-matching.md)** - 使用复杂的匹配算法

### 集成
- **[存储操作](./storage.md)** - 使用不同后端持久化CPE数据
- **[NVD集成](./nvd-integration.md)** - 下载和使用NVD数据
- **[CVE映射](./cve-mapping.md)** - 将CPE映射到漏洞

## 快速开始示例

这是一个简单的示例来帮助你开始：

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    // 解析CPE字符串
    cpeObj, err := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    // 访问组件
    fmt.Printf("供应商: %s\n", cpeObj.Vendor)
    fmt.Printf("产品: %s\n", cpeObj.ProductName)
    fmt.Printf("版本: %s\n", cpeObj.Version)
    
    // 创建匹配模式
    pattern, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:*:*:*:*:*:*:*:*:*")
    
    // 测试匹配
    if pattern.Match(cpeObj) {
        fmt.Println("CPE匹配Microsoft模式！")
    }
}
```

## 示例分类

### 🔍 解析和格式化
学习如何解析CPE字符串、处理不同格式，以及在CPE 2.2和2.3之间转换。

### 🎯 匹配和比较
发现各种匹配技术，从基本的通配符匹配到带有评分的高级模糊匹配。

### 📊 数据管理
探索如何高效地存储、检索和管理大量CPE数据集合。

### 🔗 外部集成
了解如何与国家漏洞数据库等外部数据源集成。

### 🛡️ 安全应用
学习如何使用CPE进行漏洞管理、资产清单和安全扫描。

## 运行示例

所有示例都是完整的、独立的程序。要运行它们：

1. **安装库：**
   ```bash
   go get github.com/scagogogo/cpe
   ```

2. **创建新的Go文件** 包含示例代码

3. **运行示例：**
   ```bash
   go run example.go
   ```

## 示例结构

每个示例都遵循一致的结构：

- **概览** - 示例演示的内容
- **完整代码** - 完整的、可运行的程序
- **解释** - 逐步分解
- **输出** - 预期结果
- **变化** - 替代方法或扩展

## 常见模式

### 错误处理
```go
cpeObj, err := cpe.ParseCpe23(cpeString)
if err != nil {
    if cpe.IsInvalidFormatError(err) {
        fmt.Printf("无效格式: %s\n", cpeString)
        return
    }
    log.Fatal(err)
}
```

### 资源清理
```go
storage, err := cpe.NewFileStorage("./data", true)
if err != nil {
    log.Fatal(err)
}
defer storage.Close()

err = storage.Initialize()
if err != nil {
    log.Fatal(err)
}
```

### 批处理
```go
cpeStrings := []string{
    "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
    "cpe:2.3:a:apache:tomcat:9.0:*:*:*:*:*:*:*",
}

for _, cpeStr := range cpeStrings {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
    if err != nil {
        log.Printf("解析失败 %s: %v", cpeStr, err)
        continue
    }
    
    // 处理CPE
    fmt.Printf("已处理: %s\n", cpeObj.GetURI())
}
```

## 最佳实践

### 1. 始终处理错误
```go
// 好的做法
cpeObj, err := cpe.ParseCpe23(input)
if err != nil {
    return fmt.Errorf("解析CPE失败: %w", err)
}

// 不好的做法
cpeObj, _ := cpe.ParseCpe23(input) // 忽略错误
```

### 2. 使用适当的存储
```go
// 用于测试
storage := cpe.NewMemoryStorage()

// 用于生产
storage, err := cpe.NewFileStorage("./cpe-data", true)
```

### 3. 验证输入
```go
err := cpe.ValidateCPEString(userInput)
if err != nil {
    return fmt.Errorf("无效的CPE格式: %w", err)
}
```

### 4. 对集合使用Sets
```go
// 对大集合高效
cpeSet := cpe.NewCPESet()
cpeSet.Add(cpe1, cpe2, cpe3)

// 高效过滤
microsoftCPEs := cpeSet.FilterByVendor("microsoft")
```

## 性能提示

### 1. 启用缓存
```go
storage, _ := cpe.NewFileStorage("./data", true) // 启用缓存
```

### 2. 使用批操作
```go
// 比单个操作更好
cpeSet := cpe.FromArray(cpeArray)
results := cpeSet.FilterByVendor("microsoft")
```

### 3. 重用匹配选项
```go
options := cpe.DefaultMatchOptions()
// 为多个匹配重用选项
for _, cpe := range cpes {
    if cpe.MatchCPE(pattern, cpe, options) {
        // 处理匹配
    }
}
```

## 获取帮助

如果你需要示例帮助：

1. 查看[API参考](/zh/api/)获取详细的函数文档
2. 查看完整的示例代码以获取上下文
3. 查看错误处理模式
4. 查看[GitHub仓库](https://github.com/scagogogo/cpe)的问题和讨论

## 贡献示例

我们欢迎新示例的贡献！如果你有一个有用的示例来演示特定用例，请考虑为项目贡献。

## 下一步

从[基础解析](./basic-parsing.md)示例开始学习基础知识，然后根据你的具体需求探索更高级的示例。
