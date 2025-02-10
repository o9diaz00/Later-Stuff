resource "google_storage_bucket" "bucket" {
  name          = "${var.project_id}_${var.name}"
  location      = var.location
  project       = var.project_id
  storage_class = var.storage_class
  force_destroy = var.force_destroy

  lifecycle_rule {
    condition {
      age = var.lifecycle_age
    }
    action {
      type = var.lifecycle_action
    }
  }

  versioning {
    enabled = var.versioning
  }
}
