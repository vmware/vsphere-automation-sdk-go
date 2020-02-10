/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.appliance.recovery.backup.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package backup

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
	"net/url"
)


// The ``LocationSpec`` class has fields to represent a location on the backup server. This class was added in vSphere API 6.7.
type LocationSpec struct {
    // Backup location URL. This property was added in vSphere API 6.7.
	Location url.URL
    // Username for the given location. This property was added in vSphere API 6.7.
	LocationUser *string
    // Password for the given location. This property was added in vSphere API 6.7.
	LocationPassword *string
}




func LocationSpecBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["location"] = bindings.NewUriType()
	fieldNameMap["location"] = "Location"
	fields["location_user"] = bindings.NewOptionalType(bindings.NewStringType())
	fieldNameMap["location_user"] = "LocationUser"
	fields["location_password"] = bindings.NewOptionalType(bindings.NewSecretType())
	fieldNameMap["location_password"] = "LocationPassword"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.appliance.recovery.backup.location_spec", fields, reflect.TypeOf(LocationSpec{}), fieldNameMap, validators)
}


