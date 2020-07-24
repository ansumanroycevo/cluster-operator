// +build !ignore_autogenerated

/*
RabbitMQ Cluster Operator

Copyright 2020 VMware, Inc. All Rights Reserved.

This product is licensed to you under the Mozilla Public license, Version 2.0 (the "License").  You may not use this product except in compliance with the Mozilla Public License.

This product may include a number of subcomponents with separate copyright notices and license terms. Your use of these subcomponents is subject to the terms and conditions of the subcomponent's license, as noted in the LICENSE file.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/pivotal/rabbitmq-for-kubernetes/internal/status"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClientService) DeepCopyInto(out *ClientService) {
	*out = *in
	if in.EmbeddedLabelsAnnotations != nil {
		in, out := &in.EmbeddedLabelsAnnotations, &out.EmbeddedLabelsAnnotations
		*out = new(EmbeddedLabelsAnnotations)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClientService.
func (in *ClientService) DeepCopy() *ClientService {
	if in == nil {
		return nil
	}
	out := new(ClientService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedLabelsAnnotations) DeepCopyInto(out *EmbeddedLabelsAnnotations) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedLabelsAnnotations.
func (in *EmbeddedLabelsAnnotations) DeepCopy() *EmbeddedLabelsAnnotations {
	if in == nil {
		return nil
	}
	out := new(EmbeddedLabelsAnnotations)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmbeddedObjectMeta) DeepCopyInto(out *EmbeddedObjectMeta) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmbeddedObjectMeta.
func (in *EmbeddedObjectMeta) DeepCopy() *EmbeddedObjectMeta {
	if in == nil {
		return nil
	}
	out := new(EmbeddedObjectMeta)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PersistentVolumeClaim) DeepCopyInto(out *PersistentVolumeClaim) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.EmbeddedObjectMeta.DeepCopyInto(&out.EmbeddedObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PersistentVolumeClaim.
func (in *PersistentVolumeClaim) DeepCopy() *PersistentVolumeClaim {
	if in == nil {
		return nil
	}
	out := new(PersistentVolumeClaim)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodTemplateSpec) DeepCopyInto(out *PodTemplateSpec) {
	*out = *in
	if in.EmbeddedObjectMeta != nil {
		in, out := &in.EmbeddedObjectMeta, &out.EmbeddedObjectMeta
		*out = new(EmbeddedObjectMeta)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(v1.PodSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodTemplateSpec.
func (in *PodTemplateSpec) DeepCopy() *PodTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(PodTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqCluster) DeepCopyInto(out *RabbitmqCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqCluster.
func (in *RabbitmqCluster) DeepCopy() *RabbitmqCluster {
	if in == nil {
		return nil
	}
	out := new(RabbitmqCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterAdmin) DeepCopyInto(out *RabbitmqClusterAdmin) {
	*out = *in
	if in.SecretReference != nil {
		in, out := &in.SecretReference, &out.SecretReference
		*out = new(RabbitmqClusterSecretReference)
		(*in).DeepCopyInto(*out)
	}
	if in.ServiceReference != nil {
		in, out := &in.ServiceReference, &out.ServiceReference
		*out = new(RabbitmqClusterServiceReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterAdmin.
func (in *RabbitmqClusterAdmin) DeepCopy() *RabbitmqClusterAdmin {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterAdmin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterConfigurationSpec) DeepCopyInto(out *RabbitmqClusterConfigurationSpec) {
	*out = *in
	if in.AdditionalPlugins != nil {
		in, out := &in.AdditionalPlugins, &out.AdditionalPlugins
		*out = make([]Plugin, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterConfigurationSpec.
func (in *RabbitmqClusterConfigurationSpec) DeepCopy() *RabbitmqClusterConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterList) DeepCopyInto(out *RabbitmqClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RabbitmqCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterList.
func (in *RabbitmqClusterList) DeepCopy() *RabbitmqClusterList {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RabbitmqClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterOverrideSpec) DeepCopyInto(out *RabbitmqClusterOverrideSpec) {
	*out = *in
	if in.StatefulSet != nil {
		in, out := &in.StatefulSet, &out.StatefulSet
		*out = new(StatefulSet)
		(*in).DeepCopyInto(*out)
	}
	if in.ClientService != nil {
		in, out := &in.ClientService, &out.ClientService
		*out = new(ClientService)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterOverrideSpec.
func (in *RabbitmqClusterOverrideSpec) DeepCopy() *RabbitmqClusterOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterPersistenceSpec) DeepCopyInto(out *RabbitmqClusterPersistenceSpec) {
	*out = *in
	if in.StorageClassName != nil {
		in, out := &in.StorageClassName, &out.StorageClassName
		*out = new(string)
		**out = **in
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterPersistenceSpec.
func (in *RabbitmqClusterPersistenceSpec) DeepCopy() *RabbitmqClusterPersistenceSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterPersistenceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterSecretReference) DeepCopyInto(out *RabbitmqClusterSecretReference) {
	*out = *in
	if in.Keys != nil {
		in, out := &in.Keys, &out.Keys
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterSecretReference.
func (in *RabbitmqClusterSecretReference) DeepCopy() *RabbitmqClusterSecretReference {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterSecretReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterServiceReference) DeepCopyInto(out *RabbitmqClusterServiceReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterServiceReference.
func (in *RabbitmqClusterServiceReference) DeepCopy() *RabbitmqClusterServiceReference {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterServiceReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterServiceSpec) DeepCopyInto(out *RabbitmqClusterServiceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterServiceSpec.
func (in *RabbitmqClusterServiceSpec) DeepCopy() *RabbitmqClusterServiceSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterSpec) DeepCopyInto(out *RabbitmqClusterSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Service.DeepCopyInto(&out.Service)
	in.Persistence.DeepCopyInto(&out.Persistence)
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Rabbitmq.DeepCopyInto(&out.Rabbitmq)
	out.TLS = in.TLS
	in.Override.DeepCopyInto(&out.Override)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterSpec.
func (in *RabbitmqClusterSpec) DeepCopy() *RabbitmqClusterSpec {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RabbitmqClusterStatus) DeepCopyInto(out *RabbitmqClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]status.RabbitmqClusterCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Admin != nil {
		in, out := &in.Admin, &out.Admin
		*out = new(RabbitmqClusterAdmin)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RabbitmqClusterStatus.
func (in *RabbitmqClusterStatus) DeepCopy() *RabbitmqClusterStatus {
	if in == nil {
		return nil
	}
	out := new(RabbitmqClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSet) DeepCopyInto(out *StatefulSet) {
	*out = *in
	if in.EmbeddedLabelsAnnotations != nil {
		in, out := &in.EmbeddedLabelsAnnotations, &out.EmbeddedLabelsAnnotations
		*out = new(EmbeddedLabelsAnnotations)
		(*in).DeepCopyInto(*out)
	}
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(StatefulSetSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSet.
func (in *StatefulSet) DeepCopy() *StatefulSet {
	if in == nil {
		return nil
	}
	out := new(StatefulSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StatefulSetSpec) DeepCopyInto(out *StatefulSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		*out = new(PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.VolumeClaimTemplates != nil {
		in, out := &in.VolumeClaimTemplates, &out.VolumeClaimTemplates
		*out = make([]PersistentVolumeClaim, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UpdateStrategy != nil {
		in, out := &in.UpdateStrategy, &out.UpdateStrategy
		*out = new(appsv1.StatefulSetUpdateStrategy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StatefulSetSpec.
func (in *StatefulSetSpec) DeepCopy() *StatefulSetSpec {
	if in == nil {
		return nil
	}
	out := new(StatefulSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TLSSpec) DeepCopyInto(out *TLSSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TLSSpec.
func (in *TLSSpec) DeepCopy() *TLSSpec {
	if in == nil {
		return nil
	}
	out := new(TLSSpec)
	in.DeepCopyInto(out)
	return out
}
