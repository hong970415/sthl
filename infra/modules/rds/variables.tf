# aws_db_subnet_group
variable "db_subnet_group_name" {
  description = "name for db_subnet_group"
  type        = string
  # default     = "sthl-dev"
}
variable "db_subnet_group_subnet_ids" {
  description = "subnet_ids for db_subnet_group"
  type        = list(string)
  # default     = ["10.0.1.0/24","10.0.2.0/24"]
}

# aws_db_parameter_group
variable "db_parameter_group_name" {
  description = "name for db_parameter_group"
  type        = string
  # default     = "sthl-dev"
}
variable "db_parameter_group_family" {
  description = "family for db_parameter_group"
  type        = string
  # default     = "postgres14"
}
# aws_db_instance
variable "db_instance_identifier" {
  description = "identifier for db_instance"
  type        = string
  # default     = "sthl-dev"
}
variable "db_instance_instance_class" {
  description = "instance_class for db_instance"
  type        = string
  # default     = "db.t3.micro"
}
variable "db_instance_allocated_storage" {
  description = "allocated_storage for db_instance"
  type        = number
  # default     = 5
}
variable "db_instance_engine" {
  description = "engine for db_instance"
  type        = string
  # default     = "postgres"
}
variable "db_instance_engine_version" {
  description = "engine_version for db_instance"
  type        = string
  # default     = "14.1"
}
variable "db_instance_username" {
  description = "username for db_instance"
  type        = string
  # default     = "postgres"
}
variable "db_instance_password" {
  description = "password for db_instance"
  type        = string
  sensitive   = true
  # default     = "postgres"
}
variable "db_instance_publicly_accessible" {
  description = "publicly_accessible for db_instance"
  type        = bool
  # default     = true
}
variable "db_instance_skip_final_snapshot" {
  description = "skip_final_snapshot for db_instance"
  type        = bool
  # default     = true
}
variable "db_instance_vpc_security_group_ids" {
  description = "vpc_security_group_ids for db_instance"
  type        = list(string)
}
