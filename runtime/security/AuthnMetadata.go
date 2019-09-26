package security

// represents deserialized authentication metadata from authn.ini
type AuthenticationMetadata struct {
	Authentication Authentication
}

type Authentication struct {
	Component Component
}

type Component struct {
	Name       string
	Schemes    map[string]Scheme
	Packages   map[string]interface{} //value can be string or array (for multiple schemes)
	Services   map[string]interface{} //value can be string or array (for multiple schemes)
	Operations map[string]interface{} //value can be string or array (for multiple schemes)
}

type Scheme struct {
	Type                 string
	AuthenticationScheme string
}
