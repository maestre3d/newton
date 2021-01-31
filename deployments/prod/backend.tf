terraform {
  backend "s3" {
    bucket = "newton-serverless"
    key    = "prod/tf/terraform.tfstate"
    region = "us-east-1"
  }
}