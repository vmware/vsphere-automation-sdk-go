/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
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
    "net/url"
)


// The ``WarningType`` enumeration class defines the warnings which can be raised during the OVF package deployment. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type WarningType string

const (
    // The certificate used for signing the OVF package content is self-signed. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     WarningType_SELF_SIGNED_CERTIFICATE WarningType = "SELF_SIGNED_CERTIFICATE"
    // The certificate used for signing the OVF package content is expired. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     WarningType_EXPIRED_CERTIFICATE WarningType = "EXPIRED_CERTIFICATE"
    // The certificate used for signing the OVF package content is not yet valid. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     WarningType_NOT_YET_VALID_CERTIFICATE WarningType = "NOT_YET_VALID_CERTIFICATE"
    // The certificate used for signing the OVF package content is not trusted. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     WarningType_UNTRUSTED_CERTIFICATE WarningType = "UNTRUSTED_CERTIFICATE"
)

func (w WarningType) WarningType() bool {
    switch w {
        case WarningType_SELF_SIGNED_CERTIFICATE:
            return true
        case WarningType_EXPIRED_CERTIFICATE:
            return true
        case WarningType_NOT_YET_VALID_CERTIFICATE:
            return true
        case WarningType_UNTRUSTED_CERTIFICATE:
            return true
        default:
            return false
    }
}




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
    Issuer *string
    Subject *string
    IsValid *bool
    IsSelfSigned *bool
    X509 *string
    Warnings []WarningInfo
    Type_ *string
}






// The ``WarningInfo`` class provides information about the warnings which are raised during the OVF package deployment. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type WarningInfo struct {
    Type_ WarningType
    Message std.LocalizableMessage
    Ignored bool
}






// The ``DatastoreDiskGroup`` class contains information about a storage disk group described in the OVF package. 
//
//  This class will always be returned when retrieving information about an OVF package. See DatastoreMappingParams#diskGroups. 
//
//  A disk group is a main unit of storage on a ESXi host. Each disk group includes one SSD and one or multiple HDDs (magnetic disks). 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type DatastoreDiskGroup struct {
    Id *string
    Name *string
    Description *string
    TargetProfile *string
    TargetDatastore *string
    TargetProvisioningType *DiskProvisioningType
}






// The ``DatastoreTarget`` class contains information about a datastore or datastore cluster in the target deployment environment. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type DatastoreTarget struct {
    Id *string
    Name *string
    Accessible *bool
    InaccessibleReasons []DatastoreTarget_InaccessibleReason
    StorageProfiles []string
}




    
    // The ``InaccessibleReason`` enumeration class defines the reasons why a datastore can be inaccessible.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type DatastoreTarget_InaccessibleReason string

    const (
        // The datastore is not connected.
         DatastoreTarget_InaccessibleReason_NOT_CONNECTED DatastoreTarget_InaccessibleReason = "NOT_CONNECTED"
        // The datastore is in maintenance mode.
         DatastoreTarget_InaccessibleReason_MAINTENANCE_MODE DatastoreTarget_InaccessibleReason = "MAINTENANCE_MODE"
        // The datastore is mounted read-only.
         DatastoreTarget_InaccessibleReason_READ_ONLY DatastoreTarget_InaccessibleReason = "READ_ONLY"
        // The current logged on user does not have Datastore.AllocateSpace privilege which is required to allocate to the datastore.
         DatastoreTarget_InaccessibleReason_NO_ALLOCATE_RIGHT DatastoreTarget_InaccessibleReason = "NO_ALLOCATE_RIGHT"
    )

    func (i DatastoreTarget_InaccessibleReason) DatastoreTarget_InaccessibleReason() bool {
        switch i {
            case DatastoreTarget_InaccessibleReason_NOT_CONNECTED:
                return true
            case DatastoreTarget_InaccessibleReason_MAINTENANCE_MODE:
                return true
            case DatastoreTarget_InaccessibleReason_READ_ONLY:
                return true
            case DatastoreTarget_InaccessibleReason_NO_ALLOCATE_RIGHT:
                return true
            default:
                return false
        }
    }



// The ``DatastoreMappingParams`` class contains information about the mapping from OVF storage groups to specific storage targets supported on the target platform. 
//
//  This class will only be used when deploying an OVF package to a resource pool, cluster or host. 
//
//  This is information based on the vmw:StorageGroupSection and vmw:StorageSection. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type DatastoreMappingParams struct {
    AvailableStorageProfiles []StorageProfileTarget
    StorageProfilesEnabled *bool
    AvailableDatastores []DatastoreTarget
    AvailableDiskProvisioningTypes []DiskProvisioningType
    TargetProfile *string
    TargetDatastore *string
    TargetProvisioningType *DiskProvisioningType
    DiskGroups []DatastoreDiskGroup
    Type_ *string
}






// The ``DeploymentOption`` class contains the information about a deployment option as defined in the OVF specification. 
//
//  This corresponds to the ovf:Configuration element of the ovf:DeploymentOptionSection in the specification. The ovf:DeploymentOptionSection specifies a discrete set of intended resource allocation configurations. This class represents one item from that set. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type DeploymentOption struct {
    Key *string
    Label *string
    Description *string
    DefaultChoice *bool
}






// The ``DeploymentOptionParams`` class describes the possible deployment options as well as the choice provided by the user. 
//
//  This information based on the ovf:DeploymentOptionSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type DeploymentOptionParams struct {
    DeploymentOptions []DeploymentOption
    SelectedKey *string
    Type_ *string
}






// The ``EulaParams`` class includes a array of end user licence agreements (EULAs) included in the OVF package. 
//
//  If the OVF package includes one or more EULAs, an import will only succeed if the ``allEULAAccepted`` property boolean is set to true. 
//
//  This information based on the ovf:EulaSection. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type EulaParams struct {
    Eulas []string
    AllEULAAccepted *bool
    Type_ *string
}






