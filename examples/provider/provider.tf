terraform {
  required_version = ">= 0.13"
  required_providers {
    harvester = {
      source  = "harvester/harvester"
      version = "0.5.4"
    }
  }
}

provider "harvester" {
}
