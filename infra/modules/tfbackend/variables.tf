# backend
variable "backend_bucket_name" {
  description = "s3 bucket name for remote backend state"
  type        = string
}
variable "backend_db_name" {
  description = "aws_dynamodb_table name for remote backend state locking"
  type        = string
}
