# Errors

The CPE library provides a comprehensive error handling system with structured error types for different failure scenarios.

## Error Types

### ErrorType

```go
type ErrorType int

const (
    ErrorTypeParsingFailed    ErrorType = iota // Parsing operation failed
    ErrorTypeInvalidFormat                     // Invalid format detected
    ErrorTypeInvalidPart                       // Invalid part value
    ErrorTypeInvalidAttribute                  // Invalid attribute value
    ErrorTypeNotFound                          // Resource not found
    ErrorTypeOperationFailed                   // General operation failure
)
```

Enumeration of different error types that can occur in the library.

### CPEError

```go
type CPEError struct {
    Type      ErrorType // Error type classification
    Message   string    // Human-readable error message
    CPEString string    // Related CPE string (if applicable)
    Err       error     // Underlying error (if any)
}
```

The main error type used throughout the library, providing structured error information.

#### Methods

##### Error

```go
func (e *CPEError) Error() string
```

Implements the `error` interface, returning a formatted error message.

**Returns:**
- `string` - Formatted error message

**Example:**
```go
err := cpeskills.NewInvalidFormatError("invalid:cpe:format")
fmt.Printf("Error: %s\n", err.Error())
// Output: invalid CPE format: invalid:cpe:format
```

##### Unwrap

```go
func (e *CPEError) Unwrap() error
```

Returns the underlying error for error chain unwrapping.

**Returns:**
- `error` - Underlying error or `nil`

**Example:**
```go
originalErr := errors.New("underlying issue")
cpeErr := cpeskills.NewOperationFailedError("parse", originalErr)

unwrapped := cpeErr.Unwrap()
fmt.Printf("Underlying error: %v\n", unwrapped)
```

## Error Constructors

### NewParsingError

```go
func NewParsingError(cpeString string, err error) *CPEError
```

Creates a parsing error for when CPE string parsing fails.

**Parameters:**
- `cpeString` - The CPE string that failed to parse
- `err` - Underlying parsing error

**Returns:**
- `*CPEError` - Parsing error

**Example:**
```go
err := cpeskills.NewParsingError("invalid:format", errors.New("malformed string"))
fmt.Printf("Parsing error: %v\n", err)
```

### NewInvalidFormatError

```go
func NewInvalidFormatError(cpeString string) *CPEError
```

Creates an error for invalid CPE format.

**Parameters:**
- `cpeString` - The invalid CPE string

**Returns:**
- `*CPEError` - Invalid format error

**Example:**
```go
err := cpeskills.NewInvalidFormatError("not:a:valid:cpe")
fmt.Printf("Format error: %v\n", err)
```

### NewInvalidPartError

```go
func NewInvalidPartError(part string) *CPEError
```

Creates an error for invalid CPE part values.

**Parameters:**
- `part` - The invalid part value

**Returns:**
- `*CPEError` - Invalid part error

**Example:**
```go
err := cpeskills.NewInvalidPartError("x") // Valid parts are "a", "h", "o"
fmt.Printf("Part error: %v\n", err)
```

### NewInvalidAttributeError

```go
func NewInvalidAttributeError(attribute, value string) *CPEError
```

Creates an error for invalid attribute values.

**Parameters:**
- `attribute` - Name of the invalid attribute
- `value` - The invalid value

**Returns:**
- `*CPEError` - Invalid attribute error

**Example:**
```go
err := cpeskills.NewInvalidAttributeError("vendor", "invalid\x00vendor")
fmt.Printf("Attribute error: %v\n", err)
```

### NewNotFoundError

```go
func NewNotFoundError(what string) *CPEError
```

Creates an error for when a resource is not found.

**Parameters:**
- `what` - Description of what was not found

**Returns:**
- `*CPEError` - Not found error

**Example:**
```go
err := cpeskills.NewNotFoundError("CPE dictionary")
fmt.Printf("Not found error: %v\n", err)
```

### NewOperationFailedError

```go
func NewOperationFailedError(operation string, err error) *CPEError
```

Creates an error for when an operation fails.

**Parameters:**
- `operation` - Name of the failed operation
- `err` - Underlying error

**Returns:**
- `*CPEError` - Operation failed error

**Example:**
```go
err := cpeskills.NewOperationFailedError("download", errors.New("network timeout"))
fmt.Printf("Operation error: %v\n", err)
```

## Error Checking Functions

### IsParsingError

```go
func IsParsingError(err error) bool
```

Checks if an error is a parsing error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's a parsing error

**Example:**
```go
_, err := cpeskills.ParseCpe23("invalid:format")
if cpeskills.IsParsingError(err) {
    fmt.Println("This is a parsing error")
}
```

### IsInvalidFormatError

```go
func IsInvalidFormatError(err error) bool
```

Checks if an error is an invalid format error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's an invalid format error

### IsInvalidPartError

```go
func IsInvalidPartError(err error) bool
```

Checks if an error is an invalid part error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's an invalid part error

### IsInvalidAttributeError

```go
func IsInvalidAttributeError(err error) bool
```

Checks if an error is an invalid attribute error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's an invalid attribute error

### IsNotFoundError

```go
func IsNotFoundError(err error) bool
```

Checks if an error is a not found error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's a not found error

### IsOperationFailedError

```go
func IsOperationFailedError(err error) bool
```

Checks if an error is an operation failed error.

**Parameters:**
- `err` - Error to check

**Returns:**
- `bool` - `true` if it's an operation failed error

## Storage Errors

The library also defines standard storage errors:

```go
var (
    ErrNotFound              = errors.New("record not found")
    ErrDuplicate             = errors.New("duplicate record")
    ErrInvalidData           = errors.New("invalid data")
    ErrStorageDisconnected   = errors.New("storage is disconnected")
)
```

