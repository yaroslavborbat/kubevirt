/*
This file is part of the KubeVirt project

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

Copyright The KubeVirt Authors.
*/

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1beta1 "kubevirt.io/api/snapshot/v1beta1"
	scheme "kubevirt.io/client-go/kubevirt/scheme"
)

// VirtualMachineSnapshotsGetter has a method to return a VirtualMachineSnapshotInterface.
// A group's client should implement this interface.
type VirtualMachineSnapshotsGetter interface {
	VirtualMachineSnapshots(namespace string) VirtualMachineSnapshotInterface
}

// VirtualMachineSnapshotInterface has methods to work with VirtualMachineSnapshot resources.
type VirtualMachineSnapshotInterface interface {
	Create(ctx context.Context, virtualMachineSnapshot *v1beta1.VirtualMachineSnapshot, opts v1.CreateOptions) (*v1beta1.VirtualMachineSnapshot, error)
	Update(ctx context.Context, virtualMachineSnapshot *v1beta1.VirtualMachineSnapshot, opts v1.UpdateOptions) (*v1beta1.VirtualMachineSnapshot, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, virtualMachineSnapshot *v1beta1.VirtualMachineSnapshot, opts v1.UpdateOptions) (*v1beta1.VirtualMachineSnapshot, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineSnapshot, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineSnapshotList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineSnapshot, err error)
	VirtualMachineSnapshotExpansion
}

// virtualMachineSnapshots implements VirtualMachineSnapshotInterface
type virtualMachineSnapshots struct {
	*gentype.ClientWithList[*v1beta1.VirtualMachineSnapshot, *v1beta1.VirtualMachineSnapshotList]
}

// newVirtualMachineSnapshots returns a VirtualMachineSnapshots
func newVirtualMachineSnapshots(c *SnapshotV1beta1Client, namespace string) *virtualMachineSnapshots {
	return &virtualMachineSnapshots{
		gentype.NewClientWithList[*v1beta1.VirtualMachineSnapshot, *v1beta1.VirtualMachineSnapshotList](
			"virtualmachinesnapshots",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.VirtualMachineSnapshot { return &v1beta1.VirtualMachineSnapshot{} },
			func() *v1beta1.VirtualMachineSnapshotList { return &v1beta1.VirtualMachineSnapshotList{} }),
	}
}
