package bindings

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
)

type StructValueExtractor struct {
	structType    StructType
	structValue   *data.StructValue
	typeConverter *TypeConverter
}

func NewStructValueExtractor(structType StructType, structValue *data.StructValue, typeConverter *TypeConverter) *StructValueExtractor {
	return &StructValueExtractor{structType: structType, structValue: structValue, typeConverter: typeConverter}
}

func (s *StructValueExtractor) ExtractValue(fieldName string) (interface{}, []error) {
	fieldValue, _ := s.structValue.Field(fieldName)
	fieldType := s.structType.Field(fieldName)
	return s.typeConverter.ConvertToGolang(fieldValue, fieldType)
}
