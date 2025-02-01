output "instance_ip" {
  value = aws_instance.debian_instance.public_ip
}

output "instance_id" {
  value = aws_instance.debian_instance.id
}

output "ssh_private_key" {
  value     = tls_private_key.ssh_key.private_key_pem
  sensitive = true
}