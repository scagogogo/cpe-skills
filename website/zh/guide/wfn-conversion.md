# WFN转换

本示例演示如何在CPE格式和Well-Formed Name (WFN)格式之间进行转换，以及如何使用WFN进行高效的处理和匹配操作。

## 概述

Well-Formed Name (WFN) 是CPE的内部标准表示格式，提供了一种规范化的方式来表示CPE组件，使匹配和比较操作更加高效和可靠。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== WFN转换示例 ===")
    
    // 示例1：CPE到WFN转换
    fmt.Println("\n1. CPE到WFN转换:")
    
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:/a:oracle:java:1.8.0_291",
        "cpe:2.3:o:linux:kernel:5.4.0:*:*:*:*:*:*:*",
    }
    
    for i, cpeStr := range cpeStrings {
        fmt.Printf("\n示例 %d: %s\n", i+1, cpeStr)
        
        // 解析CPE
        cpeObj, err := cpeskills.ParseCPE(cpeStr)
        if err != nil {
            log.Printf("解析CPE失败: %v", err)
            continue
        }
        
        // 转换为WFN
        wfn, err := cpeskills.CPEToWFN(cpeObj)
        if err != nil {
            log.Printf("转换为WFN失败: %v", err)
            continue
        }
        
        fmt.Printf("  原始CPE: %s\n", cpeStr)
        fmt.Printf("  WFN格式: %s\n", wfn.String())
        fmt.Printf("  部件:     %s\n", wfn.Part)
        fmt.Printf("  供应商:   %s\n", wfn.Vendor)
        fmt.Printf("  产品:     %s\n", wfn.Product)
        fmt.Printf("  版本:     %s\n", wfn.Version)
    }
    
    // 示例2：WFN到CPE转换
    fmt.Println("\n2. WFN到CPE转换:")
    
    // 手动创建WFN
    wfn := &cpeskills.WFN{
        Part:           "a",
        Vendor:         "adobe",
        Product:        "reader",
        Version:        "2021.001.20150",
        Update:         cpeskills.WFNAny,
        Edition:        cpeskills.WFNAny,
        Language:       cpeskills.WFNAny,
        SoftwareEdition: cpeskills.WFNAny,
        TargetSoftware: cpeskills.WFNAny,
        TargetHardware: cpeskills.WFNAny,
        Other:          cpeskills.WFNAny,
    }
    
    fmt.Printf("WFN: %s\n", wfn.String())
    
    // 转换为CPE 2.3
    cpe23, err := cpeskills.WFNToCPE23(wfn)
    if err != nil {
        log.Printf("转换为CPE 2.3失败: %v", err)
    } else {
        fmt.Printf("CPE 2.3: %s\n", cpe23)
    }
    
    // 转换为CPE 2.2
    cpe22, err := cpeskills.WFNToCPE22(wfn)
    if err != nil {
        log.Printf("转换为CPE 2.2失败: %v", err)
    } else {
        fmt.Printf("CPE 2.2: %s\n", cpe22)
    }
    
    // 示例3：WFN属性值
    fmt.Println("\n3. WFN属性值:")
    
    // 演示不同的WFN属性值
    examples := []struct {
        name  string
        value string
        desc  string
    }{
        {"ANY", cpeskills.WFNAny, "匹配任意值"},
        {"NA", cpeskills.WFNNotApplicable, "不适用"},
        {"字面值", "windows", "字面字符串值"},
        {"转义值", cpeskills.QuoteWFNValue("special~chars"), "转义特殊字符"},
    }
    
    for _, example := range examples {
        fmt.Printf("  %s: '%s' - %s\n", example.name, example.value, example.desc)
    }
    
    // 示例4：WFN匹配
    fmt.Println("\n4. WFN匹配:")
    
    // 创建源和目标WFN
    sourceWFN := &cpeskills.WFN{
        Part:    "a",
        Vendor:  "microsoft",
        Product: cpeskills.WFNAny, // 任意产品
        Version: cpeskills.WFNAny, // 任意版本
    }
    
    targetWFNs := []*cpeskills.WFN{
        {Part: "a", Vendor: "microsoft", Product: "windows", Version: "10"},
        {Part: "a", Vendor: "microsoft", Product: "office", Version: "2019"},
        {Part: "a", Vendor: "oracle", Product: "java", Version: "11"},
        {Part: "o", Vendor: "microsoft", Product: "windows", Version: "10"},
    }
    
    fmt.Printf("源WFN: %s\n", sourceWFN.String())
    fmt.Println("匹配目标:")
    
    for i, targetWFN := range targetWFNs {
        match := cpeskills.MatchWFN(sourceWFN, targetWFN)
        status := "❌"
        if match {
            status = "✅"
        }
        fmt.Printf("  %s 目标 %d: %s\n", status, i+1, targetWFN.String())
    }
    
    // 示例5：WFN验证
    fmt.Println("\n5. WFN验证:")
    
    validationTests := []struct {
        wfn   *cpeskills.WFN
        desc  string
        valid bool
    }{
        {
            &cpeskills.WFN{Part: "a", Vendor: "microsoft", Product: "windows"},
            "有效的应用程序WFN",
            true,
        },
        {
            &cpeskills.WFN{Part: "x", Vendor: "microsoft", Product: "windows"},
            "无效的部件值",
            false,
        },
        {
            &cpeskills.WFN{Part: "a", Vendor: "", Product: "windows"},
            "空供应商",
            false,
        },
        {
            &cpeskills.WFN{Part: "a", Vendor: "microsoft", Product: ""},
            "空产品",
            false,
        },
    }
    
    for i, test := range validationTests {
        err := cpeskills.ValidateWFN(test.wfn)
        isValid := err == nil
        
        status := "❌"
        if isValid == test.valid {
            status = "✅"
        }
        
        fmt.Printf("  %s 测试 %d: %s\n", status, i+1, test.desc)
        fmt.Printf("    WFN: %s\n", test.wfn.String())
        if err != nil {
            fmt.Printf("    错误: %v\n", err)
        }
    }
    
    // 示例6：WFN规范化
    fmt.Println("\n6. WFN规范化:")
    
    unnormalizedWFN := &cpeskills.WFN{
        Part:    "A", // 应该是小写
        Vendor:  "Microsoft", // 应该是小写
        Product: "Windows~10", // 特殊字符
        Version: "10.0.19041.1234",
    }
    
    fmt.Printf("规范化前: %s\n", unnormalizedWFN.String())
    
    normalizedWFN := cpeskills.NormalizeWFN(unnormalizedWFN)
    fmt.Printf("规范化后: %s\n", normalizedWFN.String())
    
    // 示例7：WFN比较
    fmt.Println("\n7. WFN比较:")
    
    wfn1 := &cpeskills.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.0",
    }
    wfn2 := &cpeskills.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.1",
    }
    wfn3 := &cpeskills.WFN{
        Part: "a", Vendor: "apache", Product: "tomcat", Version: "9.0.0",
    }
    
    fmt.Printf("WFN1: %s\n", wfn1.String())
    fmt.Printf("WFN2: %s\n", wfn2.String())
    fmt.Printf("WFN3: %s\n", wfn3.String())
    
    fmt.Printf("WFN1 == WFN2: %t\n", cpeskills.CompareWFN(wfn1, wfn2) == 0)
    fmt.Printf("WFN1 == WFN3: %t\n", cpeskills.CompareWFN(wfn1, wfn3) == 0)
    fmt.Printf("WFN1 < WFN2:  %t\n", cpeskills.CompareWFN(wfn1, wfn2) < 0)
    
    // 示例8：特殊字符处理
    fmt.Println("\n8. 特殊字符处理:")
    
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
    
    // 示例9：批量转换
    fmt.Println("\n9. 批量转换:")
    
    batchCPEs := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
    }
    
    fmt.Printf("批量转换 %d 个CPE:\n", len(batchCPEs))
    
    wfnList := []*cpeskills.WFN{}
    for i, cpeStr := range batchCPEs {
        cpeObj, err := cpeskills.ParseCpe23(cpeStr)
        if err != nil {
            fmt.Printf("  ❌ %d. 解析失败: %s\n", i+1, cpeStr)
            continue
        }
        
        wfn, err := cpeskills.CPEToWFN(cpeObj)
        if err != nil {
            fmt.Printf("  ❌ %d. 转换失败: %s\n", i+1, cpeStr)
            continue
        }
        
        wfnList = append(wfnList, wfn)
        fmt.Printf("  ✅ %d. %s %s %s\n", i+1, wfn.Vendor, wfn.Product, wfn.Version)
    }
    
    // 转换回CPE格式
    fmt.Printf("\n转换回CPE 2.3格式:\n")
    for i, wfn := range wfnList {
        cpe23, err := cpeskills.WFNToCPE23(wfn)
        if err != nil {
            fmt.Printf("  ❌ %d. 转换失败\n", i+1)
        } else {
            fmt.Printf("  ✅ %d. %s\n", i+1, cpe23)
        }
    }
}
```

## 预期输出

```
=== WFN转换示例 ===

