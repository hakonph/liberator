// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package nais_io_v1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicy) DeepCopyInto(out *AccessPolicy) {
	*out = *in
	if in.Inbound != nil {
		in, out := &in.Inbound, &out.Inbound
		*out = new(AccessPolicyInbound)
		(*in).DeepCopyInto(*out)
	}
	if in.Outbound != nil {
		in, out := &in.Outbound, &out.Outbound
		*out = new(AccessPolicyOutbound)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicy.
func (in *AccessPolicy) DeepCopy() *AccessPolicy {
	if in == nil {
		return nil
	}
	out := new(AccessPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyExternalRule) DeepCopyInto(out *AccessPolicyExternalRule) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]AccessPolicyPortRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyExternalRule.
func (in *AccessPolicyExternalRule) DeepCopy() *AccessPolicyExternalRule {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyExternalRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyInbound) DeepCopyInto(out *AccessPolicyInbound) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AccessPolicyRule, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyInbound.
func (in *AccessPolicyInbound) DeepCopy() *AccessPolicyInbound {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyInbound)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyOutbound) DeepCopyInto(out *AccessPolicyOutbound) {
	*out = *in
	if in.Rules != nil {
		in, out := &in.Rules, &out.Rules
		*out = make([]AccessPolicyRule, len(*in))
		copy(*out, *in)
	}
	if in.External != nil {
		in, out := &in.External, &out.External
		*out = make([]AccessPolicyExternalRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyOutbound.
func (in *AccessPolicyOutbound) DeepCopy() *AccessPolicyOutbound {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyOutbound)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyPortRule) DeepCopyInto(out *AccessPolicyPortRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyPortRule.
func (in *AccessPolicyPortRule) DeepCopy() *AccessPolicyPortRule {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyPortRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccessPolicyRule) DeepCopyInto(out *AccessPolicyRule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccessPolicyRule.
func (in *AccessPolicyRule) DeepCopy() *AccessPolicyRule {
	if in == nil {
		return nil
	}
	out := new(AccessPolicyRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Alert) DeepCopyInto(out *Alert) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Alert.
func (in *Alert) DeepCopy() *Alert {
	if in == nil {
		return nil
	}
	out := new(Alert)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Alert) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertList) DeepCopyInto(out *AlertList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Alert, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertList.
func (in *AlertList) DeepCopy() *AlertList {
	if in == nil {
		return nil
	}
	out := new(AlertList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AlertList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertSpec) DeepCopyInto(out *AlertSpec) {
	*out = *in
	out.Route = in.Route
	out.Receivers = in.Receivers
	if in.Alerts != nil {
		in, out := &in.Alerts, &out.Alerts
		*out = make([]Rule, len(*in))
		copy(*out, *in)
	}
	if in.InhibitRules != nil {
		in, out := &in.InhibitRules, &out.InhibitRules
		*out = make([]InhibitRules, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertSpec.
func (in *AlertSpec) DeepCopy() *AlertSpec {
	if in == nil {
		return nil
	}
	out := new(AlertSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AlertStatus) DeepCopyInto(out *AlertStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AlertStatus.
func (in *AlertStatus) DeepCopy() *AlertStatus {
	if in == nil {
		return nil
	}
	out := new(AlertStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdApplication) DeepCopyInto(out *AzureAdApplication) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdApplication.
func (in *AzureAdApplication) DeepCopy() *AzureAdApplication {
	if in == nil {
		return nil
	}
	out := new(AzureAdApplication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureAdApplication) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdApplicationList) DeepCopyInto(out *AzureAdApplicationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]AzureAdApplication, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdApplicationList.
func (in *AzureAdApplicationList) DeepCopy() *AzureAdApplicationList {
	if in == nil {
		return nil
	}
	out := new(AzureAdApplicationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *AzureAdApplicationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdApplicationSpec) DeepCopyInto(out *AzureAdApplicationSpec) {
	*out = *in
	if in.ReplyUrls != nil {
		in, out := &in.ReplyUrls, &out.ReplyUrls
		*out = make([]AzureAdReplyUrl, len(*in))
		copy(*out, *in)
	}
	if in.PreAuthorizedApplications != nil {
		in, out := &in.PreAuthorizedApplications, &out.PreAuthorizedApplications
		*out = make([]AccessPolicyRule, len(*in))
		copy(*out, *in)
	}
	if in.Claims != nil {
		in, out := &in.Claims, &out.Claims
		*out = new(AzureAdClaims)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdApplicationSpec.
func (in *AzureAdApplicationSpec) DeepCopy() *AzureAdApplicationSpec {
	if in == nil {
		return nil
	}
	out := new(AzureAdApplicationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdApplicationStatus) DeepCopyInto(out *AzureAdApplicationStatus) {
	*out = *in
	if in.SynchronizationTime != nil {
		in, out := &in.SynchronizationTime, &out.SynchronizationTime
		*out = (*in).DeepCopy()
	}
	if in.PasswordKeyIds != nil {
		in, out := &in.PasswordKeyIds, &out.PasswordKeyIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CertificateKeyIds != nil {
		in, out := &in.CertificateKeyIds, &out.CertificateKeyIds
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdApplicationStatus.
func (in *AzureAdApplicationStatus) DeepCopy() *AzureAdApplicationStatus {
	if in == nil {
		return nil
	}
	out := new(AzureAdApplicationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdClaims) DeepCopyInto(out *AzureAdClaims) {
	*out = *in
	if in.Extra != nil {
		in, out := &in.Extra, &out.Extra
		*out = make([]AzureAdExtraClaim, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdClaims.
func (in *AzureAdClaims) DeepCopy() *AzureAdClaims {
	if in == nil {
		return nil
	}
	out := new(AzureAdClaims)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureAdReplyUrl) DeepCopyInto(out *AzureAdReplyUrl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureAdReplyUrl.
func (in *AzureAdReplyUrl) DeepCopy() *AzureAdReplyUrl {
	if in == nil {
		return nil
	}
	out := new(AzureAdReplyUrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Email) DeepCopyInto(out *Email) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Email.
func (in *Email) DeepCopy() *Email {
	if in == nil {
		return nil
	}
	out := new(Email)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPortenClient) DeepCopyInto(out *IDPortenClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPortenClient.
func (in *IDPortenClient) DeepCopy() *IDPortenClient {
	if in == nil {
		return nil
	}
	out := new(IDPortenClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDPortenClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPortenClientList) DeepCopyInto(out *IDPortenClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]IDPortenClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPortenClientList.
func (in *IDPortenClientList) DeepCopy() *IDPortenClientList {
	if in == nil {
		return nil
	}
	out := new(IDPortenClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *IDPortenClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPortenClientSpec) DeepCopyInto(out *IDPortenClientSpec) {
	*out = *in
	if in.PostLogoutRedirectURIs != nil {
		in, out := &in.PostLogoutRedirectURIs, &out.PostLogoutRedirectURIs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SessionLifetime != nil {
		in, out := &in.SessionLifetime, &out.SessionLifetime
		*out = new(int)
		**out = **in
	}
	if in.AccessTokenLifetime != nil {
		in, out := &in.AccessTokenLifetime, &out.AccessTokenLifetime
		*out = new(int)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPortenClientSpec.
func (in *IDPortenClientSpec) DeepCopy() *IDPortenClientSpec {
	if in == nil {
		return nil
	}
	out := new(IDPortenClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IDPortenClientStatus) DeepCopyInto(out *IDPortenClientStatus) {
	*out = *in
	if in.SynchronizationTime != nil {
		in, out := &in.SynchronizationTime, &out.SynchronizationTime
		*out = (*in).DeepCopy()
	}
	if in.KeyIDs != nil {
		in, out := &in.KeyIDs, &out.KeyIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IDPortenClientStatus.
func (in *IDPortenClientStatus) DeepCopy() *IDPortenClientStatus {
	if in == nil {
		return nil
	}
	out := new(IDPortenClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InhibitRules) DeepCopyInto(out *InhibitRules) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.TargetsRegex != nil {
		in, out := &in.TargetsRegex, &out.TargetsRegex
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Sources != nil {
		in, out := &in.Sources, &out.Sources
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SourcesRegex != nil {
		in, out := &in.SourcesRegex, &out.SourcesRegex
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InhibitRules.
func (in *InhibitRules) DeepCopy() *InhibitRules {
	if in == nil {
		return nil
	}
	out := new(InhibitRules)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jwker) DeepCopyInto(out *Jwker) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jwker.
func (in *Jwker) DeepCopy() *Jwker {
	if in == nil {
		return nil
	}
	out := new(Jwker)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Jwker) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwkerList) DeepCopyInto(out *JwkerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Jwker, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwkerList.
func (in *JwkerList) DeepCopy() *JwkerList {
	if in == nil {
		return nil
	}
	out := new(JwkerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *JwkerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwkerSpec) DeepCopyInto(out *JwkerSpec) {
	*out = *in
	if in.AccessPolicy != nil {
		in, out := &in.AccessPolicy, &out.AccessPolicy
		*out = new(AccessPolicy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwkerSpec.
func (in *JwkerSpec) DeepCopy() *JwkerSpec {
	if in == nil {
		return nil
	}
	out := new(JwkerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JwkerStatus) DeepCopyInto(out *JwkerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JwkerStatus.
func (in *JwkerStatus) DeepCopy() *JwkerStatus {
	if in == nil {
		return nil
	}
	out := new(JwkerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaskinportenClient) DeepCopyInto(out *MaskinportenClient) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaskinportenClient.
func (in *MaskinportenClient) DeepCopy() *MaskinportenClient {
	if in == nil {
		return nil
	}
	out := new(MaskinportenClient)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaskinportenClient) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaskinportenClientList) DeepCopyInto(out *MaskinportenClientList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MaskinportenClient, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaskinportenClientList.
func (in *MaskinportenClientList) DeepCopy() *MaskinportenClientList {
	if in == nil {
		return nil
	}
	out := new(MaskinportenClientList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MaskinportenClientList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaskinportenClientSpec) DeepCopyInto(out *MaskinportenClientSpec) {
	*out = *in
	if in.Scopes != nil {
		in, out := &in.Scopes, &out.Scopes
		*out = make([]MaskinportenScope, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaskinportenClientSpec.
func (in *MaskinportenClientSpec) DeepCopy() *MaskinportenClientSpec {
	if in == nil {
		return nil
	}
	out := new(MaskinportenClientSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaskinportenClientStatus) DeepCopyInto(out *MaskinportenClientStatus) {
	*out = *in
	if in.SynchronizationTime != nil {
		in, out := &in.SynchronizationTime, &out.SynchronizationTime
		*out = (*in).DeepCopy()
	}
	if in.KeyIDs != nil {
		in, out := &in.KeyIDs, &out.KeyIDs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaskinportenClientStatus.
func (in *MaskinportenClientStatus) DeepCopy() *MaskinportenClientStatus {
	if in == nil {
		return nil
	}
	out := new(MaskinportenClientStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MaskinportenScope) DeepCopyInto(out *MaskinportenScope) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MaskinportenScope.
func (in *MaskinportenScope) DeepCopy() *MaskinportenScope {
	if in == nil {
		return nil
	}
	out := new(MaskinportenScope)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pushover) DeepCopyInto(out *Pushover) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pushover.
func (in *Pushover) DeepCopy() *Pushover {
	if in == nil {
		return nil
	}
	out := new(Pushover)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Receivers) DeepCopyInto(out *Receivers) {
	*out = *in
	out.Slack = in.Slack
	out.Email = in.Email
	out.SMS = in.SMS
	out.Pushover = in.Pushover
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Receivers.
func (in *Receivers) DeepCopy() *Receivers {
	if in == nil {
		return nil
	}
	out := new(Receivers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Route) DeepCopyInto(out *Route) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Route.
func (in *Route) DeepCopy() *Route {
	if in == nil {
		return nil
	}
	out := new(Route)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Rule) DeepCopyInto(out *Rule) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Rule.
func (in *Rule) DeepCopy() *Rule {
	if in == nil {
		return nil
	}
	out := new(Rule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SMS) DeepCopyInto(out *SMS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SMS.
func (in *SMS) DeepCopy() *SMS {
	if in == nil {
		return nil
	}
	out := new(SMS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Slack) DeepCopyInto(out *Slack) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Slack.
func (in *Slack) DeepCopy() *Slack {
	if in == nil {
		return nil
	}
	out := new(Slack)
	in.DeepCopyInto(out)
	return out
}
