# Dictionary

The CPE library provides comprehensive support for CPE dictionaries, including parsing NVD XML dictionaries, managing dictionary entries, and performing dictionary-based operations.

## Dictionary Types

### CPEDictionary

```go
type CPEDictionary struct {
    SchemaVersion string     // XML schema version
    GeneratedAt   time.Time  // Dictionary generation timestamp
    Items         []*CPEItem // CPE dictionary entries
}
```

Represents a complete CPE dictionary, typically from the National Vulnerability Database (NVD).

### CPEItem

```go
type CPEItem struct {
    Name         string              // CPE name in URI format
    Title        string              // Human-readable title
    References   []*CPEReference     // Associated reference links
    Deprecated   bool                // Whether the CPE is deprecated
    DeprecatedBy []*CPEDeprecation   // Replacement CPEs if deprecated
}
```

Represents a single entry in a CPE dictionary.

### CPEReference

```go
type CPEReference struct {
    Href string // Reference URL
    Text string // Reference description/text
}
```

Represents a reference link associated with a CPE entry.

### CPEDeprecation

```go
type CPEDeprecation struct {
    Name string // Name of the replacement CPE
    Type string // Type of deprecation
}
```

Represents deprecation information for a CPE entry.

## Dictionary Parsing

### ParseDictionary

```go
func ParseDictionary(r io.Reader) (*CPEDictionary, error)
```

Parses a CPE dictionary from XML data stream (typically NVD format).

**Parameters:**
- `r` - XML data stream (from file or HTTP response)

**Returns:**
- `*CPEDictionary` - Parsed dictionary
- `error` - Error if parsing fails

**Example:**
```go
// Parse dictionary from file
file, err := os.Open("official-cpe-dictionary_v2.3.xml")
if err != nil {
    log.Fatalf("Failed to open dictionary file: %v", err)
}
defer file.Close()

dictionary, err := cpe.ParseDictionary(file)
if err != nil {
    log.Fatalf("Failed to parse dictionary: %v", err)
}

fmt.Printf("Dictionary contains %d CPE items\n", len(dictionary.Items))
fmt.Printf("Generated at: %v\n", dictionary.GeneratedAt)

// Display first 5 CPE items
for i, item := range dictionary.Items[:5] {
    fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Title)
}
```

### ParseDictionaryFromFile

```go
func ParseDictionaryFromFile(filename string) (*CPEDictionary, error)
```

Convenience function to parse a dictionary directly from a file.

**Parameters:**
- `filename` - Path to the XML dictionary file

**Returns:**
- `*CPEDictionary` - Parsed dictionary
- `error` - Error if parsing fails

**Example:**
```go
dictionary, err := cpe.ParseDictionaryFromFile("cpe-dictionary.xml")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Loaded %d CPE entries\n", len(dictionary.Items))
```

## Dictionary Operations

### Search Dictionary

```go
func (d *CPEDictionary) Search(query string) []*CPEItem
```

Searches the dictionary for entries matching the query.

**Parameters:**
- `query` - Search query string

**Returns:**
- `[]*CPEItem` - Array of matching dictionary items

**Example:**
```go
// Search for Microsoft products
results := dictionary.Search("microsoft")
fmt.Printf("Found %d Microsoft entries\n", len(results))

for _, item := range results {
    fmt.Printf("- %s: %s\n", item.Name, item.Title)
}
```

### Filter by Vendor

```go
func (d *CPEDictionary) FilterByVendor(vendor string) []*CPEItem
```

Filters dictionary entries by vendor name.

**Parameters:**
- `vendor` - Vendor name to filter by

**Returns:**
- `[]*CPEItem` - Array of entries from the specified vendor

**Example:**
```go
apacheItems := dictionary.FilterByVendor("apache")
fmt.Printf("Found %d Apache products\n", len(apacheItems))
```

### Filter by Product

```go
func (d *CPEDictionary) FilterByProduct(product string) []*CPEItem
```

Filters dictionary entries by product name.

**Parameters:**
- `product` - Product name to filter by

**Returns:**
- `[]*CPEItem` - Array of entries for the specified product

### Get Statistics

```go
func (d *CPEDictionary) GetStatistics() *DictionaryStats
```

Returns statistical information about the dictionary.

**Returns:**
- `*DictionaryStats` - Dictionary statistics

```go
type DictionaryStats struct {
    TotalItems      int               // Total number of items
    VendorCount     int               // Number of unique vendors
    ProductCount    int               // Number of unique products
    DeprecatedCount int               // Number of deprecated items
    TopVendors      map[string]int    // Top vendors by product count
    TopProducts     map[string]int    // Top products by version count
}
```

**Example:**
```go
stats := dictionary.GetStatistics()
fmt.Printf("Total items: %d\n", stats.TotalItems)
fmt.Printf("Unique vendors: %d\n", stats.VendorCount)
fmt.Printf("Deprecated items: %d\n", stats.DeprecatedCount)

fmt.Println("Top vendors:")
for vendor, count := range stats.TopVendors {
    fmt.Printf("  %s: %d products\n", vendor, count)
}
```

## Dictionary Storage

### Store Dictionary

```go
func (s *Storage) StoreDictionary(dict *CPEDictionary) error
```

Stores a dictionary using the storage interface.

**Example:**
```go
// Parse dictionary
dictionary, err := cpe.ParseDictionaryFromFile("cpe-dictionary.xml")
if err != nil {
    log.Fatal(err)
}

// Store in file storage
storage, _ := cpe.NewFileStorage("./data", true)
storage.Initialize()

err = storage.StoreDictionary(dictionary)
if err != nil {
    log.Printf("Failed to store dictionary: %v", err)
} else {
    fmt.Println("Dictionary stored successfully")
}
```

