output "tfbackend_dev_s3_bucket_name" {
  description = "The name of the S3 bucket for dev remote tfstate"
  value       = module.tfbackend_dev.s3_bucket_name
}

output "tfbackend_dev_dynamodb_table_name" {
  description = "The name of the DynamoDB table for dev remote tfstate"
  value       = module.tfbackend_dev.dynamodb_table_name
}
