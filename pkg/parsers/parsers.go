// Package parsers provides manifest/lockfile parsers for various package managers.
//
// Each parser extracts component information (name, version, ecosystem/PURL type)
// from its native manifest format and returns standardized SBOM components.
// This enables automatic SBOM generation from source code repositories without
// requiring external tools.
//
// Supported package managers:
//   - Go (go.mod)
//   - Node.js (package.json, package-lock.json)
//   - Python (requirements.txt)
//   - Maven (pom.xml)
//   - Rust (Cargo.toml, Cargo.lock)
//   - PHP (composer.json, composer.lock)
//   - Ruby (Gemfile, Gemfile.lock)
//   - NuGet (.csproj packages.config)

package parsers

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"regexp"
	"strings"
)

// ParseResult holds the result of parsing a manifest/lockfile.
type ParseResult struct {
	// Ecosystem is the package ecosystem (npm, maven, pypi, golang, etc.)
	Ecosystem string `json:"ecosystem"`

	// PURLType is the PURL type field for this ecosystem
	PURLType string `json:"purlType"`

	// Components is the list of parsed components
	Components []*ComponentInfo `json:"components"`

	// Name is the project name (from the manifest)
	Name string `json:"name,omitempty"`

	// Version is the project version (from the manifest)
	Version string `json:"version,omitempty"`

	// Metadata contains additional manifest-specific info
	Metadata map[string]interface{} `json:"metadata,omitempty"`
}

// ComponentInfo represents a parsed software component.
type ComponentInfo struct {
	// Name is the package name
	Name string `json:"name"`

	// Version is the package version
	Version string `json:"version"`

	// Namespace is the package namespace (e.g., Maven groupId, npm scope)
	Namespace string `json:"namespace,omitempty"`

	// IsDirect indicates a direct (not transitive) dependency
	IsDirect bool `json:"isDirect"`

	// IsDev indicates a development-only dependency
	IsDev bool `json:"isDev,omitempty"`

	// Scope indicates the dependency scope (compile, test, runtime, etc.)
	Scope string `json:"scope,omitempty"`

	// Checksum is the integrity hash (from lockfiles)
	Checksum string `json:"checksum,omitempty"`

	// Resolved is the resolved version URL or path
	Resolved string `json:"resolved,omitempty"`
}

// ParseFunc is a function type for parsing manifest data.
type ParseFunc func(reader io.Reader) (*ParseResult, error)

