package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region, vpccidr string = "ap-northeast-1", "172.18.128.0"
)

func main() {
	sess := session.Must(session.NewSession())

	client := ec2.New(sess, aws.NewConfig().WithRegion(region))

	// create vpc
	createout, err := client.CreateVpc(&ec2.CreateVpcInput{
		CidrBlock: aws.String(vpccidr),
	})
	if err != nil {
		log.Fatalln("02", err)
	}

	vpc := createout.Vpc
	fmt.Println("VPC created:", *vpc.VpcId)
}
