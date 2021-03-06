variable "name" {}
variable "vpc_id" {}
variable "port" {}
variable "cidr_blocks" {
  type = list(string)
}

resource "aws_security_group" "terr_pres_sg" {
  name = var.name
  vpc_id = var.vpc_id

  tags = {
    "Name" = "terr_pres_sg"
  }
}

resource "aws_security_group_rule" "ingress" {
  type = "ingress"
  from_port = var.port
  to_port = var.port
  protocol = "tcp"
  cidr_blocks = var.cidr_blocks
  security_group_id = aws_security_group.terr_pres_sg.id
}

resource "aws_security_group_rule" "egress" {
  type = "egress"
  from_port = 0
  to_port = 0
  protocol = "-1"
  cidr_blocks = ["0.0.0.0/0"]
  security_group_id = aws_security_group.terr_pres_sg.id
}

output "security_group_id" {
  value = aws_security_group.terr_pres_sg.id
}