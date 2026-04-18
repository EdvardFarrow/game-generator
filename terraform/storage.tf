resource "google_storage_bucket" "bronze_layer" {
  name          = "${var.project_id}-game-analytics-bronze" 
  location      = "EU"
  force_destroy = true

  uniform_bucket_level_access = true

  lifecycle_rule {
    condition {
      age = 7
    }
    action {
      type = "Delete"
    }
  }
}

data "google_storage_project_service_account" "gcs_account" {}

resource "google_storage_bucket_iam_member" "pubsub_bucket_reader" {
  bucket = google_storage_bucket.bronze_layer.name
  role   = "roles/storage.legacyBucketReader"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

resource "google_storage_bucket_iam_member" "pubsub_object_creator" {
  bucket = google_storage_bucket.bronze_layer.name
  role   = "roles/storage.objectCreator"
  member = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-pubsub.iam.gserviceaccount.com"
}

data "google_project" "project" {}