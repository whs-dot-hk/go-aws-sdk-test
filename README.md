Follow along

## Install go
https://golang.org/doc/install#install

## Install awscli
https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-linux.html#cliv2-linux-install

## Configure aws

Enter `us-east-1` for default region name

```
$ aws configure
```

## Create key pair

```
$ go run main.go create key "my openvpn key"
Created key pair my openvpn key
Created my openvpn key.pem
```

## Create server

```
$ go run main.go create server my-openvpn-server --key "my openvpn key"
Created stack my-openvpn-server
PUBLIC_DNS_NAME=ec2-3-221-4-14.compute-1.amazonaws.com
```

## Install openvpn on server

### Update public dns name

Copy the last line of the previous output and paste it here

```
$ PUBLIC_DNS_NAME=ec2-3-221-4-14.compute-1.amazonaws.com
$ sed -i -e 's|^public_dns_name=.*$|public_dns_name='$PUBLIC_DNS_NAME'|g' install_once.sh
```

### Update key name
```
$ KEY_NAME="my openvpn key"
$ sed -i -e 's|^key_name=.*$|key_name="'"$KEY_NAME"'"|g' install_once.sh
```

### Run `install_once.sh`
```
$ cat install_once.sh
#!/bin/sh

public_dns_name=ec2-3-221-4-14.compute-1.amazonaws.com
key_name="my openvpn key"

ssh -i "$key_name.pem" ubuntu@$public_dns_name "curl https://raw.githubusercontent.com/whs-dot-hk/auto-openvpn-install/master/install_or_remove.sh | bash"

scp -i "$key_name.pem" ubuntu@$public_dns_name:/home/ubuntu/client1.ovpn .

sed '/verb 3/a script-security 2\nup /etc/openvpn/update-resolv-conf\ndown /etc/openvpn/update-resolv-conf' client1.ovpn > client1_2.ovpn
$ ls -l my\ openvpn\ key.pem
-r--------. 1 whs whs 1670 Apr 13 21:09 'my openvpn key.pem'
$ ./install_once.sh
```

You should see finished and `client1.opvn` and `client1_2.opvn` is
created

```
$ ls -l client1*
-rw-rw-r--. 1 whs whs 5080 Apr 13 21:39 client1_2.ovpn
-rw-r--r--. 1 whs whs 4990 Apr 13 21:39 client1.ovpn
```

## Install openvpn

```
$ sudo dnf install openvpn -y
```

## Install update-resolv-conf (optional)

```
$ curl https://raw.githubusercontent.com/alfredopalhares/openvpn-update-resolv-conf/master/update-resolv-conf.sh -o update-resolv-conf
$ chmod a+x update-resolv-conf
$ sudo cp update-resolv-conf /etc/openvpn
```

## Install openresolv (optional)

```
$ curl -O https://roy.marples.name/downloads/openresolv/openresolv-3.10.0.tar.xz
$ tar -xvf openresolv-3.10.0.tar.xz
$ cd openresolv-3.10.0
$ ./configure
$ sudo make install
```

## Enjoy
```
$ sudo openvpn --config client1_2.ovpn
```

Visit https://www.dnsleaktest.com to test your new vpn
