package cloudcredentials

import (
	"github.com/rancher/rancher/tests/framework/pkg/clientbase"
	"github.com/rancher/rancher/tests/framework/pkg/session"
)

type Client struct {
	apiClient *clientbase.APIBaseClient
}

func NewClient(opts *clientbase.ClientOpts, session *session.Session) (*Client, error) {
	baseClient, err := clientbase.NewAPIClient(opts)
	if err != nil {
		return nil, err
	}

	client := &Client{
		apiClient: &baseClient,
	}
	client.apiClient.Ops.Session = session

	return client, nil
}
