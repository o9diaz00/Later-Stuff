module "terraform-state-bucket" {
  source = "../templates/google_storage_bucket/"

  name             = "terraform-state-bucket"
  project_id       = "brave-aileron-450414-e0"
  storage_class    = "STANDARD"
  lifecycle_age    = 1
  lifecycle_action = "Delete"
  versioning       = true
}
