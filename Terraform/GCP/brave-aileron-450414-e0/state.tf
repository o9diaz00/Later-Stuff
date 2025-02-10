terraform {
  backend "gcs" {
    bucket = "brave-aileron-450414-e0_terraform-state-bucket"
    prefix = "storage/state"
  }
}