package template

import (
	cf "github.com/awslabs/goformation/v4/cloudformation"
	"github.com/awslabs/goformation/v4/cloudformation/ec2"
)

func addOpenvpnInstanceResource(t *cf.Template, o *OpenvpnTemplate) {
	imageId := cf.FindInMap("RegionMap", cf.Ref("AWS::Region"), "AMI")
	securityGroup := cf.Ref("OpenvpnSecurityGroup")
	t.Resources["OpenvpnInstance"] = &ec2.Instance{
		ImageId:        imageId,
		InstanceType:   "t2.nano",
		SecurityGroups: []string{securityGroup},
		KeyName:        o.KeyName,
	}
}

func addOpenvpnSecurityGroupResource(t *cf.Template) {
	openvpn := ec2.SecurityGroup_Ingress{
		CidrIp:     "0.0.0.0/0",
		FromPort:   1194,
		IpProtocol: "udp",
		ToPort:     1194,
	}
	ssh := ec2.SecurityGroup_Ingress{
		CidrIp:     "0.0.0.0/0",
		FromPort:   22,
		IpProtocol: "tcp",
		ToPort:     22,
	}
	securityGroupIngress := []ec2.SecurityGroup_Ingress{openvpn, ssh}
	t.Resources["OpenvpnSecurityGroup"] = &ec2.SecurityGroup{
		GroupDescription:     "Allow openvpn and ssh access",
		SecurityGroupIngress: securityGroupIngress,
	}
}

func addOpenvpnEipResource(t *cf.Template) {
	t.Resources["OpenvpnEip"] = &ec2.EIP{
		InstanceId: cf.Ref("OpenvpnInstance"),
	}
}

func addRegionMapMapping(t *cf.Template) {
	t.Mappings["RegionMap"] = map[string]interface{}{
		"us-east-1": map[string]interface{}{"AMI": "ami-07ebfd5b3428b6f4d"},
	}
}

func addOpenvpnInstanceIdOutput(t *cf.Template) {
	t.Outputs["OpenvpnInstanceId"] = map[string]interface{}{
		"Description": "Openvpn instance id",
		"Value":       cf.Ref("OpenvpnInstance"),
	}
}

func addResources(t *cf.Template, o *OpenvpnTemplate) {
	addOpenvpnInstanceResource(t, o)
	addOpenvpnSecurityGroupResource(t)
	addOpenvpnEipResource(t)
}

func addMappings(t *cf.Template) {
	addRegionMapMapping(t)
}

func addOutputs(t *cf.Template) {
	addOpenvpnInstanceIdOutput(t)
}

func addAll(t *cf.Template, o *OpenvpnTemplate) {
	addResources(t, o)
	addMappings(t)
	addOutputs(t)
}

func getCfTemplate(o *OpenvpnTemplate) *cf.Template {
	t := cf.NewTemplate()
	addAll(t, o)
	return t
}
