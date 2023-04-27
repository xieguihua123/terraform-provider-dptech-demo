
# terraform {
#  required_providers {
#   dptech-demo={
#      source = "xieguihua123/dptech-demo"
#      version = "1.2.10"
#    } 
#  }
#  }

# provider "dptech-demo" {
#   address="http://localhost:8080"
#   name="123"
# }

#  resource "dptech-demo_example" "a" {
#   uuid_count = "3"
# }

terraform {
  required_providers {
    dptech-demo = {
      source = "registry.terraform.io/xieguihua123/dptech-demo"
      version = "1.2.10"
    }
  }
}

provider "dptech-demo" {
  # Configuration options
}