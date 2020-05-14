// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPod) DeepCopyInto(out *EmulatedPod) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPod.
func (in *EmulatedPod) DeepCopy() *EmulatedPod {
	if in == nil {
		return nil
	}
	out := new(EmulatedPod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmulatedPod) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPodList) DeepCopyInto(out *EmulatedPodList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EmulatedPod, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPodList.
func (in *EmulatedPodList) DeepCopy() *EmulatedPodList {
	if in == nil {
		return nil
	}
	out := new(EmulatedPodList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EmulatedPodList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPodResourceUsage) DeepCopyInto(out *EmulatedPodResourceUsage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPodResourceUsage.
func (in *EmulatedPodResourceUsage) DeepCopy() *EmulatedPodResourceUsage {
	if in == nil {
		return nil
	}
	out := new(EmulatedPodResourceUsage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPodSpec) DeepCopyInto(out *EmulatedPodSpec) {
	*out = *in
	in.DirectState.DeepCopyInto(&out.DirectState)
	if in.Tasks != nil {
		in, out := &in.Tasks, &out.Tasks
		*out = make([]EmulatedPodTask, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPodSpec.
func (in *EmulatedPodSpec) DeepCopy() *EmulatedPodSpec {
	if in == nil {
		return nil
	}
	out := new(EmulatedPodSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPodState) DeepCopyInto(out *EmulatedPodState) {
	*out = *in
	if in.ResourceUsage != nil {
		in, out := &in.ResourceUsage, &out.ResourceUsage
		*out = new(EmulatedPodResourceUsage)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPodState.
func (in *EmulatedPodState) DeepCopy() *EmulatedPodState {
	if in == nil {
		return nil
	}
	out := new(EmulatedPodState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EmulatedPodTask) DeepCopyInto(out *EmulatedPodTask) {
	*out = *in
	in.State.DeepCopyInto(&out.State)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EmulatedPodTask.
func (in *EmulatedPodTask) DeepCopy() *EmulatedPodTask {
	if in == nil {
		return nil
	}
	out := new(EmulatedPodTask)
	in.DeepCopyInto(out)
	return out
}
