package main

import (
	 "github.com/scorelab/gocloud/gocloud"
)

func main() {

	/*
	    InitializeParams := map[string]string{
	    "SourceImage" : "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
	    "DiskType" :  "1",
	    "DiskSizeGb" :"1",
	  }

	    disk := []map[string]interface{}{
	    {
	    "Type": "usb",
	    "Boot":  true,
	    "Mode": "",
	    "AutoDelete": false,
	    "deviceName": "",
	    "InitializeParams": InitializeParams,
	  },

	    {
	    "Type": "uss",
	    "Boot":  true,
	    "Mode": "",
	    "AutoDelete": false,
	    "deviceName": "",
	    "InitializeParams": InitializeParams,
	  },
	  }

	    AccessConfigs := []map[string]string{{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	    },{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	    }}

	  AccessConfigs2 := []map[string]string{{
	      "Name" : "external-nat",
	      "Type" :  "ONE_TO_ONE_NAT",
	      },}


	    NetworkInterfaces := []map[string]interface{}{{
	    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	    "Subnetwork": "",
	    "AccessConfigs": AccessConfigs,
	  },

	    {
	    "Network": "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
	    "Subnetwork": "hello",
	    "AccessConfigs": AccessConfigs2,
	  },
	  }

	  scheduling := map[string]interface{}{
	    "Preemptible": false,
	    "OnHostMaintenance": "MIGRATE",
	    "AutomaticRestart": true,
	  }


	    gce := map[string]interface{}{
	    "projectid":"sheltermap-1493101612061",
	    "Name": "scorelab",
	    "MachineType":"n1-standard-1",
	    "Zone":"us-east4-c",
	    "disk": disk,
	    "NetworkInterfaces": NetworkInterfaces,
	    "description":"",
	    "scheduling" : scheduling,
	    }
	*/

	/*
	   start := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//googlecloud, _ := gocloud.CloudProvider(gocloud.Googleprovider, gocloud.Secretkey, gocloud.Secretid)
	//  googlecloud.Startnode(start)

	/*
	   stop := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//  googlecloud.Stopnode(stop)

	/*
	   reboot := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",

	   }
	*/
	//googlecloud.Rebootnode(reboot)

	/*
	   delete := map[string]string{
	     "projectid":"sheltermap-1493101612061",
	     "instance":"instance-10",
	     "Zone": "us-west1-c",
	   }

	   googlecloud.Deletenode(delete)
	*/
	/*
	    ec2 := map[string]interface{}{
	      "ImageId"     : "ami-ccf405a5",
	      "InstanceType": "t1.micro",
	    }

	    amazoncloud, _ := CloudProvider(Amazonprovider,Secretkey,Secretid)
	    amazoncloud.Createnode(ec2)


	    reboot := []string{
	        "i-0ce8be9072b64fe84",
	    }

	   amazoncloud.Stopnode(stop)

	    start := []string{
	        "i-002ebb744a5fc48e0",
	    }
	    amazoncloud.Startnode(start)


	    amazoncloud.Deletenode(start)

	    amazoncloud.Rebootnode(reboot)
	*/

/*
	InitializeParams := map[string]string{
		"SourceImage": "https://www.googleapis.com/compute/v1/projects/debian-cloud/global/images/debian-8-jessie-v20160301",
		"DiskType":    "projects/sheltermap-1493101612061/zones/us-east4-c/diskTypes/pd-standard",
		"DiskSizeGb":  "10",
	}

	disk := []map[string]interface{}{
		{
			"Boot":             true,
			"AutoDelete":       false,
			"DeviceName":       "bokya",
			"Type":             "PERSISTENT",
			"Mode":             "READ_WRITE",
			"InitializeParams": InitializeParams,
		},
	}

	AccessConfigs := []map[string]string{{
		"Name": "external-nat",
		"Type": "ONE_TO_ONE_NAT",
	},
	}

	NetworkInterfaces := []map[string]interface{}{
		{
			"Network":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/global/networks/default",
			"Subnetwork":    "projects/sheltermap-1493101612061/regions/us-east4/subnetworks/default",
			"AccessConfigs": AccessConfigs,
		},
	}

	gce := map[string]interface{}{
		"projectid":         "sheltermap-1493101612061",
		"Name":              "bokya",
		"MachineType":       "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c/machineTypes/n1-standard-1",
		"Zone":              "https://www.googleapis.com/compute/v1/projects/sheltermap-1493101612061/zones/us-east4-c",
		"disk":              disk,
		"NetworkInterfaces": NetworkInterfaces,
	}

	googlecloud.Createnode(gce)
*/

	ec2 := map[string]interface{}{
 	 "ImageId"     : "ami-ccf405a5",
 	 "InstanceType": "t1.micr",
  }

  amazoncloud, _ := gocloud.CloudProvider(gocloud.Amazonprovider,gocloud.Secretkey,gocloud.Secretid)
  amazoncloud.Createnode(ec2)
}