These can be used with `errors.Is()` for checking:

**Example:**
```go
cpe, err := storage.RetrieveCPE("non-existent")
if errors.Is(err, cpeskills.ErrNotFound) {
    fmt.Println("CPE not found in storage")
}
```

## Error Handling Patterns

### Basic Error Handling

```go
// Parse CPE and handle errors
cpeObj, err := cpeskills.ParseCpe23("invalid:format")
if err != nil {
    if cpeskills.IsInvalidFormatError(err) {
        fmt.Println("Invalid CPE format provided")
    } else {
        fmt.Printf("Other parsing error: %v\n", err)
    }
    return
}
```

### Detailed Error Information

```go
// Get detailed error information
_, err := cpeskills.ParseCpe23("cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*")
if err != nil {
    if cpeErr, ok := err.(*cpeskills.CPEError); ok {
        fmt.Printf("Error type: %d\n", cpeErr.Type)
        fmt.Printf("Message: %s\n", cpeErr.Message)
        fmt.Printf("CPE string: %s\n", cpeErr.CPEString)
        
        if cpeErr.Err != nil {
            fmt.Printf("Underlying error: %v\n", cpeErr.Err)
        }
    }
}
```

### Error Chain Unwrapping

```go
// Handle error chains
err := someComplexOperation()
if err != nil {
    // Check for specific error types in the chain
    var cpeErr *cpeskills.CPEError
    if errors.As(err, &cpeErr) {
        fmt.Printf("CPE error found: %s\n", cpeErr.Message)
    }
    
    // Check for storage errors
    if errors.Is(err, cpeskills.ErrNotFound) {
        fmt.Println("Resource not found")
    }
}
```

### Comprehensive Error Handling

```go
func handleCPEOperation(cpeString string) {
    cpeObj, err := cpeskills.ParseCpe23(cpeString)
    if err != nil {
        switch {
        case cpeskills.IsInvalidFormatError(err):
            fmt.Printf("Invalid format: %s\n", cpeString)
        case cpeskills.IsInvalidPartError(err):
            fmt.Printf("Invalid part in: %s\n", cpeString)
        case cpeskills.IsParsingError(err):
            fmt.Printf("General parsing error: %v\n", err)
        default:
            fmt.Printf("Unexpected error: %v\n", err)
        }
        return
    }
    
    // Use the parsed CPE
    fmt.Printf("Successfully parsed: %s\n", cpeObj.GetURI())
}
```

## Complete Example

```go
package main

import (
    "errors"
    "fmt"
    "github.com/scagogogo/cpe-skills"
)

func main() {
    // Test different error scenarios
    fmt.Println("=== Error Handling Examples ===")
    
    // Invalid format error
    fmt.Println("\n1. Invalid Format Error:")
    _, err := cpeskills.ParseCpe23("invalid:format")
    if err != nil {
        if cpeskills.IsInvalidFormatError(err) {
            fmt.Printf("✓ Detected invalid format error: %v\n", err)
        }
    }
    
    // Invalid part error
    fmt.Println("\n2. Invalid Part Error:")
    _, err = cpeskills.ParseCpe23("cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*")
    if err != nil {
        if cpeskills.IsInvalidPartError(err) {
            fmt.Printf("✓ Detected invalid part error: %v\n", err)
        }
    }
    
    // Storage error simulation
    fmt.Println("\n3. Storage Error:")
    storage := cpeskills.NewMemoryStorage()
    storage.Initialize()
    
    _, err = storage.RetrieveCPE("non-existent-cpe")
    if err != nil {
        if errors.Is(err, cpeskills.ErrNotFound) {
            fmt.Printf("✓ Detected not found error: %v\n", err)
        }
    }
    
    // Detailed error information
    fmt.Println("\n4. Detailed Error Information:")
    _, err = cpeskills.ParseCpe23("cpe:2.3:invalid:vendor:product:1.0:*:*:*:*:*:*:*")
    if err != nil {
        if cpeErr, ok := err.(*cpeskills.CPEError); ok {
            fmt.Printf("Error Type: %d\n", cpeErr.Type)
            fmt.Printf("Message: %s\n", cpeErr.Message)
            fmt.Printf("CPE String: %s\n", cpeErr.CPEString)
        }
    }
    
    // Error type checking
    fmt.Println("\n5. Error Type Checking:")
    testCases := []struct {
        input       string
        description string
    }{
        {"invalid", "completely invalid"},
        {"cpe:2.3:x:vendor:product:1.0:*:*:*:*:*:*:*", "invalid part"},
        {"cpe:2.3:a:vendor:product", "too few components"},
    }
    
    for _, tc := range testCases {
        _, err := cpeskills.ParseCpe23(tc.input)
        if err != nil {
            fmt.Printf("Input: %s (%s)\n", tc.input, tc.description)
            
            switch {
            case cpeskills.IsInvalidFormatError(err):
                fmt.Println("  → Invalid format error")
            case cpeskills.IsInvalidPartError(err):
                fmt.Println("  → Invalid part error")
            case cpeskills.IsParsingError(err):
                fmt.Println("  → General parsing error")
            default:
                fmt.Println("  → Other error type")
            }
        }
    }
    
    // Successful operation
    fmt.Println("\n6. Successful Operation:")
    cpeObj, err := cpeskills.ParseCpe23("cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*")
    if err != nil {
        fmt.Printf("Unexpected error: %v\n", err)
    } else {
        fmt.Printf("✓ Successfully parsed: %s\n", cpeObj.GetURI())
    }
}
```
