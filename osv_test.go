package cpeskills

import (
	"fmt"
	"testing"
	"time"
)

func TestNewOSVClient(t *testing.T) {
	client := NewOSVClient()
	if client == nil {
		t.Fatal("NewOSVClient returned nil")
	}
	if client.BaseURL != DefaultOSVBaseURL {
		t.Errorf("expected BaseURL %s, got %s", DefaultOSVBaseURL, client.BaseURL)
	}
	if client.HTTPClient == nil {
		t.Error("HTTPClient is nil")
	}
	if client.RetryCount != 3 {
		t.Errorf("expected RetryCount 3, got %d", client.RetryCount)
	}
}

func TestNewOSVClientWithOptions(t *testing.T) {
	client := NewOSVClientWithOptions("", 0, 0)
	if client == nil {
		t.Fatal("NewOSVClientWithOptions returned nil")
	}
	if client.BaseURL != DefaultOSVBaseURL {
		t.Errorf("expected default BaseURL, got %s", client.BaseURL)
	}
}

func TestOSVEntryGetFixedVersion(t *testing.T) {
	entry := &OSVEntry{
		Affected: []*OSVAffected{
			{
				Ranges: []*OSVRange{
					{
						Events: []*OSVEvent{
							{Introduced: "2.0.0"},
							{Fixed: "2.17.0"},
						},
					},
				},
			},
		},
	}

	fixed := entry.GetFixedVersion()
	if fixed != "2.17.0" {
		t.Errorf("expected fixed version '2.17.0', got %q", fixed)
	}
}

func TestOSVEntryGetFixedVersionEmpty(t *testing.T) {
	entry := &OSVEntry{}
	fixed := entry.GetFixedVersion()
	if fixed != "" {
		t.Errorf("expected empty fixed version, got %q", fixed)
	}
}

func TestOSVEntryGetFixedVersionNil(t *testing.T) {
	var entry *OSVEntry
	fixed := entry.GetFixedVersion()
	if fixed != "" {
		t.Errorf("expected empty fixed version for nil, got %q", fixed)
	}
}

func TestOSVEntryGetIntroducedVersion(t *testing.T) {
	entry := &OSVEntry{
		Affected: []*OSVAffected{
			{
				Ranges: []*OSVRange{
					{
						Events: []*OSVEvent{
							{Introduced: "2.14.0"},
							{Fixed: "2.17.0"},
						},
					},
				},
			},
		},
	}

	introduced := entry.GetIntroducedVersion()
	if introduced != "2.14.0" {
		t.Errorf("expected introduced version '2.14.0', got %q", introduced)
	}
}

func TestOSVEntryGetAffectedVersions(t *testing.T) {
	entry := &OSVEntry{
		Affected: []*OSVAffected{
			{Versions: []string{"2.14.0", "2.14.1", "2.15.0"}},
		},
	}

	versions := entry.GetAffectedVersions()
	if len(versions) != 3 {
		t.Fatalf("expected 3 versions, got %d", len(versions))
	}
	if versions[0] != "2.14.0" {
		t.Errorf("expected version '2.14.0', got %q", versions[0])
	}
}

func TestOSVEntryHasCVE(t *testing.T) {
	entry := &OSVEntry{
		Aliases: []string{"CVE-2021-44228", "GHSA-jfh8-c2jp-5v3q"},
	}

	if !entry.HasCVE() {
		t.Error("expected HasCVE=true")
	}
}

func TestOSVEntryHasCVEFalse(t *testing.T) {
	entry := &OSVEntry{
		Aliases: []string{"GHSA-jfh8-c2jp-5v3q"},
	}

	if entry.HasCVE() {
		t.Error("expected HasCVE=false for no CVE aliases")
	}
}

func TestOSVEntryGetCVEIDs(t *testing.T) {
	entry := &OSVEntry{
		Aliases: []string{"CVE-2021-44228", "GHSA-jfh8-c2jp-5v3q", "CVE-2021-45046"},
	}

	cves := entry.GetCVEIDs()
	if len(cves) != 2 {
		t.Fatalf("expected 2 CVE IDs, got %d", len(cves))
	}
	if cves[0] != "CVE-2021-44228" {
		t.Errorf("expected first CVE 'CVE-2021-44228', got %q", cves[0])
	}
	if cves[1] != "CVE-2021-45046" {
		t.Errorf("expected second CVE 'CVE-2021-45046', got %q", cves[1])
	}
}

func TestOSVEntryGetMaxCVSSScore(t *testing.T) {
	entry := &OSVEntry{
		Severity: []*OSVSeverity{
			{Type: "CVSS_V3", Score: "7.5"},
			{Type: "CVSS_V3", Score: "9.8"},
		},
	}

	score := entry.GetMaxCVSSScore()
	if score != 9.8 {
		t.Errorf("expected max CVSS 9.8, got %f", score)
	}
}

func TestOSVEntryGetMaxCVSSScoreZero(t *testing.T) {
	entry := &OSVEntry{}
	score := entry.GetMaxCVSSScore()
	if score != 0.0 {
		t.Errorf("expected CVSS 0.0, got %f", score)
	}
}

