package machinepools

import (
	"context"

	apisV1 "github.com/rancher/rancher/pkg/apis/provisioning.cattle.io/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
)

const (
	AWSKind           = "Amazonec2Config"
	AWSPoolType       = "rke-machine-config.cattle.io.amazonec2config"
	AWSResourceConfig = "amazonec2configs"
)

func NewAWSMachineConfig(generatedPoolName, namespace, region string) *unstructured.Unstructured {
	machineConfig := &unstructured.Unstructured{}
	machineConfig.SetAPIVersion("rke-machine-config.cattle.io/v1")
	machineConfig.SetKind(AWSKind)
	machineConfig.SetGenerateName(generatedPoolName)
	machineConfig.SetNamespace(namespace)
	machineConfig.Object["region"] = region
	machineConfig.Object["instanceType"] = "t3a.medium"
	machineConfig.Object["sshUser"] = "ubuntu"
	machineConfig.Object["type"] = AWSPoolType
	machineConfig.Object["vpcId"] = "vpc-bfccf4d7"
	machineConfig.Object["volumeType"] = "gp2"
	machineConfig.Object["zone"] = "a"
	machineConfig.Object["retries"] = "5"
	machineConfig.Object["rootSize"] = "16"
	machineConfig.Object["securityGroup"] = []string{
		"rancher-nodes",
	}

	return machineConfig
}

func CreateMachineConfig(resource string, machinePoolConfig *unstructured.Unstructured, client dynamic.Interface) (*unstructured.Unstructured, error) {
	groupVersionResource := schema.GroupVersionResource{
		Group:    "rke-machine-config.cattle.io",
		Version:  "v1",
		Resource: resource,
	}

	ctx := context.Background()
	podResult, err := client.Resource(groupVersionResource).Namespace(machinePoolConfig.GetNamespace()).Create(ctx, machinePoolConfig, metav1.CreateOptions{})
	return podResult, err
}

func MachinePoolSetup(controlPlaneRole, etcdRole, workerRole bool, poolName string, quantity int32, machineConfig *unstructured.Unstructured) apisV1.RKEMachinePool {
	machineConfigRef := &corev1.ObjectReference{
		Kind: machineConfig.GetKind(),
		Name: machineConfig.GetName(),
	}

	return apisV1.RKEMachinePool{
		ControlPlaneRole: controlPlaneRole,
		EtcdRole:         etcdRole,
		WorkerRole:       workerRole,
		NodeConfig:       machineConfigRef,
		Name:             poolName,
		Quantity:         &quantity,
	}
}
