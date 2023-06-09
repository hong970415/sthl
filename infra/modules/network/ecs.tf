# ecr
resource "aws_ecr_repository" "frontend" {
  name         = var.ecr_frontend_repository_name
  force_delete = true
}
resource "aws_ecr_repository" "backend" {
  name         = var.ecr_backend_repository_name
  force_delete = true
}

# # https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_cluster
resource "aws_ecs_cluster" "main" {
  name = var.ecs_cluster_name

  setting {
    name  = "containerInsights"
    value = "enabled"
  }
}
