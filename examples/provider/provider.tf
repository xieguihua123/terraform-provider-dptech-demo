
terraform {
 required_providers {
  dptech-demo={
     source = "github.com/xieguihua123/dptech-demo"
     version = "1.2.6"
   } 
 }
 }

provider "dptech-demo" {
  address="localhost"
  name=""
}


 resource "dptech-demo_example" "dptech-demo" {
  uuid_count = "3"
}