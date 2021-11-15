package rancher

import (
	"fmt"

	"github.com/rancher/rancher/tests/framework/clients/rancher/management"
	"github.com/rancher/rancher/tests/framework/pkg/clientbase"
	"github.com/rancher/rancher/tests/framework/pkg/session"
	"k8s.io/client-go/rest"
)

type Client struct {
	Management    *management.Client
	RancherConfig *Config
}

func NewClient(bearerToken string, rancherConfig *Config, session *session.Session) (*Client, error) {
	c := &Client{
		RancherConfig: rancherConfig,
	}

	var err error
	c.Management, err = management.NewClient(clientOpts(newRestConfig(bearerToken, rancherConfig), c.RancherConfig), session)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func newRestConfig(bearerToken string, rancherConfig *Config) *rest.Config {
	return &rest.Config{
		Host:        rancherConfig.RancherHost,
		BearerToken: bearerToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: *rancherConfig.Insecure,
			CAFile:   rancherConfig.CAFile,
		},
	}
}

func clientOpts(restConfig *rest.Config, rancherConfig *Config) *clientbase.ClientOpts {
	return &clientbase.ClientOpts{
		URL:      fmt.Sprintf("https://%s/v3", rancherConfig.RancherHost),
		TokenKey: restConfig.BearerToken,
		Insecure: restConfig.Insecure,
		CACerts:  rancherConfig.CACerts,
	}
}
