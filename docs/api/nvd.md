# NVD Integration

The CPE library provides comprehensive integration with the National Vulnerability Database (NVD), including downloading CPE feeds, parsing vulnerability data, and mapping CPEs to CVEs.

## NVD Data Types

### NVDCPEData

```go
type NVDCPEData struct {
    Dictionary *CPEDictionary // Official CPE dictionary
    MatchData  *CPEMatchData  // CPE match data
    LastUpdate time.Time      // Last update timestamp
}
```

Represents complete NVD CPE data including dictionary and match information.

### CPEMatchData

```go
type CPEMatchData struct {
    Matches     []*CPEMatch // CPE match entries
    GeneratedAt time.Time   // Data generation timestamp
}
```

Contains CPE match data from NVD feeds.

### CPEMatch

```go
type CPEMatch struct {
    CPE23Uri              string  // CPE 2.3 URI
    VersionStartIncluding string  // Version range start (inclusive)
    VersionStartExcluding string  // Version range start (exclusive)
    VersionEndIncluding   string  // Version range end (inclusive)
    VersionEndExcluding   string  // Version range end (exclusive)
    Vulnerable            bool    // Whether this CPE is vulnerable
    CVEs                  []string // Associated CVE IDs
}
```

Represents a single CPE match entry with version ranges and vulnerability information.

## NVD Feed Options

### NVDFeedOptions

```go
type NVDFeedOptions struct {
    BaseURL      string        // NVD base URL
    CacheDir     string        // Local cache directory
    Timeout      time.Duration // HTTP request timeout
    ShowProgress bool          // Show download progress
    UserAgent    string        // HTTP User-Agent header
    MaxRetries   int           // Maximum retry attempts
    RetryDelay   time.Duration // Delay between retries
}
```

Configuration options for NVD feed operations.

### DefaultNVDFeedOptions

```go
func DefaultNVDFeedOptions() *NVDFeedOptions
```

Returns default NVD feed options.

**Returns:**
- `*NVDFeedOptions` - Default configuration

**Example:**
```go
options := cpe.DefaultNVDFeedOptions()
options.CacheDir = "./nvd-cache"
options.ShowProgress = true
options.Timeout = 30 * time.Second
```

## Downloading NVD Data

### DownloadAllNVDData

```go
func DownloadAllNVDData(options *NVDFeedOptions) (*NVDCPEData, error)
```

Downloads and parses all NVD CPE data (dictionary and match data).

**Parameters:**
- `options` - Download options (can be `nil` for defaults)

**Returns:**
- `*NVDCPEData` - Complete NVD data
- `error` - Error if download or parsing fails

**Example:**
```go
// Download all NVD data
fmt.Println("Downloading NVD data...")
options := cpe.DefaultNVDFeedOptions()
options.CacheDir = "./nvd-cache"
options.ShowProgress = true

nvdData, err := cpe.DownloadAllNVDData(options)
if err != nil {
    log.Fatalf("Failed to download NVD data: %v", err)
}

fmt.Printf("Downloaded dictionary with %d items\n", len(nvdData.Dictionary.Items))
fmt.Printf("Downloaded %d match entries\n", len(nvdData.MatchData.Matches))
fmt.Printf("Last updated: %v\n", nvdData.LastUpdate)
```

### DownloadAndParseCPEDict

```go
func DownloadAndParseCPEDict(options *NVDFeedOptions) (*CPEDictionary, error)
```

Downloads and parses only the CPE dictionary.

**Parameters:**
- `options` - Download options

**Returns:**
- `*CPEDictionary` - CPE dictionary
- `error` - Error if download or parsing fails

**Example:**
```go
dictionary, err := cpe.DownloadAndParseCPEDict(options)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Dictionary contains %d CPE entries\n", len(dictionary.Items))
```

### DownloadAndParseCPEMatch

```go
func DownloadAndParseCPEMatch(options *NVDFeedOptions) (*CPEMatchData, error)
```

Downloads and parses only the CPE match data.

**Parameters:**
- `options` - Download options

**Returns:**
- `*CPEMatchData` - CPE match data
- `error` - Error if download or parsing fails

**Example:**
```go
matchData, err := cpe.DownloadAndParseCPEMatch(options)
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Match data contains %d entries\n", len(matchData.Matches))
```