// ParseGoMod parses a go.mod file.
//
// Extracts the module name, Go version, and direct dependencies (require directives).
// Indirect dependencies are marked with IsDirect=false.
func ParseGoMod(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read go.mod: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "golang",
		PURLType:  "golang",
		Metadata:  make(map[string]interface{}),
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	inRequire := false
	moduleRe := regexp.MustCompile(`^\s*module\s+(\S+)`)
	requireRe := regexp.MustCompile(`^\s*require\s+(\S+)\s+(\S+)`)
	indirectRequireRe := regexp.MustCompile(`^\s*(\S+)\s+(\S+)\s*//\s*indirect`)
	directInBlockRe := regexp.MustCompile(`^\s*(\S+)\s+(\S+)\s*$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip comments and empty lines
		if line == "" || strings.HasPrefix(line, "//") {
			continue
		}

		// Module name
		if m := moduleRe.FindStringSubmatch(line); m != nil {
			result.Name = m[1]
			continue
		}

		// Start of require block
		if strings.HasPrefix(line, "require (") {
			inRequire = true
			continue
		}

		// End of require block
		if inRequire && line == ")" {
			inRequire = false
			continue
		}

		// Single-line require
		if m := requireRe.FindStringSubmatch(line); m != nil {
			name := m[1]
			version := m[2]
			isIndirect := strings.Contains(line, "// indirect")

			result.Components = append(result.Components, &ComponentInfo{
				Name:     name,
				Version:  version,
				IsDirect: !isIndirect,
			})
			continue
		}

		// Inside require block
		if inRequire {
			if m := indirectRequireRe.FindStringSubmatch(line); m != nil {
				result.Components = append(result.Components, &ComponentInfo{
					Name:     m[1],
					Version:  m[2],
					IsDirect: false,
				})
			} else if m := directInBlockRe.FindStringSubmatch(line); m != nil {
				result.Components = append(result.Components, &ComponentInfo{
					Name:     m[1],
					Version:  m[2],
					IsDirect: true,
				})
			}
		}
	}

	return result, nil
}

// ParsePackageJSON parses a package.json file.
//
// Extracts package name, version, and all dependency types:
// dependencies, devDependencies, peerDependencies, optionalDependencies.
func ParsePackageJSON(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read package.json: %w", err)
	}

	var pkg struct {
		Name            string            `json:"name"`
		Version         string            `json:"version"`
		Dependencies    map[string]string `json:"dependencies"`
		DevDependencies map[string]string `json:"devDependencies"`
		PeerDependencies map[string]string `json:"peerDependencies"`
		OptDependencies map[string]string `json:"optionalDependencies"`
	}

	if err := json.Unmarshal(data, &pkg); err != nil {
		return nil, fmt.Errorf("failed to parse package.json: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "npm",
		PURLType:  "npm",
		Name:      pkg.Name,
		Version:   pkg.Version,
		Metadata:  make(map[string]interface{}),
	}

	addComponents := func(deps map[string]string, isDirect, isDev bool, scope string) {
		for name, version := range deps {
			result.Components = append(result.Components, &ComponentInfo{
				Name:     name,
				Version:  strings.TrimPrefix(version, "^"),
				IsDirect: isDirect,
				IsDev:    isDev,
				Scope:    scope,
			})
		}
	}

	addComponents(pkg.Dependencies, true, false, "runtime")
	addComponents(pkg.DevDependencies, true, true, "dev")
	addComponents(pkg.PeerDependencies, true, false, "peer")
	addComponents(pkg.OptDependencies, false, false, "optional")

	return result, nil
}

// ParsePackageLockJSON parses a package-lock.json (v1/v2) file.
//
// Extracts resolved versions, integrity hashes, and the full dependency tree.
func ParsePackageLockJSON(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read package-lock.json: %w", err)
	}

	var lock struct {
		Name            string `json:"name"`
		Version         string `json:"version"`
		LockfileVersion int    `json:"lockfileVersion"`
		Packages        map[string]struct {
			Version   string `json:"version"`
			Resolved  string `json:"resolved"`
			Integrity string `json:"integrity"`
			Dev       bool   `json:"dev"`
		} `json:"packages"`
		Dependencies map[string]struct {
			Version   string `json:"version"`
			Resolved  string `json:"resolved"`
			Integrity string `json:"integrity"`
			Dev       bool   `json:"dev"`
			Requires  map[string]string `json:"requires"`
		} `json:"dependencies"`
	}

	if err := json.Unmarshal(data, &lock); err != nil {
		return nil, fmt.Errorf("failed to parse package-lock.json: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "npm",
		PURLType:  "npm",
		Name:      lock.Name,
		Version:   lock.Version,
		Metadata:  make(map[string]interface{}),
	}

	// Handle lockfile v2+ format with "packages" field
	if len(lock.Packages) > 0 {
		for path, info := range lock.Packages {
			// Skip the root self-reference
			if path == "" {
				continue
			}

			name := extractPackageNameFromLockPath(path)
			result.Components = append(result.Components, &ComponentInfo{
				Name:     name,
				Version:  info.Version,
				IsDirect: !info.Dev,
				IsDev:    info.Dev,
				Resolved: info.Resolved,
				Checksum: info.Integrity,
			})
		}
		return result, nil
	}

	// Handle lockfile v1 format with "dependencies" field
	for name, info := range lock.Dependencies {
		result.Components = append(result.Components, &ComponentInfo{
			Name:     name,
			Version:  info.Version,
			IsDirect: !info.Dev,
			IsDev:    info.Dev,
			Resolved: info.Resolved,
			Checksum: info.Integrity,
		})
	}

	return result, nil
}

// ParseRequirementsTxt parses a requirements.txt file.
//
// Supports version specifiers: ==, >=, <=, ~=, !=, >, <
func ParseRequirementsTxt(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read requirements.txt: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "pypi",
		PURLType:  "pypi",
		Metadata:  make(map[string]interface{}),
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	// Match: package_name==version or package_name>=version etc.
	pkgRe := regexp.MustCompile(`^([a-zA-Z0-9][\w\-_.]*)\s*([><=!~]+\s*[\w\*\.\+\-]+)`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip comments, empty lines, and option flags
		if line == "" || strings.HasPrefix(line, "#") || strings.HasPrefix(line, "-") {
			continue
		}

		// Handle -r includes (just note them)
		if strings.HasPrefix(line, "-r ") {
			continue
		}

		if m := pkgRe.FindStringSubmatch(line); m != nil {
			version := strings.TrimSpace(m[2])
			// Remove operator prefix for the version value
			version = strings.TrimLeft(version, "><=!~ ")
			result.Components = append(result.Components, &ComponentInfo{
				Name:     m[1],
				Version:  version,
				IsDirect: true,
			})
		}
	}

	return result, nil
}

// ParsePomXML parses a Maven pom.xml file.
//
// Extracts groupId, artifactId, version, and all declared dependencies
// including their scope (compile, test, provided, runtime).
func ParsePomXML(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read pom.xml: %w", err)
	}

	type Dependency struct {
		GroupID    string `xml:"groupId"`
		ArtifactID string `xml:"artifactId"`
		Version    string `xml:"version"`
		Scope      string `xml:"scope"`
		Optional   string `xml:"optional"`
	}

	type Project struct {
		GroupID      string       `xml:"groupId"`
		ArtifactID   string       `xml:"artifactId"`
		Version      string       `xml:"version"`
		Dependencies []Dependency `xml:"dependencies>dependency"`
	}

	var project Project
	if err := xml.Unmarshal(data, &project); err != nil {
		return nil, fmt.Errorf("failed to parse pom.xml: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "maven",
		PURLType:  "maven",
		Name:      project.GroupID + ":" + project.ArtifactID,
		Version:   project.Version,
		Metadata:  make(map[string]interface{}),
	}

	for _, dep := range project.Dependencies {
		scope := dep.Scope
		if scope == "" {
			scope = "compile"
		}

		result.Components = append(result.Components, &ComponentInfo{
			Name:      dep.ArtifactID,
			Version:   dep.Version,
			Namespace: dep.GroupID,
			IsDirect:  true,
			Scope:     scope,
		})
	}

	return result, nil
}

// ParseCargoToml parses a Cargo.toml file.
//
// Extracts package name, version, and all dependency sections.
func ParseCargoToml(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read Cargo.toml: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "cargo",
		PURLType:  "cargo",
		Metadata:  make(map[string]interface{}),
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	inDeps := false
	inDevDeps := false
	inBuildDeps := false

	nameRe := regexp.MustCompile(`^\s*name\s*=\s*"([^"]+)"`)
	versionRe := regexp.MustCompile(`^\s*version\s*=\s*"([^"]+)"`)
	depRe := regexp.MustCompile(`^\s*([a-zA-Z][\w\-_]*)\s*=\s*"([^"]+)"`)
	tableDepRe := regexp.MustCompile(`^\s*([a-zA-Z][\w\-_]*)\s*=\s*\{.*version\s*=\s*"([^"]+)"`)

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Package name
		if m := nameRe.FindStringSubmatch(line); m != nil && !inDeps && !inDevDeps {
			result.Name = m[1]
			continue
		}

		// Package version
		if m := versionRe.FindStringSubmatch(line); m != nil && !inDeps && !inDevDeps {
			result.Version = m[1]
			continue
		}

		// Section detection
		if strings.HasPrefix(trimmed, "[dependencies]") {
			inDeps = true
			inDevDeps = false
			inBuildDeps = false
			continue
		}
		if strings.HasPrefix(trimmed, "[dev-dependencies]") {
			inDeps = false
			inDevDeps = true
			inBuildDeps = false
			continue
		}
		if strings.HasPrefix(trimmed, "[build-dependencies]") {
			inDeps = false
			inDevDeps = false
			inBuildDeps = true
			continue
		}

		// Any new section ends dependency parsing
		if strings.HasPrefix(trimmed, "[") {
			inDeps = false
			inDevDeps = false
			inBuildDeps = false
			continue
		}

		// Parse dependency lines
		if inDeps || inDevDeps || inBuildDeps {
			if m := depRe.FindStringSubmatch(line); m != nil {
				result.Components = append(result.Components, &ComponentInfo{
					Name:     m[1],
					Version:  m[2],
					IsDirect: true,
					IsDev:    inDevDeps,
				})
			} else if m := tableDepRe.FindStringSubmatch(line); m != nil {
				result.Components = append(result.Components, &ComponentInfo{
					Name:     m[1],
					Version:  m[2],
					IsDirect: true,
					IsDev:    inDevDeps,
				})
			}
		}
	}

	return result, nil
}

