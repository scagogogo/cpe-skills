package cpeskills

import (
	"testing"
	"time"
)

func TestNewKEVClient(t *testing.T) {
	client := NewKEVClient()
	if client == nil {
		t.Fatal("NewKEVClient returned nil")
	}
	if client.BaseURL != DefaultKEVBaseURL {
		t.Errorf("expected BaseURL %s, got %s", DefaultKEVBaseURL, client.BaseURL)
	}
	if client.HTTPClient == nil {
		t.Error("HTTPClient is nil")
	}
	if client.cache == nil {
		t.Error("cache is nil")
	}
}

func TestNewKEVClientWithOptions(t *testing.T) {
	client := NewKEVClientWithOptions("", 0)
	if client == nil {
		t.Fatal("NewKEVClientWithOptions returned nil")
	}
	if client.BaseURL != DefaultKEVBaseURL {
		t.Errorf("expected default BaseURL, got %s", client.BaseURL)
	}

	client2 := NewKEVClientWithOptions("https://custom.kev/api", 120e9)
	if client2.BaseURL != "https://custom.kev/api" {
		t.Errorf("expected custom BaseURL, got %s", client2.BaseURL)
	}
}

func TestKEVClientClearCache(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-44228"] = &KEVEntry{CVEID: "CVE-2021-44228"}
	client.allCache = []*KEVEntry{{CVEID: "CVE-2021-44228"}}
	client.ClearCache()
	if len(client.cache) != 0 {
		t.Errorf("cache not cleared, got %d entries", len(client.cache))
	}
	if client.allCache != nil {
		t.Error("allCache not cleared")
	}
}

func TestKEVSeverityBoost(t *testing.T) {
	tests := []struct {
		current  string
		expected string
	}{
		{"Low", "Medium"},
		{"Medium", "High"},
		{"High", "Critical"},
		{"Critical", "Critical"},
		{"Unknown", "High"},
	}

	for _, tt := range tests {
		result := KEVSeverityBoost(tt.current)
		if result != tt.expected {
			t.Errorf("KEVSeverityBoost(%q) = %q, want %q", tt.current, result, tt.expected)
		}
	}
}

func TestKEVClientFilterByVendor(t *testing.T) {
	client := NewKEVClient()
	client.allCache = []*KEVEntry{
		{CVEID: "CVE-2021-0001", VendorProject: "Apache Software Foundation", Product: "Log4j"},
		{CVEID: "CVE-2021-0002", VendorProject: "Microsoft Corporation", Product: "Windows"},
		{CVEID: "CVE-2021-0003", VendorProject: "Apache Software Foundation", Product: "Tomcat"},
	}
	client.cacheExpiry = time.Now().Add(2 * time.Hour) // future expiry

	results, err := client.FilterByVendor("apache")
	if err != nil {
		t.Fatalf("FilterByVendor failed: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 results for 'apache', got %d", len(results))
	}

	results2, err := client.FilterByVendor("microsoft")
	if err != nil {
		t.Fatalf("FilterByVendor failed: %v", err)
	}
	if len(results2) != 1 {
		t.Errorf("expected 1 result for 'microsoft', got %d", len(results2))
	}

	results3, err := client.FilterByVendor("nonexistent")
	if err != nil {
		t.Fatalf("FilterByVendor failed: %v", err)
	}
	if len(results3) != 0 {
		t.Errorf("expected 0 results for 'nonexistent', got %d", len(results3))
	}
}

func TestKEVClientFilterByProduct(t *testing.T) {
	client := NewKEVClient()
	client.allCache = []*KEVEntry{
		{CVEID: "CVE-2021-0001", VendorProject: "Apache", Product: "Log4j"},
		{CVEID: "CVE-2021-0002", VendorProject: "Microsoft", Product: "Windows"},
		{CVEID: "CVE-2021-0003", VendorProject: "Apache", Product: "Log4j"},
	}
	client.cacheExpiry = time.Now().Add(2 * time.Hour)

	results, err := client.FilterByProduct("log4j")
	if err != nil {
		t.Fatalf("FilterByProduct failed: %v", err)
	}
	if len(results) != 2 {
		t.Errorf("expected 2 results for 'log4j', got %d", len(results))
	}
}

func TestKEVClientCount(t *testing.T) {
	client := NewKEVClient()
	client.allCache = []*KEVEntry{
		{CVEID: "CVE-2021-0001"},
		{CVEID: "CVE-2021-0002"},
		{CVEID: "CVE-2021-0003"},
	}
	client.cacheExpiry = time.Now().Add(2 * time.Hour)

	count, err := client.Count()
	if err != nil {
		t.Fatalf("Count failed: %v", err)
	}
	if count != 3 {
		t.Errorf("expected count 3, got %d", count)
	}
}

