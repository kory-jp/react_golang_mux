#--------------------------
# IAM User for S3
#--------------------------

resource "aws_iam_user" "for_s3" {
  name          = var.s3_user
  path          = var.s3_path
  force_destroy = true

  tags = {
    tag-key = "terr_pres_user"
  }
}

resource "aws_iam_user_policy" "s3_policy" {
  name = "s3_policy"
  user = aws_iam_user.for_s3.name

  policy = <<EOF
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Effect": "Allow",
            "Action": [
                "s3:*",
                "s3-object-lambda:*"
            ],
            "Resource": "*"
        }
    ]
}
EOF
}

resource "aws_iam_access_key" "s3_access_key" {
  user = aws_iam_user.for_s3.name
}