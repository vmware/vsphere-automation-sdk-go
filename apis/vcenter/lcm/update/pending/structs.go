/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Pending.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package pending

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)

// Resource type for pending update
const Pending_RESOURCE_TYPE = "com.vmware.vcenter.lcm.update.pending"


// Level of severity for applying a given patch or update.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type SeverityType string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation.
     SeverityType_CRITICAL SeverityType = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers.
     SeverityType_IMPORTANT SeverityType = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources.
     SeverityType_MODERATE SeverityType = "MODERATE"
    // All other issues that may or maynot have a security impact. Vulnerabilities where exploitation is believed to be extremely difficult, or where successful exploitation would have minimal impact.
     SeverityType_LOW SeverityType = "LOW"
)

func (s SeverityType) SeverityType() bool {
    switch s {
        case SeverityType_CRITICAL:
            return true
        case SeverityType_IMPORTANT:
            return true
        case SeverityType_MODERATE:
            return true
        case SeverityType_LOW:
            return true
        default:
            return false
    }
}




// The ``Category`` enumeration class defines the type of payload this release has on top of previous release
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Category string

const (
    // Fixes vulnerabilities, doesn't change functionality
     Category_SECURITY Category = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality
     Category_FIX Category = "FIX"
    // Changes product functionality
     Category_UPDATE Category = "UPDATE"
    // Introduces new features, significantly changes product functionality
     Category_UPGRADE Category = "UPGRADE"
)

func (c Category) Category() bool {
    switch c {
        case Category_SECURITY:
            return true
        case Category_FIX:
            return true
        case Category_UPDATE:
            return true
        case Category_UPGRADE:
            return true
        default:
            return false
    }
}




// The ``UpdateType`` enumeration class defines update type
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type UpdateType string

const (
    // Fixes bugs/vulnerabilities, doesn't change functionality
     UpdateType_PATCH UpdateType = "PATCH"
    // Changes product functionality
     UpdateType_UPDATE UpdateType = "UPDATE"
    // Introduces new features, significantly changes product functionality
     UpdateType_UPGRADE UpdateType = "UPGRADE"
)

func (u UpdateType) UpdateType() bool {
    switch u {
        case UpdateType_PATCH:
            return true
        case UpdateType_UPDATE:
            return true
        case UpdateType_UPGRADE:
            return true
        default:
            return false
    }
}





// The ``Summary`` class contains basic information about the vCenter patch/update/upgrade
 type Summary struct {
    PendingUpdate string
    Version string
    ReleaseDate time.Time
    Severity SeverityType
    Build string
    UpdateType UpdateType
    Category Category
    RebootRequired bool
    ExecuteURL url.URL
    ReleaseNotes []url.URL
}






// The ``ListResult`` class contains information about the pending patch/updates for the given vCenter server.
 type ListResult struct {
    LastCheckTime time.Time
    UpdateCount *int64
    UpgradeCount *int64
    Updates []Summary
}






// The ``Info`` class contains detailed information about the vCenter patch/update.
 type Info struct {
    Description string
    PendingUpdate string
    Version string
    ReleaseDate time.Time
    Severity SeverityType
    Build string
    UpdateType UpdateType
    Category Category
    RebootRequired bool
    ExecuteURL url.URL
    ReleaseNotes []url.URL
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewReferenceType(ListResultBindingType)
}

func listRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/lcm/update/pending",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"Error": 500})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["version"] = "Version"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(InfoBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    paramsTypeMap["version"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    pathParams["version"] = "version"
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "GET",
      "/vcenter/lcm/update/pending/{version}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["pending_update"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["pending_update"] = "PendingUpdate"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(SeverityType(SeverityType_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["build"] = bindings.NewStringType()
    fieldNameMap["build"] = "Build"
    fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(UpdateType(UpdateType_PATCH)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(Category(Category_SECURITY)))
    fieldNameMap["category"] = "Category"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["execute_URL"] = bindings.NewUriType()
    fieldNameMap["execute_URL"] = "ExecuteURL"
    fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func ListResultBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["last_check_time"] = bindings.NewDateTimeType()
    fieldNameMap["last_check_time"] = "LastCheckTime"
    fields["update_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["update_count"] = "UpdateCount"
    fields["upgrade_count"] = bindings.NewOptionalType(bindings.NewIntegerType())
    fieldNameMap["upgrade_count"] = "UpgradeCount"
    fields["updates"] = bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
    fieldNameMap["updates"] = "Updates"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.list_result",fields, reflect.TypeOf(ListResult{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["description"] = bindings.NewStringType()
    fieldNameMap["description"] = "Description"
    fields["pending_update"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.update.pending"}, "")
    fieldNameMap["pending_update"] = "PendingUpdate"
    fields["version"] = bindings.NewStringType()
    fieldNameMap["version"] = "Version"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["severity"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.severity_type", reflect.TypeOf(SeverityType(SeverityType_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["build"] = bindings.NewStringType()
    fieldNameMap["build"] = "Build"
    fields["update_type"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.update_type", reflect.TypeOf(UpdateType(UpdateType_PATCH)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["category"] = bindings.NewEnumType("com.vmware.vcenter.lcm.update.pending.category", reflect.TypeOf(Category(Category_SECURITY)))
    fieldNameMap["category"] = "Category"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["execute_URL"] = bindings.NewUriType()
    fieldNameMap["execute_URL"] = "ExecuteURL"
    fields["release_notes"] = bindings.NewListType(bindings.NewUriType(), reflect.TypeOf([]url.URL{}))
    fieldNameMap["release_notes"] = "ReleaseNotes"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.update.pending.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


