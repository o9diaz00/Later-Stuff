resource "google_storage_bucket" "bucket" {
  name          = var.name
  project       = var.project_id
  storage_class = var.storage_class

  lifecycle_rule {
    condition {
      age = 1
    }
    action {
      type = "Delete"
    }
  }
}