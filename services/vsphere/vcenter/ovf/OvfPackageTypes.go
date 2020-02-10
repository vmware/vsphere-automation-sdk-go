/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.ovf.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package ovf

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``DiskProvisioningType`` enumeration class defines the virtual disk provisioning types that can be set for a disk on the target platform.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type DiskProvisioningType string

const (
    // A thin provisioned virtual disk has space allocated and zeroed on demand as the space is used.
	DiskProvisioningType_thin DiskProvisioningType = "thin"
    // A thick provisioned virtual disk has all space allocated at creation time and the space is zeroed on demand as the space is used.
	DiskProvisioningType_thick DiskProvisioningType = "thick"
    // An eager zeroed thick provisioned virtual disk has all space allocated and wiped clean of any previous contents on the physical media at creation time. 
    //
    //  Disks specified as eager zeroed thick may take longer time to create than disks specified with the other disk provisioning types.
	DiskProvisioningType_eagerZeroedThick DiskProvisioningType = "eagerZeroedThick"
)

func (d DiskProvisioningType) DiskProvisioningType() bool {
	switch d {
	case DiskProvisioningType_thin:
		return true
	case DiskProvisioningType_thick:
		return true
	case DiskProvisioningType_eagerZeroedThick:
		return true
	default:
		return false
	}
}


// The ``CertificateParams`` class contains information about the public key certificate used to sign the OVF package. This class will only be returned if the OVF package is signed. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type CertificateParams struct {
    // Certificate issuer. For example: /C=US/ST=California/L=Palo Alto/O=VMware, Inc.
	Issuer *string
    // Certificate subject. For example: /C=US/ST=Massachusetts/L=Hopkinton/O=EMC Corporation/OU=EMC Avamar/CN=EMC Corporation.
	Subject *string
    // Is the certificate chain validated.
	IsValid *bool
    // Is the certificate self-signed.
	IsSelfSigned *bool
    // The X509 representation of the certificate.
	X509 *string
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``DeploymentOption`` class contains the information about a deployment option as defined in the OVF specification. 
//
//  This corresponds to the ovf:Configuration element of the ovf:DeploymentOptionSection in the specification. The ovf:DeploymentOptionSection specifies a discrete set of intended resource allocation configurations. This class represents one item from that set. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type DeploymentOption struct {
    // The key of the deployment option, corresponding to the ovf:id attribute in the OVF descriptor.
	Key *string
    // A localizable label for the deployment option.
	Label *string
    // A localizable description for the deployment option.
	Description *string
    // A bool flag indicates whether this deployment option is the default choice.
	DefaultChoice *bool
}

// The ``DeploymentOptionParams`` class describes the possible deployment options as well as the choice provided by the user. 
//
//  This information based on the ovf:DeploymentOptionSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type DeploymentOptionParams struct {
    // Array of deployment options. This property corresponds to the ovf:Configuration elements of the ovf:DeploymentOptionSection in the specification. It is a discrete set of intended resource allocation configurations from which one can be selected.
	DeploymentOptions []DeploymentOption
    // The selected deployment option. Identifies the DeploymentOption in the list in the ``deploymentOptions`` property with a matching value in the DeploymentOption#key property.
	SelectedKey *string
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``ExtraConfig`` class contains the information about a vmw:ExtraConfig element which can be used to specify configuration settings that are transferred directly to the ``.vmx`` file. The behavior of the vmw:ExtraConfig element is similar to the ``extraConfig`` property of the ``VirtualMachineConfigSpec`` object in the VMware vSphere API. Thus, the same restrictions apply, such as you cannot set values that could otherwise be set with other properties in the ``VirtualMachineConfigSpec`` object. See the VMware vSphere API reference for details on this. 
//
//  vmw:ExtraConfig elements may occur as direct child elements of a VirtualHardwareSection, or as child elements of individual virtual hardware items. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type ExtraConfig struct {
    // The key of the ExtraConfig element.
	Key *string
    // The value of the ExtraConfig element.
	Value *string
    // The identifier of the virtual system containing the vmw:ExtraConfig element.
	VirtualSystemId *string
}

