package cpeskills

import (
	"encoding/json"
	"testing"
)

func TestParseCycloneDXJSON(t *testing.T) {
	// 最小的 CycloneDX JSON
	input := `{
		"bomFormat": "CycloneDX",
		"specVersion": "1.5",
		"version": 1,
		"metadata": {
			"timestamp": "2024-01-15T10:30:00Z",
			"tools": [
				{"name": "cpe-cli", "vendor": "scagogogo", "version": "1.0.0"}
			]
		},
		"components": [
			{
				"type": "library",
				"name": "lodash",
				"version": "4.17.21",
				"purl": "pkg:npm/lodash@4.17.21",
				"cpe": "cpe:2.3:a:lodash:lodash:4.17.21:*:*:*:*:*:*:*",
				"licenses": [
					{"license": {"id": "MIT"}}
				],
				"hashes": [
					{"alg": "SHA-256", "content": "abc123"}
				]
			}
		],
		"dependencies": [
			{"ref": "lodash@4.17.21", "dependsOn": []}
		]
	}`

	sbom, err := ParseCycloneDXJSON([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if sbom.Format != SBOMFormatCycloneDX {
		t.Errorf("expected format %s, got %s", SBOMFormatCycloneDX, sbom.Format)
	}
	if sbom.SpecVersion != "1.5" {
		t.Errorf("expected spec version '1.5', got %q", sbom.SpecVersion)
	}
	if sbom.ComponentCount() != 1 {
		t.Fatalf("expected 1 component, got %d", sbom.ComponentCount())
	}

	comp := sbom.Components[0]
	if comp.Name != "lodash" {
		t.Errorf("expected name 'lodash', got %q", comp.Name)
	}
	if comp.Version != "4.17.21" {
		t.Errorf("expected version '4.17.21', got %q", comp.Version)
	}
	if comp.Type != "library" {
		t.Errorf("expected type 'library', got %q", comp.Type)
	}
	if comp.PURL == nil || comp.PURL.Name != "lodash" {
		t.Error("expected PURL to be parsed")
	}
	if comp.CPE == nil || string(comp.CPE.Vendor) != "lodash" {
		t.Errorf("expected CPE to be parsed, got vendor %q", string(comp.CPE.Vendor))
	}
	if len(comp.Licenses) != 1 || comp.Licenses[0].SPDXID != "MIT" {
		t.Error("expected MIT license")
	}
	if comp.Hashes["sha-256"] != "abc123" {
		t.Errorf("expected hash 'abc123', got %q", comp.Hashes["sha-256"])
	}
	if sbom.DependencyCount() != 1 {
		t.Errorf("expected 1 dependency, got %d", sbom.DependencyCount())
	}

	// 元数据
	if len(sbom.Metadata.Tools) != 1 {
		t.Errorf("expected 1 tool, got %d", len(sbom.Metadata.Tools))
	}
	if sbom.Metadata.Tools[0].Name != "cpe-cli" {
		t.Errorf("expected tool name 'cpe-cli', got %q", sbom.Metadata.Tools[0].Name)
	}
}

func TestParseCycloneDXJSON_Invalid(t *testing.T) {
	_, err := ParseCycloneDXJSON([]byte("not json"))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestToCycloneDXJSON(t *testing.T) {
	sbom := NewSBOM(SBOMFormatCycloneDX, "test")
	sbom.SpecVersion = "1.5"
	sbom.SerialNumber = "urn:uuid:test-123"

	comp := NewSBOMComponent("lodash", "4.17.21")
	comp.Type = "library"
	comp.SetPURL(NewPURL("npm", "", "lodash", "4.17.21"))
	cpe, _ := Parse("cpe:2.3:a:lodash:lodash:4.17.21:*:*:*:*:*:*:*")
	comp.SetCPE(cpe)
	comp.AddHash("SHA-256", "abc123")
	comp.Supplier = "npm"
	sbom.AddComponent(comp)

	sbom.AddDependency("lodash@4.17.21", []string{})

	data, err := sbom.ToCycloneDXJSON()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// 验证输出的 JSON 结构
	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("invalid JSON output: %v", err)
	}

	if result["bomFormat"] != "CycloneDX" {
		t.Errorf("expected bomFormat 'CycloneDX', got %v", result["bomFormat"])
	}

	components, ok := result["components"].([]interface{})
	if !ok || len(components) != 1 {
		t.Fatalf("expected 1 component in output")
	}

	compMap := components[0].(map[string]interface{})
	if compMap["name"] != "lodash" {
		t.Errorf("expected name 'lodash', got %v", compMap["name"])
	}
}

func TestParseCycloneDXJSON_RoundTrip(t *testing.T) {
	original := `{
		"bomFormat": "CycloneDX",
		"specVersion": "1.5",
		"version": 1,
		"components": [
			{
				"type": "library",
				"name": "express",
				"version": "4.17.1",
				"purl": "pkg:npm/express@4.17.1"
			}
		]
	}`

	sbom, err := ParseCycloneDXJSON([]byte(original))
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}

	exported, err := sbom.ToCycloneDXJSON()
	if err != nil {
		t.Fatalf("export error: %v", err)
	}

	// Re-parse exported data
	sbom2, err := ParseCycloneDXJSON(exported)
	if err != nil {
		t.Fatalf("re-parse error: %v", err)
	}

	if sbom2.ComponentCount() != 1 {
		t.Errorf("expected 1 component after round-trip, got %d", sbom2.ComponentCount())
	}
	if sbom2.Components[0].Name != "express" {
		t.Errorf("expected name 'express', got %q", sbom2.Components[0].Name)
	}
}

func TestParseCycloneDXJSON_Empty(t *testing.T) {
	input := `{
		"bomFormat": "CycloneDX",
		"specVersion": "1.4",
		"version": 1
	}`

	sbom, err := ParseCycloneDXJSON([]byte(input))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if sbom.ComponentCount() != 0 {
		t.Errorf("expected 0 components, got %d", sbom.ComponentCount())
	}
}
