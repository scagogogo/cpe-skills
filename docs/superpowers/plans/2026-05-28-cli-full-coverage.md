# Full CLI Coverage Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: `superpowers:subagent-driven-development`
> Steps use checkbox (`- [ ]`) syntax.

**Goal:** 将 cpe 库的所有能力暴露给 CLI，新增 validate、normalize、set、evaluate、cve、nvd 六个子命令，并增强现有 match 和 dict 子命令。

**Architecture:** 用户输入 → cobra 子命令 → 调用 cpe 库 API → 格式化输出。新增 6 个子命令文件 + 修改 2 个现有文件。每个新子命令复用已有的 `parseCPEString` 和 `outputCPE` 工具函数。

**Tech Stack:** Go 1.18, github.com/spf13/cobra v1.8.1

**Scope:** Medium

**Risk:** Low — 纯增量添加，不修改库代码

**Risks:**
- Task 6 (nvd) 涉及网络下载 → 缓解：支持 --cache-dir 和 --cache-max-age 参数，测试不依赖网络
- Task 5 (cve) 的 QueryByCVE/QueryByProduct 需要输入 CVE 列表 → 缓解：支持从 stdin/文件读取 JSON

**Autonomy Level:** Full

---

### Task 1: 创建 validate 和 normalize 子命令 — CPE 验证和标准化

**Depends on:** None (已有 CLI 骨架)
**Files:**
- Create: `cmd/cpe/validate.go`
- Create: `cmd/cpe/normalize.go`

- [ ] **Step 1: 创建 validate.go — CPE 验证子命令**

```go
// cmd/cpe/validate.go
package main

import (
	"fmt"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <cpe-string>",
	Short: "Validate a CPE string",
	Long: `Validate whether a CPE string is well-formed and its components
conform to the CPE 2.3 specification.

Checks include:
  - Valid format (CPE 2.2 or 2.3)
  - Valid part value (a, h, o)
  - No illegal characters in components
  - Required fields (Vendor, ProductName) are present

Examples:
  cpe validate "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
  cpe validate "cpe:/a:apache:log4j:2.0"`,
	Args: cobra.ExactArgs(1),
	RunE: runValidate,
}

func init() {
	rootCmd.AddCommand(validateCmd)
}

func runValidate(cmd *cobra.Command, args []string) error {
	input := args[0]

	// 首先验证能否解析
	c, err := parseCPEString(input)
	if err != nil {
		if outputFormat == "json" {
			fmt.Printf(`{"valid": false, "error": "%s"}`, err.Error())
			fmt.Println()
		} else {
			fmt.Printf("INVALID: %s\n", err)
		}
		return nil
	}

	// 然后验证CPE结构是否合规
	validateErr := cpe.ValidateCPE(c)
	if validateErr != nil {
		if outputFormat == "json" {
			fmt.Printf(`{"valid": false, "parseable": true, "error": "%s"}`, validateErr.Error())
			fmt.Println()
		} else {
			fmt.Printf("INVALID: parsed successfully but validation failed: %s\n", validateErr)
		}
		return nil
	}

	if outputFormat == "json" {
		fmt.Printf(`{"valid": true, "parseable": true, "uri": "%s"}`, c.GetURI())
		fmt.Println()
	} else {
		fmt.Printf("VALID: %s\n", c.GetURI())
	}
	return nil
}
```

- [ ] **Step 2: 创建 normalize.go — CPE 标准化子命令**

```go
// cmd/cpe/normalize.go
package main

import (
	"fmt"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var normalizeCmd = &cobra.Command{
	Use:   "normalize <cpe-string>",
	Short: "Normalize a CPE string to standard form",
	Long: `Normalize a CPE string to its standard CPE 2.3 form.
Converts all component values to lowercase, replaces spaces with
underscores, and collapses multiple underscores.

Examples:
  cpe normalize "cpe:2.3:a:Microsoft:Windows 10:*:*:*:*:*:*:*"
  cpe normalize "cpe:/a:Apache:Log4j:2.0"`,
	Args: cobra.ExactArgs(1),
	RunE: runNormalize,
}

func init() {
	rootCmd.AddCommand(normalizeCmd)
}

func runNormalize(cmd *cobra.Command, args []string) error {
	c, err := parseCPEString(args[0])
	if err != nil {
		return err
	}

	normalized := cpe.NormalizeCPE(c)

	if outputFormat == "json" {
		return outputCPE(cmd.OutOrStdout(), normalized, "json")
	}

	fmt.Printf("Original:   %s\n", c.GetURI())
	fmt.Printf("Normalized: %s\n", normalized.GetURI())
	return nil
}
```

