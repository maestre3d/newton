data "aws_acm_certificate" "primary" {
  domain      = "damascus-engineering.com"
  most_recent = true
  statuses    = ["ISSUED"]
}