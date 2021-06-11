provider "google" {
  region      = "us-west1"
}

resource "google_kms_key_ring" "keyring" {
  name     = "keyring-example"
  location = "global"
}

resource "google_kms_crypto_key" "checkRotation365Days" {
  name            = "crypto-key-example"
  key_ring        = google_kms_key_ring.keyring.self_link
  rotation_period = "8775000s"

  lifecycle {
    prevent_destroy = true
  }
}

resource "google_kms_crypto_key" "checkRotation90Days" {
  name            = "crypto-key-sample"
  key_ring        = google_kms_key_ring.keyring.self_link
  rotation_period = "8775000s"

  lifecycle {
    prevent_destroy = true
  }
}
