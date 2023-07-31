terraform {
  required_providers {
    openai = {
      source  = "speakeasy/openai"
      version = "0.0.1"
    }
  }
}

variable "key" {
  type = string
}

provider "openai" {
  api_key = var.key
}

data "openai_completion" "bumblee-bee-song" {
  model  = "text-davinci-003"
  prompt = "a poetic song about a bumblebee that likes to code:"
}

output "completion" {
  value = data.openai_completion.bumblee-bee-song.choices[0].text
}