variable "s3_bucket_name" {
  description = "s3 bucket name"
  type        = string
  # default     = "sthl-dev"
}
variable "s3_bucket_force_destroy" {
  description = "s3 force destroy"
  type        = bool
  # default     = true
}
variable "s3_bucket_object_ownership" {
  description = "s3 bucket object ownership"
  type        = string
  # default     = "BucketOwnerPreferred"
}
variable "s3_bucket_cors_rule_a_allowed_headers" {
  description = "s3 bucket cors rule a allowed headers associate with backend app"
  type        = list(string)
  # default     = ["*"]
}
variable "s3_bucket_cors_rule_a_allowed_methods" {
  description = "s3 bucket cors rule a allowed methods associate with backend app"
  type        = list(string)
  # default     = ["PUT", "POST", "DELETE"]
}
variable "s3_bucket_cors_rule_a_allowed_origins" {
  description = "s3 bucket cors rule a allowed origins associate with backend app"
  type        = list(string)
  # default     = ["http://localhost:4000"]
}
variable "s3_bucket_cors_rule_a_expose_headers" {
  description = "s3 bucket cors rule a expose headers associate with backend app"
  type        = list(string)
  # default     = ["ETag"]
}
variable "s3_bucket_cors_rule_a_max_age_seconds" {
  description = "s3 bucket cors rule a max age seconds associate with backend app"
  type        = number
  # default     = 3000
}
variable "s3_bucket_cors_rule_b_allowed_methods" {
  description = "s3 bucket cors rule b allowed methods associate with frontend client"
  type        = list(string)
  # default     = ["GET"]
}
variable "s3_bucket_cors_rule_b_allowed_origins" {
  description = "s3 bucket cors rule b allowed origins associate with frontend client"
  type        = list(string)
  # default     = ["*"]
}
variable "s3_bucket_pcb_block_public_acls" {
  description = "s3_bucket_pcb_block_public_acls for s3 bucket"
  type        = bool
  # default     = false
}
variable "s3_bucket_pcb_block_public_policy" {
  description = "s3_bucket_pcb_block_public_policy for s3 bucket"
  type        = bool
  # default     = false
}
variable "s3_bucket_pcb_ignore_public_acls" {
  description = "s3_bucket_pcb_ignore_public_acls for s3 bucket"
  type        = bool
  # default     = false
}
variable "s3_bucket_pcb_restrict_public_buckets" {
  description = "s3_bucket_pcb_restrict_public_buckets for s3 bucket"
  type        = bool
  # default     = false
}
variable "s3_bucket_acl" {
  description = "s3 bucket acl"
  type        = string
  # default     = "public-read"
}
