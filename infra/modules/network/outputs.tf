# subnet
output "aws_subnet_public_a_id" {
  description = "ID of the subnet_public_a"
  value       = aws_subnet.public_a.id
}
output "aws_subnet_public_b_id" {
  description = "ID of the subnet_public_b"
  value       = aws_subnet.public_b.id
}


# security groups 
output "aws_security_group_lb_id" {
  description = "ID of the security group lb"
  value       = aws_security_group.lb.id
}
output "aws_security_group_apps_id" {
  description = "ID of the security group apps"
  value       = aws_security_group.apps.id
}
output "aws_security_group_db_id" {
  description = "ID of the security group db"
  value       = aws_security_group.db.id
}

# iam
output "aws_iam_role_arn" {
  description = "arn of aws_iam_role"
  value       = aws_iam_role.main.arn
}
output "aws_iam_role_policy_attachment_id" {
  description = "id of aws_iam_role_policy_attachment"
  value       = aws_iam_role_policy_attachment.attach.id
}

# lb
output "alb_dns_name" {
  description = "The domain name of the load balancer"
  value       = aws_lb.main.dns_name
}
output "aws_lb_target_group_frontend_arn" {
  description = "arn of aws_lb_target_group frontend"
  value       = aws_lb_target_group.frontend.arn
}
output "aws_lb_target_group_backend_arn" {
  description = "arn of aws_lb_target_group backend"
  value       = aws_lb_target_group.backend.arn
}

# ecr
output "aws_ecr_repository_frontend_id" {
  description = "id of the aws_ecr_repository frontend"
  value       = aws_ecr_repository.frontend.id
}
output "aws_ecr_repository_frontend_repository_url" {
  description = "repository_url of the aws_ecr_repository frontend"
  value       = aws_ecr_repository.frontend.repository_url
}
output "aws_ecr_repository_backend_id" {
  description = "id of the aws_ecr_repository backend"
  value       = aws_ecr_repository.backend.id
}
output "aws_ecr_repository_backend_repository_url" {
  description = "repository_url of the aws_ecr_repository backend"
  value       = aws_ecr_repository.backend.repository_url
}
output "aws_ecs_cluster_id" {
  description = "id of aws_ecs_cluster"
  value       = aws_ecs_cluster.main.id
}
