package cmd

import (
    "fmt"
    "context"
    "os"
    "sync"
    "github.com/spf13/cobra"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
)

var getSizeCmd = &cobra.Command {
    Use:    "get_size",
    Short:  "Get the size of the bucket (or directory specified with prefix)",
    Run:    func(cmd *cobra.Command, args []string) {
        getSize();
    },
}

func convert_bytes(size int) {
    t := ["B", "K", "M", "G", "T", "P"];
    n := 0;

    for (math.floor(size) > 1000) {
        size /= 1024;
        n += 1;
    }
    return (size, t[n])
}