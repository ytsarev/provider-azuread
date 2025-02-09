/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserInitParameters struct {

	// Whether or not the account should be enabled.
	// Whether or not the account should be enabled
	AccountEnabled *bool `json:"accountEnabled,omitempty" tf:"account_enabled,omitempty"`

	// The age group of the user. Supported values are Adult, NotAdult and Minor. Omit this property or specify a blank string to unset.
	// The age group of the user
	AgeGroup *string `json:"ageGroup,omitempty" tf:"age_group,omitempty"`

	// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
	// The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect
	BusinessPhones []*string `json:"businessPhones,omitempty" tf:"business_phones,omitempty"`

	// The city in which the user is located.
	// The city in which the user is located
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether consent has been obtained for minors. Supported values are Granted, Denied and NotRequired. Omit this property or specify a blank string to unset.
	// Whether consent has been obtained for minors
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty" tf:"consent_provided_for_minor,omitempty"`

	// The cost center associated with the user.
	// The cost center associated with the user.
	CostCenter *string `json:"costCenter,omitempty" tf:"cost_center,omitempty"`

	// The country/region in which the user is located, e.g. US or UK.
	// The country/region in which the user is located, e.g. `US` or `UK`
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The name for the department in which the user works.
	// The name for the department in which the user works
	Department *string `json:"department,omitempty" tf:"department,omitempty"`

	// Whether the user's password is exempt from expiring. Defaults to false.
	// Whether the users password is exempt from expiring
	DisablePasswordExpiration *bool `json:"disablePasswordExpiration,omitempty" tf:"disable_password_expiration,omitempty"`

	// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to false.
	// Whether the user is allowed weaker passwords than the default policy to be specified.
	DisableStrongPassword *bool `json:"disableStrongPassword,omitempty" tf:"disable_strong_password,omitempty"`

	// The name to display in the address book for the user.
	// The name to display in the address book for the user
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the division in which the user works.
	// The name of the division in which the user works.
	Division *string `json:"division,omitempty" tf:"division,omitempty"`

	// The employee identifier assigned to the user by the organisation.
	// The employee identifier assigned to the user by the organisation
	EmployeeID *string `json:"employeeId,omitempty" tf:"employee_id,omitempty"`

	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	EmployeeType *string `json:"employeeType,omitempty" tf:"employee_type,omitempty"`

	// The fax number of the user.
	// The fax number of the user
	FaxNumber *string `json:"faxNumber,omitempty" tf:"fax_number,omitempty"`

	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to false.
	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password
	ForcePasswordChange *bool `json:"forcePasswordChange,omitempty" tf:"force_password_change,omitempty"`

	// The given name (first name) of the user.
	// The given name (first name) of the user
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// The user’s job title.
	// The user’s job title
	JobTitle *string `json:"jobTitle,omitempty" tf:"job_title,omitempty"`

	// The SMTP address for the user. This property cannot be unset once specified.
	// The SMTP address for the user. Cannot be unset.
	Mail *string `json:"mail,omitempty" tf:"mail,omitempty"`

	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN)
	MailNickname *string `json:"mailNickname,omitempty" tf:"mail_nickname,omitempty"`

	// The object ID of the user's manager.
	// The object ID of the user's manager
	ManagerID *string `json:"managerId,omitempty" tf:"manager_id,omitempty"`

	// The primary cellular telephone number for the user.
	// The primary cellular telephone number for the user
	MobilePhone *string `json:"mobilePhone,omitempty" tf:"mobile_phone,omitempty"`

	// The office location in the user's place of business.
	// The office location in the user's place of business
	OfficeLocation *string `json:"officeLocation,omitempty" tf:"office_location,omitempty"`

	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's user_principal_name property when creating a new user account.
	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account
	OnpremisesImmutableID *string `json:"onpremisesImmutableId,omitempty" tf:"onpremises_immutable_id,omitempty"`

	// A list of additional email addresses for the user.
	// Additional email addresses for the user
	OtherMails []*string `json:"otherMails,omitempty" tf:"other_mails,omitempty"`

	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The user's preferred language, in ISO 639-1 notation.
	// The user's preferred language, in ISO 639-1 notation
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// Whether or not the Outlook global address list should include this user. Defaults to true.
	// Whether or not the Outlook global address list should include this user
	ShowInAddressList *bool `json:"showInAddressList,omitempty" tf:"show_in_address_list,omitempty"`

	// The state or province in the user's address.
	// The state or province in the user's address
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The street address of the user's place of business.
	// The street address of the user's place of business
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The user's surname (family name or last name).
	// The user's surname (family name or last name)
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: NO, JP, and GB. Cannot be reset to null once set.
	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set
	UsageLocation *string `json:"usageLocation,omitempty" tf:"usage_location,omitempty"`

	// The user principal name (UPN) of the user.
	// The user principal name (UPN) of the user
	UserPrincipalName *string `json:"userPrincipalName,omitempty" tf:"user_principal_name,omitempty"`
}

