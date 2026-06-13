# 适用性语言

本示例演示如何使用CPE适用性语言来表达复杂的匹配条件和CPE名称之间的逻辑关系。

## 概述

CPE适用性语言允许您创建复杂的表达式，定义特定信息（如漏洞）何时适用于系统。它支持逻辑运算符（AND、OR、NOT）和复杂的嵌套条件。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE适用性语言示例 ===")
    
    // 示例1：基本适用性表达式
    fmt.Println("\n1. 基本适用性表达式:")
    
    // 简单表达式：适用于Windows 10
    expr1 := "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*"
    
    // OR表达式：适用于Windows 10或Windows 11
    expr2 := `(cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR 
               cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*)`
    
    // AND表达式：适用于Windows 10和特定更新
    expr3 := `(cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* AND 
               cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*)`
    
    expressions := []struct {
        name string
        expr string
        desc string
    }{
        {"简单", expr1, "单个CPE匹配"},
        {"OR逻辑", expr2, "多个备选项"},
        {"AND逻辑", expr3, "多个要求"},
    }
    
    for _, e := range expressions {
        fmt.Printf("\n%s表达式:\n", e.name)
        fmt.Printf("  描述: %s\n", e.desc)
        fmt.Printf("  表达式: %s\n", e.expr)
        
        // 解析表达式
        parsedExpr, err := cpe.ParseApplicabilityExpression(e.expr)
        if err != nil {
            log.Printf("解析表达式失败: %v", err)
            continue
        }
        
        fmt.Printf("  解析成功: %t\n", parsedExpr != nil)
    }
    
    // 示例2：复杂嵌套表达式
    fmt.Println("\n2. 复杂嵌套表达式:")
    
    // 复杂漏洞适用性
    complexExpr := `
    (
        (cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR 
         cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*) 
        AND 
        (cpe:2.3:a:microsoft:internet_explorer:*:*:*:*:*:*:*:* OR 
         cpe:2.3:a:microsoft:edge:*:*:*:*:*:*:*:*)
        AND NOT 
        cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*
    )`
    
    fmt.Printf("复杂表达式:\n%s\n", complexExpr)
    
    parsedComplex, err := cpe.ParseApplicabilityExpression(complexExpr)
    if err != nil {
        log.Printf("解析复杂表达式失败: %v", err)
    } else {
        fmt.Printf("成功解析复杂表达式\n")
        fmt.Printf("表达式类型: %s\n", parsedComplex.Type())
        fmt.Printf("操作数数量: %d\n", len(parsedComplex.Operands()))
    }
    
    // 示例3：测试适用性
    fmt.Println("\n3. 测试适用性:")
    
    // 定义测试系统
    testSystems := []struct {
        name string
        cpes []string
    }{
        {
            "Windows 10 with IE",
            []string{
                "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:internet_explorer:11:*:*:*:*:*:*:*",
            },
        },
        {
            "Windows 11 with Edge",
            []string{
                "cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:edge:95.0.1020.44:*:*:*:*:*:*:*",
            },
        },
        {
            "Windows 10 with patch",
            []string{
                "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:internet_explorer:11:*:*:*:*:*:*:*",
                "cpe:2.3:a:microsoft:windows_update:kb5005565:*:*:*:*:*:*:*",
            },
        },
        {
            "Linux系统",
            []string{
                "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
                "cpe:2.3:a:mozilla:firefox:95.0:*:*:*:*:*:*:*",
            },
        },
    }
    
    // 针对复杂表达式测试每个系统
    for _, system := range testSystems {
        fmt.Printf("\n测试系统: %s\n", system.name)
        
        // 将CPE字符串转换为对象
        systemCPEs := make([]*cpe.CPE, 0, len(system.cpes))
        for _, cpeStr := range system.cpes {
            cpeObj, err := cpe.ParseCpe23(cpeStr)
            if err != nil {
                log.Printf("解析CPE %s失败: %v", cpeStr, err)
                continue
            }
            systemCPEs = append(systemCPEs, cpeObj)
        }
        
        // 测试适用性
        applies := cpe.EvaluateApplicability(parsedComplex, systemCPEs)
        
        status := "❌ 不适用"
        if applies {
            status = "✅ 适用"
        }
        
        fmt.Printf("  结果: %s\n", status)
        fmt.Printf("  系统CPE:\n")
        for _, cpeStr := range system.cpes {
            fmt.Printf("    - %s\n", cpeStr)
        }
    }
    
    // 示例4：版本范围适用性
    fmt.Println("\n4. 版本范围适用性:")
    
    // Java版本8.0到11.0（不包括）的表达式
    javaRangeExpr := `
    (cpe:2.3:a:oracle:java:8.*:*:*:*:*:*:*:* OR
     cpe:2.3:a:oracle:java:9.*:*:*:*:*:*:*:* OR
     cpe:2.3:a:oracle:java:10.*:*:*:*:*:*:*:*)
    `
    
    fmt.Printf("Java版本范围表达式:\n%s\n", javaRangeExpr)
    
    javaExpr, err := cpe.ParseApplicabilityExpression(javaRangeExpr)
    if err != nil {
        log.Printf("解析Java表达式失败: %v", err)
    } else {
        // 测试不同的Java版本
        javaVersions := []string{
            "cpe:2.3:a:oracle:java:7.0.80:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:8.0.291:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:9.0.4:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
            "cpe:2.3:a:oracle:java:17.0.1:*:*:*:*:*:*:*",
        }
        
        fmt.Println("\n测试Java版本:")
        for _, javaVer := range javaVersions {
            javaCPE, _ := cpe.ParseCpe23(javaVer)
            applies := cpe.EvaluateApplicability(javaExpr, []*cpe.CPE{javaCPE})
            
            status := "❌"
            if applies {
                status = "✅"
            }
            
            fmt.Printf("  %s %s\n", status, javaVer)
        }
    }
    
    // 示例5：平台特定适用性
    fmt.Println("\n5. 平台特定适用性:")
    
    // Linux上Web服务器的表达式
    webServerLinuxExpr := `
    (cpe:2.3:o:*:linux:*:*:*:*:*:*:*:* OR
     cpe:2.3:o:canonical:ubuntu:*:*:*:*:*:*:*:* OR
     cpe:2.3:o:redhat:enterprise_linux:*:*:*:*:*:*:*:*)
    AND
    (cpe:2.3:a:apache:http_server:*:*:*:*:*:*:*:* OR
     cpe:2.3:a:nginx:nginx:*:*:*:*:*:*:*:*)
    `
    
    fmt.Printf("Linux上Web服务器表达式:\n%s\n", webServerLinuxExpr)
    
    webServerExpr, err := cpe.ParseApplicabilityExpression(webServerLinuxExpr)
    if err != nil {
        log.Printf("解析Web服务器表达式失败: %v", err)
    } else {
        // 测试不同的服务器配置
        serverConfigs := []struct {
            name string
            cpes []string
        }{
            {
                "Ubuntu上的Apache",
                []string{
                    "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
                    "cpe:2.3:a:apache:http_server:2.4.41:*:*:*:*:*:*:*",
                },
            },
            {
                "RHEL上的Nginx",
                []string{
                    "cpe:2.3:o:redhat:enterprise_linux:8:*:*:*:*:*:*:*",
                    "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",
                },
            },
            {
                "Windows上的IIS",
                []string{
                    "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
                    "cpe:2.3:a:microsoft:internet_information_services:10.0:*:*:*:*:*:*:*",
                },
            },
        }
        
        fmt.Println("\n测试服务器配置:")
        for _, config := range serverConfigs {
            configCPEs := make([]*cpe.CPE, 0, len(config.cpes))
            for _, cpeStr := range config.cpes {
                cpeObj, _ := cpe.ParseCpe23(cpeStr)
                configCPEs = append(configCPEs, cpeObj)
            }
            
            applies := cpe.EvaluateApplicability(webServerExpr, configCPEs)
            
            status := "❌"
            if applies {
                status = "✅"
            }
            
            fmt.Printf("  %s %s\n", status, config.name)
        }
    }
    
    // 示例6：表达式优化
    fmt.Println("\n6. 表达式优化:")
    
    // 原始冗长表达式
    verboseExpr := `
    (cpe:2.3:a:microsoft:office:2016:*:*:*:*:*:*:* OR
     cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:* OR
     cpe:2.3:a:microsoft:office:365:*:*:*:*:*:*:*)
    AND
    (cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:* OR
     cpe:2.3:o:microsoft:windows:11:*:*:*:*:*:*:*)
    `
    
    // 使用通配符的优化表达式
    optimizedExpr := `
    cpe:2.3:a:microsoft:office:*:*:*:*:*:*:*:*
    AND
    cpe:2.3:o:microsoft:windows:*:*:*:*:*:*:*:*
    `
    
    fmt.Printf("冗长表达式:\n%s\n", verboseExpr)
    fmt.Printf("优化表达式:\n%s\n", optimizedExpr)
    
    verboseParsed, _ := cpe.ParseApplicabilityExpression(verboseExpr)
    optimizedParsed, _ := cpe.ParseApplicabilityExpression(optimizedExpr)
    
    // 测试两个表达式
    testCPEs := []*cpe.CPE{
        mustParse("cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*"),
        mustParse("cpe:2.3:a:microsoft:office:2019:*:*:*:*:*:*:*"),
    }
    
    verboseResult := cpe.EvaluateApplicability(verboseParsed, testCPEs)
    optimizedResult := cpe.EvaluateApplicability(optimizedParsed, testCPEs)
    
    fmt.Printf("冗长结果: %t\n", verboseResult)
    fmt.Printf("优化结果: %t\n", optimizedResult)
    fmt.Printf("结果匹配: %t\n", verboseResult == optimizedResult)
}

func mustParse(cpeStr string) *cpe.CPE {
    cpeObj, err := cpe.ParseCpe23(cpeStr)
    if err != nil {
        panic(err)
    }
    return cpeObj
}
```

## 关键概念

### 1. 逻辑运算符

- **AND**: 所有条件必须为真
- **OR**: 至少一个条件必须为真
- **NOT**: 条件必须为假

### 2. 表达式结构

- **简单**: 单个CPE匹配
- **复合**: 带运算符的多个CPE
- **嵌套**: 复杂的层次条件

### 3. 用例

- **漏洞适用性**: 定义受影响的系统
- **策略合规**: 指定所需配置
- **资产分类**: 分组相似系统
- **补丁管理**: 识别更新目标

## 最佳实践

1. **使用通配符**: 适当时用通配符简化表达式
2. **逻辑分组**: 将相关条件分组在一起
3. **彻底测试**: 针对已知系统验证表达式
4. **记录意图**: 为复杂表达式添加注释
5. **优化性能**: 可能时优先选择简单表达式

## 下一步

- 学习[高级匹配](./advanced-matching.md)处理复杂场景
- 探索[CPE集合](./sets.md)进行批量操作
- 查看[NVD集成](./nvd-integration.md)了解实际适用性
