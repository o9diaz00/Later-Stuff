package cmd

import (
    "fmt"
    "context"
    "log"
    "time"
    "github.com/aws/aws-sdk-go-v2/aws"
    "github.com/aws/aws-sdk-go-v2/config"
    "github.com/aws/aws-sdk-go-v2/service/s3"
    s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
    "github.com/spf13/cobra"
)

var (
    bucket string
    prefix string
    region string
    concurrency int
    profile string
    endpoint string
)

var getStatsCmd = &cobra.Command {
    Use:    "stats",
    Short:  "Get the size of the bucket (or directory specified with prefix)",
    Run:    get_bucket_stats,
}

type customResolver struct {
    endpoint string
    region string
}

type Result struct {
    Bucket string
    Prefix string
    Objects []s3types.Object
    Error error
}

func (r customResolver) ResolveEndpoint(service, region string) (aws.Endpoint, error) {
    return aws.Endpoint {
        URL:    r.endpoint,
        SigningRegion:  r.region,
        HostnameImmutable: true,
    }, nil
}

func s3Client(ctx context.Context) (*s3.Client) {
    cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(profile), config.WithEndpointResolver(customResolver{region: region, endpoint: endpoint}))
    if (err != nil) {
        log.Fatal(err);
    }

    return s3.NewFromConfig(cfg);
}

func convertBytes(size float64) (string) {
    t := []string{"B", "K", "M", "G", "T", "P"};
    var n = 0;

    for size > 1000 {
        size /= 1024;
        n += 1;
    }
    return fmt.Sprintf("%.2f%s", size, t[n]);
}

func getBucketData(client *s3.Client, bucket string, prefix string) ([]s3types.Object, error) {
    var data []s3types.Object;
    paginator := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input {
        Bucket: aws.String(bucket),
        Prefix: aws.String(prefix),
        })

    for paginator.HasMorePages() {
        page, err := paginator.NextPage(context.Background());
        if (err != nil) {
            return nil, err;
        }
        data = append(data, page.Contents...);
    }

    return data, nil;
}

func getBucketCount(data []s3types.Object) (int) {
    return len(data);
}

func getBucketSize(data []s3types.Object) (int64) {
    var size int64;
    for _, d := range data {
        size += *d.Size;
    }
    return size;
}

func getBucketLast(data []s3types.Object) (string, time.Time) {
    var latest time.Time;
    var name string;
    for _, d := range data {
        if (d.LastModified.After(latest)) {
            latest = *d.LastModified;
            name = *d.Key;
        }
    }
    return name, latest;
}

func get_bucket_stats(cmd *cobra.Command, args []string) {
    ctx := context.Background();
    client := s3Client(ctx);
    data, err := getBucketData(client, bucket, prefix);
    last_key, last_time := getBucketLast(data);
    if (err != nil) {
        log.Fatalf("%v", err);
    }
    fmt.Printf("[%v]\n", bucket);
    fmt.Printf(" - Total Objects: %d\n", getBucketCount(data));
    fmt.Printf(" - Total Size: %v\n", convertBytes(float64(getBucketSize(data))));
    fmt.Printf(" - Last Modified Object: %s\n - Last Modified Time: %s\n", last_key, last_time);
}

func init() {
    getStatsCmd.Flags().StringVarP(&bucket, "bucket", "b", "", "The bucket in which we want to check (required)");
    getStatsCmd.Flags().StringVarP(&prefix, "path", "p", "", "The subdirectory within the bucket we want to check");
    getStatsCmd.Flags().StringVarP(&region, "region", "r", "", "The region at which the bucket is located");
    getStatsCmd.Flags().StringVarP(&profile, "profile", "P", "default", "The profile used in the configuration file to authenticate with the endpoint");
    getStatsCmd.Flags().StringVarP(&endpoint, "endpoint", "", "", "The endpoint at which the bucket exists");
    _ = getStatsCmd.MarkFlagRequired("bucket");
    _ = getStatsCmd.MarkFlagRequired("region");
    _ = getStatsCmd.MarkFlagRequired("endpoint");
}