## CVE Integration

### FindCVEsForCPE

```go
func (n *NVDCPEData) FindCVEsForCPE(cpe *CPE) []*CVEReference
```

Finds all CVEs associated with a specific CPE.

**Parameters:**
- `cpe` - CPE to search for

**Returns:**
- `[]*CVEReference` - Array of associated CVEs

**Example:**
```go
// Find CVEs for Apache Log4j
log4jCPE, _ := cpe.ParseCpe23("cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*")
cves := nvdData.FindCVEsForCPE(log4jCPE)

fmt.Printf("Found %d CVEs for Apache Log4j 2.0:\n", len(cves))
for _, cve := range cves {
    fmt.Printf("- %s (CVSS: %.1f): %s\n", 
        cve.ID, cve.CVSS, cve.Description)
}
```

### FindCPEsForCVE

```go
func (n *NVDCPEData) FindCPEsForCVE(cveID string) []*CPE
```

Finds all CPEs affected by a specific CVE.

**Parameters:**
- `cveID` - CVE identifier (e.g., "CVE-2021-44228")

**Returns:**
- `[]*CPE` - Array of affected CPEs

**Example:**
```go
// Find CPEs affected by Log4Shell
affectedCPEs := nvdData.FindCPEsForCVE("CVE-2021-44228")

fmt.Printf("CVE-2021-44228 affects %d CPEs:\n", len(affectedCPEs))
for _, cpe := range affectedCPEs[:10] { // Show first 10
    fmt.Printf("- %s\n", cpe.GetURI())
}
```

### SearchVulnerabilities

```go
func (n *NVDCPEData) SearchVulnerabilities(query string) []*CVEReference
```

Searches for vulnerabilities by keyword.

**Parameters:**
- `query` - Search query

**Returns:**
- `[]*CVEReference` - Array of matching CVEs

**Example:**
```go
// Search for remote code execution vulnerabilities
rceVulns := nvdData.SearchVulnerabilities("remote code execution")

fmt.Printf("Found %d RCE vulnerabilities:\n", len(rceVulns))
for _, vuln := range rceVulns[:5] { // Show first 5
    fmt.Printf("- %s: %s\n", vuln.ID, vuln.Description)
}
```

## Vulnerability Analysis

### GetVulnerabilityStats

```go
func (n *NVDCPEData) GetVulnerabilityStats() *VulnerabilityStats
```

Returns statistical information about vulnerabilities.

**Returns:**
- `*VulnerabilityStats` - Vulnerability statistics

```go
type VulnerabilityStats struct {
    TotalCVEs        int               // Total number of CVEs
    HighSeverity     int               // High severity CVEs (CVSS >= 7.0)
    MediumSeverity   int               // Medium severity CVEs (4.0 <= CVSS < 7.0)
    LowSeverity      int               // Low severity CVEs (CVSS < 4.0)
    TopVendors       map[string]int    // Most vulnerable vendors
    TopProducts      map[string]int    // Most vulnerable products
    RecentCVEs       []*CVEReference   // Recently published CVEs
}
```

**Example:**
```go
stats := nvdData.GetVulnerabilityStats()
fmt.Printf("Total CVEs: %d\n", stats.TotalCVEs)
fmt.Printf("High severity: %d\n", stats.HighSeverity)
fmt.Printf("Medium severity: %d\n", stats.MediumSeverity)
fmt.Printf("Low severity: %d\n", stats.LowSeverity)

fmt.Println("Most vulnerable vendors:")
for vendor, count := range stats.TopVendors {
    fmt.Printf("  %s: %d CVEs\n", vendor, count)
}
```

## Data Updates

### CheckForUpdates

```go
func CheckForUpdates(options *NVDFeedOptions) (*UpdateInfo, error)
```

Checks if newer NVD data is available.

**Parameters:**
- `options` - NVD feed options

**Returns:**
- `*UpdateInfo` - Update information
- `error` - Error if check fails

```go
type UpdateInfo struct {
    DictionaryUpdated bool      // Whether dictionary has updates
    MatchDataUpdated  bool      // Whether match data has updates
    LastModified      time.Time // Last modification time
    Size              int64     // Data size in bytes
}
```

