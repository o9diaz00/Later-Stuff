## Providers

| Name | Version |
|------|---------|
| <a name="provider_linode"></a> [linode](#provider\_linode) | 2.41.0 |

## Resources

| Name | Type |
|------|------|
| [linode_object_storage_key.key](https://registry.terraform.io/providers/linode/linode/2.41.0/docs/resources/object_storage_key) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_bucket_access"></a> [bucket\_access](#input\_bucket\_access) | Map string containing 'bucket: region'.  (Must be set if 'bucket\_name' and 'permissions' is not set) | <pre>list(object({<br/>  bucket_name = string<br/>  region  = string<br/>  permissions = string }))</pre> | `null` | no |
| <a name="input_bucket_name"></a> [bucket\_name](#input\_bucket\_name) | The unique label of the bucket to which the key will grant limited access.  (Must be set if 'bucket\_access' is not set) | `string` | `""` | no |
| <a name="input_label"></a> [label](#input\_label) | The label given to this key.  For display purposes only.  | `string` | n/a | yes |
| <a name="input_permissions"></a> [permissions](#input\_permissions) | This Limited Access Key's permissions for the selected bucket.  (Must be set if 'bucket\_access' is not set) | `string` | `"read_only"` | no |
| <a name="input_region"></a> [region](#input\_region) | The region where the bucket resides.  | `string` | `""` | no |

## Outputs

| Name | Description |
|------|-------------|
| <a name="output_linode_object_storage_secret_key"></a> [linode\_object\_storage\_secret\_key](#output\_linode\_object\_storage\_secret\_key) | secret key corresponding to the created access key |
