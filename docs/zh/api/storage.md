# 存储接口

本页面描述了CPE库中用于数据持久化的存储接口和实现，包括文件存储、内存存储等多种存储后端。

## 存储接口

### Storage

核心存储接口定义。

```go
type Storage interface {
    // 初始化存储
    Initialize() error
    
    // 存储CPE对象
    Store(cpe *CPE) error
    
    // 根据ID检索CPE
    Retrieve(id string) (*CPE, error)
    
    // 删除CPE
    Delete(id string) error
    
    // 列出所有CPE
    List() ([]*CPE, error)
    
    // 搜索CPE
    Search(query string) ([]*CPE, error)
    
    // 获取存储统计信息
    Stats() (*StorageStats, error)
    
    // 关闭存储连接
    Close() error
}
```

### StorageStats

存储统计信息。

```go
type StorageStats struct {
    TotalCount    int64     // 总CPE数量
    LastModified  time.Time // 最后修改时间
    StorageSize   int64     // 存储大小（字节）
    IndexSize     int64     // 索引大小（字节）
}
```

## 文件存储

### FileStorage

基于文件系统的存储实现。

```go
type FileStorage struct {
    BaseDir     string // 基础目录
    EnableCache bool   // 是否启用缓存
    CacheSize   int    // 缓存大小
    IndexFile   string // 索引文件路径
}
```

#### 创建文件存储

```go
func NewFileStorage(baseDir string, enableCache bool) (*FileStorage, error)
```

**参数：**
- `baseDir`: 存储基础目录
- `enableCache`: 是否启用内存缓存

**示例：**
```go
storage, err := cpe.NewFileStorage("./cpe_data", true)
if err != nil {
    log.Fatal(err)
}
defer storage.Close()

// 初始化存储
err = storage.Initialize()
if err != nil {
    log.Fatal(err)
}
```

#### 文件存储方法

```go
// 存储CPE
func (fs *FileStorage) Store(cpe *CPE) error

// 检索CPE
func (fs *FileStorage) Retrieve(id string) (*CPE, error)

// 删除CPE
func (fs *FileStorage) Delete(id string) error

// 批量存储
func (fs *FileStorage) StoreBatch(cpes []*CPE) error

// 备份数据
func (fs *FileStorage) Backup(backupPath string) error

// 恢复数据
func (fs *FileStorage) Restore(backupPath string) error
```

**示例：**
```go
// 存储CPE
cpeObj, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
err = storage.Store(cpeObj)
if err != nil {
    log.Printf("存储失败: %v", err)
}

// 检索CPE
retrieved, err := storage.Retrieve(cpeObj.GetURI())
if err != nil {
    log.Printf("检索失败: %v", err)
} else {
    fmt.Printf("检索到: %s\n", retrieved.GetURI())
}
```

## 内存存储

### MemoryStorage

基于内存的存储实现，适用于测试和临时数据。

```go
type MemoryStorage struct {
    data  map[string]*CPE // 内存数据映射
    mutex sync.RWMutex    // 读写锁
}
```

#### 创建内存存储

```go
func NewMemoryStorage() *MemoryStorage
```

**示例：**
```go
storage := cpe.NewMemoryStorage()

// 内存存储不需要初始化
cpeObj, _ := cpe.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
err := storage.Store(cpeObj)
if err != nil {
    log.Printf("存储失败: %v", err)
}
```

#### 内存存储特性

- **快速访问**: 所有数据在内存中，访问速度极快
- **无持久化**: 程序重启后数据丢失
- **线程安全**: 使用读写锁保证并发安全
- **适用场景**: 测试、缓存、临时数据处理

## 数据库存储

### DatabaseStorage

基于数据库的存储实现。

```go
type DatabaseStorage struct {
    DB       *sql.DB // 数据库连接
    Driver   string  // 数据库驱动
    ConnStr  string  // 连接字符串
    TableName string // 表名
}
```

#### 支持的数据库

- **SQLite**: 轻量级文件数据库
- **PostgreSQL**: 企业级关系数据库
- **MySQL**: 流行的开源数据库
- **MongoDB**: NoSQL文档数据库

#### 创建数据库存储

```go
func NewDatabaseStorage(driver, connStr string) (*DatabaseStorage, error)
```

**示例：**
```go
// SQLite存储
storage, err := cpe.NewDatabaseStorage("sqlite3", "./cpe.db")
if err != nil {
    log.Fatal(err)
}

// PostgreSQL存储
storage, err := cpe.NewDatabaseStorage("postgres", 
    "host=localhost user=cpe dbname=cpe_db sslmode=disable")
if err != nil {
    log.Fatal(err)
}
```

