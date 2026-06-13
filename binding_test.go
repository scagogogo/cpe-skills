package cpe

import (
	"testing"
)

func TestBindToFS(t *testing.T) {
	tests := []struct {
		name     string
		wfn      *WFN
		expected string
	}{
		{
			name:     "nil WFN",
			wfn:      nil,
			expected: "",
		},
		{
			name: "simple WFN",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: "windows",
				Version: "10",
			},
			expected: "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
		},
		{
			name: "WFN with special chars",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "example.com",
				Product: "product:name",
				Version: "1.0",
			},
			expected: "cpe:2.3:a:example\\.com:product%3aname:1\\.0:*:*:*:*:*:*:*",
		},
		{
			name: "full WFN",
			wfn: &WFN{
				Part:            "a",
				Vendor:          "apache",
				Product:         "tomcat",
				Version:         "8.5.0",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			expected: "cpe:2.3:a:apache:tomcat:8\\.5\\.0:sp1:pro:en:enterprise:linux:x86:custom",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BindToFS(tt.wfn); got != tt.expected {
				t.Errorf("BindToFS() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestUnbindFS(t *testing.T) {
	tests := []struct {
		name    string
		fs      string
		wantErr bool
		part    string
		vendor  string
		product string
		version string
	}{
		{
			name:    "valid FS",
			fs:      "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
			wantErr: false,
			part:    "a",
			vendor:  "microsoft",
			product: "windows",
			version: "10",
		},
		{
			name:    "invalid prefix",
			fs:      "cpe:2.2:a:microsoft:windows:10:*:*:*:*:*:*:*",
			wantErr: true,
		},
		{
			name:    "wrong part count",
			fs:      "cpe:2.3:a:microsoft:windows",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wfn, err := UnbindFS(tt.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnbindFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if wfn.Part != tt.part {
					t.Errorf("UnbindFS().Part = %v, want %v", wfn.Part, tt.part)
				}
				if wfn.Vendor != tt.vendor {
					t.Errorf("UnbindFS().Vendor = %v, want %v", wfn.Vendor, tt.vendor)
				}
				if wfn.Product != tt.product {
					t.Errorf("UnbindFS().Product = %v, want %v", wfn.Product, tt.product)
				}
				if wfn.Version != tt.version {
					t.Errorf("UnbindFS().Version = %v, want %v", wfn.Version, tt.version)
				}
			}
		})
	}
}

func TestBindToURI(t *testing.T) {
	tests := []struct {
		name     string
		wfn      *WFN
		expected string
	}{
		{
			name:     "nil WFN",
			wfn:      nil,
			expected: "",
		},
		{
			name: "simple WFN",
			wfn: &WFN{
				Part:    "a",
				Vendor:  "microsoft",
				Product: "windows",
				Version: "10",
			},
			expected: "cpe:/a:microsoft:windows:10:*",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BindToURI(tt.wfn); got != tt.expected {
				t.Errorf("BindToURI() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConvertURIToFS(t *testing.T) {
	tests := []struct {
		name     string
		uri      string
		expected string
		wantErr  bool
	}{
		{
			name:     "basic conversion",
			uri:      "cpe:/a:microsoft:windows:10",
			expected: "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
			wantErr:  false,
		},
		{
			name:    "invalid URI",
			uri:     "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConvertURIToFS(tt.uri)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertURIToFS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.expected {
				t.Errorf("ConvertURIToFS() = %q, want %q", got, tt.expected)
			}
		})
	}
}

func TestConvertFSToURI(t *testing.T) {
	tests := []struct {
		name    string
		fs      string
		wantErr bool
	}{
		{
			name:    "valid conversion",
			fs:      "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
			wantErr: false,
		},
		{
			name:    "invalid FS",
			fs:      "invalid",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ConvertFSToURI(tt.fs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertFSToURI() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestConvertCpe22ToCpe23Binding(t *testing.T) {
	tests := []struct {
		name     string
		cpe22    string
		expected string
	}{
		{
			name:     "basic conversion",
			cpe22:    "cpe:/a:microsoft:windows:10",
			expected: "cpe:2.3:a:microsoft:windows:10:*:*:*:*:*:*:*",
		},
		{
			name:     "invalid format",
			cpe22:    "invalid",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertCpe22ToCpe23(tt.cpe22); got != tt.expected {
				t.Errorf("convertCpe22ToCpe23() = %q, want %q", got, tt.expected)
			}
		})
	}
}
