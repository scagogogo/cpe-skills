package cpe

import (
	"testing"
)

func TestRelationString(t *testing.T) {
	tests := []struct {
		relation Relation
		expected string
	}{
		{RelationDisjoint, "disjoint"},
		{RelationSubset, "subset"},
		{RelationSuperset, "superset"},
		{RelationEqual, "equal"},
		{RelationOverlap, "overlap"},
		{RelationUnknown, "unknown"},
	}

	for _, tt := range tests {
		if got := tt.relation.String(); got != tt.expected {
			t.Errorf("Relation(%d).String() = %q, want %q", tt.relation, got, tt.expected)
		}
	}
}

func TestCompareAttributes(t *testing.T) {
	tests := []struct {
		name     string
		source   string
		target   string
		expected int
	}{
		{"both ANY", ValueANY, ValueANY, 0},
		{"source ANY", ValueANY, "windows", 1},
		{"target ANY", "windows", ValueANY, -1},
		{"both NA", ValueNA, ValueNA, 0},
		{"source NA", ValueNA, "windows", -2},
		{"target NA", "windows", ValueNA, -2},
		{"equal values", "windows", "windows", 0},
		{"different values", "windows", "linux", -2},
		{"both empty", "", "", 0},
		{"source empty", "", "windows", 1},
		{"target empty", "windows", "", -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareAttributes(tt.source, tt.target); got != tt.expected {
				t.Errorf("CompareAttributes(%q, %q) = %d, want %d", tt.source, tt.target, got, tt.expected)
			}
		})
	}
}

func TestCompareWFNs(t *testing.T) {
	source := &WFN{
		Part:    "a",
		Vendor:  "microsoft",
		Product: "windows",
		Version: ValueANY, // ANY = superset for this attribute
	}
	target := &WFN{
		Part:    "a",
		Vendor:  "microsoft",
		Product: "windows",
		Version: "10",
	}

	comparisons := CompareWFNs(source, target)

	if comparisons[AttrPart] != 0 {
		t.Errorf("Part comparison = %d, want 0", comparisons[AttrPart])
	}
	if comparisons[AttrVersion] != 1 {
		t.Errorf("Version comparison = %d, want 1 (superset)", comparisons[AttrVersion])
	}
}

func TestCompareWFNRelation(t *testing.T) {
	tests := []struct {
		name        string
		comparisons map[string]int
		expected    Relation
	}{
		{
			name:        "all equal",
			comparisons: map[string]int{AttrPart: 0, AttrVendor: 0},
			expected:    RelationEqual,
		},
		{
			name:        "superset",
			comparisons: map[string]int{AttrPart: 0, AttrVersion: 1},
			expected:    RelationSuperset,
		},
		{
			name:        "subset",
			comparisons: map[string]int{AttrPart: 0, AttrVersion: -1},
			expected:    RelationSubset,
		},
		{
			name:        "disjoint",
			comparisons: map[string]int{AttrPart: 0, AttrVendor: -2},
			expected:    RelationDisjoint,
		},
		{
			name:        "overlap",
			comparisons: map[string]int{AttrPart: 1, AttrVersion: -1},
			expected:    RelationOverlap,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareWFNRelation(tt.comparisons); got != tt.expected {
				t.Errorf("CompareWFNRelation() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestCPEDisjoint(t *testing.T) {
	a := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}
	b := &CPE{
		Part:        *PartApplication,
		Vendor:      "adobe",
		ProductName: "reader",
		Version:     "10",
	}

	if !CPEDisjoint(a, b) {
		t.Error("Expected CPEs to be disjoint")
	}

	c := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}
	if CPEDisjoint(a, c) {
		t.Error("Expected equal CPEs to not be disjoint")
	}

	if !CPEDisjoint(nil, a) {
		t.Error("Expected nil CPE to be disjoint")
	}
}

func TestCPEEqual(t *testing.T) {
	a := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}
	b := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}

	if !CPEEqual(a, b) {
		t.Error("Expected CPEs to be equal")
	}

	c := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "11",
	}
	if CPEEqual(a, c) {
		t.Error("Expected different version CPEs to not be equal")
	}

	if CPEEqual(nil, a) {
		t.Error("Expected nil CPE to not be equal")
	}
}

func TestCPESubset(t *testing.T) {
	// a with specific version is subset of b with ANY version
	a := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}
	b := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     ValueANY,
	}

	if !CPESubset(a, b) {
		t.Error("Expected specific version to be subset of ANY version")
	}
}

func TestCPESuperset(t *testing.T) {
	// a with ANY version is superset of b with specific version
	a := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     ValueANY,
	}
	b := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}

	if !CPESuperset(a, b) {
		t.Error("Expected ANY version to be superset of specific version")
	}
}

func TestWildcardMatch(t *testing.T) {
	tests := []struct {
		source   string
		target   string
		expected bool
	}{
		{"windows", "windows", true},
		{"win*", "windows", true},
		{"win*", "win", true},
		{"win*dows", "windows", true},
		{"win?ows", "windows", true},
		{"win?ows", "winsows", true},
		{"windows", "linux", false},
		{"*", "anything", true},
		{"*", "", true},
	}

	for _, tt := range tests {
		if got := wildcardMatch(tt.source, tt.target); got != tt.expected {
			t.Errorf("wildcardMatch(%q, %q) = %v, want %v", tt.source, tt.target, got, tt.expected)
		}
	}
}