// The ``ExtraConfig`` class contains the information about a vmw:ExtraConfig element which can be used to specify configuration settings that are transferred directly to the ``.vmx`` file. The behavior of the vmw:ExtraConfig element is similar to the ``extraConfig`` property of the ``VirtualMachineConfigSpec`` object in the VMware vSphere API. Thus, the same restrictions apply, such as you cannot set values that could otherwise be set with other properties in the ``VirtualMachineConfigSpec`` object. See the VMware vSphere API reference for details on this. 
//
//  vmw:ExtraConfig elements may occur as direct child elements of a VirtualHardwareSection, or as child elements of individual virtual hardware items. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type ExtraConfig struct {
    Key *string
    Value *string
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
    ExtraConfigs []ExtraConfig
    ExcludeKeys []string
    IncludeKeys []string
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
    SupportedAllocationScheme []IpAllocationParams_IpAllocationScheme
    SupportedIpAllocationPolicy []IpAllocationParams_IpAllocationPolicy
    IpAllocationPolicy *IpAllocationParams_IpAllocationPolicy
    SupportedIpProtocol []IpAllocationParams_IpProtocol
    IpProtocol *IpAllocationParams_IpProtocol
    Type_ *string
}




    
    // The ``IpAllocationPolicy`` enumeration class defines the possible IP allocation policy for a deployment.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type IpAllocationParams_IpAllocationPolicy string

    const (
        // Specifies that DHCP will be used to allocate IP addresses.
         IpAllocationParams_IpAllocationPolicy_DHCP IpAllocationParams_IpAllocationPolicy = "DHCP"
        // Specifies that IP addresses are allocated from an IP pool. The IP addresses are allocated when needed, typically at power-on, and deallocated during power-off. There is no guarantee that a property will receive same IP address when restarted.
         IpAllocationParams_IpAllocationPolicy_TRANSIENT_IPPOOL IpAllocationParams_IpAllocationPolicy = "TRANSIENT_IPPOOL"
        // Specifies that IP addresses are configured manually upon deployment, and will be kept until reconfigured or the virtual appliance destroyed. This ensures that a property gets a consistent IP for its lifetime.
         IpAllocationParams_IpAllocationPolicy_STATIC_MANUAL IpAllocationParams_IpAllocationPolicy = "STATIC_MANUAL"
        // Specifies that IP addresses are allocated from the range managed by an IP pool. The IP addresses are allocated at first power-on, and remain allocated at power-off. This ensures that a virtual appliance gets a consistent IP for its life-time.
         IpAllocationParams_IpAllocationPolicy_STATIC_IPPOOL IpAllocationParams_IpAllocationPolicy = "STATIC_IPPOOL"
    )

    func (i IpAllocationParams_IpAllocationPolicy) IpAllocationParams_IpAllocationPolicy() bool {
        switch i {
            case IpAllocationParams_IpAllocationPolicy_DHCP:
                return true
            case IpAllocationParams_IpAllocationPolicy_TRANSIENT_IPPOOL:
                return true
            case IpAllocationParams_IpAllocationPolicy_STATIC_MANUAL:
                return true
            case IpAllocationParams_IpAllocationPolicy_STATIC_IPPOOL:
                return true
            default:
                return false
        }
    }

    
    // The ``IpAllocationScheme`` enumeration class defines the possible IP allocation schemes that can be supported by the guest software.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type IpAllocationParams_IpAllocationScheme string

    const (
        // It supports DHCP to acquire IP configuration.
         IpAllocationParams_IpAllocationScheme_DHCP IpAllocationParams_IpAllocationScheme = "DHCP"
        // It supports setting the IP configuration through the properties provided in the OVF environment.
         IpAllocationParams_IpAllocationScheme_OVF_ENVIRONMENT IpAllocationParams_IpAllocationScheme = "OVF_ENVIRONMENT"
    )

    func (i IpAllocationParams_IpAllocationScheme) IpAllocationParams_IpAllocationScheme() bool {
        switch i {
            case IpAllocationParams_IpAllocationScheme_DHCP:
                return true
            case IpAllocationParams_IpAllocationScheme_OVF_ENVIRONMENT:
                return true
            default:
                return false
        }
    }

    
    // The ``IpProtocol`` enumeration class defines the IP protocols supported by the guest software.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type IpAllocationParams_IpProtocol string

    const (
        // It supports the IPv4 protocol.
         IpAllocationParams_IpProtocol_IPV4 IpAllocationParams_IpProtocol = "IPV4"
        // It supports the IPv6 protocol.
         IpAllocationParams_IpProtocol_IPV6 IpAllocationParams_IpProtocol = "IPV6"
    )

    func (i IpAllocationParams_IpProtocol) IpAllocationParams_IpProtocol() bool {
        switch i {
            case IpAllocationParams_IpProtocol_IPV4:
                return true
            case IpAllocationParams_IpProtocol_IPV6:
                return true
            default:
                return false
        }
    }



// The ``NameAndProductParams`` class contains name, description, and product information about an OVF package. The information in this package is derived from the ovf:AnnotationSection, ovf:ProductSection, and the name/id information on the root content type. 
//
//  The name/description can be changed/overwritten during deployment. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type NameAndProductParams struct {
    Name *string
    Annotation *string
    VirtualApp *bool
    ProductName *string
    Vendor *string
    Version *string
    FullVersion *string
    VendorUrl *url.URL
    ProductUrl *url.URL
    Type_ *string
}






// The ``SourceNetwork`` class contains information about a network in the OVF package. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type SourceNetwork struct {
    Name *string
    Description *string
    Target *string
}






