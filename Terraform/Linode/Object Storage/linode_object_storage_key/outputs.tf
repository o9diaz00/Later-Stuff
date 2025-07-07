output "linode_object_storage_secret_key" {
  description = "secret key corresponding to the created access key"
  value       = linode_object_storage_key.key.secret_key
  sensitive   = true
}
