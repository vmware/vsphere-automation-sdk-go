/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: CompatibilityReport.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package compatibilityReport

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/esx/hcl"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)




// This ``BiosConstraint`` class contains properties that describe the BIOS that is supported for the given server and ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type BiosConstraint struct {
    Bios hcl.Firmware
    Notes []std.LocalizableMessage
}






// This ``ServerCompatibility`` class contains properties that provide the compatibility information for a server model, cpu and BIOS. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ServerCompatibility struct {
    BiosConstraints []BiosConstraint
    CpuSeries string
    SupportedReleases []string
    VcgLink url.URL
    Notes []std.LocalizableMessage
}






// This ``ServerHclInfo`` class contains properties that describe the server of a ESXi host and its compatibility information. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type ServerHclInfo struct {
    Server hcl.Server
    Matches []ServerCompatibility
    ModelCompatibility hcl.CompatibilityStatus
}






// This ``DeviceConstaint`` class contains properties that describe pair of driver and firmware that are supported for a given PCI device and ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DeviceConstaint struct {
    Driver hcl.Driver
    Firmware *hcl.Firmware
    Notes []std.LocalizableMessage
}






// This ``DeviceHclInfo`` class contains properties that describe a PCI device of a given ESXi host and its compatibility information. 
//
//  If there are multiple PCI devices of the same type on the host each one will be listed in separate instance of this class.. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type DeviceHclInfo struct {
    Compatibility hcl.CompatibilityStatus
    Device hcl.PCIDevice
    DeviceConstraints []DeviceConstaint
    SupportedReleases []string
    VcgLink *url.URL
    Notes []std.LocalizableMessage
}






// This ``HclReport`` represents the hardware compatibility report generated for a specific host and target ESXi release. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type HclReport struct {
    Host url.URL
    TargetRelease string
    ServerHcl ServerHclInfo
    DevicesHcl []DeviceHclInfo
    GeneratedAt time.Time
    Notifications hcl.Notifications
}






// The ``Result`` class contains the result of hardware compatibility report creation operation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Result struct {
    Report HclReport
    Identifier *string
}






// The ``Spec`` class contains properties to describe the input configuration for an ESXi's compatibility report generation. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Spec struct {
    Release string
}









func createInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    fields["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(SpecBindingType))
    fieldNameMap["host"] = "Host"
    fieldNameMap["spec"] = "Spec"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func createOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ResultBindingType)
}

func createRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["spec"] = bindings.NewOptionalType(bindings.NewReferenceType(SpecBindingType))
    paramsTypeMap["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    paramsTypeMap["host"] = bindings.NewIdType([]string {"HostSystem"}, "")
    pathParams["host"] = "host"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "spec",
      "POST",
      "/esx/hcl/hosts/{host}/compatibility-report",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthenticated": 401,"Unsupported": 400,"ResourceInaccessible": 500,"NotAllowedInCurrentState": 400,"Error": 500})
}



func BiosConstraintBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bios"] = bindings.NewReferenceType(hcl.FirmwareBindingType)
    fieldNameMap["bios"] = "Bios"
    fields["notes"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["notes"] = "Notes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.bios_constraint",fields, reflect.TypeOf(BiosConstraint{}), fieldNameMap, validators)
}

func ServerCompatibilityBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["bios_constraints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(BiosConstraintBindingType), reflect.TypeOf([]BiosConstraint{})))
    fieldNameMap["bios_constraints"] = "BiosConstraints"
    fields["cpu_series"] = bindings.NewStringType()
    fieldNameMap["cpu_series"] = "CpuSeries"
    fields["supported_releases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["supported_releases"] = "SupportedReleases"
    fields["vcg_link"] = bindings.NewUriType()
    fieldNameMap["vcg_link"] = "VcgLink"
    fields["notes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["notes"] = "Notes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.server_compatibility",fields, reflect.TypeOf(ServerCompatibility{}), fieldNameMap, validators)
}

func ServerHclInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["server"] = bindings.NewReferenceType(hcl.ServerBindingType)
    fieldNameMap["server"] = "Server"
    fields["matches"] = bindings.NewListType(bindings.NewReferenceType(ServerCompatibilityBindingType), reflect.TypeOf([]ServerCompatibility{}))
    fieldNameMap["matches"] = "Matches"
    fields["model_compatibility"] = bindings.NewEnumType("com.vmware.esx.hcl.compatibility_status", reflect.TypeOf(hcl.CompatibilityStatus(hcl.CompatibilityStatus_COMPATIBLE)))
    fieldNameMap["model_compatibility"] = "ModelCompatibility"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.server_hcl_info",fields, reflect.TypeOf(ServerHclInfo{}), fieldNameMap, validators)
}

func DeviceConstaintBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["driver"] = bindings.NewReferenceType(hcl.DriverBindingType)
    fieldNameMap["driver"] = "Driver"
    fields["firmware"] = bindings.NewOptionalType(bindings.NewReferenceType(hcl.FirmwareBindingType))
    fieldNameMap["firmware"] = "Firmware"
    fields["notes"] = bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{}))
    fieldNameMap["notes"] = "Notes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.device_constaint",fields, reflect.TypeOf(DeviceConstaint{}), fieldNameMap, validators)
}

func DeviceHclInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["compatibility"] = bindings.NewEnumType("com.vmware.esx.hcl.compatibility_status", reflect.TypeOf(hcl.CompatibilityStatus(hcl.CompatibilityStatus_COMPATIBLE)))
    fieldNameMap["compatibility"] = "Compatibility"
    fields["device"] = bindings.NewReferenceType(hcl.PCIDeviceBindingType)
    fieldNameMap["device"] = "Device"
    fields["device_constraints"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DeviceConstaintBindingType), reflect.TypeOf([]DeviceConstaint{})))
    fieldNameMap["device_constraints"] = "DeviceConstraints"
    fields["supported_releases"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{})))
    fieldNameMap["supported_releases"] = "SupportedReleases"
    fields["vcg_link"] = bindings.NewOptionalType(bindings.NewUriType())
    fieldNameMap["vcg_link"] = "VcgLink"
    fields["notes"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(std.LocalizableMessageBindingType), reflect.TypeOf([]std.LocalizableMessage{})))
    fieldNameMap["notes"] = "Notes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.device_hcl_info",fields, reflect.TypeOf(DeviceHclInfo{}), fieldNameMap, validators)
}

func HclReportBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["host"] = bindings.NewUriType()
    fieldNameMap["host"] = "Host"
    fields["target_release"] = bindings.NewStringType()
    fieldNameMap["target_release"] = "TargetRelease"
    fields["server_hcl"] = bindings.NewReferenceType(ServerHclInfoBindingType)
    fieldNameMap["server_hcl"] = "ServerHcl"
    fields["devices_hcl"] = bindings.NewOptionalType(bindings.NewListType(bindings.NewReferenceType(DeviceHclInfoBindingType), reflect.TypeOf([]DeviceHclInfo{})))
    fieldNameMap["devices_hcl"] = "DevicesHcl"
    fields["generated_at"] = bindings.NewDateTimeType()
    fieldNameMap["generated_at"] = "GeneratedAt"
    fields["notifications"] = bindings.NewReferenceType(hcl.NotificationsBindingType)
    fieldNameMap["notifications"] = "Notifications"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.hcl_report",fields, reflect.TypeOf(HclReport{}), fieldNameMap, validators)
}

func ResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewReferenceType(HclReportBindingType)
    fieldNameMap["report"] = "Report"
    fields["identifier"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.esx.hcl.resources.CompatibilityReport"}, ""))
    fieldNameMap["identifier"] = "Identifier"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.result",fields, reflect.TypeOf(Result{}), fieldNameMap, validators)
}

func SpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["release"] = bindings.NewStringType()
    fieldNameMap["release"] = "Release"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.hosts.compatibility_report.spec",fields, reflect.TypeOf(Spec{}), fieldNameMap, validators)
}


