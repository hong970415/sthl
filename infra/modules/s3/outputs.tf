output "s3_bucket_name" {
  description = "The name of the bucket"
  value       = aws_s3_bucket.main.bucket
}
output "s3_bucket_domain_name" {
  description = "The bucket_domain_name of the bucket"
  value       = aws_s3_bucket.main.bucket_domain_name
}
