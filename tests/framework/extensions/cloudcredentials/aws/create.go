package aws

import (
	"github.com/rancher/rancher/tests/framework/clients/rancher"
	management "github.com/rancher/rancher/tests/framework/clients/rancher/generated/management/v3"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials"
	"github.com/rancher/rancher/tests/framework/pkg/config"
)

const awsCloudCredNameBase = "awsCloudCredential"

func CreateAWSCloudCredentials(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error) {
	var aazonEC2CredentialConfig cloudcredentials.AmazonEC2CredentialConfig
	config.LoadConfig(cloudcredentials.AmazonEC2CredentialConfigurationFileKey, &aazonEC2CredentialConfig)

	cloudCredential := cloudcredentials.CloudCredential{
		Name:                      awsCloudCredNameBase,
		AmazonEC2CredentialConfig: &aazonEC2CredentialConfig,
	}

	resp := &cloudcredentials.CloudCredential{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(management.CloudCredentialType, cloudCredential, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
