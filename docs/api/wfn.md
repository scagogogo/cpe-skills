# WFN (Well-Formed Name)

The CPE library provides comprehensive support for Well-Formed Names (WFN), which are the canonical internal representation of CPE names as defined in the CPE specification.

## WFN Structure

### WFN

```go
type WFN struct {
    Part            string // Component type
    Vendor          string // Vendor name
    Product         string // Product name
    Version         string // Version
    Update          string // Update
    Edition         string // Edition
    Language        string // Language
    SoftwareEdition string // Software edition
    TargetSoftware  string // Target software
    TargetHardware  string // Target hardware
    Other           string // Other attributes
}
```

The WFN structure represents the canonical form of a CPE name with all attributes properly normalized.

## Creating WFN Objects

### FromCPE

```go
func FromCPE(cpe *CPE) *WFN
```

Creates a WFN from a CPE object.

**Parameters:**
- `cpe` - CPE object to convert

**Returns:**
- `*WFN` - WFN representation

**Example:**
```go
// Create CPE and convert to WFN
cpeObj, _ := cpe.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
wfn := cpe.FromCPE(cpeObj)

fmt.Printf("WFN Part: %s\n", wfn.Part)
fmt.Printf("WFN Vendor: %s\n", wfn.Vendor)
fmt.Printf("WFN Product: %s\n", wfn.Product)
fmt.Printf("WFN Version: %s\n", wfn.Version)
```

### FromCPE23String

```go
func FromCPE23String(cpe23 string) (*WFN, error)
```

Creates a WFN directly from a CPE 2.3 format string.

**Parameters:**
- `cpe23` - CPE 2.3 format string

**Returns:**
- `*WFN` - WFN object
- `error` - Error if parsing fails

**Example:**
```go
wfn, err := cpe.FromCPE23String("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Vendor: %s, Product: %s, Version: %s\n", 
    wfn.Vendor, wfn.Product, wfn.Version)
```

### FromCPE22String

```go
func FromCPE22String(cpe22 string) (*WFN, error)
```

Creates a WFN from a CPE 2.2 format string.

**Parameters:**
- `cpe22` - CPE 2.2 format string

**Returns:**
- `*WFN` - WFN object
- `error` - Error if parsing fails

**Example:**
```go
wfn, err := cpe.FromCPE22String("cpe:/a:apache:tomcat:8.5.0")
if err != nil {
    log.Fatal(err)
}

fmt.Printf("Converted CPE 2.2 to WFN: %s %s %s\n", 
    wfn.Vendor, wfn.Product, wfn.Version)
```

## Converting from WFN

### ToCPE

```go
func (w *WFN) ToCPE() *CPE
```

Converts a WFN to a CPE object.

**Returns:**
- `*CPE` - CPE object representation

**Example:**
```go
// Create WFN and convert to CPE
wfn := &cpe.WFN{
    Part:    "a",
    Vendor:  "microsoft",
    Product: "windows",
    Version: "10",
}

cpeObj := wfn.ToCPE()
fmt.Printf("CPE URI: %s\n", cpeObj.GetURI())
```

### ToCPE23String

```go
func (w *WFN) ToCPE23String() string
```

Converts a WFN to CPE 2.3 format string.

**Returns:**
- `string` - CPE 2.3 format string

**Example:**
```go
wfn := &cpe.WFN{
    Part:    "a",
    Vendor:  "apache",
    Product: "tomcat",
    Version: "9.0.0",
}

cpe23 := wfn.ToCPE23String()
fmt.Printf("CPE 2.3: %s\n", cpe23)
// Output: cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*
```

### ToCPE22String

```go
func (w *WFN) ToCPE22String() string
```

Converts a WFN to CPE 2.2 format string.

**Returns:**
- `string` - CPE 2.2 format string

**Example:**
```go
wfn := &cpe.WFN{
    Part:    "a",
    Vendor:  "apache",
    Product: "tomcat",
    Version: "8.5.0",
}

cpe22 := wfn.ToCPE22String()
fmt.Printf("CPE 2.2: %s\n", cpe22)
// Output: cpe:/a:apache:tomcat:8.5.0
```

## WFN Matching

### WFNMatch

```go
func WFNMatch(wfn1, wfn2 *WFN) bool
```

Performs WFN-level matching between two WFN objects.

**Parameters:**
- `wfn1` - First WFN to compare
- `wfn2` - Second WFN to compare

**Returns:**
- `bool` - `true` if WFNs match, `false` otherwise

**Example:**
```go
wfn1, _ := cpe.FromCPE23String("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
wfn2, _ := cpe.FromCPE23String("cpe:2.3:a:microsoft:windows:*:*:*:*:*:*:*:*")

if cpe.WFNMatch(wfn2, wfn1) {
    fmt.Println("WFNs match")
}
```

### Match

```go
func (w *WFN) Match(other *WFN) bool
```

Instance method for WFN matching.

**Parameters:**
- `other` - WFN to match against

**Returns:**
- `bool` - `true` if WFNs match, `false` otherwise

**Example:**
```go
pattern, _ := cpe.FromCPE23String("cpe:2.3:a:*:*:*:*:*:*:*:*:*:*")
target, _ := cpe.FromCPE23String("cpe:2.3:a:apache:tomcat:9.0:*:*:*:*:*:*:*")

if pattern.Match(target) {
    fmt.Println("Target matches pattern")
}
```

## Value Handling

WFN uses special value handling for logical values:

### Special Values

- `ANY` (`*`) - Matches any value
- `NA` (`-`) - Not applicable/undefined

### FSStringToURI

```go
func FSStringToURI(fs string) string
```

Converts a formatted string to URI format.

**Parameters:**
- `fs` - Formatted string value

**Returns:**
- `string` - URI-encoded value

