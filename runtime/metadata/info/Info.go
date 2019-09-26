package info

/* **********************************************************
 * Copyright 2019 VMware, Inc.  All rights reserved.
 *      -- VMware Confidential
 * **********************************************************/

type Info interface {
	Identifier() string
	SetFieldInfo(name string, fieldinfo *FieldInfo)
}