## 存储配置

### StorageConfig

存储配置选项。

```go
type StorageConfig struct {
    Type        string            // 存储类型
    Path        string            // 存储路径
    Options     map[string]string // 额外选项
    EnableCache bool              // 启用缓存
    CacheSize   int               // 缓存大小
    Compression bool              // 启用压缩
    Encryption  bool              // 启用加密
}
```

### 配置示例

```go
config := &cpe.StorageConfig{
    Type:        "file",
    Path:        "./cpe_data",
    EnableCache: true,
    CacheSize:   1000,
    Compression: true,
    Options: map[string]string{
        "index_type": "btree",
        "sync_mode":  "full",
    },
}

storage, err := cpe.NewStorageFromConfig(config)
```

## 高级功能

### 索引管理

```go
// 创建索引
func (s *Storage) CreateIndex(field string, indexType IndexType) error

// 删除索引
func (s *Storage) DropIndex(field string) error

// 重建索引
func (s *Storage) RebuildIndex(field string) error

// 列出索引
func (s *Storage) ListIndexes() ([]IndexInfo, error)
```

**示例：**
```go
// 为供应商字段创建索引
err = storage.CreateIndex("vendor", cpe.IndexTypeBTree)
if err != nil {
    log.Printf("创建索引失败: %v", err)
}
```

### 事务支持

```go
// 开始事务
func (s *Storage) BeginTransaction() (Transaction, error)

// 事务接口
type Transaction interface {
    Store(cpe *CPE) error
    Delete(id string) error
    Commit() error
    Rollback() error
}
```

**示例：**
```go
tx, err := storage.BeginTransaction()
if err != nil {
    log.Fatal(err)
}

// 在事务中执行操作
err = tx.Store(cpe1)
if err != nil {
    tx.Rollback()
    return
}

err = tx.Store(cpe2)
if err != nil {
    tx.Rollback()
    return
}

// 提交事务
err = tx.Commit()
if err != nil {
    log.Printf("提交失败: %v", err)
}
```

### 批量操作

```go
// 批量存储
func (s *Storage) StoreBatch(cpes []*CPE) error

// 批量删除
func (s *Storage) DeleteBatch(ids []string) error

// 批量更新
func (s *Storage) UpdateBatch(updates map[string]*CPE) error
```

**示例：**
```go
cpes := []*CPE{cpe1, cpe2, cpe3}
err = storage.StoreBatch(cpes)
if err != nil {
    log.Printf("批量存储失败: %v", err)
}
```

## 查询功能

### Query

查询构建器。

```go
type Query struct {
    Filters    []Filter    // 过滤条件
    SortBy     string      // 排序字段
    SortOrder  SortOrder   // 排序顺序
    Limit      int         // 限制数量
    Offset     int         // 偏移量
}
```

### Filter

过滤条件。

```go
type Filter struct {
    Field    string      // 字段名
    Operator Operator    // 操作符
    Value    interface{} // 值
}
```

### 查询示例

```go
// 构建查询
query := cpe.NewQuery().
    Filter("vendor", cpe.OpEquals, "microsoft").
    Filter("part", cpe.OpEquals, "a").
    SortBy("product", cpe.SortAsc).
    Limit(10)

// 执行查询
results, err := storage.Query(query)
if err != nil {
    log.Printf("查询失败: %v", err)
} else {
    fmt.Printf("找到 %d 个结果\n", len(results))
}
```

## 缓存机制

### CacheConfig

缓存配置。

```go
type CacheConfig struct {
    Enabled    bool          // 是否启用
    Size       int           // 缓存大小
    TTL        time.Duration // 生存时间
    Strategy   CacheStrategy // 缓存策略
}
```

### 缓存策略

```go
const (
    CacheLRU  CacheStrategy = "lru"  // 最近最少使用
    CacheLFU  CacheStrategy = "lfu"  // 最少使用频率
    CacheFIFO CacheStrategy = "fifo" // 先进先出
)
```

**示例：**
```go
cacheConfig := &cpe.CacheConfig{
    Enabled:  true,
    Size:     1000,
    TTL:      30 * time.Minute,
    Strategy: cpe.CacheLRU,
}

storage.SetCacheConfig(cacheConfig)
```

## 数据迁移

### Migration

数据迁移接口。

```go
type Migration interface {
    // 迁移数据
    Migrate(source, target Storage) error
    
    // 验证迁移
    Validate(source, target Storage) error
    
    // 获取进度
    Progress() MigrationProgress
}
```

