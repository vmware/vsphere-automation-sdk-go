/*
 * Copyright 2019 VMware, Inc.  All rights reserved.
 */

/*
 * AUTO GENERATED FILE -- DO NOT MODIFY!
 *
 * Data type definitions file for package: com.vmware.vcenter.deployment.migrate.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package migrate

import (
    "reflect"
    "gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)



// The ``ActiveDirectorySpec`` class contains information used to join the migrated vCenter Server appliance to the Active Directory. **Warning:** This class is part of a new feature in development. It may be changed at any time and may not have all supported functionality implemented.
type ActiveDirectorySpec struct {
    Domain string
    Username string
    Password string
}










func ActiveDirectorySpecBindingType() bindings.BindingType {
    fields := make(map[string]bindings.BindingType)
    fieldNameMap := make(map[string]string)
    fields["domain"] = bindings.NewStringType()
    fieldNameMap["domain"] = "Domain"
    fields["username"] = bindings.NewStringType()
    fieldNameMap["username"] = "Username"
    fields["password"] = bindings.NewStringType()
    fieldNameMap["password"] = "Password"
    var validators = []bindings.Validator{}
    return bindings.NewStructType("com.vmware.vcenter.deployment.migrate.active_directory_spec",fields, reflect.TypeOf(ActiveDirectorySpec{}), fieldNameMap, validators)
}


