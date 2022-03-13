package vuser

import (
	"github.com/vuuvv/orca/server"
	"strings"
)

func InAncestorOrg(token *server.AccessToken, orgPath string) bool {
	if token.IsSuper() {
		return true
	}
	return IsAncestorOrg(token.OrgPath, orgPath)
}

func IsAncestorOrg(ancestor string, descendant string) bool {
	return strings.HasPrefix(descendant, ancestor)
}