- [ ] **Step 3: 验证 validate 和 normalize 子命令**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe validate "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*" && ./cpe normalize "cpe:2.3:a:Microsoft:Windows 10:*:*:*:*:*:*:*"
```

Expected:
  - Exit code: 0
  - validate output contains: "VALID"
  - normalize output contains: "microsoft" and "windows_10"

- [ ] **Step 4: 质量门禁检查**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./cmd/cpe && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 5: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/validate.go cmd/cpe/normalize.go && git commit -m "feat(cli): add validate and normalize subcommands"
```

---

### Task 2: 创建 set 子命令 — CPE 集合操作

**Depends on:** Task 1
**Files:**
- Create: `cmd/cpe/set.go`

- [ ] **Step 1: 创建 set.go — CPE 集合操作子命令**

```go
// cmd/cpe/set.go
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "CPE set operations",
	Long: `Perform set operations on CPE collections: union, intersection,
difference, filter, and sort. Reads CPE strings from stdin or files.`,
}

var setUnionCmd = &cobra.Command{
	Use:   "union <file1> <file2>",
	Short: "Compute union of two CPE sets from files",
	Args:  cobra.ExactArgs(2),
	RunE:  runSetUnion,
}

var setIntersectCmd = &cobra.Command{
	Use:   "intersect <file1> <file2>",
	Short: "Compute intersection of two CPE sets from files",
	Args:  cobra.ExactArgs(2),
	RunE:  runSetIntersect,
}

var setDiffCmd = &cobra.Command{
	Use:   "diff <file1> <file2>",
	Short: "Compute difference (file1 - file2) of two CPE sets",
	Args:  cobra.ExactArgs(2),
	RunE:  runSetDiff,
}

var setFilterCmd = &cobra.Command{
	Use:   "filter <criteria-cpe>",
	Short: "Filter CPEs from stdin by criteria",
	Args:  cobra.ExactArgs(1),
	RunE:  runSetFilter,
}

var setSortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort CPEs from stdin",
	RunE:  runSetSort,
}

var setSortBy string
var setDescending bool

func init() {
	setSortCmd.Flags().StringVar(&setSortBy, "by", "vendor", "Sort field (part|vendor|product|version)")
	setSortCmd.Flags().BoolVar(&setDescending, "desc", false, "Sort in descending order")

	setCmd.AddCommand(setUnionCmd)
	setCmd.AddCommand(setIntersectCmd)
	setCmd.AddCommand(setDiffCmd)
	setCmd.AddCommand(setFilterCmd)
	setCmd.AddCommand(setSortCmd)
	rootCmd.AddCommand(setCmd)
}

// readCPEsFromFile 从文件读取 CPE 列表
func readCPEsFromFile(path string) ([]*cpe.CPE, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening file: %w", err)
	}
	defer f.Close()
	return scanCPEs(f)
}

// scanCPEsFromStdin 从 stdin 读取 CPE 列表
func scanCPEsFromStdin() ([]*cpe.CPE, error) {
	return scanCPEs(os.Stdin)
}

// scanCPEs 通用扫描函数
func scanCPEs(reader *os.File) ([]*cpe.CPE, error) {
	var cpes []*cpe.CPE
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		c, err := parseCPEString(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Skipping invalid CPE: %s (%v)\n", line, err)
			continue
		}
		cpes = append(cpes, c)
	}
	return cpes, scanner.Err()
}

// outputCPEList 输出 CPE 列表
func outputCPEList(cpes []*cpe.CPE) {
	if outputFormat == "json" {
		fmt.Printf("[")
		for i, c := range cpes {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf(`"%s"`, c.GetURI())
		}
		fmt.Printf("]\n")
	} else {
		for i, c := range cpes {
			fmt.Printf("%d. %s\n", i+1, c.GetURI())
		}
	}
}

// cpesToSet 将 CPE 列表转为 CPESet
func cpesToSet(cpes []*cpe.CPE, name string) *cpe.CPESet {
	set := cpe.NewCPESet(name, "")
	for _, c := range cpes {
		set.Add(c)
	}
	return set
}

func runSetUnion(cmd *cobra.Command, args []string) error {
	cpes1, err := readCPEsFromFile(args[0])
	if err != nil {
		return err
	}
	cpes2, err := readCPEsFromFile(args[1])
	if err != nil {
		return err
	}
	set1 := cpesToSet(cpes1, "set1")
	set2 := cpesToSet(cpes2, "set2")
	result := set1.Union(set2)
	outputCPEList(result.ToSlice())
	return nil
}

func runSetIntersect(cmd *cobra.Command, args []string) error {
	cpes1, err := readCPEsFromFile(args[0])
	if err != nil {
		return err
	}
	cpes2, err := readCPEsFromFile(args[1])
	if err != nil {
		return err
	}
	set1 := cpesToSet(cpes1, "set1")
	set2 := cpesToSet(cpes2, "set2")
	result := set1.Intersection(set2)
	outputCPEList(result.ToSlice())
	return nil
}

func runSetDiff(cmd *cobra.Command, args []string) error {
	cpes1, err := readCPEsFromFile(args[0])
	if err != nil {
		return err
	}
	cpes2, err := readCPEsFromFile(args[1])
	if err != nil {
		return err
	}
	set1 := cpesToSet(cpes1, "set1")
	set2 := cpesToSet(cpes2, "set2")
	result := set1.Difference(set2)
	outputCPEList(result.ToSlice())
	return nil
}

func runSetFilter(cmd *cobra.Command, args []string) error {
	criteria, err := parseCPEString(args[0])
	if err != nil {
		return fmt.Errorf("parsing criteria CPE: %w", err)
	}
	cpes, err := scanCPEsFromStdin()
	if err != nil {
		return err
	}
	set := cpesToSet(cpes, "input")
	result := set.Filter(criteria, nil)
	outputCPEList(result.ToSlice())
	return nil
}

func runSetSort(cmd *cobra.Command, args []string) error {
	cpes, err := scanCPEsFromStdin()
	if err != nil {
		return err
	}
	set := cpesToSet(cpes, "input")
	sorted := set.Sort(setSortBy, !setDescending)
	outputCPEList(sorted)
	return nil
}
```

