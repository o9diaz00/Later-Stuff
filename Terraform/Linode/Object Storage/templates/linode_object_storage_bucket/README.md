## Providers

| Name | Version |
|------|---------|
| <a name="provider_linode"></a> [linode](#provider\_linode) | 2.41.0 |

## Resources

| Name | Type |
|------|------|
| [linode_object_storage_bucket.bucket](https://registry.terraform.io/providers/linode/linode/2.41.0/docs/resources/object_storage_bucket) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_abort_incomplete_multipart_upload_days"></a> [abort\_incomplete\_multipart\_upload\_days](#input\_abort\_incomplete\_multipart\_upload\_days) | Specifies the number of days after initiating a multipart upload when the multipart upload must be completed. | `number` | `1` | no |
| <a name="input_expiration_days"></a> [expiration\_days](#input\_expiration\_days) | Specifies a period in the object's expire.  | `number` | n/a | yes |
| <a name="input_label"></a> [label](#input\_label) | The label of the Linode Object Storage Bucket. | `string` | n/a | yes |
| <a name="input_location"></a> [location](#input\_location) | Map containing region : s3-endpoint. | `map(string)` | `null` | no |
| <a name="input_noncurrent_version_expiration"></a> [noncurrent\_version\_expiration](#input\_noncurrent\_version\_expiration) | Specifies when non-current object versions expire.  | `number` | `1` | no |
| <a name="input_region"></a> [region](#input\_region) | THe region of the Linode Object Storage Bucket.  Exactly one of region and cluster is required for creating a bucket.  | `string` | `""` | no |
| <a name="input_S3_endpoint"></a> [s3\_endpoint](#input\_s3\_endpoint) | THe user's s3 endpoint URL, based on the endpoint\_type and region. | `string` | `""` | no |
