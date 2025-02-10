variable "name" {
  description = "The name of the bucket."
  type        = string
}

variable "location" {
  description = "The GCS location"
  type        = string
  default     = "US"
}

variable "project_id" {
  description = "The ID of the project in which the resource belongs. If it is not provided, the provider project is used."
  type        = string
  default     = "brave-aileron-450414-e0"
}

variable "storage_class" {
  description = " (Default: 'STANDARD' --- FREE) The Storage Class of the new bucket. Supported values include: STANDARD, MULTI_REGIONAL, REGIONAL, NEARLINE, COLDLINE, ARCHIVE."
  type        = string
  default     = "STANDARD"
}

variable "lifecycle_age" {
  description = "Minimum age of an object in days to satisfy this condition."
  type        = number
  default     = 1
}

variable "lifecycle_action" {
  description = "The Lifecycle Rule's action configuration. A single block of this type is supported."
  type        = string
  default     = "Delete"
}

variable "versioning" {
  description = "Set versioning for this bucket."
  type        = bool
  default     = false
}

variable "force_destroy" {
  description = "When deleting a bucket, this boolean option will delete all contained objects. If you try to delete a bucket that contains objects, Terraform will fail that run."
  type        = bool
  default     = true
}

