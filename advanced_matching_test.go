package cpe

import (
	"testing"
)

// TestNewAdvancedMatchOptions 测试创建高级匹配选项
func TestNewAdvancedMatchOptions(t *testing.T) {
	options := NewAdvancedMatchOptions()

	// 检查默认值
	if options.UseRegex {
		t.Errorf("NewAdvancedMatchOptions().UseRegex = %v, want false", options.UseRegex)
	}

	if options.IgnoreCase {
		t.Errorf("NewAdvancedMatchOptions().IgnoreCase = %v, want false", options.IgnoreCase)
	}

	if options.UseFuzzyMatch {
		t.Errorf("NewAdvancedMatchOptions().UseFuzzyMatch = %v, want false", options.UseFuzzyMatch)
	}

	if options.MatchCommonOnly {
		t.Errorf("NewAdvancedMatchOptions().MatchCommonOnly = %v, want false", options.MatchCommonOnly)
	}

	if options.PartialMatch {
		t.Errorf("NewAdvancedMatchOptions().PartialMatch = %v, want false", options.PartialMatch)
	}

	if options.MatchMode != "exact" {
		t.Errorf("NewAdvancedMatchOptions().MatchMode = %v, want 'exact'", options.MatchMode)
	}

	if options.VersionCompareMode != "exact" {
		t.Errorf("NewAdvancedMatchOptions().VersionCompareMode = %v, want 'exact'", options.VersionCompareMode)
	}

	if options.FieldOptions == nil {
		t.Errorf("NewAdvancedMatchOptions().FieldOptions is nil, want a map")
	}
}