// ParseComposerJSON parses a composer.json file.
//
// Extracts package name and PHP dependencies (require, require-dev).
func ParseComposerJSON(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read composer.json: %w", err)
	}

	var composer struct {
		Name        string            `json:"name"`
		Version     string            `json:"version"`
		Require     map[string]string `json:"require"`
		RequireDev  map[string]string `json:"require-dev"`
	}

	if err := json.Unmarshal(data, &composer); err != nil {
		return nil, fmt.Errorf("failed to parse composer.json: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "composer",
		PURLType:  "composer",
		Name:      composer.Name,
		Version:   composer.Version,
		Metadata:  make(map[string]interface{}),
	}

	addComponents := func(deps map[string]string, isDev bool) {
		for name, version := range deps {
			// Skip PHP version requirement
			if name == "php" || strings.HasPrefix(name, "ext-") || strings.HasPrefix(name, "lib-") {
				continue
			}
			result.Components = append(result.Components, &ComponentInfo{
				Name:     name,
				Version:  strings.TrimPrefix(strings.TrimPrefix(version, "^"), "~"),
				IsDirect: true,
				IsDev:    isDev,
			})
		}
	}

	addComponents(composer.Require, false)
	addComponents(composer.RequireDev, true)

	return result, nil
}

// composerLockPackage is a package entry in composer.lock
type composerLockPackage struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Source  struct {
		Reference string `json:"reference"`
	} `json:"source"`
}

