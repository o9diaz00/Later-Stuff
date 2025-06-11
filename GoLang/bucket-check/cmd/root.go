package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:    "bucket-check",
    Short:  "Tool for checking out stats on a particular bucket as well as performing deletions",
}

func init() {
    rootCmd.AddCommand(getStatsCmd);
    }

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err);
        os.Exit(1);
    }
}