1. CPE到WFN转换:

示例 1: cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*
  原始CPE: cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*
  WFN格式: wfn:[part="a",vendor="microsoft",product="windows",version="10",update=ANY,edition=ANY,language=ANY,sw_edition=ANY,target_sw=ANY,target_hw=ANY,other=ANY]
  部件:     a
  供应商:   microsoft
  产品:     windows
  版本:     10

2. WFN到CPE转换:
WFN: wfn:[part="a",vendor="adobe",product="reader",version="2021.001.20150",update=ANY,edition=ANY,language=ANY,sw_edition=ANY,target_sw=ANY,target_hw=ANY,other=ANY]
CPE 2.3: cpe:2.3:a:adobe:reader:2021.001.20150:*:*:*:*:*:*:*
CPE 2.2: cpe:/a:adobe:reader:2021.001.20150

3. WFN属性值:
  ANY: '*' - 匹配任意值
  NA: '-' - 不适用
  字面值: 'windows' - 字面字符串值
  转义值: 'special\~chars' - 转义特殊字符

4. WFN匹配:
源WFN: wfn:[part="a",vendor="microsoft",product=ANY,version=ANY]
匹配目标:
  ✅ 目标 1: wfn:[part="a",vendor="microsoft",product="windows",version="10"]
  ✅ 目标 2: wfn:[part="a",vendor="microsoft",product="office",version="2019"]
  ❌ 目标 3: wfn:[part="a",vendor="oracle",product="java",version="11"]
  ❌ 目标 4: wfn:[part="o",vendor="microsoft",product="windows",version="10"]

