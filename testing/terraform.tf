terraform {
  required_version = "1.9"

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "5.71.0"
    }
    vault = {
      source  = "hashicorp/vault"
      version = "4.4.0"
    }
  }
}

resource "aws_instance" "web" {
  instance_type = "t2.micro"
}

locals {
  service_account_name = "hmm"

  a_dictionary = {
    "one"  = "fish"
    "two"  = "fish",
    "red"  = "fish"
    "blue" = "fish"
  }
}

resource "vault_generic_endpoint" "user" {
  depends_on = [
    random_password.svc_acc_pass
  ]
  path                 = "auth/userpass/users/${replace(local.service_account_name, "-", "_")}"
  ignore_absent_fields = true
  disable_read         = true
  data_json            = <<EOT
{
    "policies": [
        "${local.service_account_name}-rw",
        "AA-DP-write",
        "SOHB-DP-write"
    ],
    "password": "${random_password.svc_acc_pass.result}"
}${local.a_dictionary}
EOT
}

data "aws_ami" "this" {
  most_recent = true
  owners      = ["amazon"]

  filter {
    name   = "architecture"
    values = ["arm64"]
  }
  filter {
    name = "name"
    values = [
      "al2023-ami-2023*"
    ]
  }
}

output "output1" {
  value = aws_instance.web.arn
}
output "output2" {
  value = data.aws_ami.this.arn
}
