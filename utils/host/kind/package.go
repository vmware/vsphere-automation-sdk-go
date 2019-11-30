package kind

import (
	"strings"
)

func GetHostTypeByName(name string) Info {
	for _, host := range All {
		if strings.ToLower(host.String()) == strings.ToLower(name) {
			return host
		}
	}
	return nil
}

func GetAllHostKind() []string {
	var allHostKind []string
	for _, hostKind := range All {
		allHostKind = append(allHostKind, hostKind.String())
	}
	return allHostKind
}
