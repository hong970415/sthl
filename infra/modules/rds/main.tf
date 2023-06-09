# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_subnet_group
resource "aws_db_subnet_group" "main" {
  name       = var.db_subnet_group_name
  subnet_ids = var.db_subnet_group_subnet_ids
}
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_parameter_group
resource "aws_db_parameter_group" "main" {
  name   = var.db_parameter_group_name
  family = var.db_parameter_group_family

  parameter {
    name  = "log_connections"
    value = "1"
  }
}
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/db_instance
resource "aws_db_instance" "main" {
  identifier             = var.db_instance_identifier
  instance_class         = var.db_instance_instance_class
  allocated_storage      = var.db_instance_allocated_storage
  engine                 = var.db_instance_engine
  engine_version         = var.db_instance_engine_version
  username               = var.db_instance_username
  password               = var.db_instance_password
  publicly_accessible    = var.db_instance_publicly_accessible
  skip_final_snapshot    = var.db_instance_skip_final_snapshot
  vpc_security_group_ids = var.db_instance_vpc_security_group_ids
  db_subnet_group_name   = aws_db_subnet_group.main.name
  parameter_group_name   = aws_db_parameter_group.main.name
}
