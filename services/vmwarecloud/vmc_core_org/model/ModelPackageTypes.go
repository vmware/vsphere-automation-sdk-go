// Copyright © 2019-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Data type definitions file for package: com.vmware.vmwarecloud.vmc_core_org.model.
// Includes binding types of a top level structures and enumerations.
// Shared by client-side stubs and server-side skeletons to ensure type
// compatibility.

package model

import (
	vapiBindings_ "github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	vapiData_ "github.com/vmware/vsphere-automation-sdk-go/runtime/data"
	vapiLog_ "github.com/vmware/vsphere-automation-sdk-go/runtime/log"
	"reflect"
	"time"
)

type Organization struct {
	Id          string
	Name        string
	DisplayName *string
	OrgType     *string
	States      []OrganizationState
	OrgTypes    []OrganizationType
	OrgConfigs  []OrganizationConfig
	Version     *int64
	Creator     *Creator
	Updater     *Updater
	Tags        []string
	OrgTags     []OrganizationTag
	// List of effective properties for an org, by provider.
	EffectiveProperties []EffectiveProperties
	OnboardedOnDate     *string
	OnboardedByUser     *string
	// The most recent update time to the org across all configs.
	LastUpdatedTimestamp *time.Time
}

func (s *Organization) GetType__() vapiBindings_.BindingType {
	return OrganizationBindingType()
}

func (s *Organization) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Organization._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrganizationType struct {
	Id           *string
	OrgType      string
	ProviderName *string
	Version      *int64
	Creator      *Creator
	Updater      *Updater
}

func (s *OrganizationType) GetType__() vapiBindings_.BindingType {
	return OrganizationTypeBindingType()
}

func (s *OrganizationType) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrganizationType._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrganizationConfig struct {
	Id         *string
	Name       *string
	Category   *string
	PipelineId *string
	// Name of the provider
	Provider *string
	OrgId    *string
	OrgType  *string
	DataType *string
	Version  *int64
	Creator  *Creator
	Updater  *Updater
	Data     map[string]vapiData_.DataValue
	// The most recent update time to the org's effective properties
	UpdatedTimestamp *time.Time
}

const OrganizationConfig_CATEGORY_PROVIDER = "provider"
const OrganizationConfig_CATEGORY_ORGTYPE = "orgtype"
const OrganizationConfig_CATEGORY_PIPELINE = "pipeline"
const OrganizationConfig_CATEGORY_OVERRIDE = "override"
const OrganizationConfig_CATEGORY_FEATURE_FLAG = "feature_flag"

func (s *OrganizationConfig) GetType__() vapiBindings_.BindingType {
	return OrganizationConfigBindingType()
}

