variable "ecs_task_definition_family" {
  description = "family for ecs task definition"
  type        = string
  # default     = "frontend-family-dev"
}
variable "ecs_task_cpu" {
  description = "cpu for ecs task definition"
  type        = string
  # default     = "256"
}
variable "ecs_task_memory" {
  description = "memory for ecs task definition"
  type        = string
  # default     = "512"
}
variable "iam_role_arn" {
  description = "arn for iam_role"
  type        = string
}
variable "ecs_task_container_definitions" {
  description = "container_definitions for ecs_task"
  type        = string
}
variable "ecs_service_name" {
  description = "name for ecs service"
  type        = string
  # default     = "frontend-service-dev"
}
variable "ecs_cluster_id" {
  description = "id for ecs cluster"
  type        = string
}
variable "ecs_service_desired_count" {
  description = "desired_count for ecs service"
  type        = number
}
variable "ecs_service_security_groups" {
  description = "security_groups id for ecs service"
  type        = list(string)
}
variable "ecs_service_subnets" {
  description = "subnets id for ecs service"
  type        = list(string)
}
variable "ecs_service_lb_target_group_arn" {
  description = "load_balancer target_group_arn for ecs service"
  type        = string
}
variable "ecs_service_lb_container_name" {
  description = "load_balancer container_port for ecs service"
  type        = string
}
variable "ecs_service_lb_container_port" {
  description = "load_balancer target_group_arn for ecs service"
  type        = number
}
