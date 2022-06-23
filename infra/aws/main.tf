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

variable "aws_account_id" {
  type = string
}

variable "name" {
  type = string
}

variable "policy" {
  type = string
}

variable "identifier" {
  type = string
}

variable "domain" {
  type = string
}

variable "client_image_name" {
  type = string
}

variable "api_image_name" {
  type = string
}

variable "db_image_name" {
  type = string
}

variable "mysql_database" {
  type = string
}

variable "mysql_password" {
  type = string
}

variable "mysql_root_password" {
  type = string
}

variable "mysql_user" {
  type = string
}

variable "s3_user" {
  type = string
}

variable "s3_path" {
  type = string
}

variable "s3_images_bucket_name" {
  type = string
}

variable "fargate_service_name" {
  type = string
}

variable "ecs_cluster_name" {
  type = string
}

variable "ecs_task_api_container_name" {
  type = string
}

variable "ecs_task_client_container_name" {
  type = string
}

variable "ecs_task_db_container_name" {
  type = string
}

#---------------------------
#  IAM Policy
#---------------------------

data "aws_iam_policy_document" "allow_describe_regions" {
  statement {
    effect    = "Allow"
    actions   = ["ec2:DescribeRegions"]
    resources = ["*"]
  }
}

module "describe_regions_for_ec2" {
  source     = "./iam_role"
  name       = "describe_regions_for_ec2"
  identifier = "ec2.amazonaws.com"
  policy     = data.aws_iam_policy_document.allow_describe_regions.json
}

#---------------------------
#  Security Group
#---------------------------

module "terr_pres_sg" {
  source      = "./security_group"
  name        = "module-sg"
  vpc_id      = aws_vpc.terr_pres_vpc.id
  port        = 80
  cidr_blocks = ["0.0.0.0/0"]
}
