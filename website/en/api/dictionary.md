# Dictionary

The CPE library provides support for CPE dictionaries, including parsing NVD XML dictionaries, managing dictionary entries in memory, querying entries, and exporting dictionaries back to XML.

The diagram below shows the lifecycle of a CPE dictionary, from parsing raw XML through in-memory operations to exporting and storing.

```mermaid
flowchart TD
    subgraph "Parse"
        XML["NVD XML source"]
        Parse["ParseDictionary"]
        XML --> Parse
    end

    Parse --> Dict["CPEDictionary (in memory)"]

    subgraph "Operate"
        Find["FindItemByName / FindItemsByCriteria"]
        Add["AddItem / RemoveItem"]
    end

    Dict --> Find
    Dict --> Add

    subgraph "Output"
        Export["ExportDictionary (XML)"]
        Store["StoreDictionary (Storage)"]
    end

    Dict --> Export
    Dict --> Store
```

## Dictionary Types

### CPEDictionary

```go
type CPEDictionary struct {
    Items         []*CPEItem // CPE dictionary entries
    GeneratedAt   time.Time  // Dictionary generation timestamp
    SchemaVersion string     // CPE schema version the dictionary conforms to
}
```

Represents a complete CPE dictionary, typically from the National Vulnerability Database (NVD).

### CPEItem

```go
type CPEItem struct {
    Name            string      // CPE name (usually CPE 2.3 format)
    Title           string      // Human-readable title
    References      []Reference // Associated reference links
    Deprecated      bool        // Whether the CPE is deprecated
    DeprecationDate *time.Time  // Deprecation date, if deprecated
    CPE             *CPE        // Parsed CPE object
}
```

Represents a single entry in a CPE dictionary.

### Reference

```go
type Reference struct {
    URL  string // Reference URL
    Type string // Reference type, e.g. "Vendor", "Advisory", "External"
}
```

Represents a reference link associated with a CPE entry.

## Dictionary Parsing

### ParseDictionary

```go
func ParseDictionary(r io.Reader) (*CPEDictionary, error)
```

Parses a CPE dictionary from an XML data stream (typically NVD format).

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

dictionary, err := cpeskills.ParseDictionary(file)
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

## Dictionary Operations

### NewCPEItem

```go
func NewCPEItem(cpe *CPE, title string) *CPEItem
```

Creates a new CPE item from a parsed CPE and a title.

**Parameters:**
- `cpe` - Parsed CPE object
- `title` - Human-readable title

**Returns:**
- `*CPEItem` - New CPE item

**Example:**
```go
cpe, err := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
if err != nil {
    log.Fatal(err)
}

item := cpeskills.NewCPEItem(cpe, "Microsoft Windows 10")
fmt.Printf("Created item: %s\n", item.Name)
```

### AddItem

```go
func (d *CPEDictionary) AddItem(item *CPEItem)
```

Adds a CPE item to the dictionary. If an item with the same name already exists, it is replaced.

**Parameters:**
- `item` - CPE item to add

**Example:**
```go
cpe, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
item := cpeskills.NewCPEItem(cpe, "Apache Tomcat 9.0.0")

dictionary.AddItem(item)
fmt.Printf("Dictionary now has %d items\n", len(dictionary.Items))
```

### RemoveItem

```go
func (d *CPEDictionary) RemoveItem(name string) bool
```

Removes a CPE item from the dictionary by name. Returns `true` if an item was removed.

**Parameters:**
- `name` - CPE name of the item to remove

**Returns:**
- `bool` - Whether an item was removed

**Example:**
```go
removed := dictionary.RemoveItem("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
if removed {
    fmt.Println("Item removed")
} else {
    fmt.Println("Item not found")
}
```

### FindItemByName

```go
func (d *CPEDictionary) FindItemByName(name string) *CPEItem
```

Finds a dictionary item by its CPE name. Returns `nil` if no item matches.

**Parameters:**
- `name` - CPE name to look up

**Returns:**
- `*CPEItem` - Matching item, or `nil`

**Example:**
```go
item := dictionary.FindItemByName("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
if item != nil {
    fmt.Printf("Found: %s\n", item.Title)
} else {
    fmt.Println("Not found")
}
```

### FindItemsByCriteria

```go
func (d *CPEDictionary) FindItemsByCriteria(criteria *CPE, options *MatchOptions) []*CPEItem
```

Finds dictionary items whose parsed CPE matches the given criteria, using the provided match options.

**Parameters:**
- `criteria` - CPE to match against
- `options` - Match options (use `DefaultMatchOptions()` for defaults)

**Returns:**
- `[]*CPEItem` - Matching items

**Example:**
```go
// Match all Apache products, ignoring version
criteria, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*")
options := cpeskills.DefaultMatchOptions()
options.IgnoreVersion = true

results := dictionary.FindItemsByCriteria(criteria, options)
fmt.Printf("Found %d Apache items\n", len(results))

for _, item := range results {
    fmt.Printf("- %s: %s\n", item.Name, item.Title)
}
```

