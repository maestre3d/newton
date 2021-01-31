resource "aws_cloudfront_distribution" "webapp" {
  enabled             = true
  is_ipv6_enabled     = true
  wait_for_deployment = true
  default_root_object = "index.html"
  price_class         = "PriceClass_All"

  aliases = [aws_s3_bucket.webapp.bucket]

  logging_config {
    bucket = aws_s3_bucket.logs.bucket_domain_name
    prefix = "webapp/${var.app_stage}"
  }

  default_cache_behavior {
    allowed_methods  = ["DELETE", "GET", "HEAD", "OPTIONS", "PATCH", "POST", "PUT"]
    cached_methods   = ["GET", "HEAD"]
    target_origin_id = aws_s3_bucket.webapp.id

    forwarded_values {
      query_string = false

      cookies {
        forward = "none"
      }
    }

    min_ttl                = 0
    default_ttl            = 86400
    max_ttl                = 31536000
    compress               = true
    viewer_protocol_policy = "redirect-to-https"
  }

  origin {
    domain_name = aws_s3_bucket.webapp.bucket_domain_name
    origin_path = "/${var.app_stage}"
    origin_id   = aws_s3_bucket.webapp.id
  }

  restrictions {
    geo_restriction {
      restriction_type = "none"
    }
  }

  viewer_certificate {
    acm_certificate_arn      = data.aws_acm_certificate.primary.arn
    minimum_protocol_version = "TLSv1.2_2019"
    ssl_support_method       = "sni-only"
  }

  custom_error_response {
    error_caching_min_ttl = 300
    error_code            = 404
    response_code         = 200
    response_page_path    = "/index.html"
  }

  depends_on = [
    aws_s3_bucket.logs,
    aws_s3_bucket.webapp
  ]

  tags = {
    Application = var.app_name
    Version     = var.app_version
    Stage       = var.app_stage
  }
}