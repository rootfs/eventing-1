// +build !ignore_autogenerated

/*
Copyright 2018 Google, Inc. All rights reserved.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bind) DeepCopyInto(out *Bind) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bind.
func (in *Bind) DeepCopy() *Bind {
	if in == nil {
		return nil
	}
	out := new(Bind)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Bind) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindAction) DeepCopyInto(out *BindAction) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindAction.
func (in *BindAction) DeepCopy() *BindAction {
	if in == nil {
		return nil
	}
	out := new(BindAction)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindCondition) DeepCopyInto(out *BindCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindCondition.
func (in *BindCondition) DeepCopy() *BindCondition {
	if in == nil {
		return nil
	}
	out := new(BindCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindList) DeepCopyInto(out *BindList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Bind, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindList.
func (in *BindList) DeepCopy() *BindList {
	if in == nil {
		return nil
	}
	out := new(BindList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BindList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindSpec) DeepCopyInto(out *BindSpec) {
	*out = *in
	out.Action = in.Action
	in.Trigger.DeepCopyInto(&out.Trigger)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindSpec.
func (in *BindSpec) DeepCopy() *BindSpec {
	if in == nil {
		return nil
	}
	out := new(BindSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BindStatus) DeepCopyInto(out *BindStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]BindCondition, len(*in))
		copy(*out, *in)
	}
	if in.BindContext != nil {
		in, out := &in.BindContext, &out.BindContext
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BindStatus.
func (in *BindStatus) DeepCopy() *BindStatus {
	if in == nil {
		return nil
	}
	out := new(BindStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSource) DeepCopyInto(out *EventSource) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSource.
func (in *EventSource) DeepCopy() *EventSource {
	if in == nil {
		return nil
	}
	out := new(EventSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSource) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceCondition) DeepCopyInto(out *EventSourceCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceCondition.
func (in *EventSourceCondition) DeepCopy() *EventSourceCondition {
	if in == nil {
		return nil
	}
	out := new(EventSourceCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceList) DeepCopyInto(out *EventSourceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceList.
func (in *EventSourceList) DeepCopy() *EventSourceList {
	if in == nil {
		return nil
	}
	out := new(EventSourceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventSourceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceSpec) DeepCopyInto(out *EventSourceSpec) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceSpec.
func (in *EventSourceSpec) DeepCopy() *EventSourceSpec {
	if in == nil {
		return nil
	}
	out := new(EventSourceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventSourceStatus) DeepCopyInto(out *EventSourceStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]EventSourceCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventSourceStatus.
func (in *EventSourceStatus) DeepCopy() *EventSourceStatus {
	if in == nil {
		return nil
	}
	out := new(EventSourceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTrigger) DeepCopyInto(out *EventTrigger) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.ParametersFrom != nil {
		in, out := &in.ParametersFrom, &out.ParametersFrom
		*out = make([]ParametersFromSource, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTrigger.
func (in *EventTrigger) DeepCopy() *EventTrigger {
	if in == nil {
		return nil
	}
	out := new(EventTrigger)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventType) DeepCopyInto(out *EventType) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventType.
func (in *EventType) DeepCopy() *EventType {
	if in == nil {
		return nil
	}
	out := new(EventType)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventType) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeCondition) DeepCopyInto(out *EventTypeCondition) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeCondition.
func (in *EventTypeCondition) DeepCopy() *EventTypeCondition {
	if in == nil {
		return nil
	}
	out := new(EventTypeCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeList) DeepCopyInto(out *EventTypeList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventType, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeList.
func (in *EventTypeList) DeepCopy() *EventTypeList {
	if in == nil {
		return nil
	}
	out := new(EventTypeList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventTypeList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeSpec) DeepCopyInto(out *EventTypeSpec) {
	*out = *in
	if in.SubscribeSchema != nil {
		in, out := &in.SubscribeSchema, &out.SubscribeSchema
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.EventSchema != nil {
		in, out := &in.EventSchema, &out.EventSchema
		if *in == nil {
			*out = nil
		} else {
			*out = new(runtime.RawExtension)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeSpec.
func (in *EventTypeSpec) DeepCopy() *EventTypeSpec {
	if in == nil {
		return nil
	}
	out := new(EventTypeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventTypeStatus) DeepCopyInto(out *EventTypeStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]EventTypeCondition, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventTypeStatus.
func (in *EventTypeStatus) DeepCopy() *EventTypeStatus {
	if in == nil {
		return nil
	}
	out := new(EventTypeStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ParametersFromSource) DeepCopyInto(out *ParametersFromSource) {
	*out = *in
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		if *in == nil {
			*out = nil
		} else {
			*out = new(SecretKeyReference)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ParametersFromSource.
func (in *ParametersFromSource) DeepCopy() *ParametersFromSource {
	if in == nil {
		return nil
	}
	out := new(ParametersFromSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyReference) DeepCopyInto(out *SecretKeyReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyReference.
func (in *SecretKeyReference) DeepCopy() *SecretKeyReference {
	if in == nil {
		return nil
	}
	out := new(SecretKeyReference)
	in.DeepCopyInto(out)
	return out
}
