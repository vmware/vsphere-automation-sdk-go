// Copyright © 2019-2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: BSD-2-Clause

// Auto generated code. DO NOT EDIT.

// Interface file for service: ReplicaDiskCollections
// Used by client-side stubs.

package draas

import (
	"github.com/vmware/vsphere-automation-sdk-go/lib/vapi/std/errors"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/bindings"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/core"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/lib"
	"github.com/vmware/vsphere-automation-sdk-go/runtime/protocol/client"
	"github.com/vmware/vsphere-automation-sdk-go/services/vmc/draas/model"
)

const _ = core.SupportedByRuntimeVersion1

type ReplicaDiskCollectionsClient interface {

	// Query replica disk collections
	//
	// @param orgParam Organization identifier (required)
	// @param sddcParam sddc identifier (required)
	// @param datastoreMoIdParam Represents the datastore moref id to search. (optional)
	// @throws Unauthenticated  Unauthorized
	// @throws Unauthorized  Forbidden
	// @throws NotFound  Not found
	Get(orgParam string, sddcParam string, datastoreMoIdParam *string) ([]model.ReplicaDiskCollection, error)
}

type replicaDiskCollectionsClient struct {
	connector           client.Connector
	interfaceDefinition core.InterfaceDefinition
	errorsBindingMap    map[string]bindings.BindingType
}

func NewReplicaDiskCollectionsClient(connector client.Connector) *replicaDiskCollectionsClient {
	interfaceIdentifier := core.NewInterfaceIdentifier("com.vmware.vmc.draas.replica_disk_collections")
	methodIdentifiers := map[string]core.MethodIdentifier{
		"get": core.NewMethodIdentifier(interfaceIdentifier, "get"),
	}
	interfaceDefinition := core.NewInterfaceDefinition(interfaceIdentifier, methodIdentifiers)
	errorsBindingMap := make(map[string]bindings.BindingType)

	rIface := replicaDiskCollectionsClient{interfaceDefinition: interfaceDefinition, errorsBindingMap: errorsBindingMap, connector: connector}
	return &rIface
}

func (rIface *replicaDiskCollectionsClient) GetErrorBindingType(errorName string) bindings.BindingType {
	if entry, ok := rIface.errorsBindingMap[errorName]; ok {
		return entry
	}
	return errors.ERROR_BINDINGS_MAP[errorName]
}

func (rIface *replicaDiskCollectionsClient) Get(orgParam string, sddcParam string, datastoreMoIdParam *string) ([]model.ReplicaDiskCollection, error) {
	typeConverter := rIface.connector.TypeConverter()
	executionContext := rIface.connector.NewExecutionContext()
	sv := bindings.NewStructValueBuilder(replicaDiskCollectionsGetInputType(), typeConverter)
	sv.AddStructField("Org", orgParam)
	sv.AddStructField("Sddc", sddcParam)
	sv.AddStructField("DatastoreMoId", datastoreMoIdParam)
	inputDataValue, inputError := sv.GetStructValue()
	if inputError != nil {
		var emptyOutput []model.ReplicaDiskCollection
		return emptyOutput, bindings.VAPIerrorsToError(inputError)
	}
	operationRestMetaData := replicaDiskCollectionsGetRestMetadata()
	connectionMetadata := map[string]interface{}{lib.REST_METADATA: operationRestMetaData}
	connectionMetadata["isStreamingResponse"] = false
	rIface.connector.SetConnectionMetadata(connectionMetadata)
	methodResult := rIface.connector.GetApiProvider().Invoke("com.vmware.vmc.draas.replica_disk_collections", "get", inputDataValue, executionContext)
	var emptyOutput []model.ReplicaDiskCollection
	if methodResult.IsSuccess() {
		output, errorInOutput := typeConverter.ConvertToGolang(methodResult.Output(), replicaDiskCollectionsGetOutputType())
		if errorInOutput != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInOutput)
		}
		return output.([]model.ReplicaDiskCollection), nil
	} else {
		methodError, errorInError := typeConverter.ConvertToGolang(methodResult.Error(), rIface.GetErrorBindingType(methodResult.Error().Name()))
		if errorInError != nil {
			return emptyOutput, bindings.VAPIerrorsToError(errorInError)
		}
		return emptyOutput, methodError.(error)
	}
}
