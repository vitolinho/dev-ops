data "aws_vpc" "default" {
  default = true
}

module "ec2" {
  instance_name   = "germany-instance"
  source          = "../../../ec2-instance"
  ami             = "ami-0584590e5f0e97daa"
  instance_type   = "t2.micro"
  environment     = "dev"
  allowed_ports   = [22, 80]
  region          = "eu-central-1"
}

output "instance_ip" {
  value = module.ec2.instance_ip
}

output "ssh_private_key" {
  value     = module.ec2.ssh_private_key
  sensitive = true
}