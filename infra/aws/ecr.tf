# -----------------------------
#  Elastic Container Registry
# -----------------------------

resource "aws_ecr_repository" "client" {
  name                 = var.client_image_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

resource "aws_ecr_repository" "api" {
  name                 = var.api_image_name
  image_tag_mutability = "MUTABLE"

  image_scanning_configuration {
    scan_on_push = true
  }
}

# resource "aws_ecr_repository" "db" {
#   name                 = var.db_image_name
#   image_tag_mutability = "MUTABLE"

#   image_scanning_configuration {
#     scan_on_push = true
#   }
# }


# -----------------------------
#  Local Image Push Command
# -----------------------------

resource "null_resource" "command" {
  provisioner "local-exec" {
    command = "aws ecr get-login-password --region ap-northeast-1 | docker login --username AWS --password-stdin ${var.aws_account_id}.dkr.ecr.ap-northeast-1.amazonaws.com"
  }

  # -- client ----
  provisioner "local-exec" {
    command = "docker build -t ${var.client_image_name} ../../client"
  }

  provisioner "local-exec" {
    command = "docker tag ${var.client_image_name}:latest ${aws_ecr_repository.client.repository_url}"
  }

  provisioner "local-exec" {
    command = "docker push ${aws_ecr_repository.client.repository_url}"
  }

  # --- api ---

  provisioner "local-exec" {
    command = "docker build -t ${var.api_image_name} ../../api"
  }

  provisioner "local-exec" {
    command = "docker tag ${var.api_image_name}:latest ${aws_ecr_repository.api.repository_url}"
  }

  provisioner "local-exec" {
    command = "docker push ${aws_ecr_repository.api.repository_url}"
  }

  # --- mysql ---
  # provisioner "local-exec" {
  #   command = "docker build -t ${var.db_image_name} ../../mysql"
  # }

  # provisioner "local-exec" {
  #   command = "docker tag ${var.db_image_name}:latest ${aws_ecr_repository.db.repository_url}"
  # }

  # provisioner "local-exec" {
  #   command = "docker push ${aws_ecr_repository.db.repository_url}"
  # }
}
