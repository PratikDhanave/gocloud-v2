package ec2

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"
)

//Startnode ec2 instance accept array of instance-id.
func (ec2 *EC2) Startnode(request interface{}) (resp interface{}, err error) {
	ids := request.([]string)
	params := makeParams("StartInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &StartInstanceResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//Stopnode ec2 instance accept array of instance-id.
func (ec2 *EC2) Stopnode(request interface{}) (resp interface{}, err error) {
	ids := request.([]string)
	params := makeParams("StopInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &StopInstanceResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//Rebootnode ec2 instance accept array of instance-id.
func (ec2 *EC2) Rebootnode(request interface{}) (resp interface{}, err error) {
	ids := request.([]string)
	params := makeParams("RebootInstances")
	addParamsList(params, "InstanceId", ids)
	resp = &SimpleResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

//Deletenode ec2 instance accept array of instance-id.
func (ec2 *EC2) Deletenode(request interface{}) (resp interface{}, err error) {
	instIds := request.([]string)
	params := makeParams("TerminateInstances")
	addParamsList(params, "InstanceId", instIds)
	resp = &TerminateInstancesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}

//query pass the param to query and add signature to it base on secret key and acces key.
func (ec2 *EC2) query(params map[string]string, resp interface{}) error {

	req, err := http.NewRequest("GET", USEast.EC2Endpoint, nil)
	if err != nil {
		return err
	}

	// Add the params passed in to the query string
	query := req.URL.Query()
	for varName, varVal := range params {
		query.Add(varName, varVal)
	}
	query.Add("Timestamp", timeNow().In(time.UTC).Format(time.RFC3339))
	req.URL.RawQuery = query.Encode()

	auth := Auth{"dummy", "dummy"}

	SignV2(req, auth)

	fmt.Println(req)

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if debug {
		dump, _ := httputil.DumpResponse(r, true)
		log.Printf("response:\n")
		log.Printf("%v\n}\n", string(dump))
	}
	if r.StatusCode != 200 {
		return buildError(r)
	}

	fmt.Println(xml.NewDecoder(r.Body).Decode(resp))
	return xml.NewDecoder(r.Body).Decode(resp)

}

//Createnode Ec2 instances accept map[string]interface{} with attribute Define in EC2 documentation.
func (ec2 *EC2) Createnode(request interface{}) (resp interface{}, err error) {

	var options RunInstances

	param := make(map[string]interface{})

	param = request.(map[string]interface{})

	for key, value := range param {
		switch key {
		case "ImageId":
			ImageID, _ := value.(string)
			options.ImageID = ImageID

		case "MinCount":
			MinCount, _ := value.(int)
			options.MinCount = MinCount

		case "MaxCount":
			MaxCount, _ := value.(int)
			options.MaxCount = MaxCount

		case "KeyName":
			KeyName, _ := value.(string)
			options.KeyName = KeyName

		case "KernelID":
			KernelID, _ := value.(string)
			options.KernelID = KernelID

		case "InstanceType":
			InstanceType, _ := value.(string)
			options.InstanceType = InstanceType

		case "RamdiskID":
			RamdiskID, _ := value.(string)
			options.RamdiskID = RamdiskID

		case "AvailZone":
			AvailZone, _ := value.(string)
			options.AvailZone = AvailZone

		case "PlacementGroupName":
			PlacementGroupName, _ := value.(string)
			options.PlacementGroupName = PlacementGroupName

		case "Monitoring":
			Monitoring, _ := value.(bool)
			options.Monitoring = Monitoring

		case "SubnetID":
			SubnetID, _ := value.(string)
			options.SubnetID = SubnetID

		case "DisableAPITermination":
			DisableAPITermination, _ := value.(bool)
			options.DisableAPITermination = DisableAPITermination

		case "ShutdownBehavior":
			ShutdownBehavior, _ := value.(string)
			options.ShutdownBehavior = ShutdownBehavior

		case "PrivateIPAddress":
			PrivateIPAddress, _ := value.(string)
			options.PrivateIPAddress = PrivateIPAddress

		case "SecurityGroup":
			SecurityGroupparam, _ := value.([]map[string]string)
			//fmt.Println(SecurityGroupparam)
			for i := 0; i < len(SecurityGroupparam); i++ {
				var securityGroup SecurityGroup
				securityGroup.ID = SecurityGroupparam[i]["ID"]
				securityGroup.Name = SecurityGroupparam[i]["Name"]
				options.SecurityGroups = append(options.SecurityGroups, securityGroup)
			}

		case "BlockDevice":
			BlockDeviceparam, _ := value.([]map[string]interface{})
			for i := 0; i < len(BlockDeviceparam); i++ {
				var BlockDeviceMappingParam BlockDeviceMapping
				for BlockDeviceparamkey, BlockDeviceparamvalue := range BlockDeviceparam[i] {
					switch BlockDeviceparamkey {
					case "DeviceName":
						BlockDeviceMappingParam.DeviceName = BlockDeviceparamvalue.(string)

					case "VirtualName":
						BlockDeviceMappingParam.VirtualName = BlockDeviceparamvalue.(string)

					case "SnapshotID":
						BlockDeviceMappingParam.SnapshotID = BlockDeviceparamvalue.(string)

					case "VolumeType":
						BlockDeviceMappingParam.VolumeType = BlockDeviceparamvalue.(string)

					case "VolumeSize":
						BlockDeviceMappingParam.VolumeSize = BlockDeviceparamvalue.(int64)

					case "DeleteOnTermination":
						BlockDeviceMappingParam.DeleteOnTermination = BlockDeviceparamvalue.(bool)

					case "IOPS":
						BlockDeviceMappingParam.IOPS = BlockDeviceparamvalue.(int64)

					}
				}
				options.BlockDeviceMappings = append(options.BlockDeviceMappings, BlockDeviceMappingParam)

			}

		case "RunNetworkInterface":
			RunNetworkInterfaceparam, _ := value.([]map[string]interface{})
			//fmt.Println(RunNetworkInterfaceparam)
			var runNetworkInterface RunNetworkInterface
			for i := 0; i < len(RunNetworkInterfaceparam); i++ {
				//fmt.Println(RunNetworkInterfaceparam[i])
				for RunNetworkInterfaceparamkey, RunNetworkInterfaceparamvalue := range RunNetworkInterfaceparam[i] {
					switch RunNetworkInterfaceparamkey {
					case "ID":
						runNetworkInterface.ID = RunNetworkInterfaceparamvalue.(string)
					case "DeviceIndex":
						runNetworkInterface.DeviceIndex = RunNetworkInterfaceparamvalue.(int)
					case "SubnetID":
						runNetworkInterface.SubnetID = RunNetworkInterfaceparamvalue.(string)
					case "Description":
						runNetworkInterface.Description = RunNetworkInterfaceparamvalue.(string)
					case "DeleteOnTermination":
						runNetworkInterface.DeleteOnTermination = RunNetworkInterfaceparamvalue.(bool)
					case "SecondaryPrivateIPCount":
						runNetworkInterface.SecondaryPrivateIPCount = RunNetworkInterfaceparamvalue.(int)
					case "SecurityGroupIds":
						securityGroupIds, _ := RunNetworkInterfaceparamvalue.([]string)
						runNetworkInterface.SecurityGroupIds = securityGroupIds

					case "PrivateIPs":
						privateIPsParam, _ := RunNetworkInterfaceparamvalue.([]map[string]interface{})
						for i := 0; i < len(privateIPsParam); i++ {
							var privateIP PrivateIP
							for privateIPsParamkey, privateIPsParamvalue := range privateIPsParam[i] {
								switch privateIPsParamkey {
								case "Address":
									privateIP.Address = privateIPsParamvalue.(string)
								case "DNSName":
									privateIP.DNSName = privateIPsParamvalue.(string)
								case "IsPrimary":
									privateIP.IsPrimary = privateIPsParamvalue.(bool)
								}

							}
							runNetworkInterface.PrivateIPs = append(runNetworkInterface.PrivateIPs, privateIP)
						}

					}

				}
				options.NetworkInterfaces = append(options.NetworkInterfaces, runNetworkInterface)
			}
		case "UserData":
			options.UserData = value.([]byte)
		}
	}

	fmt.Println(options)

	params := makeParams("RunInstances")

	params["ImageId"] = options.ImageId

	params["InstanceType"] = options.InstanceType
	var min, max int
	if options.MinCount == 0 && options.MaxCount == 0 {
		min = 1
		max = 1
	} else if options.MaxCount == 0 {
		min = options.MinCount
		max = min
	} else {
		min = options.MinCount
		max = options.MaxCount
	}
	params["MinCount"] = strconv.Itoa(min)
	params["MaxCount"] = strconv.Itoa(max)
	i, j := 1, 1
	for _, g := range options.SecurityGroups {
		if g.Id != "" {
			params["SecurityGroupId."+strconv.Itoa(i)] = g.Id
			i++
		} else {
			params["SecurityGroup."+strconv.Itoa(j)] = g.Name
			j++
		}
	}
	prepareBlockDevices(params, options.BlockDeviceMappings)
	prepareNetworkInterfaces(params, options.NetworkInterfaces)
	token, err := clientToken()
	if err != nil {
		return nil, err
	}
	params["ClientToken"] = token

	if options.KeyName != "" {
		params["KeyName"] = options.KeyName
	}
	if options.KernelId != "" {
		params["KernelId"] = options.KernelId
	}
	if options.RamdiskId != "" {
		params["RamdiskId"] = options.RamdiskId
	}
	if options.UserData != nil {
		userData := make([]byte, base64.StdEncoding.EncodedLen(len(options.UserData)))
		base64.StdEncoding.Encode(userData, options.UserData)
		params["UserData"] = string(userData)
	}
	if options.AvailZone != "" {
		params["Placement.AvailabilityZone"] = options.AvailZone
	}
	if options.PlacementGroupName != "" {
		params["Placement.GroupName"] = options.PlacementGroupName
	}
	if options.Monitoring {
		params["Monitoring.Enabled"] = "true"
	}
	if options.SubnetId != "" {
		params["SubnetId"] = options.SubnetId
	}
	if options.DisableAPITermination {
		params["DisableApiTermination"] = "true"
	}
	if options.ShutdownBehavior != "" {
		params["InstanceInitiatedShutdownBehavior"] = options.ShutdownBehavior
	}
	if options.PrivateIPAddress != "" {
		params["PrivateIpAddress"] = options.PrivateIPAddress
	}

	resp = &RunInstancesResp{}
	err = ec2.query(params, resp)
	if err != nil {
		return nil, err
	}
	return
}
