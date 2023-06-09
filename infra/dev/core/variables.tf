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

# rdb
variable "db_instance_username" {
  description = "username for db_instance"
  type        = string
}
variable "db_instance_password" {
  description = "password for db_instance"
  type        = string
  sensitive   = true
}
