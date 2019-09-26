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




// The ``Token`` class contains information about the token required in the HTTP GET request to generate the report.
 type Token struct {
    Token string
    Expiry time.Time
}






// The ``Location`` class contains the URI location to download the report from, as well as a token required (as a header on the HTTP request) to get the bundle. The validity of the token is 5 minutes as best attempt. After the token expires, any attempt to call the URI with said token will fail.
 type Location struct {
    Uri url.URL
    DownloadFileToken Token
}









func getInputType() bindings.StructType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
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
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
    paramsTypeMap["report"] = bindings.NewIdType([]string {"com.vmware.vcenter.lcm.report"}, "")
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
      "/vcenter/lcm/reports/{report}",
       resultHeaders,
       0,
       errorHeaders,
       map[string]int{"Unauthenticated": 401,"NotFound": 404,"Error": 500})
}



func TokenBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["token"] = bindings.NewStringType()
    fieldNameMap["token"] = "Token"
    fields["expiry"] = bindings.NewDateTimeType()
    fieldNameMap["expiry"] = "Expiry"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reports.token",fields, reflect.TypeOf(Token{}), fieldNameMap, validators)
}

func LocationBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["uri"] = bindings.NewUriType()
    fieldNameMap["uri"] = "Uri"
    fields["download_file_token"] = bindings.NewReferenceType(TokenBindingType)
    fieldNameMap["download_file_token"] = "DownloadFileToken"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.lcm.reports.location",fields, reflect.TypeOf(Location{}), fieldNameMap, validators)
}


