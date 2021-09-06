terraform {
  backend "gcs" {
    bucket = "iac-cadence-playground"
    credentials = "cadence-playground.json"
  }
}

provider "kubernetes" {
  host = "http://127.0.0.1:9874"
}

module "cassandra" {
  source = "./cadence"
}