func TestOSVEntryGetSeverityLevel(t *testing.T) {
	tests := []struct {
		score    float64
		expected string
	}{
		{9.8, "Critical"},
		{7.5, "High"},
		{5.0, "Medium"},
		{2.0, "Low"},
		{0.0, "Unknown"},
	}

	for _, tt := range tests {
		entry := &OSVEntry{
			Severity: []*OSVSeverity{
				{Type: "CVSS_V3", Score: mustFormatFloat(tt.score)},
			},
		}
		level := entry.GetSeverityLevel()
		if level != tt.expected {
			t.Errorf("expected severity %q for score %.1f, got %q", tt.expected, tt.score, level)
		}
	}
}

func TestOSVEntryGetReferenceURLs(t *testing.T) {
	entry := &OSVEntry{
		References: []*OSVReference{
			{URL: "https://nvd.nist.gov/vuln/detail/CVE-2021-44228"},
			{URL: "https://logging.apache.org/log4j/2.x/security.html"},
		},
	}

	urls := entry.GetReferenceURLs()
	if len(urls) != 2 {
		t.Fatalf("expected 2 URLs, got %d", len(urls))
	}
}

func TestOSVEntryIsWithdrawn(t *testing.T) {
	entry := &OSVEntry{
		DatabaseSpecific: map[string]interface{}{
			"withdrawn": true,
		},
	}
	if !entry.IsWithdrawn() {
		t.Error("expected IsWithdrawn=true")
	}

	entry2 := &OSVEntry{}
	if entry2.IsWithdrawn() {
		t.Error("expected IsWithdrawn=false for nil DatabaseSpecific")
	}
}

func TestOSVEntryToVulnerabilityFinding(t *testing.T) {
	entry := &OSVEntry{
		ID:       "GHSA-jfh8-c2jp-5v3q",
		Summary:  "Log4Shell",
		Published:    time.Date(2021, 12, 10, 0, 0, 0, 0, time.UTC),
		Severity: []*OSVSeverity{
			{Type: "CVSS_V3", Score: "10.0"},
		},
		Affected: []*OSVAffected{
			{
				Ranges: []*OSVRange{
					{
						Events: []*OSVEvent{
							{Introduced: "2.0.0"},
							{Fixed: "2.17.0"},
						},
					},
				},
			},
		},
	}

	finding := entry.ToVulnerabilityFinding()
	if finding == nil {
		t.Fatal("ToVulnerabilityFinding returned nil")
	}
	if finding.FixedVersion != "2.17.0" {
		t.Errorf("expected FixedVersion '2.17.0', got %q", finding.FixedVersion)
	}
	if !finding.FixAvailable {
		t.Error("expected FixAvailable=true")
	}
	if finding.Source != "OSV" {
		t.Errorf("expected Source 'OSV', got %q", finding.Source)
	}
}

func TestOSVEntryToVulnerabilityFindingNil(t *testing.T) {
	var entry *OSVEntry
	finding := entry.ToVulnerabilityFinding()
	if finding != nil {
		t.Error("expected nil for nil OSV entry")
	}
}

func TestParseOSVEntry(t *testing.T) {
	jsonData := `{
		"id": "GHSA-test",
		"summary": "Test vulnerability",
		"aliases": ["CVE-2021-0001"],
		"published": "2021-01-01T00:00:00Z"
	}`

	entry, err := ParseOSVEntry([]byte(jsonData))
	if err != nil {
		t.Fatalf("ParseOSVEntry failed: %v", err)
	}
	if entry.ID != "GHSA-test" {
		t.Errorf("expected ID 'GHSA-test', got %q", entry.ID)
	}
	if !entry.HasCVE() {
		t.Error("expected HasCVE=true")
	}
}

func TestParseOSVEntryInvalid(t *testing.T) {
	_, err := ParseOSVEntry([]byte("not json"))
	if err == nil {
		t.Error("expected error for invalid JSON")
	}
}

func TestParseOSVEntries(t *testing.T) {
	jsonData := `[{"id":"GHSA-1"},{"id":"GHSA-2"}]`

	entries, err := ParseOSVEntries([]byte(jsonData))
	if err != nil {
		t.Fatalf("ParseOSVEntries failed: %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("expected 2 entries, got %d", len(entries))
	}
}

func TestOSVClientQueryNilPURL(t *testing.T) {
	client := NewOSVClient()
	_, err := client.Query(nil)
	if err == nil {
		t.Error("expected error for nil PURL")
	}
}

func TestOSVClientGetVulnerabilityEmptyID(t *testing.T) {
	client := NewOSVClient()
	_, err := client.GetVulnerability("")
	if err == nil {
		t.Error("expected error for empty OSV ID")
	}
}

func TestOSVClientQueryByEcosystem(t *testing.T) {
	client := NewOSVClient()
	_, err := client.QueryByEcosystem("", "test", "1.0")
	if err == nil {
		t.Error("expected error for empty ecosystem")
	}
	_, err = client.QueryByEcosystem("npm", "", "1.0")
	if err == nil {
		t.Error("expected error for empty name")
	}
}

func TestOSVClientQueryByCommit(t *testing.T) {
	client := NewOSVClient()
	_, err := client.QueryByCommit("")
	if err == nil {
		t.Error("expected error for empty commit")
	}
}

func TestOSVClientQueryBatchEmpty(t *testing.T) {
	client := NewOSVClient()
	result, err := client.QueryBatch(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("expected empty result, got %d entries", len(result))
	}
}

// mustFormatFloat is a test helper
func mustFormatFloat(f float64) string {
	return fmt.Sprintf("%.1f", f)
}