// The ``ExtraConfigParams`` class contains the parameters with information about the vmw:ExtraConfig elements in an OVF package. 
//
//  vmw:ExtraConfig elements can be used to specify configuration settings that are transferred directly to the ``.vmx`` file. 
//
//  The behavior of the vmw:ExtraConfig element is similar to the ``extraConfig`` property of the ``VirtualMachineConfigSpec`` object in the VMware vSphere API. Thus, the same restrictions apply, such as you cannot set values that could otherwise be set with other properties in the ``VirtualMachineConfigSpec`` object. See the VMware vSphere API reference for details on this. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type ExtraConfigParams struct {
    // Array of vmw:ExtraConfig elements in the OVF package.
	ExtraConfigs []ExtraConfig
    // Specifies which extra configuration items in the array in the ``extraConfigs`` ``field`` should be ignored during deployment. 
    //
    //  If set, the given keys for extra configurations will be ignored during deployment. The key is defined in the ExtraConfig#key property.
	ExcludeKeys []string
    // Specifies which extra configuration items in the array in the ``extraConfigs`` ``field`` should be included during deployment. 
    //
    //  If set, all but the given keys for extra configurations will be ignored during deployment. The key is defined in the ExtraConfig#key property.
	IncludeKeys []string
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``IpAllocationParams`` class specifies how IP addresses are allocated to OVF properties. In particular, it informs the deployment platform whether the guest supports IPv4, IPv6, or both. It also specifies whether the IP addresses can be obtained through DHCP or through the properties provided in the OVF environment. 
//
//  Ovf Property elements are exposed to the guest software through the OVF environment. Each Property element exposed in the OVF environment shall be constructed from the value of the ovf:key attribute. A Property element contains a key/value pair, it may optionally specify type qualifiers using the ovf:qualifiers attribute with multiple qualifiers separated by commas. 
//
//  The settings in ``IpAllocationParams`` class are global to a deployment. Thus, if a virtual machine is part of a virtual appliance, then its settings are ignored and the settings for the virtual appliance is used. 
//
//  This information is based on the vmw:IpAssignmentSection in OVF package. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type IpAllocationParams struct {
    // Specifies the IP allocation schemes supported by the guest software. This property defines the valid values for the IP allocation policy. This setting is often configured by the virtual appliance template author or OVF package author to reflect what the guest software supports, and the IP allocation policy is configured at deployment time. See IpAllocationParams#ipAllocationPolicy.
	SupportedAllocationScheme []IpAllocationParamsIpAllocationScheme
    // Specifies the IP allocation policies supported. The set of valid options for the policy is based on the capabilities of the virtual appliance software, as specified by the IpAllocationParams#supportedAllocationScheme property.
	SupportedIpAllocationPolicy []IpAllocationParamsIpAllocationPolicy
    // Specifies how IP allocation is done through an IP Pool. This is typically specified by the deployer.
	IpAllocationPolicy *IpAllocationParamsIpAllocationPolicy
    // Specifies the IP protocols supported by the guest.
	SupportedIpProtocol []IpAllocationParamsIpProtocol
    // Specifies the chosen IP protocol for this deployment. This must be one of the IP protocols supported by the guest software. See IpAllocationParams#supportedIpProtocol.
	IpProtocol *IpAllocationParamsIpProtocol
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``IpAllocationPolicy`` enumeration class defines the possible IP allocation policy for a deployment.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type IpAllocationParamsIpAllocationPolicy string

const (
    // Specifies that DHCP will be used to allocate IP addresses.
	IpAllocationParamsIpAllocationPolicy_DHCP IpAllocationParamsIpAllocationPolicy = "DHCP"
    // Specifies that IP addresses are allocated from an IP pool. The IP addresses are allocated when needed, typically at power-on, and deallocated during power-off. There is no guarantee that a property will receive same IP address when restarted.
	IpAllocationParamsIpAllocationPolicy_TRANSIENT_IPPOOL IpAllocationParamsIpAllocationPolicy = "TRANSIENT_IPPOOL"
    // Specifies that IP addresses are configured manually upon deployment, and will be kept until reconfigured or the virtual appliance destroyed. This ensures that a property gets a consistent IP for its lifetime.
	IpAllocationParamsIpAllocationPolicy_STATIC_MANUAL IpAllocationParamsIpAllocationPolicy = "STATIC_MANUAL"
    // Specifies that IP addresses are allocated from the range managed by an IP pool. The IP addresses are allocated at first power-on, and remain allocated at power-off. This ensures that a virtual appliance gets a consistent IP for its life-time.
	IpAllocationParamsIpAllocationPolicy_STATIC_IPPOOL IpAllocationParamsIpAllocationPolicy = "STATIC_IPPOOL"
)