- [ ] **Step 2: 验证 set 子命令**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe set --help && echo -e "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*\ncpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*" | ./cpe set sort --by vendor
```

Expected:
  - Exit code: 0
  - Help output contains: "union", "intersect", "diff", "filter", "sort"
  - Sort output lists CPEs in vendor order

- [ ] **Step 3: 质量门禁检查**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./cmd/cpe && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 4: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/set.go && git commit -m "feat(cli): add set subcommand with union, intersect, diff, filter, sort"
```

---

### Task 3: 创建 evaluate 子命令 — CPE 适用性语言表达式求值

**Depends on:** Task 1
**Files:**
- Create: `cmd/cpe/evaluate.go`

- [ ] **Step 1: 创建 evaluate.go — 表达式求值子命令**

```go
// cmd/cpe/evaluate.go
package main

import (
	"fmt"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var evaluateCmd = &cobra.Command{
	Use:   "evaluate <expression> <target-cpe>",
	Short: "Evaluate a CPE applicability expression against a target",
	Long: `Evaluate a CPE applicability language expression against a
target CPE. Supports AND, OR, NOT logical operators and individual
CPE matching.

Expression syntax:
  - Single CPE: "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
  - AND: "AND(expr1, expr2, ...)"
  - OR: "OR(expr1, expr2, ...)"
  - NOT: "NOT(expr)"

Examples:
  cpe evaluate "cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*" "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
  cpe evaluate 'OR(cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*, cpe:2.3:a:microsoft:windows:11:*:*:*:*:*:*:*)' "cpe:2.3:a:microsoft:windows:11:*:*:*:*:*:*:*"
  cpe evaluate 'NOT(cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*)' "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*"`,
	Args: cobra.ExactArgs(2),
	RunE: runEvaluate,
}

func init() {
	rootCmd.AddCommand(evaluateCmd)
}

func runEvaluate(cmd *cobra.Command, args []string) error {
	expr, err := cpe.ParseExpression(args[0])
	if err != nil {
		return fmt.Errorf("parsing expression: %w", err)
	}

	target, err := parseCPEString(args[1])
	if err != nil {
		return fmt.Errorf("parsing target CPE: %w", err)
	}

	result := expr.Evaluate(target)

	if outputFormat == "json" {
		fmt.Printf(`{"matches": %t, "expression": "%s", "target": "%s"}`, result, expr.String(), target.GetURI())
		fmt.Println()
	} else {
		if result {
			fmt.Printf("MATCH: %s matches expression %s\n", target.GetURI(), expr.String())
		} else {
			fmt.Printf("NO MATCH: %s does not match expression %s\n", target.GetURI(), expr.String())
		}
	}
	return nil
}
```

- [ ] **Step 2: 验证 evaluate 子命令**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe evaluate "cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*" "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
```

Expected:
  - Exit code: 0
  - Output contains: "MATCH"

- [ ] **Step 3: 质量门禁检查**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./cmd/cpe && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 4: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/evaluate.go && git commit -m "feat(cli): add evaluate subcommand for CPE applicability expressions"
```

---

### Task 4: 创建 cve 子命令 — CVE 操作

**Depends on:** Task 1
**Files:**
- Create: `cmd/cpe/cve.go`

- [ ] **Step 1: 创建 cve.go — CVE 操作子命令**

```go
// cmd/cpe/cve.go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var cveCmd = &cobra.Command{
	Use:   "cve",
	Short: "CVE (Common Vulnerabilities and Exposures) operations",
	Long: `Operations on CVE identifiers: validate, extract from text,