## Dictionary Storage

### StoreDictionary

```go
func (fs *FileStorage) StoreDictionary(dict *CPEDictionary) error
```

Stores a dictionary using the storage backend. `StoreDictionary` is part of the `Storage` interface and is implemented by both `FileStorage` and `MemoryStorage`.

**Example:**
```go
// Store in file storage
storage, err := cpeskills.NewFileStorage("./data", true)
if err != nil {
    log.Fatal(err)
}
defer storage.Close()

if err := storage.Initialize(); err != nil {
    log.Fatal(err)
}

if err := storage.StoreDictionary(dictionary); err != nil {
    log.Printf("Failed to store dictionary: %v", err)
} else {
    fmt.Println("Dictionary stored successfully")
}
```

### RetrieveDictionary

```go
func (fs *FileStorage) RetrieveDictionary() (*CPEDictionary, error)
```

Retrieves a stored dictionary. Returns `ErrNotFound` if no dictionary has been stored.

**Example:**
```go
dictionary, err := storage.RetrieveDictionary()
if err != nil {
    if errors.Is(err, cpeskills.ErrNotFound) {
        fmt.Println("No dictionary found")
    } else {
        log.Printf("Failed to retrieve dictionary: %v", err)
    }
} else {
    fmt.Printf("Retrieved dictionary with %d items\n", len(dictionary.Items))
}
```

## Dictionary Export

### ExportDictionary

```go
func ExportDictionary(dict *CPEDictionary, w io.Writer) error
```

Exports the dictionary to NVD-style XML format.

**Parameters:**
- `dict` - Dictionary to export
- `w` - Writer to output XML data

**Returns:**
- `error` - Error if export fails

**Example:**
```go
// Export to file
file, err := os.Create("dictionary.xml")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

if err := cpeskills.ExportDictionary(dictionary, file); err != nil {
    log.Printf("Failed to export dictionary: %v", err)
} else {
    fmt.Println("Dictionary exported to XML")
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "log"
    "os"

    cpeskills "github.com/scagogogo/cpe-skills"
)

func main() {
    // Parse dictionary from NVD XML file
    fmt.Println("Parsing CPE dictionary...")
    file, err := os.Open("official-cpe-dictionary_v2.3.xml")
    if err != nil {
        log.Fatalf("Failed to open dictionary file: %v", err)
    }
    defer file.Close()

    dictionary, err := cpeskills.ParseDictionary(file)
    if err != nil {
        log.Fatalf("Failed to parse dictionary: %v", err)
    }

    // Display basic information
    fmt.Printf("Dictionary loaded successfully!\n")
    fmt.Printf("Schema version: %s\n", dictionary.SchemaVersion)
    fmt.Printf("Generated at: %v\n", dictionary.GeneratedAt)
    fmt.Printf("Total items: %d\n", len(dictionary.Items))

    // Add a new item
    cpe, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    dictionary.AddItem(cpeskills.NewCPEItem(cpe, "Apache Tomcat 9.0.0"))

    // Find items matching a criteria
    fmt.Println("\nSearching for Apache products...")
    criteria, _ := cpeskills.ParseCpe23("cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*")
    options := cpeskills.DefaultMatchOptions()
    options.IgnoreVersion = true

    apacheItems := dictionary.FindItemsByCriteria(criteria, options)
    fmt.Printf("Found %d Apache products:\n", len(apacheItems))

    for i, item := range apacheItems {
        fmt.Printf("%d. %s\n", i+1, item.Title)
        if item.Deprecated {
            fmt.Printf("   (DEPRECATED)\n")
        }

        // Show references
        for _, ref := range item.References {
            fmt.Printf("   Reference (%s): %s\n", ref.Type, ref.URL)
        }
    }

    // Look up a specific item by name
    item := dictionary.FindItemByName("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    if item != nil {
        fmt.Printf("\nFound by name: %s\n", item.Title)
    }

    // Store dictionary
    fmt.Println("\nStoring dictionary...")
    storage, err := cpeskills.NewFileStorage("./dictionary-data", true)
    if err != nil {
        log.Fatal(err)
    }
    defer storage.Close()

    if err := storage.Initialize(); err != nil {
        log.Fatal(err)
    }

    if err := storage.StoreDictionary(dictionary); err != nil {
        log.Printf("Failed to store dictionary: %v", err)
    } else {
        fmt.Println("Dictionary stored successfully!")
    }

    // Export to XML
    fmt.Println("Exporting to XML...")
    out, err := os.Create("dictionary.xml")
    if err != nil {
        log.Fatal(err)
    }
    defer out.Close()

    if err := cpeskills.ExportDictionary(dictionary, out); err != nil {
        log.Printf("Failed to export dictionary: %v", err)
    } else {
        fmt.Println("Dictionary exported to dictionary.xml")
    }
}
```
