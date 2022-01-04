/*
Copyright 2021 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha4

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	expv1beta1 "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo converts this AzureManagedControlPlane to the Hub version (v1beta1).
func (src *AzureManagedControlPlane) ConvertTo(dstRaw conversion.Hub) error { // nolint
	dst := dstRaw.(*expv1beta1.AzureManagedControlPlane)

	if err := Convert_v1alpha4_AzureManagedControlPlane_To_v1beta1_AzureManagedControlPlane(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data from annotations
	restored := &expv1beta1.AzureManagedControlPlane{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// Handle special case for conversion of CIDRBlocks and Subnets to pointer.
	if restored.Spec.VirtualNetwork.CIDRBlocks == nil {
		dst.Spec.VirtualNetwork.CIDRBlocks = nil
	}
	if restored.Spec.VirtualNetwork.CIDRBlocks != nil && len(restored.Spec.VirtualNetwork.CIDRBlocks) == 0 {
		dst.Spec.VirtualNetwork.CIDRBlocks = []string{}
	}
	if restored.Spec.VirtualNetwork.Subnets == nil {
		dst.Spec.VirtualNetwork.Subnets = nil
	}
	if restored.Spec.VirtualNetwork.Subnets != nil && len(restored.Spec.VirtualNetwork.Subnets) == 0 {
		dst.Spec.VirtualNetwork.Subnets = []expv1beta1.ManagedControlPlaneSubnet{}
	}
	for i := range dst.Spec.VirtualNetwork.Subnets {
		if restored.Spec.VirtualNetwork.Subnets[i].CIDRBlocks == nil {
			dst.Spec.VirtualNetwork.Subnets[i].CIDRBlocks = nil
		}
		if restored.Spec.VirtualNetwork.Subnets[i].CIDRBlocks != nil && len(restored.Spec.VirtualNetwork.Subnets[i].CIDRBlocks) == 0 {
			dst.Spec.VirtualNetwork.Subnets[i].CIDRBlocks = []string{}
		}
	}

	return nil
}

// ConvertFrom converts from the Hub version (v1beta1) to this version.
func (dst *AzureManagedControlPlane) ConvertFrom(srcRaw conversion.Hub) error { // nolint
	src := srcRaw.(*expv1beta1.AzureManagedControlPlane)

	if err := Convert_v1beta1_AzureManagedControlPlane_To_v1alpha4_AzureManagedControlPlane(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion.
	if err := utilconversion.MarshalData(src, dst); err != nil {
		return err
	}

	return nil
}

// Convert_v1alpha4_ManagedControlPlaneVirtualNetwork_To_v1beta1_ManagedControlPlaneVirtualNetwork converts this ManagedControlPlaneVirtualNetwork to the Hub version (v1beta1).
func Convert_v1alpha4_ManagedControlPlaneVirtualNetwork_To_v1beta1_ManagedControlPlaneVirtualNetwork(in *ManagedControlPlaneVirtualNetwork, out *expv1beta1.ManagedControlPlaneVirtualNetwork, s apiconversion.Scope) error { // nolint
	out.Name = in.Name
	out.CIDRBlocks = []string{}
	subnet := &expv1beta1.ManagedControlPlaneSubnet{}

	if err := Convert_v1alpha4_CIDRBlock_To_v1beta1_CIDRBlocks(&in.CIDRBlock, &out.CIDRBlocks, s); err != nil {
		return err
	}
	if err := Convert_v1alpha4_ManagedControlPlaneSubnet_To_v1beta1_ManagedControlPlaneSubnet(&in.Subnet, subnet, s); err != nil {
		return err
	}
	out.Subnets = []expv1beta1.ManagedControlPlaneSubnet{*subnet}

	return nil
}

// Convert_v1beta1_ManagedControlPlaneVirtualNetwork_To_v1alpha4_ManagedControlPlaneVirtualNetwork converts from the Hub version (v1beta1) of the ManagedControlPlaneVirtualNetwork to this version.
func Convert_v1beta1_ManagedControlPlaneVirtualNetwork_To_v1alpha4_ManagedControlPlaneVirtualNetwork(in *expv1beta1.ManagedControlPlaneVirtualNetwork, out *ManagedControlPlaneVirtualNetwork, s apiconversion.Scope) error { // nolint
	out.Name = in.Name

	if len(in.CIDRBlocks) > 0 {
		if err := Convert_v1beta1_CIDRBlocks_To_v1alpha4_CIDRBlock(&in.CIDRBlocks, &out.CIDRBlock, s); err != nil {
			return err
		}
	}
	if len(in.Subnets) > 0 {
		if err := Convert_v1beta1_ManagedControlPlaneSubnet_To_v1alpha4_ManagedControlPlaneSubnet(&in.Subnets[0], &out.Subnet, s); err != nil {
			return err
		}
	}

	return nil
}

// Convert_v1alpha4_CIDRBlock_To_v1beta1_CIDRBlocks converts this CIDRBlock to the Hub version (v1beta1).
func Convert_v1alpha4_CIDRBlock_To_v1beta1_CIDRBlocks(in *string, out *[]string, s apiconversion.Scope) error { // nolint
	if in != nil {
		*out = []string{*in}
	}

	return nil
}

// Convert_v1beta1_CIDRBlocks_To_v1alpha4_CIDRBlock converts from the Hub version (v1beta1) of the CIDRBlocks to this version.
func Convert_v1beta1_CIDRBlocks_To_v1alpha4_CIDRBlock(in *[]string, out *string, s apiconversion.Scope) error { // nolint
	if len(*in) > 0 {
		*out = (*in)[0]
	}

	return nil
}

// Convert_v1alpha4_ManagedControlPlaneSubnet_To_v1beta1_ManagedControlPlaneSubnet converts this ManagedControlPlaneSubnet to the Hub version (v1beta1).
func Convert_v1alpha4_ManagedControlPlaneSubnet_To_v1beta1_ManagedControlPlaneSubnet(in *ManagedControlPlaneSubnet, out *expv1beta1.ManagedControlPlaneSubnet, s apiconversion.Scope) error { // nolint
	if in != nil {
		out.Name = in.Name
		if err := Convert_v1alpha4_CIDRBlock_To_v1beta1_CIDRBlocks(&in.CIDRBlock, &out.CIDRBlocks, s); err != nil {
			return err
		}
	}

	return nil
}

// Convert_v1beta1_ManagedControlPlaneSubnet_To_v1alpha4_ManagedControlPlaneSubnet converts from the Hub version (v1beta1) of the ManagedControlPlaneSubnet to this version.
func Convert_v1beta1_ManagedControlPlaneSubnet_To_v1alpha4_ManagedControlPlaneSubnet(in *expv1beta1.ManagedControlPlaneSubnet, out *ManagedControlPlaneSubnet, s apiconversion.Scope) error { // nolint
	if in != nil {
		out.Name = in.Name
		if err := Convert_v1beta1_CIDRBlocks_To_v1alpha4_CIDRBlock(&in.CIDRBlocks, &out.CIDRBlock, s); err != nil {
			return err
		}
	}

	return nil
}
