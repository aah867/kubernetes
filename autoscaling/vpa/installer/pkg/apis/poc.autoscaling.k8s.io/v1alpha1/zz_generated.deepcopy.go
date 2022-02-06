// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerResourcePolicy) DeepCopyInto(out *ContainerResourcePolicy) {
	*out = *in
	if in.Mode != nil {
		in, out := &in.Mode, &out.Mode
		*out = new(ContainerScalingMode)
		**out = **in
	}
	if in.MinAllowed != nil {
		in, out := &in.MinAllowed, &out.MinAllowed
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.MaxAllowed != nil {
		in, out := &in.MaxAllowed, &out.MaxAllowed
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerResourcePolicy.
func (in *ContainerResourcePolicy) DeepCopy() *ContainerResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(ContainerResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HistogramCheckpoint) DeepCopyInto(out *HistogramCheckpoint) {
	*out = *in
	in.ReferenceTimestamp.DeepCopyInto(&out.ReferenceTimestamp)
	if in.BucketWeights != nil {
		in, out := &in.BucketWeights, &out.BucketWeights
		*out = make(map[int]uint32, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HistogramCheckpoint.
func (in *HistogramCheckpoint) DeepCopy() *HistogramCheckpoint {
	if in == nil {
		return nil
	}
	out := new(HistogramCheckpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodResourcePolicy) DeepCopyInto(out *PodResourcePolicy) {
	*out = *in
	if in.ContainerPolicies != nil {
		in, out := &in.ContainerPolicies, &out.ContainerPolicies
		*out = make([]ContainerResourcePolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodResourcePolicy.
func (in *PodResourcePolicy) DeepCopy() *PodResourcePolicy {
	if in == nil {
		return nil
	}
	out := new(PodResourcePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodUpdatePolicy) DeepCopyInto(out *PodUpdatePolicy) {
	*out = *in
	if in.UpdateMode != nil {
		in, out := &in.UpdateMode, &out.UpdateMode
		*out = new(UpdateMode)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodUpdatePolicy.
func (in *PodUpdatePolicy) DeepCopy() *PodUpdatePolicy {
	if in == nil {
		return nil
	}
	out := new(PodUpdatePolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendedContainerResources) DeepCopyInto(out *RecommendedContainerResources) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.LowerBound != nil {
		in, out := &in.LowerBound, &out.LowerBound
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	if in.UpperBound != nil {
		in, out := &in.UpperBound, &out.UpperBound
		*out = make(v1.ResourceList, len(*in))
		for key, val := range *in {
			(*out)[key] = val.DeepCopy()
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendedContainerResources.
func (in *RecommendedContainerResources) DeepCopy() *RecommendedContainerResources {
	if in == nil {
		return nil
	}
	out := new(RecommendedContainerResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RecommendedPodResources) DeepCopyInto(out *RecommendedPodResources) {
	*out = *in
	if in.ContainerRecommendations != nil {
		in, out := &in.ContainerRecommendations, &out.ContainerRecommendations
		*out = make([]RecommendedContainerResources, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RecommendedPodResources.
func (in *RecommendedPodResources) DeepCopy() *RecommendedPodResources {
	if in == nil {
		return nil
	}
	out := new(RecommendedPodResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscaler) DeepCopyInto(out *VerticalPodAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscaler.
func (in *VerticalPodAutoscaler) DeepCopy() *VerticalPodAutoscaler {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerticalPodAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerCheckpoint) DeepCopyInto(out *VerticalPodAutoscalerCheckpoint) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerCheckpoint.
func (in *VerticalPodAutoscalerCheckpoint) DeepCopy() *VerticalPodAutoscalerCheckpoint {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerCheckpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerticalPodAutoscalerCheckpoint) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerCheckpointList) DeepCopyInto(out *VerticalPodAutoscalerCheckpointList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerticalPodAutoscalerCheckpoint, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerCheckpointList.
func (in *VerticalPodAutoscalerCheckpointList) DeepCopy() *VerticalPodAutoscalerCheckpointList {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerCheckpointList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerticalPodAutoscalerCheckpointList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerCheckpointSpec) DeepCopyInto(out *VerticalPodAutoscalerCheckpointSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerCheckpointSpec.
func (in *VerticalPodAutoscalerCheckpointSpec) DeepCopy() *VerticalPodAutoscalerCheckpointSpec {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerCheckpointSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerCheckpointStatus) DeepCopyInto(out *VerticalPodAutoscalerCheckpointStatus) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.CPUHistogram.DeepCopyInto(&out.CPUHistogram)
	in.MemoryHistogram.DeepCopyInto(&out.MemoryHistogram)
	in.FirstSampleStart.DeepCopyInto(&out.FirstSampleStart)
	in.LastSampleStart.DeepCopyInto(&out.LastSampleStart)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerCheckpointStatus.
func (in *VerticalPodAutoscalerCheckpointStatus) DeepCopy() *VerticalPodAutoscalerCheckpointStatus {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerCheckpointStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerCondition) DeepCopyInto(out *VerticalPodAutoscalerCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerCondition.
func (in *VerticalPodAutoscalerCondition) DeepCopy() *VerticalPodAutoscalerCondition {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerList) DeepCopyInto(out *VerticalPodAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VerticalPodAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerList.
func (in *VerticalPodAutoscalerList) DeepCopy() *VerticalPodAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VerticalPodAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerSpec) DeepCopyInto(out *VerticalPodAutoscalerSpec) {
	*out = *in
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.UpdatePolicy != nil {
		in, out := &in.UpdatePolicy, &out.UpdatePolicy
		*out = new(PodUpdatePolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.ResourcePolicy != nil {
		in, out := &in.ResourcePolicy, &out.ResourcePolicy
		*out = new(PodResourcePolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerSpec.
func (in *VerticalPodAutoscalerSpec) DeepCopy() *VerticalPodAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VerticalPodAutoscalerStatus) DeepCopyInto(out *VerticalPodAutoscalerStatus) {
	*out = *in
	if in.Recommendation != nil {
		in, out := &in.Recommendation, &out.Recommendation
		*out = new(RecommendedPodResources)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]VerticalPodAutoscalerCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VerticalPodAutoscalerStatus.
func (in *VerticalPodAutoscalerStatus) DeepCopy() *VerticalPodAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(VerticalPodAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}
