package args

func InitProperties() Properties {
	return &properties{OptionsMap: make(map[string]*propValue)}
}
