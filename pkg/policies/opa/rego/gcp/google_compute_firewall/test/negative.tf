provider "google" {
  region = "us-west1"
}

resource "google_compute_firewall" "networkPortExposedToInternet" {
  provider      = google
  name          = "website-fw-2"
  network       = "some-network"
  source_ranges = ["192.164.0.0/28"]
  allow {
    protocol = "tcp"
    ports    = ["20", "22", "9000", "3389", "9090", "389", "1521", "2483", "6379", "7000", "7199", "8888", "9042", "9160", "9200", "9300", "11211", "27017", "61620", "135", "636", "1433", "2383", "2484", "3306", "5432", "7001", "3389", "11214", "11215", "23", "445", "25", "110", "137", "138", "139", "3000", "3020", "4505", "4506", "8000", "8080", "5500", "5900", "1434", "2382", "8140", "27018", "61621"]
  }
  allow {
    protocol = "udp"
    ports    = ["389", "2483", "11211", "2484", "5432", "11214", "11215", "137", "139", "161", "53", "1434"]
  }
  direction = "INGRESS"
}

resource "google_compute_firewall" "networkPortExposedToInternetPublic" {
  provider      = google
  name          = "website-fw-2"
  network       = "some-network"
  source_ranges = ["192.164.0.0/28"]
  allow {
    protocol = "tcp"
    ports    = ["20", "22", "9000", "3389", "9090", "389", "1521", "2483", "6379", "7000", "7199", "8888", "9042", "9160", "9200", "9300", "11211", "27017", "61620", "135", "636", "1433", "2383", "2484", "3306", "5432", "7001", "3389", "11214", "11215", "23", "445", "25", "110", "137", "138", "139", "3000", "3020", "4505", "4506", "8000", "8080", "5500", "5900", "1434", "2382", "8140", "27018", "61621"]
  }
  allow {
    protocol = "udp"
    ports    = ["389", "2483", "11211", "2484", "5432", "11214", "11215", "137", "139", "161", "53", "1434"]
  }
  direction = "INGRESS"
}

resource "google_compute_firewall" "networkPortExposedToInternetPrivate" {
  provider      = google
  name          = "website-fw-2"
  network       = "some-network"
  source_ranges = ["10.0.0.0/23"]
  allow {
    protocol = "tcp"
    ports    = ["20", "22", "9000", "3389", "9090", "389", "1521", "2483", "6379", "7000", "7199", "8888", "9042", "9160", "9200", "9300", "11211", "27017", "61620", "135", "636", "1433", "2383", "2484", "3306", "5432", "7001", "3389", "11214", "11215", "23", "445", "25", "110", "137", "138", "139", "3000", "3020", "4505", "4506", "8000", "8080", "5500", "5900", "1434", "2382", "8140", "27018", "61621"]
  }
  allow {
    protocol = "udp"
    ports    = ["389", "2483", "11211", "2484", "5432", "11214", "11215", "137", "139", "161", "53", "1434"]
  }
  direction = "INGRESS"
}

resource "google_compute_firewall" "unrestrictedRdpAccess" {
  name      = "another-firewall"
  network   = "another-network"
  direction = "INGRESS"

  allow {
    protocol = "icmp"
  }

  allow {
    protocol = "tcp"
    ports    = ["3389"]
  }

  source_tags = ["web"]
}