sort, group by year, and remove duplicates.`,
}

var cveValidateCmd = &cobra.Command{
	Use:   "validate <cve-id>",
	Short: "Validate a CVE ID",
	Args:  cobra.ExactArgs(1),
	RunE:  runCVEValidate,
}

var cveExtractCmd = &cobra.Command{
	Use:   "extract [text-file]",
	Short: "Extract CVE IDs from text (stdin or file)",
	Args:  cobra.MaximumNArgs(1),
	RunE:  runCVEExtract,
}

var cveSortCmd = &cobra.Command{
	Use:   "sort",
	Short: "Sort CVE IDs from stdin (one per line)",
	RunE:  runCVESort,
}

var cveGroupCmd = &cobra.Command{
	Use:   "group",
	Short: "Group CVE IDs by year from stdin",
	RunE:  runCVEGroup,
}

var cveDedupCmd = &cobra.Command{
	Use:   "dedup",
	Short: "Remove duplicate CVE IDs from stdin",
	RunE:  runCVEDedup,
}

func init() {
	cveCmd.AddCommand(cveValidateCmd)
	cveCmd.AddCommand(cveExtractCmd)
	cveCmd.AddCommand(cveSortCmd)
	cveCmd.AddCommand(cveGroupCmd)
	cveCmd.AddCommand(cveDedupCmd)
	rootCmd.AddCommand(cveCmd)
}

func runCVEValidate(cmd *cobra.Command, args []string) error {
	cveID := args[0]
	valid := cpe.ValidateCVE(cveID)
	if outputFormat == "json" {
		fmt.Printf(`{"cve_id": "%s", "valid": %t}`, cveID, valid)
		fmt.Println()
	} else {
		if valid {
			fmt.Printf("VALID: %s\n", cveID)
		} else {
			fmt.Printf("INVALID: %s\n", cveID)
		}
	}
	return nil
}

func runCVEExtract(cmd *cobra.Command, args []string) error {
	var reader io.Reader
	if len(args) == 1 {
		f, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("opening file: %w", err)
		}
		defer f.Close()
		reader = f
	} else {
		reader = os.Stdin
	}

	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("reading input: %w", err)
	}

	cveIDs := cpe.ExtractCVEsFromText(string(data))
	if outputFormat == "json" {
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		return encoder.Encode(cveIDs)
	}
	for _, id := range cveIDs {
		fmt.Println(id)
	}
	return nil
}

func runCVESort(cmd *cobra.Command, args []string) error {
	cveIDs, err := readLinesFromStdin()
	if err != nil {
		return err
	}
	sorted := cpe.SortCVEs(cveIDs)
	for _, id := range sorted {
		fmt.Println(id)
	}
	return nil
}

func runCVEGroup(cmd *cobra.Command, args []string) error {
	cveIDs, err := readLinesFromStdin()
	if err != nil {
		return err
	}
	grouped := cpe.GroupCVEsByYear(cveIDs)
	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")
	return encoder.Encode(grouped)
}

func runCVEDedup(cmd *cobra.Command, args []string) error {
	cveIDs, err := readLinesFromStdin()
	if err != nil {
		return err
	}
	unique := cpe.RemoveDuplicateCVEs(cveIDs)
	for _, id := range unique {
		fmt.Println(id)
	}
	return nil
}

// readLinesFromStdin 从 stdin 读取非空行
func readLinesFromStdin() ([]string, error) {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		return nil, fmt.Errorf("reading stdin: %w", err)
	}
	var lines []string
	for _, line := range strings.Split(string(data), "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines, nil
}
```

- [ ] **Step 2: 验证 cve 子命令**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe cve validate "CVE-2021-44228" && echo "System affected by CVE-2021-44228 and cve-2022-12345" | ./cpe cve extract
```

Expected:
  - Exit code: 0
  - validate output contains: "VALID"
  - extract output contains: "CVE-2021-44228" and "CVE-2022-12345"

- [ ] **Step 3: 质量门禁检查**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./cmd/cpe && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 4: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/cve.go && git commit -m "feat(cli): add cve subcommand with validate, extract, sort, group, dedup"
```

---

### Task 5: 创建 nvd 子命令 — NVD 数据集成

**Depends on:** Task 1
**Files:**
- Create: `cmd/cpe/nvd.go`

- [ ] **Step 1: 创建 nvd.go — NVD 数据操作子命令**

```go
// cmd/cpe/nvd.go
package main

import (
	"encoding/json"
	"fmt"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var (
	nvdCacheDir     string
	nvdCacheMaxAge  int
	nvdNoProgress   bool
)

var nvdCmd = &cobra.Command{
	Use:   "nvd",
	Short: "NVD (National Vulnerability Database) integration",
	Long: `Download and query data from the National Vulnerability Database.