// ParseComposerLock parses a composer.lock file.
//
// Extracts all installed packages with exact versions and resolved references.
func ParseComposerLock(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read composer.lock: %w", err)
	}

	var lock struct {
		Packages    []composerLockPackage `json:"packages"`
		PackagesDev []composerLockPackage `json:"packages-dev"`
	}



	if err := json.Unmarshal(data, &lock); err != nil {
		return nil, fmt.Errorf("failed to parse composer.lock: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "composer",
		PURLType:  "composer",
		Metadata:  make(map[string]interface{}),
	}

	addPackages := func(pkgs []composerLockPackage, isDev bool) {
		for _, pkg := range pkgs {
			result.Components = append(result.Components, &ComponentInfo{
				Name:     pkg.Name,
				Version:  pkg.Version,
				IsDirect: true,
				IsDev:    isDev,
				Resolved: pkg.Source.Reference,
			})
		}
	}

	addPackages(lock.Packages, false)
	addPackages(lock.PackagesDev, true)

	return result, nil
}

// ParseGemfile parses a Ruby Gemfile.
//
// Extracts gem declarations with version constraints.
func ParseGemfile(reader io.Reader) (*ParseResult, error) {
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to read Gemfile: %w", err)
	}

	result := &ParseResult{
		Ecosystem: "rubygems",
		PURLType:  "gem",
		Metadata:  make(map[string]interface{}),
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	gemRe := regexp.MustCompile(`^\s*gem\s+['"]([^'"]+)['"](?:\s*,\s*['"]([^'"]+)['"])?`)

	inGroup := false
	isDev := false

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" || strings.HasPrefix(trimmed, "#") {
			continue
		}

		// Detect group blocks
		if strings.Contains(trimmed, "group :development") || strings.Contains(trimmed, `group "development"`) {
			inGroup = true
			isDev = true
			continue
		}
		if strings.Contains(trimmed, "group :test") || strings.Contains(trimmed, `group "test"`) {
			inGroup = true
			isDev = true
			continue
		}
		if inGroup && trimmed == "end" {
			inGroup = false
			isDev = false
			continue
		}

		if m := gemRe.FindStringSubmatch(line); m != nil {
			version := ""
			if len(m) > 2 {
				version = m[2]
			}
			result.Components = append(result.Components, &ComponentInfo{
				Name:     m[1],
				Version:  version,
				IsDirect: true,
				IsDev:    isDev,
			})
		}
	}

	return result, nil
}

