package dynamic

import (
	"context"
	"fmt"

	"k8s.io/client-go/rest"

	"github.com/rancher/rancher/tests/framework/pkg/session"
	"github.com/rancher/rancher/tests/framework/pkg/wait"
	"github.com/rancher/rancher/tests/integration/pkg/defaults"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/watch"
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
	unstructuredObj, err := c.ResourceInterface.Create(ctx, obj, opts, subresources...)
	if err != nil {
		return nil, err
	}

	c.ts.RegisterCleanupFunc(func() error {
		err := c.Delete(context.TODO(), unstructuredObj.GetName(), metav1.DeleteOptions{}, subresources...)
		if errors.IsNotFound(err) {
			return nil
		}

		return err
	})

	watchInterface, err := c.Watch(context.TODO(), metav1.ListOptions{
		FieldSelector:  "metadata.name=" + unstructuredObj.GetName(),
		TimeoutSeconds: &defaults.WatchTimeoutSeconds,
	})
	if err != nil {
		fmt.Println("error with dynamic watch interface")
		return nil, err
	}

	c.ts.RegisterWaitFunc(func() error {
		return wait.WatchWait(watchInterface, func(event watch.Event) (ready bool, err error) {
			if event.Type == watch.Error {
				return false, fmt.Errorf("there was an error deleting cluster")
			} else if event.Type == watch.Deleted {
				return true, nil
			} else if unstructuredObj.GroupVersionKind().Group == "rke-machine-config.cattle.io" {
				return true, nil
			}
			return false, nil
		})
	})

	return unstructuredObj, err
}
