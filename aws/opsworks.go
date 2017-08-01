package aws

import (
	"errors"
	"strings"

	libaws "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

var (
	errorOpsworksNoResult = errors.New("No instance")
)

type FindOpsworksIpResponse struct {
	Instances []FindOpsworksIpResponse_Instance
}

type FindOpsworksIpResponse_Instance struct {
	Stack     string
	Layer     string
	ID        string
	Name      string
	PublicIp  string
	PrivateIp string
}

func FindOpsworksIp(sessions Sessions, instanceName string) (*FindOpsworksIpResponse, error) {
	instances := []FindOpsworksIpResponse_Instance{}

	filters := []*ec2.Filter{}
	filters = append(filters, &ec2.Filter{
		Name:   libaws.String("tag:opsworks:instance"),
		Values: []*string{libaws.String(instanceName)},
	})

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
				instance := FindOpsworksIpResponse_Instance{
					ID:        *i.InstanceId,
					Name:      "unknown",
					PublicIp:  "unknown",
					PrivateIp: *i.PrivateIpAddress,
				}

				for _, tag := range i.Tags {
					k, v := *tag.Key, *tag.Value
					switch {
					case k == "opsworks:instance":
						instance.Name = v
					case k == "opsworks:stack":
						instance.Stack = v
					case strings.HasPrefix(k, "opsworks:layer:"):
						instance.Layer = v
					}
				}

				if i.PublicIpAddress != nil {
					instance.PublicIp = *i.PublicIpAddress
				}

				instances = append(instances, instance)
			}
		}
	}

	if len(instances) == 0 {
		return nil, errorOpsworksNoResult
	}

	response := &FindOpsworksIpResponse{
		Instances: instances,
	}
	return response, nil
}