Supports downloading CPE dictionaries and CPE-CVE match data,
and querying CVEs for a specific CPE or CPEs for a specific CVE.

Data is cached locally to avoid repeated downloads.
Use --cache-dir to specify cache location and --cache-max-age
to control cache freshness (in hours, default 24).`,
}

var nvdDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download NVD data (CPE dictionary and match data)",
	RunE:  runNVDDownload,
}

var nvdLookupCveCmd = &cobra.Command{
	Use:   "lookup-cve <cpe-string>",
	Short: "Find CVEs associated with a CPE",
	Args:  cobra.ExactArgs(1),
	RunE:  runNVDLookupCVE,
}

var nvdLookupCpeCmd = &cobra.Command{
	Use:   "lookup-cpe <cve-id>",
	Short: "Find CPEs affected by a CVE",
	Args:  cobra.ExactArgs(1),
	RunE:  runNVDLookupCPE,
}

func init() {
	nvdCmd.PersistentFlags().StringVar(&nvdCacheDir, "cache-dir", "", "Cache directory (default: system temp dir)")
	nvdCmd.PersistentFlags().IntVar(&nvdCacheMaxAge, "cache-max-age", 24, "Cache max age in hours")
	nvdCmd.PersistentFlags().BoolVar(&nvdNoProgress, "no-progress", false, "Disable progress output")

	nvdCmd.AddCommand(nvdDownloadCmd)
	nvdCmd.AddCommand(nvdLookupCveCmd)
	nvdCmd.AddCommand(nvdLookupCpeCmd)
	rootCmd.AddCommand(nvdCmd)
}

func newNVDOptions() *cpe.NVDFeedOptions {
	opts := cpe.DefaultNVDFeedOptions()
	if nvdCacheDir != "" {
		opts.CacheDir = nvdCacheDir
	}
	opts.CacheMaxAge = nvdCacheMaxAge
	opts.ShowProgress = !nvdNoProgress
	return opts
}

func runNVDDownload(cmd *cobra.Command, args []string) error {
	opts := newNVDOptions()
	data, err := cpe.DownloadAllNVDData(opts)
	if err != nil {
		return fmt.Errorf("downloading NVD data: %w", err)
	}

	dictCount := 0
	if data.CPEDictionary != nil {
		dictCount = len(data.CPEDictionary.Items)
	}
	matchCount := 0
	if data.CPEMatchData != nil {
		matchCount = len(data.CPEMatchData.CPEToCVEs)
	}

	fmt.Printf("Download complete.\n")
	fmt.Printf("  CPE Dictionary items: %d\n", dictCount)
	fmt.Printf("  CPE-CVE mappings:     %d\n", matchCount)
	fmt.Printf("  Download time:        %s\n", data.DownloadTime.Format("2006-01-02 15:04:05"))
	return nil
}

func runNVDLookupCVE(cmd *cobra.Command, args []string) error {
	c, err := parseCPEString(args[0])
	if err != nil {
		return fmt.Errorf("parsing CPE: %w", err)
	}

	opts := newNVDOptions()
	matchData, err := cpe.DownloadAndParseCPEMatch(opts)
	if err != nil {
		return fmt.Errorf("downloading CPE match data: %w", err)
	}

	nvdData := &cpe.NVDCPEData{CPEMatchData: matchData}
	cves := nvdData.FindCVEsForCPE(c)

	if outputFormat == "json" {
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		return encoder.Encode(cves)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Found %d CVE(s) for %s:\n", len(cves), c.GetURI())
	for i, cveID := range cves {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", i+1, cveID)
	}
	return nil
}

func runNVDLookupCPE(cmd *cobra.Command, args []string) error {
	cveID := args[0]

	opts := newNVDOptions()
	matchData, err := cpe.DownloadAndParseCPEMatch(opts)
	if err != nil {
		return fmt.Errorf("downloading CPE match data: %w", err)
	}

	nvdData := &cpe.NVDCPEData{CPEMatchData: matchData}
	cpes := nvdData.FindCPEsForCVE(cveID)

	if outputFormat == "json" {
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		return encoder.Encode(cpes)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "Found %d CPE(s) affected by %s:\n", len(cpes), cveID)
	for i, c := range cpes {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", i+1, c.GetURI())
	}
	return nil
}
```

- [ ] **Step 2: 验证 nvd 子命令编译**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe nvd --help
```

Expected:
  - Exit code: 0
  - Output contains: "download", "lookup-cve", "lookup-cpe"

- [ ] **Step 3: 质量门禁检查**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./cmd/cpe && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 4: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/nvd.go && git commit -m "feat(cli): add nvd subcommand with download, lookup-cve, lookup-cpe"
```

