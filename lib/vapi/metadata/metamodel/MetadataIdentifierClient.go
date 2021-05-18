// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: MetadataIdentifier
// Used by client-side stubs.

package metamodel

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
)

const _ = core.SupportedByRuntimeVersion1

// The ``MetadataIdentifier`` interface provides string constants that can be used as identifiers for the metadata elements.
//
//  Most of the types in com.vmware.vapi.metadata.metamodel package has a metadata field whose type is ``Map<String, ElementMap>``. MetadataIdentifier contains the identifiers used in the keys of the above Map type.
type MetadataIdentifierClient interface {
}

type metadataIdentifierClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewMetadataIdentifierClient(connector client.Connector) *metadataIdentifierClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vapi.metadata.metamodel.metadata_identifier")
	methodIdentifiers := map[string]core.MethodIdentifier{}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	mIface := metadataIdentifierClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &mIface
}

func (mIface *metadataIdentifierClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := mIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}
