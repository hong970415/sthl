terraform {
  required_version = ">= 1.4.6"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }

  backend "local" {
    path = "terraform.tfstate"
  }
}


module "tfbackend_dev" {
  source = "../modules/tfbackend"

  backend_bucket_name = var.dev_backend_bucket_name
  backend_db_name     = var.dev_backend_db_name
}