---

### Task 6: 增强 match 和 dict 子命令 — 补全剩余能力

**Depends on:** Task 1
**Files:**
- Modify: `cmd/cpe/match.go`
- Modify: `cmd/cpe/dict.go`
- Modify: `cmd/cpe/parse.go`

- [ ] **Step 1: 增强 match.go — 添加高级匹配选项**

文件: `cmd/cpe/match.go`

在现有 matchCmd 的 init 函数中添加高级匹配 flags：

```go
// cmd/cpe/match.go — 增强后的完整文件
package main

import (
	"fmt"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var (
	matchIgnoreVersion bool
	matchUseRegex      bool
	matchMinVersion    string
	matchMaxVersion    string
	// 新增高级匹配选项
	matchAdvanced      bool
	matchFuzzy         bool
	matchIgnoreCase    bool
	matchCommonOnly    bool
	matchPartial       bool
	matchMode          string
	matchScoreThreshold float64
)

var matchCmd = &cobra.Command{
	Use:   "match <criteria-cpe> <target-cpe>",
	Short: "Check if two CPEs match",
	Long: `Compare two CPE strings and check if they match.
The first argument is the criteria, the second is the target.

Supports basic and advanced matching options.

Basic options:
  --ignore-version, --regex, --min-version, --max-version

Advanced options:
  --advanced          Use advanced matching algorithm
  --fuzzy             Enable fuzzy matching (requires --advanced)
  --ignore-case       Case-insensitive matching
  --match-mode        Match mode: exact, subset, superset, distance
  --score-threshold   Score threshold for distance matching (0.0-1.0)

Examples:
  cpe match "cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*" "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
  cpe match --advanced --fuzzy "cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*" "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*"
  cpe match --advanced --match-mode distance --score-threshold 0.7 "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*" "cpe:2.3:a:microsoft:windows:11:*:*:*:*:*:*:*"`,
	Args: cobra.ExactArgs(2),
	RunE: runMatch,
}

func init() {
	matchCmd.Flags().BoolVar(&matchIgnoreVersion, "ignore-version", false, "Ignore version when matching")
	matchCmd.Flags().BoolVar(&matchUseRegex, "regex", false, "Use regex matching for string fields")
	matchCmd.Flags().StringVar(&matchMinVersion, "min-version", "", "Minimum version for range matching")
	matchCmd.Flags().StringVar(&matchMaxVersion, "max-version", "", "Maximum version for range matching")
	// 高级匹配选项
	matchCmd.Flags().BoolVar(&matchAdvanced, "advanced", false, "Use advanced matching algorithm")
	matchCmd.Flags().BoolVar(&matchFuzzy, "fuzzy", false, "Enable fuzzy matching (requires --advanced)")
	matchCmd.Flags().BoolVar(&matchIgnoreCase, "ignore-case", false, "Case-insensitive matching (advanced)")
	matchCmd.Flags().BoolVar(&matchCommonOnly, "common-only", false, "Match only common fields (advanced)")
	matchCmd.Flags().BoolVar(&matchPartial, "partial", false, "Allow partial matching (advanced)")
	matchCmd.Flags().StringVar(&matchMode, "match-mode", "", "Match mode: exact, subset, superset, distance (advanced)")
	matchCmd.Flags().Float64Var(&matchScoreThreshold, "score-threshold", 0.0, "Score threshold for distance mode (0.0-1.0)")
	rootCmd.AddCommand(matchCmd)
}

func runMatch(cmd *cobra.Command, args []string) error {
	criteria, err := parseCPEString(args[0])
	if err != nil {
		return fmt.Errorf("parsing criteria CPE: %w", err)
	}

	target, err := parseCPEString(args[1])
	if err != nil {
		return fmt.Errorf("parsing target CPE: %w", err)
	}

	var result bool

	if matchAdvanced {
		opts := cpe.NewAdvancedMatchOptions()
		opts.UseFuzzyMatch = matchFuzzy
		opts.IgnoreCase = matchIgnoreCase
		opts.MatchCommonOnly = matchCommonOnly
		opts.PartialMatch = matchPartial
		if matchMode != "" {
			opts.MatchMode = matchMode
		}
		if matchScoreThreshold > 0 {
			opts.ScoreThreshold = matchScoreThreshold
		}
		result = cpe.AdvancedMatchCPE(criteria, target, opts)
	} else {
		options := &cpe.MatchOptions{
			IgnoreVersion:    matchIgnoreVersion,
			UseRegex:         matchUseRegex,
			AllowSubVersions: true,
			VersionRange:     matchMinVersion != "" || matchMaxVersion != "",
			MinVersion:       matchMinVersion,
			MaxVersion:       matchMaxVersion,
		}
		result = cpe.MatchCPE(criteria, target, options)
	}

	if outputFormat == "json" {
		fmt.Printf(`{"match": %t, "criteria": "%s", "target": "%s"}`, result, criteria.GetURI(), target.GetURI())
		fmt.Println()
	} else {
		if result {
			fmt.Printf("MATCH: %s matches %s\n", criteria.GetURI(), target.GetURI())
		} else {
			fmt.Printf("NO MATCH: %s does not match %s\n", criteria.GetURI(), target.GetURI())
		}
	}

	return nil
}
```

