/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vstats.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package vstats

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/lib/vapi/std"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``QueryPredicate`` enum describes how to query an id present in a ``RsrcId``.
//
// <p> See {@link com.vmware.vapi.bindings.ApiEnumeration enumerated types description}.

type QueryPredicate string

const (
    // Matching id-s by equality.
     QueryPredicate_EQUAL QueryPredicate = "EQUAL"
    // Matching all available id-s.
     QueryPredicate_ALL QueryPredicate = "ALL"
)

func (q QueryPredicate) QueryPredicate() bool {
    switch q {
        case QueryPredicate_EQUAL:
            return true
        case QueryPredicate_ALL:
            return true
        default:
            return false
    }
}





// The ``CidMid`` class is used to designate a counter. It contains a counter id that identifies the semantical counter. There is optional metadata identifier that identifies the particular generation of the counter. When counter metadata is not designated vStats will use a default for the counter metadata.
type CidMid struct {
    Cid string
    Mid *string
}






// The ``RsrcId`` class specifies identification of a resource to be monitored by an acquisition specification record.
type RsrcId struct {
    Key *string
    Type_ *string
    IdValue string
    Predicate *QueryPredicate
    Scheme *string
}






// The ``UserInfo`` class contains human legible, localizable description, used for VMware provided objects.
type UserInfo struct {
    Label std.LocalizableMessage
    Description std.LocalizableMessage
}










func CidMidBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["cid"] = bindings.NewIdType([]string {"com.vmware.vstats.model.Counter"}, "")
    fieldNameMap["cid"] = "Cid"
    fields["mid"] = bindings.NewOptionalType(bindings.NewIdType([]string {"com.vmware.vstats.model.CounterMetadata"}, ""))
    fieldNameMap["mid"] = "Mid"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.cid_mid",fields, reflect.TypeOf(CidMid{}), fieldNameMap, validators)
}

func RsrcIdBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["key"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["key"] = "Key"
    fields["type"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["type"] = "Type_"
    fields["id_value"] = bindings.NewStringType()
    fieldNameMap["id_value"] = "IdValue"
    fields["predicate"] = bindings.NewOptionalType(bindings.NewEnumType("com.vmware.vstats.query_predicate", reflect.TypeOf(QueryPredicate(QueryPredicate_EQUAL))))
    fieldNameMap["predicate"] = "Predicate"
    fields["scheme"] = bindings.NewOptionalType(bindings.NewStringType())
    fieldNameMap["scheme"] = "Scheme"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.rsrc_id",fields, reflect.TypeOf(RsrcId{}), fieldNameMap, validators)
}

func UserInfoBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["label"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["label"] = "Label"
    fields["description"] = bindings.NewReferenceType(std.LocalizableMessageBindingType)
    fieldNameMap["description"] = "Description"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vstats.user_info",fields, reflect.TypeOf(UserInfo{}), fieldNameMap, validators)
}


