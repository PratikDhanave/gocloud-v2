package ec2

import(
	. "github.com/scorelab/gocloud/auth"
)
const (
	debug = false

	legacyAPIVersion = "2011-12-15"

	vpcAPIVersion = "2013-10-15"
)

//Authentication struct to store AccessKey and SecretKey

//Ec2 struct

type EC2 struct {
  AWSauthbase
	Region
	Auth
	Private byte
}

// Function return EC2 instance
func New2(aWSauthbase AWSauthbase,auth Auth, region Region) *EC2 {
	return &EC2{aWSauthbase,region,auth, 0}
}
