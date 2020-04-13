#!/bin/sh

public_dns_name=
key_name=

ssh -i "$key_name.pem" ubuntu@$public_dns_name "curl https://raw.githubusercontent.com/whs-dot-hk/auto-openvpn-install/master/install_or_remove.sh | bash"

scp -i "$key_name.pem" ubuntu@$public_dns_name:/home/ubuntu/client1.ovpn .

sed '/verb 3/a script-security 2\nup /etc/openvpn/update-resolv-conf\ndown /etc/openvpn/update-resolv-conf' client1.ovpn > client1_2.ovpn
