# 存储操作

本示例演示如何使用CPE库的存储功能来持久化CPE数据，包括文件存储、内存存储和数据库存储。

## 概述

CPE存储功能提供了多种后端选项来持久化CPE数据，支持CRUD操作、批量处理、索引和查询功能。

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "time"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    fmt.Println("=== CPE存储操作示例 ===")
    
    // 示例1：文件存储
    fmt.Println("\n1. 文件存储:")
    
    // 创建文件存储
    fileStorage, err := cpeskills.NewFileStorage("./cpe_data", true) // 启用缓存
    if err != nil {
        log.Fatal(err)
    }
    defer fileStorage.Close()
    
    // 初始化存储
    err = fileStorage.Initialize()
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Println("✅ 文件存储初始化成功")
    
    // 创建测试CPE数据
    testCPEs := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
        "cpe:2.3:o:canonical:ubuntu:20.04:*:*:*:*:*:*:*",
        "cpe:2.3:h:cisco:catalyst_2960:*:*:*:*:*:*:*:*",
    }
    
    // 存储CPE数据
    fmt.Println("\n存储CPE数据:")
    for i, cpeStr := range testCPEs {
        cpeObj, err := cpeskills.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("解析CPE %s失败: %v", cpeStr, err)
            continue
        }
        
        err = fileStorage.Store(cpeObj)
        if err != nil {
            log.Printf("存储CPE失败: %v", err)
        } else {
            fmt.Printf("  ✅ %d. 已存储: %s %s\n", i+1, cpeObj.Vendor, cpeObj.ProductName)
        }
    }
    
    // 检索CPE数据
    fmt.Println("\n检索CPE数据:")
    for i, cpeStr := range testCPEs[:3] { // 检索前3个
        cpeObj, _ := cpeskills.ParseCpe23(cpeStr)
        
        retrieved, err := fileStorage.Retrieve(cpeObj.GetURI())
        if err != nil {
            log.Printf("检索失败: %v", err)
        } else {
            fmt.Printf("  ✅ %d. 检索到: %s %s %s\n", 
                i+1, retrieved.Vendor, retrieved.ProductName, retrieved.Version)
        }
    }
    
    // 列出所有CPE
    fmt.Println("\n列出所有存储的CPE:")
    allCPEs, err := fileStorage.List()
    if err != nil {
        log.Printf("列出CPE失败: %v", err)
    } else {
        fmt.Printf("总共存储了 %d 个CPE:\n", len(allCPEs))
        for i, cpeObj := range allCPEs {
            fmt.Printf("  %d. %s %s %s\n", 
                i+1, cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
        }
    }
    
    // 示例2：内存存储
    fmt.Println("\n2. 内存存储:")
    
    memoryStorage := cpeskills.NewMemoryStorage()
    
    // 批量存储到内存
    fmt.Println("批量存储到内存:")
    cpeObjects := make([]*cpeskills.CPE, 0, len(testCPEs))
    for _, cpeStr := range testCPEs {
        cpeObj, _ := cpeskills.ParseCpe23(cpeStr)
        cpeObjects = append(cpeObjects, cpeObj)
    }
    
    err = memoryStorage.StoreBatch(cpeObjects)
    if err != nil {
        log.Printf("批量存储失败: %v", err)
    } else {
        fmt.Printf("✅ 批量存储了 %d 个CPE到内存\n", len(cpeObjects))
    }
    
    // 内存存储统计
    memStats, err := memoryStorage.Stats()
    if err != nil {
        log.Printf("获取内存统计失败: %v", err)
    } else {
        fmt.Printf("内存存储统计:\n")
        fmt.Printf("  总数量: %d\n", memStats.TotalCount)
        fmt.Printf("  存储大小: %d 字节\n", memStats.StorageSize)
    }
    
    // 示例3：搜索功能
    fmt.Println("\n3. 搜索功能:")
    
    // 搜索Microsoft产品
    fmt.Println("搜索Microsoft产品:")
    microsoftResults, err := fileStorage.Search("microsoft")
    if err != nil {
        log.Printf("搜索失败: %v", err)
    } else {
        fmt.Printf("找到 %d 个Microsoft产品:\n", len(microsoftResults))
        for i, result := range microsoftResults {
            fmt.Printf("  %d. %s %s\n", i+1, result.ProductName, result.Version)
        }
    }
    
    // 搜索应用程序
    fmt.Println("\n搜索应用程序:")
    appResults, err := fileStorage.Search("part:a")
    if err != nil {
        log.Printf("搜索失败: %v", err)
    } else {
        fmt.Printf("找到 %d 个应用程序:\n", len(appResults))
        for i, result := range appResults {
            fmt.Printf("  %d. %s %s\n", i+1, result.Vendor, result.ProductName)
        }
    }
    
    // 示例4：高级查询
    fmt.Println("\n4. 高级查询:")
    
    // 构建复杂查询
    query := cpeskills.NewQuery().
        Filter("vendor", cpeskills.OpEquals, "apache").
        Filter("part", cpeskills.OpEquals, "a").
        SortBy("product", cpeskills.SortAsc).
        Limit(10)
    
    fmt.Println("执行高级查询 (Apache应用程序):")
    queryResults, err := fileStorage.Query(query)
    if err != nil {
        log.Printf("查询失败: %v", err)
    } else {
        fmt.Printf("查询结果 (%d 项):\n", len(queryResults))
        for i, result := range queryResults {
            fmt.Printf("  %d. %s %s\n", i+1, result.ProductName, result.Version)
        }
    }
    
    // 示例5：事务操作
    fmt.Println("\n5. 事务操作:")
    
    // 开始事务
    tx, err := fileStorage.BeginTransaction()
    if err != nil {
        log.Printf("开始事务失败: %v", err)
    } else {
        fmt.Println("开始事务操作:")
        
        // 在事务中添加新CPE
        newCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:mozilla:firefox:95.0:*:*:*:*:*:*:*")
        err = tx.Store(newCPE)
        if err != nil {
            tx.Rollback()
            log.Printf("事务存储失败: %v", err)
        } else {
            fmt.Printf("  ✅ 在事务中存储: %s %s\n", newCPE.Vendor, newCPE.ProductName)
            
            // 提交事务
            err = tx.Commit()
            if err != nil {
                log.Printf("提交事务失败: %v", err)
            } else {
                fmt.Println("  ✅ 事务提交成功")
            }
        }
    }
    
    // 示例6：数据备份和恢复
    fmt.Println("\n6. 数据备份和恢复:")
    
    // 备份数据
    backupPath := "./cpe_backup.json"
    err = fileStorage.Backup(backupPath)
    if err != nil {
        log.Printf("备份失败: %v", err)
    } else {
        fmt.Printf("✅ 数据已备份到: %s\n", backupPath)
    }
    
    // 创建新的存储实例用于恢复测试
    restoreStorage, err := cpeskills.NewFileStorage("./cpe_restore", false)
    if err != nil {
        log.Printf("创建恢复存储失败: %v", err)
    } else {
        defer restoreStorage.Close()
        
        err = restoreStorage.Initialize()
        if err != nil {
            log.Printf("初始化恢复存储失败: %v", err)
        } else {
            // 恢复数据
            err = restoreStorage.Restore(backupPath)
            if err != nil {
                log.Printf("恢复失败: %v", err)
            } else {
                fmt.Println("✅ 数据恢复成功")
                
                // 验证恢复的数据
                restoredCPEs, _ := restoreStorage.List()
                fmt.Printf("恢复了 %d 个CPE\n", len(restoredCPEs))
            }
        }
    }
    
    // 示例7：索引管理
    fmt.Println("\n7. 索引管理:")
    
    // 创建索引
    fmt.Println("创建索引:")
    indexes := []struct {
        field string
        type_ cpeskills.IndexType
    }{
        {"vendor", cpeskills.IndexTypeBTree},
        {"product", cpeskills.IndexTypeBTree},
        {"version", cpeskills.IndexTypeHash},
    }
    
    for _, idx := range indexes {
        err = fileStorage.CreateIndex(idx.field, idx.type_)
        if err != nil {
            log.Printf("创建索引 %s 失败: %v", idx.field, err)
        } else {
            fmt.Printf("  ✅ 创建索引: %s (%s)\n", idx.field, idx.type_)
        }
    }
    
    // 列出索引
    indexList, err := fileStorage.ListIndexes()
    if err != nil {
        log.Printf("列出索引失败: %v", err)
    } else {
        fmt.Printf("当前索引 (%d 个):\n", len(indexList))
        for i, index := range indexList {
            fmt.Printf("  %d. %s (%s)\n", i+1, index.Field, index.Type)
        }
    }
    
    // 示例8：性能监控
    fmt.Println("\n8. 性能监控:")
    
    // 执行性能测试
    fmt.Println("执行性能测试:")
    
    startTime := time.Now()
    
    // 批量插入测试
    batchSize := 100
    batchCPEs := make([]*cpeskills.CPE, 0, batchSize)
    
    for i := 0; i < batchSize; i++ {
        cpeStr := fmt.Sprintf("cpe:2.3:a:test:product_%d:1.0.%d:*:*:*:*:*:*:*", i, i)
        cpeObj, _ := cpeskills.ParseCpe23(cpeStr)
        batchCPEs = append(batchCPEs, cpeObj)
    }
    
    err = fileStorage.StoreBatch(batchCPEs)
    if err != nil {
        log.Printf("批量插入失败: %v", err)
    } else {
        insertTime := time.Since(startTime)
        fmt.Printf("  ✅ 批量插入 %d 个CPE，耗时: %v\n", batchSize, insertTime)
    }
    
    // 查询性能测试
    startTime = time.Now()
    searchResults, _ := fileStorage.Search("test")
    searchTime := time.Since(startTime)
    
    fmt.Printf("  ✅ 搜索查询返回 %d 个结果，耗时: %v\n", len(searchResults), searchTime)
    
    // 示例9：存储配置
    fmt.Println("\n9. 存储配置:")
    
    // 创建带自定义配置的存储
    config := &cpeskills.StorageConfig{
        Type:        "file",
        Path:        "./cpe_custom",
        EnableCache: true,
        CacheSize:   1000,
        Compression: true,
        Options: map[string]string{
            "index_type": "btree",
            "sync_mode":  "full",
        },
    }
    
    customStorage, err := cpeskills.NewStorageFromConfig(config)
    if err != nil {
        log.Printf("创建自定义存储失败: %v", err)
    } else {
        defer customStorage.Close()
        
        err = customStorage.Initialize()
        if err != nil {
            log.Printf("初始化自定义存储失败: %v", err)
        } else {
            fmt.Println("✅ 自定义存储配置成功")
            
            // 测试自定义存储
            testCPE, _ := cpeskills.ParseCpe23("cpe:2.3:a:custom:test:1.0:*:*:*:*:*:*:*")
            err = customStorage.Store(testCPE)
            if err != nil {
                log.Printf("自定义存储测试失败: %v", err)
            } else {
                fmt.Println("  ✅ 自定义存储测试成功")
            }
        }
    }
    
    // 示例10：存储统计和清理
    fmt.Println("\n10. 存储统计和清理:")
    
    // 获取最终统计信息
    finalStats, err := fileStorage.Stats()
    if err != nil {
        log.Printf("获取最终统计失败: %v", err)
    } else {
        fmt.Printf("最终存储统计:\n")
        fmt.Printf("  总CPE数量: %d\n", finalStats.TotalCount)
        fmt.Printf("  存储大小: %d 字节\n", finalStats.StorageSize)
        fmt.Printf("  索引大小: %d 字节\n", finalStats.IndexSize)
        fmt.Printf("  最后修改: %s\n", finalStats.LastModified.Format("2006-01-02 15:04:05"))
    }
    
    // 清理测试数据
    fmt.Println("\n清理测试数据:")
    testCPEsToDelete := []string{
        "cpe:2.3:a:test:product_0:1.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:test:product_1:1.0.1:*:*:*:*:*:*:*",
    }
    
    for i, cpeStr := range testCPEsToDelete {
        err = fileStorage.Delete(cpeStr)
        if err != nil {
            log.Printf("删除CPE失败: %v", err)
        } else {
            fmt.Printf("  ✅ %d. 已删除: %s\n", i+1, cpeStr)
        }
    }
    
    fmt.Println("\n✅ 存储操作示例完成")
}
```

## 关键概念

### 1. 存储后端

- **文件存储**: 基于文件系统，支持缓存
- **内存存储**: 快速访问，适用于测试
- **数据库存储**: 企业级，支持SQL和NoSQL

### 2. CRUD操作

- **Create**: 存储新的CPE对象
- **Read**: 检索和查询CPE数据
- **Update**: 修改现有CPE对象
- **Delete**: 删除CPE对象

### 3. 高级功能

- **事务**: 原子操作支持
- **索引**: 提高查询性能
- **备份/恢复**: 数据保护
- **批量操作**: 高效处理大量数据

### 4. 性能优化

- **缓存**: 减少磁盘I/O
- **索引**: 加速查询
- **批量处理**: 减少开销
- **压缩**: 节省存储空间

## 最佳实践

1. **选择合适的存储后端**: 根据数据量和性能需求选择
2. **使用事务**: 对关键操作使用事务确保一致性
3. **创建索引**: 为频繁查询的字段创建索引
4. **定期备份**: 实施定期备份策略
5. **监控性能**: 跟踪存储操作的性能指标

## 下一步

- 学习[NVD集成](./nvd-integration.md)来存储大规模数据
- 探索[CPE集合](./sets.md)进行批量存储操作
- 查看[高级匹配](./advanced-matching.md)来优化查询
