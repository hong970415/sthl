# provider
variable "aws_region" {
  description = "aws_region for provider"
  type        = string
}
variable "aws_access_key_id" {
  description = "aws_access_key_id for provider"
  type        = string
}
variable "aws_secret_access_key" {
  description = "aws_secret_access_key for provider"
  type        = string
}

# dev
variable "dev_backend_bucket_name" {
  description = "dev s3 bucket name for remote backend state"
  type        = string
}
variable "dev_backend_db_name" {
  description = "dev aws_dynamodb_table name for remote backend state locking"
  type        = string
}

# prod ...
