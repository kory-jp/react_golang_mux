# ------------------------
# Cluster
# ------------------------

resource "aws_ecs_cluster" "terr_pres_cluster" {
  name = var.ecs_cluster_name
}

# ------------------------
# Task
# ------------------------

resource "aws_ecs_task_definition" "terr_pres_task" {
  family                   = "terr_pres_task"
  cpu                      = "256"
  memory                   = "1024"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]

  # ====== RDS ======

  container_definitions = templatefile("./containers/app_container_definition.json", {
    db_image_uri          = aws_ecr_repository.db.repository_url
    api_image_uri         = aws_ecr_repository.api.repository_url
    client_image_uri      = aws_ecr_repository.client.repository_url
    api_container_name    = var.ecs_task_api_container_name
    client_container_name = var.ecs_task_client_container_name
    db_container_name     = var.ecs_task_db_container_name
    mysql_database        = var.mysql_database
    mysql_password        = var.mysql_password
    mysql_root_password   = var.mysql_root_password
    mysql_user            = var.mysql_user
  })

  # ====== RDS ======

  # container_definitions = templatefile("./containers/rds_container_definition.json", {
  #   api_image_uri         = aws_ecr_repository.api.repository_url
  #   client_image_uri      = aws_ecr_repository.client.repository_url
  #   api_container_name    = var.ecs_task_api_container_name
  #   client_container_name = var.ecs_task_client_container_name
  #   mysql_database        = var.mysql_database
  #   mysql_password        = var.mysql_password
  #   mysql_root_password   = var.mysql_root_password
  #   mysql_user            = var.mysql_user
  # })

  # depends_on = [
  #   aws_db_instance.instance
  # ]

  # ====== RDS ======

  execution_role_arn = aws_iam_role.ecs_task.arn
  task_role_arn      = aws_iam_role.session_manager.arn
}

# ------------------------
# Service
# ------------------------

resource "aws_ecs_service" "terr_pres_service" {
  name                              = var.fargate_service_name
  cluster                           = aws_ecs_cluster.terr_pres_cluster.arn
  task_definition                   = aws_ecs_task_definition.terr_pres_task.arn
  desired_count                     = 1
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  health_check_grace_period_seconds = 3600
  enable_execute_command            = true

  network_configuration {
    assign_public_ip = false
    security_groups  = [module.nginx_sg.security_group_id]

    subnets = [
      aws_subnet.terr_pres_private_subnet_1a.id,
      aws_subnet.terr_pres_private_subnet_1c.id,
    ]
  }

  load_balancer {
    target_group_arn = aws_lb_target_group.terr_pres_tg.arn
    container_name   = var.ecs_task_client_container_name
    container_port   = 80
  }

  lifecycle {
    ignore_changes = [task_definition]
  }
}

# ------------------------
# Security Group
# ------------------------

module "nginx_sg" {
  source      = "./security_group"
  name        = "nginx-sg"
  vpc_id      = aws_vpc.terr_pres_vpc.id
  port        = 80
  cidr_blocks = [aws_vpc.terr_pres_vpc.cidr_block]
}


# ------------------------
# ECS Iam Role
# ------------------------

data "aws_iam_policy_document" "assume_role" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["ecs-tasks.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "ecs_task" {
  name               = "MyTaskRole"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}


resource "aws_iam_policy_attachment" "ecs_task" {
  name       = "EcsPolicyAttachment"
  roles      = [aws_iam_role.ecs_task.name]
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

resource "aws_iam_policy" "allow-ecs-exec" {
  name   = "allow-ecs-exec"
  policy = file("policy/ecs-exec.json")
}

resource "aws_iam_role" "session_manager" {
  name               = "session_manager"
  assume_role_policy = data.aws_iam_policy_document.assume_role.json
}

resource "aws_iam_role_policy_attachment" "session_manager" {
  policy_arn = aws_iam_policy.allow-ecs-exec.arn
  role       = aws_iam_role.session_manager.name
}
