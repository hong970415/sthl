# when destroy this module, 
# 1. set ecs_service_desired_count to 0 then terraform apply
# 2. terraform destroy

# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_task_definition
resource "aws_ecs_task_definition" "main" {
  family                   = var.ecs_task_definition_family
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = var.ecs_task_cpu
  memory                   = var.ecs_task_memory
  execution_role_arn       = var.iam_role_arn
  container_definitions    = var.ecs_task_container_definitions
}
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/ecs_service
resource "aws_ecs_service" "main" {
  name                               = var.ecs_service_name
  cluster                            = var.ecs_cluster_id
  task_definition                    = aws_ecs_task_definition.main.arn
  desired_count                      = var.ecs_service_desired_count
  launch_type                        = "FARGATE"
  deployment_minimum_healthy_percent = 0
  deployment_maximum_percent         = 100
  wait_for_steady_state              = true
  network_configuration {
    assign_public_ip = true
    security_groups  = var.ecs_service_security_groups
    subnets          = var.ecs_service_subnets
  }

  load_balancer {
    target_group_arn = var.ecs_service_lb_target_group_arn
    container_name   = var.ecs_service_lb_container_name
    container_port   = var.ecs_service_lb_container_port
  }
}
