package aws

import (
	"errors"

	libaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	errorEc2NoResult = errors.New("No instance")
)

type FindEc2IpResponse struct {
	Instances []FindEc2IpResponse_Instance
}

type FindEc2IpResponse_Instance struct {
	ID        string
	PublicIp  string
	PrivateIp string
}

func FindEc2Ip(sessions Sessions, instanceId string) (*FindEc2IpResponse, error) {
	instances := []FindEc2IpResponse_Instance{}

	for _, session := range sessions {
		client := ec2.New(session)

		input := &ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{
				{
					Name:   libaws.String("instance-id"),
					Values: []*string{libaws.String(instanceId)},
				},
			},
		}

		output, err := client.DescribeInstances(input)
		if err != nil {
			return nil, err
		}

		for _, reservation := range output.Reservations {
			for _, instance := range reservation.Instances {
				instances = append(instances, FindEc2IpResponse_Instance{
					ID:        *instance.InstanceId,
					PublicIp:  *instance.PublicIpAddress,
					PrivateIp: *instance.PrivateIpAddress,
				})
			}
		}
	}

	if len(instances) == 0 {
		return nil, errorEc2NoResult
	}

	response := &FindEc2IpResponse{
		Instances: instances,
	}
	return response, nil
}