// The ``TargetNetwork`` class contains information about a network in the deployment environment. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type TargetNetwork struct {
    Id *string
    Accessible *bool
    InaccessibleReasons []TargetNetwork_InaccessibleReason
    Name *string
}




    
    // The ``InaccessibleReason`` enumeration class defines the reasons why a network can be inaccessible.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type TargetNetwork_InaccessibleReason string

    const (
        // The current logged on user does not have Network.Assign privilege which is required to assign to this network.
         TargetNetwork_InaccessibleReason_NO_ASSIGN_RIGHT TargetNetwork_InaccessibleReason = "NO_ASSIGN_RIGHT"
    )

    func (i TargetNetwork_InaccessibleReason) TargetNetwork_InaccessibleReason() bool {
        switch i {
            case TargetNetwork_InaccessibleReason_NO_ASSIGN_RIGHT:
                return true
            default:
                return false
        }
    }



// The ``NetworkMappingParams`` class specifies how source networks in the OVF package are mapped to target networks in the deployment environment. A mapping must be provided for each source network or the deployment will fail. 
//
//  This is based on the ovf:NetworkSection. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type NetworkMappingParams struct {
    SourceNetworks []SourceNetwork
    TargetNetworks []TargetNetwork
    Type_ *string
}






// The ``OvfMessage`` class describes a base OVF handling error message related to accessing, validating, deploying, or exporting an OVF package. 
//
//  These messages fall into different categories defined in OvfMessage_Category:
type OvfMessage struct {
    Category OvfMessage_Category
    Issues []ParseIssue
    Name *string
    Value *string
    Message *std.LocalizableMessage
    Error *data.StructValue
}




    
    // The ``Category`` enumeration class defines the categories of messages (see OvfMessage).
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type OvfMessage_Category string

    const (
        // The OVF descriptor is invalid, for example, syntax errors or schema errors.
         OvfMessage_Category_VALIDATION OvfMessage_Category = "VALIDATION"
        // The user provided input parameters are invalid.
         OvfMessage_Category_INPUT OvfMessage_Category = "INPUT"
        // Server error.
         OvfMessage_Category_SERVER OvfMessage_Category = "SERVER"
    )

    func (c OvfMessage_Category) OvfMessage_Category() bool {
        switch c {
            case OvfMessage_Category_VALIDATION:
                return true
            case OvfMessage_Category_INPUT:
                return true
            case OvfMessage_Category_SERVER:
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
    Category ParseIssue_Category
    File string
    LineNumber int64
    ColumnNumber int64
    Message std.LocalizableMessage
}




    
    // The ``Category`` enumeration class defines the categories of issues that can be found when parsing files inside an OVF package (see ParseIssue) including OVF descriptor (which is an XML document), manifest and certificate files, or exporting an OVF package.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type ParseIssue_Category string

    const (
        // Illegal value error. For example, the value is malformed, not a number, or outside of the given range, and so on.
         ParseIssue_Category_VALUE_ILLEGAL ParseIssue_Category = "VALUE_ILLEGAL"
        // Required attribute error. It indicates that a required attribute is missing from an element in the OVF descriptor.
         ParseIssue_Category_ATTRIBUTE_REQUIRED ParseIssue_Category = "ATTRIBUTE_REQUIRED"
        // Illegal attribute error. It indicates that an illegal attribute is set for an element in the OVF descriptor. For example, empty disks do not use format, parentRef, and populatedSize attributes, if these attributes are present in an empty disk element then will get this pasrse issue.
         ParseIssue_Category_ATTRIBUTE_ILLEGAL ParseIssue_Category = "ATTRIBUTE_ILLEGAL"
        // Required element error. It indicates that a required element is missing from the OVF descriptor.
         ParseIssue_Category_ELEMENT_REQUIRED ParseIssue_Category = "ELEMENT_REQUIRED"
        // Illegal element error. It indicates that an element is present in a location which is not allowed, or found multiple elements but only one is allowed at the location in the OVF descriptor.
         ParseIssue_Category_ELEMENT_ILLEGAL ParseIssue_Category = "ELEMENT_ILLEGAL"
        // Unknown element error. It indicates that an element is unsupported when parsing an OVF descriptor.
         ParseIssue_Category_ELEMENT_UNKNOWN ParseIssue_Category = "ELEMENT_UNKNOWN"
        // Section unknown error. It indicates that a section is unsupported when parsing an OVF descriptor.
         ParseIssue_Category_SECTION_UNKNOWN ParseIssue_Category = "SECTION_UNKNOWN"
        // Section restriction error. It indicates that a section appears in place in the OVF descriptor where it is not allowed, a section appears fewer times than is required, or a section appears more times than is allowed.
         ParseIssue_Category_SECTION_RESTRICTION ParseIssue_Category = "SECTION_RESTRICTION"
        // OVF package parsing error, including: 
        //
        // * OVF descriptor parsing errors, for example, syntax errors or schema errors.
        // * Manifest file parsing and verification errors.
        // * Certificate file parsing and verification errors.
         ParseIssue_Category_PARSE_ERROR ParseIssue_Category = "PARSE_ERROR"
        // OVF descriptor (which is an XML document) generating error, for example, well-formedness errors as well as unexpected processing conditions.
         ParseIssue_Category_GENERATE_ERROR ParseIssue_Category = "GENERATE_ERROR"
        // An issue with the manifest and signing.
         ParseIssue_Category_VALIDATION_ERROR ParseIssue_Category = "VALIDATION_ERROR"
        // Issue during OVF export, for example, malformed deviceId, controller not found, or file backing for a device not found.
         ParseIssue_Category_EXPORT_ERROR ParseIssue_Category = "EXPORT_ERROR"
        // Server encountered an unexpected error which prevented it from fulfilling the request.
         ParseIssue_Category_INTERNAL_ERROR ParseIssue_Category = "INTERNAL_ERROR"
    )

    func (c ParseIssue_Category) ParseIssue_Category() bool {
        switch c {
            case ParseIssue_Category_VALUE_ILLEGAL:
                return true
            case ParseIssue_Category_ATTRIBUTE_REQUIRED:
                return true
            case ParseIssue_Category_ATTRIBUTE_ILLEGAL:
                return true
            case ParseIssue_Category_ELEMENT_REQUIRED:
                return true
            case ParseIssue_Category_ELEMENT_ILLEGAL:
                return true
            case ParseIssue_Category_ELEMENT_UNKNOWN:
                return true
            case ParseIssue_Category_SECTION_UNKNOWN:
                return true
            case ParseIssue_Category_SECTION_RESTRICTION:
                return true
            case ParseIssue_Category_PARSE_ERROR:
                return true
            case ParseIssue_Category_GENERATE_ERROR:
                return true
            case ParseIssue_Category_VALIDATION_ERROR:
                return true
            case ParseIssue_Category_EXPORT_ERROR:
                return true
            case ParseIssue_Category_INTERNAL_ERROR:
                return true
            default:
                return false
        }
    }



// The ``OvfError`` class describes an error related to accessing, validating, deploying, or exporting an OVF package.
type OvfError struct {
    Category OvfMessage_Category
    Issues []ParseIssue
    Name *string
    Value *string
    Message *std.LocalizableMessage
    Error *data.StructValue
}






// The ``OvfWarning`` class describes a warning related to accessing, validating, deploying, or exporting an OVF package.
type OvfWarning struct {
    Category OvfMessage_Category
    Issues []ParseIssue
    Name *string
    Value *string
    Message *std.LocalizableMessage
    Error *data.StructValue
}






// The ``OvfInfo`` class contains informational messages related to accessing, validating, deploying, or exporting an OVF package.
type OvfInfo struct {
    Messages []std.LocalizableMessage
}






// The ``OvfFileInfo`` class contains the information regarding a single file in an OVF package or an OVA package file.
type OvfFileInfo struct {
    Name string
    FileType OvfFileInfo_FileType
    OptionalUpload bool
    FileUrl *url.URL
    SslThumbPrint *string
    UploadMethod *OvfFileInfo_HttpUploadMethod
    Size int64
    BytesTransferred int64
    Sha256 string
}




    
    // The ``FileType`` enumeration class indicates files that have special semantics in an OVF package, or indicates an OVA package file.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type OvfFileInfo_FileType string

    const (
        // The file is an OVF descriptor.
         OvfFileInfo_FileType_OVF OvfFileInfo_FileType = "OVF"
        // The file is a manifest.
         OvfFileInfo_FileType_MANIFEST OvfFileInfo_FileType = "MANIFEST"
        // The file is a certificate.
         OvfFileInfo_FileType_CERT OvfFileInfo_FileType = "CERT"
        // A file with localized messages for the OVF.
         OvfFileInfo_FileType_MSG_BUNDLE OvfFileInfo_FileType = "MSG_BUNDLE"
        // The file is a disk.
         OvfFileInfo_FileType_DISK OvfFileInfo_FileType = "DISK"
        // Anything else a file that does not need special treatment during transfer.
         OvfFileInfo_FileType_CONTENT OvfFileInfo_FileType = "CONTENT"
        // The file is an OVA file.
         OvfFileInfo_FileType_OVA OvfFileInfo_FileType = "OVA"
        // The file is a NVRAM file for EFI boot support.
         OvfFileInfo_FileType_NVRAM OvfFileInfo_FileType = "NVRAM"
    )

    func (f OvfFileInfo_FileType) OvfFileInfo_FileType() bool {
        switch f {
            case OvfFileInfo_FileType_OVF:
                return true
            case OvfFileInfo_FileType_MANIFEST:
                return true
            case OvfFileInfo_FileType_CERT:
                return true
            case OvfFileInfo_FileType_MSG_BUNDLE:
                return true
            case OvfFileInfo_FileType_DISK:
                return true
            case OvfFileInfo_FileType_CONTENT:
                return true
            case OvfFileInfo_FileType_OVA:
                return true
            case OvfFileInfo_FileType_NVRAM:
                return true
            default:
                return false
        }
    }

    
    // HTTP method that can be used to upload file content.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type OvfFileInfo_HttpUploadMethod string

    const (
        // HTTP PUT method
         OvfFileInfo_HttpUploadMethod_PUT OvfFileInfo_HttpUploadMethod = "PUT"
        // HTTP POST method
         OvfFileInfo_HttpUploadMethod_POST OvfFileInfo_HttpUploadMethod = "POST"
    )

    func (h OvfFileInfo_HttpUploadMethod) OvfFileInfo_HttpUploadMethod() bool {
        switch h {
            case OvfFileInfo_HttpUploadMethod_PUT:
                return true
            case OvfFileInfo_HttpUploadMethod_POST:
                return true
            default:
                return false
        }
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
    Type_ *string
}






// The ``Property`` class contains the information about a property in an OVF package. 
//
//  A property is uniquely identified by its [classid.]id[.instanceid] fully-qualified name (see Property#classId, Property#id, and Property#instanceId). If multiple properties in an OVF package have the same fully-qualified name, then the property is excluded and cannot be set. We do warn about this during import. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type Property struct {
    ClassId *string
    Id *string
    InstanceId *string
    Category *string
    UiOptional *bool
    Label *string
    Description *string
    Type_ *string
    Value *string
}






