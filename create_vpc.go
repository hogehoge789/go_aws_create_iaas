package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	region                                     = "ap-northeast-1"
	vpcCidr, vpcName                           = "172.1.152.0/24", "testaws01"
	dhcpsetname, dhcpdomain, dhcpdns           = "testaws01-001", "ap-northeast-1.compute.internal", "AmazonProvidedDNS"
	dhcpntp01, dhcpntp02, dhcpntp03, dhcpntp04 = "172.1.1.100", "172.1.1.101", "172.1.1.116", "172.1.1.117"
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

	if err != nil {
		log.Fatalln("03", err)
	}

	/*
		// create dhcpset
		createdhcp, err := client.CreateDhcpOptions(&ec2.CreateDhcpOptionsInput{
			DhcpConfigurations: []*ec2.DhcpConfiguration{
				&ec2.DhcpConfiguration{
					Key: aws.String("domain-name"),
					Values: []*string{
						aws.String(dhcpdomain),
					},
				},
				&ec2.DhcpConfiguration{
					Key: aws.String("domain-name-servers"),
					Values: []*string{
						aws.String(dhcpdns),
					},
				},
				&ec2.DhcpConfiguration{
					Key: aws.String("ntp-servers"),
					Values: []*string{
						aws.String(dhcpntp01),
						aws.String(dhcpntp02),
						aws.String(dhcpntp03),
						aws.String(dhcpntp04),
					},
				},
			},
		})

		_, err = client.CreateTags(&ec2.CreateTagsInput{
			Tags: []*ec2.Tag{
				{Key: aws.String("Name"), Value: aws.String(vpcName)},
			},
			Resources: []*string{vpc.VpcId},
		})

	*/

	// create subnet

}
