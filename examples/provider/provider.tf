
terraform {
 required_providers {
  dptech-demo={
     source = "xieguihua123/dptech-demo"
     version = "1.2.9"
   } 
 }
 }

provider "dptech-demo" {
  address="Http://localhost:8080"
  name="123"
}

 resource "dptech-demo_example" "a" {
  uuid_count = "3"
}