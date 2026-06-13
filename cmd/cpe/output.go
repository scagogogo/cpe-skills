package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/scagogogo/cpe"
)

// outputCPE 按指定格式输出单个 CPE 信息
func outputCPE(w io.Writer, c *cpe.CPE, format string) error {
	if format == "json" {
		return outputCPEJSON(w, c)
	}
	return outputCPEText(w, c)
}

func outputCPEText(w io.Writer, c *cpe.CPE) error {
	fmt.Fprintf(w, "CPE 2.3 URI: %s\n", c.GetURI())
	fmt.Fprintf(w, "Part:        %s (%s)\n", c.Part.ShortName, c.Part.LongName)
	fmt.Fprintf(w, "Vendor:      %s\n", c.Vendor)
	fmt.Fprintf(w, "Product:     %s\n", c.ProductName)
	fmt.Fprintf(w, "Version:     %s\n", c.Version)
	fmt.Fprintf(w, "Update:      %s\n", c.Update)
	fmt.Fprintf(w, "Edition:     %s\n", c.Edition)
	fmt.Fprintf(w, "Language:    %s\n", c.Language)
	return nil
}

func outputCPEJSON(w io.Writer, c *cpe.CPE) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(c)
}

// outputError 按统一格式输出错误
func outputError(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, "Error: "+format+"\n", args...)
}
