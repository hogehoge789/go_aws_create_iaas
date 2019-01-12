package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region           = "ap-northeast-1"
	vpcCidr, vpcName = "172.18.152.0/24", "testaws01-vpnb1"
)

func main() {
	sess := session.Must(session.NewSession())

	client := ec2.New(sess, aws.NewConfig().WithRegion(region))

	// create vpc
	createout, err := client.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock:                   aws.String(vpcCidr),
		AmazonProvidedIpv6CidrBlock: aws.Bool(false),
		InstanceTenancy:             aws.String("default"),
	})
	if err != nil {
		log.Fatalln("02", err)
	}

	vpc := createout.Vpc
	fmt.Println("VPC created:", *vpc.VpcId)

	// name this vpc
	_, err = client.CreateTags(&ec2.CreateTagsInput{
		Tags: []*ec2.Tag{
			{Key: aws.String("Name"), Value: aws.String(vpcName)},
		},
		Resources: []*string{vpc.VpcId},
	})
}
