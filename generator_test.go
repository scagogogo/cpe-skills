package cpe

import (
	"testing"
)

func TestGenerateCPE(t *testing.T) {
	cpe := GenerateCPE("a", "microsoft", "windows", "10")

	if cpe.Part.ShortName != "a" {
		t.Errorf("Part = %q, want %q", cpe.Part.ShortName, "a")
	}
	if cpe.Vendor != "microsoft" {
		t.Errorf("Vendor = %q, want %q", cpe.Vendor, "microsoft")
	}
	if cpe.ProductName != "windows" {
		t.Errorf("ProductName = %q, want %q", cpe.ProductName, "windows")
	}
	if cpe.Version != "10" {
		t.Errorf("Version = %q, want %q", cpe.Version, "10")
	}
}

func TestGenerateFromTemplate(t *testing.T) {
	template := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}

	cpe := GenerateFromTemplate(template, map[string]string{
		AttrVersion: "11",
	})

	if cpe.Vendor != "microsoft" {
		t.Errorf("Vendor = %q, want %q", cpe.Vendor, "microsoft")
	}
	if cpe.Version != "11" {
		t.Errorf("Version = %q, want %q", cpe.Version, "11")
	}
}

func TestFillDefaults(t *testing.T) {
	cpe := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
		Version:     "10",
	}

	result := FillDefaults(cpe)

	if result.Update != Update(ValueANY) {
		t.Errorf("Update = %q, want %q", result.Update, ValueANY)
	}
	if result.Edition != Edition(ValueANY) {
		t.Errorf("Edition = %q, want %q", result.Edition, ValueANY)
	}
	if result.Cpe23 == "" {
		t.Error("Cpe23 should be auto-generated")
	}
}

func TestFillDefaultsNil(t *testing.T) {
	result := FillDefaults(nil)

	if result.Part.ShortName != "a" {
		t.Errorf("Part = %q, want %q", result.Part.ShortName, "a")
	}
}

func TestMergeCPEs(t *testing.T) {
	primary := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
	}

	secondary := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "office",
		Version:     "2019",
	}

	result := MergeCPEs(primary, secondary)

	if result.Vendor != "microsoft" {
		t.Errorf("Vendor = %q, want %q", result.Vendor, "microsoft")
	}
	// Primary's non-empty value takes precedence
	if result.ProductName != "windows" {
		t.Errorf("ProductName = %q, want %q", result.ProductName, "windows")
	}
	// Secondary fills empty fields
	if result.Version != "2019" {
		t.Errorf("Version = %q, want %q", result.Version, "2019")
	}
}

func TestMergeCPEsNilPrimary(t *testing.T) {
	secondary := &CPE{
		Part:        *PartApplication,
		Vendor:      "microsoft",
		ProductName: "windows",
	}

	result := MergeCPEs(nil, secondary)

	if result.Vendor != "microsoft" {
		t.Errorf("Vendor = %q, want %q", result.Vendor, "microsoft")
	}
}

func TestFuzzyGenerateCPE(t *testing.T) {
	cpe := FuzzyGenerateCPE("A", "Microsoft Corporation", "Windows Server", "10.0")

	if cpe.Part.ShortName != "a" {
		t.Errorf("Part = %q, want %q", cpe.Part.ShortName, "a")
	}
	if cpe.Vendor != "microsoft_corporation" {
		t.Errorf("Vendor = %q, want %q", cpe.Vendor, "microsoft_corporation")
	}
}

func TestNormalizeComponentGenerator(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Windows", "windows"},
		{"Microsoft Office", "microsoft_office"},
		{"Service-Pack", "service-pack"},
		{"already_lower", "already_lower"},
	}

	for _, tt := range tests {
		if got := NormalizeComponent(tt.input); got != tt.expected {
			t.Errorf("NormalizeComponent(%q) = %q, want %q", tt.input, got, tt.expected)
		}
	}
}

func TestRandomCPE(t *testing.T) {
	cpe := RandomCPE()

	if cpe == nil {
		t.Error("RandomCPE() returned nil")
	}
	if cpe.Part.ShortName != "a" {
		t.Errorf("Part = %q, want %q", cpe.Part.ShortName, "a")
	}
}

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "CPE 2.3 format",
			input:   "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
			wantErr: false,
		},
		{
			name:    "CPE 2.2 format",
			input:   "cpe:/a:microsoft:windows:10",
			wantErr: false,
		},
		{
			name:    "invalid format",
			input:   "invalid-cpe-string",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cpe, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && cpe == nil {
				t.Error("Parse() returned nil CPE without error")
			}
		})
	}
}