### 迁移示例

```go
// 从文件存储迁移到数据库存储
sourceStorage, _ := cpe.NewFileStorage("./old_data", false)
targetStorage, _ := cpe.NewDatabaseStorage("sqlite3", "./new_data.db")

migration := cpe.NewMigration()
err = migration.Migrate(sourceStorage, targetStorage)
if err != nil {
    log.Printf("迁移失败: %v", err)
}

// 验证迁移结果
err = migration.Validate(sourceStorage, targetStorage)
if err != nil {
    log.Printf("验证失败: %v", err)
}
```

## 完整示例

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    fmt.Println("=== 存储示例 ===")
    
    // 创建文件存储
    storage, err := cpe.NewFileStorage("./cpe_storage", true)
    if err != nil {
        log.Fatal(err)
    }
    defer storage.Close()
    
    // 初始化存储
    err = storage.Initialize()
    if err != nil {
        log.Fatal(err)
    }
    
    // 创建测试CPE
    cpes := []*cpe.CPE{}
    cpeStrings := []string{
        "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:oracle:java:11.0.12:*:*:*:*:*:*:*",
    }
    
    for _, cpeStr := range cpeStrings {
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("解析失败: %v", err)
            continue
        }
        cpes = append(cpes, cpeObj)
    }
    
    // 批量存储
    fmt.Println("存储CPE数据...")
    err = storage.StoreBatch(cpes)
    if err != nil {
        log.Printf("批量存储失败: %v", err)
    } else {
        fmt.Printf("成功存储 %d 个CPE\n", len(cpes))
    }
    
    // 列出所有CPE
    fmt.Println("\n列出所有存储的CPE:")
    allCPEs, err := storage.List()
    if err != nil {
        log.Printf("列出失败: %v", err)
    } else {
        for i, cpeObj := range allCPEs {
            fmt.Printf("  %d. %s %s %s\n", i+1, 
                cpeObj.Vendor, cpeObj.ProductName, cpeObj.Version)
        }
    }
    
    // 搜索CPE
    fmt.Println("\n搜索Microsoft产品:")
    results, err := storage.Search("microsoft")
    if err != nil {
        log.Printf("搜索失败: %v", err)
    } else {
        fmt.Printf("找到 %d 个结果:\n", len(results))
        for i, cpeObj := range results {
            fmt.Printf("  %d. %s\n", i+1, cpeObj.GetURI())
        }
    }
    
    // 获取统计信息
    fmt.Println("\n存储统计信息:")
    stats, err := storage.Stats()
    if err != nil {
        log.Printf("获取统计失败: %v", err)
    } else {
        fmt.Printf("  总数量: %d\n", stats.TotalCount)
        fmt.Printf("  存储大小: %d 字节\n", stats.StorageSize)
        fmt.Printf("  最后修改: %s\n", stats.LastModified.Format("2006-01-02 15:04:05"))
    }
    
    // 演示查询功能
    fmt.Println("\n高级查询示例:")
    query := cpe.NewQuery().
        Filter("vendor", cpe.OpEquals, "apache").
        SortBy("product", cpe.SortAsc).
        Limit(5)
    
    queryResults, err := storage.Query(query)
    if err != nil {
        log.Printf("查询失败: %v", err)
    } else {
        fmt.Printf("查询到 %d 个Apache产品:\n", len(queryResults))
        for i, cpeObj := range queryResults {
            fmt.Printf("  %d. %s\n", i+1, cpeObj.ProductName)
        }
    }
    
    // 演示事务
    fmt.Println("\n事务示例:")
    tx, err := storage.BeginTransaction()
    if err != nil {
        log.Printf("开始事务失败: %v", err)
    } else {
        // 在事务中添加新CPE
        newCPE, _ := cpe.ParseCpe23("cpe:2.3:a:mozilla:firefox:95.0:*:*:*:*:*:*:*")
        err = tx.Store(newCPE)
        if err != nil {
            tx.Rollback()
            log.Printf("事务存储失败: %v", err)
        } else {
            err = tx.Commit()
            if err != nil {
                log.Printf("提交事务失败: %v", err)
            } else {
                fmt.Println("事务提交成功")
            }
        }
    }
}
```

## 下一步

- 了解[字典管理](./dictionary.md)来处理CPE字典数据
- 学习[集合操作](./sets.md)来批量处理存储的CPE
- 探索[NVD集成](./nvd.md)来存储漏洞数据
