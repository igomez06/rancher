package provisioning

import (
	"context"
	"fmt"

	apisV1 "github.com/rancher/rancher/pkg/apis/provisioning.cattle.io/v1"
	"github.com/rancher/rancher/tests/framework/pkg/wait"
	"github.com/rancher/rancher/tests/integration/pkg/defaults"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"
)

func (c *Cluster) Create(ctx context.Context, cluster *apisV1.Cluster, opts metav1.CreateOptions) (*apisV1.Cluster, error) {
	c.ts.RegisterCleanupFunc(func() error {
		err := c.Delete(context.TODO(), cluster.GetName(), metav1.DeleteOptions{})
		if errors.IsNotFound(err) {
			return nil
		}

		return err
	})
	watchInterface, err := c.Watch(context.TODO(), metav1.ListOptions{
		FieldSelector:  "metadata.name=" + cluster.GetName(),
		TimeoutSeconds: &defaults.WatchTimeoutSeconds,
	})
	if err != nil {
		return nil, err
	}

	c.ts.RegisterWaitFunc(func() error {
		return wait.WatchWait(watchInterface, func(event watch.Event) (ready bool, err error) {
			if event.Type == watch.Error {
				return false, fmt.Errorf("there was an error deleting cluster")
			} else if event.Type == watch.Deleted {
				return true, nil
			}
			return false, nil
		})
	})
	return c.ClusterInterface.Create(ctx, cluster, opts)
}
