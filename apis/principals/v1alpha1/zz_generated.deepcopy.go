//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRolesObservation) DeepCopyInto(out *AppRolesObservation) {
	*out = *in
	if in.AllowedMemberTypes != nil {
		in, out := &in.AllowedMemberTypes, &out.AllowedMemberTypes
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRolesObservation.
func (in *AppRolesObservation) DeepCopy() *AppRolesObservation {
	if in == nil {
		return nil
	}
	out := new(AppRolesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AppRolesParameters) DeepCopyInto(out *AppRolesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AppRolesParameters.
func (in *AppRolesParameters) DeepCopy() *AppRolesParameters {
	if in == nil {
		return nil
	}
	out := new(AppRolesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTagsObservation) DeepCopyInto(out *FeatureTagsObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTagsObservation.
func (in *FeatureTagsObservation) DeepCopy() *FeatureTagsObservation {
	if in == nil {
		return nil
	}
	out := new(FeatureTagsObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureTagsParameters) DeepCopyInto(out *FeatureTagsParameters) {
	*out = *in
	if in.CustomSingleSignOn != nil {
		in, out := &in.CustomSingleSignOn, &out.CustomSingleSignOn
		*out = new(bool)
		**out = **in
	}
	if in.Enterprise != nil {
		in, out := &in.Enterprise, &out.Enterprise
		*out = new(bool)
		**out = **in
	}
	if in.Gallery != nil {
		in, out := &in.Gallery, &out.Gallery
		*out = new(bool)
		**out = **in
	}
	if in.Hide != nil {
		in, out := &in.Hide, &out.Hide
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureTagsParameters.
func (in *FeatureTagsParameters) DeepCopy() *FeatureTagsParameters {
	if in == nil {
		return nil
	}
	out := new(FeatureTagsParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesObservation) DeepCopyInto(out *FeaturesObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesObservation.
func (in *FeaturesObservation) DeepCopy() *FeaturesObservation {
	if in == nil {
		return nil
	}
	out := new(FeaturesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeaturesParameters) DeepCopyInto(out *FeaturesParameters) {
	*out = *in
	if in.CustomSingleSignOnApp != nil {
		in, out := &in.CustomSingleSignOnApp, &out.CustomSingleSignOnApp
		*out = new(bool)
		**out = **in
	}
	if in.EnterpriseApplication != nil {
		in, out := &in.EnterpriseApplication, &out.EnterpriseApplication
		*out = new(bool)
		**out = **in
	}
	if in.GalleryApplication != nil {
		in, out := &in.GalleryApplication, &out.GalleryApplication
		*out = new(bool)
		**out = **in
	}
	if in.VisibleToUsers != nil {
		in, out := &in.VisibleToUsers, &out.VisibleToUsers
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeaturesParameters.
func (in *FeaturesParameters) DeepCopy() *FeaturesParameters {
	if in == nil {
		return nil
	}
	out := new(FeaturesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Oauth2PermissionScopesObservation) DeepCopyInto(out *Oauth2PermissionScopesObservation) {
	*out = *in
	if in.AdminConsentDescription != nil {
		in, out := &in.AdminConsentDescription, &out.AdminConsentDescription
		*out = new(string)
		**out = **in
	}
	if in.AdminConsentDisplayName != nil {
		in, out := &in.AdminConsentDisplayName, &out.AdminConsentDisplayName
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.UserConsentDescription != nil {
		in, out := &in.UserConsentDescription, &out.UserConsentDescription
		*out = new(string)
		**out = **in
	}
	if in.UserConsentDisplayName != nil {
		in, out := &in.UserConsentDisplayName, &out.UserConsentDisplayName
		*out = new(string)
		**out = **in
	}
	if in.Value != nil {
		in, out := &in.Value, &out.Value
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Oauth2PermissionScopesObservation.
func (in *Oauth2PermissionScopesObservation) DeepCopy() *Oauth2PermissionScopesObservation {
	if in == nil {
		return nil
	}
	out := new(Oauth2PermissionScopesObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Oauth2PermissionScopesParameters) DeepCopyInto(out *Oauth2PermissionScopesParameters) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Oauth2PermissionScopesParameters.
func (in *Oauth2PermissionScopesParameters) DeepCopy() *Oauth2PermissionScopesParameters {
	if in == nil {
		return nil
	}
	out := new(Oauth2PermissionScopesParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Principal) DeepCopyInto(out *Principal) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Principal.
func (in *Principal) DeepCopy() *Principal {
	if in == nil {
		return nil
	}
	out := new(Principal)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Principal) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalList) DeepCopyInto(out *PrincipalList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Principal, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalList.
func (in *PrincipalList) DeepCopy() *PrincipalList {
	if in == nil {
		return nil
	}
	out := new(PrincipalList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrincipalList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalObservation) DeepCopyInto(out *PrincipalObservation) {
	*out = *in
	if in.AppRoleIds != nil {
		in, out := &in.AppRoleIds, &out.AppRoleIds
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.AppRoles != nil {
		in, out := &in.AppRoles, &out.AppRoles
		*out = make([]AppRolesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ApplicationTenantID != nil {
		in, out := &in.ApplicationTenantID, &out.ApplicationTenantID
		*out = new(string)
		**out = **in
	}
	if in.DisplayName != nil {
		in, out := &in.DisplayName, &out.DisplayName
		*out = new(string)
		**out = **in
	}
	if in.HomepageURL != nil {
		in, out := &in.HomepageURL, &out.HomepageURL
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.LogoutURL != nil {
		in, out := &in.LogoutURL, &out.LogoutURL
		*out = new(string)
		**out = **in
	}
	if in.Oauth2PermissionScopeIds != nil {
		in, out := &in.Oauth2PermissionScopeIds, &out.Oauth2PermissionScopeIds
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Oauth2PermissionScopes != nil {
		in, out := &in.Oauth2PermissionScopes, &out.Oauth2PermissionScopes
		*out = make([]Oauth2PermissionScopesObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ObjectID != nil {
		in, out := &in.ObjectID, &out.ObjectID
		*out = new(string)
		**out = **in
	}
	if in.RedirectUris != nil {
		in, out := &in.RedirectUris, &out.RedirectUris
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SAMLMetadataURL != nil {
		in, out := &in.SAMLMetadataURL, &out.SAMLMetadataURL
		*out = new(string)
		**out = **in
	}
	if in.ServicePrincipalNames != nil {
		in, out := &in.ServicePrincipalNames, &out.ServicePrincipalNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.SignInAudience != nil {
		in, out := &in.SignInAudience, &out.SignInAudience
		*out = new(string)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalObservation.
func (in *PrincipalObservation) DeepCopy() *PrincipalObservation {
	if in == nil {
		return nil
	}
	out := new(PrincipalObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalParameters) DeepCopyInto(out *PrincipalParameters) {
	*out = *in
	if in.AccountEnabled != nil {
		in, out := &in.AccountEnabled, &out.AccountEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AlternativeNames != nil {
		in, out := &in.AlternativeNames, &out.AlternativeNames
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.AppRoleAssignmentRequired != nil {
		in, out := &in.AppRoleAssignmentRequired, &out.AppRoleAssignmentRequired
		*out = new(bool)
		**out = **in
	}
	if in.ApplicationID != nil {
		in, out := &in.ApplicationID, &out.ApplicationID
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.FeatureTags != nil {
		in, out := &in.FeatureTags, &out.FeatureTags
		*out = make([]FeatureTagsParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Features != nil {
		in, out := &in.Features, &out.Features
		*out = make([]FeaturesParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LoginURL != nil {
		in, out := &in.LoginURL, &out.LoginURL
		*out = new(string)
		**out = **in
	}
	if in.Notes != nil {
		in, out := &in.Notes, &out.Notes
		*out = new(string)
		**out = **in
	}
	if in.NotificationEmailAddresses != nil {
		in, out := &in.NotificationEmailAddresses, &out.NotificationEmailAddresses
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.Owners != nil {
		in, out := &in.Owners, &out.Owners
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.PreferredSingleSignOnMode != nil {
		in, out := &in.PreferredSingleSignOnMode, &out.PreferredSingleSignOnMode
		*out = new(string)
		**out = **in
	}
	if in.SAMLSingleSignOn != nil {
		in, out := &in.SAMLSingleSignOn, &out.SAMLSingleSignOn
		*out = make([]SAMLSingleSignOnParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.UseExisting != nil {
		in, out := &in.UseExisting, &out.UseExisting
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalParameters.
func (in *PrincipalParameters) DeepCopy() *PrincipalParameters {
	if in == nil {
		return nil
	}
	out := new(PrincipalParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalSpec) DeepCopyInto(out *PrincipalSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalSpec.
func (in *PrincipalSpec) DeepCopy() *PrincipalSpec {
	if in == nil {
		return nil
	}
	out := new(PrincipalSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrincipalStatus) DeepCopyInto(out *PrincipalStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrincipalStatus.
func (in *PrincipalStatus) DeepCopy() *PrincipalStatus {
	if in == nil {
		return nil
	}
	out := new(PrincipalStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLSingleSignOnObservation) DeepCopyInto(out *SAMLSingleSignOnObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLSingleSignOnObservation.
func (in *SAMLSingleSignOnObservation) DeepCopy() *SAMLSingleSignOnObservation {
	if in == nil {
		return nil
	}
	out := new(SAMLSingleSignOnObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SAMLSingleSignOnParameters) DeepCopyInto(out *SAMLSingleSignOnParameters) {
	*out = *in
	if in.RelayState != nil {
		in, out := &in.RelayState, &out.RelayState
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SAMLSingleSignOnParameters.
func (in *SAMLSingleSignOnParameters) DeepCopy() *SAMLSingleSignOnParameters {
	if in == nil {
		return nil
	}
	out := new(SAMLSingleSignOnParameters)
	in.DeepCopyInto(out)
	return out
}
