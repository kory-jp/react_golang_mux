# ---------------------
# CloudWatch
# ---------------------

resource "aws_cloudwatch_log_group" "for_db" {
  name              = "/ecs/db"
  retention_in_days = 30

  tags = {
    "Name" = "terr_pres_cw_db"
  }
}

resource "aws_cloudwatch_log_group" "for_api" {
  name              = "/ecs/api"
  retention_in_days = 30

  tags = {
    "Name" = "terr_pres_cw_api"
  }
}

resource "aws_cloudwatch_log_group" "for_client" {
  name              = "/ecs/client"
  retention_in_days = 30

  tags = {
    "Name" = "terr_pres_cw_client"
  }
}
