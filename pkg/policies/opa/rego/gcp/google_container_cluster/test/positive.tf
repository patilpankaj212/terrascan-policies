provider "google" {
  region = "us-west1"
}

resource "google_container_cluster" "stackDriverLoggingEnabled" {
  name            = "marcellus-wallace"
  logging_service = "none"
  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
}
resource "google_container_cluster" "stackDriverMonitoringEnabled" {
  name               = "marcellus-wallace1"
  monitoring_service = "none"
  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
}

resource "google_container_cluster" "legacyAuthEnabled" {
  name               = "marcellus-wallace2"
  enable_legacy_abac = true
  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
}

resource "google_container_cluster" "masterAuthEnabled" {
  name = "marcellus-wallace3"
  master_auth {
    # username = "some-username"
    # password = "some-password"

    client_certificate_config {
      issue_client_certificate = true
    }
  }
  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
}


resource "google_container_cluster" "clientCertificateEnabled" {
  name = "marcellus-wallace6"

  master_auth {
    username = "some-username"
    password = "some-password"

    client_certificate_config {
      issue_client_certificate = true
    }
  }

  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
}
resource "google_container_cluster" "networkPolicyEnabled" {
  name = "marcellus-wallace4"
  network_policy {
    enabled = false
  }
  ip_allocation_policy {
    cluster_secondary_range_name  = "some-name"
    services_secondary_range_name = "some-name"
  }
  enable_kubernetes_alpha = true
}

resource "google_container_cluster" "ipAliasingEnabled" {
  name = "marcellus-wallace5"
}

resource "google_container_cluster" "podSecurityPolicyEnabled" {
  name     = "marcellus-wallace5"
  provider = "google-beta"
  node_config {
    preemptible = false
    shielded_instance_config {
      enable_secure_boot = false
    }
  }
}
resource "google_container_cluster" "gkeControlPlaneNotPublic" {
  name     = "another-cluster"
  location = "us-central1"

  master_auth {
    username = ""
    password = ""
  }
  private_cluster_config {
    enable_private_endpoint = false
  }
}

resource "google_container_cluster" "gkeBasicAuthDisabled" {
  name     = "other-cluster"
  location = "us-central1"

  master_auth {
    username = ""
    password = ""
  }
}

resource "google_container_cluster" "privateClusterEnabled" {
  name     = "another-cluster"
  location = "us-central1"

  private_cluster_config {
    enable_private_endpoint = false
    enable_private_nodes    = false
  }
  addons_config {
    http_load_balancing {
      disabled = true
    }
  }
}

resource "google_container_cluster" "clusterLabelsEnabled" {
  name                  = "nother-cluster"
  location              = "us-central1"
  enable_shielded_nodes = false
  node_config {
    preemptible = false
    shielded_instance_config {
      enable_secure_boot = false
    }
    metadata = {
      "disable-legacy-endpoints" = "true"
    }

    workload_metadata_config {
      node_metadata = "EXPOSE"
    }
  }
  database_encryption {
    key_name = "id_rsa.pub"
    state    = "DECRYPTED"
  }
  addons_config {
    horizontal_pod_autoscaling {
      disabled = true
    }
  }
}