func (i IpAllocationParamsIpAllocationPolicy) IpAllocationParamsIpAllocationPolicy() bool {
	switch i {
	case IpAllocationParamsIpAllocationPolicy_DHCP:
		return true
	case IpAllocationParamsIpAllocationPolicy_TRANSIENT_IPPOOL:
		return true
	case IpAllocationParamsIpAllocationPolicy_STATIC_MANUAL:
		return true
	case IpAllocationParamsIpAllocationPolicy_STATIC_IPPOOL:
		return true
	default:
		return false
	}
}


// The ``IpAllocationScheme`` enumeration class defines the possible IP allocation schemes that can be supported by the guest software.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type IpAllocationParamsIpAllocationScheme string

const (
    // It supports DHCP to acquire IP configuration.
	IpAllocationParamsIpAllocationScheme_DHCP IpAllocationParamsIpAllocationScheme = "DHCP"
    // It supports setting the IP configuration through the properties provided in the OVF environment.
	IpAllocationParamsIpAllocationScheme_OVF_ENVIRONMENT IpAllocationParamsIpAllocationScheme = "OVF_ENVIRONMENT"
)

func (i IpAllocationParamsIpAllocationScheme) IpAllocationParamsIpAllocationScheme() bool {
	switch i {
	case IpAllocationParamsIpAllocationScheme_DHCP:
		return true
	case IpAllocationParamsIpAllocationScheme_OVF_ENVIRONMENT:
		return true
	default:
		return false
	}
}


// The ``IpProtocol`` enumeration class defines the IP protocols supported by the guest software.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type IpAllocationParamsIpProtocol string

const (
    // It supports the IPv4 protocol.
	IpAllocationParamsIpProtocol_IPV4 IpAllocationParamsIpProtocol = "IPV4"
    // It supports the IPv6 protocol.
	IpAllocationParamsIpProtocol_IPV6 IpAllocationParamsIpProtocol = "IPV6"
)

func (i IpAllocationParamsIpProtocol) IpAllocationParamsIpProtocol() bool {
	switch i {
	case IpAllocationParamsIpProtocol_IPV4:
		return true
	case IpAllocationParamsIpProtocol_IPV6:
		return true
	default:
		return false
	}
}


// The ``OvfMessage`` class describes a base OVF handling error message related to accessing, validating, deploying, or exporting an OVF package. 
//
//  These messages fall into different categories defined in OvfMessageCategory:
type OvfMessage struct {
    // The message category.
	Category OvfMessageCategory
    // Array of parse issues (see ParseIssue).
	Issues []ParseIssue
    // The name of input parameter.
	Name *string
    // The value of input parameter.
	Value *string
    // A localizable message.
	Message *std.LocalizableMessage
    // Represents a server errors.Error.
	Error_ *data.StructValue
}

// The ``Category`` enumeration class defines the categories of messages (see OvfMessage).
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type OvfMessageCategory string

const (
    // The OVF descriptor is invalid, for example, syntax errors or schema errors.
	OvfMessageCategory_VALIDATION OvfMessageCategory = "VALIDATION"
    // The user provided input parameters are invalid.
	OvfMessageCategory_INPUT OvfMessageCategory = "INPUT"
    // Server error.
	OvfMessageCategory_SERVER OvfMessageCategory = "SERVER"
)

func (c OvfMessageCategory) OvfMessageCategory() bool {
	switch c {
	case OvfMessageCategory_VALIDATION:
		return true
	case OvfMessageCategory_INPUT:
		return true
	case OvfMessageCategory_SERVER:
		return true
	default:
		return false
	}
}


// The ``ParseIssue`` class contains the information about the issue found when parsing an OVF package during deployment or exporting an OVF package including: 
//
// * Parsing and validation error on OVF descriptor (which is an XML document), manifest and certificate files.
// * OVF descriptor generating and device error.
// * Unexpected server error.
type ParseIssue struct {
    // The category of the parse issue.
	Category ParseIssueCategory
    // The name of the file in which the parse issue was found.
	File string
    // The line number of the line in the file (see ParseIssue#file) where the parse issue was found (or -1 if not applicable).
	LineNumber int64
    // The position in the line (see ParseIssue#lineNumber) (or -1 if not applicable).
	ColumnNumber int64
    // A localizable message describing the parse issue.
	Message std.LocalizableMessage
}

