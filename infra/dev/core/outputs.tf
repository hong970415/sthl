# subnet
output "aws_subnet_public_a_id" {
  description = "ID of the subnet_public_a"
  value       = module.network.aws_subnet_public_a_id
}
output "aws_subnet_public_b_id" {
  description = "ID of the subnet_public_b"
  value       = module.network.aws_subnet_public_b_id
}

# security groups 
output "aws_security_group_lb_id" {
  description = "ID of the security group lb"
  value       = module.network.aws_security_group_lb_id
}
output "aws_security_group_apps_id" {
  description = "ID of the security group apps"
  value       = module.network.aws_security_group_apps_id
}
output "aws_security_group_db_id" {
  description = "ID of the security group db"
  value       = module.network.aws_security_group_db_id
}

# iam
output "aws_iam_role_arn" {
  description = "arn of aws_iam_role"
  value       = module.network.aws_iam_role_arn
}
output "aws_iam_role_policy_attachment_id" {
  description = "id of aws_iam_role_policy_attachment"
  value       = module.network.aws_iam_role_policy_attachment_id
}

# lb
output "alb_dns_name" {
  value       = module.network.alb_dns_name
  description = "The domain name of the load balancer"
}
output "aws_lb_target_group_frontend_arn" {
  description = "arn of aws_lb_target_group frontend"
  value       = module.network.aws_lb_target_group_frontend_arn
}
output "aws_lb_target_group_backend_arn" {
  description = "arn of aws_lb_target_group backend"
  value       = module.network.aws_lb_target_group_backend_arn
}

# ecr
output "aws_ecr_repository_frontend_id" {
  description = "id of the aws_ecr_repository frontend"
  value       = module.network.aws_ecr_repository_frontend_id
}
output "aws_ecr_repository_frontend_repository_url" {
  description = "repository_url of the aws_ecr_repository frontend"
  value       = module.network.aws_ecr_repository_frontend_repository_url
}
output "aws_ecr_repository_backend_id" {
  description = "id of the aws_ecr_repository backend"
  value       = module.network.aws_ecr_repository_backend_id
}
output "aws_ecr_repository_backend_repository_url" {
  description = "repository_url of the aws_ecr_repository backend"
  value       = module.network.aws_ecr_repository_backend_repository_url
}
output "aws_ecs_cluster_id" {
  description = "id of aws_ecs_cluster"
  value       = module.network.aws_ecs_cluster_id
}


# #### s3
output "s3_bucket_name" {
  description = "The name of the bucket"
  value       = module.s3.s3_bucket_name
}
output "s3_bucket_domain_name" {
  description = "The bucket_domain_name of the bucket"
  value       = module.s3.s3_bucket_domain_name
}

# #### rds
output "rds_address" {
  description = "RDS instance address"
  value       = module.rds.rds_address
}
output "rds_endpoint" {
  description = "RDS instance endpoint"
  value       = module.rds.rds_endpoint
}
output "rds_port" {
  description = "RDS instance port"
  value       = module.rds.rds_port
}
output "rds_username" {
  description = "RDS instance root username"
  value       = module.rds.rds_username
}
output "rds_status" {
  description = "RDS instance status"
  value       = module.rds.rds_status
}