type UserObservation struct {

	// A freeform field for the user to describe themselves
	AboutMe *string `json:"aboutMe,omitempty" tf:"about_me,omitempty"`

	// Whether or not the account should be enabled.
	// Whether or not the account should be enabled
	AccountEnabled *bool `json:"accountEnabled,omitempty" tf:"account_enabled,omitempty"`

	// The age group of the user. Supported values are Adult, NotAdult and Minor. Omit this property or specify a blank string to unset.
	// The age group of the user
	AgeGroup *string `json:"ageGroup,omitempty" tf:"age_group,omitempty"`

	// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
	// The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect
	BusinessPhones []*string `json:"businessPhones,omitempty" tf:"business_phones,omitempty"`

	// The city in which the user is located.
	// The city in which the user is located
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether consent has been obtained for minors. Supported values are Granted, Denied and NotRequired. Omit this property or specify a blank string to unset.
	// Whether consent has been obtained for minors
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty" tf:"consent_provided_for_minor,omitempty"`

	// The cost center associated with the user.
	// The cost center associated with the user.
	CostCenter *string `json:"costCenter,omitempty" tf:"cost_center,omitempty"`

	// The country/region in which the user is located, e.g. US or UK.
	// The country/region in which the user is located, e.g. `US` or `UK`
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// Indicates whether the user account was created as a regular school or work account (null), an external account (Invitation), a local account for an Azure Active Directory B2C tenant (LocalAccount) or self-service sign-up using email verification (EmailVerified).
	// Indicates whether the user account was created as a regular school or work account (`null`), an external account (`Invitation`), a local account for an Azure Active Directory B2C tenant (`LocalAccount`) or self-service sign-up using email verification (`EmailVerified`)
	CreationType *string `json:"creationType,omitempty" tf:"creation_type,omitempty"`

	// The name for the department in which the user works.
	// The name for the department in which the user works
	Department *string `json:"department,omitempty" tf:"department,omitempty"`

	// Whether the user's password is exempt from expiring. Defaults to false.
	// Whether the users password is exempt from expiring
	DisablePasswordExpiration *bool `json:"disablePasswordExpiration,omitempty" tf:"disable_password_expiration,omitempty"`

	// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to false.
	// Whether the user is allowed weaker passwords than the default policy to be specified.
	DisableStrongPassword *bool `json:"disableStrongPassword,omitempty" tf:"disable_strong_password,omitempty"`

	// The name to display in the address book for the user.
	// The name to display in the address book for the user
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the division in which the user works.
	// The name of the division in which the user works.
	Division *string `json:"division,omitempty" tf:"division,omitempty"`

	// The employee identifier assigned to the user by the organisation.
	// The employee identifier assigned to the user by the organisation
	EmployeeID *string `json:"employeeId,omitempty" tf:"employee_id,omitempty"`

	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	EmployeeType *string `json:"employeeType,omitempty" tf:"employee_type,omitempty"`

	// For an external user invited to the tenant, this property represents the invited user's invitation status. Possible values are PendingAcceptance or Accepted.
	// For an external user invited to the tenant, this property represents the invited user's invitation status
	ExternalUserState *string `json:"externalUserState,omitempty" tf:"external_user_state,omitempty"`

	// The fax number of the user.
	// The fax number of the user
	FaxNumber *string `json:"faxNumber,omitempty" tf:"fax_number,omitempty"`

	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to false.
	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password
	ForcePasswordChange *bool `json:"forcePasswordChange,omitempty" tf:"force_password_change,omitempty"`

	// The given name (first name) of the user.
	// The given name (first name) of the user
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// A list of instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user.
	// The instant message voice over IP (VOIP) session initiation protocol (SIP) addresses for the user
	ImAddresses []*string `json:"imAddresses,omitempty" tf:"im_addresses,omitempty"`

	// The user’s job title.
	// The user’s job title
	JobTitle *string `json:"jobTitle,omitempty" tf:"job_title,omitempty"`

	// The SMTP address for the user. This property cannot be unset once specified.
	// The SMTP address for the user. Cannot be unset.
	Mail *string `json:"mail,omitempty" tf:"mail,omitempty"`

	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN)
	MailNickname *string `json:"mailNickname,omitempty" tf:"mail_nickname,omitempty"`

	// The object ID of the user's manager.
	// The object ID of the user's manager
	ManagerID *string `json:"managerId,omitempty" tf:"manager_id,omitempty"`

	// The primary cellular telephone number for the user.
	// The primary cellular telephone number for the user
	MobilePhone *string `json:"mobilePhone,omitempty" tf:"mobile_phone,omitempty"`

	// The object ID of the user.
	// The object ID of the user
	ObjectID *string `json:"objectId,omitempty" tf:"object_id,omitempty"`

	// The office location in the user's place of business.
	// The office location in the user's place of business
	OfficeLocation *string `json:"officeLocation,omitempty" tf:"office_location,omitempty"`

	// The on-premises distinguished name (DN) of the user, synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premise Active Directory distinguished name (DN) of the user
	OnpremisesDistinguishedName *string `json:"onpremisesDistinguishedName,omitempty" tf:"onpremises_distinguished_name,omitempty"`

	// The on-premises FQDN, also called dnsDomainName, synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premise FQDN (i.e. dnsDomainName) of the user
	OnpremisesDomainName *string `json:"onpremisesDomainName,omitempty" tf:"onpremises_domain_name,omitempty"`

	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's user_principal_name property when creating a new user account.
	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account
	OnpremisesImmutableID *string `json:"onpremisesImmutableId,omitempty" tf:"onpremises_immutable_id,omitempty"`

	// The on-premise SAM account name of the user.
	// The on-premise SAM account name of the user
	OnpremisesSamAccountName *string `json:"onpremisesSamAccountName,omitempty" tf:"onpremises_sam_account_name,omitempty"`

	// The on-premises security identifier (SID), synchronised from the on-premises directory when Azure AD Connect is used.
	// The on-premise security identifier (SID) of the user
	OnpremisesSecurityIdentifier *string `json:"onpremisesSecurityIdentifier,omitempty" tf:"onpremises_security_identifier,omitempty"`

	// Whether this user is synchronised from an on-premises directory (true), no longer synchronised (false), or has never been synchronised (null).
	// Whether this user is synchronized from an on-premises directory (true), no longer synchronized (false), or has never been synchronized (null)
	OnpremisesSyncEnabled *bool `json:"onpremisesSyncEnabled,omitempty" tf:"onpremises_sync_enabled,omitempty"`

	// The on-premise user principal name of the user.
	// The on-premise user principal name of the user
	OnpremisesUserPrincipalName *string `json:"onpremisesUserPrincipalName,omitempty" tf:"onpremises_user_principal_name,omitempty"`

	// A list of additional email addresses for the user.
	// Additional email addresses for the user
	OtherMails []*string `json:"otherMails,omitempty" tf:"other_mails,omitempty"`

	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The user's preferred language, in ISO 639-1 notation.
	// The user's preferred language, in ISO 639-1 notation
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// List of email addresses for the user that direct to the same mailbox.
	// Email addresses for the user that direct to the same mailbox
	ProxyAddresses []*string `json:"proxyAddresses,omitempty" tf:"proxy_addresses,omitempty"`

	// Whether or not the Outlook global address list should include this user. Defaults to true.
	// Whether or not the Outlook global address list should include this user
	ShowInAddressList *bool `json:"showInAddressList,omitempty" tf:"show_in_address_list,omitempty"`

	// The state or province in the user's address.
	// The state or province in the user's address
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The street address of the user's place of business.
	// The street address of the user's place of business
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The user's surname (family name or last name).
	// The user's surname (family name or last name)
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: NO, JP, and GB. Cannot be reset to null once set.
	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set
	UsageLocation *string `json:"usageLocation,omitempty" tf:"usage_location,omitempty"`

	// The user principal name (UPN) of the user.
	// The user principal name (UPN) of the user
	UserPrincipalName *string `json:"userPrincipalName,omitempty" tf:"user_principal_name,omitempty"`

	// The user type in the directory. Possible values are Guest or Member.
	// The user type in the directory. Possible values are `Guest` or `Member`
	UserType *string `json:"userType,omitempty" tf:"user_type,omitempty"`
}

