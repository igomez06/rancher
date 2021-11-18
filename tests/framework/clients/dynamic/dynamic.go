package dynamic

import (
	"context"

	"k8s.io/client-go/rest"

	"github.com/rancher/rancher/tests/framework/clients/rancher"
	"github.com/rancher/rancher/tests/framework/pkg/session"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

type Client struct {
	dynamic.Interface

	ts *session.Session
}

func NewForConfig(ts *session.Session, inConfig *rest.Config) (dynamic.Interface, error) {
	dynamicClient, err := dynamic.NewForConfig(inConfig)
	if err != nil {
		return nil, err
	}

	return &Client{
		Interface: dynamicClient,
		ts:        ts,
	}, nil
}

func (d *Client) Resource(resource schema.GroupVersionResource) dynamic.NamespaceableResourceInterface {
	return &NamespaceableResourceClient{
		NamespaceableResourceInterface: d.Interface.Resource(resource),
		ts:                             d.ts,
	}
}

type NamespaceableResourceClient struct {
	dynamic.NamespaceableResourceInterface
	ts *session.Session
}

func (d *NamespaceableResourceClient) Namespace(s string) dynamic.ResourceInterface {
	return &ResourceClient{
		ResourceInterface: d.NamespaceableResourceInterface.Namespace(s),
		ts:                d.ts,
	}
}

type ResourceClient struct {
	dynamic.ResourceInterface
	ts *session.Session
}

func (c *ResourceClient) Create(ctx context.Context, obj *unstructured.Unstructured, opts metav1.CreateOptions, subresources ...string) (*unstructured.Unstructured, error) {
	c.ts.RegisterCleanupFunc(func() error {
		err := c.Delete(context.TODO(), obj.GetName(), metav1.DeleteOptions{}, subresources...)
		if errors.IsNotFound(err) {
			return nil
		}

		return err
	})

	return c.ResourceInterface.Create(ctx, obj, opts, subresources...)
}

// GetRancherDynamicClient is a helper function that instantiates a dynamic client to communicate with the rancher host.
func GetRancherDynamicClient(bearerToken string, rancherConfig *rancher.Config, session *session.Session) (dynamic.Interface, error) {
	restConfig := &rest.Config{
		Host:        rancherConfig.RancherHost,
		BearerToken: bearerToken,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: *rancherConfig.Insecure,
			CAFile:   rancherConfig.CAFile,
		},
	}
	dynamic, err := NewForConfig(session, restConfig)
	if err != nil {
		return nil, err
	}
	return dynamic, nil
}