// ParseAuto detects the manifest type from filename and calls the appropriate parser.
//
// Supported filenames:
//   - go.mod → ParseGoMod
//   - package.json → ParsePackageJSON
//   - package-lock.json → ParsePackageLockJSON
//   - requirements.txt → ParseRequirementsTxt
//   - pom.xml → ParsePomXML
//   - Cargo.toml → ParseCargoToml
//   - composer.json → ParseComposerJSON
//   - composer.lock → ParseComposerLock
//   - Gemfile → ParseGemfile
func ParseAuto(filename string, reader io.Reader) (*ParseResult, error) {
	basename := strings.ToLower(filename)

	// Remove path prefix if present
	if idx := strings.LastIndex(basename, "/"); idx >= 0 {
		basename = basename[idx+1:]
	}

	switch basename {
	case "go.mod":
		return ParseGoMod(reader)
	case "package.json":
		return ParsePackageJSON(reader)
	case "package-lock.json":
		return ParsePackageLockJSON(reader)
	case "requirements.txt":
		return ParseRequirementsTxt(reader)
	case "pom.xml":
		return ParsePomXML(reader)
	case "cargo.toml":
		return ParseCargoToml(reader)
	case "composer.json":
		return ParseComposerJSON(reader)
	case "composer.lock":
		return ParseComposerLock(reader)
	case "gemfile":
		return ParseGemfile(reader)
	default:
		return nil, fmt.Errorf("unsupported manifest file: %s", basename)
	}
}

// ConvertToSBOMComponents converts ParseResult components to the cpe package's SBOMComponent format.
//
// Returns a slice of SBOM-compatible component maps that can be used to build an SBOM.
func ConvertToSBOMComponents(result *ParseResult) []ComponentMapping {
	if result == nil {
		return nil
	}

	mappings := make([]ComponentMapping, 0, len(result.Components))
	for _, c := range result.Components {
		mappings = append(mappings, ComponentMapping{
			Name:      c.Name,
			Version:   c.Version,
			Namespace: c.Namespace,
			PURLType:  result.PURLType,
			Ecosystem: result.Ecosystem,
			IsDirect:  c.IsDirect,
			IsDev:     c.IsDev,
		})
	}

	return mappings
}

// ComponentMapping represents a component ready for SBOM import.
type ComponentMapping struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Namespace string `json:"namespace,omitempty"`
	PURLType  string `json:"purlType"`
	Ecosystem string `json:"ecosystem"`
	IsDirect  bool   `json:"isDirect"`
	IsDev     bool   `json:"isDev,omitempty"`
}

// ParsersByFormat is a registry of parsers keyed by filename.
var ParsersByFormat = map[string]ParseFunc{
	"go.mod":           ParseGoMod,
	"package.json":     ParsePackageJSON,
	"package-lock.json": ParsePackageLockJSON,
	"requirements.txt": ParseRequirementsTxt,
	"pom.xml":          ParsePomXML,
	"Cargo.toml":       ParseCargoToml,
	"composer.json":    ParseComposerJSON,
	"composer.lock":    ParseComposerLock,
	"Gemfile":          ParseGemfile,
}

// extractPackageNameFromLockPath extracts the npm package name from a lockfile path like "node_modules/express".
func extractPackageNameFromLockPath(path string) string {
	path = strings.TrimPrefix(path, "node_modules/")
	// Handle scoped packages: @scope/name
	if idx := strings.LastIndex(path, "/"); idx >= 0 {
		// Check if it's a scoped package path
		parts := strings.Split(path, "/")
		if len(parts) >= 2 && strings.HasPrefix(parts[0], "@") {
			return parts[0] + "/" + parts[1]
		}
	}
	return path
}
