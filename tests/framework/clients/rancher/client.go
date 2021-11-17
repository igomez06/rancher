package rancher

import (
	"fmt"

	management "github.com/rancher/rancher/tests/framework/clients/rancher/generated/management/v3"
	provisioning "github.com/rancher/rancher/tests/framework/clients/rancher/provisioning"
	"github.com/rancher/rancher/tests/framework/pkg/clientbase"
	"github.com/rancher/rancher/tests/framework/pkg/session"
	"k8s.io/client-go/rest"
)

type Client struct {
	Management    *management.Client
	Provisioning  *provisioning.Client
	RancherConfig *Config
}

func NewClient(bearerToken string, rancherConfig *Config, session *session.Session) (*Client, error) {
	c := &Client{
		RancherConfig: rancherConfig,
	}

	var err error
	restConfig := newRestConfig(bearerToken, rancherConfig)
	c.Management, err = management.NewClient(clientOpts(restConfig, c.RancherConfig))
	if err != nil {
		return nil, err
	}

	c.Management.Ops.Session = session

	provClient, err := provisioning.NewForConfig(restConfig, session)
	if err != nil {
		return nil, err
	}

	c.Provisioning = provClient

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
