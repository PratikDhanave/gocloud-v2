package gocloud

import (
	"errors"
	"fmt"
	"github.com/scorelab/gocloud/aws"
	"github.com/scorelab/gocloud/google"
	"github.com/scorelab/gocloud/auth"

)

//gocloud is a interface which hide differece between different cloud providers

type Gocloud interface {
	Createnode(request interface{}) (resp interface{}, err error)
	Startnode(request interface{}) (resp interface{}, err error)
	Stopnode(request interface{}) (resp interface{}, err error)
	Deletenode(request interface{}) (resp interface{}, err error)
	Rebootnode(request interface{}) (resp interface{}, err error)
}

const (
	Amazonprovider = "aws"
	Googleprovider = "google"
	Secretkey      = "dummy"
	Secretid       = "dummy"
)

//cloud provider return the instance of respepted cloud and map to the Gocloud so we can call the method like createnote on CloudProvider instance
//this is a delegation of cloud provider

func CloudProvider(provider, Secretkey, Secretid string) (Gocloud, error) {

	switch provider {
	case Amazonprovider:
		awsauth1 := auth.AWSauthbase{ClientID:Secretkey,SecretKey:Secretid}
		fmt.Println("I am awsauth1",awsauth1)
		amazon := new(aws.AWS)
		fmt.Println("I am amazon",amazon.EC2)
		fmt.Println("%V",amazon.EC2.Auth)
		amazon.EC2.Auth.AccessKey = Secretkey
		fmt.Println(amazon.EC2.Auth.AccessKey)
		return amazon, nil

	case Googleprovider:
		return new(google.Google), nil

	default:
		return nil, errors.New(fmt.Sprintf("provider %s not recognized\n", provider))
	}
}
