// Copyright (c) 2019-2024 Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: BSD-2-Clause

package core

/*
InterfaceIdentifier has the information required to uniquely
address a vAPI interface
*/
type InterfaceIdentifier struct {
	/**
	 * Name of the interface
	 */
	interfaceName string
}

func NewInterfaceIdentifier(interfaceName string) InterfaceIdentifier {
	return InterfaceIdentifier{interfaceName: interfaceName}
}

func (interfaceIdentifier InterfaceIdentifier) Equals(other InterfaceIdentifier) bool {
	return interfaceIdentifier.interfaceName == other.interfaceName
}

func (interfaceIdentifier InterfaceIdentifier) Name() string {
	return (interfaceIdentifier.interfaceName)
}

func (interfaceIdentifier InterfaceIdentifier) String() string {
	return (interfaceIdentifier.interfaceName)
}
