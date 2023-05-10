
terraform {
 required_providers {
  dptech-demo={
     source = "registry.terraform.io/xieguihua123/dptech-demo"
     version = "1.2.35"
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

resource "dptech-demo_RealServiceList" "cs" {
 poollist={
  name="string"
  monitor="string"
  rs_list="string"
  schedule="string"
 }
}

resource "dptech-demo_VirtualService" "cs" {
  virtualservice={
  name ="string"
	state ="string"
	mode ="string"
	ip ="string"
	port ="string"
	protocol ="string"
	session_keep ="string"
	default_pool ="string"
	tcp_policy ="string"
	snat ="string"
	session_bkp ="string"
	vrrp ="string"      //涉及普通双机热备场景，需要关联具体的vrrp组
}
}

resource "dptech-demo_AddrPoolList" "cs" {
addrpoollist={
  name="string"
	ip_version="string"
	ip_start="string"
	ip_end="string"
	vrrp_if_name="string"//接口名称
	vrrp_id="string"    //vrid
} 
}

