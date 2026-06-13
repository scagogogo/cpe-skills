# NVD集成

本示例演示如何与美国国家漏洞数据库(NVD)集成，下载CPE字典和漏洞数据，并执行安全分析。

## 概述

NVD集成功能允许您从官方来源获取最新的CPE字典和CVE数据，进行漏洞评估和安全分析。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/scagogogo/cpe"
)

func main() {
    fmt.Println("=== NVD集成示例 ===")
    
    // 示例1：设置NVD客户端
    fmt.Println("\n1. 设置NVD客户端:")
    
    // 创建NVD客户端配置
    config := &cpe.NVDConfig{
        APIKey:         "", // 可选：NVD API密钥用于提高速率限制
        CacheDir:       "./nvd_cache",
        UpdateInterval: 24 * time.Hour,
        EnableCache:    true,
        Timeout:        30 * time.Second,
    }
    
    client := cpe.NewNVDClient(config)
    fmt.Println("✅ NVD客户端创建成功")
    
    // 示例2：下载CPE字典
    fmt.Println("\n2. 下载CPE字典:")
    
    fmt.Println("正在下载官方CPE字典...")
    dictionary, err := client.DownloadCPEDictionary()
    if err != nil {
        log.Printf("下载CPE字典失败: %v", err)
        // 使用示例字典继续演示
        dictionary = createSampleDictionary()
        fmt.Println("使用示例字典继续演示")
    } else {
        fmt.Printf("✅ 下载完成，包含 %d 个CPE条目\n", len(dictionary.Entries))
    }
    
    // 字典统计
    stats := dictionary.GetStatistics()
    fmt.Printf("字典统计:\n")
    fmt.Printf("  总条目数: %d\n", stats.TotalEntries)
    fmt.Printf("  应用程序: %d\n", stats.ApplicationCount)
    fmt.Printf("  操作系统: %d\n", stats.OperatingSystemCount)
    fmt.Printf("  硬件设备: %d\n", stats.HardwareCount)
    fmt.Printf("  供应商数量: %d\n", stats.VendorCount)
    
    // 示例3：搜索CPE字典
    fmt.Println("\n3. 搜索CPE字典:")
    
    searchTerms := []string{"apache", "microsoft", "oracle", "cisco"}
    
    for _, term := range searchTerms {
        results := dictionary.Search(term, 5) // 限制5个结果
        fmt.Printf("\n搜索 '%s' 的结果 (%d 个):\n", term, len(results))
        
        for i, entry := range results {
            fmt.Printf("  %d. %s\n", i+1, entry.Title)
            fmt.Printf("     %s\n", entry.CPE23)
        }
    }
    
    // 示例4：查询漏洞信息
    fmt.Println("\n4. 查询漏洞信息:")
    
    // 查询Apache Tomcat的漏洞
    fmt.Println("查询Apache Tomcat的高危漏洞:")
    
    query := cpe.NVDQuery{
        CPEVendor:    "apache",
        CPEProduct:   "tomcat",
        CVSSScoreMin: 7.0, // 只查询高危漏洞
        Limit:        10,
    }
    
    vulnerabilities, err := client.QueryVulnerabilities(query)
    if err != nil {
        log.Printf("查询漏洞失败: %v", err)
        // 使用示例漏洞数据
        vulnerabilities = createSampleVulnerabilities()
        fmt.Println("使用示例漏洞数据继续演示")
    }
    
    fmt.Printf("找到 %d 个Tomcat高危漏洞:\n", len(vulnerabilities))
    for i, vuln := range vulnerabilities {
        fmt.Printf("  %d. %s (CVSS: %.1f)\n", i+1, vuln.ID, vuln.BaseScore)
        fmt.Printf("     %s\n", truncateString(vuln.Description, 80))
        fmt.Printf("     发布日期: %s\n", vuln.PublishedDate.Format("2006-01-02"))
    }
    
    // 示例5：获取特定CVE详情
    fmt.Println("\n5. 获取特定CVE详情:")
    
    // 查询著名的Log4Shell漏洞
    cveID := "CVE-2021-44228"
    fmt.Printf("获取 %s 的详细信息:\n", cveID)
    
    cveDetails, err := client.GetCVEDetails(cveID)
    if err != nil {
        log.Printf("获取CVE详情失败: %v", err)
        // 使用示例CVE数据
        cveDetails = createSampleCVE(cveID)
        fmt.Println("使用示例CVE数据继续演示")
    }
    
    fmt.Printf("CVE详情:\n")
    fmt.Printf("  ID: %s\n", cveDetails.ID)
    fmt.Printf("  CVSS分数: %.1f\n", cveDetails.BaseScore)
    fmt.Printf("  严重程度: %s\n", cveDetails.Severity)
    fmt.Printf("  发布日期: %s\n", cveDetails.PublishedDate.Format("2006-01-02"))
    fmt.Printf("  描述: %s\n", truncateString(cveDetails.Description, 120))
    
    if len(cveDetails.AffectedCPEs) > 0 {
        fmt.Printf("  受影响的CPE数量: %d\n", len(cveDetails.AffectedCPEs))
        fmt.Printf("  示例受影响的CPE:\n")
        for i, affectedCPE := range cveDetails.AffectedCPEs[:min(3, len(cveDetails.AffectedCPEs))] {
            fmt.Printf("    %d. %s\n", i+1, affectedCPE)
        }
    }
    
    // 示例6：系统漏洞评估
    fmt.Println("\n6. 系统漏洞评估:")
    
    // 定义一个示例系统清单
    systemInventory := []string{
        "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.45:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:1.8.0_291:*:*:*:*:*:*:*",
        "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:nginx:nginx:1.18.0:*:*:*:*:*:*:*",
    }
    
    fmt.Printf("评估系统清单 (%d 个组件):\n", len(systemInventory))
    
    totalVulnerabilities := 0
    highSeverityCount := 0
    
    for i, componentStr := range systemInventory {
        component, _ := cpe.ParseCpe23(componentStr)
        fmt.Printf("\n  %d. %s %s %s\n", i+1, component.Vendor, component.ProductName, component.Version)
        
        // 查询该组件的漏洞
        componentQuery := cpe.NVDQuery{
            CPEVendor:  component.Vendor,
            CPEProduct: component.ProductName,
            Limit:      5,
        }
        
        componentVulns, err := client.QueryVulnerabilities(componentQuery)
        if err != nil {
            fmt.Printf("     ❌ 查询漏洞失败: %v\n", err)
            continue
        }
        
        if len(componentVulns) == 0 {
            fmt.Printf("     ✅ 未发现已知漏洞\n")
        } else {
            fmt.Printf("     ⚠️  发现 %d 个漏洞:\n", len(componentVulns))
            
            for j, vuln := range componentVulns {
                severity := "中等"
                if vuln.BaseScore >= 7.0 {
                    severity = "高危"
                    highSeverityCount++
                }
                
                fmt.Printf("       %d. %s (CVSS: %.1f, %s)\n", 
                    j+1, vuln.ID, vuln.BaseScore, severity)
                totalVulnerabilities++
            }
        }
    }
    
    fmt.Printf("\n漏洞评估摘要:\n")
    fmt.Printf("  总漏洞数: %d\n", totalVulnerabilities)
    fmt.Printf("  高危漏洞: %d\n", highSeverityCount)
    fmt.Printf("  风险等级: %s\n", getRiskLevel(highSeverityCount, totalVulnerabilities))
    
    // 示例7：自动更新
    fmt.Println("\n7. 自动更新:")
    
    // 检查更新
    fmt.Println("检查NVD数据更新:")
    updateInfo, err := client.CheckForUpdates()
    if err != nil {
        log.Printf("检查更新失败: %v", err)
    } else {
        if updateInfo.HasUpdates {
            fmt.Printf("发现更新:\n")
            fmt.Printf("  新条目: %d\n", updateInfo.NewEntriesCount)
            fmt.Printf("  更新条目: %d\n", updateInfo.UpdatedEntriesCount)
            fmt.Printf("  数据大小: %.1f MB\n", float64(updateInfo.TotalSize)/1024/1024)
            
            // 在实际应用中，您可能会选择下载更新
            // err = client.DownloadUpdates()
            fmt.Println("  (在生产环境中会自动下载更新)")
        } else {
            fmt.Println("数据已是最新版本")
        }
    }
    
    // 示例8：缓存管理
    fmt.Println("\n8. 缓存管理:")
    
    // 获取缓存统计
    cacheStats := client.GetCacheStats()
    fmt.Printf("缓存统计:\n")
    fmt.Printf("  缓存文件数: %d\n", cacheStats.FileCount)
    fmt.Printf("  缓存大小: %.1f MB\n", float64(cacheStats.TotalSize)/1024/1024)
    fmt.Printf("  缓存命中率: %.2f%%\n", cacheStats.HitRate*100)
    
    // 示例9：批量CPE验证
    fmt.Println("\n9. 批量CPE验证:")
    
    // 验证系统清单中的CPE是否在官方字典中
    fmt.Println("验证系统清单中的CPE:")
    
    validCount := 0
    for i, cpeStr := range systemInventory {
        isValid := dictionary.ValidateCPE(cpeStr)
        
        status := "❌ 不在官方字典"
        if isValid {
            status = "✅ 官方认证"
            validCount++
        }
        
        component, _ := cpe.ParseCpe23(cpeStr)
        fmt.Printf("  %d. %s %s - %s\n", 
            i+1, component.Vendor, component.ProductName, status)
    }
    
    fmt.Printf("\n验证摘要: %d/%d CPE在官方字典中\n", validCount, len(systemInventory))
    
    // 示例10：导出报告
    fmt.Println("\n10. 导出报告:")
    
    // 创建漏洞报告
    report := &VulnerabilityReport{
        GeneratedAt:        time.Now(),
        SystemComponents:   len(systemInventory),
        TotalVulnerabilities: totalVulnerabilities,
        HighSeverityVulns:  highSeverityCount,
        ValidCPEs:          validCount,
        RiskLevel:          getRiskLevel(highSeverityCount, totalVulnerabilities),
    }
    
    fmt.Printf("生成漏洞评估报告:\n")
    fmt.Printf("  生成时间: %s\n", report.GeneratedAt.Format("2006-01-02 15:04:05"))
    fmt.Printf("  系统组件: %d\n", report.SystemComponents)
    fmt.Printf("  总漏洞数: %d\n", report.TotalVulnerabilities)
    fmt.Printf("  高危漏洞: %d\n", report.HighSeverityVulns)
    fmt.Printf("  有效CPE: %d\n", report.ValidCPEs)
    fmt.Printf("  风险等级: %s\n", report.RiskLevel)
    
    // 保存报告到文件
    reportFile := "vulnerability_report.json"
    err = saveReportToFile(report, reportFile)
    if err != nil {
        log.Printf("保存报告失败: %v", err)
    } else {
        fmt.Printf("  ✅ 报告已保存到: %s\n", reportFile)
    }
    
    fmt.Println("\n✅ NVD集成示例完成")
}

