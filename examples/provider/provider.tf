
terraform {
 required_providers {
  dptech-demo={
     source = "registry.terraform.io/xieguihua123/dptech-demo"
     version = "1.2.17"
   } 
 }
 }

provider "dptech-demo" {
  address="Http://localhost:8080"
  name="123"
 
}

 resource "dptech-demo_example" "Exampleresource" {
  uuid_count="3"
}