package cpe

import (
	"testing"
)

func TestParsePart(t *testing.T) {
	tests := []struct {
		input    string
		short    string
		wantErr  bool
	}{
		{"a", "a", false},
		{"h", "h", false},
		{"o", "o", false},
		{"A", "a", false},
		{"*", "*", false},
		{"x", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			p, err := ParsePart(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePart(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if !tt.wantErr && p.ShortName != tt.short {
				t.Errorf("ParsePart(%q).ShortName = %q, want %q", tt.input, p.ShortName, tt.short)
			}
		})
	}
}

func TestVendorMethods(t *testing.T) {
	tests := []struct {
		name     string
		vendor   Vendor
		isANY    bool
		isNA     bool
		isSet    bool
	}{
		{"ANY", Vendor(ValueANY), true, false, false},
		{"NA", Vendor(ValueNA), false, true, false},
		{"empty", Vendor(""), false, false, false},
		{"set", Vendor("microsoft"), false, false, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.vendor.IsANY() != tt.isANY {
				t.Errorf("IsANY() = %v, want %v", tt.vendor.IsANY(), tt.isANY)
			}
			if tt.vendor.IsNA() != tt.isNA {
				t.Errorf("IsNA() = %v, want %v", tt.vendor.IsNA(), tt.isNA)
			}
			if tt.vendor.IsSet() != tt.isSet {
				t.Errorf("IsSet() = %v, want %v", tt.vendor.IsSet(), tt.isSet)
			}
		})
	}
}

func TestProductMethods(t *testing.T) {
	p := Product("windows")
	if p.String() != "windows" {
		t.Errorf("String() = %q, want %q", p.String(), "windows")
	}
	if p.IsANY() {
		t.Error("IsANY() should be false")
	}
	if !p.IsSet() {
		t.Error("IsSet() should be true")
	}
	if p.Normalize() != "windows" {
		t.Errorf("Normalize() = %q, want %q", p.Normalize(), "windows")
	}
}

func TestVersionMethods(t *testing.T) {
	v := Version("10.0")
	if v.String() != "10.0" {
		t.Errorf("String() = %q, want %q", v.String(), "10.0")
	}
	if v.IsANY() {
		t.Error("IsANY() should be false")
	}
	if !v.IsSet() {
		t.Error("IsSet() should be true")
	}
}

func TestEditionMethods(t *testing.T) {
	e := Edition("pro")
	if !e.IsSet() {
		t.Error("IsSet() should be true")
	}
	if e.IsANY() {
		t.Error("IsANY() should be false")
	}
}

func TestLanguageMethods(t *testing.T) {
	l := Language("en")
	if !l.IsSet() {
		t.Error("IsSet() should be true")
	}
	if l.IsNA() {
		t.Error("IsNA() should be false")
	}
}

func TestUpdateMethods(t *testing.T) {
	u := Update("sp1")
	if !u.IsSet() {
		t.Error("IsSet() should be true")
	}
	if u.IsANY() {
		t.Error("IsANY() should be false")
	}
}

func TestPartMethods(t *testing.T) {
	p := *PartApplication
	if p.IsANY() {
		t.Error("IsANY() should be false for application")
	}
	if !p.IsSet() {
		t.Error("IsSet() should be true for application")
	}
	if p.Normalize() != "a" {
		t.Errorf("Normalize() = %q, want %q", p.Normalize(), "a")
	}

	anyPart := Part{ShortName: ValueANY}
	if !anyPart.IsANY() {
		t.Error("IsANY() should be true for ANY part")
	}
}

func TestVendorNormalize(t *testing.T) {
	v := Vendor("Microsoft Corporation")
	if v.Normalize() != "microsoft_corporation" {
		t.Errorf("Normalize() = %q, want %q", v.Normalize(), "microsoft_corporation")
	}
}