// The ``PropertyParams`` class contains a array of OVF properties that can be configured when the OVF package is deployed. 
//
//  This is based on the ovf:ProductSection. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type PropertyParams struct {
    Properties []Property
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
    Id *string
    Description *string
    InstanceCount *int64
    MinimumInstanceCount *int64
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
    Groups []ScaleOutGroup
    Type_ *string
}






// The ``SizeParams`` class contains estimates of the download and deployment sizes. 
//
//  This information is based on the file references and the ovf:DiskSection in the OVF descriptor. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type SizeParams struct {
    ApproximateDownloadSize *int64
    ApproximateFlatDeploymentSize *int64
    ApproximateSparseDeploymentSize *int64
    VariableDiskSize *bool
    Type_ *string
}






// The ``StorageDiskGroup`` class contains the information about a storage disk group described in the OVF package. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type StorageDiskGroup struct {
    Id *string
    Name *string
    Description *string
    TargetProfile *string
    TargetProvisioningType *DiskProvisioningType
}






// The ``StorageProfileTarget`` class contains information about a storage policy. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type StorageProfileTarget struct {
    Id *string
    Name *string
}






// The ``StorageMappingParams`` class maps the OVF storage groups to specific storage profiles available on the target platform. 
//
//  This class will only be used when deploying an OVF package to a virtual datacenter. 
//
//  This is information based on the vmw:StorageGroupSection and vmw:StorageSection. 
//
//  See null, null, LibraryItem#deploy, and LibraryItem#filter.
type StorageMappingParams struct {
    AvailableStorageProfiles []StorageProfileTarget
    AvailableDiskProvisioningTypes []DiskProvisioningType
    TargetProfile *string
    TargetProvisioningType *DiskProvisioningType
    DiskGroups []StorageDiskGroup
    Type_ *string
}






