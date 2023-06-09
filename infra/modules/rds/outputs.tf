# rds
output "rds_address" {
  description = "RDS instance address"
  value       = aws_db_instance.main.address
}
output "rds_endpoint" {
  description = "RDS instance endpoint"
  value       = aws_db_instance.main.endpoint
}
output "rds_port" {
  description = "RDS instance port"
  value       = aws_db_instance.main.port
}
output "rds_username" {
  description = "RDS instance root username"
  value       = aws_db_instance.main.username
}
output "rds_status" {
  description = "RDS instance status"
  value       = aws_db_instance.main.status
}
