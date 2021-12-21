provider "aws" {
  region = var.region
}

terraform {
  required_version = ">= 0.12.26"
}



resource "aws_s3_bucket" "terratest-bucket" {
  bucket = "terratestvalidationnttdata"
  acl    = "private"

  versioning {
    enabled = true
  }


  tags = {
    Name        = var.tag-bucket-name
    Environment = var.tag-bucket-environment
  }
}


data "aws_caller_identity" "current" {
}

locals {
  aws_account_id = data.aws_caller_identity.current.account_id
}