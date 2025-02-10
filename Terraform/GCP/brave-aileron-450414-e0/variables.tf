variable "name" {
  description = "Name of the resource file"
  type        = string
}

variable "project" {
  description = "ID of current gcloud project"
  type        = string
  default     = "brave-aileron-450414-e0"
}

variable "storage_class" {
  description = "Storage class setting of bucket (use STANDARD for free)"
  type        = string
  default     = "STANDARD"
}