package aws

import (
	"github.com/rancher/rancher/tests/framework/clients/rancher"
	"github.com/rancher/rancher/tests/framework/extensions/cloudcredentials"
	"github.com/rancher/rancher/tests/framework/pkg/config"
)

const awsCloudCredNameBase = "awsCloudCredential"

func CreateAWSCloudCredentials(rancherClient *rancher.Client) (*cloudcredentials.CloudCredential, error) {
	var s3CredentialConfig cloudcredentials.S3CredentialConfig
	config.LoadConfig(cloudcredentials.S3CredentialConfigurationFileKey, &s3CredentialConfig)

	cloudCredential := cloudcredentials.CloudCredential{
		Name:               awsCloudCredNameBase,
		S3CredentialConfig: &s3CredentialConfig,
	}

	resp := &cloudcredentials.CloudCredential{}
	err := rancherClient.Management.APIBaseClient.Ops.DoCreate(cloudcredentials.CloudCredentialType, cloudCredential, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
