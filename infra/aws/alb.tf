resource "aws_lb" "terr_pres_alb" {
  name               = "terr-pres-alb"
  load_balancer_type = "application"
  internal           = false
  idle_timeout       = 60

  subnets = [
    aws_subnet.terr_pres_public_subnet_1a.id,
    aws_subnet.terr_pres_public_subnet_1c.id,
  ]

  access_logs {
    bucket  = aws_s3_bucket.alb_log.id
    enabled = true
  }

  security_groups = [
    module.http_sg.security_group_id,
    module.https_sg.security_group_id,
    module.http_redirect_sg.security_group_id,
  ]

  tags = {
    "Name" = "terr_pres_alb"
  }
}

output "alb_dns_name" {
  value = aws_lb.terr_pres_alb.dns_name
}

module "http_sg" {
  source      = "./security_group"
  name        = "http-sg"
  vpc_id      = aws_vpc.terr_pres_vpc.id
  port        = 80
  cidr_blocks = ["0.0.0.0/0"]
}

module "https_sg" {
  source      = "./security_group"
  name        = "https-sg"
  vpc_id      = aws_vpc.terr_pres_vpc.id
  port        = 443
  cidr_blocks = ["0.0.0.0/0"]
}

module "http_redirect_sg" {
  source      = "./security_group"
  name        = "http-redirect-sg"
  vpc_id      = aws_vpc.terr_pres_vpc.id
  port        = 8080
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_lb_listener" "http" {
  load_balancer_arn = aws_lb.terr_pres_alb.arn
  port              = 80
  protocol          = "HTTP"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "これはhttpです"
      status_code  = "200"
    }
  }
}

resource "aws_lb_listener" "https" {
  load_balancer_arn = aws_lb.terr_pres_alb.arn
  port              = 443
  protocol          = "HTTPS"
  certificate_arn   = aws_acm_certificate.terr_pres_acm.arn
  ssl_policy        = "ELBSecurityPolicy-2016-08"

  default_action {
    type = "fixed-response"

    fixed_response {
      content_type = "text/plain"
      message_body = "これはHTTPSです"
      status_code  = "200"
    }
  }
}

resource "aws_lb_listener" "redirect_http_to_https" {
  load_balancer_arn = aws_lb.terr_pres_alb.arn
  port              = 8080
  protocol          = "HTTP"

  default_action {
    type = "redirect"

    redirect {
      port        = "443"
      protocol    = "HTTPS"
      status_code = "HTTP_301"
    }
  }
}


# ----------------------------
# Target Group
# ----------------------------

resource "aws_lb_target_group" "terr_pres_tg" {
  name                 = "terr-pres-tg"
  target_type          = "ip"
  vpc_id               = aws_vpc.terr_pres_vpc.id
  port                 = 80
  protocol             = "HTTP"
  deregistration_delay = 300

  health_check {
    path                = "/"
    healthy_threshold   = 5
    unhealthy_threshold = 2
    timeout             = 5
    interval            = 30
    matcher             = 200
    port                = "traffic-port"
    protocol            = "HTTP"
  }

  depends_on = [
    aws_lb.terr_pres_alb
  ]

  tags = {
    "Name" = "terr_pres_tg"
  }
}

resource "aws_lb_listener_rule" "terr_pres_lb_listener_rule" {
  listener_arn = aws_lb_listener.https.arn
  priority     = 100

  action {
    type             = "forward"
    target_group_arn = aws_lb_target_group.terr_pres_tg.arn
  }

  condition {
    path_pattern {
      values = ["/*"]
    }
  }

}
