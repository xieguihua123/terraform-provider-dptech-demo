
terraform {
 required_providers {
  dptech-demo={
     source = "xieguihua123/dptech-demo"
     version = "1.2.7"
   } 
 }
 }

provider "dptech-demo" {
  address="Http://localhost:19090"
  name="123"
}


 resource "dptech-demo_example" "dptech-demo" {
  uuid_count = "3"
}