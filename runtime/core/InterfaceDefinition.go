// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

type InterfaceDefinition struct {
	id        InterfaceIdentifier
	methodIds map[string]MethodIdentifier
}

func NewInterfaceDefinition(id InterfaceIdentifier, methodIds map[string]MethodIdentifier) InterfaceDefinition {
	return InterfaceDefinition{id: id, methodIds: methodIds}
}

func (interfaceDefinition InterfaceDefinition) Identifier() InterfaceIdentifier {
	return interfaceDefinition.id
}

func (interfaceDefinition InterfaceDefinition) MethodIdentifiers() map[string]MethodIdentifier {
	return interfaceDefinition.methodIds
}

func (interfaceDefinition InterfaceDefinition) Equals(other InterfaceDefinition) bool {
	return interfaceDefinition.id.Equals(other.id)
	//TODO: sreeshas
	// write code to compare slice of MethodIdentifiers
}
