
terraform {
 required_providers {
  dptech-demo={
     source = "registry.terraform.io/xieguihua123/dptech-demo"
     version = "1.2.25"
   } 
 }
 }

provider "dptech-demo" {
  address="http://localhost:"
  port="123"
}

 resource "dptech-demo_RealService" "cs" {
  name="rs1"
  address="192.168.1.1"
  port="8091"
}