// The ``UnknownSection`` class contains information about an unknown section in an OVF package.
type UnknownSection struct {
    Tag string
    Info string
}






// The ``UnknownSectionParams`` class contains a array of unknown, non-required sections. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type UnknownSectionParams struct {
    UnknownSections []UnknownSection
    Type_ *string
}






// Information about a vService provider.
type VServiceProvider struct {
    Key *string
    Name *string
    Description *string
    AutoBind *bool
    Status *VServiceProvider_BindingStatus
    ValidationMessage *string
}




    
    // The state of compatibility between a vService dependency and a vService provider.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type VServiceProvider_BindingStatus string

    const (
        // The dependency on the provider cannot be bound.
         VServiceProvider_BindingStatus_RED VServiceProvider_BindingStatus = "RED"
        // The dependency on the provider can be bound, but with warnings.
         VServiceProvider_BindingStatus_YELLOW VServiceProvider_BindingStatus = "YELLOW"
        // The dependency on the provider can be bound.
         VServiceProvider_BindingStatus_GREEN VServiceProvider_BindingStatus = "GREEN"
    )

    func (b VServiceProvider_BindingStatus) VServiceProvider_BindingStatus() bool {
        switch b {
            case VServiceProvider_BindingStatus_RED:
                return true
            case VServiceProvider_BindingStatus_YELLOW:
                return true
            case VServiceProvider_BindingStatus_GREEN:
                return true
            default:
                return false
        }
    }



// Information about a vService dependency.
type VServiceDependency struct {
    Id *string
    Type_ *string
    Name *string
    Description *string
    Required *bool
    SelectedProviderKey *string
    AvailableProviders []VServiceProvider
}






// List of vService dependencies that can be configured at deployment time.
type VServiceParams struct {
    Dependencies []VServiceDependency
    Type_ *string
}






// The ``VcenterExtensionParams`` class specifies that the OVF package should be registered as a vCenter extension. The virtual machine or virtual appliance will gain unrestricted access to the vCenter Server APIs. It must be connected to a network with connectivity to the vCenter server. 
//
//  See LibraryItem#deploy and LibraryItem#filter.
type VcenterExtensionParams struct {
    Required *bool
    RegistrationAccepted *bool
    Type_ *string
}






// The ``VcenterGuestCustomization`` class defines the customization of the guest operating system for a virtual machine specified by section of type ovf:VirtualSystem in the OVF descriptor.
type VcenterGuestCustomization struct {
    Type_ VcenterGuestCustomization_Type
    SpecificationId *string
    Xml *string
}




    
    // The ``Type`` enumeration class defines the supported types of guest customization schemes for sections of type ovf:VirtualSystem in the OVF descriptor.
    //
    // <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
     
    type VcenterGuestCustomization_Type string

    const (
        // Customization setting is encapsulated in a vCenter Server guest operating system customization specification.
         VcenterGuestCustomization_Type_SPECIFICATION VcenterGuestCustomization_Type = "SPECIFICATION"
        // Customization setting is encapsulated in an XML document- formatted string. The content of the string describes a vCenter server guest operating system customization specification.
         VcenterGuestCustomization_Type_XML VcenterGuestCustomization_Type = "XML"
    )

    func (t VcenterGuestCustomization_Type) VcenterGuestCustomization_Type() bool {
        switch t {
            case VcenterGuestCustomization_Type_SPECIFICATION:
                return true
            case VcenterGuestCustomization_Type_XML:
                return true
            default:
                return false
        }
    }



// The ``VcenterGuestCustomizationParams`` class contains the parameters with information about customization of ovf:VirtualSystem using vCenter guest OS customization specifications. 
type VcenterGuestCustomizationParams struct {
    VirtualSystems []string
    Customizations map[string]VcenterGuestCustomization
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
    fields["warnings"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(WarningInfoBindingType), reflect.TypeOf([]WarningInfo{})))
    fieldNameMap["warnings"] = "Warnings"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.certificate_params",fields, reflect.TypeOf(CertificateParams{}), fieldNameMap, validators)
}

func WarningInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.warning_type", reflect.TypeOf(WarningType(WarningType_SELF_SIGNED_CERTIFICATE)))
    fieldNameMap["type"] = "Type_"
    fields["message"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["message"] = "Message"
    fields["ignored"] = bindings.NewBooleanType()
    fieldNameMap["ignored"] = "Ignored"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.warning_info",fields, reflect.TypeOf(WarningInfo{}), fieldNameMap, validators)
}

func DatastoreDiskGroupBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["target_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["target_profile"] = "TargetProfile"
    fields["target_datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["target_datastore"] = "TargetDatastore"
    fields["target_provisioning_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
    fieldNameMap["target_provisioning_type"] = "TargetProvisioningType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.datastore_disk_group",fields, reflect.TypeOf(DatastoreDiskGroup{}), fieldNameMap, validators)
}

func DatastoreTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["accessible"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["accessible"] = "Accessible"
    fields["inaccessible_reasons"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.datastore_target.inaccessible_reason", reflect.TypeOf(DatastoreTarget_InaccessibleReason(DatastoreTarget_InaccessibleReason_NOT_CONNECTED))), reflect.TypeOf([]DatastoreTarget_InaccessibleReason{})))
    fieldNameMap["inaccessible_reasons"] = "InaccessibleReasons"
    fields["storage_profiles"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewIdType([]string {"StorageProfile"}, ""), reflect.TypeOf([]string{})))
    fieldNameMap["storage_profiles"] = "StorageProfiles"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.datastore_target",fields, reflect.TypeOf(DatastoreTarget{}), fieldNameMap, validators)
}

func DatastoreMappingParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["available_storage_profiles"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageProfileTargetBindingType), reflect.TypeOf([]StorageProfileTarget{})))
    fieldNameMap["available_storage_profiles"] = "AvailableStorageProfiles"
    fields["storage_profiles_enabled"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["storage_profiles_enabled"] = "StorageProfilesEnabled"
    fields["available_datastores"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DatastoreTargetBindingType), reflect.TypeOf([]DatastoreTarget{})))
    fieldNameMap["available_datastores"] = "AvailableDatastores"
    fields["available_disk_provisioning_types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))), reflect.TypeOf([]DiskProvisioningType{})))
    fieldNameMap["available_disk_provisioning_types"] = "AvailableDiskProvisioningTypes"
    fields["target_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["target_profile"] = "TargetProfile"
    fields["target_datastore"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Datastore"}, ""))
    fieldNameMap["target_datastore"] = "TargetDatastore"
    fields["target_provisioning_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
    fieldNameMap["target_provisioning_type"] = "TargetProvisioningType"
    fields["disk_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DatastoreDiskGroupBindingType), reflect.TypeOf([]DatastoreDiskGroup{})))
    fieldNameMap["disk_groups"] = "DiskGroups"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.datastore_mapping_params",fields, reflect.TypeOf(DatastoreMappingParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.deployment_option",fields, reflect.TypeOf(DeploymentOption{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.deployment_option_params",fields, reflect.TypeOf(DeploymentOptionParams{}), fieldNameMap, validators)
}

func EulaParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["eulas"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["eulas"] = "Eulas"
    fields["all_EULA_accepted"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["all_EULA_accepted"] = "AllEULAAccepted"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.eula_params",fields, reflect.TypeOf(EulaParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.extra_config",fields, reflect.TypeOf(ExtraConfig{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.extra_config_params",fields, reflect.TypeOf(ExtraConfigParams{}), fieldNameMap, validators)
}

func IpAllocationParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["supported_allocation_scheme"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_scheme", reflect.TypeOf(IpAllocationParams_IpAllocationScheme(IpAllocationParams_IpAllocationScheme_DHCP))), reflect.TypeOf([]IpAllocationParams_IpAllocationScheme{})))
    fieldNameMap["supported_allocation_scheme"] = "SupportedAllocationScheme"
    fields["supported_ip_allocation_policy"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_policy", reflect.TypeOf(IpAllocationParams_IpAllocationPolicy(IpAllocationParams_IpAllocationPolicy_DHCP))), reflect.TypeOf([]IpAllocationParams_IpAllocationPolicy{})))
    fieldNameMap["supported_ip_allocation_policy"] = "SupportedIpAllocationPolicy"
    fields["ip_allocation_policy"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_allocation_policy", reflect.TypeOf(IpAllocationParams_IpAllocationPolicy(IpAllocationParams_IpAllocationPolicy_DHCP))))
    fieldNameMap["ip_allocation_policy"] = "IpAllocationPolicy"
    fields["supported_ip_protocol"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_protocol", reflect.TypeOf(IpAllocationParams_IpProtocol(IpAllocationParams_IpProtocol_IPV4))), reflect.TypeOf([]IpAllocationParams_IpProtocol{})))
    fieldNameMap["supported_ip_protocol"] = "SupportedIpProtocol"
    fields["ip_protocol"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.ip_allocation_params.ip_protocol", reflect.TypeOf(IpAllocationParams_IpProtocol(IpAllocationParams_IpProtocol_IPV4))))
    fieldNameMap["ip_protocol"] = "IpProtocol"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.ip_allocation_params",fields, reflect.TypeOf(IpAllocationParams{}), fieldNameMap, validators)
}

func NameAndProductParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["annotation"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["annotation"] = "Annotation"
    fields["virtual_app"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["virtual_app"] = "VirtualApp"
    fields["product_name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["product_name"] = "ProductName"
    fields["vendor"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["vendor"] = "Vendor"
    fields["version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["version"] = "Version"
    fields["full_version"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["full_version"] = "FullVersion"
    fields["vendor_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["vendor_url"] = "VendorUrl"
    fields["product_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["product_url"] = "ProductUrl"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.name_and_product_params",fields, reflect.TypeOf(NameAndProductParams{}), fieldNameMap, validators)
}

func SourceNetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["target"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["target"] = "Target"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.source_network",fields, reflect.TypeOf(SourceNetwork{}), fieldNameMap, validators)
}

func TargetNetworkBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"Network"}, ""))
    fieldNameMap["id"] = "Id"
    fields["accessible"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["accessible"] = "Accessible"
    fields["inaccessible_reasons"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.target_network.inaccessible_reason", reflect.TypeOf(TargetNetwork_InaccessibleReason(TargetNetwork_InaccessibleReason_NO_ASSIGN_RIGHT))), reflect.TypeOf([]TargetNetwork_InaccessibleReason{})))
    fieldNameMap["inaccessible_reasons"] = "InaccessibleReasons"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.target_network",fields, reflect.TypeOf(TargetNetwork{}), fieldNameMap, validators)
}

func NetworkMappingParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["source_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(SourceNetworkBindingType), reflect.TypeOf([]SourceNetwork{})))
    fieldNameMap["source_networks"] = "SourceNetworks"
    fields["target_networks"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(TargetNetworkBindingType), reflect.TypeOf([]TargetNetwork{})))
    fieldNameMap["target_networks"] = "TargetNetworks"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.network_mapping_params",fields, reflect.TypeOf(NetworkMappingParams{}), fieldNameMap, validators)
}

func OvfMessageBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessage_Category(OvfMessage_Category_VALIDATION)))
    fieldNameMap["category"] = "Category"
    fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
    fieldNameMap["issues"] = "Issues"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["value"] = "Value"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("category",
        map[string][]bindings.FieldData {
            "VALIDATION": []bindings.FieldData {
                 bindings.NewFieldData("issues", true),
            },
            "INPUT": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("value", true),
                 bindings.NewFieldData("message", true),
            },
            "SERVER": []bindings.FieldData {
                 bindings.NewFieldData("error", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_message",fields, reflect.TypeOf(OvfMessage{}), fieldNameMap, validators)
}

func ParseIssueBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.parse_issue.category", reflect.TypeOf(ParseIssue_Category(ParseIssue_Category_VALUE_ILLEGAL)))
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.parse_issue",fields, reflect.TypeOf(ParseIssue{}), fieldNameMap, validators)
}

func OvfErrorBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessage_Category(OvfMessage_Category_VALIDATION)))
    fieldNameMap["category"] = "Category"
    fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
    fieldNameMap["issues"] = "Issues"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["value"] = "Value"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("category",
        map[string][]bindings.FieldData {
            "VALIDATION": []bindings.FieldData {
                 bindings.NewFieldData("issues", true),
            },
            "INPUT": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("value", true),
                 bindings.NewFieldData("message", true),
            },
            "SERVER": []bindings.FieldData {
                 bindings.NewFieldData("error", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_error",fields, reflect.TypeOf(OvfError{}), fieldNameMap, validators)
}

func OvfWarningBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_message.category", reflect.TypeOf(OvfMessage_Category(OvfMessage_Category_VALIDATION)))
    fieldNameMap["category"] = "Category"
    fields["issues"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ParseIssueBindingType), reflect.TypeOf([]ParseIssue{})))
    fieldNameMap["issues"] = "Issues"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["value"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["value"] = "Value"
    fields["message"] = bindings.NewOptionalType(bindings.NewReferenceType(std.LocalizableMessageBindingType))
    fieldNameMap["message"] = "Message"
    fields["error"] = bindings.NewOptionalType(bindings.NewDynamicStructType([]bindings.ReferenceType {bindings.NewReferenceType(errors.ErrorBindingType),}, bindings.JSONRPC))
    fieldNameMap["error"] = "Error"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("category",
        map[string][]bindings.FieldData {
            "VALIDATION": []bindings.FieldData {
                 bindings.NewFieldData("issues", true),
            },
            "INPUT": []bindings.FieldData {
                 bindings.NewFieldData("name", true),
                 bindings.NewFieldData("value", true),
                 bindings.NewFieldData("message", true),
            },
            "SERVER": []bindings.FieldData {
                 bindings.NewFieldData("error", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_warning",fields, reflect.TypeOf(OvfWarning{}), fieldNameMap, validators)
}

func OvfInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["messages"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["messages"] = "Messages"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_info",fields, reflect.TypeOf(OvfInfo{}), fieldNameMap, validators)
}

func OvfFileInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["file_type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_file_info.file_type", reflect.TypeOf(OvfFileInfo_FileType(OvfFileInfo_FileType_OVF)))
    fieldNameMap["file_type"] = "FileType"
    fields["optional_upload"] = bindings.NewBooleanType()
    fieldNameMap["optional_upload"] = "OptionalUpload"
    fields["file_url"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["file_url"] = "FileUrl"
    fields["ssl_thumb_print"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["ssl_thumb_print"] = "SslThumbPrint"
    fields["upload_method"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.ovf_file_info.http_upload_method", reflect.TypeOf(OvfFileInfo_HttpUploadMethod(OvfFileInfo_HttpUploadMethod_PUT))))
    fieldNameMap["upload_method"] = "UploadMethod"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    fields["bytes_transferred"] = bindings.NewIntegerType()
    fieldNameMap["bytes_transferred"] = "BytesTransferred"
    fields["sha256"] = bindings.NewStringType()
    fieldNameMap["sha256"] = "Sha256"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_file_info",fields, reflect.TypeOf(OvfFileInfo{}), fieldNameMap, validators)
}

func OvfParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.ovf_params",fields, reflect.TypeOf(OvfParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.property",fields, reflect.TypeOf(Property{}), fieldNameMap, validators)
}

func PropertyParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["properties"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(PropertyBindingType), reflect.TypeOf([]Property{})))
    fieldNameMap["properties"] = "Properties"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.property_params",fields, reflect.TypeOf(PropertyParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.scale_out_group",fields, reflect.TypeOf(ScaleOutGroup{}), fieldNameMap, validators)
}

func ScaleOutParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(ScaleOutGroupBindingType), reflect.TypeOf([]ScaleOutGroup{})))
    fieldNameMap["groups"] = "Groups"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.scale_out_params",fields, reflect.TypeOf(ScaleOutParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.size_params",fields, reflect.TypeOf(SizeParams{}), fieldNameMap, validators)
}

func StorageDiskGroupBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["target_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["target_profile"] = "TargetProfile"
    fields["target_provisioning_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
    fieldNameMap["target_provisioning_type"] = "TargetProvisioningType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.storage_disk_group",fields, reflect.TypeOf(StorageDiskGroup{}), fieldNameMap, validators)
}

func StorageProfileTargetBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["id"] = "Id"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.storage_profile_target",fields, reflect.TypeOf(StorageProfileTarget{}), fieldNameMap, validators)
}

func StorageMappingParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["available_storage_profiles"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageProfileTargetBindingType), reflect.TypeOf([]StorageProfileTarget{})))
    fieldNameMap["available_storage_profiles"] = "AvailableStorageProfiles"
    fields["available_disk_provisioning_types"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))), reflect.TypeOf([]DiskProvisioningType{})))
    fieldNameMap["available_disk_provisioning_types"] = "AvailableDiskProvisioningTypes"
    fields["target_profile"] = bindings.NewOptionalType(bindings.NewIdType([]string {"StorageProfile"}, ""))
    fieldNameMap["target_profile"] = "TargetProfile"
    fields["target_provisioning_type"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.disk_provisioning_type", reflect.TypeOf(DiskProvisioningType(DiskProvisioningType_thin))))
    fieldNameMap["target_provisioning_type"] = "TargetProvisioningType"
    fields["disk_groups"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(StorageDiskGroupBindingType), reflect.TypeOf([]StorageDiskGroup{})))
    fieldNameMap["disk_groups"] = "DiskGroups"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.storage_mapping_params",fields, reflect.TypeOf(StorageMappingParams{}), fieldNameMap, validators)
}

func UnknownSectionBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["tag"] = bindings.NewStringType()
    fieldNameMap["tag"] = "Tag"
    fields["info"] = bindings.NewStringType()
    fieldNameMap["info"] = "Info"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.unknown_section",fields, reflect.TypeOf(UnknownSection{}), fieldNameMap, validators)
}

func UnknownSectionParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["unknown_sections"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(UnknownSectionBindingType), reflect.TypeOf([]UnknownSection{})))
    fieldNameMap["unknown_sections"] = "UnknownSections"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.unknown_section_params",fields, reflect.TypeOf(UnknownSectionParams{}), fieldNameMap, validators)
}

func VServiceProviderBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["key"] = "Key"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["auto_bind"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["auto_bind"] = "AutoBind"
    fields["status"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vcenter.ovf.V_Service_Provider.binding_status", reflect.TypeOf(VServiceProvider_BindingStatus(VServiceProvider_BindingStatus_RED))))
    fieldNameMap["status"] = "Status"
    fields["validation_message"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["validation_message"] = "ValidationMessage"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.V_Service_Provider",fields, reflect.TypeOf(VServiceProvider{}), fieldNameMap, validators)
}

func VServiceDependencyBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["id"] = "Id"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["name"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["name"] = "Name"
    fields["description"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["description"] = "Description"
    fields["required"] = bindings.NewOptionalType(bindings.NewBooleanType())
    fieldNameMap["required"] = "Required"
    fields["selected_provider_key"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["selected_provider_key"] = "SelectedProviderKey"
    fields["available_providers"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VServiceProviderBindingType), reflect.TypeOf([]VServiceProvider{})))
    fieldNameMap["available_providers"] = "AvailableProviders"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.V_Service_Dependency",fields, reflect.TypeOf(VServiceDependency{}), fieldNameMap, validators)
}

func VServiceParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["dependencies"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(VServiceDependencyBindingType), reflect.TypeOf([]VServiceDependency{})))
    fieldNameMap["dependencies"] = "Dependencies"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.V_Service_Params",fields, reflect.TypeOf(VServiceParams{}), fieldNameMap, validators)
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
    return bindings.NewStructType("com.vmware.vcenter.ovf.vcenter_extension_params",fields, reflect.TypeOf(VcenterExtensionParams{}), fieldNameMap, validators)
}

func VcenterGuestCustomizationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["type"] = bindings.NewEnumType("com.vmware.vcenter.ovf.vcenter_guest_customization.type", reflect.TypeOf(VcenterGuestCustomization_Type(VcenterGuestCustomization_Type_SPECIFICATION)))
    fieldNameMap["type"] = "Type_"
    fields["specification_id"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["specification_id"] = "SpecificationId"
    fields["xml"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["xml"] = "Xml"
    var validators = []bindings.Validator{}
    uv1 := bindings.NewUnionValidator("type",
        map[string][]bindings.FieldData {
            "SPECIFICATION": []bindings.FieldData {
                 bindings.NewFieldData("specification_id", true),
            },
            "XML": []bindings.FieldData {
                 bindings.NewFieldData("xml", true),
            },
        },
    )
    validators = append(validators, uv1)
    return bindings.NewStructType("com.vmware.vcenter.ovf.vcenter_guest_customization",fields, reflect.TypeOf(VcenterGuestCustomization{}), fieldNameMap, validators)
}

func VcenterGuestCustomizationParamsBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["virtual_systems"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["virtual_systems"] = "VirtualSystems"
    fields["customizations"] = bindings.NewOptionalType(bindings.NewMapType(bindings.NewStringType(), bindings.NewReferenceType(VcenterGuestCustomizationBindingType),reflect.TypeOf(map[string]VcenterGuestCustomization{})))
    fieldNameMap["customizations"] = "Customizations"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.ovf.vcenter_guest_customization_params",fields, reflect.TypeOf(VcenterGuestCustomizationParams{}), fieldNameMap, validators)
}


