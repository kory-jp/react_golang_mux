# ------------------------------
# S3 for ALB_log
# ------------------------------

resource "aws_s3_bucket" "alb_log" {
  bucket = "alb-log-pragmatic-terraform-presiden-todo"
  # force_destroy = true
}

resource "aws_s3_bucket_lifecycle_configuration" "bucket-config" {
  bucket = aws_s3_bucket.alb_log.id

  rule {
    id     = "log"
    status = "Enabled"
    expiration {
      days = "90"
    }
  }
}

resource "aws_s3_bucket_policy" "alb_log" {
  bucket = aws_s3_bucket.alb_log.id
  policy = data.aws_iam_policy_document.alb_log.json
}

data "aws_iam_policy_document" "alb_log" {
  statement {
    effect    = "Allow"
    actions   = ["s3:PutObject"]
    resources = ["arn:aws:s3:::${aws_s3_bucket.alb_log.id}/*"]

    principals {
      type        = "AWS"
      identifiers = ["582318560864"]
    }
  }
}

# ------------------------------
# S3 for Images
# ------------------------------

resource "aws_s3_bucket" "for_images" {
  bucket = var.s3_images_bucket_name
  # force_destroy = true
}

resource "aws_s3_bucket_public_access_block" "for_images" {
  bucket = aws_s3_bucket.for_images.id

  block_public_acls       = false
  block_public_policy     = false
  ignore_public_acls      = false
  restrict_public_buckets = false
}

data "aws_iam_policy_document" "allow_access_from_s3_user" {
  version = "2012-10-17"

  statement {
    sid     = "statement1"
    actions = ["s3:GetObject"]
    effect  = "Allow"
    principals {
      type        = "*"
      identifiers = ["*"]
    }
    resources = [
      "arn:aws:s3:::${var.s3_images_bucket_name}/*"
    ]
  }

  statement {
    sid     = "2"
    actions = ["s3:GetObject"]
    effect  = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::${var.aws_account_id}:user${var.s3_path}${var.s3_user}"]
    }
    resources = [
      "arn:aws:s3:::${var.s3_images_bucket_name}/*"
    ]
  }
}

resource "aws_s3_bucket_policy" "for_images" {
  bucket = aws_s3_bucket.for_images.id
  policy = data.aws_iam_policy_document.allow_access_from_s3_user.json
}