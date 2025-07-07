resource "linode_object_storage_key" "key" {
  label = join("_", [var.label, replace(var.permissions, "_", "")])

  dynamic "bucket_access" {
    for_each = var.bucket_access
    content {
      bucket_name = bucket_access.value.bucket_name
      region      = bucket_access.value.region
      permissions = bucket_access.value.permissions != "." ? bucket_access.value.permissions : var.permissions
    }
  }
}