**Example:**
```go
updateInfo, err := cpe.CheckForUpdates(options)
if err != nil {
    log.Printf("Failed to check for updates: %v", err)
} else {
    if updateInfo.DictionaryUpdated {
        fmt.Println("Dictionary updates available")
    }
    if updateInfo.MatchDataUpdated {
        fmt.Println("Match data updates available")
    }
}
```

### UpdateNVDData

```go
func UpdateNVDData(storage Storage, options *NVDFeedOptions) error
```

Updates stored NVD data if newer versions are available.

**Parameters:**
- `storage` - Storage interface for persisting data
- `options` - NVD feed options

**Returns:**
- `error` - Error if update fails

**Example:**
```go
// Check and update NVD data
err := cpe.UpdateNVDData(storage, options)
if err != nil {
    log.Printf("Failed to update NVD data: %v", err)
} else {
    fmt.Println("NVD data updated successfully")
}
```

## Caching

The library automatically caches downloaded NVD data to improve performance:

```go
// Configure caching
options := cpe.DefaultNVDFeedOptions()
options.CacheDir = "./nvd-cache"

// First download will fetch from NVD
nvdData1, _ := cpe.DownloadAllNVDData(options)

// Subsequent downloads will use cache if data is fresh
nvdData2, _ := cpe.DownloadAllNVDData(options)
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "github.com/scagogogo/cpe"
)

func main() {
    // Configure NVD options
    options := cpe.DefaultNVDFeedOptions()
    options.CacheDir = "./nvd-cache"
    options.ShowProgress = true
    
    // Download NVD data
    fmt.Println("Downloading NVD data...")
    nvdData, err := cpe.DownloadAllNVDData(options)
    if err != nil {
        log.Fatalf("Failed to download NVD data: %v", err)
    }
    
    fmt.Printf("Downloaded %d dictionary items\n", len(nvdData.Dictionary.Items))
    fmt.Printf("Downloaded %d match entries\n", len(nvdData.MatchData.Matches))
    
    // Analyze system CPEs for vulnerabilities
    systemCPEs := []string{
        "cpe:2.3:a:apache:log4j:2.0:*:*:*:*:*:*:*",
        "cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*",
        "cpe:2.3:o:microsoft:windows:10:*:*:*:*:*:*:*",
    }
    
    fmt.Println("\nVulnerability Analysis:")
    for _, cpeStr := range systemCPEs {
        cpeObj, err := cpe.ParseCpe23(cpeStr)
        if err != nil {
            log.Printf("Failed to parse %s: %v", cpeStr, err)
            continue
        }
        
        cves := nvdData.FindCVEsForCPE(cpeObj)
        fmt.Printf("\n%s:\n", cpeStr)
        fmt.Printf("  Found %d vulnerabilities\n", len(cves))
        
        // Show high-severity vulnerabilities
        highSeverity := 0
        for _, cve := range cves {
            if cve.CVSS >= 7.0 {
                highSeverity++
                fmt.Printf("  HIGH: %s (CVSS: %.1f) - %s\n", 
                    cve.ID, cve.CVSS, cve.Description[:100]+"...")
            }
        }
        
        if highSeverity == 0 {
            fmt.Printf("  No high-severity vulnerabilities found\n")
        }
    }
    
    // Get overall vulnerability statistics
    fmt.Println("\nOverall Vulnerability Statistics:")
    stats := nvdData.GetVulnerabilityStats()
    fmt.Printf("Total CVEs: %d\n", stats.TotalCVEs)
    fmt.Printf("High severity: %d (%.1f%%)\n", 
        stats.HighSeverity, 
        float64(stats.HighSeverity)/float64(stats.TotalCVEs)*100)
    
    // Search for specific vulnerabilities
    fmt.Println("\nSearching for 'remote code execution' vulnerabilities:")
    rceVulns := nvdData.SearchVulnerabilities("remote code execution")
    fmt.Printf("Found %d RCE vulnerabilities\n", len(rceVulns))
    
    for i, vuln := range rceVulns[:3] { // Show first 3
        fmt.Printf("%d. %s (CVSS: %.1f)\n", i+1, vuln.ID, vuln.CVSS)
        fmt.Printf("   %s\n", vuln.Description[:150]+"...")
    }
}
```