- [ ] **Step 2: 增强 dict.go — 添加 export 子命令**

文件: `cmd/cpe/dict.go`

在 dictCmd 的 init 中添加 export 子命令：

```go
// cmd/cpe/dict.go — 增强后完整文件
package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var dictCmd = &cobra.Command{
	Use:   "dict",
	Short: "CPE dictionary operations",
	Long: `Operations on CPE dictionaries: parse, search, and export.

A CPE dictionary is a collection of CPE items with metadata
like titles, references, and deprecation status.`,
}

var dictParseCmd = &cobra.Command{
	Use:   "parse <xml-file>",
	Short: "Parse a CPE dictionary XML file",
	Args:  cobra.ExactArgs(1),
	RunE:  runDictParse,
}

var dictSearchCmd = &cobra.Command{
	Use:   "search <xml-file> <criteria-cpe>",
	Short: "Search within a CPE dictionary",
	Args:  cobra.ExactArgs(2),
	RunE:  runDictSearch,
}

var dictExportCmd = &cobra.Command{
	Use:   "export <xml-file> [output-json]",
	Short: "Export CPE dictionary to JSON",
	Args:  cobra.MaximumNArgs(2),
	RunE:  runDictExport,
}

func init() {
	dictCmd.AddCommand(dictParseCmd)
	dictCmd.AddCommand(dictSearchCmd)
	dictCmd.AddCommand(dictExportCmd)
	rootCmd.AddCommand(dictCmd)
}

func runDictParse(cmd *cobra.Command, args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("opening dictionary file: %w", err)
	}
	defer f.Close()

	dict, err := cpe.ParseDictionary(f)
	if err != nil {
		return fmt.Errorf("parsing dictionary: %w", err)
	}

	if outputFormat == "json" {
		encoder := json.NewEncoder(cmd.OutOrStdout())
		encoder.SetIndent("", "  ")
		return encoder.Encode(dict)
	}

	fmt.Fprintf(cmd.OutOrStdout(), "CPE Dictionary\n")
	fmt.Fprintf(cmd.OutOrStdout(), "  Schema Version: %s\n", dict.SchemaVersion)
	fmt.Fprintf(cmd.OutOrStdout(), "  Generated At:   %s\n", dict.GeneratedAt)
	fmt.Fprintf(cmd.OutOrStdout(), "  Items:          %d\n\n", len(dict.Items))

	for i, item := range dict.Items {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s\n", i+1, item.Name)
		if item.Title != "" {
			fmt.Fprintf(cmd.OutOrStdout(), "   Title: %s\n", item.Title)
		}
		if item.Deprecated {
			fmt.Fprintf(cmd.OutOrStdout(), "   Deprecated: %s\n", item.DeprecationDate)
		}
	}
	return nil
}

func runDictSearch(cmd *cobra.Command, args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("opening dictionary file: %w", err)
	}
	defer f.Close()

	dict, err := cpe.ParseDictionary(f)
	if err != nil {
		return fmt.Errorf("parsing dictionary: %w", err)
	}

	criteria, err := parseCPEString(args[1])
	if err != nil {
		return fmt.Errorf("parsing criteria CPE: %w", err)
	}

	items := dict.FindItemsByCriteria(criteria, nil)

	fmt.Fprintf(cmd.OutOrStdout(), "Found %d matching item(s):\n", len(items))
	for i, item := range items {
		fmt.Fprintf(cmd.OutOrStdout(), "%d. %s", i+1, item.Name)
		if item.Title != "" {
			fmt.Fprintf(cmd.OutOrStdout(), " - %s", item.Title)
		}
		fmt.Fprintln(cmd.OutOrStdout())
	}
	return nil
}

func runDictExport(cmd *cobra.Command, args []string) error {
	f, err := os.Open(args[0])
	if err != nil {
		return fmt.Errorf("opening dictionary file: %w", err)
	}
	defer f.Close()

	dict, err := cpe.ParseDictionary(f)
	if err != nil {
		return fmt.Errorf("parsing dictionary: %w", err)
	}

	encoder := json.NewEncoder(cmd.OutOrStdout())
	encoder.SetIndent("", "  ")

	if len(args) == 2 {
		outFile, err := os.Create(args[1])
		if err != nil {
			return fmt.Errorf("creating output file: %w", err)
		}
		defer outFile.Close()
		encoder = json.NewEncoder(outFile)
		encoder.SetIndent("", "  ")
	}

	return encoder.Encode(dict)
}
```

