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

# ecs image and envs
variable "ecs_task_frontend_image_url" {
  description = "image_url for ecs task frontend image"
  type        = string
}
