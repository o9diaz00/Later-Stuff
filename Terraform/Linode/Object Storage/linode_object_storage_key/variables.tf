variable "label" {
  description = "The label given to this key.  For display purposes only."
  type        = string
}

variable "bucket_access" {
  description = "Map string containing 'bucket: region'.  (Must be set if 'bucket_name' and 'permissions' is not set.  For permissions, use '.' to inherit whatever parent permissions are set to)"
  type        = list(object({
    bucket_name = string
    region      = string
    permissions = string }))
  default = null
}

variable "region" {
  description = "The region where the bucket resides."
  type        = string
  default     = ""
}

variable "permissions" {
  description = "This Limited Access Key's permissions for the selected bucket.  (Must be set if 'bucket_access' is not set)"
  type        = string
  default     = "read_only"
}

variable "bucket_name" {
  description = "The unique label of the bucket to which the key will grant limited access.  (Must be set if 'bucket_access' is not set)"
  type        = string
  default     = ""
}