func TestKEVClientGetAll(t *testing.T) {
	client := NewKEVClient()
	client.allCache = []*KEVEntry{
		{CVEID: "CVE-2021-0001"},
		{CVEID: "CVE-2021-0002"},
	}
	client.cacheExpiry = time.Now().Add(2 * time.Hour)

	all, err := client.GetAll()
	if err != nil {
		t.Fatalf("GetAll failed: %v", err)
	}
	if len(all) != 2 {
		t.Errorf("expected 2 entries, got %d", len(all))
	}
}

func TestKEVClientGetEntryCached(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-44228"] = &KEVEntry{
		CVEID:       "CVE-2021-44228",
		VendorProject: "Apache",
		Product:     "Log4j",
	}

	entry, err := client.GetEntry("CVE-2021-44228")
	if err != nil {
		t.Fatalf("GetEntry failed: %v", err)
	}
	if entry == nil {
		t.Fatal("expected entry, got nil")
	}
	if entry.VendorProject != "Apache" {
		t.Errorf("expected VendorProject 'Apache', got %q", entry.VendorProject)
	}
}

func TestKEVClientGetEntryEmpty(t *testing.T) {
	client := NewKEVClient()
	_, err := client.GetEntry("")
	if err == nil {
		t.Error("expected error for empty CVE ID")
	}
}

func TestKEVClientGetEntriesCached(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-0001"] = &KEVEntry{CVEID: "CVE-2021-0001"}
	client.cache["CVE-2021-0002"] = &KEVEntry{CVEID: "CVE-2021-0002"}

	entries, err := client.GetEntries([]string{"CVE-2021-0001", "CVE-2021-0002", "CVE-2021-0003"})
	if err != nil {
		t.Fatalf("GetEntries failed: %v", err)
	}
	if len(entries) != 2 {
		t.Errorf("expected 2 entries from cache, got %d", len(entries))
	}
}

func TestKEVClientIsListedCached(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-44228"] = &KEVEntry{CVEID: "CVE-2021-44228"}

	listed, err := client.IsListed("CVE-2021-44228")
	if err != nil {
		t.Fatalf("IsListed failed: %v", err)
	}
	if !listed {
		t.Error("expected CVE-2021-44228 to be listed")
	}

	listed2, err := client.IsListed("CVE-2099-99999")
	if err != nil {
		t.Fatalf("IsListed failed: %v", err)
	}
	if listed2 {
		t.Error("expected CVE-2099-99999 to not be listed")
	}
}

func TestKEVClientIsRansomwareRelated(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-0001"] = &KEVEntry{
		CVEID:                       "CVE-2021-0001",
		KnownRansomwareCampaignUse:  "Known",
	}
	client.cache["CVE-2021-0002"] = &KEVEntry{
		CVEID:                       "CVE-2021-0002",
		KnownRansomwareCampaignUse:  "Unknown",
	}

	related, err := client.IsRansomwareRelated("CVE-2021-0001")
	if err != nil {
		t.Fatalf("IsRansomwareRelated failed: %v", err)
	}
	if !related {
		t.Error("expected CVE-2021-0001 to be ransomware related")
	}

	related2, err := client.IsRansomwareRelated("CVE-2021-0002")
	if err != nil {
		t.Fatalf("IsRansomwareRelated failed: %v", err)
	}
	if related2 {
		t.Error("expected CVE-2021-0002 to not be ransomware related")
	}
}

func TestKEVClientGetDueDate(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-0001"] = &KEVEntry{
		CVEID:   "CVE-2021-0001",
		DueDate: "2022-01-15",
	}

	dueDate, err := client.GetDueDate("CVE-2021-0001")
	if err != nil {
		t.Fatalf("GetDueDate failed: %v", err)
	}
	if dueDate != "2022-01-15" {
		t.Errorf("expected DueDate '2022-01-15', got %q", dueDate)
	}
}

func TestKEVClientGetRequiredAction(t *testing.T) {
	client := NewKEVClient()
	client.cache["CVE-2021-0001"] = &KEVEntry{
		CVEID:          "CVE-2021-0001",
		RequiredAction: "Apply vendor patch",
	}

	action, err := client.GetRequiredAction("CVE-2021-0001")
	if err != nil {
		t.Fatalf("GetRequiredAction failed: %v", err)
	}
	if action != "Apply vendor patch" {
		t.Errorf("expected 'Apply vendor patch', got %q", action)
	}
}
