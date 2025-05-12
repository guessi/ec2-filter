package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func getDescribeInstancesInput(key, value string) *ec2.DescribeInstancesInput {
	return &ec2.DescribeInstancesInput{
		Filters: []types.Filter{
			{
				Name: aws.String(fmt.Sprintf("tag:%s", key)),
				Values: []string{
					*aws.String(value),
				},
			},
		},
	}
}

func describeInstances(region, key, value string) *ec2.DescribeInstancesOutput {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ec2.NewFromConfig(cfg, func(o *ec2.Options) {
		o.Region = region
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var filter *ec2.DescribeInstancesInput
	if len(key) > 0 && len(value) > 0 {
		filter = getDescribeInstancesInput(key, value)
	}

	result, err := svc.DescribeInstances(ctx, filter)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func buildOutput(result *ec2.DescribeInstancesOutput) {
	var b strings.Builder

	fields := "%-32s %-16s %-12s %-16s %s\n"

	fmt.Printf(fields, "Name", "State", "InstaceType", "PrivateIP", "InstanceId")

	for _, reservation := range result.Reservations {
		for _, instance := range reservation.Instances {
			var instanceName string
			var privateIpAddress string

			if len(instance.NetworkInterfaces) > 0 {
				privateIpAddress = *instance.NetworkInterfaces[0].PrivateIpAddress
			} else {
				privateIpAddress = "n/a"
			}

			for _, tag := range instance.Tags {
				if *tag.Key == "Name" {
					instanceName = *tag.Value
					continue
				}
			}

			if len(instanceName) == 0 {
				instanceName = "n/a"
			}

			fmt.Fprintf(&b, fields,
				instanceName,
				fmt.Sprintf("%s", cases.Title(language.English).String(string(instance.State.Name))),
				instance.InstanceType,
				privateIpAddress,
				*instance.InstanceId,
			)
		}
	}
	fmt.Printf("%s", b.String())
}
