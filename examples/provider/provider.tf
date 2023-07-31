terraform {
  required_providers {
    openai = {
      source  = "speakeasy/openai"
      version = "0.0.1"
    }
  }
}

provider "openai" {
  # Configuration options
}