type UserParameters struct {

	// Whether or not the account should be enabled.
	// Whether or not the account should be enabled
	// +kubebuilder:validation:Optional
	AccountEnabled *bool `json:"accountEnabled,omitempty" tf:"account_enabled,omitempty"`

	// The age group of the user. Supported values are Adult, NotAdult and Minor. Omit this property or specify a blank string to unset.
	// The age group of the user
	// +kubebuilder:validation:Optional
	AgeGroup *string `json:"ageGroup,omitempty" tf:"age_group,omitempty"`

	// A list of telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect.
	// The telephone numbers for the user. Only one number can be set for this property. Read-only for users synced with Azure AD Connect
	// +kubebuilder:validation:Optional
	BusinessPhones []*string `json:"businessPhones,omitempty" tf:"business_phones,omitempty"`

	// The city in which the user is located.
	// The city in which the user is located
	// +kubebuilder:validation:Optional
	City *string `json:"city,omitempty" tf:"city,omitempty"`

	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from.
	// The company name which the user is associated. This property can be useful for describing the company that an external user comes from
	// +kubebuilder:validation:Optional
	CompanyName *string `json:"companyName,omitempty" tf:"company_name,omitempty"`

	// Whether consent has been obtained for minors. Supported values are Granted, Denied and NotRequired. Omit this property or specify a blank string to unset.
	// Whether consent has been obtained for minors
	// +kubebuilder:validation:Optional
	ConsentProvidedForMinor *string `json:"consentProvidedForMinor,omitempty" tf:"consent_provided_for_minor,omitempty"`

	// The cost center associated with the user.
	// The cost center associated with the user.
	// +kubebuilder:validation:Optional
	CostCenter *string `json:"costCenter,omitempty" tf:"cost_center,omitempty"`

	// The country/region in which the user is located, e.g. US or UK.
	// The country/region in which the user is located, e.g. `US` or `UK`
	// +kubebuilder:validation:Optional
	Country *string `json:"country,omitempty" tf:"country,omitempty"`

	// The name for the department in which the user works.
	// The name for the department in which the user works
	// +kubebuilder:validation:Optional
	Department *string `json:"department,omitempty" tf:"department,omitempty"`

	// Whether the user's password is exempt from expiring. Defaults to false.
	// Whether the users password is exempt from expiring
	// +kubebuilder:validation:Optional
	DisablePasswordExpiration *bool `json:"disablePasswordExpiration,omitempty" tf:"disable_password_expiration,omitempty"`

	// Whether the user is allowed weaker passwords than the default policy to be specified. Defaults to false.
	// Whether the user is allowed weaker passwords than the default policy to be specified.
	// +kubebuilder:validation:Optional
	DisableStrongPassword *bool `json:"disableStrongPassword,omitempty" tf:"disable_strong_password,omitempty"`

	// The name to display in the address book for the user.
	// The name to display in the address book for the user
	// +kubebuilder:validation:Optional
	DisplayName *string `json:"displayName,omitempty" tf:"display_name,omitempty"`

	// The name of the division in which the user works.
	// The name of the division in which the user works.
	// +kubebuilder:validation:Optional
	Division *string `json:"division,omitempty" tf:"division,omitempty"`

	// The employee identifier assigned to the user by the organisation.
	// The employee identifier assigned to the user by the organisation
	// +kubebuilder:validation:Optional
	EmployeeID *string `json:"employeeId,omitempty" tf:"employee_id,omitempty"`

	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	// Captures enterprise worker type. For example, Employee, Contractor, Consultant, or Vendor.
	// +kubebuilder:validation:Optional
	EmployeeType *string `json:"employeeType,omitempty" tf:"employee_type,omitempty"`

	// The fax number of the user.
	// The fax number of the user
	// +kubebuilder:validation:Optional
	FaxNumber *string `json:"faxNumber,omitempty" tf:"fax_number,omitempty"`

	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password. Defaults to false.
	// Whether the user is forced to change the password during the next sign-in. Only takes effect when also changing the password
	// +kubebuilder:validation:Optional
	ForcePasswordChange *bool `json:"forcePasswordChange,omitempty" tf:"force_password_change,omitempty"`

	// The given name (first name) of the user.
	// The given name (first name) of the user
	// +kubebuilder:validation:Optional
	GivenName *string `json:"givenName,omitempty" tf:"given_name,omitempty"`

	// The user’s job title.
	// The user’s job title
	// +kubebuilder:validation:Optional
	JobTitle *string `json:"jobTitle,omitempty" tf:"job_title,omitempty"`

	// The SMTP address for the user. This property cannot be unset once specified.
	// The SMTP address for the user. Cannot be unset.
	// +kubebuilder:validation:Optional
	Mail *string `json:"mail,omitempty" tf:"mail,omitempty"`

	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN).
	// The mail alias for the user. Defaults to the user name part of the user principal name (UPN)
	// +kubebuilder:validation:Optional
	MailNickname *string `json:"mailNickname,omitempty" tf:"mail_nickname,omitempty"`

	// The object ID of the user's manager.
	// The object ID of the user's manager
	// +kubebuilder:validation:Optional
	ManagerID *string `json:"managerId,omitempty" tf:"manager_id,omitempty"`

	// The primary cellular telephone number for the user.
	// The primary cellular telephone number for the user
	// +kubebuilder:validation:Optional
	MobilePhone *string `json:"mobilePhone,omitempty" tf:"mobile_phone,omitempty"`

	// The office location in the user's place of business.
	// The office location in the user's place of business
	// +kubebuilder:validation:Optional
	OfficeLocation *string `json:"officeLocation,omitempty" tf:"office_location,omitempty"`

	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's user_principal_name property when creating a new user account.
	// The value used to associate an on-premise Active Directory user account with their Azure AD user object. This must be specified if you are using a federated domain for the user's `user_principal_name` property when creating a new user account
	// +kubebuilder:validation:Optional
	OnpremisesImmutableID *string `json:"onpremisesImmutableId,omitempty" tf:"onpremises_immutable_id,omitempty"`

	// A list of additional email addresses for the user.
	// Additional email addresses for the user
	// +kubebuilder:validation:Optional
	OtherMails []*string `json:"otherMails,omitempty" tf:"other_mails,omitempty"`

	// The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user.
	// The password for the user. The password must satisfy minimum requirements as specified by the password policy. The maximum length is 256 characters. This property is required when creating a new user
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code.
	// The postal code for the user's postal address. The postal code is specific to the user's country/region. In the United States of America, this attribute contains the ZIP code
	// +kubebuilder:validation:Optional
	PostalCode *string `json:"postalCode,omitempty" tf:"postal_code,omitempty"`

	// The user's preferred language, in ISO 639-1 notation.
	// The user's preferred language, in ISO 639-1 notation
	// +kubebuilder:validation:Optional
	PreferredLanguage *string `json:"preferredLanguage,omitempty" tf:"preferred_language,omitempty"`

	// Whether or not the Outlook global address list should include this user. Defaults to true.
	// Whether or not the Outlook global address list should include this user
	// +kubebuilder:validation:Optional
	ShowInAddressList *bool `json:"showInAddressList,omitempty" tf:"show_in_address_list,omitempty"`

	// The state or province in the user's address.
	// The state or province in the user's address
	// +kubebuilder:validation:Optional
	State *string `json:"state,omitempty" tf:"state,omitempty"`

	// The street address of the user's place of business.
	// The street address of the user's place of business
	// +kubebuilder:validation:Optional
	StreetAddress *string `json:"streetAddress,omitempty" tf:"street_address,omitempty"`

	// The user's surname (family name or last name).
	// The user's surname (family name or last name)
	// +kubebuilder:validation:Optional
	Surname *string `json:"surname,omitempty" tf:"surname,omitempty"`

	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: NO, JP, and GB. Cannot be reset to null once set.
	// The usage location of the user. Required for users that will be assigned licenses due to legal requirement to check for availability of services in countries. The usage location is a two letter country code (ISO standard 3166). Examples include: `NO`, `JP`, and `GB`. Cannot be reset to null once set
	// +kubebuilder:validation:Optional
	UsageLocation *string `json:"usageLocation,omitempty" tf:"usage_location,omitempty"`

	// The user principal name (UPN) of the user.
	// The user principal name (UPN) of the user
	// +kubebuilder:validation:Optional
	UserPrincipalName *string `json:"userPrincipalName,omitempty" tf:"user_principal_name,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
	// THIS IS AN ALPHA FIELD. Do not use it in production. It is not honored
	// unless the relevant Crossplane feature flag is enabled, and may be
	// changed or removed without notice.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// User is the Schema for the Users API.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,azuread}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.displayName) || has(self.initProvider.displayName)",message="displayName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.userPrincipalName) || has(self.initProvider.userPrincipalName)",message="userPrincipalName is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