...
```

## 关键概念

### 1. WFN结构

WFN包含11个属性：
- **part**: 组件类型 (a, h, o)
- **vendor**: 供应商名称
- **product**: 产品名称
- **version**: 版本字符串
- **update**: 更新标识符
- **edition**: 版本信息
- **language**: 语言代码
- **sw_edition**: 软件版本
- **target_sw**: 目标软件
- **target_hw**: 目标硬件
- **other**: 其他信息

### 2. 特殊值

- **ANY (*)**: 匹配任意值
- **NA (-)**: 不适用/未定义
- **字面值**: 精确字符串匹配

### 3. WFN优势

- **规范形式**: 标准化表示
- **高效匹配**: 优化的比较操作
- **验证**: 内置验证规则
- **规范化**: 一致的格式化

## 最佳实践

1. **内部处理使用WFN**: 将CPE字符串转换为WFN进行操作
2. **验证WFN**: 使用前始终验证WFN对象
3. **规范化输入**: 规范化WFN以确保一致的比较
4. **处理特殊值**: 正确处理ANY和NA值
5. **转换回去**: 将WFN转换回CPE格式进行输出

## 下一步

- 学习[高级匹配](./advanced-matching.md)使用WFN
- 探索[CPE集合](./sets.md)进行批量WFN操作
- 查看[存储](./storage.md)来持久化WFN数据
