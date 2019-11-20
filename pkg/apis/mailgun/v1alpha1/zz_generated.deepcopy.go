// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailgunWebhook) DeepCopyInto(out *MailgunWebhook) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailgunWebhook.
func (in *MailgunWebhook) DeepCopy() *MailgunWebhook {
	if in == nil {
		return nil
	}
	out := new(MailgunWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MailgunWebhook) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailgunWebhookList) DeepCopyInto(out *MailgunWebhookList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MailgunWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailgunWebhookList.
func (in *MailgunWebhookList) DeepCopy() *MailgunWebhookList {
	if in == nil {
		return nil
	}
	out := new(MailgunWebhookList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MailgunWebhookList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailgunWebhookSpec) DeepCopyInto(out *MailgunWebhookSpec) {
	*out = *in
	if in.Opened != nil {
		in, out := &in.Opened, &out.Opened
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Clicked != nil {
		in, out := &in.Clicked, &out.Clicked
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailgunWebhookSpec.
func (in *MailgunWebhookSpec) DeepCopy() *MailgunWebhookSpec {
	if in == nil {
		return nil
	}
	out := new(MailgunWebhookSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MailgunWebhookStatus) DeepCopyInto(out *MailgunWebhookStatus) {
	*out = *in
	if in.Opened != nil {
		in, out := &in.Opened, &out.Opened
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Clicked != nil {
		in, out := &in.Clicked, &out.Clicked
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MailgunWebhookStatus.
func (in *MailgunWebhookStatus) DeepCopy() *MailgunWebhookStatus {
	if in == nil {
		return nil
	}
	out := new(MailgunWebhookStatus)
	in.DeepCopyInto(out)
	return out
}
