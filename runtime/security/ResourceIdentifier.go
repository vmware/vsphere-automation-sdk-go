package security

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type ResourceIdentifier struct {
	isOperation  bool
	id           string
	resourceType string
}

func NewResourceIdentifier(isOperation bool, id string, resourceType string) ResourceIdentifier {
	return ResourceIdentifier{isOperation: isOperation, id: id, resourceType: resourceType}
}