func (s *OrganizationConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrganizationConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrganizationState struct {
	Id           *string
	State        string
	ProviderName *string
	Version      *int64
	Creator      *Creator
	Updater      *Updater
}

const OrganizationState_STATE_CREATED = "CREATED"
const OrganizationState_STATE_DELETED = "DELETED"
const OrganizationState_STATE_DISABLED = "DISABLED"

func (s *OrganizationState) GetType__() vapiBindings_.BindingType {
	return OrganizationStateBindingType()
}

func (s *OrganizationState) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrganizationState._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrganizationTag struct {
	// Name of the provider that this group of tags scope to
	ProviderName *string
	// key of the tag
	Key *string
	// value of the tag
	Value *string
}

func (s *OrganizationTag) GetType__() vapiBindings_.BindingType {
	return OrganizationTagBindingType()
}

func (s *OrganizationTag) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrganizationTag._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type DataPayloadConfig struct {
	Data map[string]vapiData_.DataValue
	// The most recent update time to the org's effective properties
	UpdatedTimestamp *time.Time
}

func (s *DataPayloadConfig) GetType__() vapiBindings_.BindingType {
	return DataPayloadConfigBindingType()
}

func (s *DataPayloadConfig) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for DataPayloadConfig._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrgProfileValidationResponse struct {
	OrgValidationErrors []OrgProfileValidationStatus
	ProfileStatus       *string
}

const OrgProfileValidationResponse_PROFILE_STATUS_VALID = "VALID"
const OrgProfileValidationResponse_PROFILE_STATUS_INVALID = "INVALID"

func (s *OrgProfileValidationResponse) GetType__() vapiBindings_.BindingType {
	return OrgProfileValidationResponseBindingType()
}

func (s *OrgProfileValidationResponse) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrgProfileValidationResponse._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Creator struct {
	// Name of user that created the deployment
	UserName *string
	// Identifier of the user that created the deployment
	UserId *string
	// Identifier of the client that created the deployment
	ClientId *string
	// Timestamp of deployment creation
	Timestamp *int64
}

func (s *Creator) GetType__() vapiBindings_.BindingType {
	return CreatorBindingType()
}

func (s *Creator) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Creator._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type Updater struct {
	// Name of user that updated the deployment
	UserName *string
	// Identifier of the user that updated the deployment
	UserId *string
	// Identifier of the client that updated the deployment
	ClientId *string
	// Timestamp of deployment update
	Timestamp *int64
}

func (s *Updater) GetType__() vapiBindings_.BindingType {
	return UpdaterBindingType()
}

func (s *Updater) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for Updater._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type EffectiveProperties struct {
	ProviderName       *string
	ProviderProperties *DataPayloadConfig
}

func (s *EffectiveProperties) GetType__() vapiBindings_.BindingType {
	return EffectivePropertiesBindingType()
}

func (s *EffectiveProperties) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for EffectiveProperties._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

type OrgProfileValidationStatus struct {
	Description        *string
	OrgValidationError *string
}

const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_INVALID_ORG_ADDRESS = "INVALID_ORG_ADDRESS"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_MISSING_ORG_ADDRESS = "MISSING_ORG_ADDRESS"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_MISSING_ORG_PAYMENT_METHOD = "MISSING_ORG_PAYMENT_METHOD"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_TOS_IS_NOT_SIGNED = "TOS_IS_NOT_SIGNED"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_MISSING_VMWARE_ID_ACCOUNT = "MISSING_VMWARE_ID_ACCOUNT"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_FAILED_TRADE_COMPLIANCE_CHECK = "FAILED_TRADE_COMPLIANCE_CHECK"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_PENDING_TRADE_COMPLIANCE_CHECK = "PENDING_TRADE_COMPLIANCE_CHECK"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_MISSING_BILLING_ACCOUNT = "MISSING_BILLING_ACCOUNT"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_MISSING_BILLING_ACCOUNT_PAYMENT_METHOD = "MISSING_BILLING_ACCOUNT_PAYMENT_METHOD"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_INVALID_PAYMENT_METHOD = "INVALID_PAYMENT_METHOD"
const OrgProfileValidationStatus_ORG_VALIDATION_ERROR_UNKNOWN = "UNKNOWN"

func (s *OrgProfileValidationStatus) GetType__() vapiBindings_.BindingType {
	return OrgProfileValidationStatusBindingType()
}

func (s *OrgProfileValidationStatus) GetDataValue__() (vapiData_.DataValue, []error) {
	typeConverter := vapiBindings_.NewTypeConverter()
	dataVal, err := typeConverter.ConvertToVapi(s, s.GetType__())
	if err != nil {
		vapiLog_.Errorf("Error in ConvertToVapi for OrgProfileValidationStatus._GetDataValue method - %s",
			vapiBindings_.VAPIerrorsToError(err).Error())
		return nil, err
	}
	return dataVal, nil
}

func OrganizationBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewStringType()
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewStringType()
	fieldNameMap["name"] = "Name"
	fields["display_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["display_name"] = "DisplayName"
	fields["org_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["states"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OrganizationStateBindingType), reflect.TypeOf([]OrganizationState{})))
	fieldNameMap["states"] = "States"
	fields["org_types"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OrganizationTypeBindingType), reflect.TypeOf([]OrganizationType{})))
	fieldNameMap["org_types"] = "OrgTypes"
	fields["org_configs"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OrganizationConfigBindingType), reflect.TypeOf([]OrganizationConfig{})))
	fieldNameMap["org_configs"] = "OrgConfigs"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["tags"] = "Tags"
	fields["org_tags"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OrganizationTagBindingType), reflect.TypeOf([]OrganizationTag{})))
	fieldNameMap["org_tags"] = "OrgTags"
	fields["effective_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(EffectivePropertiesBindingType), reflect.TypeOf([]EffectiveProperties{})))
	fieldNameMap["effective_properties"] = "EffectiveProperties"
	fields["onboarded_on_date"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["onboarded_on_date"] = "OnboardedOnDate"
	fields["onboarded_by_user"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["onboarded_by_user"] = "OnboardedByUser"
	fields["last_updated_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["last_updated_timestamp"] = "LastUpdatedTimestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.organization", fields, reflect.TypeOf(Organization{}), fieldNameMap, validators)
}

func OrganizationTypeBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["org_type"] = vapiBindings_.NewStringType()
	fieldNameMap["org_type"] = "OrgType"
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.organization_type", fields, reflect.TypeOf(OrganizationType{}), fieldNameMap, validators)
}

func OrganizationConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["category"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["category"] = "Category"
	fields["pipeline_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["pipeline_id"] = "PipelineId"
	fields["provider"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider"] = "Provider"
	fields["org_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_id"] = "OrgId"
	fields["org_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_type"] = "OrgType"
	fields["data_type"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["data_type"] = "DataType"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	fields["data"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewOpaqueType(), reflect.TypeOf(map[string]vapiData_.DataValue{})))
	fieldNameMap["data"] = "Data"
	fields["updated_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["updated_timestamp"] = "UpdatedTimestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.organization_config", fields, reflect.TypeOf(OrganizationConfig{}), fieldNameMap, validators)
}

func OrganizationStateBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["state"] = vapiBindings_.NewStringType()
	fieldNameMap["state"] = "State"
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["version"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["version"] = "Version"
	fields["creator"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(CreatorBindingType))
	fieldNameMap["creator"] = "Creator"
	fields["updater"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(UpdaterBindingType))
	fieldNameMap["updater"] = "Updater"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.organization_state", fields, reflect.TypeOf(OrganizationState{}), fieldNameMap, validators)
}

func OrganizationTagBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["key"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.organization_tag", fields, reflect.TypeOf(OrganizationTag{}), fieldNameMap, validators)
}

func DataPayloadConfigBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["data"] = vapiBindings_.NewOptionalType(vapiBindings_.NewMapType(vapiBindings_.NewStringType(), vapiBindings_.NewOpaqueType(), reflect.TypeOf(map[string]vapiData_.DataValue{})))
	fieldNameMap["data"] = "Data"
	fields["updated_timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewDateTimeType())
	fieldNameMap["updated_timestamp"] = "UpdatedTimestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.data_payload_config", fields, reflect.TypeOf(DataPayloadConfig{}), fieldNameMap, validators)
}

func OrgProfileValidationResponseBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["org_validation_errors"] = vapiBindings_.NewOptionalType(vapiBindings_.NewListType(vapiBindings_.NewReferenceType(OrgProfileValidationStatusBindingType), reflect.TypeOf([]OrgProfileValidationStatus{})))
	fieldNameMap["org_validation_errors"] = "OrgValidationErrors"
	fields["profile_status"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["profile_status"] = "ProfileStatus"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.org_profile_validation_response", fields, reflect.TypeOf(OrgProfileValidationResponse{}), fieldNameMap, validators)
}

func CreatorBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["client_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["client_id"] = "ClientId"
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.creator", fields, reflect.TypeOf(Creator{}), fieldNameMap, validators)
}

func UpdaterBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["user_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_name"] = "UserName"
	fields["user_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["user_id"] = "UserId"
	fields["client_id"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["client_id"] = "ClientId"
	fields["timestamp"] = vapiBindings_.NewOptionalType(vapiBindings_.NewIntegerType())
	fieldNameMap["timestamp"] = "Timestamp"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.updater", fields, reflect.TypeOf(Updater{}), fieldNameMap, validators)
}

func EffectivePropertiesBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["provider_name"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["provider_name"] = "ProviderName"
	fields["provider_properties"] = vapiBindings_.NewOptionalType(vapiBindings_.NewReferenceType(DataPayloadConfigBindingType))
	fieldNameMap["provider_properties"] = "ProviderProperties"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.effective_properties", fields, reflect.TypeOf(EffectiveProperties{}), fieldNameMap, validators)
}

func OrgProfileValidationStatusBindingType() vapiBindings_.BindingType {
	fields := make(map[string]vapiBindings_.BindingType)
	fieldNameMap := make(map[string]string)
	fields["description"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["org_validation_error"] = vapiBindings_.NewOptionalType(vapiBindings_.NewStringType())
	fieldNameMap["org_validation_error"] = "OrgValidationError"
	var validators = []vapiBindings_.Validator{}
	return vapiBindings_.NewStructType("com.vmware.vmwarecloud.vmc_core_org.model.org_profile_validation_status", fields, reflect.TypeOf(OrgProfileValidationStatus{}), fieldNameMap, validators)
}
