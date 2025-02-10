module "terraform-state-bucket" {
  source = "git::https://github.com/o9diaz00/Later-Stuff.git//Terraform/GCP/templates/google_storage_bucket?ref=main"
  name             = "terraform-state-bucket"
  project_id       = "brave-aileron-450414-e0"
  storage_class    = "STANDARD"
  lifecycle_age    = 1
  lifecycle_action = "Delete"
  versioning       = true
}