// The ``Category`` enumeration class defines the categories of issues that can be found when parsing files inside an OVF package (see ParseIssue) including OVF descriptor (which is an XML document), manifest and certificate files, or exporting an OVF package.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
type ParseIssueCategory string

const (
    // Illegal value error. For example, the value is malformed, not a number, or outside of the given range, and so on.
	ParseIssueCategory_VALUE_ILLEGAL ParseIssueCategory = "VALUE_ILLEGAL"
    // Required attribute error. It indicates that a required attribute is missing from an element in the OVF descriptor.
	ParseIssueCategory_ATTRIBUTE_REQUIRED ParseIssueCategory = "ATTRIBUTE_REQUIRED"
    // Illegal attribute error. It indicates that an illegal attribute is set for an element in the OVF descriptor. For example, empty disks do not use format, parentRef, and populatedSize attributes, if these attributes are present in an empty disk element then will get this pasrse issue.
	ParseIssueCategory_ATTRIBUTE_ILLEGAL ParseIssueCategory = "ATTRIBUTE_ILLEGAL"
    // Required element error. It indicates that a required element is missing from the OVF descriptor.
	ParseIssueCategory_ELEMENT_REQUIRED ParseIssueCategory = "ELEMENT_REQUIRED"
    // Illegal element error. It indicates that an element is present in a location which is not allowed, or found multiple elements but only one is allowed at the location in the OVF descriptor.
	ParseIssueCategory_ELEMENT_ILLEGAL ParseIssueCategory = "ELEMENT_ILLEGAL"
    // Unknown element error. It indicates that an element is unsupported when parsing an OVF descriptor.
	ParseIssueCategory_ELEMENT_UNKNOWN ParseIssueCategory = "ELEMENT_UNKNOWN"
    // Section unknown error. It indicates that a section is unsupported when parsing an OVF descriptor.
	ParseIssueCategory_SECTION_UNKNOWN ParseIssueCategory = "SECTION_UNKNOWN"
    // Section restriction error. It indicates that a section appears in place in the OVF descriptor where it is not allowed, a section appears fewer times than is required, or a section appears more times than is allowed.
	ParseIssueCategory_SECTION_RESTRICTION ParseIssueCategory = "SECTION_RESTRICTION"
    // OVF package parsing error, including: 
    //
    // * OVF descriptor parsing errors, for example, syntax errors or schema errors.
    // * Manifest file parsing and verification errors.
    // * Certificate file parsing and verification errors.
	ParseIssueCategory_PARSE_ERROR ParseIssueCategory = "PARSE_ERROR"
    // OVF descriptor (which is an XML document) generating error, for example, well-formedness errors as well as unexpected processing conditions.
	ParseIssueCategory_GENERATE_ERROR ParseIssueCategory = "GENERATE_ERROR"
    // An issue with the manifest and signing.
	ParseIssueCategory_VALIDATION_ERROR ParseIssueCategory = "VALIDATION_ERROR"
    // Issue during OVF export, for example, malformed deviceId, controller not found, or file backing for a device not found.
	ParseIssueCategory_EXPORT_ERROR ParseIssueCategory = "EXPORT_ERROR"
    // Server encountered an unexpected error which prevented it from fulfilling the request.
	ParseIssueCategory_INTERNAL_ERROR ParseIssueCategory = "INTERNAL_ERROR"
)

func (c ParseIssueCategory) ParseIssueCategory() bool {
	switch c {
	case ParseIssueCategory_VALUE_ILLEGAL:
		return true
	case ParseIssueCategory_ATTRIBUTE_REQUIRED:
		return true
	case ParseIssueCategory_ATTRIBUTE_ILLEGAL:
		return true
	case ParseIssueCategory_ELEMENT_REQUIRED:
		return true
	case ParseIssueCategory_ELEMENT_ILLEGAL:
		return true
	case ParseIssueCategory_ELEMENT_UNKNOWN:
		return true
	case ParseIssueCategory_SECTION_UNKNOWN:
		return true
	case ParseIssueCategory_SECTION_RESTRICTION:
		return true
	case ParseIssueCategory_PARSE_ERROR:
		return true
	case ParseIssueCategory_GENERATE_ERROR:
		return true
	case ParseIssueCategory_VALIDATION_ERROR:
		return true
	case ParseIssueCategory_EXPORT_ERROR:
		return true
	case ParseIssueCategory_INTERNAL_ERROR:
		return true
	default:
		return false
	}
}


