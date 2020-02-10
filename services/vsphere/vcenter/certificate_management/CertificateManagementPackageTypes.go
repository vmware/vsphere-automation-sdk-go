/* Copyright © 2019 VMware, Inc. All Rights Reserved.
   SPDX-License-Identifier: BSD-2-Clause */

// Code generated. DO NOT EDIT.

/*
 * Data type definitions file for package: com.vmware.vcenter.certificate_management.
 * Includes binding types of a top level structures and enumerations.
 * Shared by client-side stubs and server-side skeletons to ensure type
 * compatibility.
 */

package certificate_management

import (
	"reflect"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/bindings"
)


// The ``X509CertChain`` class contains x509 certificate chain. This class was added in vSphere API 6.7.2.
type X509CertChain struct {
    // Certificate chain in base64 format. This property was added in vSphere API 6.7.2.
	CertChain []string
}




func X509CertChainBindingType() bindings.BindingType {
	fields := make(map[string]bindings.BindingType)
	fieldNameMap := make(map[string]string)
	fields["cert_chain"] = bindings.NewListType(bindings.NewStringType(), reflect.TypeOf([]string{}))
	fieldNameMap["cert_chain"] = "CertChain"
	var validators = []bindings.Validator{}
	return bindings.NewStructType("com.vmware.vcenter.certificate_management.x509_cert_chain", fields, reflect.TypeOf(X509CertChain{}), fieldNameMap, validators)
}


