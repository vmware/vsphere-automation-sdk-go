package arch

import (
	"strings"
)

var prots = []Type{
	JSONRPC,
	REST,
}

func GetByName(name string) Type {
	for _, p := range prots {
		if strings.ToLower(p.String()) == strings.ToLower(name) {
			return p
		}
	}
	return nil
}
