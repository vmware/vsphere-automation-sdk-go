package core

type SecurityContext interface {

	Property(key string) interface{}
	GetAllProperties() map[string]interface{}
	SetProperty(key string, value interface{})

}