// 辅助结构体和函数
type VulnerabilityReport struct {
    GeneratedAt          time.Time `json:"generated_at"`
    SystemComponents     int       `json:"system_components"`
    TotalVulnerabilities int       `json:"total_vulnerabilities"`
    HighSeverityVulns    int       `json:"high_severity_vulns"`
    ValidCPEs            int       `json:"valid_cpes"`
    RiskLevel            string    `json:"risk_level"`
}

func createSampleDictionary() *cpe.CPEDictionary {
    dict := cpe.NewCPEDictionary()
    
    entries := []*cpe.CPEDictionaryEntry{
        {
            CPE23: "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
            Title: "Apache Tomcat 9.0.0",
        },
        {
            CPE23: "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
            Title: "Apache Log4j 2.14.1",
        },
        {
            CPE23: "cpe:2.3:a:oracle:java:1.8.0_291:*:*:*:*:*:*:*",
            Title: "Oracle Java SE 1.8.0_291",
        },
    }
    
    for _, entry := range entries {
        dict.AddEntry(entry)
    }
    
    return dict
}

func createSampleVulnerabilities() []*cpe.CVEEntry {
    return []*cpe.CVEEntry{
        {
            ID:            "CVE-2021-25122",
            BaseScore:     7.5,
            Severity:      "HIGH",
            Description:   "Apache Tomcat request smuggling vulnerability",
            PublishedDate: time.Date(2021, 3, 1, 0, 0, 0, 0, time.UTC),
        },
        {
            ID:            "CVE-2021-30640",
            BaseScore:     8.1,
            Severity:      "HIGH", 
            Description:   "Apache Tomcat vulnerability in HTTP/2 request mix-up",
            PublishedDate: time.Date(2021, 7, 12, 0, 0, 0, 0, time.UTC),
        },
    }
}

