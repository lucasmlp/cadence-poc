terraform {
  required_providers {
    google = {
      source = "hashicorp/google"
      version = "3.5.0"
    }
  }
}

provider "google" {
  credentials = file("cadence-poc-6714c6805157.json")

  project = "cadence-poc"
  region  = "us-east1"
  zone    = "us-east1-b"
}

module "cadence" {
  source = "./cadence"
}