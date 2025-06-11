package cmd

import (
    "fmt"
    "os"
    "github.com/spf13/cobra"
)

var (
    bucket string
    prefix string
    endpoint string
    concurrency int
    )

var rootCmd = &cobra.Command{
    Use:    "bucket-check",
    Short:  "Tool for checking out stats on a particular bucket as well as performing deletions",
}

func init() {
    rootCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "The bucket in which we want to check (required)")
    rootCmd.Flags().StringVarP(&prefix, "path", "p" "", "The subdirectory within the bucket we want to check")
    rootCmd.Flags().StringVarP(&endpoint, "endpoint", "https://us-iad-8.linodeobjects.com", "The endpoint at which the bucket is located")
    rootCmd.Flags().IntVarP(&concurrency, "c", 5, "Number of concurrent workers")

    _ = rootCmd.MarkFlagRequired("bucket")
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err);
        os.Exit(1);
    }
}