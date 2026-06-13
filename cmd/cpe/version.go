package main

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	cliVersion   = "0.1.0"
	cliGitCommit = "unknown"
	cliBuildDate = "unknown"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `Print the version, git commit, build date, and Go runtime information.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("cpe CLI:     %s\n", cliVersion)
		fmt.Printf("Git Commit:  %s\n", cliGitCommit)
		fmt.Printf("Build Date:  %s\n", cliBuildDate)
		fmt.Printf("Go Version:  %s\n", runtime.Version())
		fmt.Printf("OS/Arch:     %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
