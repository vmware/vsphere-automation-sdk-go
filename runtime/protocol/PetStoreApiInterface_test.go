package protocol_test

import (
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/core"
	"gitlab.eng.vmware.com/golangsdk/vsphere-automation-sdk-go/runtime/data"
	"reflect"
)

/**
 * Pick a name for the interface. I pick "User"

 * Identify operations/methods for the interface. I choose "create", "get", "update", "delete"

 * Define each method with the help of MethodDefinition

 * Implement all methods of the ApiInterface

 * Initialize this interface

 */

type PetStoreApiInterface struct {
	interfaceName       string
	interfaceDefinition core.InterfaceDefinition
	methodNametoDefMap  map[string]*core.MethodDefinition
	methodMap           map[string]func(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult
	//ApplicationData
	userMap map[string]*User
}

type User struct {
	id         interface{}
	userName   string
	firstName  string
	lastName   string
	email      string
	password   string
	phone      string
	userStatus int64
}

func (p *PetStoreApiInterface) Identifier() core.InterfaceIdentifier {
	return core.NewInterfaceIdentifier(p.interfaceName)
}

func (p *PetStoreApiInterface) Definition() core.InterfaceDefinition {
	return p.interfaceDefinition
}

func (p *PetStoreApiInterface) MethodDefinition(mi core.MethodIdentifier) *core.MethodDefinition {
	if val, ok := p.methodNametoDefMap[mi.Name()]; ok {
		return val
	}
	return nil
}

func (p *PetStoreApiInterface) Invoke(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var method = p.methodMap[methodId.Name()]
	return method(ctx, methodId, input)
}

//Define methods with MethodDefinitions

func (p *PetStoreApiInterface) getUserDefinition() data.StructDefinition {
	var structFields = make(map[string]data.DataDefinition)
	structFields["id"] = data.OpaqueDefinition{}
	structFields["userName"] = data.NewStringDefinition()
	structFields["firstName"] = data.NewStringDefinition()
	structFields["lastName"] = data.NewStringDefinition()
	structFields["email"] = data.NewStringDefinition()
	structFields["password"] = data.NewStringDefinition()
	structFields["phone"] = data.NewStringDefinition()
	structFields["userStatus"] = data.NewIntegerDefinition()
	var input = data.NewStructDefinition("User", structFields)
	return input
}

func (p *PetStoreApiInterface) createUserMethodDefinition() *core.MethodDefinition {
	var interfaceIdentifier = core.NewInterfaceIdentifier(p.interfaceName)
	//This is input parameter to the method
	var output = data.NewVoidDefinition() // Expected output
	var input = p.getUserDefinition()
	var methodIdentifier = core.NewMethodIdentifier(interfaceIdentifier, "create")
	var methodDefinition = core.NewMethodDefinition(methodIdentifier, input, output, nil)
	return &methodDefinition
}

func (p *PetStoreApiInterface) getUserMethodDefinition() *core.MethodDefinition {
	var interfaceIdentifier = core.NewInterfaceIdentifier(p.interfaceName)
	//This is input parameter to the method
	var output = p.getUserDefinition() // Expected output
	var fields = make(map[string]data.DataDefinition)
	fields["userName"] = data.NewStringDefinition()
	var input = data.NewStructDefinition("User", fields)
	var methodIdentifier = core.NewMethodIdentifier(interfaceIdentifier, "get")
	var methodDefinition = core.NewMethodDefinition(methodIdentifier, input, output, nil)
	return &methodDefinition
}

func (p *PetStoreApiInterface) updateUserMethodDefinition() *core.MethodDefinition {
	var interfaceIdentifier = core.NewInterfaceIdentifier(p.interfaceName)
	//This is input parameter to the method
	var output = p.getUserDefinition() // Expected output
	var structFields = make(map[string]data.DataDefinition)
	structFields["userName"] = data.NewStringDefinition()
	structFields["updateduser"] = p.getUserDefinition()
	var input = data.NewStructDefinition("input", structFields)
	var methodIdentifier = core.NewMethodIdentifier(interfaceIdentifier, "update")
	var methodDefinition = core.NewMethodDefinition(methodIdentifier, input, output, nil)
	return &methodDefinition
}

func (p *PetStoreApiInterface) deleteUserMethodDefinition() *core.MethodDefinition {
	var interfaceIdentifier = core.NewInterfaceIdentifier(p.interfaceName)
	var output = data.NewVoidDefinition() // Expected output
	var structFields = make(map[string]data.DataDefinition)
	structFields["userName"] = data.NewStringDefinition()
	var input = data.NewStructDefinition("input", structFields)
	var methodIdentifier = core.NewMethodIdentifier(interfaceIdentifier, "delete")
	var methodDefinition = core.NewMethodDefinition(methodIdentifier, input, output, nil)
	return &methodDefinition
}

/**
 * Convert StructValue to User
 */
func (p *PetStoreApiInterface) getUserFromStructValue(input *data.StructValue) *User {
	var id, _ = input.Field("id")
	var idValue interface{}
	if reflect.TypeOf(id) == reflect.TypeOf((*data.StringValue)(nil)) {
		idValue = id.(*data.StringValue).Value()
	} else {
		idValue = id.(*data.IntegerValue).Value()
	}
	var userName, _ = input.String("userName")
	var firstName, _ = input.String("firstName")
	var lastName, _ = input.String("lastName")
	var email, _ = input.String("email")
	var password, _ = input.String("password")
	var phone, _ = input.String("phone")
	var userStatus, _ = input.Field("userStatus")
	var user = &User{id: idValue,
		userName:   userName,
		firstName:  firstName,
		lastName:   lastName,
		email:      email,
		password:   password,
		phone:      phone,
		userStatus: userStatus.(*data.IntegerValue).Value()}
	return user
}

/**
 * Convert User to StructValue
 */
func (p *PetStoreApiInterface) getStructValueFromUser(user *User) *data.StructValue {
	var structValue = data.NewStructValue("User", nil)
	if userIdString, ok := user.id.(string); ok {
		structValue.SetField("id", data.NewStringValue(userIdString))
	} else {
		structValue.SetField("id", data.NewIntegerValue(user.id.(int64)))
	}
	structValue.SetField("userName", data.NewStringValue(user.userName))
	structValue.SetField("firstName", data.NewStringValue(user.firstName))
	structValue.SetField("lastName", data.NewStringValue(user.lastName))
	structValue.SetField("email", data.NewStringValue(user.email))
	structValue.SetField("password", data.NewStringValue(user.password))
	structValue.SetField("phone", data.NewStringValue(user.phone))
	structValue.SetField("userStatus", data.NewIntegerValue(user.userStatus))
	return structValue
}

func (p *PetStoreApiInterface) create(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var structValue = input.(*data.StructValue)
	var user = p.getUserFromStructValue(structValue)
	p.userMap[user.userName] = user
	return core.NewMethodResult(data.NewVoidValue(), nil)
}
func (p *PetStoreApiInterface) update(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var structValue = input.(*data.StructValue)
	var userName, _ = structValue.String("userName")
	var user, _ = structValue.Field("updateduser")
	var userStructValue = user.(*data.StructValue)
	var updatedUser = p.getUserFromStructValue(userStructValue)
	var updatedUserResponse = p.getStructValueFromUser(updatedUser)
	p.userMap[userName] = updatedUser
	return core.NewMethodResult(updatedUserResponse, nil)
}
func (p *PetStoreApiInterface) get(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var structValue = input.(*data.StructValue)
	var userName, _ = structValue.String("userName")
	var user = p.userMap[userName]
	var result = p.getStructValueFromUser(user)
	return core.NewMethodResult(result, nil)
}
func (p *PetStoreApiInterface) delete(ctx *core.ExecutionContext, methodId core.MethodIdentifier, input data.DataValue) core.MethodResult {
	var structValue = input.(*data.StructValue)
	var userName, _ = structValue.String("userName")
	delete(p.userMap, userName)
	return core.NewMethodResult(data.NewVoidValue(), nil)

}
