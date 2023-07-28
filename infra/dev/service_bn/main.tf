terraform {
  required_version = ">= 1.4.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  # backend "local" {
  #   path = "terraform.tfstate"
  # }

  backend "s3" {
    bucket         = "sthl-tfstate-dev"
    dynamodb_table = "sthl-tfstate-dev"
    region         = "ap-east-1"
    encrypt        = true
    key            = "service_bn/terraform.tfstate"
  }
}

data "terraform_remote_state" "core" {
  backend = "s3"
  config = {
    bucket         = "sthl-tfstate-dev"
    dynamodb_table = "sthl-tfstate-dev"
    region         = "ap-east-1"
    encrypt        = true
    key            = "core/terraform.tfstate"
  }
}

module "backend" {
  source                     = "../../modules/ecs_task"
  ecs_task_definition_family = "backend-family-dev"
  ecs_task_cpu               = 256
  ecs_task_memory            = 512
  iam_role_arn               = data.terraform_remote_state.core.outputs.aws_iam_role_arn
  ecs_task_container_definitions = jsonencode([{
    name      = "backend"
    image     = var.ecs_task_backend_image_url
    cpu       = 256
    memory    = 512
    essential = true
    portMappings = [
      {
        containerPort = 4000
      }
    ]
    "environment" : var.ecs_task_backend_image_envs
  }])
  ecs_service_name                = "backend-service-dev"
  ecs_cluster_id                  = data.terraform_remote_state.core.outputs.aws_ecs_cluster_id
  ecs_service_desired_count       = 0
  ecs_service_security_groups     = [data.terraform_remote_state.core.outputs.aws_security_group_apps_id]
  ecs_service_subnets             = [data.terraform_remote_state.core.outputs.aws_subnet_public_a_id, data.terraform_remote_state.core.outputs.aws_subnet_public_b_id]
  ecs_service_lb_target_group_arn = data.terraform_remote_state.core.outputs.aws_lb_target_group_backend_arn
  ecs_service_lb_container_name   = "backend"
  ecs_service_lb_container_port   = 4000
}