// The ``OvfError`` class describes an error related to accessing, validating, deploying, or exporting an OVF package.
type OvfError struct {
    // The message category.
	Category OvfMessageCategory
    // Array of parse issues (see ParseIssue).
	Issues []ParseIssue
    // The name of input parameter.
	Name *string
    // The value of input parameter.
	Value *string
    // A localizable message.
	Message *std.LocalizableMessage
    // Represents a server errors.Error.
	Error_ *data.StructValue
}

// The ``OvfWarning`` class describes a warning related to accessing, validating, deploying, or exporting an OVF package.
type OvfWarning struct {
    // The message category.
	Category OvfMessageCategory
    // Array of parse issues (see ParseIssue).
	Issues []ParseIssue
    // The name of input parameter.
	Name *string
    // The value of input parameter.
	Value *string
    // A localizable message.
	Message *std.LocalizableMessage
    // Represents a server errors.Error.
	Error_ *data.StructValue
}

// The ``OvfInfo`` class contains informational messages related to accessing, validating, deploying, or exporting an OVF package.
type OvfInfo struct {
    // A array of localizable messages (see std.LocalizableMessage).
	Messages []std.LocalizableMessage
}

// The ``OvfParams`` class defines the common properties for all OVF deployment parameters. OVF parameters serve several purposes: 
//
// * Describe information about a given OVF package.
// * Describe default deployment configuration.
// * Describe possible deployment values based on the deployment environment.
// * Provide deployment-specific configuration.
//
//  Each OVF parameters class specifies a particular configurable aspect of OVF deployment. An aspect has both a query-model and a deploy-model. The query-model is used when the OVF package is queried, and the deploy-model is used when deploying an OVF package. 
//
//  Most OVF parameter classes provide both informational and deployment parameters. However, some are purely informational (for example, download size) and some are purely deployment parameters (for example, the flag to indicate whether registration as a vCenter extension is accepted). 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type OvfParams struct {
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``Property`` class contains the information about a property in an OVF package. 
//
//  A property is uniquely identified by its [classid.]id[.instanceid] fully-qualified name (see Property#classId, Property#id, and Property#instanceId). If multiple properties in an OVF package have the same fully-qualified name, then the property is excluded and cannot be set. We do warn about this during import. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type Property struct {
    // The classId of this OVF property.
	ClassId *string
    // The identifier of this OVF property.
	Id *string
    // The instanceId of this OVF property.
	InstanceId *string
    // If this is set to a non-empty string, this property starts a new category.
	Category *string
    // Whether a category is UI optional. This is only used if this property starts a new category (see Property#category). 
    //
    //  The value is stored in an optional attribute vmw:uioptional to the ovf:Category element. The default value is false. If this value is true, the properties within this category are optional. The UI renders this as a group with a check box, and the group is grayed out until the check box is selected. When the check box is selected, the input values are read and used in deployment. If properties within the same category specify conflicting values the default is used. Only implemented in vSphere Web Client 5.1 and later as of Nov 2012.
	UiOptional *bool
    // The display name of this OVF property.
	Label *string
    // A description of this OVF property.
	Description *string
    // The type of this OVF property. Refer to the configuration of a virtual appliance/virtual machine for the valid values.
	Type_ *string
    // The OVF property value. This contains the default value from ovf:defaultValue if ovf:value is not present when read.
	Value *string
}

// The ``PropertyParams`` class contains a array of OVF properties that can be configured when the OVF package is deployed. 
//
//  This is based on the ovf:ProductSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type PropertyParams struct {
    // Array of OVF properties.
	Properties []Property
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``ScaleOutGroup`` class contains information about a scale-out group. 
//
//  It allows a virtual system collection to contain a set of children that are homogeneous with respect to a prototypical virtual system or virtual system collection. It shall cause the deployment function to replicate the prototype a number of times, thus allowing the number of instantiated virtual systems to be configured dynamically at deployment time. 
//
//  This is based on the ovf2:ScaleOutSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type ScaleOutGroup struct {
    // The identifier of the scale-out group.
	Id *string
    // The description of the scale-out group.
	Description *string
    // The scaling factor to use. It defines the number of replicas of the prototypical virtual system or virtual system collection.
	InstanceCount *int64
    // The minimum scaling factor.
	MinimumInstanceCount *int64
    // The maximum scaling factor.
	MaximumInstanceCount *int64
}

