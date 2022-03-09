package vuser

import "strings"

func IsAncestorOrg(ancestor string, descendant string) bool {
	return strings.HasPrefix(descendant, ancestor)
}
