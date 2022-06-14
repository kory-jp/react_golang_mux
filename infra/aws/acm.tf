#---------------------------
#  AWS Certificate Manger
#---------------------------

resource "aws_acm_certificate" "terr_pres_acm" {
  domain_name       = aws_route53_zone.terr_pres_route53.name
  validation_method = "DNS"

  lifecycle {
    create_before_destroy = true
    prevent_destroy       = true
  }

  tags = {
    "Name" = "terr_pres_acm"
  }

  depends_on = [
    aws_route53_zone.terr_pres_route53
  ]
}

resource "aws_route53_record" "route53_acm_dns_resolve" {
  for_each = {
    for dvo in aws_acm_certificate.terr_pres_acm.domain_validation_options : dvo.domain_name => {
      name   = dvo.resource_record_name
      type   = dvo.resource_record_type
      record = dvo.resource_record_value
    }
  }

  allow_overwrite = true
  zone_id         = aws_route53_zone.terr_pres_route53.id
  name            = each.value.name
  type            = each.value.type
  ttl             = 600
  records         = [each.value.record]

  lifecycle {
    prevent_destroy = true
  }
}

resource "aws_acm_certificate_validation" "cert_valid" {
  certificate_arn         = aws_acm_certificate.terr_pres_acm.arn
  validation_record_fqdns = [for record in aws_route53_record.route53_acm_dns_resolve : record.fqdn]

  lifecycle {
    prevent_destroy = true
  }
}