// TestAdvancedMatchCPE 测试高级CPE匹配
func TestAdvancedMatchCPE(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "精确匹配模式",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchMode: "exact",
			},
			expected: true,
		},
		{
			name: "忽略大小写匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "Microsoft",
				ProductName: "Windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				IgnoreCase: true,
			},
			expected: true,
		},
		{
			name: "正则表达式匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "1.*",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				UseRegex: true,
			},
			expected: true,
		},
		{
			name: "部分匹配模式",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				PartialMatch: true,
			},
			expected: true,
		},
		{
			name: "子集匹配模式",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchMode: "subset",
			},
			expected: true,
		},
		{
			name: "超集匹配模式",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
				Update:      "pro",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				MatchMode: "superset",
			},
			expected: true,
		},
		{
			name: "仅匹配常见字段",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				SoftwareEdition: "professional",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				SoftwareEdition: "home",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: true,
			},
			expected: true,
		},
		{
			name: "版本范围匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionLower:       "9",
				VersionUpper:       "12",
			},
			expected: true,
		},
		{
			name: "不匹配 - 不同的产品",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "office",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "不匹配 - 不同的供应商",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "adobe",
				ProductName: "reader",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "reader",
			},
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvancedMatchCPE(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("AdvancedMatchCPE() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchWithRegex 测试正则表达式匹配
func TestMatchWithRegex(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "简单正则匹配",
			criteria: &CPE{
				Vendor:      "micro.*",
				ProductName: "windows",
			},
			target: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				UseRegex: true,
			},
			expected: true,
		},
		{
			name: "精确版本正则",
			criteria: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10\\.[0-9]+",
			},
			target: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10.2",
			},
			options: &AdvancedMatchOptions{
				UseRegex: true,
			},
			expected: true,
		},
		{
			name: "版本不匹配正则",
			criteria: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11\\.[0-9]+",
			},
			target: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10.2",
			},
			options: &AdvancedMatchOptions{
				UseRegex: true,
			},
			expected: false,
		},
		{
			name: "忽略大小写匹配",
			criteria: &CPE{
				Vendor:      "Microsoft",
				ProductName: "Windows",
			},
			target: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				UseRegex:   true,
				IgnoreCase: true,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchWithRegex(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchWithRegex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchNonVersionFields 测试非版本字段匹配
func TestMatchNonVersionFields(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "基本字段匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options:  &AdvancedMatchOptions{},
			expected: true,
		},
		{
			name: "忽略大小写",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "Microsoft",
				ProductName: "Windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				IgnoreCase: true,
			},
			expected: true,
		},
		{
			name: "字段不匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "office",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "需要字段权重 - 部分匹配",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Edition:     "professional",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Edition:     "home",
			},
			options: &AdvancedMatchOptions{
				FieldOptions: map[string]FieldMatchOption{
					"part":    {Weight: 0.3, Required: true},
					"vendor":  {Weight: 0.3, Required: true},
					"product": {Weight: 0.3, Required: true},
					"edition": {Weight: 0.1, Required: false},
				},
				ScoreThreshold: 0.9, // 90%匹配度
			},
			expected: true, // 因为主要字段都匹配，总分超过阈值
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchNonVersionFields(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchNonVersionFields() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestAdvancedMatchCPEExtended tests additional branches of AdvancedMatchCPE
func TestAdvancedMatchCPEExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name:     "nil criteria returns false",
			criteria: nil,
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options:  NewAdvancedMatchOptions(),
			expected: false,
		},
		{
			name: "nil target returns false",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target:   nil,
			options:  NewAdvancedMatchOptions(),
			expected: false,
		},
		{
			name: "nil options uses defaults",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options:  nil,
			expected: true,
		},
		{
			name: "distance mode match",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchMode:      "distance",
				ScoreThreshold: 0.7,
				FieldOptions:   make(map[string]FieldMatchOption),
			},
			expected: true,
		},
		{
			name: "distance mode no match - threshold too high",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "adobe",
				ProductName: "reader",
			},
			options: &AdvancedMatchOptions{
				MatchMode:      "distance",
				ScoreThreshold: 0.99,
				FieldOptions:   make(map[string]FieldMatchOption),
			},
			expected: false,
		},
		{
			name: "unknown match mode falls through to default",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchMode: "unknown_mode",
			},
			expected: true,
		},
		{
			name: "fuzzy match - containment",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "micro",
				ProductName: "win",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				UseFuzzyMatch:      true,
				VersionCompareMode: "exact",
			},
			expected: true,
		},
		{
			name: "subset mode - match with non-common fields",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
				Update:      "sp1",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
				Update:      "sp1",
			},
			options: &AdvancedMatchOptions{
				MatchMode:     "subset",
				MatchCommonOnly: false,
			},
			expected: true,
		},
		{
			name: "superset mode - criteria version empty returns false",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchMode:      "superset",
				MatchCommonOnly: false,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AdvancedMatchCPE(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("AdvancedMatchCPE() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchCommonFieldsExtended tests additional branches of matchCommonFields
func TestMatchCommonFieldsExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "version compare mode greater",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: true,
		},
		{
			name: "version compare mode greater fails",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "9",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: false,
		},
		{
			name: "version compare mode greaterOrEqual",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greaterOrEqual",
			},
			expected: true,
		},
		{
			name: "version compare mode less",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "9",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "less",
			},
			expected: true,
		},
		{
			name: "version compare mode lessOrEqual",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "lessOrEqual",
			},
			expected: true,
		},
		{
			name: "version compare mode range - in range",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionLower:       "9",
				VersionUpper:       "12",
			},
			expected: true,
		},
		{
			name: "version compare mode range - out of range high",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "13",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionLower:       "9",
				VersionUpper:       "12",
			},
			expected: false,
		},
		{
			name: "version compare mode range - out of range low",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "8",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionLower:       "9",
				VersionUpper:       "12",
			},
			expected: false,
		},
		{
			name: "version compare mode range - only lower bound",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionLower:       "9",
			},
			expected: true,
		},
		{
			name: "version compare mode range - only upper bound",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "8",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "range",
				VersionUpper:       "9",
			},
			expected: true,
		},
		{
			name: "version compare default unknown mode - exact match",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "unknown",
			},
			expected: true,
		},
		{
			name: "version compare default unknown mode - no match",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "unknown",
			},
			expected: false,
		},
		{
			name: "fuzzy match with contains check",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "soft",
				ProductName: "win",
				Version:     "1",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				UseFuzzyMatch:      true,
				VersionCompareMode: "exact",
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchCommonFields(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchCommonFields() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchWithRegexExtended tests additional branches of matchWithRegex
func TestMatchWithRegexExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "regex match with non-common fields",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp.*",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "ent.*",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom.*",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom1",
			},
			options: &AdvancedMatchOptions{
				UseRegex:      true,
				MatchCommonOnly: false,
			},
			expected: true,
		},
		{
			name: "regex match non-common field mismatch",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
				Update:      "sp3",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
				Update:      "sp1",
			},
			options: &AdvancedMatchOptions{
				UseRegex:      true,
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "regex with invalid pattern falls back to exact match",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "[invalid",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "[invalid",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				UseRegex: true,
			},
			expected: true,
		},
		{
			name: "regex with invalid pattern and ignore case falls back to EqualFold",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "[invalid",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "[INVALID",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				UseRegex:   true,
				IgnoreCase: true,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchWithRegex(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchWithRegex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchFieldExtended tests additional branches of matchField
func TestMatchFieldExtended(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name:     "both NA values match",
			a:        "-",
			b:        "-",
			options:  &AdvancedMatchOptions{},
			expected: true,
		},
		{
			name:     "NA does not match regular value",
			a:        "-",
			b:        "value",
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name:     "ignore case match",
			a:        "Windows",
			b:        "windows",
			options:  &AdvancedMatchOptions{IgnoreCase: true},
			expected: true,
		},
		{
			name:     "fuzzy match with contains",
			a:        "soft",
			b:        "microsoft",
			options:  &AdvancedMatchOptions{UseFuzzyMatch: true},
			expected: true,
		},
		{
			name:     "fuzzy match reverse contains",
			a:        "microsoft",
			b:        "soft",
			options:  &AdvancedMatchOptions{UseFuzzyMatch: true},
			expected: true,
		},
		{
			name:     "fuzzy match no containment",
			a:        "adobe",
			b:        "microsoft",
			options:  &AdvancedMatchOptions{UseFuzzyMatch: true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchField(tt.a, tt.b, tt.options); got != tt.expected {
				t.Errorf("matchField() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchFieldWithRegexExtended tests additional branches of matchFieldWithRegex
func TestMatchFieldWithRegexExtended(t *testing.T) {
	tests := []struct {
		name     string
		a        string
		b        string
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name:     "empty a pattern matches",
			a:        "",
			b:        "anything",
			options:  &AdvancedMatchOptions{},
			expected: true,
		},
		{
			name:     "star a pattern matches",
			a:        "*",
			b:        "anything",
			options:  &AdvancedMatchOptions{},
			expected: true,
		},
		{
			name:     "empty b no match",
			a:        "pattern",
			b:        "",
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name:     "NA b no match",
			a:        "pattern",
			b:        "-",
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name:     "star b matches",
			a:        "pattern",
			b:        "*",
			options:  &AdvancedMatchOptions{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchFieldWithRegex(tt.a, tt.b, tt.options); got != tt.expected {
				t.Errorf("matchFieldWithRegex() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchPartialExtended tests additional branches of matchPartial
func TestMatchPartialExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "partial match with non-common fields",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			options: &AdvancedMatchOptions{
				PartialMatch:    true,
				MatchCommonOnly: false,
			},
			expected: true,
		},
		{
			name: "partial match with wildcard criteria fields skipped",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "*",
				ProductName: "windows",
				Version:     "*",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				PartialMatch: true,
			},
			expected: true,
		},
		{
			name: "partial match version compare mode",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				PartialMatch:       true,
				VersionCompareMode: "greater",
			},
			expected: true,
		},
		{
			name: "partial match non-common field mismatch",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Edition:     "professional",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Edition:     "home",
			},
			options: &AdvancedMatchOptions{
				PartialMatch:    true,
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "partial match empty criteria fields skipped",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "",
				ProductName: "windows",
				Version:     "",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "anything",
				ProductName: "windows",
				Version:     "anything",
			},
			options: &AdvancedMatchOptions{
				PartialMatch: true,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchPartial(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchPartial() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestCompareVersionsExtended tests additional branches of compareVersions
func TestCompareVersionsExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "wildcard criteria version",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "*",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: true,
		},
		{
			name: "wildcard target version",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "*",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: true,
		},
		{
			name: "NA criteria version",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "-",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: false,
		},
		{
			name: "NA target version",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "-",
			},
			options: &AdvancedMatchOptions{
				VersionCompareMode: "greater",
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareVersions(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("compareVersions() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchSubsetExtended tests additional branches of matchSubset
func TestMatchSubsetExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name:     "nil criteria",
			criteria: nil,
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "nil target",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target:   nil,
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "subset match with all non-common fields",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: true,
		},
		{
			name: "subset mismatch on non-common field",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Update:      "sp1",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Update:      "sp2",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "subset with empty part field skipped",
			criteria: &CPE{
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchSubset(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchSubset() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchSupersetExtended tests additional branches of matchSuperset
func TestMatchSupersetExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "nil criteria",
			criteria: nil,
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "nil target",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target:   nil,
			options:  &AdvancedMatchOptions{},
			expected: false,
		},
		{
			name: "superset match with non-common fields",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: true,
		},
		{
			name: "superset mismatch on target vendor",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "adobe",
				ProductName: "windows",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "superset mismatch on target product",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "office",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "superset with target version and criteria version empty fails",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: false,
		},
		{
			name: "superset with target version and criteria wildcard version fails",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "*",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchSuperset(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchSuperset() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchDistanceExtended tests additional branches of matchDistance
func TestMatchDistanceExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "distance match with required field that fails",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "adobe",
				ProductName: "reader",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				ScoreThreshold: 0.7,
				FieldOptions: map[string]FieldMatchOption{
					"vendor": {Weight: 1.0, Required: true},
				},
			},
			expected: false,
		},
		{
			name: "distance match with version compare non-exact",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "11",
			},
			options: &AdvancedMatchOptions{
				ScoreThreshold:     0.7,
				VersionCompareMode: "greater",
				FieldOptions:       make(map[string]FieldMatchOption),
			},
			expected: true,
		},
		{
			name: "distance match with non-common fields",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Version:         "10",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			options: &AdvancedMatchOptions{
				ScoreThreshold:   0.7,
				MatchCommonOnly:  false,
				FieldOptions:     make(map[string]FieldMatchOption),
			},
			expected: true,
		},
		{
			name: "distance match with custom weights",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Version:     "10",
			},
			options: &AdvancedMatchOptions{
				ScoreThreshold: 0.7,
				FieldOptions: map[string]FieldMatchOption{
					"part":    {Weight: 2.0},
					"vendor":  {Weight: 2.0},
					"product": {Weight: 2.0},
				},
			},
			expected: true,
		},
		{
			name: "distance match below threshold",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "adobe",
				ProductName: "reader",
			},
			options: &AdvancedMatchOptions{
				ScoreThreshold: 0.9,
				FieldOptions:   make(map[string]FieldMatchOption),
			},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchDistance(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchDistance() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestMatchNonVersionFieldsExtended tests additional branches of matchNonVersionFields
func TestMatchNonVersionFieldsExtended(t *testing.T) {
	tests := []struct {
		name     string
		criteria *CPE
		target   *CPE
		options  *AdvancedMatchOptions
		expected bool
	}{
		{
			name: "non-version fields match with score below threshold",
			criteria: &CPE{
				Part:        *PartApplication,
				Vendor:      "novendor",
				ProductName: "noproduct",
				Edition:     "different",
				Language:    "different",
			},
			target: &CPE{
				Part:        *PartApplication,
				Vendor:      "microsoft",
				ProductName: "windows",
				Edition:     "pro",
				Language:    "en",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
				ScoreThreshold:  0.99,
			},
			expected: false,
		},
		{
			name: "non-version fields match all non-common fields matching",
			criteria: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			target: &CPE{
				Part:            *PartApplication,
				Vendor:          "microsoft",
				ProductName:     "windows",
				Update:          "sp1",
				Edition:         "pro",
				Language:        "en",
				SoftwareEdition: "enterprise",
				TargetSoftware:  "linux",
				TargetHardware:  "x86",
				Other:           "custom",
			},
			options: &AdvancedMatchOptions{
				MatchCommonOnly: false,
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matchNonVersionFields(tt.criteria, tt.target, tt.options); got != tt.expected {
				t.Errorf("matchNonVersionFields() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// TestIsRequiredFieldExtended tests isRequiredField edge case
func TestIsRequiredFieldExtended(t *testing.T) {
	// Test with nil FieldOptions map - should return false
	options := &AdvancedMatchOptions{
		FieldOptions: nil,
	}
	if isRequiredField(options, "vendor") {
		t.Error("isRequiredField with nil FieldOptions should return false")
	}
}
