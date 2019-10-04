/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for service: Reports.
 * Includes binding types of a structures and enumerations defined in the service.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package reports

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/protocol"
    "net/url"
    "time"
)




// The ``Token`` class contains information about the token required to be passed in the HTTP header in the HTTP GET request to generate the report. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Token struct {
    Token string
    Expiry time.Time
}






// The ``Location`` class contains the URI location to download generated compatibility report, as well as a token required (as a header on the HTTP GET request) to get the report. The validity of the token is 5 minutes. After the token expires, any attempt to call the URI with said token will fail. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
 type Location struct {
    Url url.URL
    ReportToken Token
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewIdType([]string {"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
    fieldNameMap["report"] = "Report"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("operation-input", fields, reflect.TypeOf(data.StructValue{}), fieldNameMap, validators)
}

func getOutputType() bindings.BindingType {
    return bindings.NewReferenceType(LocationBindingType)
}

func getRestMetadata() protocol.OperationRestMetadata {
    paramsTypeMap := map[string]bindings.BindingType{}
    pathParams := map[string]string{}
    queryParams := map[string]string{}
    headerParams := map[string]string{}
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.esx.hcl.resources.CompatibilityReport"}, "")
    pathParams["report"] = "report"
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
      "/esx/hcl/reports/{report}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"NotFound": 404,"Unauthenticated": 401,"ResourceInaccessible": 500,"Error": 500})
}



func TokenBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["token"] = bindings.NewSecretType()
    fieldNameMap["token"] = "Token"
    fields["expiry"] = bindings.NewDateTimeType()
    fieldNameMap["expiry"] = "Expiry"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.reports.token",fields, reflect.TypeOf(Token{}), fieldNameMap, validators)
}

func LocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["url"] = bindings.NewUriType()
    fieldNameMap["url"] = "Url"
    fields["report_token"] = bindings.NewReferenceType(TokenBindingType)
    fieldNameMap["report_token"] = "ReportToken"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.esx.hcl.reports.location",fields, reflect.TypeOf(Location{}), fieldNameMap, validators)
}