func createSampleCVE(cveID string) *cpe.CVEEntry {
    return &cpe.CVEEntry{
        ID:            cveID,
        BaseScore:     10.0,
        Severity:      "CRITICAL",
        Description:   "Apache Log4j2 JNDI features do not protect against attacker controlled LDAP and other JNDI related endpoints.",
        PublishedDate: time.Date(2021, 12, 10, 0, 0, 0, 0, time.UTC),
        AffectedCPEs: []string{
            "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*",
            "cpe:2.3:a:apache:log4j:2.14.1:*:*:*:*:*:*:*",
        },
    }
}

func truncateString(s string, maxLen int) string {
    if len(s) <= maxLen {
        return s
    }
    return s[:maxLen-3] + "..."
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func getRiskLevel(highSeverity, total int) string {
    if total == 0 {
        return "低"
    }
    
    ratio := float64(highSeverity) / float64(total)
    if ratio >= 0.5 {
        return "高"
    } else if ratio >= 0.2 {
        return "中"
    }
    return "低"
}

func saveReportToFile(report *VulnerabilityReport, filename string) error {
    // 在实际实现中，这里会将报告保存为JSON文件
    fmt.Printf("(模拟保存报告到 %s)\n", filename)
    return nil
}
```

## 关键概念

### 1. NVD数据源

- **CPE字典**: 官方CPE名称和描述
- **CVE数据**: 漏洞信息和影响的CPE
- **CVSS分数**: 漏洞严重程度评分
- **时间戳**: 发布和修改日期

### 2. API集成

- **速率限制**: 遵守NVD API限制
- **缓存**: 减少API调用
- **增量更新**: 只下载新数据
- **错误处理**: 处理网络和API错误

### 3. 安全分析

- **漏洞匹配**: 将系统组件与已知漏洞匹配
- **风险评估**: 基于CVSS分数评估风险
- **报告生成**: 创建安全评估报告
- **合规检查**: 验证CPE的官方状态

## 最佳实践

1. **使用API密钥**: 获取更高的速率限制
2. **启用缓存**: 减少重复的API调用
3. **定期更新**: 保持漏洞数据最新
4. **批量处理**: 高效处理大量CPE
5. **错误恢复**: 实现重试和降级机制

## 下一步

- 学习[CVE映射](./cve-mapping.md)进行详细的漏洞分析
- 探索[存储](./storage.md)来持久化NVD数据
- 查看[高级匹配](./advanced-matching.md)来改进漏洞检测