- [ ] **Step 3: 增强 parse.go — 添加 WFN 解析输入支持**

文件: `cmd/cpe/parse.go`

修改 `parseCPEString` 函数以支持 WFN 格式输入：

```go
// cmd/cpe/parse.go — 增强后完整文件
package main

import (
	"fmt"
	"strings"

	"github.com/scagogogo/cpe"
	"github.com/spf13/cobra"
)

var parseFormat string

var parseCmd = &cobra.Command{
	Use:   "parse <cpe-string>",
	Short: "Parse a CPE string and display its components",
	Long: `Parse a CPE 2.2, 2.3, or WFN formatted string and display
its individual components (part, vendor, product, version, etc.).

The command automatically detects the input format:
  - CPE 2.3: starts with "cpe:2.3:"
  - CPE 2.2: starts with "cpe:/"
  - WFN: starts with "wfn:["

Examples:
  cpe parse "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
  cpe parse "cpe:/a:apache:log4j:2.0"
  cpe parse -o json "cpe:2.3:a:oracle:java:1.8.0:291:*:*:*:*:*:*"
  cpe parse -t wfn "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"`,
	Args: cobra.ExactArgs(1),
	RunE: runParse,
}

func init() {
	parseCmd.Flags().StringVarP(&parseFormat, "to", "t", "", "Convert to format (2.2|2.3|wfn)")
	rootCmd.AddCommand(parseCmd)
}

func runParse(cmd *cobra.Command, args []string) error {
	input := args[0]
	c, err := parseCPEString(input)
	if err != nil {
		return err
	}

	if parseFormat != "" {
		return outputConversion(c, parseFormat)
	}

	return outputCPE(cmd.OutOrStdout(), c, outputFormat)
}

// parseCPEString 自动检测 CPE 版本并解析（支持 2.2、2.3 和 WFN）
func parseCPEString(input string) (*cpe.CPE, error) {
	if len(input) >= 6 && input[:6] == "cpe:2." {
		return cpe.ParseCpe23(input)
	}
	if len(input) >= 5 && input[:5] == "cpe:/" {
		return cpe.ParseCpe22(input)
	}
	if strings.HasPrefix(input, "wfn:[") {
		wfn, err := cpe.FromCPE23String(input)
		if err != nil {
			// WFN 格式无法直接解析为 CPE 2.3，尝试通过 WFN 转换
			return nil, fmt.Errorf("WFN format input not directly parseable, use WFN string in CPE 2.3 form: %w", err)
		}
		return wfn.ToCPE(), nil
	}
	return nil, fmt.Errorf("unrecognized CPE format: %s (expected CPE 2.2, 2.3, or WFN)", input)
}

// outputConversion 输出格式转换结果
func outputConversion(c *cpe.CPE, format string) error {
	switch format {
	case "2.2":
		fmt.Println(cpe.FormatCpe22(c))
	case "2.3":
		fmt.Println(cpe.FormatCpe23(c))
	case "wfn":
		wfn := cpe.FromCPE(c)
		fmt.Printf("wfn:[part=%s,vendor=%s,product=%s,version=%s,update=%s,edition=%s,language=%s,swEdition=%s,targetSw=%s,targetHw=%s,other=%s]\n",
			wfn.Part, wfn.Vendor, wfn.Product, wfn.Version, wfn.Update, wfn.Edition, wfn.Language,
			wfn.SoftwareEdition, wfn.TargetSoftware, wfn.TargetHardware, wfn.Other)
	default:
		return fmt.Errorf("unsupported conversion format: %s (supported: 2.2, 2.3, wfn)", format)
	}
	return nil
}
```

- [ ] **Step 4: 验证所有增强功能**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go build ./cmd/cpe && ./cpe match --advanced --fuzzy "cpe:2.3:a:apache:log4j:*:*:*:*:*:*:*:*" "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*" && ./cpe dict --help && ./cpe parse -t wfn "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*"
```

Expected:
  - Exit code: 0
  - match output contains: "MATCH"
  - dict help contains: "export"
  - parse WFN output contains: "wfn:["

- [ ] **Step 5: 全量质量门禁**

```bash
cd /home/cc11001100/github/scagogogo/cpe && go vet ./... && go test ./...
```

Expected:
  - Exit code: 0

- [ ] **Step 6: 提交**

```bash
cd /home/cc11001100/github/scagogogo/cpe && git add cmd/cpe/match.go cmd/cpe/dict.go cmd/cpe/parse.go && git commit -m "feat(cli): enhance match with advanced options, add dict export, improve WFN support"
```
