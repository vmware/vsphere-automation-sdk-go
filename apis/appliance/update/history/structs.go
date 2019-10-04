/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: History.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package history

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/apis/appliance/update"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)



// The ``Severity`` enumeration class defines the severity of the issues fixed in the update. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Severity string

const (
    // Vulnerabilities that can be exploited by an unauthenticated attacker from the Internet or those that break the guest/host Operating System isolation. The exploitation results in the complete compromise of confidentiality, integrity, and availability of user data and/or processing resources without user interaction. Exploitation could be leveraged to propagate an Internet worm or execute arbitrary code between Virtual Machines and/or the Host Operating System. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Severity_CRITICAL Severity = "CRITICAL"
    // Vulnerabilities that are not rated critical but whose exploitation results in the complete compromise of confidentiality and/or integrity of user data and/or processing resources through user assistance or by authenticated attackers. This rating also applies to those vulnerabilities which could lead to the complete compromise of availability when exploitation is by a remote unauthenticated attacker from the Internet or through a breach of virtual machine isolation. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Severity_IMPORTANT Severity = "IMPORTANT"
    // Vulnerabilities where the ability to exploit is mitigated to a significant degree by configuration or difficulty of exploitation, but in certain deployment scenarios could still lead to the compromise of confidentiality, integrity, or availability of user data and/or processing resources. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Severity_MODERATE Severity = "MODERATE"
)

func (s Severity) Severity() bool {
    switch s {
        case Severity_CRITICAL:
            return true
        case Severity_IMPORTANT:
            return true
        case Severity_MODERATE:
            return true
        default:
            return false
    }
}




// The ``Category`` enumeration class defines update type. **Warning:** This enumeration is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.
 
type Category string

const (
    // Fixes vulnerabilities, doesn't change functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Category_SECURITY Category = "SECURITY"
    // Fixes bugs/vulnerabilities, doesn't change functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Category_FIX Category = "FIX"
    // Changes product functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
     Category_UPDATE Category = "UPDATE"
    // Introduces new features, significantly changes product functionality. **Warning:** This constant field is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
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





// The ``Summary`` class contains information about the installed updates. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Summary struct {
    Name string
    InstallDate time.Time
    Version string
    Severity Severity
    ReleaseDate time.Time
    UpdateType Category
}






// The ``Info`` class contains detailed information about the installed updates. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Info struct {
    Name string
    InstallDate time.Time
    Product string
    Vendor string
    KnowledgeBase url.URL
    UpdateRepo url.URL
    Build int64
    EulaAcceptTime time.Time
    Tags *string
    Description std.LocalizableMessage
    Priority update.CommonInfo_Priority
    Severity update.CommonInfo_Severity
    UpdateType update.CommonInfo_Category
    ReleaseDate time.Time
    RebootRequired bool
    Size int64
}









func listInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func listOutputType() bindings.BindingType {
    return bindings.NewListType(bindings.NewReferenceType(SummaryBindingType), reflect.TypeOf([]Summary{}))
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
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}


func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.history"}, "")
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
    resultHeaders := map[string]string{}
    errorHeaders := map[string]string{}
    errorHeaders["Unauthenticated.challenge"] = "WWW-Authenticate"
    return protocol.NewOperationRestMetadata(
      paramsTypeMap,
      pathParams,
      queryParams,
      headerParams,
      "",
      "null",
      "",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Error": 500,"Unauthenticated": 401,"Unauthorized": 403,"NotAllowedInCurrentState": 400})
}



func SummaryBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["install_date"] = bindings.NewDateTimeType()
    fieldNameMap["install_date"] = "InstallDate"
    fields["version"] = bindings.NewIdType([]string {"com.vmware.appliance.update.history"}, "")
    fieldNameMap["version"] = "Version"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.history.severity", reflect.TypeOf(Severity(Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.history.category", reflect.TypeOf(Category(Category_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.history.summary",fields, reflect.TypeOf(Summary{}), fieldNameMap, validators)
}

func InfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["name"] = bindings.NewStringType()
    fieldNameMap["name"] = "Name"
    fields["install_date"] = bindings.NewDateTimeType()
    fieldNameMap["install_date"] = "InstallDate"
    fields["product"] = bindings.NewStringType()
    fieldNameMap["product"] = "Product"
    fields["vendor"] = bindings.NewStringType()
    fieldNameMap["vendor"] = "Vendor"
    fields["knowledge_base"] = bindings.NewUriType()
    fieldNameMap["knowledge_base"] = "KnowledgeBase"
    fields["update_repo"] = bindings.NewUriType()
    fieldNameMap["update_repo"] = "UpdateRepo"
    fields["build"] = bindings.NewIntegerType()
    fieldNameMap["build"] = "Build"
    fields["eula_accept_time"] = bindings.NewDateTimeType()
    fieldNameMap["eula_accept_time"] = "EulaAcceptTime"
    fields["tags"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["tags"] = "Tags"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    fields["priority"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.priority", reflect.TypeOf(update.CommonInfo_Priority(update.CommonInfo_Priority_HIGH)))
    fieldNameMap["priority"] = "Priority"
    fields["severity"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.severity", reflect.TypeOf(update.CommonInfo_Severity(update.CommonInfo_Severity_CRITICAL)))
    fieldNameMap["severity"] = "Severity"
    fields["update_type"] = bindings.NewEnumType("com.vmware.appliance.update.common_info.category", reflect.TypeOf(update.CommonInfo_Category(update.CommonInfo_Category_SECURITY)))
    fieldNameMap["update_type"] = "UpdateType"
    fields["release_date"] = bindings.NewDateTimeType()
    fieldNameMap["release_date"] = "ReleaseDate"
    fields["reboot_required"] = bindings.NewBooleanType()
    fieldNameMap["reboot_required"] = "RebootRequired"
    fields["size"] = bindings.NewIntegerType()
    fieldNameMap["size"] = "Size"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.appliance.update.history.info",fields, reflect.TypeOf(Info{}), fieldNameMap, validators)
}


