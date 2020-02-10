/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for service: Upgrade.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package deployment

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
)



// The ``VcsaEmbeddedSpec`` class contains information used to upgrade a Embedded vCenter Server appliance. This class was added in vSphere API 6.7.
type UpgradeVcsaEmbeddedSpec struct {
    // Customer experience improvement program should be enabled or disabled for this embedded vCenter Server upgrade. This property was added in vSphere API 6.7.
	CeipEnabled bool
}

// The ``PscSpec`` class contains information used to upgrade a Platform Service Controller appliance. This class was added in vSphere API 6.7.
type UpgradePscSpec struct {
    // Customer experience improvement program should be enabled or disabled for this Platform Services Controller upgrade. This property was added in vSphere API 6.7.
	CeipEnabled bool
}

// The ``SourceApplianceSpec`` class contains information used to connect to the appliance used as the source for an upgrade. This class was added in vSphere API 6.7.
type UpgradeSourceApplianceSpec struct {
    // The IP address or DNS resolvable name of the source appliance. This property was added in vSphere API 6.7.
	Hostname string
    // The HTTPS port of the source appliance. This property was added in vSphere API 6.7.
	HttpsPort *int64
    // SHA1 thumbprint of the server SSL certificate will be used for verification. This property was added in vSphere API 6.7.
	SslThumbprint *string
    // SSL verification should be enabled or disabled for the source appliance validations. By default it is enabled and will use SSL certificate for verification. If thumbprint is provided, will use thumbprint for the verification. This property was added in vSphere API 6.7.
	SslVerify *bool
    // The SSO administrator account on the source appliance. This property was added in vSphere API 6.7.
	SsoAdminUsername string
    // The SSO administrator account password. This property was added in vSphere API 6.7.
	SsoAdminPassword string
    // The password of the root user on the source appliance. This property was added in vSphere API 6.7.
	RootPassword string
    // Appliance SSH verification should be enabled or disabled. By default it is disabled and will not use any verification. If thumbprint is provided, thumbprint verification will be performed. This property was added in vSphere API 6.7.
	SshVerify *bool
    // MD5 thumbprint of the server SSH key will be used for verification. This property was added in vSphere API 6.7.
	SshThumbprint *string
}

// The ``UpgradeSpec`` class contains information used to configure the appliance upgrade. This class was added in vSphere API 6.7.
type UpgradeUpgradeSpec struct {
    // Source appliance spec. This property was added in vSphere API 6.7.
	SourceAppliance UpgradeSourceApplianceSpec
    // Source location spec. This property was added in vSphere API 6.7.
	SourceLocation LocationSpec
    // Determines how vCenter history will be migrated during the upgrade process. vCenter history consists of: 
    //
    // * Statistics
    // * Events
    // * Tasks
    //
    //  By default only core data will be migrated. Use this spec to define which part of vCenter history data will be migrated and when. This property was added in vSphere API 6.7.
	History *HistoryMigrationSpec
    // Information that are specific to this embedded vCenter Server. This property was added in vSphere API 6.7.
	VcsaEmbedded *UpgradeVcsaEmbeddedSpec
    // Information that are specific to this Platform Services Controller. This property was added in vSphere API 6.7.
	Psc *UpgradePscSpec
    // Use the default option for any questions that may come up during appliance configuration. This property was added in vSphere API 6.7.
	AutoAnswer *bool
}



func upgradeGetInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeGetOutputType() bindings.BindingType {
	return bindings.NewReferenceType(UpgradeUpgradeSpecBindingType)
}

func upgradeGetRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}

func upgradeCheckInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(UpgradeUpgradeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeCheckOutputType() bindings.BindingType {
	return bindings.NewReferenceType(CheckInfoBindingType)
}

func upgradeCheckRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(UpgradeUpgradeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}

func upgradeStartInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["spec"] = bindings.NewReferenceType(UpgradeUpgradeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeStartOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func upgradeStartRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	fields["spec"] = bindings.NewReferenceType(UpgradeUpgradeSpecBindingType)
	fieldNameMap["spec"] = "Spec"
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"InvalidArgument": 400,"NotAllowedInCurrentState": 400})
}

