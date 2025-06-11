package cmd

import (
    "fmt"
    "context"
    "log"
    //"os"
    "sync"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    //"github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/spf13/cobra"
)

var (
    bucket string
    prefix string
    region string
    concurrency int
    profile string
)

type customResolver struct {
    endpoint string
    region string
}

func (r customResolver) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
    return aws.Endpoint {
        URL:    "https://" + r.region + ".linodeobjects.com",
        SigningRegion:  r.region,
        HostnameImmutable: true,
    }, nil
}

func s3Client(ctx context.Context) (*s3.Client) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile), config.WithEndpointResolver(customResolver{region: region}))
    if (err != nil) {
        log.Fatal(err)
    }

    return s3.NewFromConfig(cfg);
}

var getStatsCmd = &cobra.Command {
    Use:    "stats",
    Short:  "Get the size of the bucket (or directory specified with prefix)",
    Run:    getSizeAndCount,
}

func getSizeAndCount(cmd *cobra.Command, args []string) {
    ctx := context.Background();
    client := s3Client(ctx);

    var totalSize int64;
    var totalCount int;
    var wg sync.WaitGroup;
    var mu sync.Mutex;

    pageCh := make(chan *s3.ListObjectsV2Output);

    for i := 0; i<concurrency; i+=1 {
        wg.Add(1);
        go func() {
            defer wg.Done();
            for page := range pageCh {
                var size int64;
                var count = len(page.Contents);
                for _, obj := range page.Contents {
                    size += *obj.Size;
                }
                mu.Lock();
                totalSize += size;
                totalCount += count;
                mu.Unlock();
            }
        }()
    }

    paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{
        Bucket: aws.String(bucket),
        Prefix: aws.String(prefix),
    })

    for paginator.HasMorePages() {
        page, err := paginator.NextPage(ctx);
        if (err != nil) {
            log.Fatalf("Error getting S3 page: %v", err);
        }
        pageCh <- page
    }

    close(pageCh);
    wg.Wait();

    fmt.Printf("[%v]\nTotal Size: %d\nObject Count: %d\n", bucket, totalSize, totalCount);
}


func getLast() {
    fmt.Printf("This gets the last file uploaded to the bucket: %v", bucket);
}

func init() {
    getStatsCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "The bucket in which we want to check (required)")
    getStatsCmd.Flags().StringVarP(&prefix, "path", "p", "", "The subdirectory within the bucket we want to check")
    getStatsCmd.Flags().StringVarP(&region, "region", "r", "us-iad-8", "The region at which the bucket is located")
    getStatsCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 5, "Number of concurrent workers")
    getStatsCmd.Flags().StringVarP(&profile, "profile", "P", "default", "The profile used in the configuration file to authenticate with the endpoint")
    _ = getStatsCmd.MarkFlagRequired("bucket")
}