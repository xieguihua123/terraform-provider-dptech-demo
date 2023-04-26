
terraform {
 required_providers {
  dptech-demo={
     source = "xieguihua123/dptech-demo"
     version = "1.2.4"
   } 
 }
 }

provider "dptech-demo" {
  address="localhost"
}

 resource "dptech-demo_example" "dptech-demo" {
  uuid_count = "3"
}
