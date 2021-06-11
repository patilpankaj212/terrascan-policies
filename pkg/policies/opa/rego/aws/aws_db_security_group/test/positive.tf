provider "aws" {
  region = "us-west-2"
}


resource "aws_db_security_group" "rdsHostsHigherThan256" {
  name = "rds_sg"

  ingress {
    cidr = "192.164.0.0/23"
  }
}

resource "aws_db_security_group" "rdsIsPublic" {
  name = "rds_sg"

  ingress {
    cidr = "0.0.0.0/0"
  }
}

resource "aws_db_security_group" "rdsScopeIsPublic" {
  name = "rds_sg"

  ingress {
    cidr = "192.164.0.0/32"
  }
}