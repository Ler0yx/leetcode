package defangingAnIpAddress

import "strings"

func defangIPaddr(address string) string {

	stringSlice := strings.Split(address, ".")
	return strings.Join(stringSlice, "[.]")
}
