data "aws_s3_bucket" "serverless" {
  bucket = "${var.app_name}-serverless"
}

resource "aws_s3_bucket" "logs" {
  bucket = "${var.app_name}.logs"
  acl    = "private"

  lifecycle_rule {
    id      = "log"
    enabled = true

    tags = {
      Application = var.app_name
      Version     = var.app_version
      Stage       = var.app_stage
    }


    transition {
      days          = 30
      storage_class = "STANDARD_IA"
    }

    transition {
      days          = 90
      storage_class = "GLACIER"
    }

    transition {
      days          = 180
      storage_class = "DEEP_ARCHIVE"
    }

    expiration {
      days = 365
    }
  }

  tags = {
    Application = var.app_name
    Version     = var.app_version
    Stage       = var.app_stage
  }
}

resource "aws_s3_bucket" "webapp" {
  bucket = var.webapp_domain
  acl    = "public-read"

  website_domain = var.webapp_domain
  policy         = data.aws_iam_policy_document.webapp.json

  website {
    index_document = "index.html"
    error_document = "index.html"
  }

  cors_rule {
    allowed_headers = ["*"]
    allowed_methods = ["PUT", "POST", "GET"]
    allowed_origins = ["https://${var.webapp_domain}"]
    expose_headers  = ["ETag"]
    max_age_seconds = 3000
  }

  versioning {
    enabled = true
  }

  tags = {
    Application = var.app_name
    Version     = var.app_version
    Stage       = var.app_stage
  }
}
