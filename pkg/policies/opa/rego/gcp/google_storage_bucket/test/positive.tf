provider "google" {
  region = "us-west1"
}

resource "google_storage_bucket" "checkStorageBucketConfig" {
  name          = "auto-expiring-bucket"
  location      = "US"
  force_destroy = true

  lifecycle_rule {
    condition {
      age = "3"
    }
    action {
      type = "Delete"
    }
  }

  # logging {
  #   log_bucket = "sample-log-bucket"
  # }

  versioning {
    enabled = false
  }
}

resource "google_storage_bucket" "uniformBucketEnabled" {
  name          = "auto-bucket"
  location      = "US"
  force_destroy = true

  lifecycle_rule {
    condition {
      age = "3"
    }
    action {
      type = "Delete"
    }
  }
  retention_policy {
    retention_period = 776000
  }
  bucket_policy_only = false
}