### URIToFSString

```go
func URIToFSString(uri string) string
```

Converts a URI-encoded value to formatted string.

**Parameters:**
- `uri` - URI-encoded value

**Returns:**
- `string` - Formatted string value

## Escape Handling

### EscapeValue

```go
func EscapeValue(value string) string
```

Escapes special characters in a value for WFN representation.

**Parameters:**
- `value` - Value to escape

**Returns:**
- `string` - Escaped value

**Example:**
```go
escaped := cpe.EscapeValue("product:name")
fmt.Printf("Escaped: %s\n", escaped) // product\:name
```

### UnescapeValue

```go
func UnescapeValue(value string) string
```

Unescapes a WFN value to its original form.

**Parameters:**
- `value` - Escaped value

**Returns:**
- `string` - Unescaped value

**Example:**
```go
unescaped := cpe.UnescapeValue("product\\:name")
fmt.Printf("Unescaped: %s\n", unescaped) // product:name
```

## WFN Validation

### ValidateWFN

```go
func ValidateWFN(wfn *WFN) error
```

Validates a WFN object for correctness.

**Parameters:**
- `wfn` - WFN to validate

**Returns:**
- `error` - Error if validation fails

**Example:**
```go
wfn := &cpe.WFN{
    Part:    "a",
    Vendor:  "apache",
    Product: "tomcat",
    Version: "9.0.0",
}

err := cpe.ValidateWFN(wfn)
if err != nil {
    log.Printf("WFN validation failed: %v", err)
} else {
    fmt.Println("WFN is valid")
}
```

## WFN Normalization

### NormalizeWFN

```go
func NormalizeWFN(wfn *WFN) *WFN
```

Normalizes a WFN by applying standard transformations.

**Parameters:**
- `wfn` - WFN to normalize

**Returns:**
- `*WFN` - Normalized WFN

**Example:**
```go
wfn := &cpe.WFN{
    Part:    "A",           // Will be normalized to "a"
    Vendor:  "Apache",      // Will be normalized to "apache"
    Product: "Tomcat",      // Will be normalized to "tomcat"
    Version: "9.0.0",
}

normalized := cpe.NormalizeWFN(wfn)
fmt.Printf("Normalized vendor: %s\n", normalized.Vendor) // apache
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
    // Create WFN from CPE 2.3 string
    fmt.Println("=== Creating WFN from CPE 2.3 ===")
    wfn1, err := cpe.FromCPE23String("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Part: %s\n", wfn1.Part)
    fmt.Printf("Vendor: %s\n", wfn1.Vendor)
    fmt.Printf("Product: %s\n", wfn1.Product)
    fmt.Printf("Version: %s\n", wfn1.Version)
    
    // Create WFN from CPE 2.2 string
    fmt.Println("\n=== Creating WFN from CPE 2.2 ===")
    wfn2, err := cpe.FromCPE22String("cpe:/a:microsoft:windows:10")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("Vendor: %s\n", wfn2.Vendor)
    fmt.Printf("Product: %s\n", wfn2.Product)
    fmt.Printf("Version: %s\n", wfn2.Version)
    
    // Convert WFN back to different formats
    fmt.Println("\n=== Converting WFN to different formats ===")
    cpe23 := wfn1.ToCPE23String()
    cpe22 := wfn1.ToCPE22String()
    
    fmt.Printf("CPE 2.3: %s\n", cpe23)
    fmt.Printf("CPE 2.2: %s\n", cpe22)
    
    // Convert to CPE object
    cpeObj := wfn1.ToCPE()
    fmt.Printf("CPE URI: %s\n", cpeObj.GetURI())
    
    // WFN matching
    fmt.Println("\n=== WFN Matching ===")
    pattern, _ := cpe.FromCPE23String("cpe:2.3:a:apache:*:*:*:*:*:*:*:*:*")
    target, _ := cpe.FromCPE23String("cpe:2.3:a:apache:tomcat:9.0.0:*:*:*:*:*:*:*")
    
    if pattern.Match(target) {
        fmt.Println("Target matches pattern")
    } else {
        fmt.Println("Target does not match pattern")
    }
    
    // Create WFN manually
    fmt.Println("\n=== Creating WFN manually ===")
    manualWFN := &cpe.WFN{
        Part:    "a",
        Vendor:  "oracle",
        Product: "java",
        Version: "11.0.12",
    }
    
    // Validate WFN
    err = cpe.ValidateWFN(manualWFN)
    if err != nil {
        log.Printf("WFN validation failed: %v", err)
    } else {
        fmt.Println("Manual WFN is valid")
    }
    
    // Convert manual WFN to CPE
    manualCPE := manualWFN.ToCPE()
    fmt.Printf("Manual WFN as CPE: %s\n", manualCPE.GetURI())
    
    // Demonstrate escape handling
    fmt.Println("\n=== Escape Handling ===")
    specialValue := "product:with:colons"
    escaped := cpe.EscapeValue(specialValue)
    unescaped := cpe.UnescapeValue(escaped)
    
    fmt.Printf("Original: %s\n", specialValue)
    fmt.Printf("Escaped: %s\n", escaped)
    fmt.Printf("Unescaped: %s\n", unescaped)
    
    // WFN normalization
    fmt.Println("\n=== WFN Normalization ===")
    unnormalizedWFN := &cpe.WFN{
        Part:    "A",
        Vendor:  "APACHE",
        Product: "TOMCAT",
        Version: "9.0.0",
    }
    
    normalizedWFN := cpe.NormalizeWFN(unnormalizedWFN)
    fmt.Printf("Original vendor: %s\n", unnormalizedWFN.Vendor)
    fmt.Printf("Normalized vendor: %s\n", normalizedWFN.Vendor)
}
```
