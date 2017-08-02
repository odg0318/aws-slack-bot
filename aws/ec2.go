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
	Name      string
	PublicIp  string
	PrivateIp string
	State     string
}

func FindEc2Ip(sessions Sessions, instanceId string, instanceName string) (*FindEc2IpResponse, error) {
	instances := []FindEc2IpResponse_Instance{}

	filters := []*ec2.Filter{}
	if len(instanceId) > 0 {
		filters = append(filters, &ec2.Filter{
			Name:   libaws.String("instance-id"),
			Values: []*string{libaws.String(instanceId)},
		})
	}

	if len(instanceName) > 0 {
		filters = append(filters, &ec2.Filter{
			Name:   libaws.String("tag:Name"),
			Values: []*string{libaws.String(instanceName)},
		})
	}

	input := &ec2.DescribeInstancesInput{
		Filters: filters,
	}

	for _, session := range sessions {
		client := ec2.New(session)

		output, err := client.DescribeInstances(input)
		if err != nil {
			return nil, err
		}

		for _, reservation := range output.Reservations {
			for _, i := range reservation.Instances {
				instance := FindEc2IpResponse_Instance{
					ID:        *i.InstanceId,
					Name:      "unknown",
					PublicIp:  "unknown",
					PrivateIp: "unknown",
					State:     "unknown",
				}

				for _, tag := range i.Tags {
					if *tag.Key != "Name" {
						continue
					}

					instance.Name = *tag.Value
				}

				if i.PublicIpAddress != nil {
					instance.PublicIp = *i.PublicIpAddress
				}

				if i.PrivateIpAddress != nil {
					instance.PrivateIp = *i.PrivateIpAddress
				}

				if i.State != nil {
					instance.State = *i.State.Name
				}

				instances = append(instances, instance)
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
