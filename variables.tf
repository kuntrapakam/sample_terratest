variable "region" {
    description = "the aws region"
    type = string
    default = "us-east-1"
}

variable "tag-bucket-name" {
    description = "tag name of the resource"
    type = string
    default = "TERRATEST"
}

variable "tag-bucket-environment" {
    description = "env name of the resource"
    type = string
    default = "Dev"
}