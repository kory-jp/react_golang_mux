# #-----------------------------
# #  RDS db_parameter
# #-----------------------------

# resource "aws_db_parameter_group" "db_parameter" {
#   name   = "terr-pres-mysql-parameter"
#   family = "mysql8.0"

#   parameter {
#     name  = "character_set_database"
#     value = "utf8mb4"
#   }

#   parameter {
#     name  = "character_set_server"
#     value = "utf8mb4"
#   }
# }

# #-----------------------------
# #  RDS db_option
# #-----------------------------

# resource "aws_db_option_group" "db_option" {
#   name                 = "terr-pres-mysql-db-option"
#   engine_name          = "mysql"
#   major_engine_version = "8.0"
# }

# #-----------------------------
# #  RDS subnet_group
# #-----------------------------

# resource "aws_db_subnet_group" "subnet_group" {
#   name = "terr-pres-mysql-subnet-group"
#   subnet_ids = [
#     aws_subnet.terr_pres_private_subnet_1a.id,
#     aws_subnet.terr_pres_private_subnet_1c.id
#   ]

#   tags = {
#     "Name" = "terr-pres-mysql-subnet-group"
#   }
# }

# #-----------------------------
# #  RDS instance
# #-----------------------------

# resource "aws_db_instance" "instance" {
#   engine         = "mysql"
#   engine_version = "8.0.16"

#   identifier = "terr-pres-rds"

#   username = var.mysql_user
#   password = var.mysql_password

#   instance_class = "db.t2.micro"

#   allocated_storage     = 20
#   max_allocated_storage = 50
#   storage_type          = "gp2"
#   storage_encrypted     = false

#   multi_az               = true
#   db_subnet_group_name   = aws_db_subnet_group.subnet_group.name
#   vpc_security_group_ids = [module.mysql_sg.security_group_id]
#   publicly_accessible    = false
#   port                   = 3306

#   # tfvarsを使用する場合は「name」ではなく、「db_name」
#   db_name              = var.mysql_database
#   parameter_group_name = aws_db_parameter_group.db_parameter.name
#   option_group_name    = aws_db_option_group.db_option.name

#   backup_window              = "04:00-05:00"
#   backup_retention_period    = 7
#   maintenance_window         = "Sun:03:00-Sun:04:00"
#   auto_minor_version_upgrade = false

#   deletion_protection = false
#   skip_final_snapshot = true
#   apply_immediately   = true

#   tags = {
#     "Name" = "terr-pres-rds-instance"
#   }
# }

# module "mysql_sg" {
#   source      = "./security_group"
#   name        = "mysql-sg"
#   vpc_id      = aws_vpc.terr_pres_vpc.id
#   port        = 3306
#   cidr_blocks = [aws_vpc.terr_pres_vpc.cidr_block]
# }
