# -------------------------
#  VPC
# -------------------------

resource "aws_vpc" "terr_pres_vpc" {
  cidr_block           = "10.0.0.0/16"
  instance_tenancy     = "default"
  enable_dns_support   = true
  enable_dns_hostnames = true

  tags = {
    "Name" = "terr_pres_vpc"
  }
}

# ------------------------
#  Public Subnet
# ------------------------

resource "aws_subnet" "terr_pres_public_subnet_1a" {
  vpc_id                  = aws_vpc.terr_pres_vpc.id
  cidr_block              = "10.0.1.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1a"

  tags = {
    "Name" = "terr_pres_public_subnet_1a"
  }
}

resource "aws_subnet" "terr_pres_public_subnet_1c" {
  vpc_id                  = aws_vpc.terr_pres_vpc.id
  cidr_block              = "10.0.2.0/24"
  map_public_ip_on_launch = true
  availability_zone       = "ap-northeast-1c"

  tags = {
    "Name" = "terr_pres_public_subnet_1c"
  }
}

resource "aws_internet_gateway" "terr_pres_igw" {
  vpc_id = aws_vpc.terr_pres_vpc.id

  tags = {
    "Name" = "terr_pres_igw"
  }
}

resource "aws_route_table" "terr_pres_public_route_table" {
  vpc_id = aws_vpc.terr_pres_vpc.id

  tags = {
    "Name" = "terr_pres_public_route_table"
  }
}

resource "aws_route" "terr_pres_public_route" {
  route_table_id         = aws_route_table.terr_pres_public_route_table.id
  gateway_id             = aws_internet_gateway.terr_pres_igw.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table_association" "terr_pres_public_rtAssociation_1a" {
  subnet_id      = aws_subnet.terr_pres_public_subnet_1a.id
  route_table_id = aws_route_table.terr_pres_public_route_table.id
}

resource "aws_route_table_association" "terr_pres_public_rtAssociation_1c" {
  subnet_id      = aws_subnet.terr_pres_public_subnet_1c.id
  route_table_id = aws_route_table.terr_pres_public_route_table.id
}

# ---------------------------
# Private Subnet
#  --------------------------

resource "aws_subnet" "terr_pres_private_subnet_1a" {
  vpc_id                  = aws_vpc.terr_pres_vpc.id
  cidr_block              = "10.0.65.0/24"
  availability_zone       = "ap-northeast-1a"
  map_public_ip_on_launch = false

  tags = {
    "Name" = "terr_pres_private_subnet_1a"
  }
}

resource "aws_subnet" "terr_pres_private_subnet_1c" {
  vpc_id                  = aws_vpc.terr_pres_vpc.id
  cidr_block              = "10.0.66.0/24"
  availability_zone       = "ap-northeast-1c"
  map_public_ip_on_launch = false

  tags = {
    "Name" = "terr_pres_private_subnet_1c"
  }
}

resource "aws_route_table" "terr_pres_private_route_table_1a" {
  vpc_id = aws_vpc.terr_pres_vpc.id

  tags = {
    "Name" = "terr_pres_private_route_table_1a"
  }
}

resource "aws_route_table" "terr_pres_private_route_table_1c" {
  vpc_id = aws_vpc.terr_pres_vpc.id

  tags = {
    "Name" = "terr_pres_private_route_table_1c"
  }
}

resource "aws_route" "terr_pres_private_route_1a" {
  route_table_id         = aws_route_table.terr_pres_private_route_table_1a.id
  nat_gateway_id         = aws_nat_gateway.terr_pres_ng_1a.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route" "terr_pres_private_route_1c" {
  route_table_id         = aws_route_table.terr_pres_private_route_table_1c.id
  nat_gateway_id         = aws_nat_gateway.terr_pres_ng_1c.id
  destination_cidr_block = "0.0.0.0/0"
}

resource "aws_route_table_association" "terr_pres_private_rtAssociation_1a" {
  subnet_id      = aws_subnet.terr_pres_private_subnet_1a.id
  route_table_id = aws_route_table.terr_pres_private_route_table_1a.id
}

resource "aws_route_table_association" "terr_pres_private_rtAssociation_1c" {
  subnet_id      = aws_subnet.terr_pres_private_subnet_1c.id
  route_table_id = aws_route_table.terr_pres_private_route_table_1c.id
}


# -----------------------------
#  Nat Gateway
# -----------------------------

resource "aws_eip" "terr_pres_eip_1a" {
  vpc = true

  depends_on = [aws_internet_gateway.terr_pres_igw]
}

resource "aws_eip" "terr_pres_eip_1c" {
  vpc = true

  depends_on = [aws_internet_gateway.terr_pres_igw]
}

resource "aws_nat_gateway" "terr_pres_ng_1a" {
  allocation_id = aws_eip.terr_pres_eip_1a.id
  subnet_id     = aws_subnet.terr_pres_public_subnet_1a.id

  depends_on = [aws_internet_gateway.terr_pres_igw]

  tags = {
    "Name" = "terr_pres_ng_1a"
  }
}

resource "aws_nat_gateway" "terr_pres_ng_1c" {
  allocation_id = aws_eip.terr_pres_eip_1c.id
  subnet_id     = aws_subnet.terr_pres_public_subnet_1c.id

  depends_on = [aws_internet_gateway.terr_pres_igw]

  tags = {
    "Name" = "terr_pres_ng_1c"
  }
}