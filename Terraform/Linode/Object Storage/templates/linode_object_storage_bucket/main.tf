resource "linode_object_storage_bucket" "bucket" {
  label       = var.label
  for_each    = var.location != null ? var.location : {}
  region      = each.key
  s3_endpoint = each.value

  lifecycle_rule {
    id                                     = "lifecycle_${var.expiration_days}day_expire"
    enabled                                = true
    abort_incomplete_multipart_upload_days = var.abort_incomplete_multipart_upload_days
    noncurrent_version_expiration {
      days = var.noncurrent_version_expiration
    }
    expiration {
      days = var.expiration_days
    }
  }
}
