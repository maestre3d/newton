data "aws_route53_zone" "primary" {
  name         = "damascus-engineering.com"
  private_zone = false
}

resource "aws_route53_record" "webapp" {
  zone_id = data.aws_route53_zone.primary.zone_id
  name    = var.webapp_domain
  type    = "A"

  alias {
    evaluate_target_health = false
    name                   = aws_cloudfront_distribution.webapp.domain_name
    zone_id                = aws_cloudfront_distribution.webapp.hosted_zone_id
  }

  depends_on = [
    aws_cloudfront_distribution.webapp
  ]
}