### Retrieve Dictionary

```go
func (s *Storage) RetrieveDictionary() (*CPEDictionary, error)
```

Retrieves a stored dictionary.

**Example:**
```go
dictionary, err := storage.RetrieveDictionary()
if err != nil {
    if errors.Is(err, cpe.ErrNotFound) {
        fmt.Println("No dictionary found")
    } else {
        log.Printf("Failed to retrieve dictionary: %v", err)
    }
} else {
    fmt.Printf("Retrieved dictionary with %d items\n", len(dictionary.Items))
}
```

## Dictionary Validation

### Validate Dictionary

```go
func ValidateDictionary(dict *CPEDictionary) error
```

Validates a dictionary for consistency and correctness.

**Parameters:**
- `dict` - Dictionary to validate

**Returns:**
- `error` - Error if validation fails

**Example:**
```go
err := cpe.ValidateDictionary(dictionary)
if err != nil {
    log.Printf("Dictionary validation failed: %v", err)
} else {
    fmt.Println("Dictionary is valid")
}
```

## Dictionary Merging

### Merge Dictionaries

```go
func MergeDictionaries(dict1, dict2 *CPEDictionary) *CPEDictionary
```

Merges two dictionaries into a single dictionary.

**Parameters:**
- `dict1` - First dictionary
- `dict2` - Second dictionary

**Returns:**
- `*CPEDictionary` - Merged dictionary

**Example:**
```go
// Load two dictionaries
dict1, _ := cpe.ParseDictionaryFromFile("dict1.xml")
dict2, _ := cpe.ParseDictionaryFromFile("dict2.xml")

// Merge them
merged := cpe.MergeDictionaries(dict1, dict2)
fmt.Printf("Merged dictionary has %d items\n", len(merged.Items))
```

## Dictionary Export

### Export to JSON

```go
func (d *CPEDictionary) ExportToJSON(w io.Writer) error
```

Exports the dictionary to JSON format.

**Parameters:**
- `w` - Writer to output JSON data

**Returns:**
- `error` - Error if export fails

**Example:**
```go
// Export to file
file, err := os.Create("dictionary.json")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

err = dictionary.ExportToJSON(file)
if err != nil {
    log.Printf("Failed to export dictionary: %v", err)
} else {
    fmt.Println("Dictionary exported to JSON")
}
```

### Export to CSV

```go
func (d *CPEDictionary) ExportToCSV(w io.Writer) error
```

Exports the dictionary to CSV format.

**Example:**
```go
// Export to CSV file
file, err := os.Create("dictionary.csv")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

err = dictionary.ExportToCSV(file)
if err != nil {
    log.Printf("Failed to export to CSV: %v", err)
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "os"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse dictionary from NVD XML file
    fmt.Println("Parsing CPE dictionary...")
    dictionary, err := cpe.ParseDictionaryFromFile("official-cpe-dictionary_v2.3.xml")
    if err != nil {
        log.Fatalf("Failed to parse dictionary: %v", err)
    }
    
    // Display basic information
    fmt.Printf("Dictionary loaded successfully!\n")
    fmt.Printf("Schema version: %s\n", dictionary.SchemaVersion)
    fmt.Printf("Generated at: %v\n", dictionary.GeneratedAt)
    fmt.Printf("Total items: %d\n", len(dictionary.Items))
    
    // Get statistics
    stats := dictionary.GetStatistics()
    fmt.Printf("Unique vendors: %d\n", stats.VendorCount)
    fmt.Printf("Deprecated items: %d\n", stats.DeprecatedCount)
    
    // Search for specific products
    fmt.Println("\nSearching for Apache products...")
    apacheItems := dictionary.FilterByVendor("apache")
    fmt.Printf("Found %d Apache products:\n", len(apacheItems))
    
    for i, item := range apacheItems[:10] { // Show first 10
        fmt.Printf("%d. %s\n", i+1, item.Title)
        if item.Deprecated {
            fmt.Printf("   (DEPRECATED)\n")
        }
    }
    
    // Search by query
    fmt.Println("\nSearching for 'tomcat'...")
    tomcatItems := dictionary.Search("tomcat")
    fmt.Printf("Found %d items containing 'tomcat':\n", len(tomcatItems))
    
    for _, item := range tomcatItems[:5] { // Show first 5
        fmt.Printf("- %s: %s\n", item.Name, item.Title)
        
        // Show references
        if len(item.References) > 0 {
            fmt.Printf("  References:\n")
            for _, ref := range item.References {
                fmt.Printf("    - %s: %s\n", ref.Text, ref.Href)
            }
        }
    }
    
    // Store dictionary
    fmt.Println("\nStoring dictionary...")
    storage, err := cpe.NewFileStorage("./dictionary-data", true)
    if err != nil {
        log.Fatal(err)
    }
    defer storage.Close()
    
    err = storage.Initialize()
    if err != nil {
        log.Fatal(err)
    }
    
    err = storage.StoreDictionary(dictionary)
    if err != nil {
        log.Printf("Failed to store dictionary: %v", err)
    } else {
        fmt.Println("Dictionary stored successfully!")
    }
    
    // Export to JSON
    fmt.Println("Exporting to JSON...")
    jsonFile, err := os.Create("dictionary.json")
    if err != nil {
        log.Fatal(err)
    }
    defer jsonFile.Close()
    
    err = dictionary.ExportToJSON(jsonFile)
    if err != nil {
        log.Printf("Failed to export to JSON: %v", err)
    } else {
        fmt.Println("Dictionary exported to dictionary.json")
    }
}
```
