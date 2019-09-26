/* **********************************************************
 * Copyright 2018 VMware, Inc.  All rights reserved. -- VMware Confidential
 * **********************************************************/
package data

type BlobValue struct {
	value []byte
}

func NewBlobValue(value []byte) *BlobValue {
	return &BlobValue{value: value}
}

func (b *BlobValue) Type() DataType {
	return BLOB
}

func (b *BlobValue) Value() []byte {
	return (b.value)
}
