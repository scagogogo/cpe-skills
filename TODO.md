# TODO

## 文档相关
- [x] 创建完整的API文档网站
- [x] 配置GitHub Pages自动部署
- [x] 添加使用示例和教程

## SCA Foundation Enhancement (2026-06-15)
- [x] Phase 1: PURL & Package Ecosystem Foundation (purl.go, ecosystem.go, cpe_purl_mapping.go)
- [x] Phase 2: SBOM Data Model & Serialization (sbom.go, sbom_cyclonedx.go, sbom_spdx.go)
- [x] Phase 3: Enhanced Vulnerability Integration (osv.go, vulnerability_report.go, remediation.go)
- [x] Phase 4: Dependency Graph & Resolution (dependency_graph.go)
- [x] Phase 5: Risk Scoring & Prioritization (risk_scoring.go)
- [x] Phase 6: Batch Processing & Performance (index.go, batch.go)
- [x] Phase 7-8: Export Formats & License Detection (export.go, license.go, license_detection.go)

## 功能增强
- [x] 添加更多示例文档
- [x] 性能优化 (CPEIndex 索引, BatchScanner 并发批处理)
- [x] 添加更多测试用例 (91.6% coverage 主包, 78.7% 全项目)

## 代码质量修复 (2026-06-24)
- [x] 修复 osv_test.go 编译错误 (time.Time 字段字符串赋值, stringsSprintf 未定义函数)
- [x] 修复重复测试函数声明 (TestNormalizeCPE, TestParseOSVEntry, TestParseOSVEntries)
- [x] 修复 pkg/parsers ParseGoMod 块内直接依赖无法解析
- [x] 修复 KEV 测试缓存过期时间设置错误
- [x] 修复 SBOM Diff componentKey 版本问题导致 changed 检测失败
- [x] 补充 manifest_bridge 测试 (0% → 覆盖)
- [x] 补充 reachability 测试 (0% → 覆盖)
- [x] 补充 EPSS HTTP mock 测试
- [x] 补充 CPEIndex Add/Remove/Clear 测试
- [x] 补充 ExportSBOM 导出测试
- [x] 修复 Cargo.toml 表格式依赖解析
- [x] Race detector 通过

## 维护
- [ ] 定期更新NVD数据
- [ ] 监控文档网站状态
