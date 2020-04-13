package key

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateKey(keyName string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := ec2.New(sess)

	input := &ec2.CreateKeyPairInput{
		KeyName: aws.String(keyName),
	}

	result, err := svc.CreateKeyPair(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "InvalidKeyPair.Duplicate" {
			log.Fatalf("Key pair %s exists\n", keyName)
		}
		log.Fatal(err)
	}

	fmt.Printf("Created key pair %s\n", *result.KeyName)

	ioutil.WriteFile(keyName+".pem", []byte(*result.KeyMaterial), 0400)

	fmt.Printf("Created %s.pem\n", keyName)
}
