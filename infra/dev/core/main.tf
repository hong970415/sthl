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
    key            = "core/terraform.tfstate"
  }
}

module "network" {
  source = "../../modules/network"

  vpc_cidr                 = "10.0.0.0/16"
  vpc_enable_dns_hostnames = true

  subnet_public_a_cidr                    = "10.0.1.0/24"
  subnet_public_a_az                      = "ap-east-1a"
  subnet_public_a_map_public_ip_on_launch = true
  subnet_public_b_cidr                    = "10.0.2.0/24"
  subnet_public_b_az                      = "ap-east-1b"
  subnet_public_b_map_public_ip_on_launch = true

  iam_role_name                 = "sthl-role-dev"
  lb_name                       = "sthl-alb-dev"
  lb_load_balancer_type         = "application"
  lb_frontend_target_group_name = "frontend-target-group-dev"
  lb_backend_target_group_name  = "backend-target-group-dev"

  domain_name        = "sthll.com"
  target_domain_name = "dev.sthll.com"

  ecr_frontend_repository_name = "sthl-frontend-dev"
  ecr_backend_repository_name  = "sthl-backend-dev"
  ecs_cluster_name             = "sthl-ecs-dev"
}

module "s3" {
  depends_on = [
    module.network.alb_dns_name
  ]
  source                                = "../../modules/s3"
  s3_bucket_name                        = "sthl-dev"
  s3_bucket_force_destroy               = true
  s3_bucket_object_ownership            = "BucketOwnerPreferred"
  s3_bucket_cors_rule_a_allowed_headers = ["*"]
  s3_bucket_cors_rule_a_allowed_methods = ["PUT", "POST", "DELETE"]
  s3_bucket_cors_rule_a_allowed_origins = concat(["http://localhost:4000"], [module.network.alb_dns_name])
  s3_bucket_cors_rule_a_expose_headers  = ["ETag"]
  s3_bucket_cors_rule_a_max_age_seconds = 3000
  s3_bucket_cors_rule_b_allowed_methods = ["GET"]
  s3_bucket_cors_rule_b_allowed_origins = ["*"]
  s3_bucket_pcb_block_public_acls       = false
  s3_bucket_pcb_block_public_policy     = false
  s3_bucket_pcb_ignore_public_acls      = false
  s3_bucket_pcb_restrict_public_buckets = false
  s3_bucket_acl                         = "public-read"
}

module "rds" {
  source = "../../modules/rds"
  depends_on = [
    module.network.aws_subnet_public_a_id,
    module.network.aws_subnet_public_b_id,
    module.network.aws_security_group_db_id,
  ]

  db_subnet_group_name               = "sthl-dev"
  db_subnet_group_subnet_ids         = [module.network.aws_subnet_public_a_id, module.network.aws_subnet_public_b_id]
  db_parameter_group_name            = "sthl-dev"
  db_parameter_group_family          = "postgres14"
  db_instance_identifier             = "sthl-dev"
  db_instance_instance_class         = "db.t3.micro"
  db_instance_allocated_storage      = 5
  db_instance_engine                 = "postgres"
  db_instance_engine_version         = "14.3"
  db_instance_username               = var.db_instance_username
  db_instance_password               = var.db_instance_password
  db_instance_publicly_accessible    = true
  db_instance_skip_final_snapshot    = true
  db_instance_vpc_security_group_ids = [module.network.aws_security_group_db_id]
}