func upgradeCancelInputType() bindings.StructType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	var validators = []bindings.Validator{}
	return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func upgradeCancelOutputType() bindings.BindingType {
	return bindings.NewVoidType()
}

func upgradeCancelRestMetadata() protocol.OperationRestMetadata {
	fields := map[string]bindings.BindingType{}
	fieldNameMap := map[string]string{}
	paramsTypeMap := map[string]bindings.BindingType{}
	pathParams := map[string]string{}
	queryParams := map[string]string{}
	headerParams := map[string]string{}
	resultHeaders := map[string]string{}
	errorHeaders := map[string]string{}
	errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
	return protocol.NewOperationRestMetadata(
		fields,
		fieldNameMap,
		paramsTypeMap,
		pathParams,
		queryParams,
		headerParams,
		"",
		"",
		"null",
		"",
		resultHeaders,
		0,
		errorHeaders,
		map[string]int{"Unauthenticated": 401,"NotAllowedInCurrentState": 400})
}


func UpgradeVcsaEmbeddedSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ceip_enabled"] = bindings.NewBooleanType()
	fieldNameMap["ceip_enabled"] = "CeipEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.vcsa_embedded_spec", fields, reflect.TypeOf(UpgradeVcsaEmbeddedSpec{}), fieldNameMap, validators)
}

func UpgradePscSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["ceip_enabled"] = bindings.NewBooleanType()
	fieldNameMap["ceip_enabled"] = "CeipEnabled"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.psc_spec", fields, reflect.TypeOf(UpgradePscSpec{}), fieldNameMap, validators)
}

func UpgradeSourceApplianceSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["hostname"] = bindings.NewStringType()
	fieldNameMap["hostname"] = "Hostname"
	fields["https_port"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["https_port"] = "HttpsPort"
	fields["ssl_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssl_thumbprint"] = "SslThumbprint"
	fields["ssl_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssl_verify"] = "SslVerify"
	fields["sso_admin_username"] = bindings.NewStringType()
	fieldNameMap["sso_admin_username"] = "SsoAdminUsername"
	fields["sso_admin_password"] = bindings.NewSecretType()
	fieldNameMap["sso_admin_password"] = "SsoAdminPassword"
	fields["root_password"] = bindings.NewSecretType()
	fieldNameMap["root_password"] = "RootPassword"
	fields["ssh_verify"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ssh_verify"] = "SshVerify"
	fields["ssh_thumbprint"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["ssh_thumbprint"] = "SshThumbprint"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.source_appliance_spec", fields, reflect.TypeOf(UpgradeSourceApplianceSpec{}), fieldNameMap, validators)
}

func UpgradeUpgradeSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["source_appliance"] = bindings.NewReferenceType(UpgradeSourceApplianceSpecBindingType)
	fieldNameMap["source_appliance"] = "SourceAppliance"
	fields["source_location"] = bindings.NewReferenceType(LocationSpecBindingType)
	fieldNameMap["source_location"] = "SourceLocation"
	fields["history"] = bindings.NewOptionalType(bindings.NewReferenceType(HistoryMigrationSpecBindingType))
	fieldNameMap["history"] = "History"
	fields["vcsa_embedded"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradeVcsaEmbeddedSpecBindingType))
	fieldNameMap["vcsa_embedded"] = "VcsaEmbedded"
	fields["psc"] = bindings.NewOptionalType(bindings.NewReferenceType(UpgradePscSpecBindingType))
	fieldNameMap["psc"] = "Psc"
	fields["auto_answer"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["auto_answer"] = "AutoAnswer"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.deployment.upgrade.upgrade_spec", fields, reflect.TypeOf(UpgradeUpgradeSpec{}), fieldNameMap, validators)
}

