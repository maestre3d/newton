data "aws_iam_policy_document" "webapp" {
  statement {
    sid    = "PublicRead"
    effect = "Allow"
    principals {
      identifiers = ["*"]
      type        = "AWS"
    }
    actions   = ["s3:GetObject"]
    resources = ["arn:aws:s3:::${var.webapp_domain}/*"]
  }
}