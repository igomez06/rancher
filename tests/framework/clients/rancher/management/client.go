package management

import (
	"github.com/rancher/rancher/tests/framework/clients/rancher/cloudcredentials"
	managementClient "github.com/rancher/rancher/tests/framework/clients/rancher/generated/management/v3"
	"github.com/rancher/rancher/tests/framework/pkg/clientbase"
	"github.com/rancher/rancher/tests/framework/pkg/session"
)

type Client struct {
	*managementClient.Client
	CloudCredentialUpdated cloudcredentials.CloudCredentialOperations
}

func NewClient(opts *clientbase.ClientOpts, session *session.Session) (*Client, error) {
	client, err := managementClient.NewClient(opts)
	if err != nil {
		return nil, err
	}

	client.APIBaseClient.Ops.Session = session

	cloudClient, err := cloudcredentials.NewClient(opts, session)
	if err != nil {
		return nil, err
	}

	managementClient := &Client{
		client,
		cloudClient,
	}

	return managementClient, nil
}
