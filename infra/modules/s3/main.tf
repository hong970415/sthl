# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket
resource "aws_s3_bucket" "main" {
  bucket        = var.s3_bucket_name
  force_destroy = var.s3_bucket_force_destroy
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_ownership_controls
resource "aws_s3_bucket_ownership_controls" "main" {
  bucket = aws_s3_bucket.main.id
  rule {
    object_ownership = var.s3_bucket_object_ownership
  }
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_cors_configuration
resource "aws_s3_bucket_cors_configuration" "main" {
  bucket = aws_s3_bucket.main.id

  cors_rule {
    allowed_headers = var.s3_bucket_cors_rule_a_allowed_headers
    allowed_methods = var.s3_bucket_cors_rule_a_allowed_methods
    allowed_origins = var.s3_bucket_cors_rule_a_allowed_origins
    expose_headers  = var.s3_bucket_cors_rule_a_expose_headers
    max_age_seconds = var.s3_bucket_cors_rule_a_max_age_seconds
  }
  cors_rule {
    allowed_methods = var.s3_bucket_cors_rule_b_allowed_methods
    allowed_origins = var.s3_bucket_cors_rule_b_allowed_origins
  }
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_public_access_block
resource "aws_s3_bucket_public_access_block" "main" {
  bucket = aws_s3_bucket.main.id

  block_public_acls       = var.s3_bucket_pcb_block_public_acls
  block_public_policy     = var.s3_bucket_pcb_block_public_policy
  ignore_public_acls      = var.s3_bucket_pcb_ignore_public_acls
  restrict_public_buckets = var.s3_bucket_pcb_restrict_public_buckets
}

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/s3_bucket_acl
resource "aws_s3_bucket_acl" "example" {
  depends_on = [
    aws_s3_bucket_ownership_controls.main,
    aws_s3_bucket_public_access_block.main,
  ]

  bucket = aws_s3_bucket.main.id
  acl    = var.s3_bucket_acl
}
