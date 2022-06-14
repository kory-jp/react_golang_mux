#---------------------------
#  Route53
#---------------------------

resource "aws_route53_zone" "terr_pres_route53" {
  name = var.domain

  lifecycle {
    prevent_destroy = true
  }

  tags = {
    "Name" = "terr_pres_route53"
  }
}

resource "aws_route53_record" "terr_pres_r53_record" {
  zone_id = aws_route53_zone.terr_pres_route53.id
  name    = aws_route53_zone.terr_pres_route53.name
  type    = "A"

  alias {
    name                   = aws_lb.terr_pres_alb.dns_name
    zone_id                = aws_lb.terr_pres_alb.zone_id
    evaluate_target_health = true
  }

  # lifecycle {
  #   prevent_destroy = true
  # }
}

output "domain_name" {
  value = aws_route53_record.terr_pres_r53_record.name
}
