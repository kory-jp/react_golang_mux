# ---------------------------
# Terraform configuration
# ---------------------------

terraform {
  required_version = ">=1.2"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

#---------------------------
# Provider
#---------------------------

provider "aws" {
  region = "ap-northeast-1"
}

#---------------------------
#  Variable
#---------------------------

variable "name" {
  type = string
}

variable "policy" {
  type = string
}

variable "identifier" {
  type = string
}

