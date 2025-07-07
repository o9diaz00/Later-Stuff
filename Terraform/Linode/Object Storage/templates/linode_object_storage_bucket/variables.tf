variable "label" {
  description = "The label of the Linode Object Storage Bucket."
  type        = string
}

variable "location" {
  description = "Map containing region : s3-endpoint."
  type        = map(string)
  default     = null
}

variable "region" {
  description = "The region of the Linode Object Storage Bucket.  Exactly one of region and cluster is required for creating a bucket."
  type        = string
  default     = ""
}

variable "s3_endpoint" {
  description = "The user's s3 endpoint URL, based on the endpoint_type and region."
  type        = string
  default     = ""
}

variable "abort_incomplete_multipart_upload_days" {
  description = "Specifies the number of days after initiating a multipart upload when the multipart upload must be completed."
  type        = number
  default     = 1
}

variable "noncurrent_version_expiration" {
  description = "Specifies when non-current object versions expire."
  type        = number
  default     = 1
}

variable "expiration_days" {
  description = "Specifies a period in the object's expire."
  type        = number
}
