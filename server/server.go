package server

import (
	"fmt"
	"log"

	"github.com/whs-dot-hk/go-aws-sdk-test/template"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	cf "github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateStack(stackName string, keyName string) {
	t := template.NewOpenvpnTemplate()

	t.KeyName = keyName

	cft := t.GetCfTemplate()

	yaml, err := cft.YAML()
	if err != nil {
		log.Fatal(err)
	}

	templateBody := string(yaml)

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := cf.New(sess)

	input := &cf.CreateStackInput{
		StackName:    &stackName,
		TemplateBody: &templateBody,
	}

	stack, err := svc.CreateStack(input)
	if err != nil {
		log.Fatal(err)
	}

	input2 := &cf.DescribeStacksInput{
		StackName: stack.StackId,
	}

	err = svc.WaitUntilStackCreateComplete(input2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created stack %s\n", stackName)

	stackOutput, err := svc.DescribeStacks(input2)
	if err != nil {
		log.Fatal(err)
	}

	outputs := stackOutput.Stacks[0].Outputs
	for _, output := range outputs {
		if *output.OutputKey == "OpenvpnInstanceId" {
			ec2Svc := ec2.New(sess)
			input3 := &ec2.DescribeInstancesInput{
				InstanceIds: []*string{
					aws.String(*output.OutputValue),
				},
			}

			result, err := ec2Svc.DescribeInstances(input3)
			if err != nil {
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					default:
						log.Fatal(aerr.Error())
					}
				} else {
					log.Fatal(err.Error())
				}
			}

			fmt.Printf("PUBLIC_DNS_NAME=%s\n", *result.Reservations[0].Instances[0].PublicDnsName)
		}
	}
}
