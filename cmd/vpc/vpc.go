package cmd

import (
	"encoding/json"
	"strconv"

	apiClient "fybe.com/cli/fybe/apiclient"
	"github.com/spf13/viper"
)

type PrivateNetwork struct {
	vpcId        int64    `json:"privateNetworkId"`
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	RegionName   string   `json:"regionName"`
	DataCenter   string   `json:"dataCenter"`
	Cidr         string   `json:"cidr"`
	AvailableIps int64    `json:"availableIps"`
	Instances    []string `json:"instances"`
}

func SetVpc(
	vpcId int64,
	Name string,
	Description string,
	RegionName string,
	DataCenter string,
	Cidr string,
	AvailableIps int64) *PrivateNetwork {
	vpc := new(PrivateNetwork)
	vpc.vpcId = vpcId
	vpc.Name = Name
	vpc.Description = Description
	vpc.RegionName = RegionName
	vpc.DataCenter = DataCenter
	vpc.Cidr = Cidr
	vpc.AvailableIps = AvailableIps
	vpc.Instances = []string{}
	return vpc
}

func mapInstancesToIds(data []apiClient.PrivateNetworkResponse) ([]byte, error) {
	if viper.GetString("output") != "json" && viper.GetString("output") != "yaml" {
		vpcData := data[0]
		var vpcs []*PrivateNetwork
		vpc := SetVpc(vpcData.PrivateNetworkId, vpcData.Name, vpcData.Description, vpcData.RegionName, vpcData.DataCenter, vpcData.Cidr, vpcData.AvailableIps)
		var len = len(vpcData.Instances)
		// if instance list has elements
		if len > 0 {
			var instanceIds []string
			for _, instance := range vpcData.Instances {
				instanceIds = append(instanceIds, strconv.FormatInt(instance.InstanceId, 10))
			}
			vpc.Instances = instanceIds
		}
		vpcs = append(vpcs, vpc)
		return json.Marshal(vpcs)
	} else {
		return json.Marshal(data)
	}
}
