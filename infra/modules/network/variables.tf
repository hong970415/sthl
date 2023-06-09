# VPC
variable "vpc_cidr" {
  description = "CIDR block for VPC"
  type        = string
  # default     = "10.0.0.0/16"
}
variable "vpc_enable_dns_hostnames" {
  description = "Enable dns hostnames for VPC"
  type        = bool
  # default     = true
}

# Subnet
variable "subnet_public_a_cidr" {
  description = "CIDR block for Subnet public_a"
  type        = string
  # default     = "10.0.1.0/24"
}
variable "subnet_public_a_az" {
  description = "Availability zone for Subnet public_a"
  type        = string
  # default     = "ap-east-1a"
}
variable "subnet_public_a_map_public_ip_on_launch" {
  description = "Map public ip on launch for Subnet public_a"
  type        = bool
  # default     = true
}
variable "subnet_public_b_cidr" {
  description = "CIDR block for Subnet public_b"
  type        = string
  # default     = "10.0.2.0/24"
}
variable "subnet_public_b_az" {
  description = "Availability zone for Subnet public_b"
  type        = string
  # default     = "ap-east-1b"
}
variable "subnet_public_b_map_public_ip_on_launch" {
  description = "Map public ip on launch for Subnet public_b"
  type        = bool
  # default     = true
}

# iam
variable "iam_role_name" {
  description = "name for iam role"
  type        = string
  # default     = "sthl-role-dev"
}

# lb
variable "lb_name" {
  description = "name for alb"
  type        = string
  # default     = "sthl-alb-dev"
}
variable "lb_load_balancer_type" {
  description = "load_balancer_type for alb"
  type        = string
  # default     = "application"
}
variable "lb_frontend_target_group_name" {
  description = "name for alb frontend target_group"
  type        = string
  # default     = "frontend-target-group-dev"
}
variable "lb_backend_target_group_name" {
  description = "name for alb backend target_group"
  type        = string
  # default     = "backend-target-group-dev"
}

# acm
variable "domain_name" {
  description = "owned domain name"
  type        = string
  # default     = "sthll.com"
}
variable "target_domain_name" {
  description = "specific target domain name to add to s3record"
  type        = string
  # default     = "dev.example.com"
}

# ecs
variable "ecr_frontend_repository_name" {
  description = "name for ecr frontend repository"
  type        = string
  # default     = "sthl-frontend-dev"
}
variable "ecr_backend_repository_name" {
  description = "name for ecr backend repository"
  type        = string
  # default     = "sthl-backend-dev"
}
variable "ecs_cluster_name" {
  description = "name for ecs cluster"
  type        = string
  # default     = "sthl-ecs-dev"
}
