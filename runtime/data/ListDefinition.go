/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/l10n"
	"strconv"
)

/**
 * Type definition for vAPI list built-in type
 */
type ListDefinition struct {
	/**
	 * The Definition of the type that is contained in a list.
	 */
	elementType DataDefinition
}

func NewListDefinition(elementType DataDefinition) ListDefinition {
	return ListDefinition{elementType: elementType}
}

func (listDefinition ListDefinition) Type() DataType {
	return LIST
}

func (listDefinition ListDefinition) ElementType() DataDefinition {
	return listDefinition.elementType
}

func (listDefinition ListDefinition) Validate(value DataValue) []error {
	var result = listDefinition.Type().Validate(value)
	if result != nil {
		return result
	}
	var listValue = (value).(*ListValue)
	elementType := listDefinition.elementType
	for index, listElement := range listValue.List() {
		var subErrors = elementType.Validate(listElement)
		if len(subErrors) != 0 {
			args := map[string]string{
				"value": listElement.Type().String(),
				"index": strconv.Itoa(index)}
			var msg = l10n.NewRuntimeError("vapi.data.list.invalid.entry", args)
			result = append(result, subErrors...)
			result = append(result, msg)
		}
	}
	return result
}

func (listDefinition ListDefinition) CompleteValue(value DataValue) {
	if value != nil && value.Type() == LIST {
		var listValue = (value).(*ListValue)
		for _, listElement := range listValue.List() {
			listDefinition.elementType.CompleteValue(listElement)
		}
	}
}

func (listDefinition ListDefinition) ValidInstanceOf(value DataValue) bool {
	var errors = listDefinition.Validate(value)
	if len(errors) != 0 {
		return false
	}
	return true
}

func (listDefinition ListDefinition) String() string {
	return LIST.String()
}

func (listDefinition ListDefinition) NewValue() *ListValue {
	return NewListValue()
}
