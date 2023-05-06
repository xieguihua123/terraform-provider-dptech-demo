
terraform {
 required_providers {
  dptech-demo={
     source = "registry.terraform.io/xieguihua123/dptech-demo"
     version = "1.2.30"
   } 
 }
 }

provider "dptech-demo" {
  address="http://localhost:"
  port="8080"
  username="test"
  password="jsepc123!"
}

 resource "dptech-demo_RealService" "cs" {
  rsinfo={
  name="rs1"
  address="192.168.1.1"
  port="8091"
}
}