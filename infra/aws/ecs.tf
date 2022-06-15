# ------------------------
# Cluster
# ------------------------

resource "aws_ecs_cluster" "terr_pres_cluster" {
  name = "terr_pres_cluster"
}

# ------------------------
# Task
# ------------------------

resource "aws_ecs_task_definition" "terr_pres_task" {
  family                   = "terr_pres_task"
  cpu                      = "256"
  memory                   = "512"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  container_definitions    = file("./containers/container_definitions.json")
  execution_role_arn       = aws_iam_role.ecs_task.arn
}

# ------------------------
# Service
# ------------------------

resource "aws_ecs_service" "terr_pres_service" {
  name                              = "terr_pres_service"
  cluster                           = aws_ecs_cluster.terr_pres_cluster.arn
  task_definition                   = aws_ecs_task_definition.terr_pres_task.arn
  desired_count                     = 1
  launch_type                       = "FARGATE"
  platform_version                  = "1.4.0"
  health_check_grace_period_seconds = 60

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
    container_name   = "example"
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