// The ``ScaleOutParams`` class contains information about the scale-out groups described in the OVF package. 
//
//  When deploying an OVF package, a deployment specific instance count can be specified (see ScaleOutGroup#instanceCount. 
//
//  This is based on the ovf2:ScaleOutSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type ScaleOutParams struct {
    // The array of scale-out groups.
	Groups []ScaleOutGroup
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``SizeParams`` class contains estimates of the download and deployment sizes. 
//
//  This information is based on the file references and the ovf:DiskSection in the OVF descriptor. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type SizeParams struct {
    // A best guess as to the total amount of data that must be transferred to download the OVF package. 
    //
    //  This may be inaccurate due to disk compression etc.
	ApproximateDownloadSize *int64
    // A best guess as to the total amount of space required to deploy the OVF package if using flat disks.
	ApproximateFlatDeploymentSize *int64
    // A best guess as to the total amount of space required to deploy the OVF package using sparse disks.
	ApproximateSparseDeploymentSize *int64
    // Whether the OVF uses variable disk sizes. 
    //
    //  For empty disks, rather than specifying a fixed virtual disk capacity, the capacity may be given using a reference to a ovf:Property element in a ovf:ProductSection element in OVF package.
	VariableDiskSize *bool
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``UnknownSection`` class contains information about an unknown section in an OVF package.
type UnknownSection struct {
    // A namespace-qualified tag in the form ``{ns}tag``.
	Tag string
    // The description of the Info element.
	Info string
}

// The ``UnknownSectionParams`` class contains a array of unknown, non-required sections. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type UnknownSectionParams struct {
    // Array of unknown, non-required sections.
	UnknownSections []UnknownSection
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}

// The ``VcenterExtensionParams`` class specifies that the OVF package should be registered as a vCenter extension. The virtual machine or virtual appliance will gain unrestricted access to the vCenter Server APIs. It must be connected to a network with connectivity to the vCenter server. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type VcenterExtensionParams struct {
    // Whether registration as a vCenter extension is required.
	Required *bool
    // Whether registration as a vCenter extension is accepted. 
    //
    //  If registration as a vCenter extension is required (see VcenterExtensionParams#required), this must be set to true during deployment. Defaults to false when returned from server.
	RegistrationAccepted *bool
    // Unique identifier describing the type of the OVF parameters. The value is the name of the OVF parameters class.
	Type_ *string
}




func CertificateParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["issuer"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["issuer"] = "Issuer"
	fields["subject"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["subject"] = "Subject"
	fields["is_valid"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["is_valid"] = "IsValid"
	fields["is_self_signed"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["is_self_signed"] = "IsSelfSigned"
	fields["x509"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["x509"] = "X509"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.certificate_params", fields, reflect.TypeOf(CertificateParams{}), fieldNameMap, validators)
}

func DeploymentOptionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["default_choice"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["default_choice"] = "DefaultChoice"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.deployment_option", fields, reflect.TypeOf(DeploymentOption{}), fieldNameMap, validators)
}

func DeploymentOptionParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["deployment_options"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DeploymentOptionBindingType), reflect.TypeOf([]DeploymentOption{})))
	fieldNameMap["deployment_options"] = "DeploymentOptions"
	fields["selected_key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["selected_key"] = "SelectedKey"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.deployment_option_params", fields, reflect.TypeOf(DeploymentOptionParams{}), fieldNameMap, validators)
}

func ExtraConfigBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["key"] = "Key"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["virtual_system_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["virtual_system_id"] = "VirtualSystemId"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.extra_config", fields, reflect.TypeOf(ExtraConfig{}), fieldNameMap, validators)
}

func ExtraConfigParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["extra_configs"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ExtraConfigBindingType), reflect.TypeOf([]ExtraConfig{})))
	fieldNameMap["extra_configs"] = "ExtraConfigs"
	fields["exclude_keys"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["exclude_keys"] = "ExcludeKeys"
	fields["include_keys"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
	fieldNameMap["include_keys"] = "IncludeKeys"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.extra_config_params", fields, reflect.TypeOf(ExtraConfigParams{}), fieldNameMap, validators)
}

func IpAllocationParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["supported_allocation_scheme"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_scheme", reflect.TypeOf(IpAllocationParamsIpAllocationScheme(IpAllocationParamsIpAllocationScheme_DHCP))), reflect.TypeOf([]IpAllocationParamsIpAllocationScheme{})))
	fieldNameMap["supported_allocation_scheme"] = "SupportedAllocationScheme"
	fields["supported_ip_allocation_policy"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_policy", reflect.TypeOf(IpAllocationParamsIpAllocationPolicy(IpAllocationParamsIpAllocationPolicy_DHCP))), reflect.TypeOf([]IpAllocationParamsIpAllocationPolicy{})))
	fieldNameMap["supported_ip_allocation_policy"] = "SupportedIpAllocationPolicy"
	fields["ip_allocation_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_policy", reflect.TypeOf(IpAllocationParamsIpAllocationPolicy(IpAllocationParamsIpAllocationPolicy_DHCP))))
	fieldNameMap["ip_allocation_policy"] = "IpAllocationPolicy"
	fields["supported_ip_protocol"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_protocol", reflect.TypeOf(IpAllocationParamsIpProtocol(IpAllocationParamsIpProtocol_IPV4))), reflect.TypeOf([]IpAllocationParamsIpProtocol{})))
	fieldNameMap["supported_ip_protocol"] = "SupportedIpProtocol"
	fields["ip_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_protocol", reflect.TypeOf(IpAllocationParamsIpProtocol(IpAllocationParamsIpProtocol_IPV4))))
	fieldNameMap["ip_protocol"] = "IpProtocol"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.ip_allocation_params", fields, reflect.TypeOf(IpAllocationParams{}), fieldNameMap, validators)
}

func OvfMessageBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessageCategory(OvfMessageCategory_VALIDATION)))
	fieldNameMap["category"] = "Category"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
	fieldNameMap["issues"] = "Issues"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
	fieldNameMap["error"] = "Error_"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("category",
		map[string][]bindings.FieldData{
			"VALIDATION": []bindings.FieldData{
				bindings.NewFieldData("issues", true),
			},
			"INPUT": []bindings.FieldData{
				bindings.NewFieldData("name", true),
				bindings.NewFieldData("value", true),
				bindings.NewFieldData("message", true),
			},
			"SERVER": []bindings.FieldData{
				bindings.NewFieldData("error", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_message", fields, reflect.TypeOf(OvfMessage{}), fieldNameMap, validators)
}

func ParseIssueBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.parse_issue.category", reflect.TypeOf(ParseIssueCategory(ParseIssueCategory_VALUE_ILLEGAL)))
	fieldNameMap["category"] = "Category"
	fields["file"] = bindings.NewStringType()
	fieldNameMap["file"] = "File"
	fields["line_number"] = bindings.NewIntegerType()
	fieldNameMap["line_number"] = "LineNumber"
	fields["column_number"] = bindings.NewIntegerType()
	fieldNameMap["column_number"] = "ColumnNumber"
	fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
	fieldNameMap["message"] = "Message"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.parse_issue", fields, reflect.TypeOf(ParseIssue{}), fieldNameMap, validators)
}

func OvfErrorBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessageCategory(OvfMessageCategory_VALIDATION)))
	fieldNameMap["category"] = "Category"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
	fieldNameMap["issues"] = "Issues"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
	fieldNameMap["error"] = "Error_"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("category",
		map[string][]bindings.FieldData{
			"VALIDATION": []bindings.FieldData{
				bindings.NewFieldData("issues", true),
			},
			"INPUT": []bindings.FieldData{
				bindings.NewFieldData("name", true),
				bindings.NewFieldData("value", true),
				bindings.NewFieldData("message", true),
			},
			"SERVER": []bindings.FieldData{
				bindings.NewFieldData("error", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_error", fields, reflect.TypeOf(OvfError{}), fieldNameMap, validators)
}

func OvfWarningBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessageCategory(OvfMessageCategory_VALIDATION)))
	fieldNameMap["category"] = "Category"
	fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
	fieldNameMap["issues"] = "Issues"
	fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["name"] = "Name"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
	fieldNameMap["message"] = "Message"
	fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType{bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
	fieldNameMap["error"] = "Error_"
	var validators = []bindings.Validator{}
	uv1 := bindings.NewUnionValidator("category",
		map[string][]bindings.FieldData{
			"VALIDATION": []bindings.FieldData{
				bindings.NewFieldData("issues", true),
			},
			"INPUT": []bindings.FieldData{
				bindings.NewFieldData("name", true),
				bindings.NewFieldData("value", true),
				bindings.NewFieldData("message", true),
			},
			"SERVER": []bindings.FieldData{
				bindings.NewFieldData("error", true),
			},
		},
	)
	validators = append(validators, uv1)
	return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_warning", fields, reflect.TypeOf(OvfWarning{}), fieldNameMap, validators)
}

func OvfInfoBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
	fieldNameMap["messages"] = "Messages"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_info", fields, reflect.TypeOf(OvfInfo{}), fieldNameMap, validators)
}

func OvfParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_params", fields, reflect.TypeOf(OvfParams{}), fieldNameMap, validators)
}

func PropertyBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["class_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["class_id"] = "ClassId"
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["instance_id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["instance_id"] = "InstanceId"
	fields["category"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["category"] = "Category"
	fields["ui_optional"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["ui_optional"] = "UiOptional"
	fields["label"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["label"] = "Label"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["value"] = "Value"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.property", fields, reflect.TypeOf(Property{}), fieldNameMap, validators)
}

func PropertyParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["properties"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PropertyBindingType), reflect.TypeOf([]Property{})))
	fieldNameMap["properties"] = "Properties"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.property_params", fields, reflect.TypeOf(PropertyParams{}), fieldNameMap, validators)
}

func ScaleOutGroupBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["id"] = "Id"
	fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["description"] = "Description"
	fields["instance_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["instance_count"] = "InstanceCount"
	fields["minimum_instance_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["minimum_instance_count"] = "MinimumInstanceCount"
	fields["maximum_instance_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["maximum_instance_count"] = "MaximumInstanceCount"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.scale_out_group", fields, reflect.TypeOf(ScaleOutGroup{}), fieldNameMap, validators)
}

func ScaleOutParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ScaleOutGroupBindingType), reflect.TypeOf([]ScaleOutGroup{})))
	fieldNameMap["groups"] = "Groups"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.scale_out_params", fields, reflect.TypeOf(ScaleOutParams{}), fieldNameMap, validators)
}

func SizeParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["approximate_download_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["approximate_download_size"] = "ApproximateDownloadSize"
	fields["approximate_flat_deployment_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["approximate_flat_deployment_size"] = "ApproximateFlatDeploymentSize"
	fields["approximate_sparse_deployment_size"] = bindings.NewOptionalType(bindings.NewIntegerType())
	fieldNameMap["approximate_sparse_deployment_size"] = "ApproximateSparseDeploymentSize"
	fields["variable_disk_size"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["variable_disk_size"] = "VariableDiskSize"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.size_params", fields, reflect.TypeOf(SizeParams{}), fieldNameMap, validators)
}

func UnknownSectionBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["tag"] = bindings.NewStringType()
	fieldNameMap["tag"] = "Tag"
	fields["info"] = bindings.NewStringType()
	fieldNameMap["info"] = "Info"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.unknown_section", fields, reflect.TypeOf(UnknownSection{}), fieldNameMap, validators)
}

func UnknownSectionParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["unknown_sections"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(UnknownSectionBindingType), reflect.TypeOf([]UnknownSection{})))
	fieldNameMap["unknown_sections"] = "UnknownSections"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.unknown_section_params", fields, reflect.TypeOf(UnknownSectionParams{}), fieldNameMap, validators)
}

func VcenterExtensionParamsBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["required"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["required"] = "Required"
	fields["registration_accepted"] = bindings.NewOptionalType(bindings.NewBooleanType())
	fieldNameMap["registration_accepted"] = "RegistrationAccepted"
	fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["type"] = "Type_"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.ovf.vcenter_extension_params", fields, reflect.TypeOf(VcenterExtensionParams{}), fieldNameMap, validators)
}


