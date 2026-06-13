package cpe

import (
	"testing"
)

func TestNewWFN(t *testing.T) {
	wfn := NewWFN()
	if wfn == nil {
		t.Fatal("NewWFN() returned nil")
	}

	// All attributes should default to ANY via Get
	if wfn.Get(AttrPart) != ValueANY {
		t.Errorf("NewWFN().Get(AttrPart) = %q, want %q", wfn.Get(AttrPart), ValueANY)
	}
	if wfn.Get(AttrVendor) != ValueANY {
		t.Errorf("NewWFN().Get(AttrVendor) = %q, want %q", wfn.Get(AttrVendor), ValueANY)
	}
}

func TestWFNGet(t *testing.T) {
	wfn := &WFN{
		Part:    "a",
		Vendor:  "microsoft",
		Product: "windows",
		Version: "", // empty should return ANY
	}

	tests := []struct {
		attr     string
		expected string
	}{
		{AttrPart, "a"},
		{AttrVendor, "microsoft"},
		{AttrProduct, "windows"},
		{AttrVersion, ValueANY}, // empty defaults to ANY
		{AttrUpdate, ValueANY},  // not set, defaults to ANY
		{"invalid_attr", ValueANY},
	}

	for _, tt := range tests {
		t.Run(tt.attr, func(t *testing.T) {
			if got := wfn.Get(tt.attr); got != tt.expected {
				t.Errorf("WFN.Get(%q) = %q, want %q", tt.attr, got, tt.expected)
			}
		})
	}
}

func TestWFNSet(t *testing.T) {
	wfn := NewWFN()

	wfn.Set(AttrPart, "a")
	wfn.Set(AttrVendor, "microsoft")
	wfn.Set(AttrProduct, "windows")
	wfn.Set(AttrVersion, "10")

	if wfn.Part != "a" {
		t.Errorf("Part = %q, want %q", wfn.Part, "a")
	}
	if wfn.Vendor != "microsoft" {
		t.Errorf("Vendor = %q, want %q", wfn.Vendor, "microsoft")
	}
	if wfn.Product != "windows" {
		t.Errorf("Product = %q, want %q", wfn.Product, "windows")
	}
	if wfn.Version != "10" {
		t.Errorf("Version = %q, want %q", wfn.Version, "10")
	}

	// Setting unknown attribute should be no-op
	wfn.Set("unknown_attr", "value")
	// Should not panic
}

func TestWFNString(t *testing.T) {
	tests := []struct {
		name     string
		wfn      *WFN
		expected string
	}{
		{
			name:     "empty WFN",
			wfn:      NewWFN(),
			expected: "wfn:[]",
		},
		{
			name: "partial WFN",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: "windows",
			},
			expected: `wfn:[part="a",vendor="microsoft",product="windows"]`,
		},
		{
			name: "full WFN",
			wfn: &WFN{
				Part:            "a",
				Vendor:          "microsoft",
				Product:         "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			expected: `wfn:[part="a",vendor="microsoft",product="windows",version="10",update="sp1",edition="pro",language="en",sw_edition="enterprise",target_sw="linux",target_hw="x86",other="custom"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wfn.WFNString(); got != tt.expected {
				t.Errorf("WFNString() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestIsIdentifierName(t *testing.T) {
	tests := []struct {
		name     string
		wfn      *WFN
		expected bool
	}{
		{
			name: "valid identifier",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: "windows",
				Version: "10",
			},
			expected: true,
		},
		{
			name: "ANY part - not identifier",
			wfn: &WFN{
				Part:    ValueANY,
				Vendor:  "microsoft",
				Product: "windows",
			},
			expected: false,
		},
		{
			name: "ANY vendor - not identifier",
			wfn: &WFN{
				Part:    "a",
				Vendor:  ValueANY,
				Product: "windows",
			},
			expected: false,
		},
		{
			name: "ANY product - not identifier",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: ValueANY,
			},
			expected: false,
		},
		{
			name: "NA part - not identifier",
			wfn: &WFN{
				Part:    ValueNA,
				Vendor:  "microsoft",
				Product: "windows",
			},
			expected: false,
		},
		{
			name: "wildcard in version - not identifier",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: "windows",
				Version: "1*",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.wfn.IsIdentifierName(); got != tt.expected {
				t.Errorf("IsIdentifierName() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestDefaultToANY(t *testing.T) {
	wfn := NewWFN()

	if wfn.defaultToANY("") != ValueANY {
		t.Errorf("defaultToANY('') = %q, want %q", wfn.defaultToANY(""), ValueANY)
	}
	if wfn.defaultToANY("windows") != "windows" {
		t.Errorf("defaultToANY('windows') = %q, want %q", wfn.defaultToANY("windows"), "windows")
	}
}

func TestWFNConstants(t *testing.T) {
	if ValueANY != "*" {
		t.Errorf("ValueANY = %q, want %q", ValueANY, "*")
	}
	if ValueNA != "-" {
		t.Errorf("ValueNA = %q, want %q", ValueNA, "-")
	}

	if AttrPart != "part" {
		t.Errorf("AttrPart = %q, want %q", AttrPart, "part")
	}
	if AttrVendor != "vendor" {
		t.Errorf("AttrVendor = %q, want %q", AttrVendor, "vendor")
	}

	if PartApplicationShort != "a" {
		t.Errorf("PartApplicationShort = %q, want %q", PartApplicationShort, "a")
	}
	if PartOSShort != "o" {
		t.Errorf("PartOSShort = %q, want %q", PartOSShort, "o")
	}
	if PartHardwareShort != "h" {
		t.Errorf("PartHardwareShort = %q, want %q", PartHardwareShort, "h")
	}
}

func TestValidPartValues(t *testing.T) {
	if !ValidPartValues["a"] {
		t.Error("ValidPartValues should contain 'a'")
	}
	if !ValidPartValues["o"] {
		t.Error("ValidPartValues should contain 'o'")
	}
	if !ValidPartValues["h"] {
		t.Error("ValidPartValues should contain 'h'")
	}
	if !ValidPartValues[ValueANY] {
		t.Error("ValidPartValues should contain '*'")
	}
	if ValidPartValues["x"] {
		t.Error("ValidPartValues should not contain 'x'")
	}
}

func TestAllAttributes(t *testing.T) {
	expected := []string{
		AttrPart, AttrVendor, AttrProduct, AttrVersion, AttrUpdate,
		AttrEdition, AttrLanguage, AttrSoftwareEdition, AttrTargetSoftware,
		AttrTargetHardware, AttrOther,
	}

	if len(allAttributes) != len(expected) {
		t.Errorf("allAttributes length = %d, want %d", len(allAttributes), len(expected))
	}

	for i, attr := range allAttributes {
		if attr != expected[i] {
			t.Errorf("allAttributes[%d] = %q, want %q", i, attr, expected[i])
		}
	}
}
